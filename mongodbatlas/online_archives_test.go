// Copyright 2019 MongoDB Inc
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

func TestOnlineArchiveServiceOp_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "5e2211c17a3e5a48f5497de3"
	clusterName := "myTestClstr"
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/onlineArchives", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{"links": [
			 {
				 "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5e2211c17a3e5a48f5497de3/clusters/myTestCluster/onlineArchives?pageNum=1&itemsPerPage=100",
				 "rel": "self"
			 }
		  ],
		  "results": [
		  {
			 "_id": "5ebad3c1fe9c0ab8d37d61e1",
			 "clusterName": "myTestClstr",
			 "collName": "employees",
			 "criteria": {
				 "dateField": "created",
				 "expireAfterDays": 5
			 },
			 "dbName": "people",
			 "groupId": "5e2211c17a3e5a48f5497de3",
			 "partitionFields": [
				 {
					 "fieldName": "firstName",
					 "fieldType": "string",
					 "order": 0
				 },
				 {
					 "fieldName": "lastName",
					 "fieldType": "string",
					 "order": 1
				 },
				 {
					 "fieldName": "created",
					 "fieldType": "date",
					 "order": 2
				 }
			 ],
			 "paused": false
		  },
		  {
			 "_id": "5ebc2789fe9c0ab8d33b96cd",
			 "clusterName": "myTestClstr",
			 "collName": "invoices",
			 "criteria": {
				 "dateField": "year",
				 "expireAfterDays": 5
			 },
			 "dbName": "accounting",
			 "groupId": "5e2211c17a3e5a48f5497de3",
			 "partitionFields": [
				 {
					 "fieldName": "year",
					 "fieldType": "date",
					 "order": 0
				 },
				 {
					 "fieldName": "month",
					 "fieldType": "string",
					 "order": 1
				 },
				 {
					 "fieldName": "invoiceNumber",
					 "fieldType": "int",
					 "order": 2
				 }
			 ],
			 "paused": false
		  }
		],
		"totalCount": 2 
		}`)
	})

	archives, _, err := client.OnlineArchives.List(ctx, groupID, clusterName, nil)
	if err != nil {
		t.Fatalf("OnlineArchives.List returned error: %v", err)
	}

	expected := &OnlineArchives{
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5e2211c17a3e5a48f5497de3/clusters/myTestCluster/onlineArchives?pageNum=1&itemsPerPage=100",
			},
		},
		Results: []*OnlineArchive{
			{
				ID:          "5ebad3c1fe9c0ab8d37d61e1",
				ClusterName: clusterName,
				CollName:    "employees",
				Criteria: &OnlineArchiveCriteria{
					DateField:       "created",
					ExpireAfterDays: pointy.Float64(5),
				},
				DBName:  "people",
				GroupID: groupID,
				PartitionFields: []*PartitionFields{
					{
						FieldName: "firstName",
						FieldType: "string",
						Order:     pointy.Float64(0),
					},
					{
						FieldName: "lastName",
						FieldType: "string",
						Order:     pointy.Float64(1),
					},
					{
						FieldName: "created",
						FieldType: "date",
						Order:     pointy.Float64(2),
					},
				},
				Paused: pointy.Bool(false),
			},
			{
				ID:          "5ebc2789fe9c0ab8d33b96cd",
				ClusterName: clusterName,
				CollName:    "invoices",
				Criteria: &OnlineArchiveCriteria{
					DateField:       "year",
					ExpireAfterDays: pointy.Float64(5),
				},
				DBName:  "accounting",
				GroupID: groupID,
				PartitionFields: []*PartitionFields{
					{
						FieldName: "year",
						FieldType: "date",
						Order:     pointy.Float64(0),
					},
					{
						FieldName: "month",
						FieldType: "string",
						Order:     pointy.Float64(1),
					},
					{
						FieldName: "invoiceNumber",
						FieldType: "int",
						Order:     pointy.Float64(2),
					},
				},
				Paused: pointy.Bool(false),
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(archives, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOnlineArchiveServiceOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "5e2211c17a3e5a48f5497de3"
	clusterName := "myTestClstr"
	archiveID := "5ebad3c1fe9c0ab8d37d61e1"
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/onlineArchives/%s", groupID, clusterName, archiveID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
			 "_id": "5ebad3c1fe9c0ab8d37d61e1",
			 "clusterName": "myTestClstr",
			 "collName": "employees",
			 "criteria": {
				 "dateField": "created",
				 "expireAfterDays": 5
			 },
			 "dbName": "people",
			 "groupId": "5e2211c17a3e5a48f5497de3",
			 "partitionFields": [
				 {
					 "fieldName": "firstName",
					 "fieldType": "string",
					 "order": 0
				 },
				 {
					 "fieldName": "lastName",
					 "fieldType": "string",
					 "order": 1
				 },
				 {
					 "fieldName": "created",
					 "fieldType": "date",
					 "order": 2
				 }
			 ],
			 "paused": false
		  }`)
	})

	archive, _, err := client.OnlineArchives.Get(ctx, groupID, clusterName, archiveID)
	if err != nil {
		t.Fatalf("OnlineArchives.Get returned error: %v", err)
	}

	expected := &OnlineArchive{
		ID:          archiveID,
		ClusterName: clusterName,
		CollName:    "employees",
		Criteria: &OnlineArchiveCriteria{
			DateField:       "created",
			ExpireAfterDays: pointy.Float64(5),
		},
		DBName:  "people",
		GroupID: groupID,
		PartitionFields: []*PartitionFields{
			{
				FieldName: "firstName",
				FieldType: "string",
				Order:     pointy.Float64(0),
			},
			{
				FieldName: "lastName",
				FieldType: "string",
				Order:     pointy.Float64(1),
			},
			{
				FieldName: "created",
				FieldType: "date",
				Order:     pointy.Float64(2),
			},
		},
		Paused: pointy.Bool(false),
	}

	if diff := deep.Equal(archive, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOnlineArchiveServiceOp_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "5e2211c17a3e5a48f5497de3"
	clusterName := "myTestClstr"

	createRequest := &OnlineArchive{
		CollName: "employees",
		Criteria: &OnlineArchiveCriteria{
			DateField:       "created",
			ExpireAfterDays: pointy.Float64(5),
		},
		DBName: "people",
		PartitionFields: []*PartitionFields{
			{
				FieldName: "created",
				FieldType: "date",
				Order:     pointy.Float64(0),
			},
		},
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/onlineArchives", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"collName": "employees",
			"criteria": map[string]interface{}{
				"dateField":       "created",
				"expireAfterDays": float64(5),
			},
			"dbName": "people",
			"partitionFields": []interface{}{
				map[string]interface{}{
					"fieldName": "created",
					"fieldType": "date",
					"order":     float64(0),
				},
			},
		}

		var v map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			t.Fatalf("decode json: %v", err)
		}
		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Clusters.Create Request Body = %v", diff)
		}

		jsonBlob := `{
			"_id": "1",
			"clusterName": "myTestClstr",
			"collName":"employees",
			"criteria": {
                "dateField": "created",
				"expireAfterDays": 5
            },
            "dbName": "people",
			"groupId": "5e2211c17a3e5a48f5497de3",
            "partitionFields": [{
                "fieldName": "created",
                "fieldType": "date",
				"order": 0
            }],
			"paused": false
		}`
		fmt.Fprint(w, jsonBlob)
	})

	archive, _, err := client.OnlineArchives.Create(ctx, groupID, clusterName, createRequest)
	if err != nil {
		t.Fatalf("OnlineArchives.Create returned error: %v", err)
	}

	const expectedDBName = "people"
	if archive.DBName != expectedDBName {
		t.Errorf("expected name '%s', received '%s'", expectedDBName, archive.DBName)
	}

	const expectedColName = "employees"
	if archive.CollName != expectedColName {
		t.Errorf("expected name '%s', received '%s'", expectedColName, archive.CollName)
	}

	if id := archive.GroupID; id != groupID {
		t.Errorf("expected groupId '%s', received '%s'", groupID, id)
	}
}

