/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"fmt"
)

// NDSAuditTypeForOrg Unique identifier of event type.
type NDSAuditTypeForOrg string

// List of NDSAuditTypeForOrg
const (
	NDSAUDITTYPEFORORG_ORG_LIMIT_UPDATED NDSAuditTypeForOrg = "ORG_LIMIT_UPDATED"
)

// All allowed values of NDSAuditTypeForOrg enum
var AllowedNDSAuditTypeForOrgEnumValues = []NDSAuditTypeForOrg{
	"ORG_LIMIT_UPDATED",
}

func (v *NDSAuditTypeForOrg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NDSAuditTypeForOrg(value)
	for _, existing := range AllowedNDSAuditTypeForOrgEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NDSAuditTypeForOrg", value)
}

// NewNDSAuditTypeForOrgFromValue returns a pointer to a valid NDSAuditTypeForOrg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNDSAuditTypeForOrgFromValue(v string) (*NDSAuditTypeForOrg, error) {
	ev := NDSAuditTypeForOrg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NDSAuditTypeForOrg: valid values are %v", v, AllowedNDSAuditTypeForOrgEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NDSAuditTypeForOrg) IsValid() bool {
	for _, existing := range AllowedNDSAuditTypeForOrgEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NDSAuditTypeForOrg value
func (v NDSAuditTypeForOrg) Ptr() *NDSAuditTypeForOrg {
	return &v
}

type NullableNDSAuditTypeForOrg struct {
	value *NDSAuditTypeForOrg
	isSet bool
}

func (v NullableNDSAuditTypeForOrg) Get() *NDSAuditTypeForOrg {
	return v.value
}

func (v *NullableNDSAuditTypeForOrg) Set(val *NDSAuditTypeForOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSAuditTypeForOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSAuditTypeForOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSAuditTypeForOrg(val *NDSAuditTypeForOrg) *NullableNDSAuditTypeForOrg {
	return &NullableNDSAuditTypeForOrg{value: val, isSet: true}
}

func (v NullableNDSAuditTypeForOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSAuditTypeForOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
