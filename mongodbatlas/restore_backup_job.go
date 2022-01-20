// Copyright 2022 MongoDB Inc
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

const restoreClusterBackupJobPath = "api/atlas/v1.0/groups/%s/clusters/%s/backup/restoreJobs"

const restoreServerlessBackupJobPath = "api/atlas/v1.0/groups/%s/serverless/%s/backup/restoreJobs"

// RestoreBackupJobService provides access to the cloud restore backup jobs related functions in the Atlas API.
//
// See more: https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/restores/
type RestoreBackupJobService interface {
	List(context.Context, string, string, string, *ListOptions) (*RestoreBackupJobs, *Response, error)
	Get(context.Context, string, string, string, string) (*RestoreBackupJob, *Response, error)
	Create(context.Context, string, string, *RestoreBackupJobRequest) (*RestoreBackupJob, *Response, error)
	Delete(context.Context, string, string, string) (*Response, error)
}

// RestoreBackupJobServiceOp handles communication with the Cloud Backup Restore Jobs related methods
// of the MongoDB Atlas API.
type RestoreBackupJobServiceOp service

// List Get All Cloud Backup Restore Jobs.
//
// See more: https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-all-restore-jobs/
// See more: https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/return-all-restore-jobs-for-one-serverless-instance/
func (r RestoreBackupJobServiceOp) List(ctx context.Context, projectID string, instanceName string, deploymentType string, options *ListOptions) (*RestoreBackupJobs, *Response, error) {
	if projectID == "" {
		return nil, nil, NewArgError("projectID", "must be set")
	}
	if instanceName == "" {
		return nil, nil, NewArgError("instanceName", "must be set")
	}

	var path string

	switch deploymentType {
	case "SERVERLESS":
		path = fmt.Sprintf(restoreServerlessBackupJobPath, projectID, instanceName)
	case "DEDICATED":
		path = fmt.Sprintf(restoreClusterBackupJobPath, projectID, instanceName)
	default:
		return nil, nil, fmt.Errorf("deployment type: %s not supported, choose SERVERLESS or DEDICATED", deploymentType)
	}

	path, err := setListOptions(path, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := r.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var root RestoreBackupJobs
	resp, err := r.Client.Do(ctx, req, &root)

	return &root, resp, err
}

// Get One Cloud Backup Restore Job.
//
// See more: https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-one-restore-job/
// See more: https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/return-one-restore-job-for-one-serverless-instance/
func (r RestoreBackupJobServiceOp) Get(ctx context.Context, projectID string, instanceName string, restoreJobID string, deploymentType string) (*RestoreBackupJob, *Response, error) {
	if projectID == "" {
		return nil, nil, NewArgError("projectID", "must be set")
	}
	if instanceName == "" {
		return nil, nil, NewArgError("instanceName", "must be set")
	}
	if restoreJobID == "" {
		return nil, nil, NewArgError("restoreJobID", "must be set")
	}

	var path string

	switch deploymentType {
	case "SERVERLESS":
		path = fmt.Sprintf(restoreServerlessBackupJobPath, projectID, instanceName)
	case "DEDICATED":
		path = fmt.Sprintf(restoreClusterBackupJobPath, projectID, instanceName)
	default:
		return nil, nil, fmt.Errorf("deployment type: %s not supported, choose SERVERLESS or DEDICATED", deploymentType)
	}

	path = fmt.Sprintf("%s/%s", path, restoreJobID)

	req, err := r.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var responseObj *RestoreBackupJob
	resp, err := r.Client.Do(ctx, req, &responseObj)

	return responseObj, resp, err
}

// Create a Cloud Backup Restore Job.
//
// See more: https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/create-one-restore-job/
// See more: https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/restore-one-snapshot-of-one-serverless-instance/
func (r RestoreBackupJobServiceOp) Create(ctx context.Context, projectID string, instanceName string, request *RestoreBackupJobRequest) (*RestoreBackupJob, *Response, error) {
	if projectID == "" {
		return nil, nil, NewArgError("projectID", "must be set")
	}

	if instanceName == "" {
		return nil, nil, NewArgError("instanceName", "must be set")
	}

	var path string

	switch request.DeploymentType {
	case "SERVERLESS":
		path = fmt.Sprintf(restoreServerlessBackupJobPath, projectID, instanceName)
	case "DEDICATED":
		path = fmt.Sprintf(restoreClusterBackupJobPath, projectID, instanceName)
	default:
		return nil, nil, fmt.Errorf("deployment type: %s not supported, choose SERVERLESS or DEDICATED", request.DeploymentType)
	}

	req, err := r.Client.NewRequest(ctx, http.MethodPost, path, request)
	if err != nil {
		return nil, nil, err
	}

	responseObj := new(RestoreBackupJob)
	resp, err := r.Client.Do(ctx, req, responseObj)
	if err != nil {
		return nil, resp, err
	}

	return responseObj, resp, err
}

// Delete Cancel One Cloud Backup Restore Job.
//
// See more: https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/delete-one-restore-job/
func (r RestoreBackupJobServiceOp) Delete(ctx context.Context, projectID string, instanceName string, restoreJobID string) (*Response, error) {
	if projectID == "" {
		return nil, NewArgError("projectID", "must be set")
	}

	if instanceName == "" {
		return nil, NewArgError("instanceName", "must be set")
	}

	if restoreJobID == "" {
		return nil, NewArgError("restoreJobID", "must be set")
	}

	path := fmt.Sprintf(restoreClusterBackupJobPath, projectID, instanceName)
	path = fmt.Sprintf("%s/%s", path, restoreJobID)

	req, err := r.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := r.Client.Do(ctx, req, nil)

	return resp, err
}

var _ RestoreBackupJobService = &RestoreBackupJobServiceOp{}

type RestoreBackupJob struct {
	DeliveryType          string                       `json:"deliveryType"`
	SnapshotID            string                       `json:"snapshotId"`
	TargetGroupID         string                       `json:"targetGroupId"`
	DeploymentType        string                       `json:"deploymentType"`
	ID                    string                       `json:"id,omitempty"`
	Failed *bool `json:"failed,omitempty"`
	SourceClusterName     string                       `json:"sourceClusterName,omitempty"`
	Cancelled             *bool                        `json:"cancelled,omitempty"`
	CreatedAt             string                       `json:"createdAt,omitempty"`
	Expired               *bool                        `json:"expired,omitempty"`
	ExpiresAt             string                       `json:"expiresAt,omitempty"`
	FinishedAt            string                       `json:"finishedAt,omitempty"`
	Links                 []Link                       `json:"links,omitempty"`
	Timestamp             string                       `json:"timestamp,omitempty"`
	Components            []*RestoreBackupJobComponent `json:"components,omitempty"`
	DeliveryURL           []string                     `json:"deliveryUrl,omitempty"`
	OplogTs               int                          `json:"oplogTs,omitempty"`
	OplogInc              int                          `json:"oplogInc,omitempty"`
	PointInTimeUTCSeconds int                          `json:"pointInTimeUtcSeconds,omitempty"`
	TargetClusterName     string                       `json:"targetClusterName,omitempty"`
}

type RestoreBackupJobRequest struct {
	DeliveryType      string `json:"deliveryType"`
	SnapshotID        string `json:"snapshotId"`
	TargetClusterName string `json:"targetClusterName"`
	TargetGroupID     string `json:"targetGroupId"`
	DeploymentType    string `json:"deploymentType"`
}

type RestoreBackupJobs struct {
	Results    []*RestoreBackupJob `json:"results,omitempty"`
	Links      []*Link             `json:"links,omitempty"`
	TotalCount int64               `json:"totalCount,omitempty"`
}

type RestoreBackupJobComponent struct {
	DownloadURL    string `json:"downloadURL,omitempty"`
	ReplicaSetName string `json:"replicaSetName,omitempty"`
}
