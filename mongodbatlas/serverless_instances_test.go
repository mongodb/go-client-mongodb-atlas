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

const (
	id = "1"
)

func TestServerlessInstances_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/"+serverlessInstancesPath, groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/serverless/5c49e72980eef544a218f8f8/accessList/?pretty=true&pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [{
				"connectionStrings": {
				  "standardSrv": "mongodb+srv://instance1.example.com"
				},
				"createDate": "2021-06-25T21:32:06Z",
				"groupId" : "1",
				"id": "1",
				"links": [ {
				  "href": "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}",
				  "rel": "self"
				} ],
				"mongoDBVersion": "5.0.0",
				"name": "test1",
				"providerSettings": {
				  "providerName": "SERVERLESS",
				  "backingProviderName": "AWS",
				  "regionName": "US_EAST_1"
				},
				"stateName": "IDLE",
				"tags": [ { "key": "key1", "value": "value1" } ]
			  },{ 
				"connectionStrings": {
				  "standardSrv" : "mongodb+srv://instance1.example.com"
				},
				"createDate": "2021-06-25T21:32:06Z",
				"groupId": "1",
				"id": "2",
				"links": [{
				  "href": "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}",
				  "rel": "self"
				}],
				"mongoDBVersion": "5.0.0",
				"name": "test2",
				"providerSettings": {
				  "providerName": "SERVERLESS",
				  "backingProviderName": "AWS",
				  "regionName": "US_EAST_1"
				},
				"stateName": "IDLE",
				"tags": [ { "key": "key1", "value": "value1" } ]
			  }],
			"totalCount": 2
		}`)
	})

	serverlessInstances, _, err := client.ServerlessInstances.List(ctx, groupID, nil)
	if err != nil {
		t.Fatalf("ServerlessInstances.List returned error: %v", err)
	}

	expected := &ClustersResponse{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/serverless/5c49e72980eef544a218f8f8/accessList/?pretty=true&pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*Cluster{
			{
				ID:                id,
				GroupID:           groupID,
				MongoDBVersion:    "5.0.0",
				Name:              "test1",
				ProviderSettings:  &ProviderSettings{RegionName: "US_EAST_1", BackingProviderName: "AWS", ProviderName: "SERVERLESS"},
				StateName:         "IDLE",
				ConnectionStrings: &ConnectionStrings{StandardSrv: "mongodb+srv://instance1.example.com"},
				CreateDate:        "2021-06-25T21:32:06Z",
				Tags: &[]*Tag{
					{
						Key:   "key1",
						Value: "value1",
					},
				},
				Links: []*Link{
					{
						Rel:  "self",
						Href: "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}",
					},
				},
			},
			{
				ID:                "2",
				GroupID:           groupID,
				MongoDBVersion:    "5.0.0",
				Name:              "test2",
				ProviderSettings:  &ProviderSettings{RegionName: "US_EAST_1", BackingProviderName: "AWS", ProviderName: "SERVERLESS"},
				StateName:         "IDLE",
				ConnectionStrings: &ConnectionStrings{StandardSrv: "mongodb+srv://instance1.example.com"},
				CreateDate:        "2021-06-25T21:32:06Z",
				Tags: &[]*Tag{
					{
						Key:   "key1",
						Value: "value1",
					},
				},
				Links: []*Link{
					{
						Rel:  "self",
						Href: "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}",
					},
				},
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(serverlessInstances, expected); diff != nil {
		t.Error(diff)
	}
}

func TestServerlessInstances_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/"+serverlessInstancesPath+"/%s", groupID, id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				"connectionStrings": {
				  "standardSrv": "mongodb+srv://instance1.example.com"
				},
				"createDate": "2021-06-25T21:32:06Z",
				"groupId" : "1",
				"id": "1",
				"links": [ {
				  "href": "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}",
				  "rel": "self"
				} ],
				"mongoDBVersion": "5.0.0",
				"name": "test1",
				"providerSettings": {
				  "providerName": "SERVERLESS",
				  "backingProviderName": "AWS",
				  "regionName": "US_EAST_1"
				},
				"terminationProtectionEnabled": false,
				"stateName": "IDLE"
		}`)
	})

	serverlessInstance, _, err := client.ServerlessInstances.Get(ctx, groupID, id)
	if err != nil {
		t.Fatalf("ServerlessInstances.Get returned error: %v", err)
	}

	expected := &Cluster{
		ID:                           id,
		GroupID:                      groupID,
		MongoDBVersion:               "5.0.0",
		Name:                         "test1",
		ProviderSettings:             &ProviderSettings{RegionName: "US_EAST_1", BackingProviderName: "AWS", ProviderName: "SERVERLESS"},
		StateName:                    "IDLE",
		TerminationProtectionEnabled: pointer(false),
		ConnectionStrings:            &ConnectionStrings{StandardSrv: "mongodb+srv://instance1.example.com"},
		CreateDate:                   "2021-06-25T21:32:06Z",
		Links: []*Link{
			{
				Rel:  "self",
				Href: "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}",
			},
		},
	}

	if diff := deep.Equal(serverlessInstance, expected); diff != nil {
		t.Error(diff)
	}
}

