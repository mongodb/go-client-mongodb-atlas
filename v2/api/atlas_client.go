package api // import "go.mongodb.org/atlas/mongodbatlas/v2"

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/mongodb-forks/digest"
)

const (
	// CloudURL is default base URL for the services.
	DefaultCloudURL = "https://cloud.mongodb.com/"
	// Version the version of the current API client. Should be set to the next version planned to be released.
	Version    = "1.20230201.0-dev1"
	ClientName = "go-mongodbatlas-versioned"
)

// NewClient returns a new MongoDBAtlas API Client.
func NewClient(httpClient *http.Client) *APIClient {
	return NewClientWithUrl(httpClient, DefaultCloudURL)
}

// NewClientWithUrl returns a new MongoDBAtlas API Client using custom base url
func NewClientWithUrl(httpClient *http.Client, baseURL string) *APIClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	userAgent := fmt.Sprintf("%s/%s (%s;%s)", ClientName, Version, runtime.GOOS, runtime.GOARCH)
	cfg := &Configuration{
		HTTPClient: httpClient,
		Servers: ServerConfigurations{ServerConfiguration{
			URL: baseURL,
		},
		},
		UserAgent: userAgent,
	}
	client := NewAPIClient(cfg)
	return client
}

/**
*  NewSDKClientWithCredentials Helper method used to create an SDK object for provided API keys
 */
func NewSDKClientWithCredentials(apiKey string, apiSecret string) (*APIClient, error) {
	transport := digest.NewTransport(apiKey, apiSecret)
	httpClient, err := transport.Client()
	if err != nil {
		return nil, err
	}
	return NewClient(httpClient), nil
}
