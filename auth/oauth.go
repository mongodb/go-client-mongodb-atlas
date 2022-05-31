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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const defaultBaseURL = atlas.CloudURL

var (
	userAgent = fmt.Sprintf("go-mongodbatlas/%s (%s;%s)", atlas.Version, runtime.GOOS, runtime.GOARCH)
)

type Config struct {
	client    *http.Client
	ClientID  string
	AuthURL   *url.URL
	UserAgent string
	Scopes    []string

	// copy raw server response to the Response struct
	withRaw bool
}

type ConfigOpt func(*Config) error

func NewConfig(httpClient *http.Client) *Config {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Config{
		client:    httpClient,
		AuthURL:   baseURL,
		UserAgent: userAgent,
	}
	return c
}

func NewConfigWithOptions(httpClient *http.Client, opts ...ConfigOpt) (*Config, error) {
	c := NewConfig(httpClient)
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// SetAuthURL is a config option for setting the base URL.
func SetAuthURL(bu string) ConfigOpt {
	return func(c *Config) error {
		u, err := url.Parse(bu)
		if err != nil {
			return err
		}

		c.AuthURL = u
		return nil
	}
}

// SetWithRaw is a client option for getting raw atlas server response within Response structure.
func SetWithRaw() ConfigOpt {
	return func(c *Config) error {
		c.withRaw = true
		return nil
	}
}

// SetUserAgent is a config option for setting the user agent.
func SetUserAgent(ua string) ConfigOpt {
	return func(c *Config) error {
		c.UserAgent = fmt.Sprintf("%s %s", ua, userAgent)
		return nil
	}
}

// SetClientID is a config option for setting the ClientID.
func SetClientID(clientID string) ConfigOpt {
	return func(c *Config) error {
		c.ClientID = clientID
		return nil
	}
}

// SetScopes is a config option for setting the Scopes.
func SetScopes(scopes []string) ConfigOpt {
	return func(c *Config) error {
		c.Scopes = scopes
		return nil
	}
}

// A TokenSource is anything that can return a token.
type TokenSource interface {
	// Token returns a token or an error.
	// Token must be safe for concurrent use by multiple goroutines.
	// The returned Token must not be modified.
	Token() (*Token, error)
}

func (c *Config) Do(ctx context.Context, req *http.Request, v interface{}) (*atlas.Response, error) {
	resp, err := atlas.DoRequestWithClient(ctx, c.client, req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer resp.Body.Close()

	r := &atlas.Response{Response: resp}
	body := resp.Body

	if c.withRaw {
		raw := new(bytes.Buffer)
		_, err = io.Copy(raw, body)
		if err != nil {
			return r, err
		}

		r.Raw = raw.Bytes()
		body = io.NopCloser(raw)
	}

	if err2 := r.CheckResponse(body); err2 != nil {
		return r, err2
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, body)
			if err != nil {
				return nil, err
			}
		} else {
			decErr := json.NewDecoder(body).Decode(v)
			if errors.Is(decErr, io.EOF) {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}
	return r, err
}

func (c *Config) NewRequest(ctx context.Context, method, urlStr string, v url.Values) (*http.Request, error) {
	if !strings.HasSuffix(c.AuthURL.Path, "/") {
		return nil, fmt.Errorf("base URL must have a trailing slash, but %q does not", c.AuthURL)
	}
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.AuthURL.ResolveReference(rel)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}
