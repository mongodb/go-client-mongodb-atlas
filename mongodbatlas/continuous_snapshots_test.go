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

	"github.com/openlyinc/pointy"

	"github.com/go-test/deep"
)

func TestContinuousSnapshots_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	clusterName := "Cluster0"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/snapshots", groupID, clusterName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				"links": [{
				   "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/clusters/Cluster0/snapshots?pageNum=1&itemsPerPage=100",
				   "rel": "self"
				}],
				"results": [{
				   "clusterId": "7c2487d833e9e75286093696",
				   "complete": true,
				   "created": {
					  "date": "2017-12-26T16:32:16Z",
					  "increment": 1
				   },
				   "doNotDelete": false,
				   "expires": "2018-12-25T16:32:16Z",
				   "groupId": "6c7498dg87d9e6526801572b",
				   "id": "5a4279d4fcc178500596745a",
				   "lastOplogAppliedTimestamp": {
					  "date": "2017-12-26T16:32:15Z",
					  "increment": 1
				   },
				   "links": [{
					  "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/clusters/Cluster0/snapshots/5a4279d4fcc178500596745a",
					  "rel": "self"
				   }],
				   "parts": [{
					  "clusterId": "7c2487d833e9e75286093696",
					  "compressionSetting": "GZIP",
					  "dataSizeBytes": 4502,
					  "encryptionEnabled": false,
					  "fileSizeBytes": 324760,
					  "mongodVersion": "3.6.10",
					  "replicaSetName": "Cluster0-shard-0",
					  "storageSizeBytes": 53248,
					  "typeName": "REPLICA_SET"
				   }]
				}],
				"totalCount": 1
			}`)
	})

	snapshots, _, err := client.ContinuousSnapshots.List(ctx, groupID, clusterName, nil)
	if err != nil {
		t.Fatalf("ContinuousSnapshots.List returned error: %v", err)
	}

	expected := &ContinuousSnapshots{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/clusters/Cluster0/snapshots?pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*ContinuousSnapshot{
			{
				ClusterID: "7c2487d833e9e75286093696",
				Complete:  true,
				Created: &SnapshotTimestamp{
					Date:      "2017-12-26T16:32:16Z",
					Increment: 1,
				},
				DoNotDelete: pointy.Bool(false),
				Expires:     "2018-12-25T16:32:16Z",
				GroupID:     "6c7498dg87d9e6526801572b",
				ID:          "5a4279d4fcc178500596745a",
				LastOplogAppliedTimestamp: &SnapshotTimestamp{
					Date:      "2017-12-26T16:32:15Z",
					Increment: 1,
				},
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/clusters/Cluster0/snapshots/5a4279d4fcc178500596745a",
						Rel:  "self",
					},
				},
				Parts: []*Part{
					{
						ReplicaSetName: "Cluster0-shard-0",
						TypeName:       "REPLICA_SET",
						SnapshotPart: SnapshotPart{
							ClusterID:          "7c2487d833e9e75286093696",
							CompressionSetting: "GZIP",
							DataSizeBytes:      4502,
							EncryptionEnabled:  false,
							FileSizeBytes:      324760,
							MongodVersion:      "3.6.10",
							StorageSizeBytes:   53248},
					},
				},
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(snapshots, expected); diff != nil {
		t.Error(diff)
	}
}

func TestContinuousSnapshots_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	clusterName := "Cluster0"
	snapshotID := "6b5380e6jvn128560506942b"
	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/snapshots/%s", groupID, clusterName, snapshotID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"clusterId": "7c2487d833e9e75286093696",
			"complete": true,
			"created": {
			   "date": "2017-12-26T16:32:16Z",
			   "increment": 1
			},
			"doNotDelete": false,
			"expires": "2018-12-25T16:32:16Z",
			"groupId": "6c7498dg87d9e6526801572b",
			"id": "6b5380e6jvn128560506942b",
			"lastOplogAppliedTimestamp": {
			   "date": "2017-12-26T16:32:15Z",
			   "increment": 1
			},
			"links": [{
			   "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/clusters/Cluster0/snapshots/6b5380e6jvn128560506942b",
			   "rel": "self"
			}],
			"parts": [{
			   "clusterId": "7c2487d833e9e75286093696",
			   "compressionSetting": "GZIP",
			   "dataSizeBytes": 4502,
			   "encryptionEnabled": false,
			   "fileSizeBytes": 324760,
			   "mongodVersion": "3.6.10",
			   "replicaSetName": "Cluster0-shard-0",
			   "storageSizeBytes": 53248,
			   "typeName": "REPLICA_SET"
			}]
		}`)
	})

	cloudProviderSnapshot, _, err := client.ContinuousSnapshots.Get(ctx, groupID, clusterName, snapshotID)
	if err != nil {
		t.Fatalf("ContinuousSnapshots.Get returned error: %v", err)
	}

	expected := &ContinuousSnapshot{
		ClusterID: "7c2487d833e9e75286093696",
		Complete:  true,
		Created: &SnapshotTimestamp{
			Date:      "2017-12-26T16:32:16Z",
			Increment: 1,
		},
		DoNotDelete: pointy.Bool(false),
		Expires:     "2018-12-25T16:32:16Z",
		GroupID:     "6c7498dg87d9e6526801572b",
		ID:          "6b5380e6jvn128560506942b",
		LastOplogAppliedTimestamp: &SnapshotTimestamp{
			Date:      "2017-12-26T16:32:15Z",
			Increment: 1,
		},
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/clusters/Cluster0/snapshots/6b5380e6jvn128560506942b",
				Rel:  "self",
			},
		},
		Parts: []*Part{
			{
				ReplicaSetName: "Cluster0-shard-0",
				TypeName:       "REPLICA_SET",
				SnapshotPart: SnapshotPart{
					ClusterID:          "7c2487d833e9e75286093696",
					CompressionSetting: "GZIP",
					DataSizeBytes:      4502,
					EncryptionEnabled:  false,
					FileSizeBytes:      324760,
					MongodVersion:      "3.6.10",
					StorageSizeBytes:   53248},
			},
		},
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}

