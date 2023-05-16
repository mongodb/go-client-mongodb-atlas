// Copyright 2023 MongoDB Inc
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

func TestDataFederation_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataFederation", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"cloudProviderConfig": {
					"aws": {
					  	"iamAssumedRoleARN": "arn:aws:iam::123456789012:role/ReadS3BucketRole",
						"externalId": "test",
						"iamUserARN": "test",
						"roleId": "test"
				  	}
			  	},
			  	"dataProcessRegion": {
					"cloudProvider" : "AWS",
					"region" : "VIRGINIA_USA"
			  	},
			  	"name": "UserMetricData",
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
												"path" : "/path",
												"allowInsecure": false,
												"database": "test",
												"databaseRegex": "test",
												"collection": "test",
												"collectionRegex": "test",
												"provenanceFieldName": "test",
												"urls": ["test"]
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

	dataFederationInstances, _, err := client.DataFederation.List(ctx, groupID)
	if err != nil {
		t.Fatalf("DataFederation.List returned error: %v", err)
	}

	expected := []*DataFederationInstance{
		{
			CloudProviderConfig: &CloudProviderConfig{
				AWSConfig: AwsCloudProviderConfig{
					ExternalID:        "test",
					IAMAssumedRoleARN: "arn:aws:iam::123456789012:role/ReadS3BucketRole",
					IAMUserARN:        "test",
					RoleID:            "test",
				},
			},
			DataProcessRegion: &DataProcessRegion{
				CloudProvider: "AWS",
				Region:        "VIRGINIA_USA",
			},
			Name: "UserMetricData",
			Storage: &DataFederationStorage{
				Databases: []*DataFederationDatabase{
					{
						Name: "my.database",
						Collections: []*DataFederationCollection{
							{
								Name: "my.collection",
								DataSources: []*DataFederationDataSource{
									{
										AllowInsecure:       pointer(false),
										Collection:          "test",
										CollectionRegex:     "test",
										Database:            "test",
										DatabaseRegex:       "test",
										DefaultFormat:       ".json",
										Path:                "/path",
										ProvenanceFieldName: "test",
										StoreName:           "store",
										Urls:                []*string{pointer("test")},
									},
								},
							},
						},
						Views: []*DataFederationDatabaseView{
							{
								Name:     "my.view",
								Source:   "source",
								Pipeline: "my.pipeline",
							},
						},
						MaxWildcardCollections: 92,
					},
				},
				Stores: []*DataFederationStore{
					{
						Name:                     "datacenter-alpha",
						Provider:                 "s3",
						Region:                   "us-east-1",
						Bucket:                   "datacenter-alpha",
						Prefix:                   "/metrics",
						Delimiter:                "/",
						IncludeTags:              pointer(false),
						AdditionalStorageClasses: []*string{pointer("STANDARD_IA")},
					},
				},
			},
		},
	}

	if diff := deep.Equal(dataFederationInstances, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataFederation_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	tenantName := "test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataFederation/%s", groupID, tenantName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
			{
				"cloudProviderConfig": {
					"aws": {
					  	"iamAssumedRoleARN": "arn:aws:iam::123456789012:role/ReadS3BucketRole",
						"externalId": "test",
						"iamUserARN": "test",
						"roleId": "test"
				  	}
			  	},
			  	"dataProcessRegion": {
					"cloudProvider" : "AWS",
					"region" : "VIRGINIA_USA"
			  	},
			  	"name": "UserMetricData",
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
												"path" : "/path",
												"allowInsecure": false,
												"database": "test",
												"databaseRegex": "test",
												"collection": "test",
												"collectionRegex": "test",
												"provenanceFieldName": "test",
												"urls": ["test"]
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
			}`)
	})

	dataFederationInstance, _, err := client.DataFederation.Get(ctx, groupID, tenantName)
	if err != nil {
		t.Fatalf("DataFederation.Get returned error: %v", err)
	}

	expected := &DataFederationInstance{
		CloudProviderConfig: &CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				ExternalID:        "test",
				IAMAssumedRoleARN: "arn:aws:iam::123456789012:role/ReadS3BucketRole",
				IAMUserARN:        "test",
				RoleID:            "test",
			},
		},
		DataProcessRegion: &DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "VIRGINIA_USA",
		},
		Name: "UserMetricData",
		Storage: &DataFederationStorage{
			Databases: []*DataFederationDatabase{
				{
					Name: "my.database",
					Collections: []*DataFederationCollection{
						{
							Name: "my.collection",
							DataSources: []*DataFederationDataSource{
								{
									AllowInsecure:       pointer(false),
									Collection:          "test",
									CollectionRegex:     "test",
									Database:            "test",
									DatabaseRegex:       "test",
									DefaultFormat:       ".json",
									Path:                "/path",
									ProvenanceFieldName: "test",
									StoreName:           "store",
									Urls:                []*string{pointer("test")},
								},
							},
						},
					},
					Views: []*DataFederationDatabaseView{
						{
							Name:     "my.view",
							Source:   "source",
							Pipeline: "my.pipeline",
						},
					},
					MaxWildcardCollections: 92,
				},
			},
			Stores: []*DataFederationStore{
				{
					Name:                     "datacenter-alpha",
					Provider:                 "s3",
					Region:                   "us-east-1",
					Bucket:                   "datacenter-alpha",
					Prefix:                   "/metrics",
					Delimiter:                "/",
					IncludeTags:              pointer(false),
					AdditionalStorageClasses: []*string{pointer("STANDARD_IA")},
				},
			},
		},
	}

	if diff := deep.Equal(dataFederationInstance, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataFederation_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataFederation", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
			{
				"cloudProviderConfig": {
					"aws": {
					  	"iamAssumedRoleARN": "arn:aws:iam::123456789012:role/ReadS3BucketRole",
						"externalId": "test",
						"iamUserARN": "test",
						"roleId": "test"
				  	}
			  	},
			  	"dataProcessRegion": {
					"cloudProvider" : "AWS",
					"region" : "VIRGINIA_USA"
			  	},
			  	"name": "UserMetricData",
			  	"state": "ACTIVE",
			  	"hostnames": ["test.mongodb-dev.net"],
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
												"path" : "/path",
												"allowInsecure": false,
												"database": "test",
												"databaseRegex": "test",
												"collection": "test",
												"collectionRegex": "test",
												"provenanceFieldName": "test",
												"urls": ["test"]
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
			}`)
	})

	requestBody := &DataFederationInstance{
		CloudProviderConfig: &CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				ExternalID:        "test",
				IAMAssumedRoleARN: "arn:aws:iam::123456789012:role/ReadS3BucketRole",
				IAMUserARN:        "test",
				RoleID:            "test",
			},
		},
		DataProcessRegion: &DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "VIRGINIA_USA",
		},
		Name: "UserMetricData",
		Storage: &DataFederationStorage{
			Databases: []*DataFederationDatabase{
				{
					Name: "my.database",
					Collections: []*DataFederationCollection{
						{
							Name: "my.collection",
							DataSources: []*DataFederationDataSource{
								{
									AllowInsecure:       pointer(false),
									Collection:          "test",
									CollectionRegex:     "test",
									Database:            "test",
									DatabaseRegex:       "test",
									DefaultFormat:       ".json",
									Path:                "/path",
									ProvenanceFieldName: "test",
									StoreName:           "store",
									Urls:                []*string{pointer("test")},
								},
							},
						},
					},
					Views: []*DataFederationDatabaseView{
						{
							Name:     "my.view",
							Source:   "source",
							Pipeline: "my.pipeline",
						},
					},
					MaxWildcardCollections: 92,
				},
			},
			Stores: []*DataFederationStore{
				{
					Name:                     "datacenter-alpha",
					Provider:                 "s3",
					Region:                   "us-east-1",
					Bucket:                   "datacenter-alpha",
					Prefix:                   "/metrics",
					Delimiter:                "/",
					IncludeTags:              pointer(false),
					AdditionalStorageClasses: []*string{pointer("STANDARD_IA")},
				},
			},
		},
	}
	dataFederationInstance, _, err := client.DataFederation.Create(ctx, groupID, requestBody)
	if err != nil {
		t.Fatalf("DataFederation.Create returned error: %v", err)
	}

	expected := &DataFederationInstance{
		CloudProviderConfig: &CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				ExternalID:        "test",
				IAMAssumedRoleARN: "arn:aws:iam::123456789012:role/ReadS3BucketRole",
				IAMUserARN:        "test",
				RoleID:            "test",
			},
		},
		DataProcessRegion: &DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "VIRGINIA_USA",
		},
		Name:      "UserMetricData",
		State:     "ACTIVE",
		Hostnames: []string{"test.mongodb-dev.net"},
		Storage: &DataFederationStorage{
			Databases: []*DataFederationDatabase{
				{
					Name: "my.database",
					Collections: []*DataFederationCollection{
						{
							Name: "my.collection",
							DataSources: []*DataFederationDataSource{
								{
									AllowInsecure:       pointer(false),
									Collection:          "test",
									CollectionRegex:     "test",
									Database:            "test",
									DatabaseRegex:       "test",
									DefaultFormat:       ".json",
									Path:                "/path",
									ProvenanceFieldName: "test",
									StoreName:           "store",
									Urls:                []*string{pointer("test")},
								},
							},
						},
					},
					Views: []*DataFederationDatabaseView{
						{
							Name:     "my.view",
							Source:   "source",
							Pipeline: "my.pipeline",
						},
					},
					MaxWildcardCollections: 92,
				},
			},
			Stores: []*DataFederationStore{
				{
					Name:                     "datacenter-alpha",
					Provider:                 "s3",
					Region:                   "us-east-1",
					Bucket:                   "datacenter-alpha",
					Prefix:                   "/metrics",
					Delimiter:                "/",
					IncludeTags:              pointer(false),
					AdditionalStorageClasses: []*string{pointer("STANDARD_IA")},
				},
			},
		},
	}

	if diff := deep.Equal(dataFederationInstance, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataFederation_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	tenantName := "test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataFederation/%s", groupID, tenantName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `
			{
				"cloudProviderConfig": {
					"aws": {
					  	"iamAssumedRoleARN": "arn:aws:iam::123456789012:role/ReadS3BucketRole",
						"externalId": "test",
						"iamUserARN": "test",
						"roleId": "test"
				  	}
			  	},
			  	"dataProcessRegion": {
					"cloudProvider" : "AWS",
					"region" : "VIRGINIA_USA"
			  	},
			  	"name": "UserMetricData",
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
												"path" : "/path",
												"allowInsecure": false,
												"database": "test",
												"databaseRegex": "test",
												"collection": "test",
												"collectionRegex": "test",
												"provenanceFieldName": "test",
												"urls": ["test"]
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
			}`)
	})

	requestBody := &DataFederationInstance{
		CloudProviderConfig: &CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				ExternalID:        "test",
				IAMAssumedRoleARN: "arn:aws:iam::123456789012:role/ReadS3BucketRole",
				IAMUserARN:        "test",
				RoleID:            "test",
			},
		},
		DataProcessRegion: &DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "VIRGINIA_USA",
		},
		Name: "UserMetricData",
		Storage: &DataFederationStorage{
			Databases: []*DataFederationDatabase{
				{
					Name: "my.database",
					Collections: []*DataFederationCollection{
						{
							Name: "my.collection",
							DataSources: []*DataFederationDataSource{
								{
									AllowInsecure:       pointer(false),
									Collection:          "test",
									CollectionRegex:     "test",
									Database:            "test",
									DatabaseRegex:       "test",
									DefaultFormat:       ".json",
									Path:                "/path",
									ProvenanceFieldName: "test",
									StoreName:           "store",
									Urls:                []*string{pointer("test")},
								},
							},
						},
					},
					Views: []*DataFederationDatabaseView{
						{
							Name:     "my.view",
							Source:   "source",
							Pipeline: "my.pipeline",
						},
					},
					MaxWildcardCollections: 92,
				},
			},
			Stores: []*DataFederationStore{
				{
					Name:                     "datacenter-alpha",
					Provider:                 "s3",
					Region:                   "us-east-1",
					Bucket:                   "datacenter-alpha",
					Prefix:                   "/metrics",
					Delimiter:                "/",
					IncludeTags:              pointer(false),
					AdditionalStorageClasses: []*string{pointer("STANDARD_IA")},
				},
			},
		},
	}
	dataFederationInstance, _, err := client.DataFederation.Update(ctx, groupID, tenantName, requestBody, nil)
	if err != nil {
		t.Fatalf("DataFederation.Update returned error: %v", err)
	}

	expected := &DataFederationInstance{
		CloudProviderConfig: &CloudProviderConfig{
			AWSConfig: AwsCloudProviderConfig{
				ExternalID:        "test",
				IAMAssumedRoleARN: "arn:aws:iam::123456789012:role/ReadS3BucketRole",
				IAMUserARN:        "test",
				RoleID:            "test",
			},
		},
		DataProcessRegion: &DataProcessRegion{
			CloudProvider: "AWS",
			Region:        "VIRGINIA_USA",
		},
		Name: "UserMetricData",
		Storage: &DataFederationStorage{
			Databases: []*DataFederationDatabase{
				{
					Name: "my.database",
					Collections: []*DataFederationCollection{
						{
							Name: "my.collection",
							DataSources: []*DataFederationDataSource{
								{
									AllowInsecure:       pointer(false),
									Collection:          "test",
									CollectionRegex:     "test",
									Database:            "test",
									DatabaseRegex:       "test",
									DefaultFormat:       ".json",
									Path:                "/path",
									ProvenanceFieldName: "test",
									StoreName:           "store",
									Urls:                []*string{pointer("test")},
								},
							},
						},
					},
					Views: []*DataFederationDatabaseView{
						{
							Name:     "my.view",
							Source:   "source",
							Pipeline: "my.pipeline",
						},
					},
					MaxWildcardCollections: 92,
				},
			},
			Stores: []*DataFederationStore{
				{
					Name:                     "datacenter-alpha",
					Provider:                 "s3",
					Region:                   "us-east-1",
					Bucket:                   "datacenter-alpha",
					Prefix:                   "/metrics",
					Delimiter:                "/",
					IncludeTags:              pointer(false),
					AdditionalStorageClasses: []*string{pointer("STANDARD_IA")},
				},
			},
		},
	}

	if diff := deep.Equal(dataFederationInstance, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataFederation_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	tenantName := "test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataFederation/%s", groupID, tenantName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.DataFederation.Delete(ctx, groupID, tenantName)
	if err != nil {
		t.Fatalf("DataFederation.Delete returned error: %v", err)
	}
}

