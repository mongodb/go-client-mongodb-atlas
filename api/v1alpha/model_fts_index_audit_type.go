/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"fmt"
)

// FTSIndexAuditType Unique identifier of event type.
type FTSIndexAuditType string

// List of FTSIndexAuditType
const (
	FTSINDEXAUDITTYPE_DELETION_FAILED FTSIndexAuditType = "FTS_INDEX_DELETION_FAILED"
	FTSINDEXAUDITTYPE_BUILD_COMPLETE FTSIndexAuditType = "FTS_INDEX_BUILD_COMPLETE"
	FTSINDEXAUDITTYPE_BUILD_FAILED FTSIndexAuditType = "FTS_INDEX_BUILD_FAILED"
	FTSINDEXAUDITTYPE_CREATED FTSIndexAuditType = "FTS_INDEX_CREATED"
	FTSINDEXAUDITTYPE_UPDATED FTSIndexAuditType = "FTS_INDEX_UPDATED"
	FTSINDEXAUDITTYPE_DELETED FTSIndexAuditType = "FTS_INDEX_DELETED"
)

// All allowed values of FTSIndexAuditType enum
var AllowedFTSIndexAuditTypeEnumValues = []FTSIndexAuditType{
	"FTS_INDEX_DELETION_FAILED",
	"FTS_INDEX_BUILD_COMPLETE",
	"FTS_INDEX_BUILD_FAILED",
	"FTS_INDEX_CREATED",
	"FTS_INDEX_UPDATED",
	"FTS_INDEX_DELETED",
}

func (v *FTSIndexAuditType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FTSIndexAuditType(value)
	for _, existing := range AllowedFTSIndexAuditTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FTSIndexAuditType", value)
}

// NewFTSIndexAuditTypeFromValue returns a pointer to a valid FTSIndexAuditType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFTSIndexAuditTypeFromValue(v string) (*FTSIndexAuditType, error) {
	ev := FTSIndexAuditType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FTSIndexAuditType: valid values are %v", v, AllowedFTSIndexAuditTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FTSIndexAuditType) IsValid() bool {
	for _, existing := range AllowedFTSIndexAuditTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FTSIndexAuditType value
func (v FTSIndexAuditType) Ptr() *FTSIndexAuditType {
	return &v
}

type NullableFTSIndexAuditType struct {
	value *FTSIndexAuditType
	isSet bool
}

func (v NullableFTSIndexAuditType) Get() *FTSIndexAuditType {
	return v.value
}

func (v *NullableFTSIndexAuditType) Set(val *FTSIndexAuditType) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSIndexAuditType) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSIndexAuditType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSIndexAuditType(val *FTSIndexAuditType) *NullableFTSIndexAuditType {
	return &NullableFTSIndexAuditType{value: val, isSet: true}
}

func (v NullableFTSIndexAuditType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSIndexAuditType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
