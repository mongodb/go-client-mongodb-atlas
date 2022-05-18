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

const federationSettingsOrganizationConnectionBasePath = "api/atlas/v1.0/federationSettings/%s/connectedOrgConfigs/%s/roleMappings"

// FederatedSettingsIdentityProviderService is an interface for working with the Federation Settings Role Mapping
// endpoints of the MongoDB Atlas API.
// See more: https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/
type FederatedSettingsOrganizationConnectionService interface {
	List(context.Context, *ListOptions, string, string) (*FederatedSettingsOrganizationConnections, *Response, error)
	Get(context.Context, string, string, string) (*FederatedSettingsOrganizationConnection, *Response, error)
	Update(context.Context, string, string, string, *FederatedSettingsOrganizationConnection) (*FederatedSettingsOrganizationConnection, *Response, error)
	Delete(context.Context, string, string, string) (*Response, error)
}

// FederatedSettingsIdentityProviderServiceOp handles communication with the FederatedSettings related methods of the
// MongoDB Atlas API.
type FederatedSettingsOrganizationConnectionSeviceOp service

var _ FederatedSettingsOrganizationConnectionService = &FederatedSettingsOrganizationConnectionSeviceOp{}

// A Resource describes a specific resource the Role will allow operating on.

// FederatedSettings represents a FederatedSettings Organization Connection..
type FederatedSettingsOrganizationConnections struct {
	Links      []*Link                                    `json:"links,omitempty"`
	Results    []*FederatedSettingsOrganizationConnection `json:"results,omitempty"`
	TotalCount int                                        `json:"totalCount,omitempty"`
}

type FederatedSettingsOrganizationConnection struct {
	ExternalGroupName string          `json:"externalGroupName,omitempty"`
	ID                string          `json:"id,omitempty"`
	RoleAssignments   RoleAssignments `json:"roleAssignments,omitempty"`
}

type RoleAssignments []struct {
	GroupID string `json:"groupId,omitempty"`
	OrgID   string `json:"orgId,omitempty"`
	Role    string `json:"role,omitempty"`
}

// List gets all Federated Settings Identity Providers for an organization.
//
// See more: https://www.mongodb.com/docs/atlas/reference/api/role-mapping-return-all/
func (s *FederatedSettingsOrganizationConnectionSeviceOp) List(ctx context.Context, opts *ListOptions, federationSettingsID, orgID string) (*FederatedSettingsOrganizationConnections, *Response, error) {

	basePath := fmt.Sprintf(federationSettingsOrganizationConnectionBasePath, federationSettingsID, orgID)
	path, err := setListOptions(basePath, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(FederatedSettingsOrganizationConnections)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root, resp, nil
}

// Get gets Federated Settings Identity Providers for an organization.
//
// See more: https://www.mongodb.com/docs/atlas/reference/api/role-mapping-return-one/
func (s *FederatedSettingsOrganizationConnectionSeviceOp) Get(ctx context.Context, federationSettingsID, orgID, roleMappingID string) (*FederatedSettingsOrganizationConnection, *Response, error) {
	if federationSettingsID == "" {
		return nil, nil, NewArgError("federationSettingsID", "must be set")
	}

	basePath := fmt.Sprintf(federationSettingsOrganizationConnectionBasePath, federationSettingsID, orgID)
	path := fmt.Sprintf("%s/%s", basePath, roleMappingID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(FederatedSettingsOrganizationConnection)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Update updates Federated Settings Identity Providers for an organization.
//
// See more: https://www.mongodb.com/docs/atlas/reference/api/role-mapping-create-one/
func (s *FederatedSettingsOrganizationConnectionSeviceOp) Update(ctx context.Context, federationSettingsID, orgID, roleMappingID string, updateRequest *FederatedSettingsOrganizationConnection) (*FederatedSettingsOrganizationConnection, *Response, error) {
	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	basePath := fmt.Sprintf(federationSettingsOrganizationConnectionBasePath, federationSettingsID, orgID)
	path := fmt.Sprintf("%s/%s", basePath, roleMappingID)

	req, err := s.Client.NewRequest(ctx, http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(FederatedSettingsOrganizationConnection)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Delete deletes federation setting.
//
// See more: https://www.mongodb.com/docs/atlas/reference/api/role-mapping-delete-one/
func (s *FederatedSettingsOrganizationConnectionSeviceOp) Delete(ctx context.Context, federationSettingsID string, orgID, roleMappingID string) (*Response, error) {
	if federationSettingsID == "" {
		return nil, NewArgError("federationSettingsID", "must be set")
	}

	basePath := fmt.Sprintf(federationSettingsOrganizationConnectionBasePath, federationSettingsID, orgID)
	path := fmt.Sprintf("%s/%s", basePath, roleMappingID)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}
