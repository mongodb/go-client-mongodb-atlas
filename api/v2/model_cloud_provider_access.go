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

// CloudProviderAccess struct for CloudProviderAccess
type CloudProviderAccess struct {
	// List that contains the Amazon Web Services (AWS) IAM roles registered and authorized with MongoDB Cloud.
	AwsIamRoles []CloudProviderAccessAWSIAMRole `json:"awsIamRoles,omitempty"`
}

// NewCloudProviderAccess instantiates a new CloudProviderAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccess() *CloudProviderAccess {
	this := CloudProviderAccess{}
	return &this
}

// NewCloudProviderAccessWithDefaults instantiates a new CloudProviderAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessWithDefaults() *CloudProviderAccess {
	this := CloudProviderAccess{}
	return &this
}

// GetAwsIamRoles returns the AwsIamRoles field value if set, zero value otherwise.
func (o *CloudProviderAccess) GetAwsIamRoles() []CloudProviderAccessAWSIAMRole {
	if o == nil || o.AwsIamRoles == nil {
		var ret []CloudProviderAccessAWSIAMRole
		return ret
	}
	return o.AwsIamRoles
}

// GetAwsIamRolesOk returns a tuple with the AwsIamRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccess) GetAwsIamRolesOk() ([]CloudProviderAccessAWSIAMRole, bool) {
	if o == nil || o.AwsIamRoles == nil {
		return nil, false
	}
	return o.AwsIamRoles, true
}

// HasAwsIamRoles returns a boolean if a field has been set.
func (o *CloudProviderAccess) HasAwsIamRoles() bool {
	if o != nil && o.AwsIamRoles != nil {
		return true
	}

	return false
}

// SetAwsIamRoles gets a reference to the given []CloudProviderAccessAWSIAMRole and assigns it to the AwsIamRoles field.
func (o *CloudProviderAccess) SetAwsIamRoles(v []CloudProviderAccessAWSIAMRole) {
	o.AwsIamRoles = v
}

func (o CloudProviderAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsIamRoles != nil {
		toSerialize["awsIamRoles"] = o.AwsIamRoles
	}
	return json.Marshal(toSerialize)
}

type NullableCloudProviderAccess struct {
	value *CloudProviderAccess
	isSet bool
}

func (v NullableCloudProviderAccess) Get() *CloudProviderAccess {
	return v.value
}

func (v *NullableCloudProviderAccess) Set(val *CloudProviderAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccess(val *CloudProviderAccess) *NullableCloudProviderAccess {
	return &NullableCloudProviderAccess{value: val, isSet: true}
}

func (v NullableCloudProviderAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

