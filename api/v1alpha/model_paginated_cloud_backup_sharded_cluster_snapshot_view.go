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

// PaginatedCloudBackupShardedClusterSnapshotView struct for PaginatedCloudBackupShardedClusterSnapshotView
type PaginatedCloudBackupShardedClusterSnapshotView struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []DiskBackupShardedClusterSnapshot `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedCloudBackupShardedClusterSnapshotView instantiates a new PaginatedCloudBackupShardedClusterSnapshotView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCloudBackupShardedClusterSnapshotView() *PaginatedCloudBackupShardedClusterSnapshotView {
	this := PaginatedCloudBackupShardedClusterSnapshotView{}
	return &this
}

// NewPaginatedCloudBackupShardedClusterSnapshotViewWithDefaults instantiates a new PaginatedCloudBackupShardedClusterSnapshotView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCloudBackupShardedClusterSnapshotViewWithDefaults() *PaginatedCloudBackupShardedClusterSnapshotView {
	this := PaginatedCloudBackupShardedClusterSnapshotView{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetResults() []DiskBackupShardedClusterSnapshot {
	if o == nil || o.Results == nil {
		var ret []DiskBackupShardedClusterSnapshot
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetResultsOk() ([]DiskBackupShardedClusterSnapshot, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DiskBackupShardedClusterSnapshot and assigns it to the Results field.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) SetResults(v []DiskBackupShardedClusterSnapshot) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedCloudBackupShardedClusterSnapshotView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedCloudBackupShardedClusterSnapshotView struct {
	value *PaginatedCloudBackupShardedClusterSnapshotView
	isSet bool
}

func (v NullablePaginatedCloudBackupShardedClusterSnapshotView) Get() *PaginatedCloudBackupShardedClusterSnapshotView {
	return v.value
}

func (v *NullablePaginatedCloudBackupShardedClusterSnapshotView) Set(val *PaginatedCloudBackupShardedClusterSnapshotView) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCloudBackupShardedClusterSnapshotView) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCloudBackupShardedClusterSnapshotView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCloudBackupShardedClusterSnapshotView(val *PaginatedCloudBackupShardedClusterSnapshotView) *NullablePaginatedCloudBackupShardedClusterSnapshotView {
	return &NullablePaginatedCloudBackupShardedClusterSnapshotView{value: val, isSet: true}
}

func (v NullablePaginatedCloudBackupShardedClusterSnapshotView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCloudBackupShardedClusterSnapshotView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

