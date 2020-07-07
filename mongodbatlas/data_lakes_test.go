package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestDataLakes_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/groups/%s/datalakes", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"cloudProviderConfig": {
					"aws": {}
				},
				"dataProcessRegion": {},
				"groupId": "6c7498dg87d9e6526801572b",
				"hostnames": [
					"datalake0-ta1ir.a.query.mongodb.net"
				],
				"name": "DataLake0",
				"state": "UNVERIFIED",
				"storage": {
					"databases": []
			}
			{
			  "cloudProviderConfig": {
				  "aws": {
					  "iamAssumedRoleARN": "arn:aws:iam::123456789012:role/ReadS3BucketRole"
				  }
			  },
			  "dataProcessRegion": {
				"cloudProvider" : "AWS",
				"region" : "VIRGINIA_USA"
			  },
			  "groupId": "6c7498dg87d9e6526801572b",
			  "hostnames": [
				  "usermetricdata.mongodb.example.net"
			  ],
			  "name": "UserMetricData",
			  "state": "ACTIVE",
			  "storage": {
				  "databases": {},
				  "stores": []
			  }
			}
		]`)
	})

	snapshots, _, err := client.DataLakes.List(ctx, &DataLakeReqPathParameters{
		GroupID: groupID,
	}, nil)
	if err != nil {
		t.Fatalf("DataLake.List returned error: %v", err)
	}

	expected := []*DataLake{
		{
			CloudProviderConfig: CloudProviderConfig{},
			DataProcessRegion:   DataProcessRegion{},
			GroupID:             groupID,
			Hostnames:           []string{"my.data.lake"},
			Name:                "dataLake1",
			State:               "UNVERIFIED",
			Storage:             Storage{},
		},
		{
			CloudProviderConfig: CloudProviderConfig{},
			DataProcessRegion: DataProcessRegion{
				CloudProvider: "AWS",
				Region:        "VIRGINIA_USA",
			},
			GroupID:   groupID,
			Hostnames: []string{"usermetricdata.mongodb.example.net"},
			Name:      "UserMetricData",
			State:     "ACTIVE",
			Storage: Storage{
				Databases: nil,
				Stores:    []DataLakeStore{},
			},
		},
	}

	if diff := deep.Equal(snapshots, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLake_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	dataLakeName := "UserMetricData"
	path := fmt.Sprintf("/groups/%s/datalakes/%s", groupID, dataLakeName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			  "cloudProviderConfig": {
				  "aws": {
					  "iamAssumedRoleARN": "arn:aws:iam::123456789012:role/ReadS3BucketRole"
				  }
			  },
			  "dataProcessRegion": {
				"cloudProvider" : "AWS",
				"region" : "VIRGINIA_USA"
			  },
			  "groupId": "6c7498dg87d9e6526801572b",
			  "hostnames": [
				  "usermetricdata.mongodb.example.net"
			  ],
			  "name": "UserMetricData",
			  "state": "ACTIVE",
			  "storage": {
				  "databases": {},
				  "stores": []
			  }
			}`)
	})

	cloudProviderSnapshot, _, err := client.DataLakes.Get(ctx, &DataLakeReqPathParameters{
		GroupID: groupID,
		Name:    dataLakeName,
	})
	if err != nil {
		t.Fatalf("DataLake.Get returned error: %v", err)
	}

	expected := DataLake{
		CloudProviderConfig: CloudProviderConfig{},
		DataProcessRegion: DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "VIRGINIA_USA",
		},
		GroupID:   groupID,
		Hostnames: []string{"usermetricdata.mongodb.example.net"},
		Name:      "UserMetricData",
		State:     "ACTIVE",
		Storage: Storage{
			Databases: nil,
			Stores:    []DataLakeStore{},
		},
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLake_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	dataLakeName := "UserMetricData"

	updateRequest := &DataLakeUpdateRequest{
		CloudProviderConfig: CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				IAMAssumedRoleARN: "new_arn",
				TestS3Bucket:      "new_bucket",
			},
		},
		DataProcessRegion: DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "VIRGINIA_USA",
		},
	}

	path := fmt.Sprintf("/groups/%s/datalakes/%s", groupID, dataLakeName)
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

	updatedDataLake, _, err := client.DataLakes.Update(ctx, &DataLakeReqPathParameters{
		GroupID: groupID,
		Name:    dataLakeName,
	}, updateRequest)
	if err != nil {
		t.Fatalf("DataLake.Update returned error: %v", err)
	}

	expected := DataLake{
		CloudProviderConfig: CloudProviderConfig{},
		DataProcessRegion: DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "VIRGINIA_USA",
		},
		GroupID:   groupID,
		Hostnames: []string{"usermetricdata.mongodb.example.net"},
		Name:      "UserMetricData",
		State:     "ACTIVE",
		Storage: Storage{
			Databases: nil,
			Stores:    []DataLakeStore{},
		},
	}

	if diff := deep.Equal(updatedDataLake, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLake_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	dataLakeName := "dataLake"

	path := fmt.Sprintf("/groups/%s/clusters/%s", groupID, dataLakeName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.DataLakes.Delete(ctx, &DataLakeReqPathParameters{
		GroupID: groupID,
		Name:    dataLakeName,
	})
	if err != nil {
		t.Fatalf("DataLakes.Delete returned error: %v", err)
	}
}
