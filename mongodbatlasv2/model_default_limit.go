/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the DefaultLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultLimit{}

// DefaultLimit Details of user managed limits
type DefaultLimit struct {
	// Amount that indicates the current usage of the limit.
	CurrentUsage *int64 `json:"currentUsage,omitempty"`
	// Default value of the limit.
	DefaultLimit *int64 `json:"defaultLimit,omitempty"`
	// Maximum value of the limit.
	MaximumLimit *int64 `json:"maximumLimit,omitempty"`
	// Human-readable label that identifies the user-managed limit to modify.
	Name string `json:"name"`
	// Amount to set the limit to.
	Value int64 `json:"value"`
}

// NewDefaultLimit instantiates a new DefaultLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultLimit(name string, value int64) *DefaultLimit {
	this := DefaultLimit{}
	this.Name = name
	this.Value = value
	return &this
}

// NewDefaultLimitWithDefaults instantiates a new DefaultLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultLimitWithDefaults() *DefaultLimit {
	this := DefaultLimit{}
	return &this
}

// GetCurrentUsage returns the CurrentUsage field value if set, zero value otherwise.
func (o *DefaultLimit) GetCurrentUsage() int64 {
	if o == nil || IsNil(o.CurrentUsage) {
		var ret int64
		return ret
	}
	return *o.CurrentUsage
}

// GetCurrentUsageOk returns a tuple with the CurrentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultLimit) GetCurrentUsageOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentUsage) {
		return nil, false
	}
	return o.CurrentUsage, true
}

// HasCurrentUsage returns a boolean if a field has been set.
func (o *DefaultLimit) HasCurrentUsage() bool {
	if o != nil && !IsNil(o.CurrentUsage) {
		return true
	}

	return false
}

// SetCurrentUsage gets a reference to the given int64 and assigns it to the CurrentUsage field.
func (o *DefaultLimit) SetCurrentUsage(v int64) {
	o.CurrentUsage = &v
}

// GetDefaultLimit returns the DefaultLimit field value if set, zero value otherwise.
func (o *DefaultLimit) GetDefaultLimit() int64 {
	if o == nil || IsNil(o.DefaultLimit) {
		var ret int64
		return ret
	}
	return *o.DefaultLimit
}

// GetDefaultLimitOk returns a tuple with the DefaultLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultLimit) GetDefaultLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultLimit) {
		return nil, false
	}
	return o.DefaultLimit, true
}

// HasDefaultLimit returns a boolean if a field has been set.
func (o *DefaultLimit) HasDefaultLimit() bool {
	if o != nil && !IsNil(o.DefaultLimit) {
		return true
	}

	return false
}

// SetDefaultLimit gets a reference to the given int64 and assigns it to the DefaultLimit field.
func (o *DefaultLimit) SetDefaultLimit(v int64) {
	o.DefaultLimit = &v
}

// GetMaximumLimit returns the MaximumLimit field value if set, zero value otherwise.
func (o *DefaultLimit) GetMaximumLimit() int64 {
	if o == nil || IsNil(o.MaximumLimit) {
		var ret int64
		return ret
	}
	return *o.MaximumLimit
}

// GetMaximumLimitOk returns a tuple with the MaximumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultLimit) GetMaximumLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumLimit) {
		return nil, false
	}
	return o.MaximumLimit, true
}

// HasMaximumLimit returns a boolean if a field has been set.
func (o *DefaultLimit) HasMaximumLimit() bool {
	if o != nil && !IsNil(o.MaximumLimit) {
		return true
	}

	return false
}

// SetMaximumLimit gets a reference to the given int64 and assigns it to the MaximumLimit field.
func (o *DefaultLimit) SetMaximumLimit(v int64) {
	o.MaximumLimit = &v
}

// GetName returns the Name field value
func (o *DefaultLimit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DefaultLimit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DefaultLimit) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *DefaultLimit) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DefaultLimit) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DefaultLimit) SetValue(v int64) {
	o.Value = v
}

func (o DefaultLimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: currentUsage is readOnly
	// skip: defaultLimit is readOnly
	// skip: maximumLimit is readOnly
	// skip: name is readOnly
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableDefaultLimit struct {
	value *DefaultLimit
	isSet bool
}

func (v NullableDefaultLimit) Get() *DefaultLimit {
	return v.value
}

func (v *NullableDefaultLimit) Set(val *DefaultLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultLimit(val *DefaultLimit) *NullableDefaultLimit {
	return &NullableDefaultLimit{value: val, isSet: true}
}

func (v NullableDefaultLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

