/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"time"
)

// DiskBackupSnapshotSchedule struct for DiskBackupSnapshotSchedule
type DiskBackupSnapshotSchedule struct {
	// Flag that indicates whether MongoDB Cloud automatically exports cloud backup snapshots to the AWS bucket.
	AutoExportEnabled *bool `json:"autoExportEnabled,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the cluster with the snapshot you want to return.
	ClusterId *string `json:"clusterId,omitempty"`
	// Human-readable label that identifies the cluster with the snapshot you want to return.
	ClusterName *string `json:"clusterName,omitempty"`
	// List that contains a document for each copy setting item in the desired backup policy.
	CopySettings []DiskBackupCopySetting `json:"copySettings,omitempty"`
	// List that contains a document for each deleted copy setting whose backup copies you want to delete.
	DeleteCopiedBackups []ApiDeleteCopiedBackupsView `json:"deleteCopiedBackups,omitempty"`
	Export *AutoExportPolicyView `json:"export,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Date and time when MongoDB Cloud takes the next snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	NextSnapshot *time.Time `json:"nextSnapshot,omitempty"`
	// Rules set for this backup schedule.
	Policies []ApiPolicyView `json:"policies"`
	// Hour of day in Coordinated Universal Time (UTC) that represents when MongoDB Cloud takes the snapshot.
	ReferenceHourOfDay *int32 `json:"referenceHourOfDay,omitempty"`
	// Minute of the **referenceHourOfDay** that represents when MongoDB Cloud takes the snapshot.
	ReferenceMinuteOfHour *int32 `json:"referenceMinuteOfHour,omitempty"`
	// Number of previous days that you can restore back to with Continuous Cloud Backup accuracy. You must specify a positive, non-zero integer. This parameter applies to continuous cloud backups only.
	RestoreWindowDays *int32 `json:"restoreWindowDays,omitempty"`
	// Flag that indicates whether to apply the retention changes in the updated backup policy to snapshots that MongoDB Cloud took previously.
	UpdateSnapshots *bool `json:"updateSnapshots,omitempty"`
	// Flag that indicates whether to use organization and project names instead of organization and project UUIDs in the path to the metadata files that MongoDB Cloud uploads to your AWS bucket.
	UseOrgAndGroupNamesInExportPrefix *bool `json:"useOrgAndGroupNamesInExportPrefix,omitempty"`
}

// NewDiskBackupSnapshotSchedule instantiates a new DiskBackupSnapshotSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupSnapshotSchedule() *DiskBackupSnapshotSchedule {
	this := DiskBackupSnapshotSchedule{}
	return &this
}

// NewDiskBackupSnapshotScheduleWithDefaults instantiates a new DiskBackupSnapshotSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupSnapshotScheduleWithDefaults() *DiskBackupSnapshotSchedule {
	this := DiskBackupSnapshotSchedule{}
	return &this
}

// GetAutoExportEnabled returns the AutoExportEnabled field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetAutoExportEnabled() bool {
	if o == nil || o.AutoExportEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutoExportEnabled
}

// GetAutoExportEnabledOk returns a tuple with the AutoExportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetAutoExportEnabledOk() (*bool, bool) {
	if o == nil || o.AutoExportEnabled == nil {
		return nil, false
	}
	return o.AutoExportEnabled, true
}

// HasAutoExportEnabled returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasAutoExportEnabled() bool {
	if o != nil && o.AutoExportEnabled != nil {
		return true
	}

	return false
}

// SetAutoExportEnabled gets a reference to the given bool and assigns it to the AutoExportEnabled field.
func (o *DiskBackupSnapshotSchedule) SetAutoExportEnabled(v bool) {
	o.AutoExportEnabled = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *DiskBackupSnapshotSchedule) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *DiskBackupSnapshotSchedule) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCopySettings returns the CopySettings field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetCopySettings() []DiskBackupCopySetting {
	if o == nil || o.CopySettings == nil {
		var ret []DiskBackupCopySetting
		return ret
	}
	return o.CopySettings
}

