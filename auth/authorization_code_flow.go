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
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	atlas "go.mongodb.org/atlas/mongodbatlas"
)

// PKCEParams holds the PKCE code verifier and challenge for an authorization request.
type PKCEParams struct {
	CodeVerifier  string
	CodeChallenge string
}

// GeneratePKCE creates a PKCE code verifier and its S256 challenge.
func GeneratePKCE() (*PKCEParams, error) {
	// 32 random bytes -> 43 base64url characters
	verifierBytes := make([]byte, 32)
	if _, err := rand.Read(verifierBytes); err != nil {
		return nil, fmt.Errorf("failed to generate code verifier: %w", err)
	}
	verifier := base64.RawURLEncoding.EncodeToString(verifierBytes)

	hash := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(hash[:])

	return &PKCEParams{
		CodeVerifier:  verifier,
		CodeChallenge: challenge,
	}, nil
}

// GenerateState creates a random state value for CSRF protection.
func GenerateState() (string, error) {
	stateBytes := make([]byte, 32)
	if _, err := rand.Read(stateBytes); err != nil {
		return "", fmt.Errorf("failed to generate state: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(stateBytes), nil
}

// AuthorizationURL builds the authorization endpoint URL using the ClientID from the Config.
func (c *Config) AuthorizationURL(authorizationEndpoint, redirectURI, state string, pkce *PKCEParams) (string, error) {
	u, err := url.Parse(authorizationEndpoint)
	if err != nil {
		return "", fmt.Errorf("invalid authorization endpoint: %w", err)
	}

	q := u.Query()
	q.Set("response_type", "code")
	q.Set("client_id", c.ClientID)
	q.Set("redirect_uri", redirectURI)
	q.Set("scope", "atlas")
	q.Set("state", state)
	q.Set("code_challenge", pkce.CodeChallenge)
	q.Set("code_challenge_method", "S256")
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// ExchangeCode exchanges an authorization code for tokens at the token endpoint.
func (c *Config) ExchangeCode(ctx context.Context, tokenEndpoint, code, redirectURI, codeVerifier string) (*Token, error) {
	v := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {redirectURI},
		"client_id":     {c.ClientID},
		"code_verifier": {codeVerifier},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenEndpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp atlas.ErrorResponse
		if decErr := json.NewDecoder(resp.Body).Decode(&errResp); decErr == nil && errResp.Detail != "" {
			return nil, fmt.Errorf("token exchange failed (HTTP %d): %s", resp.StatusCode, errResp.Detail)
		}
		return nil, fmt.Errorf("token exchange failed with HTTP %d", resp.StatusCode)
	}

	var token Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}

	if token.ExpiresIn > 0 {
		token.Expiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}

	return &token, nil
}

// RefreshAccessToken exchanges a refresh token for a new access token at the
// given token endpoint, using the ClientID from the Config.
func (c *Config) RefreshAccessToken(ctx context.Context, tokenEndpoint, refreshToken string) (*Token, error) {
	v := url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {c.ClientID},
		"refresh_token": {refreshToken},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenEndpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("refresh request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var oauthErr struct {
			Error       string `json:"error"`
			Description string `json:"error_description"`
		}
		if decErr := json.NewDecoder(resp.Body).Decode(&oauthErr); decErr == nil && oauthErr.Error != "" {
			if oauthErr.Description != "" {
				return nil, fmt.Errorf("token refresh failed (HTTP %d): %s: %s", resp.StatusCode, oauthErr.Error, oauthErr.Description)
			}
			return nil, fmt.Errorf("token refresh failed (HTTP %d): %s", resp.StatusCode, oauthErr.Error)
		}
		return nil, fmt.Errorf("token refresh failed with HTTP %d", resp.StatusCode)
	}

	var token Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, fmt.Errorf("failed to parse refresh response: %w", err)
	}

	if token.ExpiresIn > 0 {
		token.Expiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}

	return &token, nil
}

// RevokeToken revokes a token at the given revocation endpoint per RFC 7009,
// using the ClientID from the Config.
func (c *Config) RevokeAuthServerToken(ctx context.Context, revocationEndpoint, token, tokenTypeHint string) error {
	v := url.Values{
		"client_id":       {c.ClientID},
		"token":           {token},
		"token_type_hint": {tokenTypeHint},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, revocationEndpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create revocation request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("revocation request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var oauthErr struct {
			Error       string `json:"error"`
			Description string `json:"error_description"`
		}
		if decErr := json.NewDecoder(resp.Body).Decode(&oauthErr); decErr == nil && oauthErr.Error != "" {
			if oauthErr.Description != "" {
				return fmt.Errorf("token revocation failed (HTTP %d): %s: %s", resp.StatusCode, oauthErr.Error, oauthErr.Description)
			}
			return fmt.Errorf("token revocation failed (HTTP %d): %s", resp.StatusCode, oauthErr.Error)
		}
		return fmt.Errorf("token revocation failed with HTTP %d", resp.StatusCode)
	}

	return nil
}
