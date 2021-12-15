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
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/openlyinc/pointy"
)

func TestDataLakes_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataLakes", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
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
				  	"databases": [
						{
							"name": "my.database",
							"collections": [
								{
									"name": "my.collection",
									"dataSources": [
										{
												"storeName" : "store",
												"defaultFormat" : ".json",
												"path" : "/path"
										}
									]
								}
							],
							"views": [
								{
									"name" : "my.view",
									"source" : "source",
									"pipeline" : "my.pipeline"
								}
							],
							"maxWildcardCollections" : 92
						}
					],
					"stores": [
						{
							"name": "datacenter-alpha",
							"provider": "s3",
						  	"region": "us-east-1",
						  	"bucket": "datacenter-alpha",
						  	"prefix": "/metrics",
						  	"delimiter": "/",
						  	"includeTags": false,
							"additionalStorageClasses" : ["STANDARD_IA"]
						}
					]
				}
			}
		]`)
	})

	snapshots, _, err := client.DataLakes.List(ctx, groupID)
	if err != nil {
		t.Fatalf("DataLake.List returned error: %v", err)
	}

	expected := []DataLake{
		{
			CloudProviderConfig: CloudProviderConfig{
				AWSConfig: AwsCloudProviderConfig{
					IAMAssumedRoleARN: "arn:aws:iam::123456789012:role/ReadS3BucketRole",
					TestS3Bucket:      "",
				},
			},
			DataProcessRegion: DataProcessRegion{
				CloudProvider: "AWS",
				Region:        "VIRGINIA_USA",
			},
			GroupID:   groupID,
			Hostnames: []string{"usermetricdata.mongodb.example.net"},
			Name:      "UserMetricData",
			State:     "ACTIVE",
			Storage: Storage{
				Databases: []DataLakeDatabase{
					{
						Name: "my.database",
						Collections: []DataLakeCollection{
							{
								Name: "my.collection",
								DataSources: []DataLakeDataSource{
									{
										StoreName:     "store",
										DefaultFormat: ".json",
										Path:          "/path",
									},
								},
							},
						},
						Views: []DataLakeDatabaseView{
							{
								Name:     "my.view",
								Source:   "source",
								Pipeline: "my.pipeline",
							},
						},
						MaxWildcardCollections: pointy.Int64(92),
					},
				},
				Stores: []DataLakeStore{
					{
						Name:                     "datacenter-alpha",
						Provider:                 "s3",
						Region:                   "us-east-1",
						Bucket:                   "datacenter-alpha",
						Prefix:                   "/metrics",
						Delimiter:                "/",
						IncludeTags:              pointy.Bool(false),
						AdditionalStorageClasses: []string{"STANDARD_IA"},
					},
				},
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
	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataLakes/%s", groupID, dataLakeName)

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
				  "databases": [],
				  "stores": []
			  }
		}`)
	})

	cloudProviderSnapshot, _, err := client.DataLakes.Get(ctx, groupID, dataLakeName)
	if err != nil {
		t.Fatalf("DataLake.Get returned error: %v", err)
	}

	expected := DataLake{
		CloudProviderConfig: CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				IAMAssumedRoleARN: "arn:aws:iam::123456789012:role/ReadS3BucketRole",
				TestS3Bucket:      "",
			},
		},
		DataProcessRegion: DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "VIRGINIA_USA",
		},
		GroupID:   groupID,
		Hostnames: []string{"usermetricdata.mongodb.example.net"},
		Name:      "UserMetricData",
		State:     "ACTIVE",
		Storage: Storage{
			Databases: []DataLakeDatabase{},
			Stores:    []DataLakeStore{},
		},
	}

	if diff := deep.Equal(cloudProviderSnapshot, &expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLake_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	dataLakeName := "UserMetricData"

	updateRequest := &DataLakeUpdateRequest{
		CloudProviderConfig: &CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				IAMAssumedRoleARN: "new_arn",
				TestS3Bucket:      "new_bucket",
			},
		},
		DataProcessRegion: &DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "DUBLIN_IRL",
		},
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataLakes/%s", groupID, dataLakeName)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		expected := map[string]interface{}{
			"cloudProviderConfig": map[string]interface{}{
				"aws": map[string]interface{}{
					"iamAssumedRoleARN": "new_arn",
					"testS3Bucket":      "new_bucket",
				},
			},
			"dataProcessRegion": map[string]interface{}{
				"cloudProvider": "AWS",
				"region":        "DUBLIN_IRL",
			},
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
			  "cloudProviderConfig": {
				  "aws": {
					  "iamAssumedRoleARN": "new_arn",
					  "testS3Bucket": "new_bucket"
				  }
			  },
			  "dataProcessRegion": {
				"cloudProvider" : "AWS",
				"region" : "DUBLIN_IRL"
			  },
			  "groupId": "6c7498dg87d9e6526801572b",
			  "hostnames": [
				  "usermetricdata.mongodb.example.net"
			  ],
			  "name": "UserMetricData",
			  "state": "ACTIVE",
			  "storage": {
				  "databases": [],
				  "stores": []
			  }
		}`)
	})

	updatedDataLake, _, err := client.DataLakes.Update(ctx, groupID, dataLakeName, updateRequest)
	if err != nil {
		t.Fatalf("DataLake.Update returned error: %v", err)
	}

	expected := DataLake{
		CloudProviderConfig: CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				IAMAssumedRoleARN: "new_arn",
				TestS3Bucket:      "new_bucket",
			},
		},
		DataProcessRegion: DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "DUBLIN_IRL",
		},
		GroupID:   groupID,
		Hostnames: []string{"usermetricdata.mongodb.example.net"},
		Name:      "UserMetricData",
		State:     "ACTIVE",
		Storage: Storage{
			Databases: []DataLakeDatabase{},
			Stores:    []DataLakeStore{},
		},
	}

	if diff := deep.Equal(updatedDataLake, &expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLake_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	dataLakeName := "UserMetricData"

	createRequest := &DataLakeCreateRequest{
		Name: dataLakeName,
		CloudProviderConfig: &CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				RoleID:       "1a234bcd5e67f89a12b345c6",
				TestS3Bucket: "user-metric-data-bucket",
			},
		},
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataLakes", groupID)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		expected := map[string]interface{}{
			"name": "UserMetricData",
			"cloudProviderConfig": map[string]interface{}{
				"aws": map[string]interface{}{
					"roleId":       "1a234bcd5e67f89a12b345c6",
					"testS3Bucket": "user-metric-data-bucket",
				},
			},
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
			  "cloudProviderConfig": {
				"aws": {
				  "externalId" : "12a3bc45-de6f-7890-12gh-3i45jklm6n7o",
				  "iamAssumedRoleARN": "arn:aws:iam::123456789012:role/ReadS3BucketRole",
				  "iamUserARN": "arn:aws:iam::1234567890123:root",
				  "roleId": "1a234bcd5e67f89a12b345c6"
				}
			  },
			  "dataProcessRegion": null,
			  "groupId": "6c7498dg87d9e6526801572b",
			  "hostnames": [
				"hardwaremetricdata.mongodb.example.net"
			  ],
			  "name": "UserMetricData",
			  "state": "ACTIVE",
			  "storage": {
				"databases": [],
				"stores": []
			  }
			}`)
	})

	updatedDataLake, _, err := client.DataLakes.Create(ctx, groupID, createRequest)
	if err != nil {
		t.Fatalf("DataLake.Create returned error: %v", err)
	}

	expected := DataLake{
		CloudProviderConfig: CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				ExternalID:        "12a3bc45-de6f-7890-12gh-3i45jklm6n7o",
				IAMAssumedRoleARN: "arn:aws:iam::123456789012:role/ReadS3BucketRole",
				IAMUserARN:        "arn:aws:iam::1234567890123:root",
				RoleID:            "1a234bcd5e67f89a12b345c6",
			},
		},
		GroupID:   groupID,
		Hostnames: []string{"hardwaremetricdata.mongodb.example.net"},
		Name:      "UserMetricData",
		State:     "ACTIVE",
		Storage: Storage{
			Databases: []DataLakeDatabase{},
			Stores:    []DataLakeStore{},
		},
	}

	if diff := deep.Equal(updatedDataLake, &expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLake_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	dataLakeName := "dataLake"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataLakes/%s", groupID, dataLakeName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.DataLakes.Delete(ctx, groupID, dataLakeName)
	if err != nil {
		t.Fatalf("DataLakes.Delete returned error: %v", err)
	}
}