func TestContinuousSnapshots_ChangeExpiry(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	clusterName := "Cluster0"
	snapshotID := "6b5380e6jvn128560506942b"

	updateRequest := &ContinuousSnapshot{
		Expires: "2018-12-01",
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/snapshots/%s", groupID, clusterName, snapshotID)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		expected := map[string]interface{}{
			"expires": "2018-12-01",
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
		  "clusterId": "57c2487d833e9e75286093696",
		  "complete": true,
		  "created": {
			"date": "2017-12-26T16:32:16Z",
			"increment": 1
		  },
		  "doNotDelete": false,
		  "expires": "2018-12-01T00:00:00Z",
		  "groupId": "6c7498dg87d9e6526801572b",
		  "id": "6b5380e6jvn128560506942b",
		  "lastOplogAppliedTimestamp": {
			"date": "2017-12-26T16:32:15Z",
			"increment": 1
		  },
		  "links": [
			{
			  "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/clusters/Cluster0/snapshots/6b5380e6jvn128560506942b",
			  "rel": "self"
			}
		  ],
		  "parts": [
			{
			  "clusterId": "57c2487d833e9e75286093696",
			  "compressionSetting": "GZIP",
			  "dataSizeBytes": 4502,
			  "encryptionEnabled": false,
			  "fileSizeBytes": 324760,
			  "mongodVersion": "3.6.10",
			  "replicaSetName": "Cluster0-shard-0",
			  "storageSizeBytes": 53248,
			  "typeName": "REPLICA_SET"
			}
		  ]
		}`)
	})

	cloudProviderSnapshot, _, err := client.ContinuousSnapshots.ChangeExpiry(ctx, groupID, clusterName, snapshotID, updateRequest)
	if err != nil {
		t.Fatalf("ContinuousSnapshots.ChangeExpiry returned error: %v", err)
	}

	expected := &ContinuousSnapshot{
		ClusterID: "57c2487d833e9e75286093696",
		Complete:  true,
		Created: &SnapshotTimestamp{
			Date:      "2017-12-26T16:32:16Z",
			Increment: 1,
		},
		DoNotDelete: pointy.Bool(false),
		Expires:     "2018-12-01T00:00:00Z",
		GroupID:     "6c7498dg87d9e6526801572b",
		ID:          "6b5380e6jvn128560506942b",
		LastOplogAppliedTimestamp: &SnapshotTimestamp{
			Date:      "2017-12-26T16:32:15Z",
			Increment: 1,
		},
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/clusters/Cluster0/snapshots/6b5380e6jvn128560506942b",
				Rel:  "self",
			},
		},
		Parts: []*Part{
			{
				ReplicaSetName: "Cluster0-shard-0",
				TypeName:       "REPLICA_SET",
				SnapshotPart: SnapshotPart{
					ClusterID:          "57c2487d833e9e75286093696",
					CompressionSetting: "GZIP",
					DataSizeBytes:      4502,
					EncryptionEnabled:  false,
					FileSizeBytes:      324760,
					MongodVersion:      "3.6.10",
					StorageSizeBytes:   53248,
				}},
		},
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}

func TestContinuousSnapshots_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	clusterName := "Cluster0"
	snapshotID := "6b5380e6jvn128560506942b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/snapshots/%s", groupID, clusterName, snapshotID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.ContinuousSnapshots.Delete(ctx, groupID, clusterName, snapshotID)
	if err != nil {
		t.Fatalf("ContinuousSnapshots.Delete returned error: %v", err)
	}
}
