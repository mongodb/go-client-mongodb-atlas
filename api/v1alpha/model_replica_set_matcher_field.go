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

// ReplicaSetMatcherField Name of the parameter in the target object that MongoDB Cloud checks. The parameter must match all rules for MongoDB Cloud to check for alert configurations.
type ReplicaSetMatcherField string

// List of ReplicaSetMatcherField
const (
	REPLICASETMATCHERFIELD_REPLICA_SET_NAME ReplicaSetMatcherField = "REPLICA_SET_NAME"
	REPLICASETMATCHERFIELD_SHARD_NAME ReplicaSetMatcherField = "SHARD_NAME"
	REPLICASETMATCHERFIELD_CLUSTER_NAME ReplicaSetMatcherField = "CLUSTER_NAME"
)

// All allowed values of ReplicaSetMatcherField enum
var AllowedReplicaSetMatcherFieldEnumValues = []ReplicaSetMatcherField{
	"REPLICA_SET_NAME",
	"SHARD_NAME",
	"CLUSTER_NAME",
}

func (v *ReplicaSetMatcherField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReplicaSetMatcherField(value)
	for _, existing := range AllowedReplicaSetMatcherFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReplicaSetMatcherField", value)
}

// NewReplicaSetMatcherFieldFromValue returns a pointer to a valid ReplicaSetMatcherField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReplicaSetMatcherFieldFromValue(v string) (*ReplicaSetMatcherField, error) {
	ev := ReplicaSetMatcherField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReplicaSetMatcherField: valid values are %v", v, AllowedReplicaSetMatcherFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReplicaSetMatcherField) IsValid() bool {
	for _, existing := range AllowedReplicaSetMatcherFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReplicaSetMatcherField value
func (v ReplicaSetMatcherField) Ptr() *ReplicaSetMatcherField {
	return &v
}

type NullableReplicaSetMatcherField struct {
	value *ReplicaSetMatcherField
	isSet bool
}

func (v NullableReplicaSetMatcherField) Get() *ReplicaSetMatcherField {
	return v.value
}

func (v *NullableReplicaSetMatcherField) Set(val *ReplicaSetMatcherField) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicaSetMatcherField) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicaSetMatcherField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicaSetMatcherField(val *ReplicaSetMatcherField) *NullableReplicaSetMatcherField {
	return &NullableReplicaSetMatcherField{value: val, isSet: true}
}

func (v NullableReplicaSetMatcherField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicaSetMatcherField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
