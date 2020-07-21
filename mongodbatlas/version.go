package mongodbatlas

import "fmt"

const (
	// UserAgent the default user agent string reported by this client
	UserAgent = "go-mongodbatlas"
	// ClientVersion the version of the current API client
	ClientVersion = "0.4.0" // Should be set to the next version planned to be released
)

// DefaultUserAgent constructs the default user agent
func DefaultUserAgent() string {
	return fmt.Sprintf("%s/%s", UserAgent, ClientVersion)
}
