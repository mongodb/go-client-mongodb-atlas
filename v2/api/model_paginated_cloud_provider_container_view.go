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

// PaginatedCloudProviderContainerView List of Network Peering Containers that Amazon Web Services serves.
type PaginatedCloudProviderContainerView struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []CloudProviderContainer `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedCloudProviderContainerView instantiates a new PaginatedCloudProviderContainerView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCloudProviderContainerView() *PaginatedCloudProviderContainerView {
	this := PaginatedCloudProviderContainerView{}
	return &this
}

// NewPaginatedCloudProviderContainerViewWithDefaults instantiates a new PaginatedCloudProviderContainerView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCloudProviderContainerViewWithDefaults() *PaginatedCloudProviderContainerView {
	this := PaginatedCloudProviderContainerView{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedCloudProviderContainerView) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudProviderContainerView) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedCloudProviderContainerView) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedCloudProviderContainerView) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedCloudProviderContainerView) GetResults() []CloudProviderContainer {
	if o == nil || o.Results == nil {
		var ret []CloudProviderContainer
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudProviderContainerView) GetResultsOk() ([]CloudProviderContainer, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedCloudProviderContainerView) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CloudProviderContainer and assigns it to the Results field.
func (o *PaginatedCloudProviderContainerView) SetResults(v []CloudProviderContainer) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedCloudProviderContainerView) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedCloudProviderContainerView) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedCloudProviderContainerView) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedCloudProviderContainerView) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedCloudProviderContainerView) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedCloudProviderContainerView struct {
	value *PaginatedCloudProviderContainerView
	isSet bool
}

func (v NullablePaginatedCloudProviderContainerView) Get() *PaginatedCloudProviderContainerView {
	return v.value
}

func (v *NullablePaginatedCloudProviderContainerView) Set(val *PaginatedCloudProviderContainerView) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCloudProviderContainerView) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCloudProviderContainerView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCloudProviderContainerView(val *PaginatedCloudProviderContainerView) *NullablePaginatedCloudProviderContainerView {
	return &NullablePaginatedCloudProviderContainerView{value: val, isSet: true}
}

func (v NullablePaginatedCloudProviderContainerView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCloudProviderContainerView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

