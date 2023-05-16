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
	"io"
	"net/http"
)

const (
	logsPath     = "api/atlas/v1.0/groups/%s/clusters/%s/logs/%s"
	logsJobsPath = "api/atlas/v1.0/groups/%s/logCollectionJobs"
)

// LogsService is an interface for interfacing with the Logs
// endpoints of the MongoDB Atlas API.
// See more: https://docs.atlas.mongodb.com/reference/api/logs/
type LogsService interface {
	Get(context.Context, string, string, string, io.Writer, *DateRangetOptions) (*Response, error)
	Download(context.Context, string, string, io.Writer) (*Response, error)
}

// LogsServiceOp handles communication with the Logs related methods of the
// MongoDB Atlas API.
type LogsServiceOp struct {
	Client GZipRequestDoer
}

// LogCollectionService is an interface for interfacing with the Log Collection Jobs
// endpoints of the Atlas Admin API. These endpoints are not publicly documented, as
// they are only available to MongoDB employees with Global access.
//
// See more: https://wiki.corp.mongodb.com/display/MMS/Atlas+Private+API#AtlasPrivateAPI-GatherFTDCData(LogCollectionJobs)
type LogCollectionService interface {
	List(context.Context, string, *LogListOptions) (*LogCollectionJobs, *Response, error)
	Get(context.Context, string, string, *LogListOptions) (*LogCollectionJob, *Response, error)
	Retry(context.Context, string, string) (*Response, error)
	Create(context.Context, string, *LogCollectionJob) (*LogCollectionJob, *Response, error)
	Extend(context.Context, string, string, *LogCollectionJob) (*Response, error)
	Delete(context.Context, string, string) (*Response, error)
}

// LogCollectionServiceOp provides an implementation of the LogCollectionService interface.
type LogCollectionServiceOp service

var _ LogCollectionService = &LogCollectionServiceOp{}

// DateRangetOptions specifies an optional date range query.
type DateRangetOptions struct {
	StartDate string `url:"startDate,omitempty"`
	EndDate   string `url:"endDate,omitempty"`
}

// Get gets a compressed (.gz) log file that contains a range of log messages for a particular host.
// Note: The input parameter out (io.Writer) is not closed by this function.
//
// See more: https://docs.atlas.mongodb.com/reference/api/logs/
func (s *LogsServiceOp) Get(ctx context.Context, groupID, hostName, logName string, out io.Writer, opts *DateRangetOptions) (*Response, error) {
	if groupID == "" {
		return nil, NewArgError("groupID", "must be set")
	}

	if hostName == "" {
		return nil, NewArgError("hostName", "must be set")
	}

	if logName == "" {
		return nil, NewArgError("logName", "must be set")
	}

	basePath := fmt.Sprintf(logsPath, groupID, hostName, logName)

	// Add query params
	path, err := setListOptions(basePath, opts)
	if err != nil {
		return nil, err
	}

	req, err := s.Client.NewGZipRequest(ctx, http.MethodGet, path)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, out)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// LogCollectionJob represents a Log Collection Job in the MongoDB Atlas API.
type LogCollectionJob struct {
	ID                         string      `json:"id,omitempty"`
	GroupID                    string      `json:"groupId,omitempty"`
	UserID                     string      `json:"userId,omitempty"`
	CreationDate               string      `json:"creationDate,omitempty"`
	ExpirationDate             string      `json:"expirationDate,omitempty"`
	Status                     string      `json:"status,omitempty"`
	ResourceType               string      `json:"resourceType,omitempty"`
	ResourceName               string      `json:"resourceName,omitempty"`
	RootResourceName           string      `json:"rootResourceName,omitempty"`
	RootResourceType           string      `json:"rootResourceType,omitempty"`
	DownloadURL                string      `json:"downloadUrl,omitempty"`
	Redacted                   *bool       `json:"redacted,omitempty"`
	LogTypes                   []string    `json:"logTypes,omitempty"`
	SizeRequestedPerFileBytes  int64       `json:"sizeRequestedPerFileBytes,omitempty"`
	UncompressedSizeTotalBytes int64       `json:"uncompressedSizeTotalBytes,omitempty"`
	ChildJobs                  []*ChildJob `json:"childJobs,omitempty"` // ChildJobs included if verbose is true
}

// ChildJob represents a ChildJob in the MongoDB Ops Manager API.
type ChildJob struct {
	AutomationAgentID          string `json:"automationAgentId,omitempty"`
	ErrorMessage               string `json:"errorMessage,omitempty"`
	FinishDate                 string `json:"finishDate"`
	HostName                   string `json:"hostName"`
	LogCollectionType          string `json:"logCollectionType"`
	Path                       string `json:"path"`
	StartDate                  string `json:"startDate"`
	Status                     string `json:"status"`
	UncompressedDiskSpaceBytes int64  `json:"uncompressedDiskSpaceBytes"`
}

// LogCollectionJobs represents a array of LogCollectionJobs.
type LogCollectionJobs struct {
	Links      []*Link             `json:"links"`
	Results    []*LogCollectionJob `json:"results"`
	TotalCount int                 `json:"totalCount"`
}

