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

// GCPProviderSettings struct for GCPProviderSettings
type GCPProviderSettings struct {
	AutoScaling *GCPAutoScaling `json:"autoScaling,omitempty"`
	// Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster.
	InstanceSizeName *string `json:"instanceSizeName,omitempty"`
	// Google Compute Regions.
	RegionName *string `json:"regionName,omitempty"`
}

// NewGCPProviderSettings instantiates a new GCPProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPProviderSettings() *GCPProviderSettings {
	this := GCPProviderSettings{}
	return &this
}

// NewGCPProviderSettingsWithDefaults instantiates a new GCPProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPProviderSettingsWithDefaults() *GCPProviderSettings {
	this := GCPProviderSettings{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *GCPProviderSettings) GetAutoScaling() GCPAutoScaling {
	if o == nil || o.AutoScaling == nil {
		var ret GCPAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPProviderSettings) GetAutoScalingOk() (*GCPAutoScaling, bool) {
	if o == nil || o.AutoScaling == nil {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *GCPProviderSettings) HasAutoScaling() bool {
	if o != nil && o.AutoScaling != nil {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given GCPAutoScaling and assigns it to the AutoScaling field.
func (o *GCPProviderSettings) SetAutoScaling(v GCPAutoScaling) {
	o.AutoScaling = &v
}

// GetInstanceSizeName returns the InstanceSizeName field value if set, zero value otherwise.
func (o *GCPProviderSettings) GetInstanceSizeName() string {
	if o == nil || o.InstanceSizeName == nil {
		var ret string
		return ret
	}
	return *o.InstanceSizeName
}

// GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPProviderSettings) GetInstanceSizeNameOk() (*string, bool) {
	if o == nil || o.InstanceSizeName == nil {
		return nil, false
	}
	return o.InstanceSizeName, true
}

// HasInstanceSizeName returns a boolean if a field has been set.
func (o *GCPProviderSettings) HasInstanceSizeName() bool {
	if o != nil && o.InstanceSizeName != nil {
		return true
	}

	return false
}

// SetInstanceSizeName gets a reference to the given string and assigns it to the InstanceSizeName field.
func (o *GCPProviderSettings) SetInstanceSizeName(v string) {
	o.InstanceSizeName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *GCPProviderSettings) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPProviderSettings) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *GCPProviderSettings) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *GCPProviderSettings) SetRegionName(v string) {
	o.RegionName = &v
}

func (o GCPProviderSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoScaling != nil {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if o.InstanceSizeName != nil {
		toSerialize["instanceSizeName"] = o.InstanceSizeName
	}
	if o.RegionName != nil {
		toSerialize["regionName"] = o.RegionName
	}
	return json.Marshal(toSerialize)
}

type NullableGCPProviderSettings struct {
	value *GCPProviderSettings
	isSet bool
}

func (v NullableGCPProviderSettings) Get() *GCPProviderSettings {
	return v.value
}

func (v *NullableGCPProviderSettings) Set(val *GCPProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPProviderSettings(val *GCPProviderSettings) *NullableGCPProviderSettings {
	return &NullableGCPProviderSettings{value: val, isSet: true}
}

func (v NullableGCPProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