func TestDataLake_CreatePrivateLinkEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateNetworkSettings/endpointIds", groupID)

	createRequest := &PrivateLinkEndpointDataLake{
		EndpointID: "vpce-jjg5e24qp93513h03",
		Type:       "DATA_LAKE",
		Provider:   "AWS",
		Comment:    "Private endpoint for Application Server A",
	}

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"endpointId": "vpce-jjg5e24qp93513h03",
			"type":       "DATA_LAKE",
			"provider":   "AWS",
			"comment":    "Private endpoint for Application Server A",
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Request body\n got=%#v\nwant=%#v \n\ndiff=%#v", v, expected, diff)
		}

		fmt.Fprint(w, `{
  "links" : [ {
    "href" : "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/privateNetworkSettings/endpointIds?pretty=true&pageNum=1&itemsPerPage=100",
    "rel" : "self"
  } ],
  "results" : [ {
    "comment" : "Private endpoint for Application Server A",
    "endpointId" : "vpce-jjg5e24qp93513h03",
    "provider" : "AWS",
    "type" : "DATA_LAKE"
  } ],
  "totalCount" : 1
}`)
	})

	privateEndpointConnection, _, err := client.DataLakes.CreatePrivateLinkEndpoint(ctx, groupID, createRequest)
	if err != nil {
		t.Errorf("DataLakes.CreatePrivateLinkEndpoint returned error: %v", err)
		return
	}

	expected := &PrivateLinkEndpointDataLakeResponse{
		Results: []*PrivateLinkEndpointDataLake{
			{
				EndpointID: "vpce-jjg5e24qp93513h03",
				Type:       "DATA_LAKE",
				Provider:   "AWS",
				Comment:    "Private endpoint for Application Server A",
			},
		},
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/privateNetworkSettings/endpointIds?pretty=true&pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		TotalCount: 1,
	}

	if !reflect.DeepEqual(privateEndpointConnection, expected) {
		t.Errorf("DataLakes.CreatePrivateLinkEndpoint\n got=%#v\nwant=%#v", privateEndpointConnection, expected)
	}
}

