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

// checks if the MeasurementsGeneralViewAtlas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeasurementsGeneralViewAtlas{}

// MeasurementsGeneralViewAtlas struct for MeasurementsGeneralViewAtlas
type MeasurementsGeneralViewAtlas struct {
	// Human-readable label that identifies the database that the specified MongoDB process serves.
	DatabaseName *string `json:"databaseName,omitempty"`
	// Date and time that specifies when to stop retrieving measurements. If you set **end**, you must set **start**. You can't set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	End *time.Time `json:"end,omitempty"`
	// Duration that specifies the interval between measurement data points. The parameter expresses its value in ISO 8601 timestamp format in UTC. If you set this parameter, you must set either **period** or **start** and **end**.
	Granularity *string `json:"granularity,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project. The project contains MongoDB processes that you want to return. The MongoDB process can be either the `mongod` or `mongos`.
	GroupId *string `json:"groupId,omitempty"`
	// Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	HostId *string `json:"hostId,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []LinkAtlas `json:"links,omitempty"`
	// List that contains measurements and their data points.
	Measurements []MeasurementViewAtlas `json:"measurements,omitempty"`
	// Human-readable label of the disk or partition to which the measurements apply.
	PartitionName *string `json:"partitionName,omitempty"`
	// Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	ProcessId *string `json:"processId,omitempty"`
	// Date and time that specifies when to start retrieving measurements. If you set **start**, you must set **end**. You can't set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Start *time.Time `json:"start,omitempty"`
}

// NewMeasurementsGeneralViewAtlas instantiates a new MeasurementsGeneralViewAtlas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurementsGeneralViewAtlas() *MeasurementsGeneralViewAtlas {
	this := MeasurementsGeneralViewAtlas{}
	return &this
}

