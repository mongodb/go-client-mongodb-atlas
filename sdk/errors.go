package sdk

import (
	"errors"
	"fmt"

	mongodbatlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

// APIError checks if error is an error returned by API.
func APIError(err error) *mongodbatlasv2.Error {
	var openapiError mongodbatlasv2.GenericOpenAPIError

	if ok := errors.As(err, &openapiError); !ok {
		return nil
	}
	errModel := openapiError.Model()
	transformedError, ok := errModel.(mongodbatlasv2.Error)
	if !ok {
		return nil
	}
	return &transformedError
}

// APIErrorToErr returns APIError as instance of Go error or nil if passed error is not an API error.
func APIErrorToErr(err error) error {
	apiError := APIError(err)
	if apiError != nil {
		return fmt.Errorf("%v", apiError)
	}
	return nil
}

// IsErrorCode returns true if the error contains the errCode.
// Error code is an code that is returned by the API.
func IsErrorCode(err error, code string) bool {
	mappedErr := APIError(err)
	if mappedErr == nil {
		return false
	}

	return mappedErr.GetErrorCode() == code
}
