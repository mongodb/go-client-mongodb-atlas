/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// HostEventTypeViewForNdsGroup Unique identifier of event type.
type HostEventTypeViewForNdsGroup string

// List of HostEventTypeViewForNdsGroup
const (
	HOSTEVENTTYPEVIEWFORNDSGROUP_AUTO_CREATED_INDEX_AUDIT HostEventTypeViewForNdsGroup = "AUTO_CREATED_INDEX_AUDIT"
	HOSTEVENTTYPEVIEWFORNDSGROUP_ATTEMPT_KILLOP_AUDIT HostEventTypeViewForNdsGroup = "ATTEMPT_KILLOP_AUDIT"
	HOSTEVENTTYPEVIEWFORNDSGROUP_ATTEMPT_KILLSESSION_AUDIT HostEventTypeViewForNdsGroup = "ATTEMPT_KILLSESSION_AUDIT"
	HOSTEVENTTYPEVIEWFORNDSGROUP_HOST_UP HostEventTypeViewForNdsGroup = "HOST_UP"
	HOSTEVENTTYPEVIEWFORNDSGROUP_HOST_DOWN HostEventTypeViewForNdsGroup = "HOST_DOWN"
	HOSTEVENTTYPEVIEWFORNDSGROUP_HOST_HAS_INDEX_SUGGESTIONS HostEventTypeViewForNdsGroup = "HOST_HAS_INDEX_SUGGESTIONS"
	HOSTEVENTTYPEVIEWFORNDSGROUP_HOST_MONGOT_RECOVERED_OOM HostEventTypeViewForNdsGroup = "HOST_MONGOT_RECOVERED_OOM"
	HOSTEVENTTYPEVIEWFORNDSGROUP_HOST_MONGOT_CRASHING_OOM HostEventTypeViewForNdsGroup = "HOST_MONGOT_CRASHING_OOM"
	HOSTEVENTTYPEVIEWFORNDSGROUP_HOST_DISK_SPACE_SUFFICIENT_FOR_SEARCH_INDEX_REBUILD HostEventTypeViewForNdsGroup = "HOST_DISK_SPACE_SUFFICIENT_FOR_SEARCH_INDEX_REBUILD"
	HOSTEVENTTYPEVIEWFORNDSGROUP_HOST_ENOUGH_DISK_SPACE HostEventTypeViewForNdsGroup = "HOST_ENOUGH_DISK_SPACE"
	HOSTEVENTTYPEVIEWFORNDSGROUP_HOST_DISK_SPACE_INSUFFICIENT_FOR_SEARCH_INDEX_REBUILD HostEventTypeViewForNdsGroup = "HOST_DISK_SPACE_INSUFFICIENT_FOR_SEARCH_INDEX_REBUILD"
	HOSTEVENTTYPEVIEWFORNDSGROUP_HOST_NOT_ENOUGH_DISK_SPACE HostEventTypeViewForNdsGroup = "HOST_NOT_ENOUGH_DISK_SPACE"
)

// All allowed values of HostEventTypeViewForNdsGroup enum
var AllowedHostEventTypeViewForNdsGroupEnumValues = []HostEventTypeViewForNdsGroup{
	"AUTO_CREATED_INDEX_AUDIT",
	"ATTEMPT_KILLOP_AUDIT",
	"ATTEMPT_KILLSESSION_AUDIT",
	"HOST_UP",
	"HOST_DOWN",
	"HOST_HAS_INDEX_SUGGESTIONS",
	"HOST_MONGOT_RECOVERED_OOM",
	"HOST_MONGOT_CRASHING_OOM",
	"HOST_DISK_SPACE_SUFFICIENT_FOR_SEARCH_INDEX_REBUILD",
	"HOST_ENOUGH_DISK_SPACE",
	"HOST_DISK_SPACE_INSUFFICIENT_FOR_SEARCH_INDEX_REBUILD",
	"HOST_NOT_ENOUGH_DISK_SPACE",
}

func (v *HostEventTypeViewForNdsGroup) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HostEventTypeViewForNdsGroup(value)
	for _, existing := range AllowedHostEventTypeViewForNdsGroupEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HostEventTypeViewForNdsGroup", value)
}

// NewHostEventTypeViewForNdsGroupFromValue returns a pointer to a valid HostEventTypeViewForNdsGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHostEventTypeViewForNdsGroupFromValue(v string) (*HostEventTypeViewForNdsGroup, error) {
	ev := HostEventTypeViewForNdsGroup(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HostEventTypeViewForNdsGroup: valid values are %v", v, AllowedHostEventTypeViewForNdsGroupEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HostEventTypeViewForNdsGroup) IsValid() bool {
	for _, existing := range AllowedHostEventTypeViewForNdsGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostEventTypeViewForNdsGroup value
func (v HostEventTypeViewForNdsGroup) Ptr() *HostEventTypeViewForNdsGroup {
	return &v
}

