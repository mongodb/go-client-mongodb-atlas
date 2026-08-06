// Copyright 2026 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
)

const callbackPath = "/atlas-cli/callback"

// callbackPage is the browser-facing page for the loopback redirect.
// Detail beyond success/failure intentionally stays out of the browser; the
// CLI surfaces the specifics.
//
//go:embed callback_page.html
var callbackPage string

func writeCallbackPage(w http.ResponseWriter, status int, title, message string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	accent := "#00ed64"
	if status != http.StatusOK {
		accent = "#970606"
	}
	fmt.Fprintf(w, callbackPage, title, accent, message)
}

// NoBrowserRedirectURI returns the redirect URI for the manual paste flow,
// using the registered loopback URI without a port.
func NoBrowserRedirectURI() string {
	return "http://127.0.0.1" + callbackPath
}

// ParseCodeFromRedirectURL reads a pasted URL from r (pluggable for
// testing), extracts the authorization code, and validates the state
// parameter.
func ParseCodeFromRedirectURL(r io.Reader, expectedState string) (string, error) {
	var raw string
	if _, err := fmt.Fscanln(r, &raw); err != nil {
		return "", fmt.Errorf("failed to read URL: %w", err)
	}
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	query := u.Query()

	// Validate state first (CSRF protection — must precede surfacing any
	// AS response data, otherwise an unsolicited callback could inject
	// attacker-controlled content into the user-facing error message).
	if query.Get("state") != expectedState {
		return "", fmt.Errorf("state mismatch")
	}

	// OAuth error response from the AS (RFC 6749 §4.1.2.1).
	if query.Get("error") != "" {
		return "", fmt.Errorf("%s", formatOAuthError(query))
	}

	code := query.Get("code")
	if code == "" {
		return "", fmt.Errorf("no authorization code in URL")
	}

	return code, nil
}

// formatOAuthError builds the user-facing message for an RFC 6749 §4.1.2.1 /
// §5.2 error response, including error_description and error_uri when present.
func formatOAuthError(q url.Values) string {
	msg := q.Get("error")
	if desc := q.Get("error_description"); desc != "" {
		msg = fmt.Sprintf("%s: %s", msg, desc)
	}
	if uri := q.Get("error_uri"); uri != "" {
		msg = fmt.Sprintf("%s (see %s)", msg, uri)
	}
	return msg
}

// CallbackServer listens on the loopback address with an OS-assigned port
// to receive the authorization code redirect from the OAuth Authorization Server.
type CallbackServer struct {
	listener net.Listener
	server   *http.Server
	code     string
	state    string
	errMsg   string
	done     chan struct{}
}

// StartCallbackServer binds a listener on 127.0.0.1 with a dynamic port and
// registers a handler for the callback path. The server is ready to receive
// the redirect once this function returns.
func StartCallbackServer(expectedState string) (*CallbackServer, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("failed to start callback listener: %w", err)
	}

	cs := &CallbackServer{
		listener: listener,
		state:    expectedState,
		done:     make(chan struct{}),
	}

	mux := http.NewServeMux()
	mux.HandleFunc(callbackPath, cs.handleCallback)

	cs.server = &http.Server{
		Handler: mux,
	}

	go func() {
		_ = cs.server.Serve(listener)
	}()

	return cs, nil
}

// RedirectURI returns the full redirect URI including the dynamically assigned port,
// suitable for inclusion in the authorization request.
func (cs *CallbackServer) RedirectURI() string {
	return fmt.Sprintf("http://127.0.0.1:%d%s", cs.Port(), callbackPath)
}

// Port returns the dynamically assigned port the server is listening on.
func (cs *CallbackServer) Port() int {
	return cs.listener.Addr().(*net.TCPAddr).Port
}

// WaitForCallback blocks until the callback is received or the context is cancelled.
// Returns the authorization code on success.
func (cs *CallbackServer) WaitForCallback(ctx context.Context) (string, error) {
	select {
	case <-cs.done:
		if cs.errMsg != "" {
			return "", fmt.Errorf("authorization failed: %s", cs.errMsg)
		}
		return cs.code, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// Close shuts down the callback server.
func (cs *CallbackServer) Close() error {
	return cs.server.Close()
}

func (cs *CallbackServer) handleCallback(w http.ResponseWriter, r *http.Request) {
	defer close(cs.done)

	query := r.URL.Query()

	// Validate state first (CSRF protection — must precede surfacing any
	// AS response data, otherwise an unsolicited callback could inject
	// attacker-controlled content into the user-facing error message).
	if query.Get("state") != cs.state {
		cs.errMsg = "state mismatch"
		writeCallbackPage(w, http.StatusBadRequest, "Connection failed",
			"The response did not match this connection attempt. Return to the terminal and retry the connection.")
		return
	}

	// OAuth error response from the AS (RFC 6749 §4.1.2.1).
	if query.Get("error") != "" {
		cs.errMsg = formatOAuthError(query)
		writeCallbackPage(w, http.StatusBadRequest, "Connection failed",
			"The connection attempt failed. Return to the terminal for details.")
		return
	}

	code := query.Get("code")
	if code == "" {
		cs.errMsg = "no authorization code in response"
		writeCallbackPage(w, http.StatusBadRequest, "Connection failed",
			"No authorization code was received. Return to the terminal and retry the connection.")
		return
	}

	cs.code = code
	writeCallbackPage(w, http.StatusOK, "Connected to Atlas",
		"The Atlas CLI is now connected. Return to the terminal to continue.")
}
