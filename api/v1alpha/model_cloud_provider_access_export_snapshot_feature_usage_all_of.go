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

// checks if the CloudProviderAccessExportSnapshotFeatureUsageAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAccessExportSnapshotFeatureUsageAllOf{}

// CloudProviderAccessExportSnapshotFeatureUsageAllOf struct for CloudProviderAccessExportSnapshotFeatureUsageAllOf
type CloudProviderAccessExportSnapshotFeatureUsageAllOf struct {
	FeatureId *CloudProviderAccessFeatureUsageExportSnapshotFeatureId `json:"featureId,omitempty"`
}

// NewCloudProviderAccessExportSnapshotFeatureUsageAllOf instantiates a new CloudProviderAccessExportSnapshotFeatureUsageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessExportSnapshotFeatureUsageAllOf() *CloudProviderAccessExportSnapshotFeatureUsageAllOf {
	this := CloudProviderAccessExportSnapshotFeatureUsageAllOf{}
	return &this
}

// NewCloudProviderAccessExportSnapshotFeatureUsageAllOfWithDefaults instantiates a new CloudProviderAccessExportSnapshotFeatureUsageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessExportSnapshotFeatureUsageAllOfWithDefaults() *CloudProviderAccessExportSnapshotFeatureUsageAllOf {
	this := CloudProviderAccessExportSnapshotFeatureUsageAllOf{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *CloudProviderAccessExportSnapshotFeatureUsageAllOf) GetFeatureId() CloudProviderAccessFeatureUsageExportSnapshotFeatureId {
	if o == nil || IsNil(o.FeatureId) {
		var ret CloudProviderAccessFeatureUsageExportSnapshotFeatureId
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessExportSnapshotFeatureUsageAllOf) GetFeatureIdOk() (*CloudProviderAccessFeatureUsageExportSnapshotFeatureId, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *CloudProviderAccessExportSnapshotFeatureUsageAllOf) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given CloudProviderAccessFeatureUsageExportSnapshotFeatureId and assigns it to the FeatureId field.
func (o *CloudProviderAccessExportSnapshotFeatureUsageAllOf) SetFeatureId(v CloudProviderAccessFeatureUsageExportSnapshotFeatureId) {
	o.FeatureId = &v
}

func (o CloudProviderAccessExportSnapshotFeatureUsageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudProviderAccessExportSnapshotFeatureUsageAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureId) {
		toSerialize["featureId"] = o.FeatureId
	}
	return toSerialize, nil
}

type NullableCloudProviderAccessExportSnapshotFeatureUsageAllOf struct {
	value *CloudProviderAccessExportSnapshotFeatureUsageAllOf
	isSet bool
}

func (v NullableCloudProviderAccessExportSnapshotFeatureUsageAllOf) Get() *CloudProviderAccessExportSnapshotFeatureUsageAllOf {
	return v.value
}

func (v *NullableCloudProviderAccessExportSnapshotFeatureUsageAllOf) Set(val *CloudProviderAccessExportSnapshotFeatureUsageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessExportSnapshotFeatureUsageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessExportSnapshotFeatureUsageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessExportSnapshotFeatureUsageAllOf(val *CloudProviderAccessExportSnapshotFeatureUsageAllOf) *NullableCloudProviderAccessExportSnapshotFeatureUsageAllOf {
	return &NullableCloudProviderAccessExportSnapshotFeatureUsageAllOf{value: val, isSet: true}
}

func (v NullableCloudProviderAccessExportSnapshotFeatureUsageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessExportSnapshotFeatureUsageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

