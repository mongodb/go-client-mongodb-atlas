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

// checks if the AzureComputeAutoScaling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureComputeAutoScaling{}

// AzureComputeAutoScaling Collection of settings that configures how a cluster might scale its cluster tier and whether the cluster can scale down. Cluster tier auto-scaling is unavailable for clusters using Low CPU or NVME storage classes.
type AzureComputeAutoScaling struct {
	// Maximum instance size to which your cluster can automatically scale.
	MaxInstanceSize *string `json:"maxInstanceSize,omitempty"`
	// Minimum instance size to which your cluster can automatically scale.
	MinInstanceSize *string `json:"minInstanceSize,omitempty"`
}

// NewAzureComputeAutoScaling instantiates a new AzureComputeAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureComputeAutoScaling() *AzureComputeAutoScaling {
	this := AzureComputeAutoScaling{}
	return &this
}

// NewAzureComputeAutoScalingWithDefaults instantiates a new AzureComputeAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureComputeAutoScalingWithDefaults() *AzureComputeAutoScaling {
	this := AzureComputeAutoScaling{}
	return &this
}

// GetMaxInstanceSize returns the MaxInstanceSize field value if set, zero value otherwise.
func (o *AzureComputeAutoScaling) GetMaxInstanceSize() string {
	if o == nil || IsNil(o.MaxInstanceSize) {
		var ret string
		return ret
	}
	return *o.MaxInstanceSize
}

// GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeAutoScaling) GetMaxInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxInstanceSize) {
		return nil, false
	}
	return o.MaxInstanceSize, true
}

// HasMaxInstanceSize returns a boolean if a field has been set.
func (o *AzureComputeAutoScaling) HasMaxInstanceSize() bool {
	if o != nil && !IsNil(o.MaxInstanceSize) {
		return true
	}

	return false
}

// SetMaxInstanceSize gets a reference to the given string and assigns it to the MaxInstanceSize field.
func (o *AzureComputeAutoScaling) SetMaxInstanceSize(v string) {
	o.MaxInstanceSize = &v
}

// GetMinInstanceSize returns the MinInstanceSize field value if set, zero value otherwise.
func (o *AzureComputeAutoScaling) GetMinInstanceSize() string {
	if o == nil || IsNil(o.MinInstanceSize) {
		var ret string
		return ret
	}
	return *o.MinInstanceSize
}

// GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeAutoScaling) GetMinInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MinInstanceSize) {
		return nil, false
	}
	return o.MinInstanceSize, true
}

// HasMinInstanceSize returns a boolean if a field has been set.
func (o *AzureComputeAutoScaling) HasMinInstanceSize() bool {
	if o != nil && !IsNil(o.MinInstanceSize) {
		return true
	}

	return false
}

// SetMinInstanceSize gets a reference to the given string and assigns it to the MinInstanceSize field.
func (o *AzureComputeAutoScaling) SetMinInstanceSize(v string) {
	o.MinInstanceSize = &v
}

func (o AzureComputeAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AzureComputeAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxInstanceSize) {
		toSerialize["maxInstanceSize"] = o.MaxInstanceSize
	}
	if !IsNil(o.MinInstanceSize) {
		toSerialize["minInstanceSize"] = o.MinInstanceSize
	}
	return toSerialize, nil
}

type NullableAzureComputeAutoScaling struct {
	value *AzureComputeAutoScaling
	isSet bool
}

func (v NullableAzureComputeAutoScaling) Get() *AzureComputeAutoScaling {
	return v.value
}

func (v *NullableAzureComputeAutoScaling) Set(val *AzureComputeAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureComputeAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureComputeAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureComputeAutoScaling(val *AzureComputeAutoScaling) *NullableAzureComputeAutoScaling {
	return &NullableAzureComputeAutoScaling{value: val, isSet: true}
}

func (v NullableAzureComputeAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureComputeAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

