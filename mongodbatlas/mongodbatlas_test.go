// Copyright 2021 MongoDB Inc
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

package mongodbatlas

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
)

var (
	ctx = context.TODO()
)

const (
	// baseURLPath is a non-empty Client.BaseURL path to use during tests,
	// to ensure relative URLs are used for all endpoints.
	baseURLPath = "/api-v1"
)

// setup sets up a test HTTP server along with a mongodbatlas.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, teardown func()) {
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
	client = NewClient(nil)
	u, _ := url.Parse(server.URL + baseURLPath + "/")
	client.BaseURL = u

	return client, mux, server.Close
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	t.Helper()
	if expected != r.Method {
		t.Fatalf("Request method = %q, expected %q", r.Method, expected)
	}
}

func testURLParseError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	var target *url.Error
	if errors.As(err, &target) && target.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testClientDefaultBaseURL(t *testing.T, c *Client) {
	t.Helper()
	if c.BaseURL == nil || c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, expected %v", c.BaseURL, defaultBaseURL)
	}
}

func testClientDefaultUserAgent(t *testing.T, c *Client) {
	t.Helper()
	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, expected %v", c.UserAgent, userAgent)
	}
}

func testClientDefaults(t *testing.T, c *Client) {
	t.Helper()
	testClientDefaultBaseURL(t, c)
	testClientDefaultUserAgent(t, c)
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)
	testClientDefaults(t, c)
}

func TestNew(t *testing.T) {
	c, err := New(nil)

	if err != nil {
		t.Fatalf("New(): %v", err)
	}
	testClientDefaults(t, c)
}

type testRequestBody struct {
	TestName    string `json:"testName"`
	TestCounter int64  `json:"testCounter"`
	TestData    string `json:"testUserData"`
}

// If a nil body is passed to mongodbatlas.NewRequest, make sure that nil is also
// passed to http.NewRequest. In most cases, passing an io.Reader that returns
// no content is fine, since there is no difference between an HTTP request
// body that is an empty string versus one that is not set at all. However in
// certain cases, intermediate systems may treat these differently resulting in
// subtle errors.
func TestNewRequest_emptyBody(t *testing.T) {
	c := NewClient(nil)
	req, err := c.NewRequest(ctx, http.MethodGet, ".", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if req.Body != nil {
		t.Fatalf("constructed request contains a non-nil Body")
	}
}

func TestNewRequest_withUserData(t *testing.T) {
	c := NewClient(nil)

	requestPath := "foo"

	inURL, outURL := requestPath, defaultBaseURL+requestPath
	inBody, outBody := &testRequestBody{TestName: "l", TestData: "u"},
		`{"testName":"l","testCounter":0,`+
			`"testUserData":"u"}`+"\n"
	req, _ := c.NewRequest(ctx, http.MethodGet, inURL, inBody)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test body was JSON encoded
	body, _ := io.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v)Body = %v, expected %v", inBody, string(body), outBody)
	}

	// test default user-agent is attached to the request
	uA := req.Header.Get("User-Agent")
	if c.UserAgent != uA {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", uA, c.UserAgent)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest(ctx, http.MethodGet, ":", nil)
	testURLParseError(t, err)
}

func TestNewRequest_withCustomUserAgent(t *testing.T) {
	ua := fmt.Sprintf("testing/%s", Version)
	c, err := New(nil, SetUserAgent(ua))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	req, _ := c.NewRequest(ctx, http.MethodGet, "/foo", nil)

	expected := fmt.Sprintf("%s %s", ua, userAgent)
	if got := req.Header.Get("User-Agent"); got != expected {
		t.Errorf("New() UserAgent = %s; expected %s", got, expected)
	}
}

