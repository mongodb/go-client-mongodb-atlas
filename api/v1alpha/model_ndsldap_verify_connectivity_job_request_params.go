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

// checks if the NDSLDAPVerifyConnectivityJobRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NDSLDAPVerifyConnectivityJobRequestParams{}

// NDSLDAPVerifyConnectivityJobRequestParams Request information needed to verify an Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration. The response does not return the **bindPassword**.
type NDSLDAPVerifyConnectivityJobRequestParams struct {
	// Lightweight Directory Access Protocol (LDAP) query template that MongoDB Cloud applies to create an LDAP query to return the LDAP groups associated with the authenticated MongoDB user. MongoDB Cloud uses this parameter only for user authorization.  Use the `{USER}` placeholder in the Uniform Resource Locator (URL) to substitute the authenticated username. The query relates to the host specified with the hostname. Format this query per [RFC 4515](https://tools.ietf.org/search/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516).
	AuthzQueryTemplate *string `json:"authzQueryTemplate,omitempty"`
	// Password that MongoDB Cloud uses to authenticate the **bindUsername**.
	BindPassword string `json:"bindPassword"`
	// Full Distinguished Name (DN) of the Lightweight Directory Access Protocol (LDAP) user that MongoDB Cloud uses to connect to the LDAP host. LDAP distinguished names must be formatted according to RFC 2253.
	BindUsername string `json:"bindUsername"`
	// Certificate Authority (CA) certificate that MongoDB Cloud uses to verify the identity of the Lightweight Directory Access Protocol (LDAP) host. MongoDB Cloud allows self-signed certificates. To delete an assigned value, pass an empty string: `\"caCertificate\": \"\"`.
	CaCertificate *string `json:"caCertificate,omitempty"`
	// Human-readable label that identifies the hostname or Internet Protocol (IP) address of the Lightweight Directory Access Protocol (LDAP) host. This host must have access to the internet or have a Virtual Private Cloud (VPC) peering connection to your cluster.
	Hostname string `json:"hostname"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// IANA port to which the Lightweight Directory Access Protocol (LDAP) host listens for client connections.
	Port int32 `json:"port"`
}

// NewNDSLDAPVerifyConnectivityJobRequestParams instantiates a new NDSLDAPVerifyConnectivityJobRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNDSLDAPVerifyConnectivityJobRequestParams() *NDSLDAPVerifyConnectivityJobRequestParams {
	this := NDSLDAPVerifyConnectivityJobRequestParams{}
	var authzQueryTemplate string = "{USER}?memberOf?base"
	this.AuthzQueryTemplate = &authzQueryTemplate
	return &this
}

// NewNDSLDAPVerifyConnectivityJobRequestParamsWithDefaults instantiates a new NDSLDAPVerifyConnectivityJobRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNDSLDAPVerifyConnectivityJobRequestParamsWithDefaults() *NDSLDAPVerifyConnectivityJobRequestParams {
	this := NDSLDAPVerifyConnectivityJobRequestParams{}
	var authzQueryTemplate string = "{USER}?memberOf?base"
	this.AuthzQueryTemplate = &authzQueryTemplate
	var port int32 = 636
	this.Port = port
	return &this
}

// GetAuthzQueryTemplate returns the AuthzQueryTemplate field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetAuthzQueryTemplate() string {
	if o == nil || IsNil(o.AuthzQueryTemplate) {
		var ret string
		return ret
	}
	return *o.AuthzQueryTemplate
}

// GetAuthzQueryTemplateOk returns a tuple with the AuthzQueryTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetAuthzQueryTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.AuthzQueryTemplate) {
		return nil, false
	}
	return o.AuthzQueryTemplate, true
}

// HasAuthzQueryTemplate returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) HasAuthzQueryTemplate() bool {
	if o != nil && !IsNil(o.AuthzQueryTemplate) {
		return true
	}

	return false
}

// SetAuthzQueryTemplate gets a reference to the given string and assigns it to the AuthzQueryTemplate field.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetAuthzQueryTemplate(v string) {
	o.AuthzQueryTemplate = &v
}

// GetBindPassword returns the BindPassword field value
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetBindPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BindPassword
}

// GetBindPasswordOk returns a tuple with the BindPassword field value
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetBindPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BindPassword, true
}

// SetBindPassword sets field value
func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetBindPassword(v string) {
	o.BindPassword = v
}

// GetBindUsername returns the BindUsername field value
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetBindUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BindUsername
}

// GetBindUsernameOk returns a tuple with the BindUsername field value
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetBindUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BindUsername, true
}

// SetBindUsername sets field value
func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetBindUsername(v string) {
	o.BindUsername = v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetCaCertificate() string {
	if o == nil || IsNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetCaCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.CaCertificate) {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) HasCaCertificate() bool {
	if o != nil && !IsNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

// GetHostname returns the Hostname field value
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetHostname(v string) {
	o.Hostname = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetLinks(v []Link) {
	o.Links = v
}

// GetPort returns the Port field value
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetPort(v int32) {
	o.Port = v
}

func (o NDSLDAPVerifyConnectivityJobRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NDSLDAPVerifyConnectivityJobRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthzQueryTemplate) {
		toSerialize["authzQueryTemplate"] = o.AuthzQueryTemplate
	}
	toSerialize["bindPassword"] = o.BindPassword
	toSerialize["bindUsername"] = o.BindUsername
	if !IsNil(o.CaCertificate) {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	toSerialize["hostname"] = o.Hostname
	// skip: links is readOnly
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

type NullableNDSLDAPVerifyConnectivityJobRequestParams struct {
	value *NDSLDAPVerifyConnectivityJobRequestParams
	isSet bool
}

func (v NullableNDSLDAPVerifyConnectivityJobRequestParams) Get() *NDSLDAPVerifyConnectivityJobRequestParams {
	return v.value
}

func (v *NullableNDSLDAPVerifyConnectivityJobRequestParams) Set(val *NDSLDAPVerifyConnectivityJobRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSLDAPVerifyConnectivityJobRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSLDAPVerifyConnectivityJobRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSLDAPVerifyConnectivityJobRequestParams(val *NDSLDAPVerifyConnectivityJobRequestParams) *NullableNDSLDAPVerifyConnectivityJobRequestParams {
	return &NullableNDSLDAPVerifyConnectivityJobRequestParams{value: val, isSet: true}
}

func (v NullableNDSLDAPVerifyConnectivityJobRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSLDAPVerifyConnectivityJobRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

