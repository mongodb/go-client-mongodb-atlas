// Copyright 2022 MongoDB Inc
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
	"testing"

	"github.com/go-test/deep"
)

func TestConfig_RequestCode(t *testing.T) {
	config, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/private/unauth/account/device/authorize", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		fmt.Fprintf(w, `{
		  "user_code": "QW3PYV7R",
		  "verification_uri": "%s/account/connect",
		  "device_code": "61eef18e310968047ff5e02a",
		  "expires_in": 600,
		  "interval": 10
		}`, baseURLPath)
	})

	results, _, err := config.RequestCode(ctx)
	if err != nil {
		t.Fatalf("RequestCode returned error: %v", err)
	}

	expected := &DeviceCode{
		UserCode:        "QW3PYV7R",
		VerificationURI: baseURLPath + "/account/connect",
		DeviceCode:      "61eef18e310968047ff5e02a",
		ExpiresIn:       600,
		Interval:        10,
	}

	if diff := deep.Equal(results, expected); diff != nil {
		t.Error(diff)
	}
}

func TestConfig_GetToken(t *testing.T) {
	config, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/private/unauth/account/device/token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		fmt.Fprint(w, `{
		  "access_token": "secret1",
		  "refresh_token": "secret2",
		  "scope": "openid",
		  "id_token": "idtoken",
		  "token_type": "Bearer",
		  "expires_in": 3600
		}`)
	})
	code := &DeviceCode{
		DeviceCode: "61eef18e310968047ff5e02a",
		ExpiresIn:  600,
		Interval:   10,
	}
	results, _, err := config.GetToken(ctx, code.DeviceCode)
	if err != nil {
		t.Fatalf("GetToken returned error: %v", err)
	}

	expected := &Token{
		AccessToken:  "secret1",
		RefreshToken: "secret2",
		Scope:        "openid",
		IDToken:      "idtoken",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
	}

	if diff := deep.Equal(results, expected); diff != nil {
		t.Error(diff)
	}
}

func TestConfig_RefreshToken(t *testing.T) {
	config, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/private/unauth/account/device/token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		fmt.Fprint(w, `{
		  "access_token": "secret1",
		  "refresh_token": "secret2",
		  "scope": "openid",
		  "id_token": "idtoken",
		  "token_type": "Bearer",
		  "expires_in": 3600
		}`)
	})

	results, _, err := config.RefreshToken(ctx, "secret2")
	if err != nil {
		t.Fatalf("RefreshToken returned error: %v", err)
	}

	expected := &Token{
		AccessToken:  "secret1",
		RefreshToken: "secret2",
		Scope:        "openid",
		IDToken:      "idtoken",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
	}

	if diff := deep.Equal(results, expected); diff != nil {
		t.Error(diff)
	}
}

func TestConfig_PollToken(t *testing.T) {
	config, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/private/unauth/account/device/token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		fmt.Fprint(w, `{
		  "access_token": "secret1",
		  "refresh_token": "secret2",
		  "scope": "openid",
		  "id_token": "idtoken",
		  "token_type": "Bearer",
		  "expires_in": 3600
		}`)
	})
	code := &DeviceCode{
		DeviceCode: "61eef18e310968047ff5e02a",
		ExpiresIn:  600,
		Interval:   10,
	}
	results, _, err := config.PollToken(ctx, code)
	if err != nil {
		t.Fatalf("PollToken returned error: %v", err)
	}

	expected := &Token{
		AccessToken:  "secret1",
		RefreshToken: "secret2",
		Scope:        "openid",
		IDToken:      "idtoken",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
	}

	if diff := deep.Equal(results, expected); diff != nil {
		t.Error(diff)
	}
}

func TestConfig_RevokeToken(t *testing.T) {
	config, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/private/unauth/account/device/revoke", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
	})

	_, err := config.RevokeToken(ctx, "a", "refresh_token")
	if err != nil {
		t.Fatalf("RequestCode returned error: %v", err)
	}
}

func TestConfig_RegistrationConfig(t *testing.T) {
	config, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/private/unauth/account/device/registration", func(w http.ResponseWriter, r *http.Request) {
		if http.MethodGet != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, http.MethodGet)
		}

		fmt.Fprint(w, `{
		  "registrationUrl": "http://localhost:8080/account/register/cli"
		}`)
	})

	results, _, err := config.RegistrationConfig(ctx)
	if err != nil {
		t.Fatalf("RegistrationConfig returned error: %v", err)
	}

	expected := &RegistrationConfig{
		RegistrationURL: "http://localhost:8080/account/register/cli",
	}

	if diff := deep.Equal(results, expected); diff != nil {
		t.Error(diff)
	}
}