func TestNewRequest_noErrorForNoTrailingSlash(t *testing.T) {
	tests := []struct {
		rawurl    string
		wantError bool
	}{
		{rawurl: "https://example.com/api/v1", wantError: false},
		{rawurl: "https://example.com/api/v1/", wantError: false},
	}
	c := NewClient(nil)
	for _, test := range tests {
		u, err := url.Parse(test.rawurl)
		if err != nil {
			t.Fatalf("url.Parse returned unexpected error: %v.", err)
		}
		c.BaseURL = u
		if _, err := c.NewRequest(ctx, http.MethodGet, "test", nil); test.wantError && err == nil {
			t.Fatalf("Expected error to be returned.")
		} else if !test.wantError && err != nil {
			t.Fatalf("NewRequest returned unexpected error: %v.", err)
		}
	}
}

func TestNewRequest_correctURLWithNoTrailingSlash(t *testing.T) {
	tests := []struct {
		rawURL      string
		expectedURL string
	}{
		{rawURL: "http://test.com", expectedURL: "http://test.com/"},
		{rawURL: "http://home.base.com/", expectedURL: "http://home.base.com/"},
	}

	c := NewClient(nil)
	for _, test := range tests {
		u, err := url.Parse(test.rawURL)
		if err != nil {
			t.Fatalf("url.Parse returned unexpected error: %v.", err)
		}
		c.BaseURL = u
		req, err := c.NewRequest(ctx, http.MethodGet, "", nil)
		if err != nil {
			t.Fatalf("NewRequest return unexpected error: %v.", err)
		}
		if req.URL.String() != test.expectedURL {
			t.Fatalf("Incorrectly added trailing slash. Expected: %v; Got: %v.", test.expectedURL, req.URL.String())
		}
	}
}

func TestNewGZipRequest_emptyBody(t *testing.T) {
	c := NewClient(nil)
	req, err := c.NewGZipRequest(ctx, http.MethodGet, ".")
	if err != nil {
		t.Fatalf("NewGZipRequest returned unexpected error: %v", err)
	}
	if req.Body != nil {
		t.Fatalf("constructed request contains a non-nil Body")
	}
}

func TestNewGZipRequest_withCustomUserAgent(t *testing.T) {
	ua := fmt.Sprintf("testing/%s", Version)
	c, err := New(nil, SetUserAgent(ua))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	req, _ := c.NewGZipRequest(ctx, http.MethodGet, "/foo")

	expected := fmt.Sprintf("%s %s", ua, userAgent)
	if got := req.Header.Get("User-Agent"); got != expected {
		t.Errorf("NewGZipRequest() UserAgent = %s; expected %s", got, expected)
	}
}

func TestNewGZipRequest_badURL(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewGZipRequest(ctx, http.MethodGet, ":")
	testURLParseError(t, err)
}

func TestNewGZipRequest(t *testing.T) {
	c := NewClient(nil)

	requestPath := "foo"

	inURL, outURL := requestPath, defaultBaseURL+requestPath
	req, _ := c.NewGZipRequest(ctx, http.MethodGet, inURL)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewGZipRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test accept content type is correct
	accept := req.Header.Get("Accept")
	if gzipMediaType != accept {
		t.Errorf("NewGZipRequest() Accept = %v, expected %v", accept, gzipMediaType)
	}

	// test default user-agent is attached to the request
	uA := req.Header.Get("User-Agent")
	if c.UserAgent != uA {
		t.Errorf("NewGZipRequest() User-Agent = %v, expected %v", uA, c.UserAgent)
	}
}

func TestNewPlainRequest_emptyBody(t *testing.T) {
	c := NewClient(nil)
	req, err := c.NewPlainRequest(ctx, http.MethodGet, ".")
	if err != nil {
		t.Fatalf("NewPlainRequest returned unexpected error: %v", err)
	}
	if req.Body != nil {
		t.Fatalf("constructed request contains a non-nil Body")
	}
}

