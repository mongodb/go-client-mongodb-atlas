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

// ServerlessAWSTenantEndpoint View for a serverless AWS tenant endpoint.
type ServerlessAWSTenantEndpoint struct {
	// Unique 24-hexadecimal digit string that identifies the private endpoint.
	Id *string `json:"_id,omitempty"`
	// Unique string that identifies the private endpoint's network interface.
	CloudProviderEndpointId *string `json:"cloudProviderEndpointId,omitempty"`
	// Human-readable comment associated with the private endpoint.
	Comment *string `json:"comment,omitempty"`
	// Unique string that identifies the Amazon Web Services (AWS) PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service.
	EndpointServiceName *string `json:"endpointServiceName,omitempty"`
	// Human-readable error message that indicates error condition associated with establishing the private endpoint connection.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Human-readable label that identifies the cloud service provider.
	ProviderName *string `json:"providerName,omitempty"`
	// Human-readable label that indicates the current operating status of the private endpoint.
	Status *string `json:"status,omitempty"`
}

// NewServerlessAWSTenantEndpoint instantiates a new ServerlessAWSTenantEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessAWSTenantEndpoint() *ServerlessAWSTenantEndpoint {
	this := ServerlessAWSTenantEndpoint{}
	return &this
}

// NewServerlessAWSTenantEndpointWithDefaults instantiates a new ServerlessAWSTenantEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessAWSTenantEndpointWithDefaults() *ServerlessAWSTenantEndpoint {
	this := ServerlessAWSTenantEndpoint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerlessAWSTenantEndpoint) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAWSTenantEndpoint) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerlessAWSTenantEndpoint) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerlessAWSTenantEndpoint) SetId(v string) {
	o.Id = &v
}

// GetCloudProviderEndpointId returns the CloudProviderEndpointId field value if set, zero value otherwise.
func (o *ServerlessAWSTenantEndpoint) GetCloudProviderEndpointId() string {
	if o == nil || o.CloudProviderEndpointId == nil {
		var ret string
		return ret
	}
	return *o.CloudProviderEndpointId
}

// GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAWSTenantEndpoint) GetCloudProviderEndpointIdOk() (*string, bool) {
	if o == nil || o.CloudProviderEndpointId == nil {
		return nil, false
	}
	return o.CloudProviderEndpointId, true
}

// HasCloudProviderEndpointId returns a boolean if a field has been set.
func (o *ServerlessAWSTenantEndpoint) HasCloudProviderEndpointId() bool {
	if o != nil && o.CloudProviderEndpointId != nil {
		return true
	}

	return false
}

// SetCloudProviderEndpointId gets a reference to the given string and assigns it to the CloudProviderEndpointId field.
func (o *ServerlessAWSTenantEndpoint) SetCloudProviderEndpointId(v string) {
	o.CloudProviderEndpointId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ServerlessAWSTenantEndpoint) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAWSTenantEndpoint) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ServerlessAWSTenantEndpoint) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ServerlessAWSTenantEndpoint) SetComment(v string) {
	o.Comment = &v
}

// GetEndpointServiceName returns the EndpointServiceName field value if set, zero value otherwise.
func (o *ServerlessAWSTenantEndpoint) GetEndpointServiceName() string {
	if o == nil || o.EndpointServiceName == nil {
		var ret string
		return ret
	}
	return *o.EndpointServiceName
}

// GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAWSTenantEndpoint) GetEndpointServiceNameOk() (*string, bool) {
	if o == nil || o.EndpointServiceName == nil {
		return nil, false
	}
	return o.EndpointServiceName, true
}

// HasEndpointServiceName returns a boolean if a field has been set.
func (o *ServerlessAWSTenantEndpoint) HasEndpointServiceName() bool {
	if o != nil && o.EndpointServiceName != nil {
		return true
	}

	return false
}

// SetEndpointServiceName gets a reference to the given string and assigns it to the EndpointServiceName field.
func (o *ServerlessAWSTenantEndpoint) SetEndpointServiceName(v string) {
	o.EndpointServiceName = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ServerlessAWSTenantEndpoint) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAWSTenantEndpoint) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ServerlessAWSTenantEndpoint) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ServerlessAWSTenantEndpoint) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *ServerlessAWSTenantEndpoint) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAWSTenantEndpoint) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *ServerlessAWSTenantEndpoint) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *ServerlessAWSTenantEndpoint) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServerlessAWSTenantEndpoint) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAWSTenantEndpoint) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServerlessAWSTenantEndpoint) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ServerlessAWSTenantEndpoint) SetStatus(v string) {
	o.Status = &v
}

func (o ServerlessAWSTenantEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.CloudProviderEndpointId != nil {
		toSerialize["cloudProviderEndpointId"] = o.CloudProviderEndpointId
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.EndpointServiceName != nil {
		toSerialize["endpointServiceName"] = o.EndpointServiceName
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.ProviderName != nil {
		toSerialize["providerName"] = o.ProviderName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableServerlessAWSTenantEndpoint struct {
	value *ServerlessAWSTenantEndpoint
	isSet bool
}

func (v NullableServerlessAWSTenantEndpoint) Get() *ServerlessAWSTenantEndpoint {
	return v.value
}

func (v *NullableServerlessAWSTenantEndpoint) Set(val *ServerlessAWSTenantEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessAWSTenantEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessAWSTenantEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessAWSTenantEndpoint(val *ServerlessAWSTenantEndpoint) *NullableServerlessAWSTenantEndpoint {
	return &NullableServerlessAWSTenantEndpoint{value: val, isSet: true}
}

func (v NullableServerlessAWSTenantEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessAWSTenantEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

