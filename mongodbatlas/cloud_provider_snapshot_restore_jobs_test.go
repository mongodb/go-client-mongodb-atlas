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

func TestCloudProviderSnapshotRestoreJobs_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	requestParameters := &SnapshotReqPathParameters{
		GroupID:     "5b6212af90dc76637950a2c6",
		ClusterName: "MyCluster",
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/restoreJobs", requestParameters.GroupID, requestParameters.ClusterName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs?pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				{
					"cancelled": false,
					"deliveryType": "automated",
					"expired": false,
					"expiresAt": "2018-08-02T02:08:48Z",
					"id": "5b622f7087d9d6039fafe03f",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622f7087d9d6039fafe03f",
							"rel": "self"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
							"rel": "http://mms.mongodb.com/snapshot"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
							"rel": "http://mms.mongodb.com/cluster"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
							"rel": "http://mms.mongodb.com/group"
						}
					],
					"snapshotId": "5b6211ff87d9d663c59d3feg",
					"targetClusterName": "MyOtherCluster",
					"targetGroupId": "5b6212af90dc76637950a2c6",
					"timestamp": "2018-08-01T20:02:07Z"
				},
				{
					"cancelled": false,
					"createdAt": "2018-08-01T22:05:41Z",
					"components": [
						{
							"downloadUrl": "https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f-config.tar.gz",
							"replicaSetName": "atlas-abcdef-config-0"
						},
						{
							"downloadUrl": "https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f.tar.gz",
							"replicaSetName": "atlas-abcdef-shard-0"
						}
					],
					"deliveryType": "download",
					"deliveryUrl": ["https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f.tar.gz"],
					"expired": false,
					"expiresAt": "2018-08-02T02:03:33Z",
					"id": "5b622e3587d9d6039faf713f",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622e3587d9d6039faf713f",
							"rel": "self"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
							"rel": "http://mms.mongodb.com/snapshot"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
							"rel": "http://mms.mongodb.com/cluster"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
							"rel": "http://mms.mongodb.com/group"
						}
					],
					"snapshotId": "5b6211ff87d9d663c59d3feg",
					"timestamp": "2018-08-01T20:02:07Z",
					"oplogTs":	1583753751,
					"oplogInc":	1
				}

			],
			"totalCount": 2
		}`)
	})

	cloudProviderSnapshots, _, err := client.CloudProviderSnapshotRestoreJobs.List(ctx, requestParameters, nil)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotRestoreJobs.List returned error: %v", err)
	}

	expected := &CloudProviderSnapshotRestoreJobs{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs?pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*CloudProviderSnapshotRestoreJob{
			{
				Cancelled:    false,
				DeliveryType: "automated",
				Expired:      false,
				ExpiresAt:    "2018-08-02T02:08:48Z",
				ID:           "5b622f7087d9d6039fafe03f",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622f7087d9d6039fafe03f",
						Rel:  "self",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
						Rel:  "http://mms.mongodb.com/snapshot",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
						Rel:  "http://mms.mongodb.com/cluster",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
						Rel:  "http://mms.mongodb.com/group",
					},
				},
				SnapshotID:        "5b6211ff87d9d663c59d3feg",
				TargetClusterName: "MyOtherCluster",
				TargetGroupID:     "5b6212af90dc76637950a2c6",
				Timestamp:         "2018-08-01T20:02:07Z",
			},
			{
				Cancelled: false,
				CreatedAt: "2018-08-01T22:05:41Z",
				Components: []*Component{
					{
						DownloadURL:    "https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f-config.tar.gz",
						ReplicaSetName: "atlas-abcdef-config-0",
					},
					{
						DownloadURL:    "https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f.tar.gz",
						ReplicaSetName: "atlas-abcdef-shard-0",
					},
				},
				DeliveryType: "download",
				DeliveryURL:  []string{"https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f.tar.gz"},
				Expired:      false,
				ExpiresAt:    "2018-08-02T02:03:33Z",
				ID:           "5b622e3587d9d6039faf713f",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622e3587d9d6039faf713f",
						Rel:  "self",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
						Rel:  "http://mms.mongodb.com/snapshot",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
						Rel:  "http://mms.mongodb.com/cluster",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
						Rel:  "http://mms.mongodb.com/group",
					},
				},
				SnapshotID: "5b6211ff87d9d663c59d3feg",
				Timestamp:  "2018-08-01T20:02:07Z",
				OplogTs:    1583753751,
				OplogInc:   1,
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(cloudProviderSnapshots, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotRestoreJobs_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	requestParameters := &SnapshotReqPathParameters{
		GroupID:     "5b6212af90dc76637950a2c6",
		ClusterName: "MyCluster",
		JobID:       "5b622f7087d9d6039fafe03f",
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/restoreJobs/%s", requestParameters.GroupID, requestParameters.ClusterName, requestParameters.JobID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
  "cancelled": false,
  "deliveryType": "automated",
  "expired": false,
  "expiresAt": "2018-08-02T02:08:48Z",
  "id": "{RESTORE-JOB-ID}",
  "links": [
    {
      "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/clusters/{CLUSTER-NAME}/backup/restoreJobs/{RESTORE-JOB-ID}",
      "rel": "self"
    },
    {
      "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/clusters/{CLUSTER-NAME}/backup/snapshots/{SNAPSHOT-ID}",
      "rel": "http://mms.mongodb.com/snapshot"
    },
    {
      "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/clusters/{CLUSTER-NAME}",
      "rel": "http://mms.mongodb.com/cluster"
    },
    {
      "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}",
      "rel": "http://mms.mongodb.com/group"
    }
  ],
  "snapshotId": "{SNAPSHOT-ID}",
  "targetClusterName": "MyOtherCluster",
  "targetGroupId": "{GROUP-ID}",
  "timestamp": "2018-08-01T20:02:07Z"
}`)
	})

	cloudProviderSnapshot, _, err := client.CloudProviderSnapshotRestoreJobs.Get(ctx, requestParameters)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotRestoreJobs.Get returned error: %v", err)
	}

	expected := &CloudProviderSnapshotRestoreJob{
		Cancelled:    false,
		DeliveryType: "automated",
		Expired:      false,
		ExpiresAt:    "2018-08-02T02:08:48Z",
		ID:           "{RESTORE-JOB-ID}",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/clusters/{CLUSTER-NAME}/backup/restoreJobs/{RESTORE-JOB-ID}",
				Rel:  "self",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/clusters/{CLUSTER-NAME}/backup/snapshots/{SNAPSHOT-ID}",
				Rel:  "http://mms.mongodb.com/snapshot",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/clusters/{CLUSTER-NAME}",
				Rel:  "http://mms.mongodb.com/cluster",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}",
				Rel:  "http://mms.mongodb.com/group",
			},
		},
		SnapshotID:        "{SNAPSHOT-ID}",
		TargetClusterName: "MyOtherCluster",
		TargetGroupID:     "{GROUP-ID}",
		Timestamp:         "2018-08-01T20:02:07Z",
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotRestoreJobs_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	requestParameters := &SnapshotReqPathParameters{
		GroupID:     "5b6212af90dc76637950a2c6",
		ClusterName: "MyClusterName",
	}

	createRequest := &CloudProviderSnapshotRestoreJob{
		SnapshotID:        "5b6211ff87d9d663c59d3feg",
		DeliveryType:      "automated",
		TargetClusterName: "MyOtherCluster",
		TargetGroupID:     "5b6212af90dc76637950a2c6",
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/restoreJobs", requestParameters.GroupID, requestParameters.ClusterName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"snapshotId":        "5b6211ff87d9d663c59d3feg",
			"deliveryType":      "automated",
			"targetClusterName": "MyOtherCluster",
			"targetGroupId":     "5b6212af90dc76637950a2c6",
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
			"cancelled": false,
			"deliveryType": "automated",
			"expired": false,
			"expiresAt": "2018-08-02T02:08:48Z",
			"id": "5b622f7087d9d6039fafe03f",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622f7087d9d6039fafe03f",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
					"rel": "http://mms.mongodb.com/snapshot"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
					"rel": "http://mms.mongodb.com/cluster"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
					"rel": "http://mms.mongodb.com/group"
				}
			],
			"snapshotId": "5b6211ff87d9d663c59d3feg",
			"targetClusterName": "MyOtherCluster",
			"targetGroupId": "5b6212af90dc76637950a2c6",
			"timestamp": "2018-08-01T20:02:07Z",
			"oplogTs":	1583753751,
			"oplogInc":	1
		}`)
	})

	cloudProviderSnapshot, _, err := client.CloudProviderSnapshotRestoreJobs.Create(ctx, requestParameters, createRequest)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotRestoreJobs.Create returned error: %v", err)
	}

	expected := &CloudProviderSnapshotRestoreJob{
		Cancelled:    false,
		DeliveryType: "automated",
		Expired:      false,
		ExpiresAt:    "2018-08-02T02:08:48Z",
		ID:           "5b622f7087d9d6039fafe03f",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622f7087d9d6039fafe03f",
				Rel:  "self",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
				Rel:  "http://mms.mongodb.com/snapshot",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
				Rel:  "http://mms.mongodb.com/cluster",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
				Rel:  "http://mms.mongodb.com/group",
			},
		},
		SnapshotID:        "5b6211ff87d9d663c59d3feg",
		TargetClusterName: "MyOtherCluster",
		TargetGroupID:     "5b6212af90dc76637950a2c6",
		Timestamp:         "2018-08-01T20:02:07Z",
		OplogTs:           1583753751,
		OplogInc:          1,
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotRestoreJobs_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	requestParameters := &SnapshotReqPathParameters{
		GroupID:     "5b6212af90dc76637950a2c6",
		ClusterName: "MyCluster",
		JobID:       "5b622f7087d9d6039fafe03f",
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/restoreJobs/%s", requestParameters.GroupID, requestParameters.ClusterName, requestParameters.JobID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.CloudProviderSnapshotRestoreJobs.Delete(ctx, requestParameters)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotRestoreJobs.Delete returned error: %v", err)
	}
}

func TestCloudProviderSnapshotRestoreJobs_ServerlessList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID, instanceName := "5b6212af90dc76637950a2c6", "MyCluster"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/serverless/%s/backup/restoreJobs", projectID, instanceName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs?pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				{
					"cancelled": false,
					"deliveryType": "automated",
					"expired": false,
					"expiresAt": "2018-08-02T02:08:48Z",
					"id": "5b622f7087d9d6039fafe03f",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622f7087d9d6039fafe03f",
							"rel": "self"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
							"rel": "http://mms.mongodb.com/snapshot"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
							"rel": "http://mms.mongodb.com/cluster"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
							"rel": "http://mms.mongodb.com/group"
						}
					],
					"snapshotId": "5b6211ff87d9d663c59d3feg",
					"targetClusterName": "MyOtherCluster",
					"targetGroupId": "5b6212af90dc76637950a2c6",
					"timestamp": "2018-08-01T20:02:07Z"
				},
				{
					"cancelled": false,
					"createdAt": "2018-08-01T22:05:41Z",
					"components": [
						{
							"downloadUrl": "https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f-config.tar.gz",
							"replicaSetName": "atlas-abcdef-config-0"
						},
						{
							"downloadUrl": "https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f.tar.gz",
							"replicaSetName": "atlas-abcdef-shard-0"
						}
					],
					"deliveryType": "download",
					"deliveryUrl": ["https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f.tar.gz"],
					"expired": false,
					"expiresAt": "2018-08-02T02:03:33Z",
					"id": "5b622e3587d9d6039faf713f",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622e3587d9d6039faf713f",
							"rel": "self"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
							"rel": "http://mms.mongodb.com/snapshot"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
							"rel": "http://mms.mongodb.com/cluster"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
							"rel": "http://mms.mongodb.com/group"
						}
					],
					"snapshotId": "5b6211ff87d9d663c59d3feg",
					"timestamp": "2018-08-01T20:02:07Z",
					"oplogTs":	1583753751,
					"oplogInc":	1
				}

			],
			"totalCount": 2
		}`)
	})

	cloudProviderSnapshots, _, err := client.CloudProviderSnapshotRestoreJobs.ListForServerlessBackupRestore(ctx, projectID, instanceName, nil)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotRestoreJobs.List returned error: %v", err)
	}

	expected := &CloudProviderSnapshotRestoreJobs{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs?pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*CloudProviderSnapshotRestoreJob{
			{
				Cancelled:    false,
				DeliveryType: "automated",
				Expired:      false,
				ExpiresAt:    "2018-08-02T02:08:48Z",
				ID:           "5b622f7087d9d6039fafe03f",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622f7087d9d6039fafe03f",
						Rel:  "self",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
						Rel:  "http://mms.mongodb.com/snapshot",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
						Rel:  "http://mms.mongodb.com/cluster",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
						Rel:  "http://mms.mongodb.com/group",
					},
				},
				SnapshotID:        "5b6211ff87d9d663c59d3feg",
				TargetClusterName: "MyOtherCluster",
				TargetGroupID:     "5b6212af90dc76637950a2c6",
				Timestamp:         "2018-08-01T20:02:07Z",
			},
			{
				Cancelled: false,
				CreatedAt: "2018-08-01T22:05:41Z",
				Components: []*Component{
					{
						DownloadURL:    "https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f-config.tar.gz",
						ReplicaSetName: "atlas-abcdef-config-0",
					},
					{
						DownloadURL:    "https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f.tar.gz",
						ReplicaSetName: "atlas-abcdef-shard-0",
					},
				},
				DeliveryType: "download",
				DeliveryURL:  []string{"https://restore.mongodb.net:27017/restore-5b622e3587d9d6039faf713f.tar.gz"},
				Expired:      false,
				ExpiresAt:    "2018-08-02T02:03:33Z",
				ID:           "5b622e3587d9d6039faf713f",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622e3587d9d6039faf713f",
						Rel:  "self",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
						Rel:  "http://mms.mongodb.com/snapshot",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
						Rel:  "http://mms.mongodb.com/cluster",
					},
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
						Rel:  "http://mms.mongodb.com/group",
					},
				},
				SnapshotID: "5b6211ff87d9d663c59d3feg",
				Timestamp:  "2018-08-01T20:02:07Z",
				OplogTs:    1583753751,
				OplogInc:   1,
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(cloudProviderSnapshots, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotRestoreJobs_ServerlessGet(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID, instanceName, jobID := "5b6212af90dc76637950a2c6", "MyCluster", "5b622f7087d9d6039fafe03f"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/serverless/%s/backup/restoreJobs/%s", projectID, instanceName, jobID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"cancelled": false,
			"deliveryType": "automated",
			"expired": false,
			"expiresAt": "2018-08-02T02:08:48Z",
			"id": "5b622f7087d9d6039fafe03f",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622f7087d9d6039fafe03f",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
					"rel": "http://mms.mongodb.com/snapshot"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
					"rel": "http://mms.mongodb.com/cluster"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
					"rel": "http://mms.mongodb.com/group"
				}
			],
			"snapshotId": "5b6211ff87d9d663c59d3feg",
			"targetClusterName": "MyOtherCluster",
			"targetGroupId": "5b6212af90dc76637950a2c6",
			"timestamp": "2018-08-01T20:02:07Z",
			"oplogTs":	1583753751,
			"oplogInc":	1
		}`)
	})

	cloudProviderSnapshot, _, err := client.CloudProviderSnapshotRestoreJobs.GetForServerlessBackupRestore(ctx, projectID, instanceName, jobID)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotRestoreJobs.Get returned error: %v", err)
	}

	expected := &CloudProviderSnapshotRestoreJob{
		Cancelled:    false,
		DeliveryType: "automated",
		Expired:      false,
		ExpiresAt:    "2018-08-02T02:08:48Z",
		ID:           "5b622f7087d9d6039fafe03f",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622f7087d9d6039fafe03f",
				Rel:  "self",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
				Rel:  "http://mms.mongodb.com/snapshot",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
				Rel:  "http://mms.mongodb.com/cluster",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
				Rel:  "http://mms.mongodb.com/group",
			},
		},
		SnapshotID:        "5b6211ff87d9d663c59d3feg",
		TargetClusterName: "MyOtherCluster",
		TargetGroupID:     "5b6212af90dc76637950a2c6",
		Timestamp:         "2018-08-01T20:02:07Z",
		OplogTs:           1583753751,
		OplogInc:          1,
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotRestoreJobs_ServerlessCreate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID, instanceName := "5b6212af90dc76637950a2c6", "MyClusterName"

	createRequest := &CloudProviderSnapshotRestoreJob{
		SnapshotID:        "5b6211ff87d9d663c59d3feg",
		DeliveryType:      "automated",
		TargetClusterName: "MyOtherCluster",
		TargetGroupID:     "5b6212af90dc76637950a2c6",
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/serverless/%s/backup/restoreJobs", projectID, instanceName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"snapshotId":        "5b6211ff87d9d663c59d3feg",
			"deliveryType":      "automated",
			"targetClusterName": "MyOtherCluster",
			"targetGroupId":     "5b6212af90dc76637950a2c6",
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
			"cancelled": false,
			"deliveryType": "automated",
			"expired": false,
			"expiresAt": "2018-08-02T02:08:48Z",
			"id": "5b622f7087d9d6039fafe03f",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622f7087d9d6039fafe03f",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
					"rel": "http://mms.mongodb.com/snapshot"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
					"rel": "http://mms.mongodb.com/cluster"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
					"rel": "http://mms.mongodb.com/group"
				}
			],
			"snapshotId": "5b6211ff87d9d663c59d3feg",
			"targetClusterName": "MyOtherCluster",
			"targetGroupId": "5b6212af90dc76637950a2c6",
			"timestamp": "2018-08-01T20:02:07Z",
			"oplogTs":	1583753751,
			"oplogInc":	1
		}`)
	})

	cloudProviderSnapshot, _, err := client.CloudProviderSnapshotRestoreJobs.CreateForServerlessBackupRestore(ctx, projectID, instanceName, createRequest)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotRestoreJobs.Create returned error: %v", err)
	}

	expected := &CloudProviderSnapshotRestoreJob{
		Cancelled:    false,
		DeliveryType: "automated",
		Expired:      false,
		ExpiresAt:    "2018-08-02T02:08:48Z",
		ID:           "5b622f7087d9d6039fafe03f",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/restoreJobs/5b622f7087d9d6039fafe03f",
				Rel:  "self",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0/backup/snapshots/5b6211ff87d9d663c59d3dee",
				Rel:  "http://mms.mongodb.com/snapshot",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6/clusters/Cluster0",
				Rel:  "http://mms.mongodb.com/cluster",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b6212af90dc76637950a2c6",
				Rel:  "http://mms.mongodb.com/group",
			},
		},
		SnapshotID:        "5b6211ff87d9d663c59d3feg",
		TargetClusterName: "MyOtherCluster",
		TargetGroupID:     "5b6212af90dc76637950a2c6",
		Timestamp:         "2018-08-01T20:02:07Z",
		OplogTs:           1583753751,
		OplogInc:          1,
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}