func TestNewPlainRequest_withCustomUserAgent(t *testing.T) {
	ua := fmt.Sprintf("testing/%s", Version)
	c, err := New(nil, SetUserAgent(ua))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	req, _ := c.NewPlainRequest(ctx, http.MethodGet, "/foo")

	expected := fmt.Sprintf("%s %s", ua, userAgent)
	if got := req.Header.Get("User-Agent"); got != expected {
		t.Errorf("NewPlainRequest() UserAgent = %s; expected %s", got, expected)
	}
}

func TestNewPlainRequest_badURL(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewPlainRequest(ctx, http.MethodGet, ":")
	testURLParseError(t, err)
}

func TestNewPlainRequest(t *testing.T) {
	c := NewClient(nil)

	requestPath := "foo"

	inURL, outURL := requestPath, defaultBaseURL+requestPath
	req, _ := c.NewPlainRequest(ctx, http.MethodGet, inURL)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewPlainRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test accept content type is correct
	accept := req.Header.Get("Accept")
	if plainMediaType != accept {
		t.Errorf("NewPlainRequest() Accept = %v, expected %v", accept, plainMediaType)
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

func TestDo_httpError(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, ".", nil)
	_, err := client.Do(context.Background(), req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}
}

// Test handling of an error caused by the internal http Client's Do()
// function.
func TestDo_redirectLoop(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, baseURLPath, http.StatusFound)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, ".", nil)
	_, err := client.Do(context.Background(), req, nil)

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	var target *url.Error
	if !errors.As(err, &target) {
		t.Errorf("Expected a URL error; got %#v.", err)
	}
}

func TestDo_noContent(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var body json.RawMessage

	req, _ := client.NewRequest(ctx, http.MethodGet, ".", nil)
	_, err := client.Do(context.Background(), req, &body)
	if err != nil {
		t.Fatalf("Do returned unexpected error: %v", err)
	}
}

func TestCheckResponse(t *testing.T) {
	res := &Response{
		Response: &http.Response{
			Request:    &http.Request{},
			StatusCode: http.StatusBadRequest,
			Body: io.NopCloser(
				strings.NewReader(
					`{"error":409, "errorCode": "GROUP_ALREADY_EXISTS", "reason":"Conflict", "detail":"A group with name \"Test\" already exists"}`,
				),
			),
		},
	}
	var target *ErrorResponse
	if !errors.As(res.CheckResponse(res.Body), &target) {
		t.Fatalf("Expected error response.")
	}

	expected := &ErrorResponse{
		Response:  res.Response,
		HTTPCode:  409,
		ErrorCode: "GROUP_ALREADY_EXISTS",
		Reason:    "Conflict",
		Detail:    `A group with name "Test" already exists`,
	}
	if !errors.Is(target, expected) {
		t.Errorf("Got = %#v, expected %#v", target, expected)
	}
}

// ensure that we properly handle API errors that do not contain a response
// body.
func TestCheckResponse_noBody(t *testing.T) {
	res := &Response{
		Response: &http.Response{
			Request:    &http.Request{},
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(strings.NewReader("")),
		},
	}
	var target *ErrorResponse
	if !errors.As(res.CheckResponse(res.Body), &target) {
		t.Errorf("Expected error response.")
	}

	expected := &ErrorResponse{
		Response: res.Response,
	}
	if !errors.Is(target, expected) {
		t.Errorf("Got = %#v, expected %#v", target, expected)
	}
}

func TestErrorResponse_Error(t *testing.T) {
	res := &http.Response{Request: &http.Request{}}
	err := ErrorResponse{Reason: "m", Response: res}
	if err.Error() == "" {
		t.Errorf("Expected non-empty ErrorResponse.Error()")
	}
}

func TestDo_completion_callback(t *testing.T) {
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
	var completedReq *http.Request
	var completedResp string
	client.OnRequestCompleted(func(req *http.Request, resp *http.Response) {
		completedReq = req
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			t.Errorf("Failed to dump response: %s", err)
		}
		completedResp = string(b)
	})
	_, err := client.Do(context.Background(), req, body)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}
	if !reflect.DeepEqual(req, completedReq) {
		t.Errorf("Completed request = %v, expected %v", completedReq, req)
	}
	const expected = `{"A":"a"}`
	if !strings.Contains(completedResp, expected) {
		t.Errorf("expected response to contain %v, Response = %v", expected, completedResp)
	}
}

