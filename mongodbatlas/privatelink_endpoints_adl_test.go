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
)

func TestPrivateLinkEndpointsADLServiceOp_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateNetworkSettings/endpointIds", groupID)

	createRequest := &PrivateLinkEndpointADL{
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

	privateEndpointConnection, _, err := client.PrivateLinkEndpointsADL.Create(ctx, groupID, createRequest)
	if err != nil {
		t.Errorf("PrivateLinkEndpointADL.Create returned error: %v", err)
		return
	}

	expected := &PrivateLinkEndpointADLResponse{
		Results: []*PrivateLinkEndpointADL{
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
		t.Errorf("PrivateLinksEndpointsADL.Create\n got=%#v\nwant=%#v", privateEndpointConnection, expected)
	}
}

func TestPrivateLinkEndpointsADLServiceOp_Get(t *testing.T) {
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

	privateLinkEndpointADL, _, err := client.PrivateLinkEndpointsADL.Get(ctx, groupID, "1")
	if err != nil {
		t.Fatalf("PrivateLinksEndpointsADL.Get returned error: %v", err)
	}

	expected := PrivateLinkEndpointADL{
		Comment:    "Private endpoint for Application Server A",
		EndpointID: "vpce-jjg5e24qp93513h03",
		Provider:   "AWS",
		Type:       "DATA_LAKE",
	}

	if diff := deep.Equal(privateLinkEndpointADL, &expected); diff != nil {
		t.Error(diff)
	}
}

func TestPrivateLinkEndpointsADLServiceOp_List(t *testing.T) {
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

	privateLinkEndpoints, _, err := client.PrivateLinkEndpointsADL.List(ctx, groupID)
	if err != nil {
		t.Fatalf("PrivateLinksEndpointsADL.List returned error: %v", err)
	}

	expected := &PrivateLinkEndpointADLResponse{
		Results: []*PrivateLinkEndpointADL{
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
		t.Errorf("PrivateLinksEndpointsADL.List\n got=%#v\nwant=%#v", privateLinkEndpoints, expected)
	}
}

func TestPrivateLinkEndpointsADLServiceOp_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateNetworkSettings/endpointIds/1", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PrivateLinkEndpointsADL.Delete(ctx, groupID, "1")
	if err != nil {
		t.Errorf("PrivateLinksEndpointsADL.Delete returned error: %v", err)
	}
}
