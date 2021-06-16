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

func TestPrivateEndpointDeprecated_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	createRequest := &PrivateEndpointConnectionDeprecated{
		ProviderName: "AWS",
		Region:       "us-east-1",
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateEndpoint", groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"providerName": "AWS",
			"region":       "us-east-1",
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
			"endpointServiceName": "com.amazonaws.vpce.us-east-1.vpce-svc-0aee615d3fe32c14e",
			"id": "5df264b8f10fab7d2cad2f0d",
			"interfaceEndpoints": [
				"vpce-08fb7e9319909ec7b"
			],
			"status": "WAITING_FOR_USER"
		}`)
	})

	privateEndpointConnection, _, err := client.PrivateEndpointsDeprecated.Create(ctx, groupID, createRequest)
	if err != nil {
		t.Errorf("PrivateEndpointsDeprecated.Create returned error: %v", err)
		return
	}

	expected := &PrivateEndpointConnectionDeprecated{
		EndpointServiceName: "com.amazonaws.vpce.us-east-1.vpce-svc-0aee615d3fe32c14e",
		ID:                  "5df264b8f10fab7d2cad2f0d",
		InterfaceEndpoints:  []string{"vpce-08fb7e9319909ec7b"},
		Status:              "WAITING_FOR_USER",
	}

	if !reflect.DeepEqual(privateEndpointConnection, expected) {
		t.Errorf("PrivateEndpointsDeprecated.Create\n got=%#v\nwant=%#v", privateEndpointConnection, expected)
	}
}

func TestPrivateEndpointsDeprecated_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateEndpoint/%s", groupID, privateLinkID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"endpointServiceName": "com.amazonaws.vpce.us-east-1.vpce-svc-0aee615d3fe32c14e",
			"id": "5df264b8f10fab7d2cad2f0d",
			"interfaceEndpoints": [
				"vpce-08fb7e9319909ec7b"
			],
			"status": "WAITING_FOR_USER"
		}`)
	})

	privateEndpointConnection, _, err := client.PrivateEndpointsDeprecated.Get(ctx, groupID, privateLinkID)
	if err != nil {
		t.Errorf("PrivateEndpointsDeprecated.Get returned error: %v", err)
	}

	expected := &PrivateEndpointConnectionDeprecated{
		EndpointServiceName: "com.amazonaws.vpce.us-east-1.vpce-svc-0aee615d3fe32c14e",
		ID:                  "5df264b8f10fab7d2cad2f0d",
		InterfaceEndpoints:  []string{"vpce-08fb7e9319909ec7b"},
		Status:              "WAITING_FOR_USER",
	}

	if !reflect.DeepEqual(privateEndpointConnection, expected) {
		t.Errorf("PrivateEndpointsDeprecated.Get\n got=%#v\nwant=%#v", privateEndpointConnection, expected)
	}
}

