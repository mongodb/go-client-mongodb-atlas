/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"time"
)

// checks if the ApiUserEventViewForOrg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUserEventViewForOrg{}

// ApiUserEventViewForOrg API User event identifies different activities around user api keys.
type ApiUserEventViewForOrg struct {
	// Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **userId** parameter.
	ApiKeyId *string `json:"apiKeyId,omitempty"`
	// Date and time when this event occurred. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created time.Time `json:"created"`
	EventTypeName ApiUserEventTypeViewForOrg `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the event.
	Id string `json:"id"`
	// Flag that indicates whether a MongoDB employee triggered the specified event.
	IsGlobalAdmin *bool `json:"isGlobalAdmin,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization to which these events apply.
	OrgId *string `json:"orgId,omitempty"`
	// Public part of the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **username** parameter.
	PublicKey *string `json:"publicKey,omitempty"`
	Raw *Raw `json:"raw,omitempty"`
	// IPv4 or IPv6 address from which the user triggered this event.
	RemoteAddress *string `json:"remoteAddress,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn't return the **apiKeyId** parameter.
	UserId *string `json:"userId,omitempty"`
	// Email address for the user who triggered this event. If this resource returns this parameter, it doesn't return the **publicApiKey** parameter.
	Username *string `json:"username,omitempty"`
}

// NewApiUserEventViewForOrg instantiates a new ApiUserEventViewForOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserEventViewForOrg() *ApiUserEventViewForOrg {
	this := ApiUserEventViewForOrg{}
	return &this
}

// NewApiUserEventViewForOrgWithDefaults instantiates a new ApiUserEventViewForOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserEventViewForOrgWithDefaults() *ApiUserEventViewForOrg {
	this := ApiUserEventViewForOrg{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise.
func (o *ApiUserEventViewForOrg) GetApiKeyId() string {
	if o == nil || IsNil(o.ApiKeyId) {
		var ret string
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetApiKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKeyId) {
		return nil, false
	}
	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *ApiUserEventViewForOrg) HasApiKeyId() bool {
	if o != nil && !IsNil(o.ApiKeyId) {
		return true
	}

	return false
}

// SetApiKeyId gets a reference to the given string and assigns it to the ApiKeyId field.
func (o *ApiUserEventViewForOrg) SetApiKeyId(v string) {
	o.ApiKeyId = &v
}

// GetCreated returns the Created field value
func (o *ApiUserEventViewForOrg) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ApiUserEventViewForOrg) SetCreated(v time.Time) {
	o.Created = v
}

// GetEventTypeName returns the EventTypeName field value
func (o *ApiUserEventViewForOrg) GetEventTypeName() ApiUserEventTypeViewForOrg {
	if o == nil {
		var ret ApiUserEventTypeViewForOrg
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetEventTypeNameOk() (*ApiUserEventTypeViewForOrg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *ApiUserEventViewForOrg) SetEventTypeName(v ApiUserEventTypeViewForOrg) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiUserEventViewForOrg) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiUserEventViewForOrg) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiUserEventViewForOrg) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value
func (o *ApiUserEventViewForOrg) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiUserEventViewForOrg) SetId(v string) {
	o.Id = v
}

// GetIsGlobalAdmin returns the IsGlobalAdmin field value if set, zero value otherwise.
func (o *ApiUserEventViewForOrg) GetIsGlobalAdmin() bool {
	if o == nil || IsNil(o.IsGlobalAdmin) {
		var ret bool
		return ret
	}
	return *o.IsGlobalAdmin
}

// GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetIsGlobalAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobalAdmin) {
		return nil, false
	}
	return o.IsGlobalAdmin, true
}

// HasIsGlobalAdmin returns a boolean if a field has been set.
func (o *ApiUserEventViewForOrg) HasIsGlobalAdmin() bool {
	if o != nil && !IsNil(o.IsGlobalAdmin) {
		return true
	}

	return false
}

// SetIsGlobalAdmin gets a reference to the given bool and assigns it to the IsGlobalAdmin field.
func (o *ApiUserEventViewForOrg) SetIsGlobalAdmin(v bool) {
	o.IsGlobalAdmin = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiUserEventViewForOrg) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiUserEventViewForOrg) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ApiUserEventViewForOrg) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ApiUserEventViewForOrg) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ApiUserEventViewForOrg) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ApiUserEventViewForOrg) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *ApiUserEventViewForOrg) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *ApiUserEventViewForOrg) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *ApiUserEventViewForOrg) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *ApiUserEventViewForOrg) GetRaw() Raw {
	if o == nil || IsNil(o.Raw) {
		var ret Raw
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetRawOk() (*Raw, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *ApiUserEventViewForOrg) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given Raw and assigns it to the Raw field.
func (o *ApiUserEventViewForOrg) SetRaw(v Raw) {
	o.Raw = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *ApiUserEventViewForOrg) GetRemoteAddress() string {
	if o == nil || IsNil(o.RemoteAddress) {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetRemoteAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteAddress) {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *ApiUserEventViewForOrg) HasRemoteAddress() bool {
	if o != nil && !IsNil(o.RemoteAddress) {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *ApiUserEventViewForOrg) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ApiUserEventViewForOrg) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ApiUserEventViewForOrg) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ApiUserEventViewForOrg) SetUserId(v string) {
	o.UserId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiUserEventViewForOrg) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserEventViewForOrg) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiUserEventViewForOrg) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiUserEventViewForOrg) SetUsername(v string) {
	o.Username = &v
}

func (o ApiUserEventViewForOrg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUserEventViewForOrg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: apiKeyId is readOnly
	// skip: created is readOnly
	toSerialize["eventTypeName"] = o.EventTypeName
	// skip: groupId is readOnly
	// skip: id is readOnly
	// skip: isGlobalAdmin is readOnly
	// skip: links is readOnly
	// skip: orgId is readOnly
	// skip: publicKey is readOnly
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	// skip: remoteAddress is readOnly
	// skip: userId is readOnly
	// skip: username is readOnly
	return toSerialize, nil
}

type NullableApiUserEventViewForOrg struct {
	value *ApiUserEventViewForOrg
	isSet bool
}

func (v NullableApiUserEventViewForOrg) Get() *ApiUserEventViewForOrg {
	return v.value
}

func (v *NullableApiUserEventViewForOrg) Set(val *ApiUserEventViewForOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserEventViewForOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserEventViewForOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserEventViewForOrg(val *ApiUserEventViewForOrg) *NullableApiUserEventViewForOrg {
	return &NullableApiUserEventViewForOrg{value: val, isSet: true}
}

func (v NullableApiUserEventViewForOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserEventViewForOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

