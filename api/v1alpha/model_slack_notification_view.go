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

// SlackNotificationView Slack notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type SlackNotificationView struct {
	// Slack API token or Bot token that MongoDB Cloud needs to send alert notifications via Slack. The resource requires this parameter when `\"notifications.[n].typeName\" : \"SLACK\"`. If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes the token.   **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ApiToken *string `json:"apiToken,omitempty"`
	// Name of the Slack channel to which MongoDB Cloud sends alert notifications. The resource requires this parameter when `\"notifications.[n].typeName\" : \"SLACK\"`.
	ChannelName *string `json:"channelName,omitempty"`
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewSlackNotificationView instantiates a new SlackNotificationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackNotificationView() *SlackNotificationView {
	this := SlackNotificationView{}
	return &this
}

// NewSlackNotificationViewWithDefaults instantiates a new SlackNotificationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackNotificationViewWithDefaults() *SlackNotificationView {
	this := SlackNotificationView{}
	return &this
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise.
func (o *SlackNotificationView) GetApiToken() string {
	if o == nil || o.ApiToken == nil {
		var ret string
		return ret
	}
	return *o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackNotificationView) GetApiTokenOk() (*string, bool) {
	if o == nil || o.ApiToken == nil {
		return nil, false
	}
	return o.ApiToken, true
}

// HasApiToken returns a boolean if a field has been set.
func (o *SlackNotificationView) HasApiToken() bool {
	if o != nil && o.ApiToken != nil {
		return true
	}

	return false
}

// SetApiToken gets a reference to the given string and assigns it to the ApiToken field.
func (o *SlackNotificationView) SetApiToken(v string) {
	o.ApiToken = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *SlackNotificationView) GetChannelName() string {
	if o == nil || o.ChannelName == nil {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackNotificationView) GetChannelNameOk() (*string, bool) {
	if o == nil || o.ChannelName == nil {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *SlackNotificationView) HasChannelName() bool {
	if o != nil && o.ChannelName != nil {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *SlackNotificationView) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *SlackNotificationView) GetDelayMin() int32 {
	if o == nil || o.DelayMin == nil {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackNotificationView) GetDelayMinOk() (*int32, bool) {
	if o == nil || o.DelayMin == nil {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *SlackNotificationView) HasDelayMin() bool {
	if o != nil && o.DelayMin != nil {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *SlackNotificationView) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *SlackNotificationView) GetIntervalMin() int32 {
	if o == nil || o.IntervalMin == nil {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackNotificationView) GetIntervalMinOk() (*int32, bool) {
	if o == nil || o.IntervalMin == nil {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *SlackNotificationView) HasIntervalMin() bool {
	if o != nil && o.IntervalMin != nil {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *SlackNotificationView) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetTypeName returns the TypeName field value
func (o *SlackNotificationView) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *SlackNotificationView) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *SlackNotificationView) SetTypeName(v string) {
	o.TypeName = v
}

func (o SlackNotificationView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiToken != nil {
		toSerialize["apiToken"] = o.ApiToken
	}
	if o.ChannelName != nil {
		toSerialize["channelName"] = o.ChannelName
	}
	if o.DelayMin != nil {
		toSerialize["delayMin"] = o.DelayMin
	}
	if o.IntervalMin != nil {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	if true {
		toSerialize["typeName"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableSlackNotificationView struct {
	value *SlackNotificationView
	isSet bool
}

func (v NullableSlackNotificationView) Get() *SlackNotificationView {
	return v.value
}

func (v *NullableSlackNotificationView) Set(val *SlackNotificationView) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackNotificationView) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackNotificationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackNotificationView(val *SlackNotificationView) *NullableSlackNotificationView {
	return &NullableSlackNotificationView{value: val, isSet: true}
}

func (v NullableSlackNotificationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackNotificationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

