/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ClusterAlertableEventType Incident that triggered this alert.
type ClusterAlertableEventType string

// List of ClusterAlertableEventType
const (
	CLUSTERALERTABLEEVENTTYPE_CLUSTER_MONGOS_IS_MISSING ClusterAlertableEventType = "CLUSTER_MONGOS_IS_MISSING"
)

// All allowed values of ClusterAlertableEventType enum
var AllowedClusterAlertableEventTypeEnumValues = []ClusterAlertableEventType{
	"CLUSTER_MONGOS_IS_MISSING",
}

func (v *ClusterAlertableEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterAlertableEventType(value)
	for _, existing := range AllowedClusterAlertableEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterAlertableEventType", value)
}

// NewClusterAlertableEventTypeFromValue returns a pointer to a valid ClusterAlertableEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterAlertableEventTypeFromValue(v string) (*ClusterAlertableEventType, error) {
	ev := ClusterAlertableEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterAlertableEventType: valid values are %v", v, AllowedClusterAlertableEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterAlertableEventType) IsValid() bool {
	for _, existing := range AllowedClusterAlertableEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterAlertableEventType value
func (v ClusterAlertableEventType) Ptr() *ClusterAlertableEventType {
	return &v
}

type NullableClusterAlertableEventType struct {
	value *ClusterAlertableEventType
	isSet bool
}

func (v NullableClusterAlertableEventType) Get() *ClusterAlertableEventType {
	return v.value
}

func (v *NullableClusterAlertableEventType) Set(val *ClusterAlertableEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAlertableEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAlertableEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAlertableEventType(val *ClusterAlertableEventType) *NullableClusterAlertableEventType {
	return &NullableClusterAlertableEventType{value: val, isSet: true}
}

func (v NullableClusterAlertableEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAlertableEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
