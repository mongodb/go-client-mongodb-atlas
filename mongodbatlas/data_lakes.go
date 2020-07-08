package mongodbatlas

import (
	"context"
	"fmt"
	"net/http"
)

const (
	dataLakesBasePath = "groups"
)

// DataLakeService is an interface for interfacing with the Data Lake endpoints of the MongoDB Atlas API.
// See more: https://docs.mongodb.com/datalake/reference/api/datalakes-api
type DataLakeService interface {
	List(context.Context, *DataLakeReqPathParameters, *ListOptions) ([]DataLake, *Response, error)
	Get(context.Context, *DataLakeReqPathParameters) (*DataLake, *Response, error)
	Create(context.Context, *DataLakeReqPathParameters, *DataLakeCreateRequest) (*DataLake, *Response, error)
	Update(context.Context, *DataLakeReqPathParameters, *DataLakeUpdateRequest) (*DataLake, *Response, error)
	Delete(context.Context, *DataLakeReqPathParameters) (*Response, error)
}

// DataLakeServiceOp handles communication with the DataLakeService related methods of the
// MongoDB Atlas API
type DataLakeServiceOp service

type AwsCloudProviderConfig struct {
	IAMAssumedRoleARN string `json:"iamAssumedRoleARN,omitempty"`
	TestS3Bucket      string `json:"testS3Bucket,omitempty"`
}

type CloudProviderConfig struct {
	AWSConfig AwsCloudProviderConfig `json:"aws,omitempty"`
}

type DataProcessRegion struct {
	CloudProvider string `json:"cloudProvider,omitempty"`
	Region        string `json:"region,omitempty"`
}

type DataLakeStore struct {
	Name        string `json:"name,omitempty"`
	Provider    string `json:"provider,omitempty"`
	Region      string `json:"region,omitempty"`
	Bucket      string `json:"bucket,omitempty"`
	Prefix      string `json:"prefix,omitempty"`
	Delimiter   string `json:"delimiter,omitempty"`
	IncludeTags bool   `json:"includeTags,omitempty"`
}

type DataLakeDataSource struct {
	StoreName     string `json:"storeName,omitempty"`
	DefaultFormat string `json:"defaultFormat,omitempty"`
	Path          string `json:"path,omitempty"`
}

type DataLakeCollection struct {
	Name        string               `json:"name,omitempty"`
	DataSources []DataLakeDataSource `json:"dataSources,omitempty"`
}

type DataLakeDatabaseView struct {
	Name     string `json:"name,omitempty"`
	Source   string `json:"source,omitempty"`
	Pipeline string `json:"pipeline,omitempty"`
}

type DataLakeDatabase struct {
	Name        string                 `json:"name,omitempty"`
	Collections []DataLakeCollection   `json:"collections,omitempty"`
	Views       []DataLakeDatabaseView `json:"views,omitempty"`
}

type Storage struct {
	Databases map[string]DataLakeDatabase `json:"databases,omitempty"`
	Stores    []DataLakeStore             `json:"stores,omitempty"`
}

// DataLake represents a data lake.
type DataLake struct {
	CloudProviderConfig CloudProviderConfig `json:"cloudProviderConfig,omitempty"` // Configuration for the cloud service where Data Lake source data is stored.
	DataProcessRegion   DataProcessRegion   `json:"dataProcessRegion,omitempty"`   // Cloud provider region which clients are routed to for data processing.
	GroupID             string              `json:"groupId,omitempty"`             // Unique identifier for the project.
	Hostnames           []string            `json:"hostnames,omitempty"`           // List of hostnames for the data lake.
	Name                string              `json:"name,omitempty"`                // Name of the data lake.
	State               string              `json:"state,omitempty"`               // Current state of the data lake.
	Storage             Storage             `json:"storage,omitempty"`             // Configuration for each data store and its mapping to MongoDB collections / databases.
}

// DataLakeReqPathParameters represents all the possible parameters to make the request
type DataLakeReqPathParameters struct {
	GroupID string `json:"groupId,omitempty"` // The unique identifier of the project for the Atlas cluster.
	Name    string `json:"name,omitempty"`    // The name of the Data Lake.
}

// DataLakeReqPathParameters represents all possible fields that can be updated in a data lake
type DataLakeUpdateRequest struct {
	CloudProviderConfig CloudProviderConfig `json:"cloudProviderConfig,omitempty"`
	DataProcessRegion   DataProcessRegion   `json:"dataProcessRegion,omitempty"`
}

// DataLakeReqPathParameters represents the required fields to create a new data lake
type DataLakeCreateRequest struct {
	Name string `json:"name,omitempty"`
}

// List gets all data lakes for the specified group.
// See more: https://docs.mongodb.com/datalake/reference/api/dataLakes-get-all-tenants
func (s *DataLakeServiceOp) List(ctx context.Context, requestParameters *DataLakeReqPathParameters, listOptions *ListOptions) ([]DataLake, *Response, error) {
	if requestParameters.GroupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}

	path := fmt.Sprintf("%s/%s/dataLakes", dataLakesBasePath, requestParameters.GroupID)

	path, err := setListOptions(path, listOptions)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new([]DataLake)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return *root, resp, nil
}

// Get gets the data laked associated with a specific name.
// See more: https://docs.mongodb.com/datalake/reference/api/dataLakes-get-one-tenant/
func (s *DataLakeServiceOp) Get(ctx context.Context, requestParameters *DataLakeReqPathParameters) (*DataLake, *Response, error) {
	if requestParameters.GroupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}
	if requestParameters.Name == "" {
		return nil, nil, NewArgError("name", "must be set")
	}

	path := fmt.Sprintf("%s/%s/dataLakes/%s", dataLakesBasePath, requestParameters.GroupID, requestParameters.Name)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(DataLake)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Create creates a new Data Lake.
// See more: https://docs.mongodb.com/datalake/reference/api/dataLakes-create-one-tenant/
func (s *DataLakeServiceOp) Create(ctx context.Context, requestParameters *DataLakeReqPathParameters, createRequest *DataLakeCreateRequest) (*DataLake, *Response, error) {
	if requestParameters.GroupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}
	if createRequest == nil {
		return nil, nil, NewArgError("name", "must be set")
	}

	path := fmt.Sprintf("%s/%s/dataLakes", dataLakesBasePath, requestParameters.GroupID)

	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(DataLake)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Update updates an existing Data Lake.
// See more: https://docs.mongodb.com/datalake/reference/api/dataLakes-update-one-tenant/
func (s *DataLakeServiceOp) Update(ctx context.Context, requestParameters *DataLakeReqPathParameters, updateRequest *DataLakeUpdateRequest) (*DataLake, *Response, error) {
	if requestParameters.GroupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}
	if requestParameters.Name == "" {
		return nil, nil, NewArgError("name", "must be set")
	}
	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s/%s/dataLakes/%s", dataLakesBasePath, requestParameters.GroupID, requestParameters.Name)

	req, err := s.Client.NewRequest(ctx, http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(DataLake)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Delete deletes the Data Lake with a given name.
// See more: https://docs.mongodb.com/datalake/reference/api/dataLakes-delete-one-tenant/
func (s *DataLakeServiceOp) Delete(ctx context.Context, requestParameters *DataLakeReqPathParameters) (*Response, error) {
	if requestParameters.GroupID == "" {
		return nil, NewArgError("groupId", "must be set")
	}
	if requestParameters.Name == "" {
		return nil, NewArgError("name", "must be set")
	}

	path := fmt.Sprintf("%s/%s/dataLakes/%s", dataLakesBasePath, requestParameters.GroupID, requestParameters.Name)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}
