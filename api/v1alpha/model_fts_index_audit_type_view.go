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

// FTSIndexAuditTypeView Unique identifier of event type.
type FTSIndexAuditTypeView string

// List of FTSIndexAuditTypeView
const (
	FTSINDEXAUDITTYPEVIEW_DELETION_FAILED FTSIndexAuditTypeView = "FTS_INDEX_DELETION_FAILED"
	FTSINDEXAUDITTYPEVIEW_BUILD_COMPLETE FTSIndexAuditTypeView = "FTS_INDEX_BUILD_COMPLETE"
	FTSINDEXAUDITTYPEVIEW_BUILD_FAILED FTSIndexAuditTypeView = "FTS_INDEX_BUILD_FAILED"
	FTSINDEXAUDITTYPEVIEW_CREATED FTSIndexAuditTypeView = "FTS_INDEX_CREATED"
	FTSINDEXAUDITTYPEVIEW_UPDATED FTSIndexAuditTypeView = "FTS_INDEX_UPDATED"
	FTSINDEXAUDITTYPEVIEW_DELETED FTSIndexAuditTypeView = "FTS_INDEX_DELETED"
)

// All allowed values of FTSIndexAuditTypeView enum
var AllowedFTSIndexAuditTypeViewEnumValues = []FTSIndexAuditTypeView{
	"FTS_INDEX_DELETION_FAILED",
	"FTS_INDEX_BUILD_COMPLETE",
	"FTS_INDEX_BUILD_FAILED",
	"FTS_INDEX_CREATED",
	"FTS_INDEX_UPDATED",
	"FTS_INDEX_DELETED",
}

func (v *FTSIndexAuditTypeView) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FTSIndexAuditTypeView(value)
	for _, existing := range AllowedFTSIndexAuditTypeViewEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FTSIndexAuditTypeView", value)
}

// NewFTSIndexAuditTypeViewFromValue returns a pointer to a valid FTSIndexAuditTypeView
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFTSIndexAuditTypeViewFromValue(v string) (*FTSIndexAuditTypeView, error) {
	ev := FTSIndexAuditTypeView(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FTSIndexAuditTypeView: valid values are %v", v, AllowedFTSIndexAuditTypeViewEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FTSIndexAuditTypeView) IsValid() bool {
	for _, existing := range AllowedFTSIndexAuditTypeViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FTSIndexAuditTypeView value
func (v FTSIndexAuditTypeView) Ptr() *FTSIndexAuditTypeView {
	return &v
}

type NullableFTSIndexAuditTypeView struct {
	value *FTSIndexAuditTypeView
	isSet bool
}

func (v NullableFTSIndexAuditTypeView) Get() *FTSIndexAuditTypeView {
	return v.value
}

func (v *NullableFTSIndexAuditTypeView) Set(val *FTSIndexAuditTypeView) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSIndexAuditTypeView) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSIndexAuditTypeView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSIndexAuditTypeView(val *FTSIndexAuditTypeView) *NullableFTSIndexAuditTypeView {
	return &NullableFTSIndexAuditTypeView{value: val, isSet: true}
}

func (v NullableFTSIndexAuditTypeView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSIndexAuditTypeView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
