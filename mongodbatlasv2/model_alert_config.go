/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"time"
)

// checks if the AlertConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertConfig{}

// AlertConfig Alert settings allows to select which conditions trigger alerts and how users are notified.
type AlertConfig struct {
	// Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Flag that indicates whether someone enabled this alert configuration for the specified project.
	Enabled *bool `json:"enabled,omitempty"`
	EventTypeName *string `json:"eventTypeName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this alert configuration.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	Matchers []Matcher `json:"matchers,omitempty"`
	// List that contains the targets that MongoDB Cloud sends notifications.
	Notifications []Notification `json:"notifications,omitempty"`
	// Date and time when someone last updated this alert configuration. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewAlertConfig instantiates a new AlertConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertConfig() *AlertConfig {
	this := AlertConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewAlertConfigWithDefaults instantiates a new AlertConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertConfigWithDefaults() *AlertConfig {
	this := AlertConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AlertConfig) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfig) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AlertConfig) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AlertConfig) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AlertConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AlertConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AlertConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypeName returns the EventTypeName field value if set, zero value otherwise.
func (o *AlertConfig) GetEventTypeName() string {
	if o == nil || IsNil(o.EventTypeName) {
		var ret string
		return ret
	}
	return *o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfig) GetEventTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventTypeName) {
		return nil, false
	}
	return o.EventTypeName, true
}

// HasEventTypeName returns a boolean if a field has been set.
func (o *AlertConfig) HasEventTypeName() bool {
	if o != nil && !IsNil(o.EventTypeName) {
		return true
	}

	return false
}

// SetEventTypeName gets a reference to the given string and assigns it to the EventTypeName field.
func (o *AlertConfig) SetEventTypeName(v string) {
	o.EventTypeName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AlertConfig) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfig) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AlertConfig) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AlertConfig) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlertConfig) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlertConfig) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfig) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlertConfig) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *AlertConfig) SetLinks(v []Link) {
	o.Links = v
}

// GetMatchers returns the Matchers field value if set, zero value otherwise.
func (o *AlertConfig) GetMatchers() []Matcher {
	if o == nil || IsNil(o.Matchers) {
		var ret []Matcher
		return ret
	}
	return o.Matchers
}

// GetMatchersOk returns a tuple with the Matchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfig) GetMatchersOk() ([]Matcher, bool) {
	if o == nil || IsNil(o.Matchers) {
		return nil, false
	}
	return o.Matchers, true
}

// HasMatchers returns a boolean if a field has been set.
func (o *AlertConfig) HasMatchers() bool {
	if o != nil && !IsNil(o.Matchers) {
		return true
	}

	return false
}

// SetMatchers gets a reference to the given []Matcher and assigns it to the Matchers field.
func (o *AlertConfig) SetMatchers(v []Matcher) {
	o.Matchers = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *AlertConfig) GetNotifications() []Notification {
	if o == nil || IsNil(o.Notifications) {
		var ret []Notification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfig) GetNotificationsOk() ([]Notification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *AlertConfig) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []Notification and assigns it to the Notifications field.
func (o *AlertConfig) SetNotifications(v []Notification) {
	o.Notifications = v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AlertConfig) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertConfig) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AlertConfig) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *AlertConfig) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o AlertConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: created is readOnly
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.EventTypeName) {
		toSerialize["eventTypeName"] = o.EventTypeName
	}
	// skip: groupId is readOnly
	// skip: id is readOnly
	// skip: links is readOnly
	if !IsNil(o.Matchers) {
		toSerialize["matchers"] = o.Matchers
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	// skip: updated is readOnly
	return toSerialize, nil
}

type NullableAlertConfig struct {
	value *AlertConfig
	isSet bool
}

func (v NullableAlertConfig) Get() *AlertConfig {
	return v.value
}

func (v *NullableAlertConfig) Set(val *AlertConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertConfig(val *AlertConfig) *NullableAlertConfig {
	return &NullableAlertConfig{value: val, isSet: true}
}

func (v NullableAlertConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