// GetCopySettingsOk returns a tuple with the CopySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetCopySettingsOk() ([]DiskBackupCopySetting, bool) {
	if o == nil || o.CopySettings == nil {
		return nil, false
	}
	return o.CopySettings, true
}

// HasCopySettings returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasCopySettings() bool {
	if o != nil && o.CopySettings != nil {
		return true
	}

	return false
}

// SetCopySettings gets a reference to the given []DiskBackupCopySetting and assigns it to the CopySettings field.
func (o *DiskBackupSnapshotSchedule) SetCopySettings(v []DiskBackupCopySetting) {
	o.CopySettings = v
}

// GetDeleteCopiedBackups returns the DeleteCopiedBackups field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetDeleteCopiedBackups() []ApiDeleteCopiedBackupsView {
	if o == nil || o.DeleteCopiedBackups == nil {
		var ret []ApiDeleteCopiedBackupsView
		return ret
	}
	return o.DeleteCopiedBackups
}

// GetDeleteCopiedBackupsOk returns a tuple with the DeleteCopiedBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetDeleteCopiedBackupsOk() ([]ApiDeleteCopiedBackupsView, bool) {
	if o == nil || o.DeleteCopiedBackups == nil {
		return nil, false
	}
	return o.DeleteCopiedBackups, true
}

// HasDeleteCopiedBackups returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasDeleteCopiedBackups() bool {
	if o != nil && o.DeleteCopiedBackups != nil {
		return true
	}

	return false
}

// SetDeleteCopiedBackups gets a reference to the given []ApiDeleteCopiedBackupsView and assigns it to the DeleteCopiedBackups field.
func (o *DiskBackupSnapshotSchedule) SetDeleteCopiedBackups(v []ApiDeleteCopiedBackupsView) {
	o.DeleteCopiedBackups = v
}

// GetExport returns the Export field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetExport() AutoExportPolicyView {
	if o == nil || o.Export == nil {
		var ret AutoExportPolicyView
		return ret
	}
	return *o.Export
}

// GetExportOk returns a tuple with the Export field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetExportOk() (*AutoExportPolicyView, bool) {
	if o == nil || o.Export == nil {
		return nil, false
	}
	return o.Export, true
}

// HasExport returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasExport() bool {
	if o != nil && o.Export != nil {
		return true
	}

	return false
}

// SetExport gets a reference to the given AutoExportPolicyView and assigns it to the Export field.
func (o *DiskBackupSnapshotSchedule) SetExport(v AutoExportPolicyView) {
	o.Export = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetLinks() []Link {
	if o == nil || o.Links == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetLinksOk() ([]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DiskBackupSnapshotSchedule) SetLinks(v []Link) {
	o.Links = v
}

// GetNextSnapshot returns the NextSnapshot field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetNextSnapshot() time.Time {
	if o == nil || o.NextSnapshot == nil {
		var ret time.Time
		return ret
	}
	return *o.NextSnapshot
}

// GetNextSnapshotOk returns a tuple with the NextSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetNextSnapshotOk() (*time.Time, bool) {
	if o == nil || o.NextSnapshot == nil {
		return nil, false
	}
	return o.NextSnapshot, true
}

// HasNextSnapshot returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasNextSnapshot() bool {
	if o != nil && o.NextSnapshot != nil {
		return true
	}

	return false
}

// SetNextSnapshot gets a reference to the given time.Time and assigns it to the NextSnapshot field.
func (o *DiskBackupSnapshotSchedule) SetNextSnapshot(v time.Time) {
	o.NextSnapshot = &v
}

