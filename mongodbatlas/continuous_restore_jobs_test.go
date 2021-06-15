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

func TestContinuousBackupRestore_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups/6b77777887d9d61443b41645/clusters/Cluster0/restoreJobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
					{
					  "links": [
						{
						  "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/6b77777887d9d61443b41645/clusters/Cluster0/restoreJobs?pageNum=1&itemsPerPage=100",
						  "rel": "self"
						}
					  ],
					  "results": [
						{
						  "batchId": "5a66783b80eef5354c77ee13",
						  "clusterId": "7c88887880eef52e5f4d0e2d",
						  "clusterName": "Cluster0",
						  "created": "2018-01-22T23:48:11Z",
						  "delivery": {
							"expirationHours": 48,
							"expires": "2018-01-22T23:49:38Z",
							"maxDownloads": 2147483647,
							"methodName": "HTTP",
							"statusName": "EXPIRED",
							"url": "https://api-backup.mongodb.com/backup/restore/v2/pull/5a66783b80eef5354c77ee16/MGViYTUwZGQ4YWVkNDY4MGE2YWE4NGQzODY0MzAzYTU=/config-Cluster0-config-0-5a66689487d9d61443b46149-1516661094-5a66783b80eef5354c77ee16.tar.gz"
						  },
						  "encryptionEnabled": false,
						  "groupId": "6b77777887d9d61443b41645",
						  "hashes": [
							{
							  "fileName": "config-Cluster0-config-0-5a66689487d9d61443b46149-1516661094-5a66783b80eef5354c77ee16.tar.gz",
							  "hash": "a98af3c1f85a9eb3061423cda0fad8b4d0a48209",
							  "typeName": "SHA1"
							}
						  ],
						  "id": "5a66783b80eef5354c77ee16",
						  "links": [
							{
							  "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/6b77777887d9d61443b41645/clusters/Cluster0/restoreJobs/5a66783b80eef5354c77ee16",
							  "rel": "self"
							}
						  ],
						  "pointInTime": false,
						  "snapshotId": "5a6669d5e2bfb3461861360c",
						  "statusName": "FINISHED",
						  "timestamp": {
							"date": "2018-01-22T22:44:54Z",
							"increment": 1
						  }
						}
					  ],
					  "totalCount": 1
					}`,
		)
	})

	customDBRoles, _, err := client.ContinuousRestoreJobs.List(ctx, "6b77777887d9d61443b41645", "Cluster0", nil)
	if err != nil {
		t.Fatalf("ContinuousRestoreJobs.List returned error: %v", err)
	}
	pointInTimeValue := false

	expected := &ContinuousJobs{
		Results: []*ContinuousJob{{
			BatchID:     "5a66783b80eef5354c77ee13",
			ClusterID:   "7c88887880eef52e5f4d0e2d",
			ClusterName: "Cluster0",
			Created:     "2018-01-22T23:48:11Z",
			Delivery: &Delivery{
				Expires:         "2018-01-22T23:49:38Z",
				ExpirationHours: 48,
				MaxDownloads:    2147483647,
				MethodName:      "HTTP",
				StatusName:      "EXPIRED",
				URL:             "https://api-backup.mongodb.com/backup/restore/v2/pull/5a66783b80eef5354c77ee16/MGViYTUwZGQ4YWVkNDY4MGE2YWE4NGQzODY0MzAzYTU=/config-Cluster0-config-0-5a66689487d9d61443b46149-1516661094-5a66783b80eef5354c77ee16.tar.gz",
			},
			EncryptionEnabled: false,
			GroupID:           "6b77777887d9d61443b41645",
			Hashes: []*Hash{
				{
					TypeName: "SHA1",
					FileName: "config-Cluster0-config-0-5a66689487d9d61443b46149-1516661094-5a66783b80eef5354c77ee16.tar.gz",
					Hash:     "a98af3c1f85a9eb3061423cda0fad8b4d0a48209",
				},
			},
			ID: "5a66783b80eef5354c77ee16",
			Links: []*Link{
				{
					Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6b77777887d9d61443b41645/clusters/Cluster0/restoreJobs/5a66783b80eef5354c77ee16",
					Rel:  "self",
				},
			},
			PointInTime: &pointInTimeValue,
			SnapshotID:  "5a6669d5e2bfb3461861360c",
			StatusName:  "FINISHED",
			Timestamp: SnapshotTimestamp{
				Date:      "2018-01-22T22:44:54Z",
				Increment: 1,
			},
		}},
		TotalCount: 1,
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6b77777887d9d61443b41645/clusters/Cluster0/restoreJobs?pageNum=1&itemsPerPage=100",
			},
		},
	}

	if diff := deep.Equal(customDBRoles, expected); diff != nil {
		t.Error(diff)
	}
}

func TestContinuousBackupRestore_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups/5a66666887d9d61443b41645/clusters/Cluster0/restoreJobs/5a6669d9fcc178211a0d86b9", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
					  "batchId": "5a66783b80eef5354c77ee13",
					  "clusterId": "5a66689487d9d61443b46149",
					  "clusterName": "Cluster0",
					  "created": "2018-01-22T23:48:11Z",
					  "delivery": {
						"expirationHours": 48,
						"expires": "2018-01-22T23:49:38Z",
						"maxDownloads": 2147483647,
						"methodName": "HTTP",
						"statusName": "EXPIRED",
						"url": "https://api-backup.mongodb.com/backup/restore/v2/pull/6b77893b80eef5354c77ee15/N2Y1NDhkMjg0Mzk4NGU1Mzk3NTkwMjA0M2ZhODVkNDk=/Cluster0-shard-1-1516661094-6b77893b80eef5354c77ee15.tar.gz"
					  },
					  "encryptionEnabled": false,
					  "groupId": "5a66666887d9d61443b41645",
					  "hashes": [
						{
						  "fileName": "Cluster0-shard-1-1516661094-6b77893b80eef5354c77ee15.tar.gz",
						  "hash": "86bc2f505c0874cdc0eaaa82ead2ef48aaf56d67",
						  "typeName": "SHA1"
						}
					  ],
					  "id": "6b77893b80eef5354c77ee15",
					  "links": [
						{
						  "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a66666887d9d61443b41645/clusters/Cluster0/restoreJobs/6b77893b80eef5354c77ee15",
						  "rel": "self"
						}
					  ],
					  "pointInTime": false,
					  "snapshotId": "5a6669d9fcc178211a0d86b9",
					  "statusName": "FINISHED",
					  "timestamp": {
						"date": "2018-01-22T22:44:54Z",
						"increment": 1
					  }
					}`,
		)
	})

	customDBRoles, _, err := client.ContinuousRestoreJobs.Get(ctx, "5a66666887d9d61443b41645", "Cluster0", "5a6669d9fcc178211a0d86b9")
	if err != nil {
		t.Fatalf("ContinuousRestoreJobs.Get returned error: %v", err)
	}
	pointInTimeValue := false

	expected := &ContinuousJob{
		BatchID:     "5a66783b80eef5354c77ee13",
		ClusterID:   "5a66689487d9d61443b46149",
		ClusterName: "Cluster0",
		Created:     "2018-01-22T23:48:11Z",
		Delivery: &Delivery{
			Expires:         "2018-01-22T23:49:38Z",
			ExpirationHours: 48,
			MaxDownloads:    2147483647,
			MethodName:      "HTTP",
			StatusName:      "EXPIRED",
			URL:             "https://api-backup.mongodb.com/backup/restore/v2/pull/6b77893b80eef5354c77ee15/N2Y1NDhkMjg0Mzk4NGU1Mzk3NTkwMjA0M2ZhODVkNDk=/Cluster0-shard-1-1516661094-6b77893b80eef5354c77ee15.tar.gz",
		},
		EncryptionEnabled: false,
		GroupID:           "5a66666887d9d61443b41645",
		Hashes: []*Hash{
			{
				TypeName: "SHA1",
				FileName: "Cluster0-shard-1-1516661094-6b77893b80eef5354c77ee15.tar.gz",
				Hash:     "86bc2f505c0874cdc0eaaa82ead2ef48aaf56d67",
			},
		},
		ID: "6b77893b80eef5354c77ee15",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a66666887d9d61443b41645/clusters/Cluster0/restoreJobs/6b77893b80eef5354c77ee15",
				Rel:  "self",
			},
		},
		PointInTime: &pointInTimeValue,
		SnapshotID:  "5a6669d9fcc178211a0d86b9",
		StatusName:  "FINISHED",
		Timestamp: SnapshotTimestamp{
			Date:      "2018-01-22T22:44:54Z",
			Increment: 1,
		},
	}

	if diff := deep.Equal(customDBRoles, expected); diff != nil {
		t.Error(diff)
	}
}