func TestOnlineArchiveServiceOp_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "5e2211c17a3e5a48f5497de3"
	clusterName := "myTestClstr"
	archiveID := "5ebad3c1fe9c0ab8d37d61e1"

	updateRequest := &OnlineArchive{
		Criteria: &OnlineArchiveCriteria{
			ExpireAfterDays: pointy.Float64(6),
		},
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/onlineArchives/%s", groupID, clusterName, archiveID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"criteria": map[string]interface{}{
				"expireAfterDays": float64(6),
			},
		}

		var v map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			t.Fatalf("decode json: %v", err)
		}
		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Clusters.Create Request Body = %v", diff)
		}

		jsonBlob := `{
			"_id": "5ebad3c1fe9c0ab8d37d61e1",
			"clusterName": "myTestClstr",
			"collName":"employees",
			"criteria": {
                "dateField": "created",
				"expireAfterDays": 6
            },
            "dbName": "people",
			"groupId": "5e2211c17a3e5a48f5497de3",
            "partitionFields": [{
                "fieldName": "created",
                "fieldType": "date",
				"order": 0
            }],
			"paused": false
		}`
		fmt.Fprint(w, jsonBlob)
	})

	archive, _, err := client.OnlineArchives.Update(ctx, groupID, clusterName, archiveID, updateRequest)
	if err != nil {
		t.Fatalf("OnlineArchives.Update returned error: %v", err)
	}

	expectedExpireAfterDays := pointy.Float64(6)
	if *(archive.Criteria.ExpireAfterDays) != *expectedExpireAfterDays {
		t.Errorf("expected expireAfterDays '%f', received '%f'", *expectedExpireAfterDays, *archive.Criteria.ExpireAfterDays)
	}
}