// GetPolicies returns the Policies field value
func (o *DiskBackupSnapshotSchedule) GetPolicies() []ApiPolicyView {
	if o == nil {
		var ret []ApiPolicyView
		return ret
	}

	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetPoliciesOk() ([]ApiPolicyView, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policies, true
}

// SetPolicies sets field value
func (o *DiskBackupSnapshotSchedule) SetPolicies(v []ApiPolicyView) {
	o.Policies = v
}

// GetReferenceHourOfDay returns the ReferenceHourOfDay field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetReferenceHourOfDay() int32 {
	if o == nil || o.ReferenceHourOfDay == nil {
		var ret int32
		return ret
	}
	return *o.ReferenceHourOfDay
}

// GetReferenceHourOfDayOk returns a tuple with the ReferenceHourOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetReferenceHourOfDayOk() (*int32, bool) {
	if o == nil || o.ReferenceHourOfDay == nil {
		return nil, false
	}
	return o.ReferenceHourOfDay, true
}

// HasReferenceHourOfDay returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasReferenceHourOfDay() bool {
	if o != nil && o.ReferenceHourOfDay != nil {
		return true
	}

	return false
}

// SetReferenceHourOfDay gets a reference to the given int32 and assigns it to the ReferenceHourOfDay field.
func (o *DiskBackupSnapshotSchedule) SetReferenceHourOfDay(v int32) {
	o.ReferenceHourOfDay = &v
}

// GetReferenceMinuteOfHour returns the ReferenceMinuteOfHour field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetReferenceMinuteOfHour() int32 {
	if o == nil || o.ReferenceMinuteOfHour == nil {
		var ret int32
		return ret
	}
	return *o.ReferenceMinuteOfHour
}

// GetReferenceMinuteOfHourOk returns a tuple with the ReferenceMinuteOfHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetReferenceMinuteOfHourOk() (*int32, bool) {
	if o == nil || o.ReferenceMinuteOfHour == nil {
		return nil, false
	}
	return o.ReferenceMinuteOfHour, true
}

// HasReferenceMinuteOfHour returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasReferenceMinuteOfHour() bool {
	if o != nil && o.ReferenceMinuteOfHour != nil {
		return true
	}

	return false
}

// SetReferenceMinuteOfHour gets a reference to the given int32 and assigns it to the ReferenceMinuteOfHour field.
func (o *DiskBackupSnapshotSchedule) SetReferenceMinuteOfHour(v int32) {
	o.ReferenceMinuteOfHour = &v
}

// GetRestoreWindowDays returns the RestoreWindowDays field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetRestoreWindowDays() int32 {
	if o == nil || o.RestoreWindowDays == nil {
		var ret int32
		return ret
	}
	return *o.RestoreWindowDays
}

// GetRestoreWindowDaysOk returns a tuple with the RestoreWindowDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetRestoreWindowDaysOk() (*int32, bool) {
	if o == nil || o.RestoreWindowDays == nil {
		return nil, false
	}
	return o.RestoreWindowDays, true
}

// HasRestoreWindowDays returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasRestoreWindowDays() bool {
	if o != nil && o.RestoreWindowDays != nil {
		return true
	}

	return false
}

// SetRestoreWindowDays gets a reference to the given int32 and assigns it to the RestoreWindowDays field.
func (o *DiskBackupSnapshotSchedule) SetRestoreWindowDays(v int32) {
	o.RestoreWindowDays = &v
}

// GetUpdateSnapshots returns the UpdateSnapshots field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetUpdateSnapshots() bool {
	if o == nil || o.UpdateSnapshots == nil {
		var ret bool
		return ret
	}
	return *o.UpdateSnapshots
}

// GetUpdateSnapshotsOk returns a tuple with the UpdateSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetUpdateSnapshotsOk() (*bool, bool) {
	if o == nil || o.UpdateSnapshots == nil {
		return nil, false
	}
	return o.UpdateSnapshots, true
}

// HasUpdateSnapshots returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasUpdateSnapshots() bool {
	if o != nil && o.UpdateSnapshots != nil {
		return true
	}

	return false
}

