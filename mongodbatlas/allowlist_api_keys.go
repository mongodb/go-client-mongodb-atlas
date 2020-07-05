package mongodbatlas

import (
	"context"
	"fmt"
	"net/http"
)

const allowlistAPIKeysPath = "orgs/%s/apiKeys/%s/whitelist"

// AllowlistAPIKeysService is an interface for interfacing with the Allowlist API Keys
// endpoints of the MongoDB Atlas API.
// See more: https://docs.atlas.mongodb.com/reference/api/apiKeys/#organization-api-key-endpoints
type AllowlistAPIKeysService interface {
	List(context.Context, string, string, *ListOptions) (*AllowlistAPIKeys, *Response, error)
	Get(context.Context, string, string, string) (*AllowlistAPIKey, *Response, error)
	Create(context.Context, string, string, []*AllowlistAPIKeysReq) (*AllowlistAPIKeys, *Response, error)
	Delete(context.Context, string, string, string) (*Response, error)
}

// AllowlistAPIKeysServiceOp handles communication with the Allowlist API keys related methods of the
// MongoDB Atlas API
type AllowlistAPIKeysServiceOp service

var _ AllowlistAPIKeysService = &AllowlistAPIKeysServiceOp{}

// AllowlistAPIKey represents an Allowlist API key.
type AllowlistAPIKey struct {
	CidrBlock       string  `json:"cidrBlock,omitempty"`       // CIDR-notated range of allowlisted IP addresses.
	Count           int     `json:"count,omitempty"`           // Total number of requests that have originated from this IP address.
	Created         string  `json:"created,omitempty"`         // Date this IP address was added to the allowlist.
	IPAddress       string  `json:"ipAddress,omitempty"`       // Allowed IP address.
	LastUsed        string  `json:"lastUsed,omitempty"`        // Date of the most recent request that originated from this IP address. This field only appears if at least one request has originated from this IP address, and is only updated when an allowlisted resource is accessed.
	LastUsedAddress string  `json:"lastUsedAddress,omitempty"` // IP address from which the last call to the API was issued. This field only appears if at least one request has originated from this IP address.
	Links           []*Link `json:"links,omitempty"`           // An array of documents, representing a link to one or more sub-resources and/or related resources such as list pagination. See Linking for more information.}
}

// AllowlistAPIKeys represents all Allowlisted API keys.
type AllowlistAPIKeys struct {
	Results    []*AllowlistAPIKey `json:"results,omitempty"`    // Includes one AllowlistAPIKey object for each item detailed in the results array section.
	Links      []*Link            `json:"links,omitempty"`      // One or more links to sub-resources and/or related resources.
	TotalCount int                `json:"totalCount,omitempty"` // Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated.
}

// AllowlistAPIKeysReq represents the request to the mehtod create
type AllowlistAPIKeysReq struct {
	IPAddress string `json:"ipAddress,omitempty"` // IP address to be added to the allowlist for the API key.
	CidrBlock string `json:"cidrBlock,omitempty"` // Allowed entry in CIDR notation to be added for the API key.
}

// List gets all Allowlisted API keys.
// See more: https://docs.atlas.mongodb.com/reference/api/apiKeys-org-whitelist-get-all/
func (s *AllowlistAPIKeysServiceOp) List(ctx context.Context, orgID, apiKeyID string, listOptions *ListOptions) (*AllowlistAPIKeys, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}
	if apiKeyID == "" {
		return nil, nil, NewArgError("apiKeyID", "must be set")
	}

	path := fmt.Sprintf(allowlistAPIKeysPath, orgID, apiKeyID)
	path, err := setListOptions(path, listOptions)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AllowlistAPIKeys)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root, resp, nil
}

// Get gets the Allowlist API keys.
// See more: https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-get-one/
func (s *AllowlistAPIKeysServiceOp) Get(ctx context.Context, orgID, apiKeyID, ipAddress string) (*AllowlistAPIKey, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}
	if apiKeyID == "" {
		return nil, nil, NewArgError("apiKeyID", "must be set")
	}
	if ipAddress == "" {
		return nil, nil, NewArgError("ipAddress", "must be set")
	}

	path := fmt.Sprintf(allowlistAPIKeysPath+"/%s", orgID, apiKeyID, ipAddress)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AllowlistAPIKey)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Create a submit a POST request containing ipAddress or cidrBlock values which are not already present in the allowlist, Atlas adds those entries to the list of existing entries in the allowlist.
// See more: https://docs.atlas.mongodb.com/reference/api/apiKeys-org-whitelist-create/
func (s *AllowlistAPIKeysServiceOp) Create(ctx context.Context, orgID, apiKeyID string, createRequest []*AllowlistAPIKeysReq) (*AllowlistAPIKeys, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}
	if apiKeyID == "" {
		return nil, nil, NewArgError("apiKeyID", "must be set")
	}
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := fmt.Sprintf(allowlistAPIKeysPath, orgID, apiKeyID)

	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(AllowlistAPIKeys)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Delete deletes the Allowlist API keys.
// See more: https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-delete-one/
func (s *AllowlistAPIKeysServiceOp) Delete(ctx context.Context, orgID, apiKeyID, ipAddress string) (*Response, error) {
	if orgID == "" {
		return nil, NewArgError("groupId", "must be set")
	}
	if apiKeyID == "" {
		return nil, NewArgError("clusterName", "must be set")
	}
	if ipAddress == "" {
		return nil, NewArgError("snapshotId", "must be set")
	}

	path := fmt.Sprintf(allowlistAPIKeysPath+"/%s", orgID, apiKeyID, ipAddress)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}
