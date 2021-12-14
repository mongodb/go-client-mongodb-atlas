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

const (
	privateLinkEndpointsADLPath = "api/atlas/v1.0/groups/%s/privateNetworkSettings/endpointIds"
)

// PrivateLinkEndpointsADLService is an interface for interfacing with the Private Link Endpoints ADL
// of the MongoDB Atlas API datalake and online archive.
//
// See more: https://docs.atlas.mongodb.com/reference/api/private-endpoints/
type PrivateLinkEndpointsADLService interface {
	Create(context.Context, string, *PrivateLinkEndpointADL) (*PrivateLinkEndpointADLResponse, *Response, error)
	Get(context.Context, string, string) (*PrivateLinkEndpointADL, *Response, error)
	List(context.Context, string) (*PrivateLinkEndpointADLResponse, *Response, error)
	Delete(context.Context, string, string) (*Response, error)
}

// PrivateLinkEndpointsADLServiceOp handles communication with the PrivateLinkEndpointsADL related methods
// of the MongoDB Atlas API.
type PrivateLinkEndpointsADLServiceOp service

var _ PrivateLinkEndpointsADLService = &PrivateLinkEndpointsADLServiceOp{}

// PrivateLinkEndpointADLResponse represents MongoDB Private Endpoint Connection to DataLake and Online Archive.
type PrivateLinkEndpointADLResponse struct {
	Links      []*Link                   `json:"links,omitempty"`
	Results    []*PrivateLinkEndpointADL `json:"results"`
	TotalCount int                       `json:"totalCount"`
}

//PrivateLinkEndpointADL represents the private link result for data lake and online archive.
type PrivateLinkEndpointADL struct {
	Comment    string `json:"comment,omitempty"`
	EndpointID string `json:"endpointId,omitempty"`
	Provider   string `json:"provider,omitempty"`
	Type       string `json:"type,omitempty"`
}

//Create creates one private link endpoint in Data Lake and Online Archive Atlas project.
//
// See more: https://docs.mongodb.com/datalake/reference/api/dataLakes-private-link-create-one/#std-label-api-pvt-link-create-one
func (s *PrivateLinkEndpointsADLServiceOp) Create(ctx context.Context, groupID string, createRequest *PrivateLinkEndpointADL) (*PrivateLinkEndpointADLResponse, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "must be set")
	}

	path := fmt.Sprintf(privateLinkEndpointsADLPath, groupID)

	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(PrivateLinkEndpointADLResponse)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Delete deletes the Data Lake and Online Archive private link endpoint with a given endpoint id.
//
// See more: https://docs.mongodb.com/datalake/reference/api/dataLakes-private-link-delete-one/#std-label-api-pvt-link-delete-one
func (s *PrivateLinkEndpointsADLServiceOp) Delete(ctx context.Context, groupID, endpointID string) (*Response, error) {
	if groupID == "" {
		return nil, NewArgError("groupId", "must be set")
	}
	if endpointID == "" {
		return nil, NewArgError("endpointID", "must be set")
	}

	path := fmt.Sprintf("%s/%s", fmt.Sprintf(privateLinkEndpointsADLPath, groupID), endpointID)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}

// List gets all private link endpoints for data lake and online archive for the specified group.
//
// See more: https://docs.atlas.mongodb.com/reference/api/online-archive-private-link-get-all/#std-label-api-online-archive-pvt-link-get-all
func (s *PrivateLinkEndpointsADLServiceOp) List(ctx context.Context, groupID string) (*PrivateLinkEndpointADLResponse, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}

	path := fmt.Sprintf(privateLinkEndpointsADLPath, groupID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var root = new(PrivateLinkEndpointADLResponse)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return root, resp, err
	}

	return root, resp, nil
}

// Get gets the data laked and online archive private link endpoint associated with a specific group and endpointID.
//
// See more: https://docs.mongodb.com/datalake/reference/api/dataLakes-private-link-get-one/#std-label-api-pvt-link-get-one
func (s *PrivateLinkEndpointsADLServiceOp) Get(ctx context.Context, groupID, endpointID string) (*PrivateLinkEndpointADL, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}
	if endpointID == "" {
		return nil, nil, NewArgError("endpointID", "must be set")
	}

	path := fmt.Sprintf("%s/%s", fmt.Sprintf(privateLinkEndpointsADLPath, groupID), endpointID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(PrivateLinkEndpointADL)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
