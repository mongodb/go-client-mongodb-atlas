/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"fmt"
)

// HostEventTypeForGroup Unique identifier of event type.
type HostEventTypeForGroup string

// List of HostEventTypeForGroup
const (
	HOSTEVENTTYPEFORGROUP_ADD_HOST_AUDIT HostEventTypeForGroup = "ADD_HOST_AUDIT"
	HOSTEVENTTYPEFORGROUP_DELETE_HOST_AUDIT HostEventTypeForGroup = "DELETE_HOST_AUDIT"
	HOSTEVENTTYPEFORGROUP_UNDELETE_HOST_AUDIT HostEventTypeForGroup = "UNDELETE_HOST_AUDIT"
	HOSTEVENTTYPEFORGROUP_ADD_HOST_TO_REPLICA_SET_AUDIT HostEventTypeForGroup = "ADD_HOST_TO_REPLICA_SET_AUDIT"
	HOSTEVENTTYPEFORGROUP_REMOVE_HOST_FROM_REPLICA_SET_AUDIT HostEventTypeForGroup = "REMOVE_HOST_FROM_REPLICA_SET_AUDIT"
	HOSTEVENTTYPEFORGROUP_HIDE_HOST_AUDIT HostEventTypeForGroup = "HIDE_HOST_AUDIT"
	HOSTEVENTTYPEFORGROUP_HIDE_AND_DISABLE_HOST_AUDIT HostEventTypeForGroup = "HIDE_AND_DISABLE_HOST_AUDIT"
	HOSTEVENTTYPEFORGROUP_DISABLE_HOST_AUDIT HostEventTypeForGroup = "DISABLE_HOST_AUDIT"
	HOSTEVENTTYPEFORGROUP_PAUSE_HOST_AUDIT HostEventTypeForGroup = "PAUSE_HOST_AUDIT"
	HOSTEVENTTYPEFORGROUP_RESUME_HOST_AUDIT HostEventTypeForGroup = "RESUME_HOST_AUDIT"
	HOSTEVENTTYPEFORGROUP_DB_PROFILER_ENABLE_AUDIT HostEventTypeForGroup = "DB_PROFILER_ENABLE_AUDIT"
	HOSTEVENTTYPEFORGROUP_DB_PROFILER_DISABLE_AUDIT HostEventTypeForGroup = "DB_PROFILER_DISABLE_AUDIT"
	HOSTEVENTTYPEFORGROUP_HOST_IP_CHANGED_AUDIT HostEventTypeForGroup = "HOST_IP_CHANGED_AUDIT"
	HOSTEVENTTYPEFORGROUP_AUTO_CREATED_INDEX_AUDIT HostEventTypeForGroup = "AUTO_CREATED_INDEX_AUDIT"
	HOSTEVENTTYPEFORGROUP_ATTEMPT_KILLOP_AUDIT HostEventTypeForGroup = "ATTEMPT_KILLOP_AUDIT"
	HOSTEVENTTYPEFORGROUP_ATTEMPT_KILLSESSION_AUDIT HostEventTypeForGroup = "ATTEMPT_KILLSESSION_AUDIT"
	HOSTEVENTTYPEFORGROUP_HOST_UP HostEventTypeForGroup = "HOST_UP"
	HOSTEVENTTYPEFORGROUP_HOST_DOWN HostEventTypeForGroup = "HOST_DOWN"
	HOSTEVENTTYPEFORGROUP_HOST_ROLLBACK HostEventTypeForGroup = "HOST_ROLLBACK"
	HOSTEVENTTYPEFORGROUP_HOST_RECOVERED HostEventTypeForGroup = "HOST_RECOVERED"
	HOSTEVENTTYPEFORGROUP_HOST_RECOVERING HostEventTypeForGroup = "HOST_RECOVERING"
	HOSTEVENTTYPEFORGROUP_VERSION_CURRENT HostEventTypeForGroup = "VERSION_CURRENT"
	HOSTEVENTTYPEFORGROUP_VERSION_BEHIND HostEventTypeForGroup = "VERSION_BEHIND"
	HOSTEVENTTYPEFORGROUP_VALUE_NO_LONGER_ANOMALOUS HostEventTypeForGroup = "VALUE_NO_LONGER_ANOMALOUS"
	HOSTEVENTTYPEFORGROUP_HOST_LOCKED_DOWN HostEventTypeForGroup = "HOST_LOCKED_DOWN"
	HOSTEVENTTYPEFORGROUP_HOST_EXPOSED HostEventTypeForGroup = "HOST_EXPOSED"
	HOSTEVENTTYPEFORGROUP_VERSION_CHANGED HostEventTypeForGroup = "VERSION_CHANGED"
	HOSTEVENTTYPEFORGROUP_HOST_SSL_CERTIFICATE_CURRENT HostEventTypeForGroup = "HOST_SSL_CERTIFICATE_CURRENT"
	HOSTEVENTTYPEFORGROUP_HOST_SSL_CERTIFICATE_STALE HostEventTypeForGroup = "HOST_SSL_CERTIFICATE_STALE"
	HOSTEVENTTYPEFORGROUP_HOST_HAS_INDEX_SUGGESTIONS HostEventTypeForGroup = "HOST_HAS_INDEX_SUGGESTIONS"
	HOSTEVENTTYPEFORGROUP_HOST_SECURITY_CHECKUP_MET HostEventTypeForGroup = "HOST_SECURITY_CHECKUP_MET"
	HOSTEVENTTYPEFORGROUP_HOST_SECURITY_CHECKUP_NOT_MET HostEventTypeForGroup = "HOST_SECURITY_CHECKUP_NOT_MET"
	HOSTEVENTTYPEFORGROUP_HOST_MONGOT_RECOVERED_OOM HostEventTypeForGroup = "HOST_MONGOT_RECOVERED_OOM"
	HOSTEVENTTYPEFORGROUP_HOST_MONGOT_CRASHING_OOM HostEventTypeForGroup = "HOST_MONGOT_CRASHING_OOM"
	HOSTEVENTTYPEFORGROUP_HOST_DISK_SPACE_SUFFICIENT_FOR_SEARCH_INDEX_REBUILD HostEventTypeForGroup = "HOST_DISK_SPACE_SUFFICIENT_FOR_SEARCH_INDEX_REBUILD"
	HOSTEVENTTYPEFORGROUP_HOST_DISK_SPACE_INSUFFICIENT_FOR_SEARCH_INDEX_REBUILD HostEventTypeForGroup = "HOST_DISK_SPACE_INSUFFICIENT_FOR_SEARCH_INDEX_REBUILD"
	HOSTEVENTTYPEFORGROUP_NEW_HOST HostEventTypeForGroup = "NEW_HOST"
	HOSTEVENTTYPEFORGROUP_HOST_RESTARTED HostEventTypeForGroup = "HOST_RESTARTED"
	HOSTEVENTTYPEFORGROUP_HOST_UPGRADED HostEventTypeForGroup = "HOST_UPGRADED"
	HOSTEVENTTYPEFORGROUP_HOST_DOWNGRADED HostEventTypeForGroup = "HOST_DOWNGRADED"
	HOSTEVENTTYPEFORGROUP_HOST_NOW_PRIMARY HostEventTypeForGroup = "HOST_NOW_PRIMARY"
	HOSTEVENTTYPEFORGROUP_HOST_NOW_SECONDARY HostEventTypeForGroup = "HOST_NOW_SECONDARY"
	HOSTEVENTTYPEFORGROUP_HOST_NOW_STANDALONE HostEventTypeForGroup = "HOST_NOW_STANDALONE"
)

