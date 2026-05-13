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
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

const callbackPath = "/atlas-cli/callback"

// NoBrowserRedirectURI returns the redirect URI for the manual paste flow,
// using the registered loopback URI without a port.
func NoBrowserRedirectURI() string {
	return "http://127.0.0.1" + callbackPath
}

// ParseCodeFromRedirectURL reads a pasted URL from stdin, extracts the
// authorization code, and validates the state parameter.
func ParseCodeFromRedirectURL(expectedState string) (string, error) {
	var raw string
	if _, err := fmt.Scanln(&raw); err != nil {
		return "", fmt.Errorf("failed to read URL: %w", err)
	}

	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	query := u.Query()

	if errCode := query.Get("error"); errCode != "" {
		desc := query.Get("error_description")
		if desc != "" {
			return "", fmt.Errorf("%s: %s", errCode, desc)
		}
		return "", fmt.Errorf("%s", errCode)
	}

	if query.Get("state") != expectedState {
		return "", fmt.Errorf("state mismatch")
	}

	code := query.Get("code")
	if code == "" {
		return "", fmt.Errorf("no authorization code in URL")
	}

	return code, nil
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

	// OAuth error response from the AS (RFC 6749 §4.1.2.1) — surface before validating state so the upstream failure reaches the user.
	if errCode := query.Get("error"); errCode != "" {
		cs.errMsg = errCode
		if desc := query.Get("error_description"); desc != "" {
			cs.errMsg = fmt.Sprintf("%s: %s", errCode, desc)
		}
		http.Error(w, "Authorization failed. You may close this window.", http.StatusBadRequest)
		return
	}

	// Validate state to prevent CSRF
	if query.Get("state") != cs.state {
		cs.errMsg = "state mismatch"
		http.Error(w, "Authorization failed: state mismatch. You may close this window.", http.StatusBadRequest)
		return
	}

	code := query.Get("code")
	if code == "" {
		cs.errMsg = "no authorization code in response"
		http.Error(w, "Authorization failed: no code received. You may close this window.", http.StatusBadRequest)
		return
	}

	cs.code = code
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<html><body><p>Authorization successful. You may close this window.</p></body></html>")
}
