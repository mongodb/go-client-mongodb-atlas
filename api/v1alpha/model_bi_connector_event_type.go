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

// BiConnectorEventType Unique identifier of event type.
type BiConnectorEventType string

// List of BiConnectorEventType
const (
	BICONNECTOREVENTTYPE_UP BiConnectorEventType = "BI_CONNECTOR_UP"
	BICONNECTOREVENTTYPE_DOWN BiConnectorEventType = "BI_CONNECTOR_DOWN"
)

// All allowed values of BiConnectorEventType enum
var AllowedBiConnectorEventTypeEnumValues = []BiConnectorEventType{
	"BI_CONNECTOR_UP",
	"BI_CONNECTOR_DOWN",
}

func (v *BiConnectorEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BiConnectorEventType(value)
	for _, existing := range AllowedBiConnectorEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BiConnectorEventType", value)
}

// NewBiConnectorEventTypeFromValue returns a pointer to a valid BiConnectorEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBiConnectorEventTypeFromValue(v string) (*BiConnectorEventType, error) {
	ev := BiConnectorEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BiConnectorEventType: valid values are %v", v, AllowedBiConnectorEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BiConnectorEventType) IsValid() bool {
	for _, existing := range AllowedBiConnectorEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BiConnectorEventType value
func (v BiConnectorEventType) Ptr() *BiConnectorEventType {
	return &v
}

type NullableBiConnectorEventType struct {
	value *BiConnectorEventType
	isSet bool
}

func (v NullableBiConnectorEventType) Get() *BiConnectorEventType {
	return v.value
}

func (v *NullableBiConnectorEventType) Set(val *BiConnectorEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableBiConnectorEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableBiConnectorEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiConnectorEventType(val *BiConnectorEventType) *NullableBiConnectorEventType {
	return &NullableBiConnectorEventType{value: val, isSet: true}
}

func (v NullableBiConnectorEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiConnectorEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
