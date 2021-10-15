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
)

func TestCloudProviderSnapshots_GetAllCloudProviderSnapshots(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	requestParameters := &SnapshotReqPathParameters{
		GroupID:     "5b6212af90dc76637950a2c6",
		ClusterName: "MyCluster",
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/snapshots", requestParameters.GroupID, requestParameters.ClusterName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/MyCluster/backup/snapshots?pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				{
					"cloudProvider": "AZURE",
					"createdAt": "2018-08-01T20:02:07Z",
					"expiresAt": "2018-08-04T20:03:11Z",
					"id": "5b6211ff87d9d663c59d3feg",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/MyCluster/backup/snapshots/5b6211ff87d9d663c59d3feg",
							"rel": "self"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/MyCluster",
							"rel": "http://mms.mongodb.com/cluster"
						}
					],
					"mongodVersion": "3.6.6",
					"replicaSetName": "newCluster",
					"snapshotType": "scheduled",
					"storageSizeBytes": 1778012160,
					"type": "replicaSet"
				}
			],
			"totalCount": 1
		}`)
	})

	cloudProviderSnapshots, _, err := client.CloudProviderSnapshots.GetAllCloudProviderSnapshots(ctx, requestParameters, nil)
	if err != nil {
		t.Fatalf("CloudProviderSnapshots.GetAllCloudProviderSnapshots returned error: %v", err)
	}

	expected := &CloudProviderSnapshots{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/MyCluster/backup/snapshots?pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*CloudProviderSnapshot{
			{
				CloudProvider: "AZURE",
				CreatedAt:     "2018-08-01T20:02:07Z",
				ExpiresAt:     "2018-08-04T20:03:11Z",
				ID:            "5b6211ff87d9d663c59d3feg",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/MyCluster/backup/snapshots/5b6211ff87d9d663c59d3feg",
						Rel:  "self",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/MyCluster",
						Rel:  "http://mms.mongodb.com/cluster",
					},
				},
				MongodVersion:    "3.6.6",
				ReplicaSetName:   "newCluster",
				SnapshotType:     "scheduled",
				StorageSizeBytes: 1778012160,
				Type:             "replicaSet",
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(cloudProviderSnapshots, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshots_GetOneCloudProviderSnapshot(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	requestParameters := &SnapshotReqPathParameters{
		GroupID:     "5b6212af90dc76637950a2c6",
		ClusterName: "MyCluster",
		SnapshotID:  "5b6211ff87d9d663c59d3feg",
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/snapshots/%s", requestParameters.GroupID, requestParameters.ClusterName, requestParameters.SnapshotID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"cloudProvider": "AZURE",
			"createdAt": "2018-08-01T20:02:07Z",
			"description": "SomeDescription",
			"expiresAt": "2018-08-04T20:03:11Z",
			"id": "5b6211ff87d9d663c59d3feg",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/MyCluster/backup/snapshots/5b6211ff87d9d663c59d3feg",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/MyCluster",
					"rel": "http://mms.mongodb.com/cluster"
				}
			],
			"mongodVersion": "4.0.3",
			"replicaSetName": "newCluster",
			"snapshotType": "onDemand",
			"status": "queued",
			"storageSizeBytes": 1778012160,
			"type": "shardedCluster"
		}`)
	})

	cloudProviderSnapshot, _, err := client.CloudProviderSnapshots.GetOneCloudProviderSnapshot(ctx, requestParameters)
	if err != nil {
		t.Fatalf("CloudProviderSnapshots.GetOneCloudProviderSnapshot returned error: %v", err)
	}

	expected := &CloudProviderSnapshot{
		CloudProvider: "AZURE",
		CreatedAt:     "2018-08-01T20:02:07Z",
		Description:   "SomeDescription",
		ExpiresAt:     "2018-08-04T20:03:11Z",
		ID:            "5b6211ff87d9d663c59d3feg",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/MyCluster/backup/snapshots/5b6211ff87d9d663c59d3feg",
				Rel:  "self",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/MyCluster",
				Rel:  "http://mms.mongodb.com/cluster",
			},
		},
		MongodVersion:    "4.0.3",
		ReplicaSetName:   "newCluster",
		SnapshotType:     "onDemand",
		Status:           "queued",
		StorageSizeBytes: 1778012160,
		Type:             "shardedCluster",
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshots_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	requestParameters := &SnapshotReqPathParameters{
		GroupID:     "6d2065c687d9d64ae7acdg41",
		ClusterName: "MyClusterName",
	}

	createRequest := &CloudProviderSnapshot{
		Description:     "SomeDescription",
		RetentionInDays: 5,
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/snapshots", requestParameters.GroupID, requestParameters.ClusterName)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"description":     "SomeDescription",
			"retentionInDays": float64(5),
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
			"cloudProvider": "AZURE",
			"createdAt": "2018-12-31T20:54:03Z",
			"description": "SomeDescription ",
			"expiresAt": "2019-01-05T20:54:03Z",
			"id": "6d3b81eb87d9d61e37598558",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/6d2065c687d9d64ae7acdg41/clusters/Cluster0/backup/snapshots/6d3b81eb87d9d61e37598558",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/6d2065c687d9d64ae7acdg41/clusters/Cluster0",
					"rel": "http://mms.mongodb.com/cluster"
				}
			],
			"mongodVersion": "4.0.4",
			"replicaSetName": "newCluster",
			"snapshotType": "onDemand",
			"status": "queued",
			"storageSizeBytes": 0,
			"type": "replicaSet"
		}`)
	})

	cloudProviderSnapshot, _, err := client.CloudProviderSnapshots.Create(ctx, requestParameters, createRequest)
	if err != nil {
		t.Fatalf("CloudProviderSnapshots.Create returned error: %v", err)
	}

	expected := &CloudProviderSnapshot{
		CloudProvider: "AZURE",
		CreatedAt:     "2018-12-31T20:54:03Z",
		Description:   "SomeDescription ",
		ExpiresAt:     "2019-01-05T20:54:03Z",
		ID:            "6d3b81eb87d9d61e37598558",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6d2065c687d9d64ae7acdg41/clusters/Cluster0/backup/snapshots/6d3b81eb87d9d61e37598558",
				Rel:  "self",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6d2065c687d9d64ae7acdg41/clusters/Cluster0",
				Rel:  "http://mms.mongodb.com/cluster",
			},
		},
		MongodVersion:    "4.0.4",
		ReplicaSetName:   "newCluster",
		SnapshotType:     "onDemand",
		Status:           "queued",
		StorageSizeBytes: 0,
		Type:             "replicaSet",
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshots_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	requestParameters := &SnapshotReqPathParameters{
		GroupID:     "5b6212af90dc76637950a2c6",
		ClusterName: "MyCluster",
		SnapshotID:  "5b6211ff87d9d663c59d3feg",
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/snapshots/%s", requestParameters.GroupID, requestParameters.ClusterName, requestParameters.SnapshotID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.CloudProviderSnapshots.Delete(ctx, requestParameters)
	if err != nil {
		t.Fatalf("CloudProviderSnapshots.Delete returned error: %v", err)
	}
}
