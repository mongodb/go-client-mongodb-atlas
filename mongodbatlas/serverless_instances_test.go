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
	projectID = "PROJECT-ID"
	id        = "1"
)

func TestServerlessInstances_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/"+serverlessInstancesPath, projectID), func(w http.ResponseWriter, r *http.Request) {
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
				"groupId" : "PROJECT-ID",
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
				"stateName": "IDLE"
			  },{ 
				"connectionStrings": {
				  "standardSrv" : "mongodb+srv://instance1.example.com"
				},
				"createDate": "2021-06-25T21:32:06Z",
				"groupId": "PROJECT-ID",
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
				"stateName": "IDLE"
			  }],
			"totalCount": 2
		}`)
	})

	serverlessInstances, _, err := client.ServerlessInstances.List(ctx, projectID, nil)
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
				GroupID:           projectID,
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
			},
			{
				ID:                "2",
				GroupID:           projectID,
				MongoDBVersion:    "5.0.0",
				Name:              "test2",
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

	mux.HandleFunc(fmt.Sprintf("/"+serverlessInstancesPath+"/%s", projectID, id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				"connectionStrings": {
				  "standardSrv": "mongodb+srv://instance1.example.com"
				},
				"createDate": "2021-06-25T21:32:06Z",
				"groupId" : "PROJECT-ID",
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
				"stateName": "IDLE"
		}`)
	})

	serverlessInstance, _, err := client.ServerlessInstances.Get(ctx, projectID, id)
	if err != nil {
		t.Fatalf("ServerlessInstances.Get returned error: %v", err)
	}

	expected := &Cluster{
		ID:                id,
		GroupID:           projectID,
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
	}

	if diff := deep.Equal(serverlessInstance, expected); diff != nil {
		t.Error(diff)
	}
}

func TestServerlessInstances_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/"+serverlessInstancesPath, projectID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
				"connectionStrings": {
				  "standardSrv": "mongodb+srv://instance1.example.com"
				},
				"createDate": "2021-06-25T21:32:06Z",
				"groupId" : "PROJECT-ID",
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
				"stateName": "IDLE"
		}`)
	})

	cluster := &Cluster{
		ID:                id,
		GroupID:           projectID,
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
	}

	serverlessInstance, _, err := client.ServerlessInstances.Create(ctx, projectID, cluster)
	if err != nil {
		t.Fatalf("ServerlessInstances.Get returned error: %v", err)
	}

	expected := &Cluster{
		ID:                id,
		GroupID:           projectID,
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
	}

	if diff := deep.Equal(serverlessInstance, expected); diff != nil {
		t.Error(diff)
	}
}

func TestServerlessInstances_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/"+serverlessInstancesPath+"/%s", projectID, id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.ServerlessInstances.Delete(ctx, projectID, id)
	if err != nil {
		t.Fatalf("ServerlessInstances.Delete returned error: %v", err)
	}
}