func TestOnlineArchiveServiceOp_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "5e2211c17a3e5a48f5497de3"
	clusterName := "myTestClstr"
	archiveID := "5ebad3c1fe9c0ab8d37d61e1"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/onlineArchives/%s", groupID, clusterName, archiveID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.OnlineArchives.Delete(ctx, groupID, clusterName, archiveID)
	if err != nil {
		t.Fatalf("OnlineArchives.Delete returned error: %v", err)
	}
}

func TestOnlineArchiveServiceOp_CreatePrivateLinkEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateNetworkSettings/endpointIds", groupID)

	createRequest := &PrivateLinkEndpointOnlineArchive{
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

	privateEndpointConnection, _, err := client.OnlineArchives.CreatePrivateLinkEndpoint(ctx, groupID, createRequest)
	if err != nil {
		t.Errorf("OnlineArchives.CreatePrivateLinkEndpoint returned error: %v", err)
		return
	}

	expected := &PrivateLinkEndpointOnlineArchiveResponse{
		Results: []*PrivateLinkEndpointOnlineArchive{
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
		t.Errorf("OnlineArchives.CreatePrivateLinkEndpoint\n got=%#v\nwant=%#v", privateEndpointConnection, expected)
	}
}

func TestOnlineArchiveServiceOp_GetPrivateLinkEndpoint(t *testing.T) {
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

	privateLinkEndpointADL, _, err := client.OnlineArchives.GetPrivateLinkEndpoint(ctx, groupID, "1")
	if err != nil {
		t.Fatalf("OnlineArchives.GetPrivateLinkEndpoint returned error: %v", err)
	}

	expected := PrivateLinkEndpointOnlineArchive{
		Comment:    "Private endpoint for Application Server A",
		EndpointID: "vpce-jjg5e24qp93513h03",
		Provider:   "AWS",
		Type:       "DATA_LAKE",
	}

	if diff := deep.Equal(privateLinkEndpointADL, &expected); diff != nil {
		t.Error(diff)
	}
}

func TestOnlineArchiveServiceOp_ListPrivateLinkEndpoint(t *testing.T) {
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

	privateLinkEndpoints, _, err := client.OnlineArchives.ListPrivateLinkEndpoint(ctx, groupID)
	if err != nil {
		t.Fatalf("OnlineArchives.ListPrivateLinkEndpoint returned error: %v", err)
	}

	expected := &PrivateLinkEndpointOnlineArchiveResponse{
		Results: []*PrivateLinkEndpointOnlineArchive{
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
		t.Errorf("OnlineArchives.ListPrivateLinkEndpoint\n got=%#v\nwant=%#v", privateLinkEndpoints, expected)
	}
}

func TestOnlineArchiveServiceOp_DeletePrivateLinkEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateNetworkSettings/endpointIds/1", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.OnlineArchives.DeletePrivateLinkEndpoint(ctx, groupID, "1")
	if err != nil {
		t.Errorf("OnlineArchives.DeletePrivateLinkEndpoint returned error: %v", err)
	}
}