// LogListOptions specifies the optional parameters to List methods that
// support pagination.
type LogListOptions struct {
	ListOptions
	Verbose bool `url:"verbose,omitempty"`
}

// List gets all collection log jobs.
//
// See more: https://wiki.corp.mongodb.com/display/MMS/Atlas+Private+API#AtlasPrivateAPI-GatherFTDCData(LogCollectionJobs)
func (s *LogCollectionServiceOp) List(ctx context.Context, groupID string, opts *LogListOptions) (*LogCollectionJobs, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}

	basePath := fmt.Sprintf(logsJobsPath, groupID)
	path, err := setListOptions(basePath, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(LogCollectionJobs)
	resp, err := s.Client.Do(ctx, req, root)

	return root, resp, err
}

// Get gets a log collection job.
//
// See more: https://wiki.corp.mongodb.com/display/MMS/Atlas+Private+API#AtlasPrivateAPI-GatherFTDCData(LogCollectionJobs)
func (s *LogCollectionServiceOp) Get(ctx context.Context, groupID, jobID string, opts *LogListOptions) (*LogCollectionJob, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}

	if jobID == "" {
		return nil, nil, NewArgError("jobID", "must be set")
	}

	basePath := fmt.Sprintf(logsJobsPath, groupID)
	path := fmt.Sprintf("%s/%s", basePath, jobID)

	path, err := setListOptions(path, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(LogCollectionJob)
	resp, err := s.Client.Do(ctx, req, root)

	return root, resp, err
}

// Create creates a log collection job.
//
// See more: https://wiki.corp.mongodb.com/display/MMS/Atlas+Private+API#AtlasPrivateAPI-GatherFTDCData(LogCollectionJobs)
func (s *LogCollectionServiceOp) Create(ctx context.Context, groupID string, log *LogCollectionJob) (*LogCollectionJob, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}

	if log == nil {
		return nil, nil, NewArgError("log", "must be set")
	}

	path := fmt.Sprintf(logsJobsPath, groupID)
	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, log)
	if err != nil {
		return nil, nil, err
	}

	root := new(LogCollectionJob)
	resp, err := s.Client.Do(ctx, req, root)

	return root, resp, err
}

// Extend extends the expiration data of a log collection job.
//
// See more: https://wiki.corp.mongodb.com/display/MMS/Atlas+Private+API#AtlasPrivateAPI-GatherFTDCData(LogCollectionJobs)
func (s *LogCollectionServiceOp) Extend(ctx context.Context, groupID, jobID string, log *LogCollectionJob) (*Response, error) {
	if groupID == "" {
		return nil, NewArgError("groupID", "must be set")
	}

	if jobID == "" {
		return nil, NewArgError("jobID", "must be set")
	}

	if log == nil {
		return nil, NewArgError("log", "must be set")
	}

	basePath := fmt.Sprintf(logsJobsPath, groupID)
	path := fmt.Sprintf("%s/%s", basePath, jobID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, path, log)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}

// Retry retries a single failed log collection job.
//
// See more: https://docs.opsmanager.mongodb.com/current/reference/api/log-collections/log-collections-update-one/
func (s *LogCollectionServiceOp) Retry(ctx context.Context, groupID, jobID string) (*Response, error) {
	if groupID == "" {
		return nil, NewArgError("groupID", "must be set")
	}

	if jobID == "" {
		return nil, NewArgError("jobID", "must be set")
	}

	basePath := fmt.Sprintf(logsJobsPath, groupID)
	path := fmt.Sprintf("%s/%s/retry", basePath, jobID)

	req, err := s.Client.NewRequest(ctx, http.MethodPut, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}

// Delete removes a log collection job.
//
// See more: https://docs.opsmanager.mongodb.com/current/reference/api/log-collections/log-collections-retry/
func (s *LogCollectionServiceOp) Delete(ctx context.Context, groupID, jobID string) (*Response, error) {
	if groupID == "" {
		return nil, NewArgError("groupID", "must be set")
	}

	if jobID == "" {
		return nil, NewArgError("jobID", "must be set")
	}

	basePath := fmt.Sprintf(logsJobsPath, groupID)
	path := fmt.Sprintf("%s/%s", basePath, jobID)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}

// Download allows to download a log from a collection job.
//
// See more: https://docs.opsmanager.mongodb.com/current/reference/api/log-collections/log-collections-retry/
func (s *LogsServiceOp) Download(ctx context.Context, groupID, jobID string, out io.Writer) (*Response, error) {
	if groupID == "" {
		return nil, NewArgError("groupID", "must be set")
	}

	if jobID == "" {
		return nil, NewArgError("jobID", "must be set")
	}

	basePath := fmt.Sprintf(logsJobsPath, groupID)
	path := fmt.Sprintf("%s/%s/download", basePath, jobID)

	req, err := s.Client.NewGZipRequest(ctx, http.MethodGet, path)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, out)

	return resp, err
}
