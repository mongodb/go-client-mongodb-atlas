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

// AdvancedClustersService is an interface for interfacing with the Clusters (Advanced)
// endpoints of the MongoDB Atlas API.
//
// See more: https://docs-atlas-staging.mongodb.com/cloud-docs/docsworker-xlarge/DOCSP-14695/reference/api/clusters-advanced/
type AdvancedClustersService interface {
	List(ctx context.Context, groupID string, options *ListOptions) (*AdvancedClustersResponse, *Response, error)
	Get(ctx context.Context, groupID, clusterName string) (*AdvancedCluster, *Response, error)
	Create(ctx context.Context, groupID string, cluster *AdvancedCluster) (*AdvancedCluster, *Response, error)
	Update(ctx context.Context, groupID, clusterName string, cluster *AdvancedCluster) (*AdvancedCluster, *Response, error)
}

// AdvancedClustersServiceOp handles communication with the Advanced Cluster related methods
// of the MongoDB Atlas API
type AdvancedClustersServiceOp service

var _ AdvancedClustersService = &AdvancedClustersServiceOp{}

// AdvancedCluster represents MongoDB cluster.
type AdvancedCluster struct {
	BaseCluster
	ReplicationSpecs []*AdvancedRegionSpec `json:"replicationSpecs,omitempty"`
	CreateDate       string                `json:"createDate,omitempty"`
	RootCertType     string                `json:"rootCertType,omitempty"`
	StateName        string                `json:"stateName,omitempty"`
}

type AdvancedRegionSpec struct {
	NumShards     int `json:"numShards,omitempty"`
	RegionConfigs []*AdvancedRegionConfig
}

type AdvancedRegionConfig struct {
	AnalyticsSpecs      *Specs `json:"analyticsSpecs,omitempty"`
	ElectableSpecs      *Specs `json:"electableSpecs,omitempty"`
	ReadOnlySpecs       *Specs `json:"readOnlySpecs,omitempty"`
	BackingProviderName string `json:"backingProviderName,omitempty"`
	Priority            int    `json:"priority,omitempty"`
	ProviderName        string `json:"providerName,omitempty"`
	RegionName          string `json:"regionName,omitempty"`
}

type AdvancedAutoScaling struct {
	DiskGB  *DiskGB  `json:"diskGB,omitempty"`
	Compute *Compute `json:"compute,omitempty"`
}

type DiskGB struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type Specs struct {
	DiskIOPS      float64 `json:"diskIOPS,omitempty"`
	EbsVolumeType string  `json:"ebsVolumeType,omitempty"`
	InstanceSize  string  `json:"instanceSize,omitempty"`
	NodeCount     int     `json:"nodeCount,omitempty"`
}

// AdvancedClustersResponse is the response from the ClustersService.List.
type AdvancedClustersResponse struct {
	Links      []*Link            `json:"links,omitempty"`
	Results    []*AdvancedCluster `json:"results,omitempty"`
	TotalCount int                `json:"totalCount,omitempty"`
}

// List all clusters in the project associated to {GROUP-ID}.
//
// See more: https://docs.atlas.mongodb.com/reference/api/clusters-get-all/
func (s *AdvancedClustersServiceOp) List(ctx context.Context, groupID string, listOptions *ListOptions) (*AdvancedClustersResponse, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}
	path := fmt.Sprintf(clustersPath, groupID)

	// Add query params from listOptions
	path, err := setListOptions(path, listOptions)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequestAndSetBaseUrl(ctx, ApiV15, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AdvancedClustersResponse)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root, resp, nil
}

// Get gets the cluster specified to {ClUSTER-NAME} from the project associated to {GROUP-ID}.
//
// See more: https://docs.atlas.mongodb.com/reference/api/clusters-get-one/
func (s *AdvancedClustersServiceOp) Get(ctx context.Context, groupID, clusterName string) (*AdvancedCluster, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}
	if err := checkClusterNameParam(clusterName); err != nil {
		return nil, nil, err
	}

	basePath := fmt.Sprintf(clustersPath, groupID)
	escapedEntry := url.PathEscape(clusterName)
	path := fmt.Sprintf("%s/%s", basePath, escapedEntry)

	req, err := s.Client.NewRequestAndSetBaseUrl(ctx, ApiV15, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(AdvancedCluster)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Create adds a cluster to the project associated to {GROUP-ID}.
//
// See more: https://docs.atlas.mongodb.com/reference/api/clusters-create-one/
func (s *AdvancedClustersServiceOp) Create(ctx context.Context, groupID string, createRequest *AdvancedCluster) (*AdvancedCluster, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := fmt.Sprintf(clustersPath, groupID)

	req, err := s.Client.NewRequestAndSetBaseUrl(ctx, ApiV15, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(AdvancedCluster)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Update a cluster in the project associated to {GROUP-ID}
//
// See more: https://docs.atlas.mongodb.com/reference/api/clusters-modify-one/
func (s *AdvancedClustersServiceOp) Update(ctx context.Context, groupID, clusterName string, updateRequest *AdvancedCluster) (*AdvancedCluster, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}
	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	basePath := fmt.Sprintf(clustersPath, groupID)
	path := fmt.Sprintf("%s/%s", basePath, clusterName)

	req, err := s.Client.NewRequestAndSetBaseUrl(ctx, ApiV15, http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(AdvancedCluster)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
