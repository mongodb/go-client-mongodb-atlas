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
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"

	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	// baseURLPath is a non-empty Client.BaseURL path to use during tests,
	// to ensure relative URLs are used for all endpoints.
	baseURLPath = "/api-v1"
)

var (
	ctx = context.TODO()
)

// setup sets up a test HTTP server along with an auth.Config that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (config *Config, mux *http.ServeMux, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// We want to ensure that tests catch mistakes where the endpoint URL is
	// specified as absolute rather than relative. It only makes a difference
	// when there's a non-empty base URL path. So, use that.
	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "FAIL: Client.BaseURL path prefix is not preserved in the request URL:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\tDid you accidentally use an absolute endpoint URL rather than relative?")
		http.Error(w, "Client.BaseURL path prefix is not preserved in the request URL.", http.StatusInternalServerError)
	})

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the Atlas client being tested and is
	// configured to use test server.
	config = NewConfig(nil)
	u, _ := url.Parse(server.URL + baseURLPath + "/")
	config.AuthURL = u

	return config, mux, server.Close
}

func testMethod(t *testing.T, r *http.Request) {
	t.Helper()
	if http.MethodPost != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, http.MethodPost)
	}
}

func testClientDefaultBaseURL(t *testing.T, c *Config) {
	t.Helper()
	if c.AuthURL == nil || c.AuthURL.String() != defaultBaseURL {
		t.Errorf("NewConfig BaseURL = %v, expected %v", c.AuthURL, defaultBaseURL)
	}
}

func testClientDefaultUserAgent(t *testing.T, c *Config) {
	t.Helper()
	if c.UserAgent != userAgent {
		t.Errorf("NewConfig UserAgent = %v, expected %v", c.UserAgent, userAgent)
	}
}

func testClientDefaults(t *testing.T, c *Config) {
	t.Helper()
	testClientDefaultBaseURL(t, c)
	testClientDefaultUserAgent(t, c)
}

func TestNewConfig(t *testing.T) {
	c := NewConfig(nil)
	testClientDefaults(t, c)
}

func TestNewConfigWithOptions(t *testing.T) {
	c, err := NewConfigWithOptions(nil)

	if err != nil {
		t.Fatalf("NewConfigWithOptions(): %v", err)
	}
	testClientDefaults(t, c)
}

func TestNewRequest_withCustomUserAgent(t *testing.T) {
	ua := fmt.Sprintf("testing/%s", atlas.Version)
	c, err := NewConfigWithOptions(nil, SetUserAgent(ua))

	if err != nil {
		t.Fatalf("NewConfigWithOptions() unexpected error: %v", err)
	}

	req, _ := c.NewRequest(ctx, http.MethodGet, "/foo", nil)

	expected := fmt.Sprintf("%s %s", ua, userAgent)
	if got := req.Header.Get("User-Agent"); got != expected {
		t.Errorf("NewConfigWithOptions() UserAgent = %s; expected %s", got, expected)
	}
}

func TestNewPlainRequest(t *testing.T) {
	c := NewConfig(nil)

	requestPath := "foo"

	inURL, outURL := requestPath, defaultBaseURL+requestPath
	req, _ := c.NewRequest(ctx, http.MethodGet, inURL, nil)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewPlainRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test accept content type is correct
	accept := req.Header.Get("Accept")
	if accept != "application/json" {
		t.Errorf("NewPlainRequest() Accept = %v, expected %v", accept, "application/json")
	}
	contentType := req.Header.Get("Content-Type")
	if contentType != "application/x-www-form-urlencoded" {
		t.Errorf("NewPlainRequest() Accept = %v, expected %v", contentType, "application/x-www-form-urlencoded")
	}
	// test default user-agent is attached to the request
	uA := req.Header.Get("User-Agent")
	if c.UserAgent != uA {
		t.Errorf("NewPlainRequest() User-Agent = %v, expected %v", uA, c.UserAgent)
	}
}

func TestDo(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodGet; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, ".", nil)
	body := new(foo)
	_, err := client.Do(context.Background(), req, body)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	expected := &foo{"a"}
	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Response body = %v, expected %v", body, expected)
	}
}

func TestCustomUserAgent(t *testing.T) {
	ua := fmt.Sprintf("testing/%s", atlas.Version)
	c, err := NewConfigWithOptions(nil, SetUserAgent(ua))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	expected := fmt.Sprintf("%s %s", ua, userAgent)
	if got := c.UserAgent; got != expected {
		t.Errorf("New() UserAgent = %s; expected %s", got, expected)
	}
}

func TestCustomBaseURL(t *testing.T) {
	const baseURL = "http://localhost/foo"
	c, err := NewConfigWithOptions(nil, SetAuthURL(baseURL))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}
	if got := c.AuthURL.String(); got != baseURL {
		t.Errorf("New() BaseURL = %s; expected %s", got, baseURL)
	}
}
