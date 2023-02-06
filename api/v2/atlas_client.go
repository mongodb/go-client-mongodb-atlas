package v2 // import "go.mongodb.org/atlas/mongodbatlas/v2"

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	// CloudURL is default base URL for the services.
	DefaultCloudURL = "https://cloud.mongodb.com/"
	// Version the version of the current API client inherited from.
	Version = mongodbatlas.Version
	// Name of the v2 API client.
	ClientName = "go-mongodbatlas-v2"
)

// NewClient returns a new MongoDBAtlas API Client.
func NewClient(modifiers ...ClientModifier) *APIClient {
	userAgent := fmt.Sprintf("%s/%s (%s;%s)", ClientName, Version, runtime.GOOS, runtime.GOARCH)
	defaultConfig := &Configuration{
		HTTPClient: http.DefaultClient,
		Servers: ServerConfigurations{ServerConfiguration{
			URL: DefaultCloudURL,
		},
		},
		UserAgent: userAgent,
	}
	for _, modifierFunction := range modifiers {
		modifierFunction(defaultConfig)
	}

	return NewAPIClient(defaultConfig)
}

// ClientModifiers lets you create function that controls configuration before creating client.
type ClientModifier func(*Configuration) error

func UseDigestAuth(apiKey, apiSecret string) ClientModifier {
	return func(c *Configuration) error {
		transport := digest.NewTransport(apiKey, apiSecret)
		httpClient, err := transport.Client()
		if err != nil {
			return err
		}
		c.HTTPClient = httpClient
		return nil
	}
}

// Advanced modifiers.

// UseHttpClient set custom http client implementation.
func UseHTTPClient(client *http.Client) ClientModifier {
	return func(c *Configuration) error {
		c.HTTPClient = client
		return nil
	}
}

// UseDebug enable debug level logging.
func UseDebug(debug bool) ClientModifier {
	return func(c *Configuration) error {
		c.Debug = debug
		return nil
	}
}

// UseBaseURL set custom base url.
func UseBaseURL(baseURL string) ClientModifier {
	return func(c *Configuration) error {
		c.Servers = ServerConfigurations{ServerConfiguration{
			URL: baseURL,
		}}
		return nil
	}
}