// NewMeasurementsGeneralViewAtlasWithDefaults instantiates a new MeasurementsGeneralViewAtlas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementsGeneralViewAtlasWithDefaults() *MeasurementsGeneralViewAtlas {
	this := MeasurementsGeneralViewAtlas{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *MeasurementsGeneralViewAtlas) GetDatabaseName() string {
	if o == nil || IsNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsGeneralViewAtlas) GetDatabaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseName) {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *MeasurementsGeneralViewAtlas) HasDatabaseName() bool {
	if o != nil && !IsNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *MeasurementsGeneralViewAtlas) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *MeasurementsGeneralViewAtlas) GetEnd() time.Time {
	if o == nil || IsNil(o.End) {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsGeneralViewAtlas) GetEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *MeasurementsGeneralViewAtlas) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *MeasurementsGeneralViewAtlas) SetEnd(v time.Time) {
	o.End = &v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *MeasurementsGeneralViewAtlas) GetGranularity() string {
	if o == nil || IsNil(o.Granularity) {
		var ret string
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsGeneralViewAtlas) GetGranularityOk() (*string, bool) {
	if o == nil || IsNil(o.Granularity) {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *MeasurementsGeneralViewAtlas) HasGranularity() bool {
	if o != nil && !IsNil(o.Granularity) {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given string and assigns it to the Granularity field.
func (o *MeasurementsGeneralViewAtlas) SetGranularity(v string) {
	o.Granularity = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *MeasurementsGeneralViewAtlas) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsGeneralViewAtlas) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *MeasurementsGeneralViewAtlas) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *MeasurementsGeneralViewAtlas) SetGroupId(v string) {
	o.GroupId = &v
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *MeasurementsGeneralViewAtlas) GetHostId() string {
	if o == nil || IsNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsGeneralViewAtlas) GetHostIdOk() (*string, bool) {
	if o == nil || IsNil(o.HostId) {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *MeasurementsGeneralViewAtlas) HasHostId() bool {
	if o != nil && !IsNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *MeasurementsGeneralViewAtlas) SetHostId(v string) {
	o.HostId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MeasurementsGeneralViewAtlas) GetLinks() []LinkAtlas {
	if o == nil || IsNil(o.Links) {
		var ret []LinkAtlas
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsGeneralViewAtlas) GetLinksOk() ([]LinkAtlas, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MeasurementsGeneralViewAtlas) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkAtlas and assigns it to the Links field.
func (o *MeasurementsGeneralViewAtlas) SetLinks(v []LinkAtlas) {
	o.Links = v
}

// GetMeasurements returns the Measurements field value if set, zero value otherwise.
func (o *MeasurementsGeneralViewAtlas) GetMeasurements() []MeasurementViewAtlas {
	if o == nil || IsNil(o.Measurements) {
		var ret []MeasurementViewAtlas
		return ret
	}
	return o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsGeneralViewAtlas) GetMeasurementsOk() ([]MeasurementViewAtlas, bool) {
	if o == nil || IsNil(o.Measurements) {
		return nil, false
	}
	return o.Measurements, true
}

// HasMeasurements returns a boolean if a field has been set.
func (o *MeasurementsGeneralViewAtlas) HasMeasurements() bool {
	if o != nil && !IsNil(o.Measurements) {
		return true
	}

	return false
}

// SetMeasurements gets a reference to the given []MeasurementViewAtlas and assigns it to the Measurements field.
func (o *MeasurementsGeneralViewAtlas) SetMeasurements(v []MeasurementViewAtlas) {
	o.Measurements = v
}

// GetPartitionName returns the PartitionName field value if set, zero value otherwise.
func (o *MeasurementsGeneralViewAtlas) GetPartitionName() string {
	if o == nil || IsNil(o.PartitionName) {
		var ret string
		return ret
	}
	return *o.PartitionName
}

// GetPartitionNameOk returns a tuple with the PartitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsGeneralViewAtlas) GetPartitionNameOk() (*string, bool) {
	if o == nil || IsNil(o.PartitionName) {
		return nil, false
	}
	return o.PartitionName, true
}

// HasPartitionName returns a boolean if a field has been set.
func (o *MeasurementsGeneralViewAtlas) HasPartitionName() bool {
	if o != nil && !IsNil(o.PartitionName) {
		return true
	}

	return false
}

// SetPartitionName gets a reference to the given string and assigns it to the PartitionName field.
func (o *MeasurementsGeneralViewAtlas) SetPartitionName(v string) {
	o.PartitionName = &v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *MeasurementsGeneralViewAtlas) GetProcessId() string {
	if o == nil || IsNil(o.ProcessId) {
		var ret string
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsGeneralViewAtlas) GetProcessIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessId) {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *MeasurementsGeneralViewAtlas) HasProcessId() bool {
	if o != nil && !IsNil(o.ProcessId) {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given string and assigns it to the ProcessId field.
func (o *MeasurementsGeneralViewAtlas) SetProcessId(v string) {
	o.ProcessId = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *MeasurementsGeneralViewAtlas) GetStart() time.Time {
	if o == nil || IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementsGeneralViewAtlas) GetStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *MeasurementsGeneralViewAtlas) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *MeasurementsGeneralViewAtlas) SetStart(v time.Time) {
	o.Start = &v
}

func (o MeasurementsGeneralViewAtlas) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeasurementsGeneralViewAtlas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: databaseName is readOnly
	// skip: end is readOnly
	// skip: granularity is readOnly
	// skip: groupId is readOnly
	// skip: hostId is readOnly
	// skip: links is readOnly
	// skip: measurements is readOnly
	// skip: partitionName is readOnly
	// skip: processId is readOnly
	// skip: start is readOnly
	return toSerialize, nil
}

type NullableMeasurementsGeneralViewAtlas struct {
	value *MeasurementsGeneralViewAtlas
	isSet bool
}

func (v NullableMeasurementsGeneralViewAtlas) Get() *MeasurementsGeneralViewAtlas {
	return v.value
}

func (v *NullableMeasurementsGeneralViewAtlas) Set(val *MeasurementsGeneralViewAtlas) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementsGeneralViewAtlas) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementsGeneralViewAtlas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementsGeneralViewAtlas(val *MeasurementsGeneralViewAtlas) *NullableMeasurementsGeneralViewAtlas {
	return &NullableMeasurementsGeneralViewAtlas{value: val, isSet: true}
}

func (v NullableMeasurementsGeneralViewAtlas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementsGeneralViewAtlas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

