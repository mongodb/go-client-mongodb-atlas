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

// checks if the PaginatedCloudBackupReplicaSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedCloudBackupReplicaSet{}

// PaginatedCloudBackupReplicaSet struct for PaginatedCloudBackupReplicaSet
type PaginatedCloudBackupReplicaSet struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []DiskBackupReplicaSet `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedCloudBackupReplicaSet instantiates a new PaginatedCloudBackupReplicaSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCloudBackupReplicaSet() *PaginatedCloudBackupReplicaSet {
	this := PaginatedCloudBackupReplicaSet{}
	return &this
}

// NewPaginatedCloudBackupReplicaSetWithDefaults instantiates a new PaginatedCloudBackupReplicaSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCloudBackupReplicaSetWithDefaults() *PaginatedCloudBackupReplicaSet {
	this := PaginatedCloudBackupReplicaSet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedCloudBackupReplicaSet) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupReplicaSet) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedCloudBackupReplicaSet) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedCloudBackupReplicaSet) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedCloudBackupReplicaSet) GetResults() []DiskBackupReplicaSet {
	if o == nil || IsNil(o.Results) {
		var ret []DiskBackupReplicaSet
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupReplicaSet) GetResultsOk() ([]DiskBackupReplicaSet, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedCloudBackupReplicaSet) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DiskBackupReplicaSet and assigns it to the Results field.
func (o *PaginatedCloudBackupReplicaSet) SetResults(v []DiskBackupReplicaSet) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedCloudBackupReplicaSet) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupReplicaSet) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedCloudBackupReplicaSet) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedCloudBackupReplicaSet) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedCloudBackupReplicaSet) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedCloudBackupReplicaSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePaginatedCloudBackupReplicaSet struct {
	value *PaginatedCloudBackupReplicaSet
	isSet bool
}

func (v NullablePaginatedCloudBackupReplicaSet) Get() *PaginatedCloudBackupReplicaSet {
	return v.value
}

func (v *NullablePaginatedCloudBackupReplicaSet) Set(val *PaginatedCloudBackupReplicaSet) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCloudBackupReplicaSet) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCloudBackupReplicaSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCloudBackupReplicaSet(val *PaginatedCloudBackupReplicaSet) *NullablePaginatedCloudBackupReplicaSet {
	return &NullablePaginatedCloudBackupReplicaSet{value: val, isSet: true}
}

func (v NullablePaginatedCloudBackupReplicaSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCloudBackupReplicaSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

