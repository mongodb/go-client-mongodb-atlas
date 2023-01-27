// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/openlyinc/pointy"
)

func TestAlert_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	alertID := "57b76ddc96e8215c017ceafb"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alerts/%s", groupID, alertID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": "533dc40ae4b00835ff81eaee",
			"groupId": "535683b3794d371327b",
			"eventTypeName": "OUTSIDE_METRIC_THRESHOLD",
			"created": "2016-08-23T20:26:50Z",
			"updated": "2016-08-23T20:26:50Z",
			"enabled": true,
			"matchers": [
				{
					"fieldName": "HOSTNAME_AND_PORT",
					"operator": "EQUALS",
					"value": "mongo.example.com:27017"
				}
			],
			"notifications": [
				{
					"typeName": "SMS",
					"intervalMin": 5,
					"delayMin": 0,
					"mobileNumber": "2343454567"
				}
			],
			"metricThreshold": {
				"metricName": "ASSERT_REGULAR",
				"operator": "LESS_THAN",
				"threshold": 99.0,
				"units": "RAW",
				"mode": "AVERAGE"
			}
		}`)
	})

	alert, _, err := client.Alerts.Get(ctx, groupID, alertID)
	if err != nil {
		t.Fatalf("Alerts.Get returned error: %v", err)
	}

	expected := &Alert{
		ID:            "533dc40ae4b00835ff81eaee",
		GroupID:       "535683b3794d371327b",
		EventTypeName: "OUTSIDE_METRIC_THRESHOLD",
		Created:       "2016-08-23T20:26:50Z",
		Updated:       "2016-08-23T20:26:50Z",
		Enabled:       pointy.Bool(true),
		Matchers: []Matcher{
			{
				FieldName: "HOSTNAME_AND_PORT",
				Operator:  "EQUALS",
				Value:     "mongo.example.com:27017",
			},
		},
		Notifications: []Notification{
			{
				TypeName:     "SMS",
				IntervalMin:  5,
				DelayMin:     pointy.Int(0),
				MobileNumber: "2343454567",
			},
		},
		MetricThreshold: &MetricThreshold{
			MetricName: "ASSERT_REGULAR",
			Operator:   "LESS_THAN",
			Threshold:  99.0,
			Units:      "RAW",
			Mode:       "AVERAGE",
		},
	}

	if diff := deep.Equal(alert, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAlerts_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alerts", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"results": [
				{
					"id": "533dc40ae4b00835ff81eaee",
					"groupId": "535683b3794d371327b",
					"eventTypeName": "OUTSIDE_METRIC_THRESHOLD",
					"created": "2016-08-23T20:26:50Z",
					"updated": "2016-08-23T20:26:50Z",
					"enabled": true,
					"matchers": [
						{
							"fieldName": "HOSTNAME_AND_PORT",
							"operator": "EQUALS",
							"value": "mongo.example.com:27017"
						}
					],
					"notifications": [
						{
							"typeName": "SMS",
							"intervalMin": 5,
							"delayMin": 0,
							"mobileNumber": "2343454567"
						}
					],
					"metricThreshold": {
						"metricName": "ASSERT_REGULAR",
						"operator": "LESS_THAN",
						"threshold": 99.0,
						"units": "RAW",
						"mode": "AVERAGE"
					}
				},
				{
					"id": "533dc40ae4b00835ff81eaee",
					"groupId": "535683b3794d371327b",
					"eventTypeName": "OUTSIDE_METRIC_THRESHOLD",
					"created": "2016-08-23T20:26:50Z",
					"updated": "2016-08-23T20:26:50Z",
					"enabled": true,
					"matchers": [
						{
							"fieldName": "HOSTNAME_AND_PORT",
							"operator": "EQUALS",
							"value": "mongo.example.com:27017"
						}
					],
					"notifications": [
						{
							"typeName": "SMS",
							"intervalMin": 5,
							"delayMin": 0,
							"mobileNumber": "2343454567"
						}
					],
					"metricThreshold": {
						"metricName": "ASSERT_REGULAR",
						"operator": "LESS_THAN",
						"threshold": 99.0,
						"units": "RAW",
						"mode": "AVERAGE"
					}
				}
			],
			"totalCount": 2
		}`)
	})

	alerts, _, err := client.Alerts.List(ctx, groupID, nil)
	if err != nil {
		t.Fatalf("Alerts.List returned error: %v", err)
	}

	expected := &AlertsResponse{
		Links: nil,
		Results: []Alert{
			{
				ID:            "533dc40ae4b00835ff81eaee",
				GroupID:       "535683b3794d371327b",
				EventTypeName: "OUTSIDE_METRIC_THRESHOLD",
				Created:       "2016-08-23T20:26:50Z",
				Updated:       "2016-08-23T20:26:50Z",
				Enabled:       pointy.Bool(true),
				Matchers: []Matcher{
					{
						FieldName: "HOSTNAME_AND_PORT",
						Operator:  "EQUALS",
						Value:     "mongo.example.com:27017",
					},
				},
				Notifications: []Notification{
					{
						TypeName:     "SMS",
						IntervalMin:  5,
						DelayMin:     pointy.Int(0),
						MobileNumber: "2343454567",
					},
				},
				MetricThreshold: &MetricThreshold{
					MetricName: "ASSERT_REGULAR",
					Operator:   "LESS_THAN",
					Threshold:  99.0,
					Units:      "RAW",
					Mode:       "AVERAGE",
				},
			},
			{
				ID:            "533dc40ae4b00835ff81eaee",
				GroupID:       "535683b3794d371327b",
				EventTypeName: "OUTSIDE_METRIC_THRESHOLD",
				Created:       "2016-08-23T20:26:50Z",
				Updated:       "2016-08-23T20:26:50Z",
				Enabled:       pointy.Bool(true),
				Matchers: []Matcher{
					{
						FieldName: "HOSTNAME_AND_PORT",
						Operator:  "EQUALS",
						Value:     "mongo.example.com:27017",
					},
				},
				Notifications: []Notification{
					{
						TypeName:     "SMS",
						IntervalMin:  5,
						DelayMin:     pointy.Int(0),
						MobileNumber: "2343454567",
					},
				},
				MetricThreshold: &MetricThreshold{
					MetricName: "ASSERT_REGULAR",
					Operator:   "LESS_THAN",
					Threshold:  99.0,
					Units:      "RAW",
					Mode:       "AVERAGE",
				},
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(alerts, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAlert_Acknowledge(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	alertID := "533dc40ae4b00835ff81eaee"
	acknowledgedUntil := "2026-10-01T00:00:00-0400"

	params := AcknowledgeRequest{
		AcknowledgedUntil:      &acknowledgedUntil,
		AcknowledgementComment: "This is normal. Please ignore.",
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alerts/%s", groupID, alertID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
			"id": "533dc40ae4b00835ff81eaee",
			"groupId": "535683b3794d371327b",
			"eventTypeName": "OUTSIDE_METRIC_THRESHOLD",
			"status" : "OPEN",
			"acknowledgedUntil" : "2016-10-01T04:00:00Z",
			"acknowledgementComment" : "This is normal. Please ignore.",
			"acknowledgingUsername" : "someuser@example.com",
			"created": "2016-08-23T20:26:50Z",
			"updated": "2016-08-23T20:26:50Z",
			"lastNotified" : "2016-08-23T20:33:23Z",
			"metricName" : "ASSERTS_REGULAR",
			"currentValue" : {
				"number" : 0.0,
				"units" : "RAW"
  			}
		}`)
	})

	alert, _, err := client.Alerts.Acknowledge(ctx, groupID, alertID, &params)
	if err != nil {
		t.Fatalf("Alerts.Acknowledge returned error: %v", err)
	}

	expected := &Alert{
		ID:                     "533dc40ae4b00835ff81eaee",
		GroupID:                "535683b3794d371327b",
		EventTypeName:          "OUTSIDE_METRIC_THRESHOLD",
		Status:                 "OPEN",
		AcknowledgedUntil:      "2016-10-01T04:00:00Z",
		AcknowledgementComment: "This is normal. Please ignore.",
		AcknowledgingUsername:  "someuser@example.com",
		Created:                "2016-08-23T20:26:50Z",
		Updated:                "2016-08-23T20:26:50Z",
		LastNotified:           "2016-08-23T20:33:23Z",
		MetricName:             "ASSERTS_REGULAR",
		CurrentValue: &CurrentValue{
			Number: pointy.Float64(0.0),
			Units:  "RAW",
		},
	}

	if diff := deep.Equal(alert, expected); diff != nil {
		t.Error(diff)
	}
}
