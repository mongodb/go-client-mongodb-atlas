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

// ClusterView Settings that describe the clusters in each project that the API key is authorized to view.
type ClusterView struct {
	// Whole number that indicates the quantity of alerts open on the cluster.
	AlertCount *int32 `json:"alertCount,omitempty"`
	// Flag that indicates whether authentication is required to access the nodes in this cluster.
	AuthEnabled *bool `json:"authEnabled,omitempty"`
	// Term that expresses how many nodes of the cluster can be accessed when MongoDB Cloud receives this request. This parameter returns `available` when all nodes are accessible, `warning` only when some nodes in the cluster can be accessed, `unavailable` when the cluster can't be accessed, or `dead` when the cluster has been deactivated.
	Availability *string `json:"availability,omitempty"`
	// Flag that indicates whether the cluster can perform backups. If set to `true`, the cluster can perform backups. You must set this value to `true` for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to `false`, the cluster doesn't use MongoDB Cloud backups.
	BackupEnabled *bool `json:"backupEnabled,omitempty"`
	// Unique 24-hexadecimal character string that identifies the cluster.
	ClusterId *string `json:"clusterId,omitempty"`
	// Total size of the data stored on each node in the cluster. The resource expresses this value in bytes.
	DataSizeBytes *int64 `json:"dataSizeBytes,omitempty"`
	// Human-readable label that identifies the cluster.
	Name *string `json:"name,omitempty"`
	// Whole number that indicates the quantity of nodes that comprise the cluster.
	NodeCount *int32 `json:"nodeCount,omitempty"`
	// Flag that indicates whether TLS authentication is required to access the nodes in this cluster.
	SslEnabled *bool `json:"sslEnabled,omitempty"`
	// Human-readable label that indicates the cluster type.
	Type *string `json:"type,omitempty"`
	// List that contains the versions of MongoDB that each node in the cluster runs.
	Versions []string `json:"versions,omitempty"`
}

// NewClusterView instantiates a new ClusterView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterView() *ClusterView {
	this := ClusterView{}
	return &this
}

// NewClusterViewWithDefaults instantiates a new ClusterView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterViewWithDefaults() *ClusterView {
	this := ClusterView{}
	return &this
}

// GetAlertCount returns the AlertCount field value if set, zero value otherwise.
func (o *ClusterView) GetAlertCount() int32 {
	if o == nil || o.AlertCount == nil {
		var ret int32
		return ret
	}
	return *o.AlertCount
}

// GetAlertCountOk returns a tuple with the AlertCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetAlertCountOk() (*int32, bool) {
	if o == nil || o.AlertCount == nil {
		return nil, false
	}
	return o.AlertCount, true
}

// HasAlertCount returns a boolean if a field has been set.
func (o *ClusterView) HasAlertCount() bool {
	if o != nil && o.AlertCount != nil {
		return true
	}

	return false
}

// SetAlertCount gets a reference to the given int32 and assigns it to the AlertCount field.
func (o *ClusterView) SetAlertCount(v int32) {
	o.AlertCount = &v
}

// GetAuthEnabled returns the AuthEnabled field value if set, zero value otherwise.
func (o *ClusterView) GetAuthEnabled() bool {
	if o == nil || o.AuthEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthEnabled
}

// GetAuthEnabledOk returns a tuple with the AuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetAuthEnabledOk() (*bool, bool) {
	if o == nil || o.AuthEnabled == nil {
		return nil, false
	}
	return o.AuthEnabled, true
}

// HasAuthEnabled returns a boolean if a field has been set.
func (o *ClusterView) HasAuthEnabled() bool {
	if o != nil && o.AuthEnabled != nil {
		return true
	}

	return false
}

