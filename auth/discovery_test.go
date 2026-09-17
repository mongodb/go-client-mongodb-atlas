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
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestConfig_DiscoverAuthServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/.well-known/oauth-authorization-server" {
			t.Errorf("path = %s, expected the well-known endpoint", r.URL.Path)
		}
		w.Header().Set("Cache-Control", "max-age=3600")
		fmt.Fprint(w, `{
			"issuer": "https://as.example.com",
			"authorization_endpoint": "https://as.example.com/authorize",
			"token_endpoint": "https://as.example.com/token",
			"revocation_endpoint": "https://as.example.com/revoke"
		}`)
	}))
	defer server.Close()

	c := NewConfig(nil)
	c.AuthServerURL, _ = url.Parse(server.URL)

	before := time.Now()
	result, err := c.DiscoverAuthServer(ctx)
	if err != nil {
		t.Fatalf("DiscoverAuthServer returned error: %v", err)
	}

	metadata, ok := result["metadata"].(map[string]any)
	if !ok {
		t.Fatalf(`result["metadata"] is not a map: %T`, result["metadata"])
	}
	if metadata["issuer"] != "https://as.example.com" {
		t.Errorf("issuer = %v, expected https://as.example.com", metadata["issuer"])
	}

	expiryStr, ok := result["expiry"].(string)
	if !ok {
		t.Fatalf(`result["expiry"] is not a string: %T`, result["expiry"])
	}
	expiry, err := time.Parse(time.RFC3339, expiryStr)
	if err != nil {
		t.Fatalf("expiry %q not RFC 3339: %v", expiryStr, err)
	}
	if expiry.Before(before.Add(3590*time.Second)) || expiry.After(before.Add(3610*time.Second)) {
		t.Errorf("expiry %v outside expected window from Cache-Control max-age=3600", expiry)
	}
}