func TestDataFederationQueryLimit_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	tenantName := "TestInstance"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataFederation/%s/limits", groupID, tenantName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{
        	"currentUsage": 0,
        	"lastModifiedDate": "2023-05-12T10:41:03Z",
			"name": "bytesProcessed.daily",
        	"overrunPolicy": "BLOCK",
        	"tenantName": "TestInstance",
        	"value": 2147483648
    	}]`)
	})

	dataFederationQueryLimits, _, err := client.DataFederation.ListQueryLimits(ctx, groupID, tenantName)
	if err != nil {
		t.Fatalf("DataFederation.ListQueryLimit returned error: %v", err)
	}

	expected := []*DataFederationQueryLimit{
		{
			CurrentUsage:     0,
			LastModifiedDate: "2023-05-12T10:41:03Z",
			Name:             "bytesProcessed.daily",
			OverrunPolicy:    "BLOCK",
			TenantName:       "TestInstance",
			Value:            2147483648,
		},
	}

	if diff := deep.Equal(dataFederationQueryLimits, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataFederationQueryLimit_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	tenantName := "TestInstance"
	limitName := "bytesProcessed.daily"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataFederation/%s/limits/%s", groupID, tenantName, limitName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
    		"currentUsage": 0,
    		"lastModifiedDate": "2023-05-12T10:41:03Z",
    		"name": "bytesProcessed.daily",
    		"overrunPolicy": "BLOCK",
    		"tenantName": "TestInstance",
    		"value": 2147483648
		}`)
	})

	dataFederationQueryLimits, _, err := client.DataFederation.GetQueryLimit(ctx, groupID, tenantName, limitName)
	if err != nil {
		t.Fatalf("DataFederation.GetQueryLimit returned error: %v", err)
	}

	expected := DataFederationQueryLimit{
		CurrentUsage:     0,
		LastModifiedDate: "2023-05-12T10:41:03Z",
		Name:             "bytesProcessed.daily",
		OverrunPolicy:    "BLOCK",
		TenantName:       "TestInstance",
		Value:            2147483648,
	}

	if diff := deep.Equal(*dataFederationQueryLimits, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataFederation_Configure(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	tenantName := "TestInstance"
	limitName := "bytesProcessed.daily"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataFederation/%s/limits/%s", groupID, tenantName, limitName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
    		"currentUsage": 0,
    		"lastModifiedDate": "2023-05-12T10:41:03Z",
    		"name": "bytesProcessed.daily",
    		"overrunPolicy": "BLOCK",
    		"tenantName": "TestInstance",
    		"value": 2147483648
		}`)
	})

	requestBody := &DataFederationQueryLimit{OverrunPolicy: "BLOCK", Value: 2147483648}
	dataFederationQueryLimit, _, err := client.DataFederation.ConfigureQueryLimit(ctx, groupID, tenantName, limitName, requestBody)
	if err != nil {
		t.Fatalf("DataFederation.ConfigureQueryLimit returned error: %v", err)
	}

	expected := DataFederationQueryLimit{
		CurrentUsage:     0,
		LastModifiedDate: "2023-05-12T10:41:03Z",
		Name:             "bytesProcessed.daily",
		OverrunPolicy:    "BLOCK",
		TenantName:       "TestInstance",
		Value:            2147483648,
	}

	if diff := deep.Equal(*dataFederationQueryLimit, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataFederationQueryLimit_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	tenantName := "TestInstance"
	limitName := "bytesProcessed.daily"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/dataFederation/%s/limits/%s", groupID, tenantName, limitName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.DataFederation.DeleteQueryLimit(ctx, groupID, tenantName, limitName)
	if err != nil {
		t.Fatalf("DataFederation.DeleteQueryLimit returned error: %v", err)
	}
}