func TestDataLake_GetPrivateLinkEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateNetworkSettings/endpointIds/1", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
   "comment" : "Private endpoint for Application Server A",
   "endpointId" : "vpce-jjg5e24qp93513h03",
   "provider" : "AWS",
   "type" : "DATA_LAKE"
}`)
	})

	privateLinkEndpointADL, _, err := client.DataLakes.GetPrivateLinkEndpoint(ctx, groupID, "1")
	if err != nil {
		t.Fatalf("DataLakes.GetPrivateLinkEndpoint returned error: %v", err)
	}

	expected := PrivateLinkEndpointDataLake{
		Comment:    "Private endpoint for Application Server A",
		EndpointID: "vpce-jjg5e24qp93513h03",
		Provider:   "AWS",
		Type:       "DATA_LAKE",
	}

	if diff := deep.Equal(privateLinkEndpointADL, &expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLake_ListPrivateLinkEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateNetworkSettings/endpointIds", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
  "links" : [ {
    "href" : "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/privateNetworkSettings/endpointIds?pretty=true&pageNum=1&itemsPerPage=100",
     "rel" : "self"
  } ],
  "results" : [ {
    "comment" : "Private endpoint for Application Server A",
    "endpointId" : "vpce-jjg5e24qp93513h03",
     "provider" : "AWS",
     "type" : "DATA_LAKE"
   } ],
   "totalCount" : 1
 }`)
	})

	privateLinkEndpoints, _, err := client.DataLakes.ListPrivateLinkEndpoint(ctx, groupID)
	if err != nil {
		t.Fatalf("DataLakes.ListPrivateLinkEndpoint returned error: %v", err)
	}

	expected := &PrivateLinkEndpointDataLakeResponse{
		Results: []*PrivateLinkEndpointDataLake{
			{
				EndpointID: "vpce-jjg5e24qp93513h03",
				Type:       "DATA_LAKE",
				Provider:   "AWS",
				Comment:    "Private endpoint for Application Server A",
			},
		},
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/privateNetworkSettings/endpointIds?pretty=true&pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		TotalCount: 1,
	}

	if !reflect.DeepEqual(privateLinkEndpoints, expected) {
		t.Errorf("DataLakes.ListPrivateLinkEndpoint\n got=%#v\nwant=%#v", privateLinkEndpoints, expected)
	}
}

func TestDataLake_DeletePrivateLinkEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateNetworkSettings/endpointIds/1", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.DataLakes.DeletePrivateLinkEndpoint(ctx, groupID, "1")
	if err != nil {
		t.Errorf("DataLakes.DeletePrivateLinkEndpoint returned error: %v", err)
	}
}
