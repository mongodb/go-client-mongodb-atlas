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

// AppServiceAlertableEventType Incident that triggered this alert.
type AppServiceAlertableEventType string

// List of AppServiceAlertableEventType
const (
	APPSERVICEALERTABLEEVENTTYPE_URL_CONFIRMATION AppServiceAlertableEventType = "URL_CONFIRMATION"
	APPSERVICEALERTABLEEVENTTYPE_SUCCESSFUL_DEPLOY AppServiceAlertableEventType = "SUCCESSFUL_DEPLOY"
	APPSERVICEALERTABLEEVENTTYPE_DEPLOYMENT_FAILURE AppServiceAlertableEventType = "DEPLOYMENT_FAILURE"
	APPSERVICEALERTABLEEVENTTYPE_REQUEST_RATE_LIMIT AppServiceAlertableEventType = "REQUEST_RATE_LIMIT"
	APPSERVICEALERTABLEEVENTTYPE_LOG_FORWARDER_FAILURE AppServiceAlertableEventType = "LOG_FORWARDER_FAILURE"
	APPSERVICEALERTABLEEVENTTYPE_OUTSIDE_REALM_METRIC_THRESHOLD AppServiceAlertableEventType = "OUTSIDE_REALM_METRIC_THRESHOLD"
	APPSERVICEALERTABLEEVENTTYPE_SYNC_FAILURE AppServiceAlertableEventType = "SYNC_FAILURE"
	APPSERVICEALERTABLEEVENTTYPE_TRIGGER_FAILURE AppServiceAlertableEventType = "TRIGGER_FAILURE"
	APPSERVICEALERTABLEEVENTTYPE_TRIGGER_AUTO_RESUMED AppServiceAlertableEventType = "TRIGGER_AUTO_RESUMED"
)

// All allowed values of AppServiceAlertableEventType enum
var AllowedAppServiceAlertableEventTypeEnumValues = []AppServiceAlertableEventType{
	"URL_CONFIRMATION",
	"SUCCESSFUL_DEPLOY",
	"DEPLOYMENT_FAILURE",
	"REQUEST_RATE_LIMIT",
	"LOG_FORWARDER_FAILURE",
	"OUTSIDE_REALM_METRIC_THRESHOLD",
	"SYNC_FAILURE",
	"TRIGGER_FAILURE",
	"TRIGGER_AUTO_RESUMED",
}

func (v *AppServiceAlertableEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppServiceAlertableEventType(value)
	for _, existing := range AllowedAppServiceAlertableEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppServiceAlertableEventType", value)
}

// NewAppServiceAlertableEventTypeFromValue returns a pointer to a valid AppServiceAlertableEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppServiceAlertableEventTypeFromValue(v string) (*AppServiceAlertableEventType, error) {
	ev := AppServiceAlertableEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppServiceAlertableEventType: valid values are %v", v, AllowedAppServiceAlertableEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppServiceAlertableEventType) IsValid() bool {
	for _, existing := range AllowedAppServiceAlertableEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppServiceAlertableEventType value
func (v AppServiceAlertableEventType) Ptr() *AppServiceAlertableEventType {
	return &v
}

type NullableAppServiceAlertableEventType struct {
	value *AppServiceAlertableEventType
	isSet bool
}

func (v NullableAppServiceAlertableEventType) Get() *AppServiceAlertableEventType {
	return v.value
}

func (v *NullableAppServiceAlertableEventType) Set(val *AppServiceAlertableEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableAppServiceAlertableEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableAppServiceAlertableEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppServiceAlertableEventType(val *AppServiceAlertableEventType) *NullableAppServiceAlertableEventType {
	return &NullableAppServiceAlertableEventType{value: val, isSet: true}
}

func (v NullableAppServiceAlertableEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppServiceAlertableEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
