/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the AWSCustomDNSEnabled type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSCustomDNSEnabled{}

// AWSCustomDNSEnabled struct for AWSCustomDNSEnabled
type AWSCustomDNSEnabled struct {
	// Flag that indicates whether the project's clusters deployed to Amazon Web Services (AWS) use a custom Domain Name System (DNS). When `\"enabled\": true`, connect to your cluster using Private IP for Peering connection strings.
	Enabled bool `json:"enabled"`
}

// NewAWSCustomDNSEnabled instantiates a new AWSCustomDNSEnabled object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSCustomDNSEnabled(enabled bool) *AWSCustomDNSEnabled {
	this := AWSCustomDNSEnabled{}
	this.Enabled = enabled
	return &this
}

// NewAWSCustomDNSEnabledWithDefaults instantiates a new AWSCustomDNSEnabled object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSCustomDNSEnabledWithDefaults() *AWSCustomDNSEnabled {
	this := AWSCustomDNSEnabled{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *AWSCustomDNSEnabled) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AWSCustomDNSEnabled) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AWSCustomDNSEnabled) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AWSCustomDNSEnabled) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSCustomDNSEnabled) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAWSCustomDNSEnabled struct {
	value *AWSCustomDNSEnabled
	isSet bool
}

func (v NullableAWSCustomDNSEnabled) Get() *AWSCustomDNSEnabled {
	return v.value
}

func (v *NullableAWSCustomDNSEnabled) Set(val *AWSCustomDNSEnabled) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSCustomDNSEnabled) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSCustomDNSEnabled) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSCustomDNSEnabled(val *AWSCustomDNSEnabled) *NullableAWSCustomDNSEnabled {
	return &NullableAWSCustomDNSEnabled{value: val, isSet: true}
}

func (v NullableAWSCustomDNSEnabled) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSCustomDNSEnabled) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

