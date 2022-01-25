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

// DeviceCode holds information about the authorization-in-progress.
type DeviceCode struct {
	UserCode        string `json:"user_code"`        // UserCode is the code presented to users
	VerificationURI string `json:"verification_uri"` // VerificationURI is the URI where users will need to confirm the code
	DeviceCode      string `json:"device_code"`      // DeviceCode is the internal code to confirm the status of the flow
	ExpiresIn       int    `json:"expires_in"`       // ExpiresIn when the code will expire
	Interval        int    `json:"interval"`         // Interval who often to verify the status of the code

	timeNow   func() time.Time
	timeSleep func(time.Duration)
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
	var r *Token
	resp, err2 := c.Do(ctx, req, &r)
	if err2 != nil {
		return nil, resp, err2
	}
	return r, resp, err2
}

// ErrTimeout is thrown when polling the server for the granted token has timed out.
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
		token, resp, err2 := c.GetToken(ctx, code.DeviceCode)
		var target *atlas.ErrorResponse
		if errors.As(err2, &target) && target.ErrorCode == "DEVICE_AUTHORIZATION_PENDING" {
			continue
		}
		if err2 != nil {
			return nil, resp, err2
		}

		if timeNow().After(expiresAt) {
			return nil, nil, ErrTimeout
		}
		return token, resp, nil
	}
}
