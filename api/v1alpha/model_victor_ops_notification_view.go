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

// checks if the VictorOpsNotificationView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VictorOpsNotificationView{}

// VictorOpsNotificationView VictorOps notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type VictorOpsNotificationView struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
	// API key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when `\"notifications.[n].typeName\" : \"VICTOR_OPS\"`. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	VictorOpsApiKey *string `json:"victorOpsApiKey,omitempty"`
	// Routing key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when `\"notifications.[n].typeName\" : \"VICTOR_OPS\"`. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	VictorOpsRoutingKey *string `json:"victorOpsRoutingKey,omitempty"`
}

// NewVictorOpsNotificationView instantiates a new VictorOpsNotificationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVictorOpsNotificationView() *VictorOpsNotificationView {
	this := VictorOpsNotificationView{}
	return &this
}

// NewVictorOpsNotificationViewWithDefaults instantiates a new VictorOpsNotificationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVictorOpsNotificationViewWithDefaults() *VictorOpsNotificationView {
	this := VictorOpsNotificationView{}
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *VictorOpsNotificationView) GetDelayMin() int32 {
	if o == nil || IsNil(o.DelayMin) {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOpsNotificationView) GetDelayMinOk() (*int32, bool) {
	if o == nil || IsNil(o.DelayMin) {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *VictorOpsNotificationView) HasDelayMin() bool {
	if o != nil && !IsNil(o.DelayMin) {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *VictorOpsNotificationView) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *VictorOpsNotificationView) GetIntervalMin() int32 {
	if o == nil || IsNil(o.IntervalMin) {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOpsNotificationView) GetIntervalMinOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalMin) {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *VictorOpsNotificationView) HasIntervalMin() bool {
	if o != nil && !IsNil(o.IntervalMin) {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *VictorOpsNotificationView) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetTypeName returns the TypeName field value
func (o *VictorOpsNotificationView) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *VictorOpsNotificationView) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *VictorOpsNotificationView) SetTypeName(v string) {
	o.TypeName = v
}

// GetVictorOpsApiKey returns the VictorOpsApiKey field value if set, zero value otherwise.
func (o *VictorOpsNotificationView) GetVictorOpsApiKey() string {
	if o == nil || IsNil(o.VictorOpsApiKey) {
		var ret string
		return ret
	}
	return *o.VictorOpsApiKey
}

// GetVictorOpsApiKeyOk returns a tuple with the VictorOpsApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOpsNotificationView) GetVictorOpsApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.VictorOpsApiKey) {
		return nil, false
	}
	return o.VictorOpsApiKey, true
}

// HasVictorOpsApiKey returns a boolean if a field has been set.
func (o *VictorOpsNotificationView) HasVictorOpsApiKey() bool {
	if o != nil && !IsNil(o.VictorOpsApiKey) {
		return true
	}

	return false
}

// SetVictorOpsApiKey gets a reference to the given string and assigns it to the VictorOpsApiKey field.
func (o *VictorOpsNotificationView) SetVictorOpsApiKey(v string) {
	o.VictorOpsApiKey = &v
}

// GetVictorOpsRoutingKey returns the VictorOpsRoutingKey field value if set, zero value otherwise.
func (o *VictorOpsNotificationView) GetVictorOpsRoutingKey() string {
	if o == nil || IsNil(o.VictorOpsRoutingKey) {
		var ret string
		return ret
	}
	return *o.VictorOpsRoutingKey
}

// GetVictorOpsRoutingKeyOk returns a tuple with the VictorOpsRoutingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOpsNotificationView) GetVictorOpsRoutingKeyOk() (*string, bool) {
	if o == nil || IsNil(o.VictorOpsRoutingKey) {
		return nil, false
	}
	return o.VictorOpsRoutingKey, true
}

// HasVictorOpsRoutingKey returns a boolean if a field has been set.
func (o *VictorOpsNotificationView) HasVictorOpsRoutingKey() bool {
	if o != nil && !IsNil(o.VictorOpsRoutingKey) {
		return true
	}

	return false
}

// SetVictorOpsRoutingKey gets a reference to the given string and assigns it to the VictorOpsRoutingKey field.
func (o *VictorOpsNotificationView) SetVictorOpsRoutingKey(v string) {
	o.VictorOpsRoutingKey = &v
}

func (o VictorOpsNotificationView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VictorOpsNotificationView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelayMin) {
		toSerialize["delayMin"] = o.DelayMin
	}
	if !IsNil(o.IntervalMin) {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	toSerialize["typeName"] = o.TypeName
	if !IsNil(o.VictorOpsApiKey) {
		toSerialize["victorOpsApiKey"] = o.VictorOpsApiKey
	}
	if !IsNil(o.VictorOpsRoutingKey) {
		toSerialize["victorOpsRoutingKey"] = o.VictorOpsRoutingKey
	}
	return toSerialize, nil
}

type NullableVictorOpsNotificationView struct {
	value *VictorOpsNotificationView
	isSet bool
}

func (v NullableVictorOpsNotificationView) Get() *VictorOpsNotificationView {
	return v.value
}

func (v *NullableVictorOpsNotificationView) Set(val *VictorOpsNotificationView) {
	v.value = val
	v.isSet = true
}

func (v NullableVictorOpsNotificationView) IsSet() bool {
	return v.isSet
}

func (v *NullableVictorOpsNotificationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVictorOpsNotificationView(val *VictorOpsNotificationView) *NullableVictorOpsNotificationView {
	return &NullableVictorOpsNotificationView{value: val, isSet: true}
}

func (v NullableVictorOpsNotificationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVictorOpsNotificationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