func TestDo_response_processed_callback(t *testing.T) {
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

	client.withRaw = true
	req, _ := client.NewRequest(ctx, http.MethodGet, ".", nil)
	body := new(foo)
	var completedReq *http.Request
	var completedResp string
	client.OnResponseProcessed(func(resp *Response) {
		completedReq = req
		completedResp = string(resp.Raw)
	})
	_, err := client.Do(context.Background(), req, body)
	client.withRaw = false
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}
	if !reflect.DeepEqual(req, completedReq) {
		t.Errorf("Completed request = %v, expected %v", completedReq, req)
	}
	const expected = `{"A":"a"}`
	if !strings.Contains(completedResp, expected) {
		t.Errorf("expected response to contain %v, Response = %v", expected, completedResp)
	}
}

func TestCustomUserAgent(t *testing.T) {
	ua := fmt.Sprintf("testing/%s", Version)
	c, err := New(nil, SetUserAgent(ua))

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
	c, err := New(nil, SetBaseURL(baseURL))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}
	if got := c.BaseURL.String(); got != baseURL {
		t.Errorf("New() BaseURL = %s; expected %s", got, baseURL)
	}
}

func TestCustomBaseURL_badURL(t *testing.T) {
	baseURL := ":"
	_, err := New(nil, SetBaseURL(baseURL))

	testURLParseError(t, err)
}

func checkCurrentPage(t *testing.T, resp *Response, expectedPage int) { //nolint:unparam // currently we always use expectedPage with value 2 but that may change
	t.Helper()
	p, err := resp.CurrentPage()
	if err != nil {
		t.Fatal(err)
	}

	if p != expectedPage {
		t.Fatalf("expected current page to be '%d', was '%d'", expectedPage, p)
	}
}

func TestSetListOptions_EmptyBoolPtr(t *testing.T) {
	type testListOptions struct {
		Exists *bool `url:"exists,omitempty"`
	}

	boolPtr := func(b bool) *bool {
		return &b
	}
	testBaseURL := "www.example.com"

	tests := []struct {
		testName string
		arg      *testListOptions
		expected string
	}{
		{
			testName: "true value",
			arg: &testListOptions{
				Exists: boolPtr(true),
			},
			expected: testBaseURL + "?exists=true",
		},
		{
			testName: "false value",
			arg: &testListOptions{
				Exists: boolPtr(false),
			},
			expected: testBaseURL + "?exists=false",
		},
		{
			testName: "omitted null value",
			arg:      &testListOptions{},
			expected: testBaseURL,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Helper()
			t.Log(tt.testName)

			actual, err := setListOptions(testBaseURL, tt.arg)
			if err != nil {
				t.Errorf("setListOptions() error = %v", err)
			}
			if actual != tt.expected {
				t.Errorf("setListOptions() actual = %s, expected = %s", actual, tt.expected)
			}
		})
	}
}

const testResponse = `{"A":"a"}`

func TestClient_withRaw(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	type foo struct {
		A string
	}

	client.withRaw = true

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodGet; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		_, _ = fmt.Fprint(w, testResponse)
	})

	body := new(foo)
	req, _ := client.NewRequest(ctx, http.MethodGet, ".", nil)
	resp, err := client.Do(context.Background(), req, body)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	if string(resp.Raw) != testResponse {
		t.Errorf("expected response to be %v, Response = %v", testResponse, string(resp.Raw))
	}
}

func TestSetWithRaw(t *testing.T) {
	c, err := New(nil, SetWithRaw())

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if !c.withRaw {
		t.Errorf("New() withRaw = %v", c.withRaw)
	}
}
