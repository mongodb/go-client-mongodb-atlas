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

// CustomerX509 Settings to configure TLS Certificates for database users.
type CustomerX509 struct {
	// Concatenated list of customer certificate authority (CA) certificates needed to authenticate database users. MongoDB Cloud expects this as a PEM-formatted certificate.
	Cas *string `json:"cas,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
}

// NewCustomerX509 instantiates a new CustomerX509 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerX509() *CustomerX509 {
	this := CustomerX509{}
	return &this
}

// NewCustomerX509WithDefaults instantiates a new CustomerX509 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerX509WithDefaults() *CustomerX509 {
	this := CustomerX509{}
	return &this
}

// GetCas returns the Cas field value if set, zero value otherwise.
func (o *CustomerX509) GetCas() string {
	if o == nil || o.Cas == nil {
		var ret string
		return ret
	}
	return *o.Cas
}

// GetCasOk returns a tuple with the Cas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerX509) GetCasOk() (*string, bool) {
	if o == nil || o.Cas == nil {
		return nil, false
	}
	return o.Cas, true
}

// HasCas returns a boolean if a field has been set.
func (o *CustomerX509) HasCas() bool {
	if o != nil && o.Cas != nil {
		return true
	}

	return false
}

// SetCas gets a reference to the given string and assigns it to the Cas field.
func (o *CustomerX509) SetCas(v string) {
	o.Cas = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CustomerX509) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerX509) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CustomerX509) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *CustomerX509) SetLinks(v []Link) {
	o.Links = v
}

func (o CustomerX509) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cas != nil {
		toSerialize["cas"] = o.Cas
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerX509 struct {
	value *CustomerX509
	isSet bool
}

func (v NullableCustomerX509) Get() *CustomerX509 {
	return v.value
}

func (v *NullableCustomerX509) Set(val *CustomerX509) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerX509) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerX509) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerX509(val *CustomerX509) *NullableCustomerX509 {
	return &NullableCustomerX509{value: val, isSet: true}
}

func (v NullableCustomerX509) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerX509) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