// SetAuthEnabled gets a reference to the given bool and assigns it to the AuthEnabled field.
func (o *ClusterView) SetAuthEnabled(v bool) {
	o.AuthEnabled = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *ClusterView) GetAvailability() string {
	if o == nil || o.Availability == nil {
		var ret string
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetAvailabilityOk() (*string, bool) {
	if o == nil || o.Availability == nil {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *ClusterView) HasAvailability() bool {
	if o != nil && o.Availability != nil {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given string and assigns it to the Availability field.
func (o *ClusterView) SetAvailability(v string) {
	o.Availability = &v
}

// GetBackupEnabled returns the BackupEnabled field value if set, zero value otherwise.
func (o *ClusterView) GetBackupEnabled() bool {
	if o == nil || o.BackupEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BackupEnabled
}

// GetBackupEnabledOk returns a tuple with the BackupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetBackupEnabledOk() (*bool, bool) {
	if o == nil || o.BackupEnabled == nil {
		return nil, false
	}
	return o.BackupEnabled, true
}

// HasBackupEnabled returns a boolean if a field has been set.
func (o *ClusterView) HasBackupEnabled() bool {
	if o != nil && o.BackupEnabled != nil {
		return true
	}

	return false
}

// SetBackupEnabled gets a reference to the given bool and assigns it to the BackupEnabled field.
func (o *ClusterView) SetBackupEnabled(v bool) {
	o.BackupEnabled = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *ClusterView) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *ClusterView) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *ClusterView) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetDataSizeBytes returns the DataSizeBytes field value if set, zero value otherwise.
func (o *ClusterView) GetDataSizeBytes() int64 {
	if o == nil || o.DataSizeBytes == nil {
		var ret int64
		return ret
	}
	return *o.DataSizeBytes
}

// GetDataSizeBytesOk returns a tuple with the DataSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetDataSizeBytesOk() (*int64, bool) {
	if o == nil || o.DataSizeBytes == nil {
		return nil, false
	}
	return o.DataSizeBytes, true
}

// HasDataSizeBytes returns a boolean if a field has been set.
func (o *ClusterView) HasDataSizeBytes() bool {
	if o != nil && o.DataSizeBytes != nil {
		return true
	}

	return false
}

// SetDataSizeBytes gets a reference to the given int64 and assigns it to the DataSizeBytes field.
func (o *ClusterView) SetDataSizeBytes(v int64) {
	o.DataSizeBytes = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterView) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterView) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterView) SetName(v string) {
	o.Name = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *ClusterView) GetNodeCount() int32 {
	if o == nil || o.NodeCount == nil {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetNodeCountOk() (*int32, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *ClusterView) HasNodeCount() bool {
	if o != nil && o.NodeCount != nil {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *ClusterView) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetSslEnabled returns the SslEnabled field value if set, zero value otherwise.
func (o *ClusterView) GetSslEnabled() bool {
	if o == nil || o.SslEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SslEnabled
}

// GetSslEnabledOk returns a tuple with the SslEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetSslEnabledOk() (*bool, bool) {
	if o == nil || o.SslEnabled == nil {
		return nil, false
	}
	return o.SslEnabled, true
}

// HasSslEnabled returns a boolean if a field has been set.
func (o *ClusterView) HasSslEnabled() bool {
	if o != nil && o.SslEnabled != nil {
		return true
	}

	return false
}

// SetSslEnabled gets a reference to the given bool and assigns it to the SslEnabled field.
func (o *ClusterView) SetSslEnabled(v bool) {
	o.SslEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterView) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterView) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterView) SetType(v string) {
	o.Type = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ClusterView) GetVersions() []string {
	if o == nil || o.Versions == nil {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterView) GetVersionsOk() ([]string, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ClusterView) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *ClusterView) SetVersions(v []string) {
	o.Versions = v
}

func (o ClusterView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlertCount != nil {
		toSerialize["alertCount"] = o.AlertCount
	}
	if o.AuthEnabled != nil {
		toSerialize["authEnabled"] = o.AuthEnabled
	}
	if o.Availability != nil {
		toSerialize["availability"] = o.Availability
	}
	if o.BackupEnabled != nil {
		toSerialize["backupEnabled"] = o.BackupEnabled
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	if o.DataSizeBytes != nil {
		toSerialize["dataSizeBytes"] = o.DataSizeBytes
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NodeCount != nil {
		toSerialize["nodeCount"] = o.NodeCount
	}
	if o.SslEnabled != nil {
		toSerialize["sslEnabled"] = o.SslEnabled
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	return json.Marshal(toSerialize)
}

type NullableClusterView struct {
	value *ClusterView
	isSet bool
}

func (v NullableClusterView) Get() *ClusterView {
	return v.value
}

func (v *NullableClusterView) Set(val *ClusterView) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterView) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterView(val *ClusterView) *NullableClusterView {
	return &NullableClusterView{value: val, isSet: true}
}

func (v NullableClusterView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

