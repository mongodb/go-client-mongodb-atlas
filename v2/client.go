package v2 // import "go.mongodb.org/atlas/mongodbatlas/v2"

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/mongodb-forks/digest"
	apiv2latest "go.mongodb.org/atlas/mongodbatlas/v2/api"
)

const (
	// CloudURL is default base URL for the services.
	DefaultCloudURL = "https://cloud.mongodb.com/"
	// Version the version of the current API client. Should be set to the next version planned to be released.
	Version    = "0.1.0"
	ClientName = "go-mongodbatlas-versioned"
)

// NewClient returns a new MongoDBAtlas API Client.
func NewClient(httpClient *http.Client) *apiv2latest.APIClient {
	return NewClientWithUrl(httpClient, DefaultCloudURL)
}

// NewClientWithUrl returns a new MongoDBAtlas API Client using custom base url
func NewClientWithUrl(httpClient *http.Client, baseURL string) *apiv2latest.APIClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	userAgent := fmt.Sprintf("%s/%s (%s;%s)", ClientName, Version, runtime.GOOS, runtime.GOARCH)
	cfg := &apiv2latest.Configuration{
		HTTPClient: httpClient,
		Servers: apiv2latest.ServerConfigurations{{
			URL: baseURL,
		}},
		UserAgent: userAgent,
	}
	client := apiv2latest.NewAPIClient(cfg)
	return client
}

/**
*  NewSDKClientWithCredentials Helper method used to create an SDK object for provided API keys
 */
func NewSDKClientWithCredentials(apiKey string, apiSecret string) (*apiv2latest.APIClient, error) {
	transport := digest.NewTransport(apiKey, apiSecret)
	httpClient, err := transport.Client()
	if err != nil {
		return nil, err
	}
	return NewClient(httpClient), nil
}
