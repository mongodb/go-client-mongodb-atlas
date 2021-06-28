// Copyright 2019 MongoDB Inc
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
	"fmt"
	"net/http"
)

const (
	orgsBasePath       = "api/atlas/v1.0/orgs"
	invitationBasePath = orgsBasePath + "/%s/invites"
)

// OrganizationsService provides access to the organization related functions in the Atlas API.
//
// See more: https://docs.atlas.mongodb.com/reference/api/organizations/
type OrganizationsService interface {
	List(context.Context, *OrganizationsListOptions) (*Organizations, *Response, error)
	ListUnacceptedInvitations(context.Context, string, *InvitationOptions) (*[]Invitation, *Response, error)
	Get(context.Context, string) (*Organization, *Response, error)
	GetUnacceptedInvitation(context.Context, string, string) (*Invitation, *Response, error)
	Projects(context.Context, string, *ListOptions) (*Projects, *Response, error)
	Users(context.Context, string, *ListOptions) (*AtlasUsersResponse, *Response, error)
	Delete(context.Context, string) (*Response, error)
	InviteUser(context.Context, *Invitation) (*Invitation, *Response, error)
	UpdateInvitation(context.Context, *Invitation) (*Invitation, *Response, error)
	DeleteInvitation(context.Context, string, string) (*Response, error)
}

// OrganizationsServiceOp provides an implementation of the OrganizationsService interface.
type OrganizationsServiceOp service

var _ OrganizationsService = &OrganizationsServiceOp{}

// OrganizationsListOptions filtering options for organizations.
type OrganizationsListOptions struct {
	Name               string `url:"name,omitempty"`
	IncludeDeletedOrgs *bool  `url:"includeDeletedOrgs,omitempty"`
	ListOptions
}

// Organization represents the structure of an organization.
type Organization struct {
	ID        string  `json:"id,omitempty"`
	IsDeleted *bool   `json:"isDeleted,omitempty"`
	Links     []*Link `json:"links,omitempty"`
	Name      string  `json:"name,omitempty"`
}

// Organizations represents an array of organization.
type Organizations struct {
	Links      []*Link         `json:"links"`
	Results    []*Organization `json:"results"`
	TotalCount int             `json:"totalCount"`
}

// InvitationOptions filtering options for invitations.
type InvitationOptions struct {
	Username string `url:"username,omitempty"`
}

// Invitation represents the structure of an Invitation.
type Invitation struct {
	ID              string   `json:"id,omitempty"`
	OrgID           string   `json:"orgId,omitempty"`
	OrgName         string   `json:"orgName,omitempty"`
	CreatedAt       string   `json:"createdAt,omitempty"`
	ExpiresAt       string   `json:"expiresAt,omitempty"`
	InviterUserName string   `json:"inviterUserName,omitempty"`
	Username        string   `json:"username,omitempty"`
	Roles           []string `json:"roles,omitempty"`
	TeamIDs         []string `json:"teamIds,omitempty"`
}

// List gets all organizations.
//
// See more: https://docs.atlas.mongodb.com/reference/api/organization-get-all/
func (s *OrganizationsServiceOp) List(ctx context.Context, opts *OrganizationsListOptions) (*Organizations, *Response, error) {
	path, err := setListOptions(orgsBasePath, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Organizations)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root, resp, nil
}

