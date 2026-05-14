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
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestCallbackServer_SuccessfulCallback(t *testing.T) {
	state := "test-state-123"
	cs, err := StartCallbackServer(state)
	if err != nil {
		t.Fatalf("failed to start callback server: %v", err)
	}
	defer cs.Close()

	// Simulate the AS redirecting the browser with a code and matching state
	go func() {
		url := fmt.Sprintf("http://127.0.0.1:%d%s?code=test-auth-code&state=%s", cs.Port(), callbackPath, state)
		resp, err := http.Get(url)
		if err != nil {
			t.Errorf("callback request failed: %v", err)
			return
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	code, err := cs.WaitForCallback(ctx)
	if err != nil {
		t.Fatalf("WaitForCallback returned error: %v", err)
	}
	if code != "test-auth-code" {
		t.Errorf("expected code 'test-auth-code', got '%s'", code)
	}
}

func TestCallbackServer_StateMismatch(t *testing.T) {
	cs, err := StartCallbackServer("expected-state")
	if err != nil {
		t.Fatalf("failed to start callback server: %v", err)
	}
	defer cs.Close()

	go func() {
		url := fmt.Sprintf("http://127.0.0.1:%d%s?code=test-auth-code&state=wrong-state", cs.Port(), callbackPath)
		resp, err := http.Get(url)
		if err != nil {
			t.Errorf("callback request failed: %v", err)
			return
		}
		resp.Body.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = cs.WaitForCallback(ctx)
	if err == nil {
		t.Fatal("expected error for state mismatch, got nil")
	}
}

func TestCallbackServer_ErrorResponse(t *testing.T) {
	cs, err := StartCallbackServer("test-state")
	if err != nil {
		t.Fatalf("failed to start callback server: %v", err)
	}
	defer cs.Close()

	go func() {
		url := fmt.Sprintf("http://127.0.0.1:%d%s?error=access_denied&error_description=user+denied+access", cs.Port(), callbackPath)
		resp, err := http.Get(url)
		if err != nil {
			t.Errorf("callback request failed: %v", err)
			return
		}
		resp.Body.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = cs.WaitForCallback(ctx)
	if err == nil {
		t.Fatal("expected error for error response, got nil")
	}
}

func TestCallbackServer_NoCode(t *testing.T) {
	state := "test-state"
	cs, err := StartCallbackServer(state)
	if err != nil {
		t.Fatalf("failed to start callback server: %v", err)
	}
	defer cs.Close()

	go func() {
		url := fmt.Sprintf("http://127.0.0.1:%d%s?state=%s", cs.Port(), callbackPath, state)
		resp, err := http.Get(url)
		if err != nil {
			t.Errorf("callback request failed: %v", err)
			return
		}
		resp.Body.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = cs.WaitForCallback(ctx)
	if err == nil {
		t.Fatal("expected error for missing code, got nil")
	}
}

func TestCallbackServer_ContextCancelled(t *testing.T) {
	cs, err := StartCallbackServer("test-state")
	if err != nil {
		t.Fatalf("failed to start callback server: %v", err)
	}
	defer cs.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err = cs.WaitForCallback(ctx)
	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
	}
}

func TestCallbackServer_DynamicPort(t *testing.T) {
	cs, err := StartCallbackServer("test-state")
	if err != nil {
		t.Fatalf("failed to start callback server: %v", err)
	}
	defer cs.Close()

	if cs.Port() == 0 {
		t.Error("expected non-zero port")
	}

	expected := fmt.Sprintf("http://127.0.0.1:%d/atlas-cli/callback", cs.Port())
	if cs.RedirectURI() != expected {
		t.Errorf("expected redirect URI '%s', got '%s'", expected, cs.RedirectURI())
	}
}

func TestNoBrowserRedirectURI(t *testing.T) {
	if got := NoBrowserRedirectURI(); got != "http://127.0.0.1/atlas-cli/callback" {
		t.Errorf("NoBrowserRedirectURI() = %q", got)
	}
}

func TestParseCodeFromRedirectURL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		input := strings.NewReader("http://127.0.0.1/atlas-cli/callback?code=auth-code&state=expected-state\n")
		code, err := ParseCodeFromRedirectURL(input, "expected-state")
		if err != nil {
			t.Fatalf("ParseCodeFromRedirectURL returned error: %v", err)
		}
		if code != "auth-code" {
			t.Errorf("code = %q, expected auth-code", code)
		}
	})

	t.Run("state mismatch", func(t *testing.T) {
		input := strings.NewReader("http://127.0.0.1/atlas-cli/callback?code=auth-code&state=wrong-state\n")
		_, err := ParseCodeFromRedirectURL(input, "expected-state")
		if err == nil {
			t.Fatal("expected error for state mismatch, got nil")
		}
		if !strings.Contains(err.Error(), "state") {
			t.Errorf("error %q does not mention state", err.Error())
		}
	})
}
