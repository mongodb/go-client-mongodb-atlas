/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the DLSIngestionSinkAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DLSIngestionSinkAllOf{}

// DLSIngestionSinkAllOf struct for DLSIngestionSinkAllOf
type DLSIngestionSinkAllOf struct {
	// Target cloud provider for this Data Lake Pipeline.
	MetadataProvider *string `json:"metadataProvider,omitempty"`
	// Target cloud provider region for this Data Lake Pipeline.
	MetadataRegion *string `json:"metadataRegion,omitempty"`
	// Ordered fields used to physically organize data in the destination.
	PartitionFields []PartitionField `json:"partitionFields,omitempty"`
}

// NewDLSIngestionSinkAllOf instantiates a new DLSIngestionSinkAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDLSIngestionSinkAllOf() *DLSIngestionSinkAllOf {
	this := DLSIngestionSinkAllOf{}
	return &this
}

// NewDLSIngestionSinkAllOfWithDefaults instantiates a new DLSIngestionSinkAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDLSIngestionSinkAllOfWithDefaults() *DLSIngestionSinkAllOf {
	this := DLSIngestionSinkAllOf{}
	return &this
}

// GetMetadataProvider returns the MetadataProvider field value if set, zero value otherwise.
func (o *DLSIngestionSinkAllOf) GetMetadataProvider() string {
	if o == nil || IsNil(o.MetadataProvider) {
		var ret string
		return ret
	}
	return *o.MetadataProvider
}

// GetMetadataProviderOk returns a tuple with the MetadataProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLSIngestionSinkAllOf) GetMetadataProviderOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataProvider) {
		return nil, false
	}
	return o.MetadataProvider, true
}

// HasMetadataProvider returns a boolean if a field has been set.
func (o *DLSIngestionSinkAllOf) HasMetadataProvider() bool {
	if o != nil && !IsNil(o.MetadataProvider) {
		return true
	}

	return false
}

// SetMetadataProvider gets a reference to the given string and assigns it to the MetadataProvider field.
func (o *DLSIngestionSinkAllOf) SetMetadataProvider(v string) {
	o.MetadataProvider = &v
}

// GetMetadataRegion returns the MetadataRegion field value if set, zero value otherwise.
func (o *DLSIngestionSinkAllOf) GetMetadataRegion() string {
	if o == nil || IsNil(o.MetadataRegion) {
		var ret string
		return ret
	}
	return *o.MetadataRegion
}

// GetMetadataRegionOk returns a tuple with the MetadataRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLSIngestionSinkAllOf) GetMetadataRegionOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataRegion) {
		return nil, false
	}
	return o.MetadataRegion, true
}

// HasMetadataRegion returns a boolean if a field has been set.
func (o *DLSIngestionSinkAllOf) HasMetadataRegion() bool {
	if o != nil && !IsNil(o.MetadataRegion) {
		return true
	}

	return false
}

// SetMetadataRegion gets a reference to the given string and assigns it to the MetadataRegion field.
func (o *DLSIngestionSinkAllOf) SetMetadataRegion(v string) {
	o.MetadataRegion = &v
}

// GetPartitionFields returns the PartitionFields field value if set, zero value otherwise.
func (o *DLSIngestionSinkAllOf) GetPartitionFields() []PartitionField {
	if o == nil || IsNil(o.PartitionFields) {
		var ret []PartitionField
		return ret
	}
	return o.PartitionFields
}

// GetPartitionFieldsOk returns a tuple with the PartitionFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLSIngestionSinkAllOf) GetPartitionFieldsOk() ([]PartitionField, bool) {
	if o == nil || IsNil(o.PartitionFields) {
		return nil, false
	}
	return o.PartitionFields, true
}

// HasPartitionFields returns a boolean if a field has been set.
func (o *DLSIngestionSinkAllOf) HasPartitionFields() bool {
	if o != nil && !IsNil(o.PartitionFields) {
		return true
	}

	return false
}

// SetPartitionFields gets a reference to the given []PartitionField and assigns it to the PartitionFields field.
func (o *DLSIngestionSinkAllOf) SetPartitionFields(v []PartitionField) {
	o.PartitionFields = v
}

func (o DLSIngestionSinkAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DLSIngestionSinkAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetadataProvider) {
		toSerialize["metadataProvider"] = o.MetadataProvider
	}
	if !IsNil(o.MetadataRegion) {
		toSerialize["metadataRegion"] = o.MetadataRegion
	}
	if !IsNil(o.PartitionFields) {
		toSerialize["partitionFields"] = o.PartitionFields
	}
	return toSerialize, nil
}

type NullableDLSIngestionSinkAllOf struct {
	value *DLSIngestionSinkAllOf
	isSet bool
}

func (v NullableDLSIngestionSinkAllOf) Get() *DLSIngestionSinkAllOf {
	return v.value
}

func (v *NullableDLSIngestionSinkAllOf) Set(val *DLSIngestionSinkAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDLSIngestionSinkAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDLSIngestionSinkAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDLSIngestionSinkAllOf(val *DLSIngestionSinkAllOf) *NullableDLSIngestionSinkAllOf {
	return &NullableDLSIngestionSinkAllOf{value: val, isSet: true}
}

func (v NullableDLSIngestionSinkAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDLSIngestionSinkAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

