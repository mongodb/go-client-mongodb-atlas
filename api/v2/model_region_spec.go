/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// RegionSpec Physical location where MongoDB Cloud provisions cluster nodes.
type RegionSpec struct {
	// Number of analytics nodes in the region. Analytics nodes handle analytic data such as reporting queries from MongoDB Connector for Business Intelligence on MongoDB Cloud. Analytics nodes are read-only, and can never become the primary. Use **replicationSpecs[n].{region}.analyticsNodes** instead.
	AnalyticsNodes *int32 `json:"analyticsNodes,omitempty"`
	// Number of electable nodes to deploy in the specified region. Electable nodes can become the primary and can facilitate local reads. Use **replicationSpecs[n].{region}.electableNodes** instead.
	ElectableNodes *int32 `json:"electableNodes,omitempty"`
	// Number that indicates the election priority of the region. To identify the Preferred Region of the cluster, set this parameter to `7`. The primary node runs in the **Preferred Region**. To identify a read-only region, set this parameter to `0`.
	Priority *int32 `json:"priority,omitempty"`
	// Number of read-only nodes in the region. Read-only nodes can never become the primary member, but can facilitate local reads. Use **replicationSpecs[n].{region}.readOnlyNodes** instead.
	ReadOnlyNodes *int32 `json:"readOnlyNodes,omitempty"`
}

// NewRegionSpec instantiates a new RegionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionSpec() *RegionSpec {
	this := RegionSpec{}
	return &this
}

// NewRegionSpecWithDefaults instantiates a new RegionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionSpecWithDefaults() *RegionSpec {
	this := RegionSpec{}
	return &this
}

// GetAnalyticsNodes returns the AnalyticsNodes field value if set, zero value otherwise.
func (o *RegionSpec) GetAnalyticsNodes() int32 {
	if o == nil || o.AnalyticsNodes == nil {
		var ret int32
		return ret
	}
	return *o.AnalyticsNodes
}

// GetAnalyticsNodesOk returns a tuple with the AnalyticsNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSpec) GetAnalyticsNodesOk() (*int32, bool) {
	if o == nil || o.AnalyticsNodes == nil {
		return nil, false
	}
	return o.AnalyticsNodes, true
}

// HasAnalyticsNodes returns a boolean if a field has been set.
func (o *RegionSpec) HasAnalyticsNodes() bool {
	if o != nil && o.AnalyticsNodes != nil {
		return true
	}

	return false
}

// SetAnalyticsNodes gets a reference to the given int32 and assigns it to the AnalyticsNodes field.
func (o *RegionSpec) SetAnalyticsNodes(v int32) {
	o.AnalyticsNodes = &v
}

// GetElectableNodes returns the ElectableNodes field value if set, zero value otherwise.
func (o *RegionSpec) GetElectableNodes() int32 {
	if o == nil || o.ElectableNodes == nil {
		var ret int32
		return ret
	}
	return *o.ElectableNodes
}

// GetElectableNodesOk returns a tuple with the ElectableNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSpec) GetElectableNodesOk() (*int32, bool) {
	if o == nil || o.ElectableNodes == nil {
		return nil, false
	}
	return o.ElectableNodes, true
}

// HasElectableNodes returns a boolean if a field has been set.
func (o *RegionSpec) HasElectableNodes() bool {
	if o != nil && o.ElectableNodes != nil {
		return true
	}

	return false
}

// SetElectableNodes gets a reference to the given int32 and assigns it to the ElectableNodes field.
func (o *RegionSpec) SetElectableNodes(v int32) {
	o.ElectableNodes = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RegionSpec) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSpec) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RegionSpec) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *RegionSpec) SetPriority(v int32) {
	o.Priority = &v
}

// GetReadOnlyNodes returns the ReadOnlyNodes field value if set, zero value otherwise.
func (o *RegionSpec) GetReadOnlyNodes() int32 {
	if o == nil || o.ReadOnlyNodes == nil {
		var ret int32
		return ret
	}
	return *o.ReadOnlyNodes
}

// GetReadOnlyNodesOk returns a tuple with the ReadOnlyNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSpec) GetReadOnlyNodesOk() (*int32, bool) {
	if o == nil || o.ReadOnlyNodes == nil {
		return nil, false
	}
	return o.ReadOnlyNodes, true
}

// HasReadOnlyNodes returns a boolean if a field has been set.
func (o *RegionSpec) HasReadOnlyNodes() bool {
	if o != nil && o.ReadOnlyNodes != nil {
		return true
	}

	return false
}

// SetReadOnlyNodes gets a reference to the given int32 and assigns it to the ReadOnlyNodes field.
func (o *RegionSpec) SetReadOnlyNodes(v int32) {
	o.ReadOnlyNodes = &v
}

func (o RegionSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnalyticsNodes != nil {
		toSerialize["analyticsNodes"] = o.AnalyticsNodes
	}
	if o.ElectableNodes != nil {
		toSerialize["electableNodes"] = o.ElectableNodes
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.ReadOnlyNodes != nil {
		toSerialize["readOnlyNodes"] = o.ReadOnlyNodes
	}
	return json.Marshal(toSerialize)
}

type NullableRegionSpec struct {
	value *RegionSpec
	isSet bool
}

func (v NullableRegionSpec) Get() *RegionSpec {
	return v.value
}

func (v *NullableRegionSpec) Set(val *RegionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionSpec(val *RegionSpec) *NullableRegionSpec {
	return &NullableRegionSpec{value: val, isSet: true}
}

func (v NullableRegionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