func TestServerlessInstances_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/"+serverlessInstancesPath, groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
				"connectionStrings": {
				  "standardSrv": "mongodb+srv://instance1.example.com"
				},
				"createDate": "2021-06-25T21:32:06Z",
				"groupId" : "1",
				"id": "1",
				"links": [ {
				  "href": "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}",
				  "rel": "self"
				} ],
				"mongoDBVersion": "5.0.0",
				"name": "test1",
				"providerSettings": {
				  "providerName": "SERVERLESS",
				  "backingProviderName": "AWS",
				  "regionName": "US_EAST_1"
				},
				"stateName": "IDLE",
				"tags": [ { "key": "key1", "value": "value1" } ]
		}`)
	})

	bodyParam := &ServerlessCreateRequestParams{
		Name:             "test1",
		ProviderSettings: &ServerlessProviderSettings{RegionName: "US_EAST_1", BackingProviderName: "AWS", ProviderName: "SERVERLESS"},
	}

	serverlessInstance, _, err := client.ServerlessInstances.Create(ctx, groupID, bodyParam)
	if err != nil {
		t.Fatalf("ServerlessInstances.Get returned error: %v", err)
	}

	expected := &Cluster{
		ID:                id,
		GroupID:           groupID,
		MongoDBVersion:    "5.0.0",
		Name:              "test1",
		ProviderSettings:  &ProviderSettings{RegionName: "US_EAST_1", BackingProviderName: "AWS", ProviderName: "SERVERLESS"},
		StateName:         "IDLE",
		ConnectionStrings: &ConnectionStrings{StandardSrv: "mongodb+srv://instance1.example.com"},
		CreateDate:        "2021-06-25T21:32:06Z",
		Links: []*Link{
			{
				Rel:  "self",
				Href: "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}",
			},
		},
		Tags: &[]*Tag{
			{
				Key:   "key1",
				Value: "value1",
			},
		},
	}

	if diff := deep.Equal(serverlessInstance, expected); diff != nil {
		t.Error(diff)
	}
}

func TestServerlessInstances_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/"+serverlessInstancesPath+"/%s", groupID, "sample"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
			"connectionStrings" : {
			  "standardSrv" : "mongodb+srv://instanceName1.example.com"
			},
			"createDate" : "2021-06-25T21:31:10Z",
			"groupId" : "1",
			"id" : "1",
			"links" : [ {
			  "href" : "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}",
			  "rel" : "self"
			}, {
			  "href" : "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}/backup/restoreJobs",
			  "rel" : "http://cloud.mongodb.com/restoreJobs"
			}, {
			  "href" : "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}/backup/snapshots",
			  "rel" : "http://cloud.mongodb.com/snapshots"
			}],
			"mongoDBVersion" : "5.0.0",
			"name" : "sample",
			"providerSettings" : {
			  "providerName" : "SERVERLESS",
			  "backingProviderName" : "AWS",
			  "regionName" : "US_EAST_1"
			},
			"serverlessBackupOptions" : {
			  "serverlessContinuousBackupEnabled" : true
			},
			"stateName" : "IDLE",
			"terminationProtectionEnabled": true,
			"tags": [ { "key": "key1", "value": "value1" } ]
		}`)
	})

	bodyParam := &ServerlessUpdateRequestParams{
		ServerlessBackupOptions:      &ServerlessBackupOptions{ServerlessContinuousBackupEnabled: pointer(true)},
		TerminationProtectionEnabled: pointer(true),
		Tag:                          &[]*Tag{{Key: "key1", Value: "value1"}},
	}

	serverlessInstance, _, err := client.ServerlessInstances.Update(ctx, groupID, "sample", bodyParam)
	if err != nil {
		t.Fatalf("ServerlessInstances.Get returned error: %v", err)
	}

	expected := &Cluster{
		ID:                           id,
		GroupID:                      groupID,
		MongoDBVersion:               "5.0.0",
		Name:                         "sample",
		ProviderSettings:             &ProviderSettings{RegionName: "US_EAST_1", BackingProviderName: "AWS", ProviderName: "SERVERLESS"},
		StateName:                    "IDLE",
		ConnectionStrings:            &ConnectionStrings{StandardSrv: "mongodb+srv://instanceName1.example.com"},
		CreateDate:                   "2021-06-25T21:31:10Z",
		ServerlessBackupOptions:      &ServerlessBackupOptions{ServerlessContinuousBackupEnabled: pointer(true)},
		TerminationProtectionEnabled: pointer(true),
		Links: []*Link{
			{
				Rel:  "self",
				Href: "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}",
			},
			{
				Rel:  "http://cloud.mongodb.com/restoreJobs",
				Href: "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}/backup/restoreJobs",
			},
			{
				Rel:  "http://cloud.mongodb.com/snapshots",
				Href: "http://cloud.mongodb.com/api/atlas/v1.0/groups/{groupId}/serverless/{instanceName1}/backup/snapshots",
			},
		},
		Tags: &[]*Tag{
			{
				Key:   "key1",
				Value: "value1",
			},
		},
	}

	if diff := deep.Equal(serverlessInstance, expected); diff != nil {
		t.Error(diff)
	}

	// Testing for removing tags
	client, mux, teardown = setup()
	defer teardown()

	if diff := deep.Equal(serverlessInstance, expected); diff != nil {
		t.Error(diff)
	}

	mux.HandleFunc(fmt.Sprintf("/"+serverlessInstancesPath+"/%s", groupID, "sample"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
			"tags": []
		}`)
	})

	bodyParam = &ServerlessUpdateRequestParams{
		Tag: &[]*Tag{},
	}

	serverlessInstance, _, err = client.ServerlessInstances.Update(ctx, groupID, "sample", bodyParam)
	if err != nil {
		t.Fatalf("ServerlessInstances.Get returned error: %v", err)
	}

	expected = &Cluster{
		Tags: &[]*Tag{},
	}

	if diff := deep.Equal(serverlessInstance, expected); diff != nil {
		t.Error(diff)
	}
}

func TestServerlessInstances_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/"+serverlessInstancesPath+"/%s", groupID, id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.ServerlessInstances.Delete(ctx, groupID, id)
	if err != nil {
		t.Fatalf("ServerlessInstances.Delete returned error: %v", err)
	}
}
