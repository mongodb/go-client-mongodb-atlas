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

// checks if the AvailableDeploymentView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableDeploymentView{}

// AvailableDeploymentView Deployments that can be migrated to MongoDB Atlas.
type AvailableDeploymentView struct {
	// Version of MongoDB Agent that monitors/manages the cluster.
	AgentVersion *string `json:"agentVersion,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the cluster.
	ClusterId *string `json:"clusterId,omitempty"`
	// Size of this database on disk at the time of the request expressed in bytes.
	DbSizeBytes *int64 `json:"dbSizeBytes,omitempty"`
	// Version of MongoDB [features](https://docs.mongodb.com/manual/reference/command/setFeatureCompatibilityVersion) that this cluster supports.
	FeatureCompatibilityVersion string `json:"featureCompatibilityVersion"`
	// Flag that indicates whether Automation manages this cluster.
	Managed bool `json:"managed"`
	// Version of MongoDB that this cluster runs.
	MongoDBVersion string `json:"mongoDBVersion"`
	// Human-readable label that identifies this cluster.
	Name string `json:"name"`
	// Size of the Oplog on disk at the time of the request expressed in MB.
	OplogSizeMB *int32 `json:"oplogSizeMB,omitempty"`
	// Flag that indicates whether someone configured this cluster as a sharded cluster.  - If `true`, this cluster serves as a sharded cluster. - If `false`, this cluster serves as a replica set.
	Sharded bool `json:"sharded"`
	// Number of shards that comprise this cluster.
	ShardsSize *int32 `json:"shardsSize,omitempty"`
	// Flag that indicates whether someone enabled TLS for this cluster.
	TlsEnabled bool `json:"tlsEnabled"`
}

// NewAvailableDeploymentView instantiates a new AvailableDeploymentView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableDeploymentView() *AvailableDeploymentView {
	this := AvailableDeploymentView{}
	return &this
}

// NewAvailableDeploymentViewWithDefaults instantiates a new AvailableDeploymentView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableDeploymentViewWithDefaults() *AvailableDeploymentView {
	this := AvailableDeploymentView{}
	return &this
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *AvailableDeploymentView) GetAgentVersion() string {
	if o == nil || IsNil(o.AgentVersion) {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetAgentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AgentVersion) {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *AvailableDeploymentView) HasAgentVersion() bool {
	if o != nil && !IsNil(o.AgentVersion) {
		return true
	}

	return false
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *AvailableDeploymentView) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *AvailableDeploymentView) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *AvailableDeploymentView) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *AvailableDeploymentView) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetDbSizeBytes returns the DbSizeBytes field value if set, zero value otherwise.
func (o *AvailableDeploymentView) GetDbSizeBytes() int64 {
	if o == nil || IsNil(o.DbSizeBytes) {
		var ret int64
		return ret
	}
	return *o.DbSizeBytes
}

// GetDbSizeBytesOk returns a tuple with the DbSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetDbSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DbSizeBytes) {
		return nil, false
	}
	return o.DbSizeBytes, true
}

// HasDbSizeBytes returns a boolean if a field has been set.
func (o *AvailableDeploymentView) HasDbSizeBytes() bool {
	if o != nil && !IsNil(o.DbSizeBytes) {
		return true
	}

	return false
}

// SetDbSizeBytes gets a reference to the given int64 and assigns it to the DbSizeBytes field.
func (o *AvailableDeploymentView) SetDbSizeBytes(v int64) {
	o.DbSizeBytes = &v
}

// GetFeatureCompatibilityVersion returns the FeatureCompatibilityVersion field value
func (o *AvailableDeploymentView) GetFeatureCompatibilityVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureCompatibilityVersion
}

// GetFeatureCompatibilityVersionOk returns a tuple with the FeatureCompatibilityVersion field value
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetFeatureCompatibilityVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureCompatibilityVersion, true
}

// SetFeatureCompatibilityVersion sets field value
func (o *AvailableDeploymentView) SetFeatureCompatibilityVersion(v string) {
	o.FeatureCompatibilityVersion = v
}

// GetManaged returns the Managed field value
func (o *AvailableDeploymentView) GetManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Managed
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetManagedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Managed, true
}

// SetManaged sets field value
func (o *AvailableDeploymentView) SetManaged(v bool) {
	o.Managed = v
}

// GetMongoDBVersion returns the MongoDBVersion field value
func (o *AvailableDeploymentView) GetMongoDBVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MongoDBVersion
}

// GetMongoDBVersionOk returns a tuple with the MongoDBVersion field value
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetMongoDBVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MongoDBVersion, true
}

// SetMongoDBVersion sets field value
func (o *AvailableDeploymentView) SetMongoDBVersion(v string) {
	o.MongoDBVersion = v
}

// GetName returns the Name field value
func (o *AvailableDeploymentView) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AvailableDeploymentView) SetName(v string) {
	o.Name = v
}

// GetOplogSizeMB returns the OplogSizeMB field value if set, zero value otherwise.
func (o *AvailableDeploymentView) GetOplogSizeMB() int32 {
	if o == nil || IsNil(o.OplogSizeMB) {
		var ret int32
		return ret
	}
	return *o.OplogSizeMB
}

// GetOplogSizeMBOk returns a tuple with the OplogSizeMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetOplogSizeMBOk() (*int32, bool) {
	if o == nil || IsNil(o.OplogSizeMB) {
		return nil, false
	}
	return o.OplogSizeMB, true
}

// HasOplogSizeMB returns a boolean if a field has been set.
func (o *AvailableDeploymentView) HasOplogSizeMB() bool {
	if o != nil && !IsNil(o.OplogSizeMB) {
		return true
	}

	return false
}

// SetOplogSizeMB gets a reference to the given int32 and assigns it to the OplogSizeMB field.
func (o *AvailableDeploymentView) SetOplogSizeMB(v int32) {
	o.OplogSizeMB = &v
}

// GetSharded returns the Sharded field value
func (o *AvailableDeploymentView) GetSharded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sharded
}

// GetShardedOk returns a tuple with the Sharded field value
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetShardedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sharded, true
}

// SetSharded sets field value
func (o *AvailableDeploymentView) SetSharded(v bool) {
	o.Sharded = v
}

// GetShardsSize returns the ShardsSize field value if set, zero value otherwise.
func (o *AvailableDeploymentView) GetShardsSize() int32 {
	if o == nil || IsNil(o.ShardsSize) {
		var ret int32
		return ret
	}
	return *o.ShardsSize
}

// GetShardsSizeOk returns a tuple with the ShardsSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetShardsSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ShardsSize) {
		return nil, false
	}
	return o.ShardsSize, true
}

// HasShardsSize returns a boolean if a field has been set.
func (o *AvailableDeploymentView) HasShardsSize() bool {
	if o != nil && !IsNil(o.ShardsSize) {
		return true
	}

	return false
}

// SetShardsSize gets a reference to the given int32 and assigns it to the ShardsSize field.
func (o *AvailableDeploymentView) SetShardsSize(v int32) {
	o.ShardsSize = &v
}

// GetTlsEnabled returns the TlsEnabled field value
func (o *AvailableDeploymentView) GetTlsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TlsEnabled
}

// GetTlsEnabledOk returns a tuple with the TlsEnabled field value
// and a boolean to check if the value has been set.
func (o *AvailableDeploymentView) GetTlsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TlsEnabled, true
}

// SetTlsEnabled sets field value
func (o *AvailableDeploymentView) SetTlsEnabled(v bool) {
	o.TlsEnabled = v
}

func (o AvailableDeploymentView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailableDeploymentView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: agentVersion is readOnly
	// skip: clusterId is readOnly
	// skip: dbSizeBytes is readOnly
	// skip: featureCompatibilityVersion is readOnly
	// skip: managed is readOnly
	// skip: mongoDBVersion is readOnly
	// skip: name is readOnly
	// skip: oplogSizeMB is readOnly
	// skip: sharded is readOnly
	// skip: shardsSize is readOnly
	// skip: tlsEnabled is readOnly
	return toSerialize, nil
}

type NullableAvailableDeploymentView struct {
	value *AvailableDeploymentView
	isSet bool
}

func (v NullableAvailableDeploymentView) Get() *AvailableDeploymentView {
	return v.value
}

func (v *NullableAvailableDeploymentView) Set(val *AvailableDeploymentView) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableDeploymentView) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableDeploymentView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableDeploymentView(val *AvailableDeploymentView) *NullableAvailableDeploymentView {
	return &NullableAvailableDeploymentView{value: val, isSet: true}
}

func (v NullableAvailableDeploymentView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableDeploymentView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