// SetUpdateSnapshots gets a reference to the given bool and assigns it to the UpdateSnapshots field.
func (o *DiskBackupSnapshotSchedule) SetUpdateSnapshots(v bool) {
	o.UpdateSnapshots = &v
}

// GetUseOrgAndGroupNamesInExportPrefix returns the UseOrgAndGroupNamesInExportPrefix field value if set, zero value otherwise.
func (o *DiskBackupSnapshotSchedule) GetUseOrgAndGroupNamesInExportPrefix() bool {
	if o == nil || o.UseOrgAndGroupNamesInExportPrefix == nil {
		var ret bool
		return ret
	}
	return *o.UseOrgAndGroupNamesInExportPrefix
}

// GetUseOrgAndGroupNamesInExportPrefixOk returns a tuple with the UseOrgAndGroupNamesInExportPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetUseOrgAndGroupNamesInExportPrefixOk() (*bool, bool) {
	if o == nil || o.UseOrgAndGroupNamesInExportPrefix == nil {
		return nil, false
	}
	return o.UseOrgAndGroupNamesInExportPrefix, true
}

// HasUseOrgAndGroupNamesInExportPrefix returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasUseOrgAndGroupNamesInExportPrefix() bool {
	if o != nil && o.UseOrgAndGroupNamesInExportPrefix != nil {
		return true
	}

	return false
}

// SetUseOrgAndGroupNamesInExportPrefix gets a reference to the given bool and assigns it to the UseOrgAndGroupNamesInExportPrefix field.
func (o *DiskBackupSnapshotSchedule) SetUseOrgAndGroupNamesInExportPrefix(v bool) {
	o.UseOrgAndGroupNamesInExportPrefix = &v
}

func (o DiskBackupSnapshotSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoExportEnabled != nil {
		toSerialize["autoExportEnabled"] = o.AutoExportEnabled
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.CopySettings != nil {
		toSerialize["copySettings"] = o.CopySettings
	}
	if o.DeleteCopiedBackups != nil {
		toSerialize["deleteCopiedBackups"] = o.DeleteCopiedBackups
	}
	if o.Export != nil {
		toSerialize["export"] = o.Export
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.NextSnapshot != nil {
		toSerialize["nextSnapshot"] = o.NextSnapshot
	}
	if true {
		toSerialize["policies"] = o.Policies
	}
	if o.ReferenceHourOfDay != nil {
		toSerialize["referenceHourOfDay"] = o.ReferenceHourOfDay
	}
	if o.ReferenceMinuteOfHour != nil {
		toSerialize["referenceMinuteOfHour"] = o.ReferenceMinuteOfHour
	}
	if o.RestoreWindowDays != nil {
		toSerialize["restoreWindowDays"] = o.RestoreWindowDays
	}
	if o.UpdateSnapshots != nil {
		toSerialize["updateSnapshots"] = o.UpdateSnapshots
	}
	if o.UseOrgAndGroupNamesInExportPrefix != nil {
		toSerialize["useOrgAndGroupNamesInExportPrefix"] = o.UseOrgAndGroupNamesInExportPrefix
	}
	return json.Marshal(toSerialize)
}

type NullableDiskBackupSnapshotSchedule struct {
	value *DiskBackupSnapshotSchedule
	isSet bool
}

func (v NullableDiskBackupSnapshotSchedule) Get() *DiskBackupSnapshotSchedule {
	return v.value
}

func (v *NullableDiskBackupSnapshotSchedule) Set(val *DiskBackupSnapshotSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskBackupSnapshotSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskBackupSnapshotSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskBackupSnapshotSchedule(val *DiskBackupSnapshotSchedule) *NullableDiskBackupSnapshotSchedule {
	return &NullableDiskBackupSnapshotSchedule{value: val, isSet: true}
}

func (v NullableDiskBackupSnapshotSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskBackupSnapshotSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

