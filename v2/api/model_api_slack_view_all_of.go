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

// ApiSlackViewAllOf struct for ApiSlackViewAllOf
type ApiSlackViewAllOf struct {
	// Key that allows MongoDB Cloud to access your Slack account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.  **IMPORTANT**: Slack integrations now use the OAuth2 verification method and must  be initially configured, or updated from a legacy integration, through the Atlas  third-party service integrations page. Legacy tokens will soon no longer be  supported.  
	ApiToken *string `json:"apiToken,omitempty"`
	// Name of the Slack channel to which MongoDB Cloud sends alert notifications.
	ChannelName NullableString `json:"channelName,omitempty"`
	// Human-readable label that identifies your Slack team. Set this parameter when you configure a legacy Slack integration.
	TeamName *string `json:"teamName,omitempty"`
}

// NewApiSlackViewAllOf instantiates a new ApiSlackViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSlackViewAllOf() *ApiSlackViewAllOf {
	this := ApiSlackViewAllOf{}
	return &this
}

// NewApiSlackViewAllOfWithDefaults instantiates a new ApiSlackViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSlackViewAllOfWithDefaults() *ApiSlackViewAllOf {
	this := ApiSlackViewAllOf{}
	return &this
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise.
func (o *ApiSlackViewAllOf) GetApiToken() string {
	if o == nil || o.ApiToken == nil {
		var ret string
		return ret
	}
	return *o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSlackViewAllOf) GetApiTokenOk() (*string, bool) {
	if o == nil || o.ApiToken == nil {
		return nil, false
	}
	return o.ApiToken, true
}

// HasApiToken returns a boolean if a field has been set.
func (o *ApiSlackViewAllOf) HasApiToken() bool {
	if o != nil && o.ApiToken != nil {
		return true
	}

	return false
}

// SetApiToken gets a reference to the given string and assigns it to the ApiToken field.
func (o *ApiSlackViewAllOf) SetApiToken(v string) {
	o.ApiToken = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiSlackViewAllOf) GetChannelName() string {
	if o == nil || o.ChannelName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChannelName.Get()
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiSlackViewAllOf) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelName.Get(), o.ChannelName.IsSet()
}

// HasChannelName returns a boolean if a field has been set.
func (o *ApiSlackViewAllOf) HasChannelName() bool {
	if o != nil && o.ChannelName.IsSet() {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given NullableString and assigns it to the ChannelName field.
func (o *ApiSlackViewAllOf) SetChannelName(v string) {
	o.ChannelName.Set(&v)
}
// SetChannelNameNil sets the value for ChannelName to be an explicit nil
func (o *ApiSlackViewAllOf) SetChannelNameNil() {
	o.ChannelName.Set(nil)
}

// UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
func (o *ApiSlackViewAllOf) UnsetChannelName() {
	o.ChannelName.Unset()
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *ApiSlackViewAllOf) GetTeamName() string {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSlackViewAllOf) GetTeamNameOk() (*string, bool) {
	if o == nil || o.TeamName == nil {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *ApiSlackViewAllOf) HasTeamName() bool {
	if o != nil && o.TeamName != nil {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *ApiSlackViewAllOf) SetTeamName(v string) {
	o.TeamName = &v
}

func (o ApiSlackViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiToken != nil {
		toSerialize["apiToken"] = o.ApiToken
	}
	if o.ChannelName.IsSet() {
		toSerialize["channelName"] = o.ChannelName.Get()
	}
	if o.TeamName != nil {
		toSerialize["teamName"] = o.TeamName
	}
	return json.Marshal(toSerialize)
}

type NullableApiSlackViewAllOf struct {
	value *ApiSlackViewAllOf
	isSet bool
}

func (v NullableApiSlackViewAllOf) Get() *ApiSlackViewAllOf {
	return v.value
}

func (v *NullableApiSlackViewAllOf) Set(val *ApiSlackViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSlackViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSlackViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSlackViewAllOf(val *ApiSlackViewAllOf) *NullableApiSlackViewAllOf {
	return &NullableApiSlackViewAllOf{value: val, isSet: true}
}

func (v NullableApiSlackViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSlackViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