// Get gets a single organization.
//
// See more: https://docs.atlas.mongodb.com/reference/api/organization-get-one/
func (s *OrganizationsServiceOp) Get(ctx context.Context, orgID string) (*Organization, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	path := fmt.Sprintf("%s/%s", orgsBasePath, orgID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Organization)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Projects gets all projects for the given organization ID.
//
// See more: https://docs.atlas.mongodb.com/reference/api/organization-get-all-projects/
func (s *OrganizationsServiceOp) Projects(ctx context.Context, orgID string, opts *ListOptions) (*Projects, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}
	basePath := fmt.Sprintf("%s/%s/groups", orgsBasePath, orgID)

	path, err := setListOptions(basePath, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Projects)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Users gets all users for the given organization ID.
//
// See more: https://docs.atlas.mongodb.com/reference/api/organization-users-get-all-users/
func (s *OrganizationsServiceOp) Users(ctx context.Context, orgID string, opts *ListOptions) (*AtlasUsersResponse, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}
	basePath := fmt.Sprintf("%s/%s/users", orgsBasePath, orgID)

	path, err := setListOptions(basePath, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AtlasUsersResponse)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Delete deletes an organization.
//
// See more: https://docs.atlas.mongodb.com/reference/api/organization-delete-one/
func (s *OrganizationsServiceOp) Delete(ctx context.Context, orgID string) (*Response, error) {
	if orgID == "" {
		return nil, NewArgError("orgID", "must be set")
	}

	basePath := fmt.Sprintf("%s/%s", orgsBasePath, orgID)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, basePath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}

// ListUnacceptedInvitations gets all unaccepted invitations to the specified Atlas organization.
//
// See more: https://docs.atlas.mongodb.com/reference/api/organization-get-invitations/
func (s *OrganizationsServiceOp) ListUnacceptedInvitations(ctx context.Context, orgID string, opts *InvitationOptions) (*[]Invitation, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	basePath := fmt.Sprintf(invitationBasePath, orgID)
	path, err := setListOptions(basePath, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new([]Invitation)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// GetUnacceptedInvitation gets details for one unaccepted invitation to the specified Atlas organization.
//
// See more: https://docs.atlas.mongodb.com/reference/api/organization-get-one-invitation/
func (s *OrganizationsServiceOp) GetUnacceptedInvitation(ctx context.Context, orgID, invitationID string) (*Invitation, *Response, error) {
	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	basePath := fmt.Sprintf(invitationBasePath, orgID)
	path := fmt.Sprintf("%s/%s", basePath, invitationID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Invitation)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// InviteUser invites one user to the Atlas organization that you specify.
//
// See more: https://docs-atlas-staging.mongodb.com/cloud-docs/docsworker-xlarge/DOCSP-14695/reference/api/organization-create-one-invitation/
func (s *OrganizationsServiceOp) InviteUser(ctx context.Context, invitation *Invitation) (*Invitation, *Response, error) {
	if invitation.OrgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	path := fmt.Sprintf(invitationBasePath, invitation.OrgID)

	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, invitation)
	if err != nil {
		return nil, nil, err
	}

	root := new(Invitation)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// InviteUser invites one user to the Atlas organization that you specify.
//
// See more: https://docs-atlas-staging.mongodb.com/cloud-docs/docsworker-xlarge/DOCSP-14695/reference/api/organization-create-one-invitation/
func (s *OrganizationsServiceOp) UpdateInvitation(ctx context.Context, invitation *Invitation) (*Invitation, *Response, error) {
	if invitation.OrgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	path := fmt.Sprintf(invitationBasePath, invitation.OrgID)

	if invitation.ID != "" {
		path = fmt.Sprintf("%s/%s", path, invitation.ID)
	}

	req, err := s.Client.NewRequest(ctx, http.MethodPatch, path, invitation)
	if err != nil {
		return nil, nil, err
	}

	root := new(Invitation)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// Delete deletes one unaccepted invitation to the specified Atlas organization. You can't delete an invitation that a user has accepted.
//
// See more: https://docs-atlas-staging.mongodb.com/cloud-docs/docsworker-xlarge/DOCSP-14695/reference/api/organization-delete-invitation/
func (s *OrganizationsServiceOp) DeleteInvitation(ctx context.Context, orgID, invitationID string) (*Response, error) {
	if orgID == "" {
		return nil, NewArgError("orgID", "must be set")
	}

	basePath := fmt.Sprintf(invitationBasePath, orgID)
	path := fmt.Sprintf("%s/%s", basePath, invitationID)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}
