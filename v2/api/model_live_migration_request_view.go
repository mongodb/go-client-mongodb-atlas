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

// LiveMigrationRequestView struct for LiveMigrationRequestView
type LiveMigrationRequestView struct {
	// Unique 24-hexadecimal digit string that identifies the migration request.
	Id *string `json:"_id,omitempty"`
	Destination Destination `json:"destination"`
	// Flag that indicates whether the migration process drops all collections from the destination cluster before the migration starts.
	DropEnabled bool `json:"dropEnabled"`
	// List of migration hosts used for this migration.
	MigrationHosts []string `json:"migrationHosts"`
	Source Source `json:"source"`
}

// NewLiveMigrationRequestView instantiates a new LiveMigrationRequestView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveMigrationRequestView() *LiveMigrationRequestView {
	this := LiveMigrationRequestView{}
	return &this
}

// NewLiveMigrationRequestViewWithDefaults instantiates a new LiveMigrationRequestView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveMigrationRequestViewWithDefaults() *LiveMigrationRequestView {
	this := LiveMigrationRequestView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LiveMigrationRequestView) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequestView) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LiveMigrationRequestView) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LiveMigrationRequestView) SetId(v string) {
	o.Id = &v
}

// GetDestination returns the Destination field value
func (o *LiveMigrationRequestView) GetDestination() Destination {
	if o == nil {
		var ret Destination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequestView) GetDestinationOk() (*Destination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *LiveMigrationRequestView) SetDestination(v Destination) {
	o.Destination = v
}

// GetDropEnabled returns the DropEnabled field value
func (o *LiveMigrationRequestView) GetDropEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DropEnabled
}

// GetDropEnabledOk returns a tuple with the DropEnabled field value
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequestView) GetDropEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DropEnabled, true
}

// SetDropEnabled sets field value
func (o *LiveMigrationRequestView) SetDropEnabled(v bool) {
	o.DropEnabled = v
}

// GetMigrationHosts returns the MigrationHosts field value
func (o *LiveMigrationRequestView) GetMigrationHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MigrationHosts
}

// GetMigrationHostsOk returns a tuple with the MigrationHosts field value
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequestView) GetMigrationHostsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MigrationHosts, true
}

// SetMigrationHosts sets field value
func (o *LiveMigrationRequestView) SetMigrationHosts(v []string) {
	o.MigrationHosts = v
}

// GetSource returns the Source field value
func (o *LiveMigrationRequestView) GetSource() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequestView) GetSourceOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *LiveMigrationRequestView) SetSource(v Source) {
	o.Source = v
}

func (o LiveMigrationRequestView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["dropEnabled"] = o.DropEnabled
	}
	if true {
		toSerialize["migrationHosts"] = o.MigrationHosts
	}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableLiveMigrationRequestView struct {
	value *LiveMigrationRequestView
	isSet bool
}

func (v NullableLiveMigrationRequestView) Get() *LiveMigrationRequestView {
	return v.value
}

func (v *NullableLiveMigrationRequestView) Set(val *LiveMigrationRequestView) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveMigrationRequestView) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveMigrationRequestView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveMigrationRequestView(val *LiveMigrationRequestView) *NullableLiveMigrationRequestView {
	return &NullableLiveMigrationRequestView{value: val, isSet: true}
}

func (v NullableLiveMigrationRequestView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveMigrationRequestView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

