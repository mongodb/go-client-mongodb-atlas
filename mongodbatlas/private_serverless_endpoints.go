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
	"net/url"
)

const (
	serverlessPrivateEndpointsPath = "api/atlas/v1.0/groups/%s/privateEndpoint/serverless/instance/%s/endpoint"
)

// PrivateEndpointsService is an interface for interfacing with the Private Endpoints
// of the MongoDB Atlas API.
//
// See more: https://docs.atlas.mongodb.com/reference/api/private-endpoints/
type PrivateServerlessEndpointsService interface {
	ListPrivateServerlessEndpoint(context.Context, string, string, *ListOptions) ([]PrivateServerlessEndpointConnection, *Response, error)
	AddOnePrivateServerlessEndpoint(context.Context, string, string, *PrivateServerlessEndpointConnection) (*PrivateServerlessEndpointConnection, *Response, error)
	GetOnePrivateServerlessEndpoint(context.Context, string, string, string) (*PrivateServerlessEndpointConnection, *Response, error)
	DeleteOnePrivateServerlessEndpoint(context.Context, string, string, string) (*Response, error)
	UpdateOnePrivateServerlessEndpoint(context.Context, string, string, string, *PrivateServerlessEndpointConnection) (*PrivateServerlessEndpointConnection, *Response, error)
}

// PrivateServerlessEndpointsServiceOp handles communication with the PrivateServerlessEndpoints related methods
// of the MongoDB Atlas API.
type PrivateServerlessEndpointsServiceOp service

var _ PrivateServerlessEndpointsService = &PrivateServerlessEndpointsServiceOp{}

// PrivateEndpointServerlessConnection represents MongoDB Private Endpoint Connection.
type PrivateServerlessEndpointConnection struct {
	ID                      string `json:"_id,omitempty"` // Unique identifier of the Serverless PrivateLink Service.
	CloudProviderEndpointID string `json:"cloudProviderEndpointId,omitempty"`
	Comment                 string `json:"comment,omitempty"`
	EndpointServiceName     string `json:"endpointServiceName,omitempty"` // Name of the PrivateLink endpoint service in AWS. Returns null while the endpoint service is being created.
	ErrorMessage            string `json:"errorMessage,omitempty"`        // Error message pertaining to the AWS Service Connect. Returns null if there are no errors.
	Status                  string `json:"status,omitempty"`              // Status of the AWS Serverless PrivateLink connection: INITIATING, WAITING_FOR_USER, FAILED, DELETING, AVAILABLE.
	ProviderName            string `json:"providerName,omitempty"`        // Human-readable label that identifies the cloud provider. Values include AWS or AZURE. Atlas currently supports only AWS.
}

// List retrieve details for all private Serverless endpoint services in one Atlas project.
//
// See more: https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/returnAllPrivateEndpointsForOneServerlessInstance
func (s *PrivateServerlessEndpointsServiceOp) ListPrivateServerlessEndpoint(ctx context.Context, groupID, instanceID string, listOptions *ListOptions) ([]PrivateServerlessEndpointConnection, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}
	if instanceID == "" {
		return nil, nil, NewArgError("instanceID", "must be set")
	}

	path := fmt.Sprintf(serverlessPrivateEndpointsPath, groupID, instanceID) // Add query params from listOptions
	path, err := setListOptions(path, listOptions)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new([]PrivateServerlessEndpointConnection)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return *root, resp, nil
}

// DeleteOnePrivateServerlessEndpoint one private serverless endpoint service in an Atlas project.
//
// See more https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/removeOnePrivateEndpointFromOneServerlessInstance
func (s *PrivateServerlessEndpointsServiceOp) DeleteOnePrivateServerlessEndpoint(ctx context.Context, groupID, instanceID, privateEndpointID string) (*Response, error) {
	if groupID == "" {
		return nil, NewArgError("groupID", "must be set")
	}
	if privateEndpointID == "" {
		return nil, NewArgError("PrivateEndpointID", "must be set")
	}
	if instanceID == "" {
		return nil, NewArgError("instanceID", "must be set")
	}

	basePath := fmt.Sprintf(serverlessPrivateEndpointsPath, groupID, instanceID)
	path := fmt.Sprintf("%s/%s", basePath, url.PathEscape(privateEndpointID))

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.Client.Do(ctx, req, nil)
}

// AddOnePrivateServerlessEndpoint Adds one serverless  private endpoint in an Atlas project.
//
// See more: https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/createOnePrivateEndpointForOneServerlessInstance
func (s *PrivateServerlessEndpointsServiceOp) AddOnePrivateServerlessEndpoint(ctx context.Context, groupID, instanceID string, createRequest *PrivateServerlessEndpointConnection) (*PrivateServerlessEndpointConnection, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}

	if instanceID == "" {
		return nil, nil, NewArgError("instanceID", "must be set")
	}
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := fmt.Sprintf(serverlessPrivateEndpointsPath, groupID, instanceID)

	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(PrivateServerlessEndpointConnection)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetOnePrivateServerlessEndpoint retrieve details for one private serverless endpoint in an Atlas project.
//
// See more: https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/returnOnePrivateEndpointForOneServerlessInstance
func (s *PrivateServerlessEndpointsServiceOp) GetOnePrivateServerlessEndpoint(ctx context.Context, groupID, instanceID, privateEndpointID string) (*PrivateServerlessEndpointConnection, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}

	if instanceID == "" {
		return nil, nil, NewArgError("instanceID", "must be set")
	}
	if privateEndpointID == "" {
		return nil, nil, NewArgError("privateEndpointID", "must be set")
	}

	basePath := fmt.Sprintf(serverlessPrivateEndpointsPath, groupID, instanceID)
	path := fmt.Sprintf("%s/%s", basePath, url.PathEscape(privateEndpointID))

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PrivateServerlessEndpointConnection)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UpdateOnePrivateServerlessEndpoint updates the private serverless endpoint setting for one Atlas project.
//
// See more: https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/updateOnePrivateEndpointForOneServerlessInstance
func (s *PrivateServerlessEndpointsServiceOp) UpdateOnePrivateServerlessEndpoint(ctx context.Context, groupID, instanceID, privateEndpointID string, updateRequest *PrivateServerlessEndpointConnection) (*PrivateServerlessEndpointConnection, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}
	if instanceID == "" {
		return nil, nil, NewArgError("instanceID", "must be set")
	}
	if privateEndpointID == "" {
		return nil, nil, NewArgError("privateEndpointID", "must be set")
	}

	basePath := fmt.Sprintf(serverlessPrivateEndpointsPath, groupID, instanceID)
	path := fmt.Sprintf("%s/%s", basePath, url.PathEscape(privateEndpointID))
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(PrivateServerlessEndpointConnection)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
