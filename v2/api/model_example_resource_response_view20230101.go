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

// ExampleResourceResponseView20230101 struct for ExampleResourceResponseView20230101
type ExampleResourceResponseView20230101 struct {
	// Dummy data added as response.
	Data string `json:"data"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
}

// NewExampleResourceResponseView20230101 instantiates a new ExampleResourceResponseView20230101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExampleResourceResponseView20230101() *ExampleResourceResponseView20230101 {
	this := ExampleResourceResponseView20230101{}
	return &this
}

// NewExampleResourceResponseView20230101WithDefaults instantiates a new ExampleResourceResponseView20230101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExampleResourceResponseView20230101WithDefaults() *ExampleResourceResponseView20230101 {
	this := ExampleResourceResponseView20230101{}
	return &this
}

// GetData returns the Data field value
func (o *ExampleResourceResponseView20230101) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponseView20230101) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExampleResourceResponseView20230101) SetData(v string) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ExampleResourceResponseView20230101) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponseView20230101) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExampleResourceResponseView20230101) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ExampleResourceResponseView20230101) SetLinks(v []Link) {
	o.Links = v
}

func (o ExampleResourceResponseView20230101) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableExampleResourceResponseView20230101 struct {
	value *ExampleResourceResponseView20230101
	isSet bool
}

func (v NullableExampleResourceResponseView20230101) Get() *ExampleResourceResponseView20230101 {
	return v.value
}

func (v *NullableExampleResourceResponseView20230101) Set(val *ExampleResourceResponseView20230101) {
	v.value = val
	v.isSet = true
}

func (v NullableExampleResourceResponseView20230101) IsSet() bool {
	return v.isSet
}

func (v *NullableExampleResourceResponseView20230101) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExampleResourceResponseView20230101(val *ExampleResourceResponseView20230101) *NullableExampleResourceResponseView20230101 {
	return &NullableExampleResourceResponseView20230101{value: val, isSet: true}
}

func (v NullableExampleResourceResponseView20230101) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExampleResourceResponseView20230101) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

