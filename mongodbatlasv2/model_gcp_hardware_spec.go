/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the GCPHardwareSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GCPHardwareSpec{}

// GCPHardwareSpec struct for GCPHardwareSpec
type GCPHardwareSpec struct {
	// Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size.
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Number of nodes of the given type for MongoDB Cloud to deploy to the region.
	NodeCount *int32 `json:"nodeCount,omitempty"`
}

// NewGCPHardwareSpec instantiates a new GCPHardwareSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPHardwareSpec() *GCPHardwareSpec {
	this := GCPHardwareSpec{}
	return &this
}

// NewGCPHardwareSpecWithDefaults instantiates a new GCPHardwareSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPHardwareSpecWithDefaults() *GCPHardwareSpec {
	this := GCPHardwareSpec{}
	return &this
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise.
func (o *GCPHardwareSpec) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPHardwareSpec) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}
	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *GCPHardwareSpec) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *GCPHardwareSpec) SetInstanceSize(v string) {
	o.InstanceSize = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *GCPHardwareSpec) GetNodeCount() int32 {
	if o == nil || IsNil(o.NodeCount) {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPHardwareSpec) GetNodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *GCPHardwareSpec) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *GCPHardwareSpec) SetNodeCount(v int32) {
	o.NodeCount = &v
}

func (o GCPHardwareSpec) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GCPHardwareSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceSize) {
		toSerialize["instanceSize"] = o.InstanceSize
	}
	if !IsNil(o.NodeCount) {
		toSerialize["nodeCount"] = o.NodeCount
	}
	return toSerialize, nil
}

type NullableGCPHardwareSpec struct {
	value *GCPHardwareSpec
	isSet bool
}

func (v NullableGCPHardwareSpec) Get() *GCPHardwareSpec {
	return v.value
}

func (v *NullableGCPHardwareSpec) Set(val *GCPHardwareSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPHardwareSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPHardwareSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPHardwareSpec(val *GCPHardwareSpec) *NullableGCPHardwareSpec {
	return &NullableGCPHardwareSpec{value: val, isSet: true}
}

func (v NullableGCPHardwareSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPHardwareSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

