// Copyright 2022 MongoDB Inc
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

func TestRestoreBackupJob_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "1"
	clusterName := "test"
	deploymentType := "DEDICATED"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/restoreJobs", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
			{
  "results": [
    {
      "cancelled": false,
      "deliveryType": "automated",
      "expired": false,
      "expiresAt": "2020-08-02T02:08:48Z",
      "id": "{RESTORE-JOB-ID-1}",
      "snapshotId": "{SNAPSHOT-ID}",
      "targetClusterName": "MyOtherCluster",
      "targetGroupId": "{GROUP-ID}",
      "timestamp": "2020-08-01T20:02:07Z"
    },
    {
      "cancelled": false,
      "createdAt": "2020-08-01T22:05:41Z",
      "deliveryType": "download",
      "deliveryUrl": ["https://restore.example.com:27017/restore-{RESTORE-JOB-ID-2}.tar.gz"],
      "expired": false,
      "expiresAt": "2020-08-02T02:03:33Z",
      "id": "{RESTORE-JOB-ID-2}",
      "snapshotId": "{SNAPSHOT-ID}",
      "timestamp": "2020-08-01T20:02:07Z"
    }
  ],
  "totalCount": 2
}`)
	})

	responseObj, _, err := client.RestoreBackupJob.List(ctx, groupID, clusterName, deploymentType, nil)
	if err != nil {
		t.Fatalf("RestoreBackupJob.List returned error: %v", err)
	}

	expected := &RestoreBackupJobs{
		TotalCount: 2,
		Results: []*RestoreBackupJob{
			{
				Cancelled:         pointy.Bool(false),
				DeliveryType:      "automated",
				Expired:           pointy.Bool(false),
				ExpiresAt:         "2020-08-02T02:08:48Z",
				ID:                "{RESTORE-JOB-ID-1}",
				SnapshotID:        "{SNAPSHOT-ID}",
				TargetClusterName: "MyOtherCluster",
				TargetGroupID:     "{GROUP-ID}",
				Timestamp:         "2020-08-01T20:02:07Z",
			},
			{
				Cancelled:    pointy.Bool(false),
				CreatedAt:    "2020-08-01T22:05:41Z",
				DeliveryType: "download",
				DeliveryURL:  []string{"https://restore.example.com:27017/restore-{RESTORE-JOB-ID-2}.tar.gz"},
				Expired:      pointy.Bool(false),
				ExpiresAt:    "2020-08-02T02:03:33Z",
				ID:           "{RESTORE-JOB-ID-2}",
				SnapshotID:   "{SNAPSHOT-ID}",
				Timestamp:    "2020-08-01T20:02:07Z",
			},
		},
	}

	if diff := deep.Equal(responseObj, expected); diff != nil {
		t.Error(diff)
	}
}

func TestRestoreBackupJob_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "1"
	clusterName := "test"
	deploymentType := "DEDICATED"
	restoreJobID := "jobID1"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/restoreJobs/%s", groupID, clusterName, restoreJobID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
			{
  "cancelled": false,
  "deliveryType": "automated",
  "expired": false,
  "expiresAt": "2018-08-02T02:08:48Z",
  "failed" : false,
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

	responseObj, _, err := client.RestoreBackupJob.Get(ctx, groupID, clusterName, restoreJobID, deploymentType)
	if err != nil {
		t.Fatalf("RestoreBackupJob.Get returned error: %v", err)
	}

	expected := &RestoreBackupJob{
		Cancelled:    pointy.Bool(false),
		DeliveryType: "automated",
		Expired:      pointy.Bool(false),
		ExpiresAt:    "2018-08-02T02:08:48Z",
		Failed:       pointy.Bool(false),
		ID:           "{RESTORE-JOB-ID}",
		Links: []Link{
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
			}},

		SnapshotID:        "{SNAPSHOT-ID}",
		TargetClusterName: "MyOtherCluster",
		TargetGroupID:     "{GROUP-ID}",
		Timestamp:         "2018-08-01T20:02:07Z",
	}

	if diff := deep.Equal(responseObj, expected); diff != nil {
		t.Error(diff)
	}
}

func TestRestoreBackupJob_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "1"
	clusterName := "test"

	mux.HandleFunc(fmt.Sprintf("/"+restoreServerlessBackupJobPath, groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			  "cancelled": false,
			  "deliveryType": "automated",
			  "expired": false,
			  "expiresAt": "2021-04-20T00:00:00Z",
			  "failed": false,
			  "id": "{snapshotId}",
			  "snapshotId": "{snapshotId}",
			  "targetClusterName": "{clusterName}",
			  "targetDeploymentItemName": "{deploymentName}",
			  "targetGroupId": "{targetGroupId}",
			  "timestamp": "2020-12-21T00:00:00Z"
			}`)
	})

	bodyParam := &RestoreBackupJobRequest{
		SnapshotID:        "{snapshotId}",
		DeliveryType:      "automated",
		TargetGroupID:     "{targetGroupId}",
		TargetClusterName: "{clusterName}",
		DeploymentType:    "SERVERLESS",
	}

	serverlessInstance, _, err := client.RestoreBackupJob.Create(ctx, groupID, clusterName, bodyParam)
	if err != nil {
		t.Fatalf("RestoreBackupJob.Create returned error: %v", err)
	}

	expected := &RestoreBackupJob{
		Cancelled:         pointy.Bool(false),
		DeliveryType:      "automated",
		Expired:           pointy.Bool(false),
		ExpiresAt:         "2021-04-20T00:00:00Z",
		Failed:            pointy.Bool(false),
		ID:                "{snapshotId}",
		SnapshotID:        "{snapshotId}",
		TargetClusterName: "{clusterName}",
		TargetGroupID:     "{targetGroupId}",
		Timestamp:         "2020-12-21T00:00:00Z",
	}

	if diff := deep.Equal(serverlessInstance, expected); diff != nil {
		t.Error(diff)
	}
}

func TestRestoreBackupJob_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "1"
	clusterName := "test"
	restoreJobID := "jobrestore1"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/restoreJobs/%s", groupID, clusterName, restoreJobID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.RestoreBackupJob.Delete(ctx, groupID, clusterName, restoreJobID)
	if err != nil {
		t.Fatalf("RestoreBackupJob.Delete returned error: %v", err)
	}
}
