/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// NDSLDAPVerifyConnectivityJobRequestValidation One test that MongoDB Cloud runs to test verification of the provided Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration details.
type NDSLDAPVerifyConnectivityJobRequestValidation struct {
	// Human-readable string that indicates the result of this verification test.
	Status *string `json:"status,omitempty"`
	// Human-readable label that identifies this verification test that MongoDB Cloud runs.
	ValidationType *string `json:"validationType,omitempty"`
}

// NewNDSLDAPVerifyConnectivityJobRequestValidation instantiates a new NDSLDAPVerifyConnectivityJobRequestValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNDSLDAPVerifyConnectivityJobRequestValidation() *NDSLDAPVerifyConnectivityJobRequestValidation {
	this := NDSLDAPVerifyConnectivityJobRequestValidation{}
	return &this
}

// NewNDSLDAPVerifyConnectivityJobRequestValidationWithDefaults instantiates a new NDSLDAPVerifyConnectivityJobRequestValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNDSLDAPVerifyConnectivityJobRequestValidationWithDefaults() *NDSLDAPVerifyConnectivityJobRequestValidation {
	this := NDSLDAPVerifyConnectivityJobRequestValidation{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequestValidation) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestValidation) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestValidation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NDSLDAPVerifyConnectivityJobRequestValidation) SetStatus(v string) {
	o.Status = &v
}

// GetValidationType returns the ValidationType field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequestValidation) GetValidationType() string {
	if o == nil || o.ValidationType == nil {
		var ret string
		return ret
	}
	return *o.ValidationType
}

// GetValidationTypeOk returns a tuple with the ValidationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestValidation) GetValidationTypeOk() (*string, bool) {
	if o == nil || o.ValidationType == nil {
		return nil, false
	}
	return o.ValidationType, true
}

// HasValidationType returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestValidation) HasValidationType() bool {
	if o != nil && o.ValidationType != nil {
		return true
	}

	return false
}

// SetValidationType gets a reference to the given string and assigns it to the ValidationType field.
func (o *NDSLDAPVerifyConnectivityJobRequestValidation) SetValidationType(v string) {
	o.ValidationType = &v
}

func (o NDSLDAPVerifyConnectivityJobRequestValidation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ValidationType != nil {
		toSerialize["validationType"] = o.ValidationType
	}
	return json.Marshal(toSerialize)
}

type NullableNDSLDAPVerifyConnectivityJobRequestValidation struct {
	value *NDSLDAPVerifyConnectivityJobRequestValidation
	isSet bool
}

func (v NullableNDSLDAPVerifyConnectivityJobRequestValidation) Get() *NDSLDAPVerifyConnectivityJobRequestValidation {
	return v.value
}

func (v *NullableNDSLDAPVerifyConnectivityJobRequestValidation) Set(val *NDSLDAPVerifyConnectivityJobRequestValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSLDAPVerifyConnectivityJobRequestValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSLDAPVerifyConnectivityJobRequestValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSLDAPVerifyConnectivityJobRequestValidation(val *NDSLDAPVerifyConnectivityJobRequestValidation) *NullableNDSLDAPVerifyConnectivityJobRequestValidation {
	return &NullableNDSLDAPVerifyConnectivityJobRequestValidation{value: val, isSet: true}
}

func (v NullableNDSLDAPVerifyConnectivityJobRequestValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSLDAPVerifyConnectivityJobRequestValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

