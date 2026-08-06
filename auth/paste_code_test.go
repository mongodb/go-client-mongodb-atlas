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
	"crypto/hkdf"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"testing"
)

// Vector produced by the helper page's WebCrypto implementation
// (HKDF-SHA256, empty salt, info "atlas-cli-paste-v1").
const (
	vectorState = "kXx9uP4v2mJq8LwZ0aBcDeFgHiJkLmN_r7sT1uVwXyZ"
	vectorCode  = "yJ2mo9wLh4jBjrqctwbGP3Y6d8kQfz1N.a-Ok7aCode"
	vectorPaste = "TxxEQh30Wmf-zwfLNiHZ-BoG0KkYONsxmWHFb6QuvQ8yUv9olw0Zc4sn6Qoj46"
)

// encodePaste mirrors the helper page's encoding, for round-trip tests.
func encodePaste(t *testing.T, state, code string) string {
	t.Helper()
	keystream, err := hkdf.Key(sha256.New, []byte(state), nil, pasteInfo, len(code))
	if err != nil {
		t.Fatal(err)
	}
	cipher := make([]byte, len(code))
	for i := range cipher {
		cipher[i] = code[i] ^ keystream[i]
	}
	return base64.RawURLEncoding.EncodeToString(cipher) + codeChecksum([]byte(code))
}

func TestParsePastedCallback(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		state         string
		expectedCode  string
		expectedError string
	}{
		{
			name:         "helper page vector",
			input:        vectorPaste,
			state:        vectorState,
			expectedCode: vectorCode,
		},
		{
			name:         "full URL fallback",
			input:        "https://example.cloudfront.net/?code=abc123&state=some-state",
			state:        "some-state",
			expectedCode: "abc123",
		},
		{
			name:          "URL fallback with wrong state",
			input:         "https://example.cloudfront.net/?code=abc123&state=other",
			state:         "some-state",
			expectedError: "state mismatch",
		},
		{
			name:          "URL fallback with AS error",
			input:         "https://example.cloudfront.net/?error=access_denied&state=some-state",
			state:         "some-state",
			expectedError: "access_denied",
		},
		{
			name:          "typo in paste",
			input:         vectorPaste[:10] + "X" + vectorPaste[11:],
			state:         vectorState,
			expectedError: errPasteMalformed.Error(),
		},
		{
			name:          "paste from a different session",
			input:         vectorPaste,
			state:         "a-different-state-value",
			expectedError: errPasteMalformed.Error(),
		},
		{
			name:          "paste too short",
			input:         "abc",
			state:         vectorState,
			expectedError: errPasteMalformed.Error(),
		},
		{
			name:          "paste is not base64url",
			input:         "not!valid!base64!urlZZZZ",
			state:         vectorState,
			expectedError: errPasteMalformed.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, err := ParsePastedCallback(strings.NewReader(tt.input+"\n"), tt.state)
			if tt.expectedError != "" {
				if err == nil || !strings.Contains(err.Error(), tt.expectedError) {
					t.Fatalf("error = %v, want containing %q", err, tt.expectedError)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if code != tt.expectedCode {
				t.Errorf("code = %q, want %q", code, tt.expectedCode)
			}
		})
	}
}

func TestParsePastedCallbackRoundTrip(t *testing.T) {
	state, err := GenerateState()
	if err != nil {
		t.Fatal(err)
	}
	const code = "yJ2mo9wLh4jBjrqctwbGP3Y6d8kQfz1N"

	paste := encodePaste(t, state, code)
	got, err := ParsePastedCallback(strings.NewReader(paste+"\n"), state)
	if err != nil {
		t.Fatal(err)
	}
	if got != code {
		t.Errorf("code = %q, want %q", got, code)
	}
}
