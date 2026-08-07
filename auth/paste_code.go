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
	"errors"
	"fmt"
	"io"
	"strings"
)

// Wire format shared with the hosted no-browser callback page: the page
// derives a keystream from the state via HKDF-SHA256 (empty salt, versioned
// info string), XORs it with the authorization code, and appends a short
// checksum over the code so typos and cross-session pastes fail locally
// instead of at the token endpoint. The state must be fresh CSPRNG output,
// unique per flow: the encoding leans on it never being reused.
const (
	// URL of the hosted helper page for --noBrowser interaction.
	NoBrowserRedirectURI = "https://dvtm994tafpir.cloudfront.net/"

	pasteInfo          = "atlas-cli-paste-v1"
	pasteChecksumChars = 4
)

var errPasteMalformed = errors.New("the pasted code doesn't match this connection attempt; check for missing or mistyped characters and try again")

// ParsePastedCallback reads a single pasted token from r and extracts the
// authorization code from it. It accepts either the compact code shown by
// the callback helper page or, as a fallback, the full redirect URL.
func ParsePastedCallback(r io.Reader, expectedState string) (string, error) {
	var raw string
	if _, err := fmt.Fscanln(r, &raw); err != nil {
		return "", fmt.Errorf("failed to read pasted value: %w", err)
	}
	raw = strings.TrimSpace(raw)

	if strings.HasPrefix(raw, NoBrowserRedirectURI) {
		return parseRedirectURL(raw, expectedState)
	}
	return decodePastedCode(raw, expectedState)
}

// decodePastedCode reverses the helper page's encoding and verifies the
// embedded checksum before returning the authorization code.
func decodePastedCode(paste, state string) (string, error) {
	if len(paste) <= pasteChecksumChars {
		return "", errPasteMalformed
	}
	sum := paste[len(paste)-pasteChecksumChars:]
	cipher, err := base64.RawURLEncoding.DecodeString(paste[:len(paste)-pasteChecksumChars])
	if err != nil {
		return "", errPasteMalformed
	}

	keystream, err := hkdf.Key(sha256.New, []byte(state), nil, pasteInfo, len(cipher))
	if err != nil {
		return "", fmt.Errorf("failed to derive paste keystream: %w", err)
	}
	code := make([]byte, len(cipher))
	for i, b := range cipher {
		code[i] = b ^ keystream[i]
	}

	if codeChecksum(code) != sum {
		return "", errPasteMalformed
	}
	return string(code), nil
}

func codeChecksum(code []byte) string {
	digest := sha256.Sum256(code)
	return base64.RawURLEncoding.EncodeToString(digest[:])[:pasteChecksumChars]
}
