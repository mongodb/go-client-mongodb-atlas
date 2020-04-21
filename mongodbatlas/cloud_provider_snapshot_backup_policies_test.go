package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestCloudProviderSnapshotBackupPolicies_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	groupID := "5b6212af90dc76637950a2c6"
	clusterName := "myCluster"

	path := fmt.Sprintf("/groups/%s/clusters/%s/backup/schedule", groupID, clusterName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"clusterId": "5e2f1bcaf38990fab9227b8",
			"clusterName": "myCluster",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5dd5a6a472efab1d71a58495/clusters/myCluster/backup/schedule",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/public/v1.0/groups/5dd5a6a472efab1d71a58495",
					"rel": "http://mms.mongodb.com/group"
				}
			],
			"nextSnapshot": "2020-01-28T05:24:25Z",
			"policies": [
				{
					"id": "5e2f1bcaf38990fab9227b8",
					"policyItems": [
						{
							"frequencyInterval": 6,
							"frequencyType": "hourly",
							"id": "5e2f1cc8105eef6d6bd9005b",
							"retentionUnit": "days",
							"retentionValue": 7
						},
						{
							"frequencyInterval": 1,
							"frequencyType": "daily",
							"id": "5e2f1cc8105eef6d6bd9005c",
							"retentionUnit": "days",
							"retentionValue": 7
						},
						{
							"frequencyInterval": 6,
							"frequencyType": "weekly",
							"id": "5e2f1cc8105eef6d6bd9005d",
							"retentionUnit": "weeks",
							"retentionValue": 4
						},
						{
							"frequencyInterval": 40,
							"frequencyType": "monthly",
							"id": "5e2f1cc8105eef6d6bd9005e",
							"retentionUnit": "months",
							"retentionValue": 12
						}
					]
				}
			],
			"referenceHourOfDay": 17,
			"referenceMinuteOfHour": 24,
			"restoreWindowDays": 7
		}`)
	})

	snapshotBackupPolicy, _, err := client.CloudProviderSnapshotBackupPolicies.Get(ctx, groupID, clusterName)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotBackupPolicies.Get returned error: %v", err)
	}

	expected := &CloudProviderSnapshotBackupPolicy{
		ClusterID:    "5e2f1bcaf38990fab9227b8",
		ClusterName:  "myCluster",
		NextSnapshot: "2020-01-28T05:24:25Z",
		Policies: []Policy{
			{
				ID: "5e2f1bcaf38990fab9227b8",
				PolicyItems: []PolicyItem{
					{
						FrequencyInterval: 6,
						FrequencyType:     "hourly",
						ID:                "5e2f1cc8105eef6d6bd9005b",
						RetentionUnit:     "days",
						RetentionValue:    7,
					},
					{
						FrequencyInterval: 1,
						FrequencyType:     "daily",
						ID:                "5e2f1cc8105eef6d6bd9005c",
						RetentionUnit:     "days",
						RetentionValue:    7,
					},
					{
						FrequencyInterval: 6,
						FrequencyType:     "weekly",
						ID:                "5e2f1cc8105eef6d6bd9005d",
						RetentionUnit:     "weeks",
						RetentionValue:    4,
					},
					{
						FrequencyInterval: 40,
						FrequencyType:     "monthly",
						ID:                "5e2f1cc8105eef6d6bd9005e",
						RetentionUnit:     "months",
						RetentionValue:    12,
					},
				},
			},
		},
		ReferenceHourOfDay:    17,
		ReferenceMinuteOfHour: 24,
		RestoreWindowDays:     7,
	}

	if diff := deep.Equal(snapshotBackupPolicy, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotBackupPolicies_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	groupID := "5b6212af90dc76637950a2c6"
	clusterName := "myCluster"

	path := fmt.Sprintf("/groups/%s/clusters/%s/backup/schedule", groupID, clusterName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"referenceHourOfDay":    float64(12),
			"referenceMinuteOfHour": float64(30),
			"policies": []interface{}{
				map[string]interface{}{
					"id": "5c95242c87d9d636e70c28ef",
					"policyItems": []interface{}{
						map[string]interface{}{
							"id":                "5c95242c87d9d636e70c28f0",
							"frequencyType":     "hourly",
							"frequencyInterval": float64(6),
							"retentionValue":    float64(2),
							"retentionUnit":     "days",
						},
						map[string]interface{}{
							"id":                "5c95242c87d9d636e70c28f2",
							"frequencyType":     "weekly",
							"frequencyInterval": float64(1),
							"retentionValue":    float64(3),
							"retentionUnit":     "weeks",
						},
					},
				},
			},
			"updateSnapshots": true,
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("Decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, `{
			"clusterId": "5c94f6ea80eef5617167224d",
			"clusterName": "Cluster0",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/schedule",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
					"rel": "http://mms.mongodb.com/group"
				}
			],
			"nextSnapshot": "2019-04-03T18:30:08Z",
			"policies": [
				{
					"id": "5c95242c87d9d636e70c28ef",
					"policyItems": [
						{
							"frequencyInterval": 6,
							"frequencyType": "hourly",
							"id": "5c95242c87d9d636e70c28f0",
							"retentionUnit": "days",
							"retentionValue": 2
						},
						{
							"frequencyInterval": 1,
							"frequencyType": "weekly",
							"id": "5c95242c87d9d636e70c28f2",
							"retentionUnit": "weeks",
							"retentionValue": 3
						}
					]
				}
			],
			"referenceHourOfDay": 12,
			"referenceMinuteOfHour": 30
		}`)
	})

	updateRequest := &CloudProviderSnapshotBackupPolicy{
		ReferenceHourOfDay:    12,
		ReferenceMinuteOfHour: 30,
		Policies: []Policy{
			{
				ID: "5c95242c87d9d636e70c28ef",
				PolicyItems: []PolicyItem{
					{
						ID:                "5c95242c87d9d636e70c28f0",
						FrequencyType:     "hourly",
						FrequencyInterval: 6,
						RetentionValue:    2,
						RetentionUnit:     "days",
					},
					{
						ID:                "5c95242c87d9d636e70c28f2",
						FrequencyType:     "weekly",
						FrequencyInterval: 1,
						RetentionValue:    3,
						RetentionUnit:     "weeks",
					},
				},
			},
		},
		UpdateSnapshots: true,
	}

	cloudProviderSnapshot, _, err := client.CloudProviderSnapshotBackupPolicies.Update(ctx, groupID, clusterName, updateRequest)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotBackupPolicies.Update returned error: %v", err)
	}

	expected := &CloudProviderSnapshotBackupPolicy{
		ClusterID:    "5c94f6ea80eef5617167224d",
		ClusterName:  "Cluster0",
		NextSnapshot: "2019-04-03T18:30:08Z",
		Policies: []Policy{
			{
				ID: "5c95242c87d9d636e70c28ef",
				PolicyItems: []PolicyItem{
					{
						FrequencyInterval: 6,
						FrequencyType:     "hourly",
						ID:                "5c95242c87d9d636e70c28f0",
						RetentionUnit:     "days",
						RetentionValue:    2,
					},
					{
						FrequencyInterval: 1,
						FrequencyType:     "weekly",
						ID:                "5c95242c87d9d636e70c28f2",
						RetentionUnit:     "weeks",
						RetentionValue:    3,
					},
				},
			},
		},
		ReferenceHourOfDay:    12,
		ReferenceMinuteOfHour: 30,
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}
