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
	cloudProviderRegionsBasePath = "groups/%s/clusters/provider/regions"
)

// CloudProviderRegionsService is an interface for interfacing with the Cloud Provider Regions
// endpoints of the MongoDB Atlas API.
type CloudProviderRegionsService interface {
	List(context.Context, string, *CloudProviderRegionsOptions) (*CloudProviders, *Response, error)
}

// CloudProviderRegionsServiceOp handles communication with the CloudProviderRegionsService related methods of the
// MongoDB Atlas API
type CloudProviderRegionsServiceOp service

var _ CloudProviderRegionsService = &CloudProviderRegionsServiceOp{}

// CloudProvider represents a cloud provider of the MongoDB Atlas API
type CloudProvider struct {
	Provider      string          `json:"provider,omitempty"`
	InstanceSizes []*InstanceSize `json:"instanceSizes,omitempty"`
}

// InstanceSize represents an instance size of the MongoDB Atlas API
type InstanceSize struct {
	Name             string             `json:"name,omitempty"`
	AvailableRegions []*AvailableRegion `json:"availableRegions,omitempty"`
}

// AvailableRegion represents an available region of the MongoDB Atlas API
type AvailableRegion struct {
	Name    string `json:"name,omitempty"`
	Default bool   `json:"default,omitempty"`
}

// CloudProviders represents the response from CloudProviderRegionsService.Get
type CloudProviders struct {
	Links      []*Link          `json:"links,omitempty"`
	Results    []*CloudProvider `json:"results,omitempty"`
	TotalCount int              `json:"totalCount,omitempty"`
}

// CloudProviderRegionsOptions specifies the optional parameters to the CloudProviderRegions Get method.
type CloudProviderRegionsOptions struct {
	Provider           string `url:"provider,omitempty"`
	Tier               string `url:"tier,omitempty"`
	CrossCloudProvider bool   `url:"maxDate"`
}

// List gets the available regions for each cloud provider
func (s *CloudProviderRegionsServiceOp) List(ctx context.Context, groupID string, options *CloudProviderRegionsOptions) (*CloudProviders, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}

	path := fmt.Sprintf(cloudProviderRegionsBasePath, groupID)

	path, err := setListOptions(path, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(CloudProviders)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
