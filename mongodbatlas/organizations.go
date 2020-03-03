package mongodbatlas

import (
	"context"
	"fmt"
	"net/http"
)

const (
	organizationsBasePath = "orgs"
)

// OrganizationsService is an interface for interfacing with the Organization
// endpoints of the MongoDB Atlas API.
// See more: https://docs.atlas.mongodb.com/reference/api/organizations/
type OrganizationsService interface {
	List(context.Context, *ListOptions) ([]Organization, *Response, error)
	Get(context.Context, string) (*Organization, *Response, error)
	ListUsers(context.Context, string, *ListOptions) ([]AtlasUser, *Response, error)
	Rename(context.Context, string, string) (*Organization, *Response, error)
	Delete(context.Context, string) (*Organization, *Response, error)
	GetAllProjects(context.Context, string, *ListOptions) (*Projects, *Response, error)
}

//OrganizationsServiceOp handles communication with the Organization related methos of the
//MongoDB Atlas API
type OrganizationsServiceOp struct {
	client *Client
}

var _ OrganizationsService = &OrganizationsServiceOp{}

//OrganizationResponse represents a array of organization
type OrganizationResponse struct {
	Links      []*Link        `json:"links"`
	Results    []Organization `json:"results"`
	TotalCount int            `json:"totalCount"`
}

// Organization
type Organization struct {
	Name  string `json:"name"`
	ID    string `json:"id,omitempty"`
	Links string `json:"links"`
}

//List gets all organizations.
//See more: https://docs.atlas.mongodb.com/reference/api/organization-get-all/
func (s *OrganizationsServiceOp) List(ctx context.Context, listOptions *ListOptions) ([]Organization, *Response, error) {
	path := fmt.Sprintf(organizationsBasePath)

	//Add query params from listOptions
	path, err := setListOptions(path, listOptions)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(OrganizationResponse)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root.Results, resp, nil
}

//Get gets a single organization.
//See more: https://docs.atlas.mongodb.com/reference/api/organization-get-one/
func (s *OrganizationsServiceOp) Get(ctx context.Context, orgID string) (*Organization, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	path := fmt.Sprintf("orgs/%s", orgID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Organization)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

//ListUsers list all users in a organization.
//See more: https://docs.atlas.mongodb.com/reference/api/organization-users-get-all-users/
func (s *OrganizationsServiceOp) ListUsers(ctx context.Context, orgID string, listOptions *ListOptions) ([]AtlasUser, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	path := fmt.Sprintf("%s/%s/users", organizationsBasePath, orgID)

	//Add query params from listOptions
	path, err := setListOptions(path, listOptions)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AtlasUsersResponse)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root.Results, resp, nil
}

//Rename an organization.
//See more: https://docs.atlas.mongodb.com/reference/api/organization-rename/
func (s *OrganizationsServiceOp) Rename(ctx context.Context, orgID string, newName string) (*Organization, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	if newName == "" {
		return nil, nil, NewArgError("newName", "must be set")
	}

	path := fmt.Sprintf("orgs/%s", orgID)

	req, err := s.client.NewRequest(ctx, http.MethodPatch, path, newName)
	if err != nil {
		return nil, nil, err
	}

	root := new(Organization)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

//Delete an organization.
//See more: https://docs.atlas.mongodb.com/reference/api/organization-delete-one/
func (s *OrganizationsServiceOp) Delete(ctx context.Context, orgID string) (*Organization, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	path := fmt.Sprintf("orgs/%s", orgID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Organization)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

//ListProjects list all projects in a organization.
//See more: https://docs.atlas.mongodb.com/reference/api/organization-get-all-projects/
func (s *OrganizationsServiceOp) GetAllProjects(ctx context.Context, orgID string, listOptions *ListOptions) (*Projects, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	path := fmt.Sprintf("%s/%s/groups", organizationsBasePath, orgID)

	//Add query params from listOptions
	path, err := setListOptions(path, listOptions)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Projects)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root, resp, nil
}
