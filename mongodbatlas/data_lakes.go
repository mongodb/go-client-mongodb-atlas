package mongodbatlas

import (
	"context"
	"fmt"
	"net/http"
)

const (
	dataLakesBasePath = "groups"
	Active = "ACTIVE"
	Unverified = "UNVERIFIED"
)

// DataLakeService is an interface for interfacing with the Data Lake endpoints of the MongoDB Atlas API.
// See more: https://docs.mongodb.com/datalake/reference/api/datalakes-api
type DataLakeService interface {
	GetAllDataLakes(context.Context, *DataLakeReqPathParameters, *ListOptions) ([]DataLake, *Response, error)
	GetOneDataLake(context.Context, *DataLakeReqPathParameters) (*DataLake, *Response, error)
	Create(context.Context, *DataLakeReqPathParameters, *DataLakeCreateRequest) (*DataLake, *Response, error)
	Update(context.Context, *DataLakeReqPathParameters, *DataLake) (*DataLake, *Response, error)
	Delete(context.Context, *DataLakeReqPathParameters) (*Response, error)
}

// DataLakeServiceOp handles communication with the DataLakeService related methods of the
// MongoDB Atlas API
type DataLakeServiceOp service

var _ DataLakeServiceOp = DataLakeServiceOp{}

type AwsCloudProviderConfig struct {
	IAMAssumedRoleARN string `json:"iamAssumedRoleARN,omitempty"`
	TestS3Bucket string `json:"testS3Bucket,omitempty"`
}

type CloudProviderConfig struct {
	AWSConfig AwsCloudProviderConfig `json:"aws,omitempty"`
}

type DataProcessRegion struct {
	CloudProvider string `json:"cloudProvider,omitempty"`
	Region string `json:"region,omitempty"`
}

type DataLakeStore struct {
	Name string `json:"cloudProvider,omitempty"`
	Provider string `json:"cloudProvider,omitempty"`
	Region string `json:"cloudProvider,omitempty"`
	Bucket string `json:"cloudProvider,omitempty"`
	Prefix string `json:"cloudProvider,omitempty"`
	Delimiter string `json:"cloudProvider,omitempty"`
	IncludeTags bool `json:"cloudProvider,omitempty"`
}

type Storage struct {
	Databases interface{} `json:"databases,omitempty"`
	Stores []DataLakeStore `json:"stores,omitempty"`
}

// DataLake represents a data lake.
type DataLake struct {
	CloudProviderConfig CloudProviderConfig  `json:"cloudProviderConfig,omitempty"`               // Configuration for the cloud service where Data Lake source data is stored.
	DataProcessRegion DataProcessRegion  `json:"dataProcessRegion,omitempty"`               // Cloud provider region which clients are routed to for data processing.
	GroupID string  `json:"groupId,omitempty"`               // Unique identifier for the project.
	Hostnames []string  `json:"hostnames,omitempty"`               // List of hostnames for the data lake.
	Name string  `json:"name,omitempty"`               // Name of the data lake.
	State string  `json:"state,omitempty"`               // Current state of the data lake.
	Storage Storage  `json:"storage,omitempty"`               // Configuration for each data store and its mapping to MongoDB collections / databases.
}

// DataLakeReqPathParameters represents all the possible parameters to make the request
type DataLakeReqPathParameters struct {
	GroupID     string `json:"groupId,omitempty"`     	// The unique identifier of the project for the Atlas cluster.
	Name 		string `json:"name,omitempty"` 			// The name of the Data Lake.
}

type DataLakeCreateRequest struct {
	Name		string `json:"name,omitempty"`
}

// GetAllDataLakes gets all data lakes for the specified group.
// See more: https://docs.mongodb.com/datalake/reference/api/dataLakes-get-all-tenants
func (s *DataLakeServiceOp) GetAllDataLakes(ctx context.Context, requestParameters *DataLakeReqPathParameters, listOptions *ListOptions) ([]DataLake, *Response, error) {
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

// GetOneDataLake gets the data laked associated with a specific name.
// See more: https://docs.mongodb.com/datalake/reference/api/dataLakes-get-one-tenant/
func (s *DataLakeServiceOp) GetOneDataLake(ctx context.Context, requestParameters *DataLakeReqPathParameters) (*DataLake, *Response, error) {
	if requestParameters.GroupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}
	if requestParameters.Name == "" {
		return nil, nil, NewArgError("dataLakeName", "must be set")
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
		return nil, nil, NewArgError("clusterName", "must be set")
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
func (s *DataLakeServiceOp) Update(ctx context.Context, requestParameters *DataLakeReqPathParameters, updateRequest *DataLake) (*DataLake, *Response, error) {
	if requestParameters.GroupID == "" {
		return nil, nil, NewArgError("groupId", "must be set")
	}
	if requestParameters.Name == "" {
		return nil, nil, NewArgError("clusterName", "must be set")
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
		return nil, NewArgError("clusterName", "must be set")
	}

	path := fmt.Sprintf("%s/%s/dataLakes/%s", dataLakesBasePath, requestParameters.GroupID, requestParameters.Name)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}
