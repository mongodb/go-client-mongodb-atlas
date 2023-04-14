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

// IsErrorCode returns true if the error contains the errCode.
// Error code is an code that is returned by the API.
func IsErrorCode(err error, code string) bool {
	mappedErr := APIError(err)
	if mappedErr == nil {
		return false
	}

	return mappedErr.GetErrorCode() == code
}
