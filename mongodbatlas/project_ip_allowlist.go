package mongodbatlas

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const projectIPAllowlistPath = "groups/%s/whitelist"

// ProjectIPAllowlistService is an interface for interfacing with the Project IP Allowlist
// endpoints of the MongoDB Atlas API.
// See more: https://docs.atlas.mongodb.com/reference/api/whitelist/
type ProjectIPAllowlistService interface {
	List(context.Context, string, *ListOptions) ([]ProjectIPAllowlist, *Response, error)
	Get(context.Context, string, string) (*ProjectIPAllowlist, *Response, error)
	Create(context.Context, string, []*ProjectIPAllowlist) ([]ProjectIPAllowlist, *Response, error)
	Update(context.Context, string, []*ProjectIPAllowlist) ([]ProjectIPAllowlist, *Response, error)
	Delete(context.Context, string, string) (*Response, error)
}

// ProjectIPAllowlistServiceOp handles communication with the ProjectIPAllowlist related methods
// of the MongoDB Atlas API
type ProjectIPAllowlistServiceOp service

var _ ProjectIPAllowlistService = &ProjectIPAllowlistServiceOp{}

// ProjectIPAllowlist represents MongoDB project's IP allowlist.
type ProjectIPAllowlist struct {
	GroupID          string `json:"groupId,omitempty"`          // The unique identifier for the project for which you want to update one or more allowlist entries.
	AwsSecurityGroup string `json:"awsSecurityGroup,omitempty"` // ID of the allowed AWS security group to update. Mutually exclusive with cidrBlock and ipAddress.
	CIDRBlock        string `json:"cidrBlock,omitempty"`        // Allowlist entry in Classless Inter-Domain Routing (CIDR) notation to update. Mutually exclusive with awsSecurityGroup and ipAddress.
	IPAddress        string `json:"ipAddress,omitempty"`        // Allowlisted IP address to update. Mutually exclusive with awsSecurityGroup and cidrBlock.
	Comment          string `json:"comment,omitempty"`          // Optional The comment associated with the allowlist entry. Specify an empty string "" to delete the comment associated to an IP address.
	DeleteAfterDate  string `json:"deleteAfterDate,omitempty"`  // Optional The ISO-8601-formatted UTC date after which Atlas removes the entry from the allowlist. The specified date must be in the future and within one week of the time you make the API request. To update a temporary allowlist entry to be permanent, set the value of this field to null
}

// projectIPAllowlistsResponse is the response from the ProjectIPAllowlistService.List.
type projectIPAllowlistsResponse struct {
	Links      []*Link              `json:"links"`
	Results    []ProjectIPAllowlist `json:"results"`
	TotalCount int                  `json:"totalCount"`
}

// Create adds one or more allowed entries to the project associated to {GROUP-ID}.
// See more: https://docs.atlas.mongodb.com/reference/api/database-users-create-a-user/
func (s *ProjectIPAllowlistServiceOp) Create(ctx context.Context, groupID string, createRequest []*ProjectIPAllowlist) ([]ProjectIPAllowlist, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := fmt.Sprintf(projectIPAllowlistPath, groupID)

	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(projectIPAllowlistsResponse)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root.Results, resp, err
}

// Get gets the allowlist entry specified to {WHITELIST-ENTRY} from the project associated to {GROUP-ID}.
// See more: https://docs.atlas.mongodb.com/reference/api/whitelist-get-one-entry/
func (s *ProjectIPAllowlistServiceOp) Get(ctx context.Context, groupID, allowListEntry string) (*ProjectIPAllowlist, *Response, error) {
	if allowListEntry == "" {
		return nil, nil, NewArgError("allowListEntry", "must be set")
	}

	basePath := fmt.Sprintf(projectIPAllowlistPath, groupID)
	escapedEntry := url.PathEscape(allowListEntry)
	path := fmt.Sprintf("%s/%s", basePath, escapedEntry)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ProjectIPAllowlist)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// List all allowlist entries in the project associated to {GROUP-ID}.
// See more: https://docs.atlas.mongodb.com/reference/api/whitelist-get-all/
func (s *ProjectIPAllowlistServiceOp) List(ctx context.Context, groupID string, listOptions *ListOptions) ([]ProjectIPAllowlist, *Response, error) {
	path := fmt.Sprintf(projectIPAllowlistPath, groupID)

	// Add query params from listOptions
	path, err := setListOptions(path, listOptions)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(projectIPAllowlistsResponse)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root.Results, resp, nil
}

// Update one or more allowlist entries in the project associated to {GROUP-ID}
// See more: https://docs.atlas.mongodb.com/reference/api/whitelist-update-one/
func (s *ProjectIPAllowlistServiceOp) Update(ctx context.Context, groupID string, updateRequest []*ProjectIPAllowlist) ([]ProjectIPAllowlist, *Response, error) {
	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf(projectIPAllowlistPath, groupID)

	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(projectIPAllowlistsResponse)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root.Results, resp, err
}

// Delete the allowlist entry specified to {WHITELIST-ENTRY} from the project associated to {GROUP-ID}.
// See more: https://docs.atlas.mongodb.com/reference/api/whitelist-delete-one/
func (s *ProjectIPAllowlistServiceOp) Delete(ctx context.Context, groupID, allowListEntry string) (*Response, error) {
	if allowListEntry == "" {
		return nil, NewArgError("allowListEntry", "must be set")
	}

	basePath := fmt.Sprintf(projectIPAllowlistPath, groupID)
	escapedEntry := url.PathEscape(allowListEntry)
	path := fmt.Sprintf("%s/%s", basePath, escapedEntry)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}
