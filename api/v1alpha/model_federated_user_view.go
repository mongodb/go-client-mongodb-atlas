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

// checks if the FederatedUserView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederatedUserView{}

// FederatedUserView MongoDB Cloud user linked to this federated authentication.
type FederatedUserView struct {
	// Email address of the MongoDB Cloud user linked to the federated organization.
	EmailAddress string `json:"emailAddress"`
	// Unique 24-hexadecimal digit string that identifies the federation to which this MongoDB Cloud user belongs.
	FederationSettingsId string `json:"federationSettingsId"`
	// First or given name that belongs to the MongoDB Cloud user.
	FirstName string `json:"firstName"`
	// Last name, family name, or surname that belongs to the MongoDB Cloud user.
	LastName string `json:"lastName"`
	// Unique 24-hexadecimal digit string that identifies this user.
	UserId *string `json:"userId,omitempty"`
}

// NewFederatedUserView instantiates a new FederatedUserView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederatedUserView() *FederatedUserView {
	this := FederatedUserView{}
	return &this
}

// NewFederatedUserViewWithDefaults instantiates a new FederatedUserView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederatedUserViewWithDefaults() *FederatedUserView {
	this := FederatedUserView{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *FederatedUserView) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *FederatedUserView) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *FederatedUserView) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetFederationSettingsId returns the FederationSettingsId field value
func (o *FederatedUserView) GetFederationSettingsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FederationSettingsId
}

// GetFederationSettingsIdOk returns a tuple with the FederationSettingsId field value
// and a boolean to check if the value has been set.
func (o *FederatedUserView) GetFederationSettingsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FederationSettingsId, true
}

// SetFederationSettingsId sets field value
func (o *FederatedUserView) SetFederationSettingsId(v string) {
	o.FederationSettingsId = v
}

// GetFirstName returns the FirstName field value
func (o *FederatedUserView) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *FederatedUserView) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *FederatedUserView) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *FederatedUserView) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *FederatedUserView) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *FederatedUserView) SetLastName(v string) {
	o.LastName = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FederatedUserView) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedUserView) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FederatedUserView) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FederatedUserView) SetUserId(v string) {
	o.UserId = &v
}

func (o FederatedUserView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederatedUserView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailAddress"] = o.EmailAddress
	toSerialize["federationSettingsId"] = o.FederationSettingsId
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	// skip: userId is readOnly
	return toSerialize, nil
}

type NullableFederatedUserView struct {
	value *FederatedUserView
	isSet bool
}

func (v NullableFederatedUserView) Get() *FederatedUserView {
	return v.value
}

func (v *NullableFederatedUserView) Set(val *FederatedUserView) {
	v.value = val
	v.isSet = true
}

func (v NullableFederatedUserView) IsSet() bool {
	return v.isSet
}

func (v *NullableFederatedUserView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederatedUserView(val *FederatedUserView) *NullableFederatedUserView {
	return &NullableFederatedUserView{value: val, isSet: true}
}

func (v NullableFederatedUserView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederatedUserView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

