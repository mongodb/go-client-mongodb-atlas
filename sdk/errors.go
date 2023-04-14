package sdk

import (
	"errors"

	mongodbatlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

// GetAPIError gets a strongly typed error from an error
func GetAPIError(err error) *mongodbatlasv2.Error {
	var openapiError mongodbatlasv2.GenericOpenAPIError

	if ok := errors.As(err, &openapiError); ok {
		errModel := openapiError.Model()

		transformedError, ok := errModel.(mongodbatlasv2.Error)
		if !ok {
			return nil
		}
		return &transformedError
	}

	return nil
}

// IsAPIError returns true if the error contains the errCode
// Error code is an code that is returned by the API
func IsAPIError(err error, code string) bool {
	mappedErr := GetAPIError(err)
	if mappedErr == nil {
		return false
	}

	return mappedErr.GetErrorCode() == string(code)
}