// All allowed values of HostEventTypeForGroup enum
var AllowedHostEventTypeForGroupEnumValues = []HostEventTypeForGroup{
	"ADD_HOST_AUDIT",
	"DELETE_HOST_AUDIT",
	"UNDELETE_HOST_AUDIT",
	"ADD_HOST_TO_REPLICA_SET_AUDIT",
	"REMOVE_HOST_FROM_REPLICA_SET_AUDIT",
	"HIDE_HOST_AUDIT",
	"HIDE_AND_DISABLE_HOST_AUDIT",
	"DISABLE_HOST_AUDIT",
	"PAUSE_HOST_AUDIT",
	"RESUME_HOST_AUDIT",
	"DB_PROFILER_ENABLE_AUDIT",
	"DB_PROFILER_DISABLE_AUDIT",
	"HOST_IP_CHANGED_AUDIT",
	"AUTO_CREATED_INDEX_AUDIT",
	"ATTEMPT_KILLOP_AUDIT",
	"ATTEMPT_KILLSESSION_AUDIT",
	"HOST_UP",
	"HOST_DOWN",
	"HOST_ROLLBACK",
	"HOST_RECOVERED",
	"HOST_RECOVERING",
	"VERSION_CURRENT",
	"VERSION_BEHIND",
	"VALUE_NO_LONGER_ANOMALOUS",
	"HOST_LOCKED_DOWN",
	"HOST_EXPOSED",
	"VERSION_CHANGED",
	"HOST_SSL_CERTIFICATE_CURRENT",
	"HOST_SSL_CERTIFICATE_STALE",
	"HOST_HAS_INDEX_SUGGESTIONS",
	"HOST_SECURITY_CHECKUP_MET",
	"HOST_SECURITY_CHECKUP_NOT_MET",
	"HOST_MONGOT_RECOVERED_OOM",
	"HOST_MONGOT_CRASHING_OOM",
	"HOST_DISK_SPACE_SUFFICIENT_FOR_SEARCH_INDEX_REBUILD",
	"HOST_DISK_SPACE_INSUFFICIENT_FOR_SEARCH_INDEX_REBUILD",
	"NEW_HOST",
	"HOST_RESTARTED",
	"HOST_UPGRADED",
	"HOST_DOWNGRADED",
	"HOST_NOW_PRIMARY",
	"HOST_NOW_SECONDARY",
	"HOST_NOW_STANDALONE",
}

func (v *HostEventTypeForGroup) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HostEventTypeForGroup(value)
	for _, existing := range AllowedHostEventTypeForGroupEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HostEventTypeForGroup", value)
}

// NewHostEventTypeForGroupFromValue returns a pointer to a valid HostEventTypeForGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHostEventTypeForGroupFromValue(v string) (*HostEventTypeForGroup, error) {
	ev := HostEventTypeForGroup(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HostEventTypeForGroup: valid values are %v", v, AllowedHostEventTypeForGroupEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HostEventTypeForGroup) IsValid() bool {
	for _, existing := range AllowedHostEventTypeForGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostEventTypeForGroup value
func (v HostEventTypeForGroup) Ptr() *HostEventTypeForGroup {
	return &v
}

type NullableHostEventTypeForGroup struct {
	value *HostEventTypeForGroup
	isSet bool
}

func (v NullableHostEventTypeForGroup) Get() *HostEventTypeForGroup {
	return v.value
}

func (v *NullableHostEventTypeForGroup) Set(val *HostEventTypeForGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableHostEventTypeForGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableHostEventTypeForGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostEventTypeForGroup(val *HostEventTypeForGroup) *NullableHostEventTypeForGroup {
	return &NullableHostEventTypeForGroup{value: val, isSet: true}
}

func (v NullableHostEventTypeForGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostEventTypeForGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
