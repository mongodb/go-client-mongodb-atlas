/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// ApiPagerDutyViewAllOf struct for ApiPagerDutyViewAllOf
type ApiPagerDutyViewAllOf struct {
	// Service key associated with your PagerDuty account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ServiceKey *string `json:"serviceKey,omitempty"`
}

// NewApiPagerDutyViewAllOf instantiates a new ApiPagerDutyViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPagerDutyViewAllOf() *ApiPagerDutyViewAllOf {
	this := ApiPagerDutyViewAllOf{}
	return &this
}

// NewApiPagerDutyViewAllOfWithDefaults instantiates a new ApiPagerDutyViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPagerDutyViewAllOfWithDefaults() *ApiPagerDutyViewAllOf {
	this := ApiPagerDutyViewAllOf{}
	return &this
}

// GetServiceKey returns the ServiceKey field value if set, zero value otherwise.
func (o *ApiPagerDutyViewAllOf) GetServiceKey() string {
	if o == nil || o.ServiceKey == nil {
		var ret string
		return ret
	}
	return *o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPagerDutyViewAllOf) GetServiceKeyOk() (*string, bool) {
	if o == nil || o.ServiceKey == nil {
		return nil, false
	}
	return o.ServiceKey, true
}

// HasServiceKey returns a boolean if a field has been set.
func (o *ApiPagerDutyViewAllOf) HasServiceKey() bool {
	if o != nil && o.ServiceKey != nil {
		return true
	}

	return false
}

// SetServiceKey gets a reference to the given string and assigns it to the ServiceKey field.
func (o *ApiPagerDutyViewAllOf) SetServiceKey(v string) {
	o.ServiceKey = &v
}

func (o ApiPagerDutyViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceKey != nil {
		toSerialize["serviceKey"] = o.ServiceKey
	}
	return json.Marshal(toSerialize)
}

type NullableApiPagerDutyViewAllOf struct {
	value *ApiPagerDutyViewAllOf
	isSet bool
}

func (v NullableApiPagerDutyViewAllOf) Get() *ApiPagerDutyViewAllOf {
	return v.value
}

func (v *NullableApiPagerDutyViewAllOf) Set(val *ApiPagerDutyViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPagerDutyViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPagerDutyViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPagerDutyViewAllOf(val *ApiPagerDutyViewAllOf) *NullableApiPagerDutyViewAllOf {
	return &NullableApiPagerDutyViewAllOf{value: val, isSet: true}
}

func (v NullableApiPagerDutyViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPagerDutyViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