func TestContinuousBackupRestore_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups/5a66666887d9d61443b41645/clusters/Cluster0/restoreJobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		var pointInTimeUTCMillis float64 = 1536610288000
		expected := map[string]interface{}{
			"pointInTimeUTCMillis": pointInTimeUTCMillis,
			"delivery": map[string]interface{}{
				"methodName":        "AUTOMATED_RESTORE",
				"targetGroupId":     "5a66666887d9d61443b41645",
				"targetClusterName": "Cluster0",
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
		  "links" : [ {
			"href" : "http://cloud.mongodb.com/api/atlas/v1.0/groups/5a66666887d9d61443b41645/clusters/Cluster0/restoreJobs",
			"rel" : "self"
		  } ],
		  "results" : [{
				  "batchId": "5a66783b80eef5354c77ee13",
				  "clusterId": "5a66689487d9d61443b46149",
				  "clusterName": "Cluster0",
				  "created": "2018-09-20T15:02:00Z",
				  "delivery": {
						"methodName": "HTTP",
						"statusName": "READY"
				  },
				  "encryptionEnabled": false,
				  "groupId": "5a66666887d9d61443b41645",
				  "id": "6b77893b80eef5354c77ee15",
				  "links": [{
						"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a66666887d9d61443b41645/clusters/Cluster0/restoreJobs/1",
						"rel": "self"
				  }],
				  "snapshotId": "6b77893b80eef5354c77ee15",
				  "statusName": "FINISHED",
				  "timestamp": {
					"date": "2018-09-15T15:53:00Z",
					"increment": 1
				  }
			}],
		  "totalCount" : 1
		}}`,
		)
	})

	request := &ContinuousJobRequest{
		Delivery: Delivery{
			MethodName:        "AUTOMATED_RESTORE",
			TargetClusterName: "Cluster0",
			TargetGroupID:     "5a66666887d9d61443b41645",
		},
		PointInTimeUTCMillis: 1536610288000,
	}

	customDBRoles, _, err := client.ContinuousRestoreJobs.Create(ctx, "5a66666887d9d61443b41645", "Cluster0", request)
	if err != nil {
		t.Fatalf("ContinuousRestoreJobs.Create returned error: %v", err)
	}

	expected := &ContinuousJobs{
		Links: []*Link{
			{
				Href: "http://cloud.mongodb.com/api/atlas/v1.0/groups/5a66666887d9d61443b41645/clusters/Cluster0/restoreJobs",
				Rel:  "self",
			},
		},
		Results: []*ContinuousJob{
			{
				BatchID:     "5a66783b80eef5354c77ee13",
				ClusterID:   "5a66689487d9d61443b46149",
				ClusterName: "Cluster0",
				Created:     "2018-09-20T15:02:00Z",
				Delivery: &Delivery{
					MethodName: "HTTP",
					StatusName: "READY",
				},
				EncryptionEnabled: false,
				GroupID:           "5a66666887d9d61443b41645",
				ID:                "6b77893b80eef5354c77ee15",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a66666887d9d61443b41645/clusters/Cluster0/restoreJobs/1",
						Rel:  "self",
					},
				},
				SnapshotID: "6b77893b80eef5354c77ee15",
				StatusName: "FINISHED",
				Timestamp: SnapshotTimestamp{
					Date:      "2018-09-15T15:53:00Z",
					Increment: 1,
				},
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(customDBRoles, expected); diff != nil {
		t.Error(diff)
	}
}
