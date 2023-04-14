package sdk

import (
	"errors"

	mongodbatlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

// APIError gets a strongly typed error from an error.
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

// IsAPIError returns true if the error contains the errCode.
// Error code is an code that is returned by the API.
func IsAPIError(err error, code string) bool {
	mappedErr := GetAPIError(err)
	if mappedErr == nil {
		return false
	}

	return mappedErr.GetErrorCode() == code
}