func TestPrivateEndpointsDeprecated_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateEndpoint", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
				{
					"endpointServiceName": "com.amazonaws.vpce.us-east-1.vpce-svc-0aee615d3fe32c14e",
					"id": "5df264b8f10fab7d2cad2f0d",
					"interfaceEndpoints": [
						"vpce-05934e68436496102"
					],
					"status": "WAITING_FOR_USER"
				},
				{
					"endpointServiceName": "com.amazonaws.vpce.us-west-1.vpce-svc-4fbdd5faa86d2e7538",
					"id": "5758160cc363b548152501c2",
					"interfaceEndpoints": [
						"vpce-08fb7e9319909ec7b"
					],
					"status": "WAITING_FOR_USER"
				}
			]`)
	})

	privateEndpoints, _, err := client.PrivateEndpointsDeprecated.List(ctx, groupID, nil)
	if err != nil {
		t.Errorf("PrivateEndpointsDeprecated.List returned error: %v", err)
	}

	expected := []PrivateEndpointConnectionDeprecated{
		{
			EndpointServiceName: "com.amazonaws.vpce.us-east-1.vpce-svc-0aee615d3fe32c14e",
			ID:                  "5df264b8f10fab7d2cad2f0d",
			InterfaceEndpoints:  []string{"vpce-05934e68436496102"},
			Status:              "WAITING_FOR_USER",
		},
		{
			EndpointServiceName: "com.amazonaws.vpce.us-west-1.vpce-svc-4fbdd5faa86d2e7538",
			ID:                  "5758160cc363b548152501c2",
			InterfaceEndpoints:  []string{"vpce-08fb7e9319909ec7b"},
			Status:              "WAITING_FOR_USER",
		},
	}

	if !reflect.DeepEqual(privateEndpoints, expected) {
		t.Errorf("PrivateEndpointsDeprecated.List\n got=%#v\nwant=%#v", privateEndpoints, expected)
	}
}

func TestPrivateEndpointsDeprecated_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateEndpoint/%s", groupID, privateLinkID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PrivateEndpointsDeprecated.Delete(ctx, groupID, privateLinkID)
	if err != nil {
		t.Errorf("PrivateEndpointsDeprecated.Delete returned error: %v", err)
	}
}

func TestPrivateEndpointDeprecated_AddOneInterfaceEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	interfaceEndpointID := "vpce-0b9c5701325cb15dd"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateEndpoint/%s/interfaceEndpoints", groupID, privateLinkID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"interfaceEndpointId": "vpce-0b9c5701325cb15dd",
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
			"interfaceEndpointId": "vpce-08fb7e9319909ec7b",
			"connectionStatus": "PENDING",
			"deleteRequested": false
		}`)
	})

	interfaceEndpoint, _, err := client.PrivateEndpointsDeprecated.AddOneInterfaceEndpoint(ctx, groupID, privateLinkID, interfaceEndpointID)
	if err != nil {
		t.Errorf("PrivateEndpointsDeprecated.AddOnePrivateEndpoint returned error: %v", err)
		return
	}

	expected := &InterfaceEndpointConnectionDeprecated{
		ID:               "vpce-08fb7e9319909ec7b",
		ConnectionStatus: "PENDING",
		DeleteRequested:  pointy.Bool(false),
	}

	if !reflect.DeepEqual(interfaceEndpoint, expected) {
		t.Errorf("PrivateEndpointsDeprecated.AddOnePrivateEndpoint\n got=%#v\nwant=%#v", interfaceEndpoint, expected)
	}
}

func TestPrivateEndpointsDeprecated_GetOneInterfaceEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	interfaceEndpointID := "vpce-0b9c5701325cb15dd"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateEndpoint/%s/interfaceEndpoints/%s", groupID, privateLinkID, interfaceEndpointID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"interfaceEndpointId": "vpce-08fb7e9319909ec7b",
			"connectionStatus": "PENDING",
			"deleteRequested": false
		}`)
	})

	interfaceEndpoint, _, err := client.PrivateEndpointsDeprecated.GetOneInterfaceEndpoint(ctx, groupID, privateLinkID, interfaceEndpointID)
	if err != nil {
		t.Errorf("PrivateEndpointsDeprecated.GetOnePrivateEndpoint returned error: %v", err)
	}

	expected := &InterfaceEndpointConnectionDeprecated{
		ID:               "vpce-08fb7e9319909ec7b",
		ConnectionStatus: "PENDING",
		DeleteRequested:  pointy.Bool(false),
	}

	if !reflect.DeepEqual(interfaceEndpoint, expected) {
		t.Errorf("PrivateEndpointsDeprecated.GetOnePrivateEndpoint\n got=%#v\nwant=%#v", interfaceEndpoint, expected)
	}
}

func TestPrivateEndpointsDeprecated_DeleteOneInterfaceEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	interfaceEndpointID := "vpce-0b9c5701325cb15dd"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateEndpoint/%s/interfaceEndpoints/%s", groupID, privateLinkID, interfaceEndpointID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PrivateEndpointsDeprecated.DeleteOneInterfaceEndpoint(ctx, groupID, privateLinkID, interfaceEndpointID)
	if err != nil {
		t.Errorf("PrivateEndpointsDeprecated.DeleteOnePrivateEndpoint returned error: %v", err)
	}
}
