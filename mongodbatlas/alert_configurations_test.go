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
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/openlyinc/pointy"
)

func TestAlertConfiguration_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alertConfigs", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		expected := map[string]interface{}{
			"eventTypeName": "NO_PRIMARY",
			"enabled":       true,
			"notifications": []interface{}{
				map[string]interface{}{
					"typeName":     "GROUP",
					"intervalMin":  float64(5),
					"delayMin":     float64(0),
					"smsEnabled":   false,
					"emailEnabled": true,
					"roles":        []interface{}{"GROUP_CHARTS_ADMIN", "GROUP_CLUSTER_MANAGER"},
				},
			},
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err.Error())
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, `{
			"id": "57b76ddc96e8215c017ceafb",
			"created": "2016-08-19T20:36:44Z",
			"enabled": true,
			"eventTypeName": "NO_PRIMARY",
			"groupId": "535683b3794d371327b",
			"matchers": [],
			"notifications": [
				{
					"delayMin": 0,
					"emailEnabled": true,
					"intervalMin": 5,
					"smsEnabled": false,
					"typeName": "GROUP",
					"roles" : ["GROUP_CHARTS_ADMIN", "GROUP_CLUSTER_MANAGER"]
				}
			],
			"updated": "2016-08-19T20:36:44Z"
		}`)
	})

	alertReq := &AlertConfiguration{
		EventTypeName: "NO_PRIMARY",
		Enabled:       pointy.Bool(true),
		Notifications: []Notification{
			{
				TypeName:     "GROUP",
				IntervalMin:  5,
				DelayMin:     pointy.Int(0),
				SMSEnabled:   pointy.Bool(false),
				EmailEnabled: pointy.Bool(true),
				Roles:        []string{"GROUP_CHARTS_ADMIN", "GROUP_CLUSTER_MANAGER"},
			},
		},
	}

	alertConfiguration, _, err := client.AlertConfigurations.Create(ctx, groupID, alertReq)
	if err != nil {
		t.Fatalf("AlertConfiguration.Create returned error: %v", err)
		return
	}

	expected := &AlertConfiguration{
		ID:            "57b76ddc96e8215c017ceafb",
		Created:       "2016-08-19T20:36:44Z",
		Enabled:       pointy.Bool(true),
		EventTypeName: "NO_PRIMARY",
		GroupID:       "535683b3794d371327b",
		Matchers:      []Matcher{},
		Notifications: []Notification{
			{
				DelayMin:     pointy.Int(0),
				EmailEnabled: pointy.Bool(true),
				IntervalMin:  5,
				SMSEnabled:   pointy.Bool(false),
				TypeName:     "GROUP",
				Roles:        []string{"GROUP_CHARTS_ADMIN", "GROUP_CLUSTER_MANAGER"},
			},
		},
		Updated: "2016-08-19T20:36:44Z",
	}

	if diff := deep.Equal(alertConfiguration, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAlertConfiguration_EnableAnAlertConfig(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"
	alertConfigID := "57b76ddc96e8215c017ceafb"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alertConfigs/%s", groupID, alertConfigID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		expected := map[string]interface{}{
			"enabled": true,
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, `{
			"created": "2016-08-19T20:45:29Z",
			"enabled": false,
			"eventTypeName": "NO_PRIMARY",
			"groupId": "535683b3794d371327b",
			"id": "57b76ddc96e8215c017ceafb",
			"matchers": [],
			"notifications": [
				{
					"delayMin": 5,
					"emailEnabled": false,
					"intervalMin": 5,
					"smsEnabled": true,
					"typeName": "GROUP"
				}
			],
			"updated": "2016-08-19T20:51:49Z"
		}`)
	})

	alertConfiguration, _, err := client.AlertConfigurations.EnableAnAlertConfig(ctx, groupID, alertConfigID, pointy.Bool(true))
	if err != nil {
		t.Fatalf("AlertConfiguration.EnableAnAlertConfig returned error: %v", err)
	}

	expected := &AlertConfiguration{
		Created:       "2016-08-19T20:45:29Z",
		Enabled:       pointy.Bool(false),
		EventTypeName: "NO_PRIMARY",
		GroupID:       "535683b3794d371327b",
		ID:            "57b76ddc96e8215c017ceafb",
		Matchers:      []Matcher{},
		Notifications: []Notification{
			{
				DelayMin:     pointy.Int(5),
				EmailEnabled: pointy.Bool(false),
				IntervalMin:  5,
				SMSEnabled:   pointy.Bool(true),
				TypeName:     "GROUP",
			},
		},
		Updated: "2016-08-19T20:51:49Z",
	}

	if diff := deep.Equal(alertConfiguration, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAlertConfiguration_GetAnAlertConfig(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"
	alertConfigID := "57b76ddc96e8215c017ceafb"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alertConfigs/%s", groupID, alertConfigID), func(w http.ResponseWriter, r *http.Request) {
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
					"mobileNumber": "2343454567",
					"roles" : ["GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_ONLY"]
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

	alertConfiguration, _, err := client.AlertConfigurations.GetAnAlertConfig(ctx, groupID, alertConfigID)
	if err != nil {
		t.Fatalf("AlertConfigurations.GetAnAlertConfig returned error: %v", err)
	}

	expected := &AlertConfiguration{
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
				Roles:        []string{"GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_ONLY"},
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

	if diff := deep.Equal(alertConfiguration, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAlertConfiguration_GetOpenAlertsConfig(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"
	alertConfigID := "57b76ddc96e8215c017ceafb"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alertConfigs/%s/alerts", groupID, alertConfigID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"totalCount": 2,
			"results": [
				{
					"id": "53569159300495c7702ee3a3",
					"groupId": "535683b3794d371327b",
					"eventTypeName": "OUTSIDE_METRIC_THRESHOLD",
					"status": "OPEN",
					"acknowledgedUntil": "2016-09-01T14:00:00Z",
					"created": "2016-08-22T15:57:13.562Z",
					"updated": "2016-08-22T20:14:11.388Z",
					"lastNotified": "2016-08-22T15:57:24.126Z",
					"metricName": "ASSERT_REGULAR",
					"currentValue": {
						"number": 0.0,
						"units": "RAW"
					}
				},
				{
					"id": "5356ca0e300495c770333340",
					"groupId": "535683b3794d371327b",
					"eventTypeName": "OUTSIDE_METRIC_THRESHOLD",
					"status": "OPEN",
					"created": "2016-08-22T19:59:10.657Z",
					"updated": "2016-08-22T20:14:11.388Z",
					"lastNotified": "2016-08-22T20:14:19.313Z",
					"metricName": "ASSERT_REGULAR",
					"currentValue": {
						"number": 0.0,
						"units": "RAW"
					}
				}
			]
		}`)
	})

	alertConfigurations, _, err := client.AlertConfigurations.GetOpenAlertsConfig(ctx, groupID, alertConfigID)
	if err != nil {
		t.Fatalf("AlertConfigurations.GetOpenAlertsConfig returned error: %v", err)
	}

	expected := []AlertConfiguration{
		{
			ID:                "53569159300495c7702ee3a3",
			GroupID:           "535683b3794d371327b",
			EventTypeName:     "OUTSIDE_METRIC_THRESHOLD",
			Status:            "OPEN",
			AcknowledgedUntil: "2016-09-01T14:00:00Z",
			Created:           "2016-08-22T15:57:13.562Z",
			Updated:           "2016-08-22T20:14:11.388Z",
			LastNotified:      "2016-08-22T15:57:24.126Z",
			MetricName:        "ASSERT_REGULAR",
			CurrentValue: &CurrentValue{
				Number: pointy.Float64(0.0),
				Units:  "RAW",
			},
		},
		{
			ID:            "5356ca0e300495c770333340",
			GroupID:       "535683b3794d371327b",
			EventTypeName: "OUTSIDE_METRIC_THRESHOLD",
			Status:        "OPEN",
			Created:       "2016-08-22T19:59:10.657Z",
			Updated:       "2016-08-22T20:14:11.388Z",
			LastNotified:  "2016-08-22T20:14:19.313Z",
			MetricName:    "ASSERT_REGULAR",
			CurrentValue: &CurrentValue{
				Number: pointy.Float64(0.0),
				Units:  "RAW",
			},
		},
	}

	if diff := deep.Equal(alertConfigurations, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAlertConfiguration_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alertConfigs", groupID), func(w http.ResponseWriter, r *http.Request) {
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
							"mobileNumber": "2343454567",
							"roles" : ["GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_ONLY"]
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
							"mobileNumber": "2343454567",
							"roles" : ["GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_ONLY"]
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

	alertConfigurations, _, err := client.AlertConfigurations.List(ctx, groupID, nil)
	if err != nil {
		t.Fatalf("AlertConfigurations.List returned error: %v", err)
	}

	expected := []AlertConfiguration{
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
					Roles:        []string{"GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_ONLY"},
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
					Roles:        []string{"GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_ONLY"},
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
	}

	if diff := deep.Equal(alertConfigurations, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAlertConfiguration_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6d2065c687d9d64ae7acdg41"
	alertConfigID := "57b76ddc96e8215c017ceafb"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alertConfigs/%s", groupID, alertConfigID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)

		expected := map[string]interface{}{
			"eventTypeName": "NO_PRIMARY",
			"enabled":       true,
			"notifications": []interface{}{
				map[string]interface{}{
					"typeName":     "GROUP",
					"intervalMin":  float64(5),
					"delayMin":     float64(5),
					"smsEnabled":   true,
					"emailEnabled": false,
					"roles":        []interface{}{"GROUP_CHARTS_ADMIN", "GROUP_CLUSTER_MANAGER"},
				},
			},
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err.Error())
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Request body\n got=%#v\nwant=%#v \n\ndiff=%#v", v, expected, diff)
		}

		fmt.Fprint(w, `{
			"created": "2016-08-19T20:45:29Z",
			"enabled": true,
			"eventTypeName": "NO_PRIMARY",
			"groupId": "535683b3794d371327b",
			"id": "57b76ddc96e8215c017ceafb",
			"notifications": [
				{
					"typeName": "GROUP",
					"intervalMin": 5,
					"delayMin": 5,
					"smsEnabled": true,
					"emailEnabled": false,
					"roles" : ["GROUP_CHARTS_ADMIN", "GROUP_CLUSTER_MANAGER"]
				}
			],
			"updated": "2016-08-19T20:45:29Z"
		}`)
	})

	alertReq := &AlertConfiguration{
		EventTypeName: "NO_PRIMARY",
		Enabled:       pointy.Bool(true),
		Notifications: []Notification{
			{
				TypeName:     "GROUP",
				IntervalMin:  5,
				DelayMin:     pointy.Int(5),
				SMSEnabled:   pointy.Bool(true),
				EmailEnabled: pointy.Bool(false),
				Roles:        []string{"GROUP_CHARTS_ADMIN", "GROUP_CLUSTER_MANAGER"},
			},
		},
	}

	alertConfiguration, _, err := client.AlertConfigurations.Update(ctx, groupID, alertConfigID, alertReq)
	if err != nil {
		t.Fatalf("AlertConfiguration.Update returned error: %v", err)
		return
	}

	expected := &AlertConfiguration{
		Created:       "2016-08-19T20:45:29Z",
		Enabled:       pointy.Bool(true),
		EventTypeName: "NO_PRIMARY",
		GroupID:       "535683b3794d371327b",
		ID:            "57b76ddc96e8215c017ceafb",
		Notifications: []Notification{
			{
				TypeName:     "GROUP",
				IntervalMin:  5,
				DelayMin:     pointy.Int(5),
				SMSEnabled:   pointy.Bool(true),
				EmailEnabled: pointy.Bool(false),
				Roles:        []string{"GROUP_CHARTS_ADMIN", "GROUP_CLUSTER_MANAGER"},
			},
		},
		Updated: "2016-08-19T20:45:29Z",
	}

	if diff := deep.Equal(alertConfiguration, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAlertConfiguration_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"
	alertConfigID := "57b76ddc96e8215c017ceafb"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alertConfigs/%s", groupID, alertConfigID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.AlertConfigurations.Delete(ctx, groupID, alertConfigID)
	if err != nil {
		t.Fatalf("AlertConfigurations.Delete returned error: %v", err)
	}
}

func TestAlertConfiguration_ListMatcherFields(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/alertConfigs/matchers/fieldNames", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		fmt.Fprint(w, `
				[
					"TYPE_NAME",
					"HOSTNAME",
					"PORT",
					"HOSTNAME_AND_PORT",
					"REPLICA_SET_NAME",
					"SHARD_NAME",
					"CLUSTER_NAME"
				]
		`)
	})

	alertConfigurations, _, err := client.AlertConfigurations.ListMatcherFields(ctx)
	if err != nil {
		t.Fatalf("AlertConfigurations.ListMatcherFields returned error: %v", err)
	}

	expected := []string{
		"TYPE_NAME",
		"HOSTNAME",
		"PORT",
		"HOSTNAME_AND_PORT",
		"REPLICA_SET_NAME",
		"SHARD_NAME",
		"CLUSTER_NAME",
	}

	if diff := deep.Equal(alertConfigurations, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAlertConfiguration_GetAnAlertConfigTeams(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"
	alertConfigID := "57b76ddc96e8215c017ceafb"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alertConfigs/%s", groupID, alertConfigID), func(w http.ResponseWriter, r *http.Request) {
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
					"typeName": "MICROSOFT_TEAMS",
					"intervalMin": 5,
					"delayMin": 0,
					"microsoftTeamsWebhookUrl": "http://941a-47-225-212-178.ngrok.io",
					"roles" : ["GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_ONLY"]
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

	alertConfiguration, _, err := client.AlertConfigurations.GetAnAlertConfig(ctx, groupID, alertConfigID)
	if err != nil {
		t.Fatalf("AlertConfigurations.GetAnAlertConfigTeams returned error: %v", err)
	}

	expected := &AlertConfiguration{
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
				TypeName:                 "MICROSOFT_TEAMS",
				IntervalMin:              5,
				DelayMin:                 pointy.Int(0),
				MicrosoftTeamsWebhookURL: "http://941a-47-225-212-178.ngrok.io",
				Roles:                    []string{"GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_ONLY"},
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

	if diff := deep.Equal(alertConfiguration, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAlertConfiguration_GetAnAlertConfigWebhook(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"
	alertConfigID := "57b76ddc96e8215c017ceafb"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/alertConfigs/%s", groupID, alertConfigID), func(w http.ResponseWriter, r *http.Request) {
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
					"typeName": "WEBHOOK",
					"intervalMin": 5,
					"delayMin": 0,
					"webhookSecret": "SECRET",
					"webhookUrl": "http://941a-47-225-212-178.ngrok.io",
					"roles" : ["GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_ONLY"]
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

	alertConfiguration, _, err := client.AlertConfigurations.GetAnAlertConfig(ctx, groupID, alertConfigID)
	if err != nil {
		t.Fatalf("AlertConfigurations.GetAnAlertConfig returned error: %v", err)
	}

	expected := &AlertConfiguration{
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
				TypeName:      "WEBHOOK",
				IntervalMin:   5,
				DelayMin:      pointy.Int(0),
				WebhookSecret: "SECRET",
				WebhookURL:    "http://941a-47-225-212-178.ngrok.io",
				Roles:         []string{"GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_ONLY"},
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

	if diff := deep.Equal(alertConfiguration, expected); diff != nil {
		t.Error(diff)
	}
}
