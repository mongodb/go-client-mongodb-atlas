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
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestGeneratePKCE(t *testing.T) {
	pkce, err := GeneratePKCE()
	if err != nil {
		t.Fatalf("GeneratePKCE returned error: %v", err)
	}

	if len(pkce.CodeVerifier) != 43 {
		t.Errorf("verifier length = %d, expected 43", len(pkce.CodeVerifier))
	}
	if len(pkce.CodeChallenge) != 43 {
		t.Errorf("challenge length = %d, expected 43", len(pkce.CodeChallenge))
	}

	hash := sha256.Sum256([]byte(pkce.CodeVerifier))
	expected := base64.RawURLEncoding.EncodeToString(hash[:])
	if pkce.CodeChallenge != expected {
		t.Errorf("challenge does not match S256(verifier): got %q, expected %q", pkce.CodeChallenge, expected)
	}

	other, _ := GeneratePKCE()
	if other.CodeVerifier == pkce.CodeVerifier {
		t.Error("two calls produced the same verifier")
	}
}

func TestGenerateState(t *testing.T) {
	state, err := GenerateState()
	if err != nil {
		t.Fatalf("GenerateState returned error: %v", err)
	}
	if len(state) != 43 {
		t.Errorf("state length = %d, expected 43", len(state))
	}

	other, _ := GenerateState()
	if other == state {
		t.Error("two calls produced the same state")
	}
}

func TestConfig_AuthorizationURL(t *testing.T) {
	c := NewConfig(nil)
	c.ClientID = "test-client"

	pkce := &PKCEParams{CodeVerifier: "verifier", CodeChallenge: "challenge"}
	got, err := c.AuthorizationURL("https://as.example.com/authorize", "http://127.0.0.1:1234/cb", "state-123", pkce)
	if err != nil {
		t.Fatalf("AuthorizationURL returned error: %v", err)
	}

	u, err := url.Parse(got)
	if err != nil {
		t.Fatalf("returned URL did not parse: %v", err)
	}
	q := u.Query()
	expected := map[string]string{
		"response_type":         "code",
		"client_id":             "test-client",
		"redirect_uri":          "http://127.0.0.1:1234/cb",
		"scope":                 "atlas",
		"state":                 "state-123",
		"code_challenge":        "challenge",
		"code_challenge_method": "S256",
	}
	for k, v := range expected {
		if q.Get(k) != v {
			t.Errorf("query[%q] = %q, expected %q", k, q.Get(k), v)
		}
	}
}

func TestConfig_ExchangeCode(t *testing.T) {
	var gotForm url.Values
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, expected POST", r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Fatalf("ParseForm: %v", err)
		}
		gotForm = r.PostForm
		fmt.Fprint(w, `{"access_token":"at","refresh_token":"rt","token_type":"Bearer","expires_in":3600}`)
	}))
	defer server.Close()

	c := NewConfig(nil)
	c.ClientID = "test-client"

	before := time.Now()
	token, err := c.ExchangeCode(ctx, server.URL, "auth-code", "http://127.0.0.1/cb", "verifier")
	if err != nil {
		t.Fatalf("ExchangeCode returned error: %v", err)
	}

	if token.AccessToken != "at" || token.RefreshToken != "rt" {
		t.Errorf("token mismatch: %+v", token)
	}
	if token.Expiry.Before(before.Add(3590*time.Second)) || token.Expiry.After(before.Add(3610*time.Second)) {
		t.Errorf("Expiry not within expected range: %v", token.Expiry)
	}
	wantForm := map[string]string{
		"grant_type":    "authorization_code",
		"code":          "auth-code",
		"redirect_uri":  "http://127.0.0.1/cb",
		"client_id":     "test-client",
		"code_verifier": "verifier",
	}
	for k, v := range wantForm {
		if gotForm.Get(k) != v {
			t.Errorf("form[%q] = %q, expected %q", k, gotForm.Get(k), v)
		}
	}
}

func TestConfig_RefreshAccessToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var gotForm url.Values
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_ = r.ParseForm()
			gotForm = r.PostForm
			fmt.Fprint(w, `{"access_token":"new-at","refresh_token":"new-rt","token_type":"Bearer","expires_in":1800}`)
		}))
		defer server.Close()

		c := NewConfig(nil)
		c.ClientID = "test-client"
		token, err := c.RefreshAccessToken(ctx, server.URL, "old-rt")
		if err != nil {
			t.Fatalf("RefreshAccessToken returned error: %v", err)
		}
		if token.AccessToken != "new-at" {
			t.Errorf("AccessToken = %q, expected new-at", token.AccessToken)
		}
		if gotForm.Get("grant_type") != "refresh_token" || gotForm.Get("refresh_token") != "old-rt" || gotForm.Get("client_id") != "test-client" {
			t.Errorf("unexpected form: %v", gotForm)
		}
	})

	t.Run("oauth error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"error":"invalid_grant","error_description":"token revoked"}`)
		}))
		defer server.Close()

		c := NewConfig(nil)
		_, err := c.RefreshAccessToken(ctx, server.URL, "old-rt")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if got := err.Error(); !strings.Contains(got, "invalid_grant") || !strings.Contains(got, "token revoked") {
			t.Errorf("error %q missing error code or description", got)
		}
	})
}

func TestConfig_RevokeAuthServerToken(t *testing.T) {
	var gotForm url.Values
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		gotForm = r.PostForm
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := NewConfig(nil)
	c.ClientID = "test-client"
	if err := c.RevokeAuthServerToken(ctx, server.URL, "rt", "refresh_token"); err != nil {
		t.Fatalf("RevokeAuthServerToken returned error: %v", err)
	}
	wantForm := map[string]string{
		"client_id":       "test-client",
		"token":           "rt",
		"token_type_hint": "refresh_token",
	}
	for k, v := range wantForm {
		if gotForm.Get(k) != v {
			t.Errorf("form[%q] = %q, expected %q", k, gotForm.Get(k), v)
		}
	}
}

