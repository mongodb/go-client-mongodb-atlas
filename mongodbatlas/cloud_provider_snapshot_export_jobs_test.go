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
)

func TestCloudProviderSnapshotExportJobs_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/exports", groupID, clusterName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
  "links": [
    {
      "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/clusters/{CLUSTER-NAME}/backup/exports/?pageNum=1&itemsPerPage=100",
      "rel": "self"
    }
  ],
  "results": [
    {
      "components": [
        {
          "exportId": "60675896c37241327a557130",
          "replicaSetName": "atlas-xxxxxx-shard-0"
        },
        {
          "exportId": "60675896c37241327a557132",
          "replicaSetName": "atlas-wnvwfd-shard-1"
        }
      ],
      "createdAt": "2021-04-02T17:47:02Z",
      "customData": [],
      "exportBucketId": "{BUCKET-ID}",
      "id": "60675896c37241327a55712f",
      "prefix": "/exported_snapshots/{ORG-NAME}/{PROJECT-NAME}/{CLUSTER-NAME}/2021-04-02T1742/1617385622",
      "snapshotId": "{SNAPSHOT-ID}",
      "state": "Queued"
    },
    {
      "components": [
        {
          "exportId": "606758180d97065aabbe5b86",
          "replicaSetName": "atlas-xxxxxx-shard-0"
        },
        {
          "exportId": "606758180d97065aabbe5b88",
          "replicaSetName": "atlas-xxxxxx-shard-1"
        }
      ],
      "createdAt": "2021-04-02T17:44:56Z",
      "customData": [],
      "exportBucketId": "{BUCKET-ID}",
      "id": "606758180d97065aabbe5b85",
      "prefix": "/exported_snapshots/{ORG-NAME}/{PROJECT-NAME}/{CLUSTER-NAME}/2021-04-02T1742/1617385496",
      "snapshotId": "{SNAPSHOT-ID}",
      "state": "Successful"
    }
  ],
  "totalCount": 2
}`)
	})

	cloudProviderSnapshots, _, err := client.CloudProviderSnapshotExportJobs.List(ctx, groupID, clusterName, nil)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotExportJobs.List returned error: %v", err)
	}

	expected := &CloudProviderSnapshotExportJobs{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/clusters/{CLUSTER-NAME}/backup/exports/?pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*CloudProviderSnapshotExportJob{
			{
				Components: []*CloudProviderSnapshotExportJobComponent{
					{
						ExportID:       "60675896c37241327a557130",
						ReplicaSetName: "atlas-xxxxxx-shard-0",
					},
					{
						ExportID:       "60675896c37241327a557132",
						ReplicaSetName: "atlas-wnvwfd-shard-1",
					},
				},
				CreatedAt:      "2021-04-02T17:47:02Z",
				CustomData:     []*CloudProviderSnapshotExportJobCustomData{},
				ExportBucketID: "{BUCKET-ID}",
				ID:             "60675896c37241327a55712f",
				Prefix:         "/exported_snapshots/{ORG-NAME}/{PROJECT-NAME}/{CLUSTER-NAME}/2021-04-02T1742/1617385622",
				SnapshotID:     "{SNAPSHOT-ID}",
				State:          "Queued",
			},
			{
				Components: []*CloudProviderSnapshotExportJobComponent{
					{
						ExportID:       "606758180d97065aabbe5b86",
						ReplicaSetName: "atlas-xxxxxx-shard-0",
					},
					{
						ExportID:       "606758180d97065aabbe5b88",
						ReplicaSetName: "atlas-xxxxxx-shard-1",
					},
				},
				CreatedAt:      "2021-04-02T17:44:56Z",
				CustomData:     []*CloudProviderSnapshotExportJobCustomData{},
				ExportBucketID: "{BUCKET-ID}",
				ID:             "606758180d97065aabbe5b85",
				Prefix:         "/exported_snapshots/{ORG-NAME}/{PROJECT-NAME}/{CLUSTER-NAME}/2021-04-02T1742/1617385496",
				SnapshotID:     "{SNAPSHOT-ID}",
				State:          "Successful",
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(cloudProviderSnapshots, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotExportJobs_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	exportID := "job-id-test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/exports/%s", groupID, clusterName, exportID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
  "createdAt": "2021-03-25T16:49:23Z",
  "exportBucketId": "{BUCKET-ID}",
  "finishedAt": "2021-03-25T17:48:20Z",
  "id": "605ccba3c1c7613f423e1585",
  "prefix": "/exported_snapshots/{ORG-NAME}/{PROJECT-NAME}/{CLUSTER-NAME}/2021-03-25T1649+0000/1616694179",
  "snapshotId": "605cbac9b498841007af7d9f",
  "state": "Successful"
}`)
	})

	cloudProviderSnapshotBucket, _, err := client.CloudProviderSnapshotExportJobs.Get(ctx, groupID, clusterName, exportID)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotExportJobs.Get returned error: %v", err)
	}

	expected := &CloudProviderSnapshotExportJob{
		CreatedAt:      "2021-03-25T16:49:23Z",
		ExportBucketID: "{BUCKET-ID}",
		FinishedAt:     "2021-03-25T17:48:20Z",
		ID:             "605ccba3c1c7613f423e1585",
		Prefix:         "/exported_snapshots/{ORG-NAME}/{PROJECT-NAME}/{CLUSTER-NAME}/2021-03-25T1649+0000/1616694179",
		SnapshotID:     "605cbac9b498841007af7d9f",
		State:          "Successful",
	}

	if diff := deep.Equal(cloudProviderSnapshotBucket, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotExportJobs_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &CloudProviderSnapshotExportJob{
		SnapshotID:     "{SNAPSHOT-ID}",
		ExportBucketID: "{BUCKET-ID}",
		CustomData: []*CloudProviderSnapshotExportJobCustomData{
			{
				Key:   "exported by",
				Value: "myName",
			},
		},
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/backup/exports", groupID, clusterName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
  "createdAt": "2021-05-04T14:43:40Z",
  "customData": [
    {
      "key": "exported by",
      "value": "myName"
    }
  ],
  "exportBucketId": "604f6322dc786a5341d4f7fb",
  "exportStatus": {
    "exportedCollections": 0,
    "totalCollections": 0
  },
  "id": "60915d9cef8ca80cd31646e5",
  "prefix": "exported_snapshots/KS+Sandbox/mySbx/testCluster/2021-03-27T1409/1620139419",
  "snapshotId": "605f3bdf16d634087d0e15ce",
  "state": "Queued"
}`)
	})

	cloudProviderSnapshot, _, err := client.CloudProviderSnapshotExportJobs.Create(ctx, groupID, clusterName, createRequest)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotExportBuckets.Create returned error: %v", err)
	}

	expected := &CloudProviderSnapshotExportJob{
		CreatedAt:      "2021-05-04T14:43:40Z",
		ExportBucketID: "604f6322dc786a5341d4f7fb",
		CustomData: []*CloudProviderSnapshotExportJobCustomData{
			{
				Key:   "exported by",
				Value: "myName",
			},
		},
		ExportStatus: &CloudProviderSnapshotExportJobStatus{
			ExportedCollections: 0,
			TotalCollections:    0,
		},
		ID:         "60915d9cef8ca80cd31646e5",
		Prefix:     "exported_snapshots/KS+Sandbox/mySbx/testCluster/2021-03-27T1409/1620139419",
		SnapshotID: "605f3bdf16d634087d0e15ce",
		State:      "Queued",
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}
