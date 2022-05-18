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
	"fmt"
	"net/http"
)

const federationSettingsConnectedOrganizationBasePath = "api/atlas/v1.0/federationSettings/%s/connectedOrgConfigs"

// FederatedSettingsIdentityProviderService is an interface for working with the Federation Settings Connected Organization
// endpoints of the MongoDB Atlas API.
// See more: https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/
type FederatedSettingsConnectedOrganizationService interface {
	List(context.Context, *ListOptions, string) (*FederatedSettingsConnectedOrganizations, *Response, error)
	Get(context.Context, string, string) (*FederatedSettingsConnectedOrganization, *Response, error)
	Update(context.Context, string, string, *FederatedSettingsConnectedOrganization) (*FederatedSettingsConnectedOrganization, *Response, error)
	Delete(context.Context, string, string) (*Response, error)
}

// FederatedSettingsIdentityProviderServiceOp handles communication with the FederatedSettings Connected Organization related methods of the
// MongoDB Atlas API.
type FederatedSettingsConnectedOrganizationSeviceOp service

var _ FederatedSettingsConnectedOrganizationService = &FederatedSettingsConnectedOrganizationSeviceOp{}

// A Resource describes a specific resource the Role will allow operating on.

// FederatedSettings represents a FederatedSettings Connected Organization.
type FederatedSettingsConnectedOrganizations struct {
	Links      []*Link                                   `json:"links,omitempty"`
	Results    []*FederatedSettingsConnectedOrganization `json:"results,omitempty"`
	TotalCount int                                       `json:"totalCount,omitempty"`
}

type FederatedSettingsConnectedOrganization struct {
	DomainAllowList          []string     `json:"domainAllowList,omitempty"`
	DomainRestrictionEnabled bool         `json:"domainRestrictionEnabled,omitempty"`
	IdentityProviderID       string       `json:"identityProviderId,omitempty"`
	OrgID                    string       `json:"orgId,omitempty"`
	PostAuthRoleGrants       []string     `json:"postAuthRoleGrants,omitempty"`
	RoleMappings             RoleMappings `json:"roleMappings,omitempty"`
	UserConflicts            []string     `json:"userConflicts,omitempty"`
}

type RoleMappings []struct {
	ExternalGroupName string `json:"externalGroupName,omitempty"`
	ID                string `json:"id,omitempty"`
	RoleAssignments   []struct {
		GroupID string `json:"groupId,omitempty"`
		OrgID   string `json:"orgId,omitempty"`
		Role    string `json:"role,omitempty"`
	} `json:"roleAssignments,omitempty"`
}

// List gets all Federated Settings Connected Organization (Org-Mappings).
//
// See more: https://www.mongodb.com/docs/atlas/reference/api/org-mappings-return-all/
func (s *FederatedSettingsConnectedOrganizationSeviceOp) List(ctx context.Context, opts *ListOptions, federationSettingsID string) (*FederatedSettingsConnectedOrganizations, *Response, error) {
	basePath := fmt.Sprintf(federationSettingsConnectedOrganizationBasePath, federationSettingsID)
	path, err := setListOptions(basePath, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(FederatedSettingsConnectedOrganizations)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root, resp, nil
}

// Get getsFederated Settings Connected Organization (Org-Mapping).
//
// See more: https://www.mongodb.com/docs/atlas/reference/api/org-mapping-return-one/
func (s *FederatedSettingsConnectedOrganizationSeviceOp) Get(ctx context.Context, federationSettingsID, orgID string) (*FederatedSettingsConnectedOrganization, *Response, error) {
	if federationSettingsID == "" {
		return nil, nil, NewArgError("federationSettingsID", "must be set")
	}

	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	basePath := fmt.Sprintf(federationSettingsConnectedOrganizationBasePath, federationSettingsID)
	path := fmt.Sprintf("%s/%s", basePath, orgID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(FederatedSettingsConnectedOrganization)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Update updates Federated Settings Connected Organization (Org-Mapping).
//
// See more: https://www.mongodb.com/docs/atlas/reference/api/org-mapping-update-one/
func (s *FederatedSettingsConnectedOrganizationSeviceOp) Update(ctx context.Context, federationSettingsID, orgID string, updateRequest *FederatedSettingsConnectedOrganization) (*FederatedSettingsConnectedOrganization, *Response, error) {
	if federationSettingsID == "" {
		return nil, nil, NewArgError("federationSettingsID", "must be set")
	}

	if orgID == "" {
		return nil, nil, NewArgError("orgID", "must be set")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	basePath := fmt.Sprintf(federationSettingsConnectedOrganizationBasePath, federationSettingsID)
	path := fmt.Sprintf("%s/%s", basePath, orgID)

	req, err := s.Client.NewRequest(ctx, http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(FederatedSettingsConnectedOrganization)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Delete deletes federation setting.
//
// See more: https://www.mongodb.com/docs/atlas/reference/api/org-mapping-remove-one/
func (s *FederatedSettingsConnectedOrganizationSeviceOp) Delete(ctx context.Context, federationSettingsID, orgID string) (*Response, error) {
	if federationSettingsID == "" {
		return nil, NewArgError("federationSettingsID", "must be set")
	}

	if orgID == "" {
		return nil, NewArgError("orgID", "must be set")
	}

	basePath := fmt.Sprintf(federationSettingsConnectedOrganizationBasePath, federationSettingsID)
	path := fmt.Sprintf("%s/%s", basePath, orgID)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}
