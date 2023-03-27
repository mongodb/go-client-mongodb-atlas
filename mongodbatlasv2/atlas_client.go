package mongodbatlasv2 // import "go.mongodb.org/atlas/mongodbatlas/v2"

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	// CloudURL is default base URL for the services.
	DefaultCloudURL = "https://cloud.mongodb.com"
	// Version the version of the current API client inherited from.
	Version = mongodbatlas.Version
	// Name of the v2 API client.
	ClientName = "go-mongodbatlas-v2"
)

// NewClient returns a new MongoDBAtlas API Client.
func NewClient(modifiers ...ClientModifier) (*APIClient, error) {
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
		err := modifierFunction(defaultConfig)
		if err != nil {
			return nil, err
		}
	}

	return NewAPIClient(defaultConfig), nil
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
		urlWithoutSuffix := strings.TrimSuffix(baseURL, "/")
		c.Servers = ServerConfigurations{ServerConfiguration{
			URL: urlWithoutSuffix,
		}}
		return nil
	}
}

// UseUserAgent set custom UserAgent header.
func UseUserAgent(userAgent string) ClientModifier {
	return func(c *Configuration) error {
		c.UserAgent = userAgent
		return nil
	}
}

// ListOptions specifies the optional parameters to List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	PageNum int32 `url:"pageNum,omitempty"`

	// For paginated result sets, the number of results to include per page.
	ItemsPerPage int32 `url:"itemsPerPage,omitempty"`

	// Flag that indicates whether Atlas returns the totalCount parameter in the response body.
	IncludeCount bool `url:"includeCount,omitempty"`
}
