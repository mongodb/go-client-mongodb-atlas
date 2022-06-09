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
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const authExpiredError = "DEVICE_AUTHORIZATION_EXPIRED"

// DeviceCode holds information about the authorization-in-progress.
type DeviceCode struct {
	UserCode        string `json:"user_code"`        // UserCode is the code presented to users
	VerificationURI string `json:"verification_uri"` // VerificationURI is the URI where users will need to confirm the code
	DeviceCode      string `json:"device_code"`      // DeviceCode is the internal code to confirm the status of the flow
	ExpiresIn       int    `json:"expires_in"`       // ExpiresIn when the code will expire
	Interval        int    `json:"interval"`         // Interval how often to verify the status of the code

	timeNow   func() time.Time
	timeSleep func(time.Duration)
}

type RegistrationConfig struct {
	RegistrationURL string `json:"registrationUrl"`
}

const deviceBasePath = "api/private/unauth/account/device"

// RequestCode initiates the authorization flow by requesting a code.
func (c Config) RequestCode(ctx context.Context) (*DeviceCode, *atlas.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, deviceBasePath+"/authorize",
		url.Values{
			"client_id": {c.ClientID},
			"scope":     {strings.Join(c.Scopes, " ")},
		},
	)
	if err != nil {
		return nil, nil, err
	}
	var r *DeviceCode
	resp, err2 := c.Do(ctx, req, &r)
	return r, resp, err2
}

// GetToken gets a device token.
func (c Config) GetToken(ctx context.Context, deviceCode string) (*Token, *atlas.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, deviceBasePath+"/token",
		url.Values{
			"client_id":   {c.ClientID},
			"device_code": {deviceCode},
			"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
		},
	)
	if err != nil {
		return nil, nil, err
	}
	var t *Token
	resp, err2 := c.Do(ctx, req, &t)
	if err2 != nil {
		return nil, resp, err2
	}
	return t, resp, err2
}

// ErrTimeout is returned when polling the server for the granted token has timed out.
var ErrTimeout = errors.New("authentication timed out")

// PollToken polls the server until an access token is granted or denied.
func (c Config) PollToken(ctx context.Context, code *DeviceCode) (*Token, *atlas.Response, error) {
	timeNow := code.timeNow
	if timeNow == nil {
		timeNow = time.Now
	}
	timeSleep := code.timeSleep
	if timeSleep == nil {
		timeSleep = time.Sleep
	}

	checkInterval := time.Duration(code.Interval) * time.Second
	expiresAt := timeNow().Add(time.Duration(code.ExpiresIn) * time.Second)

	for {
		timeSleep(checkInterval)
		token, resp, err := c.GetToken(ctx, code.DeviceCode)
		var target *atlas.ErrorResponse
		if errors.As(err, &target) && target.ErrorCode == "DEVICE_AUTHORIZATION_PENDING" {
			continue
		}
		if err != nil {
			return nil, resp, err
		}

		if timeNow().After(expiresAt) {
			return nil, nil, ErrTimeout
		}
		return token, resp, nil
	}
}

// RefreshToken takes a refresh token and gets a new access token.
func (c Config) RefreshToken(ctx context.Context, token string) (*Token, *atlas.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, deviceBasePath+"/token",
		url.Values{
			"client_id":     {c.ClientID},
			"refresh_token": {token},
			"scope":         {strings.Join(c.Scopes, " ")},
			"grant_type":    {"refresh_token"},
		},
	)
	if err != nil {
		return nil, nil, err
	}
	var t *Token
	resp, err2 := c.Do(ctx, req, &t)
	if err2 != nil {
		return nil, resp, err2
	}
	return t, resp, err2
}

// RevokeToken takes an access or refresh token and revokes it.
func (c Config) RevokeToken(ctx context.Context, token, tokenTypeHint string) (*atlas.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, deviceBasePath+"/revoke",
		url.Values{
			"client_id":       {c.ClientID},
			"token":           {token},
			"token_type_hint": {tokenTypeHint},
		},
	)
	if err != nil {
		return nil, err
	}

	return c.Do(ctx, req, nil)
}

// RegistrationConfig retrieves the config used for registration.
func (c Config) RegistrationConfig(ctx context.Context) (*RegistrationConfig, *atlas.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, deviceBasePath+"/registration", url.Values{})
	if err != nil {
		return nil, nil, err
	}
	var rc *RegistrationConfig
	resp, err := c.Do(ctx, req, &rc)
	if err != nil {
		return nil, resp, err
	}
	return rc, resp, err
}

func IsTimeoutErr(err error) bool {
	var target *atlas.ErrorResponse
	return errors.Is(err, ErrTimeout) || (errors.As(err, &target) && target.ErrorCode == authExpiredError)
}
