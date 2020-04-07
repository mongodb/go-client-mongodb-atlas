package mongodbatlas

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

const logsPath = "groups/%s/clusters/%s/logs/%s"

// LogsService is an interface for interfacing with the Logs
// endpoints of the MongoDB Atlas API.
// See more: https://docs.atlas.mongodb.com/reference/api/process-measurements/
type LogsService interface {
	Get(context.Context, string, string, string, *LogsListOptions) (*Response, error)
}

// LogsServiceOp handles communication with the Logs related methods of the
// MongoDB Atlas API
type LogsServiceOp struct {
	Client RequestDoer
}

// EventListOptions specifies the optional parameters to the Event List methods.
type LogsListOptions struct {
	StartDate string `url:"startDate,omitempty"`
	EndDate   string `url:"endDate,omitempty"`
}

// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// Get gets a compressed (.gz) log file that contains a range of log messages for a particular host in an Atlas cluster.
// See more: https://docs.atlas.mongodb.com/reference/api/logs/
func (s *LogsServiceOp) Get(ctx context.Context, groupID string, hostName string, logName string, opts *LogsListOptions) (*Response, error) {

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

	//Add query params from listOptions
	path, err := setListOptions(basePath, opts)
	if err != nil {
		return nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
