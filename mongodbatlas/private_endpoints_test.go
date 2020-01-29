package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/mwielbut/pointy"
)

func TestPrivateEndpoint_Create(t *testing.T) {
	setup()
	defer teardown()

	groupID := "1"

	createRequest := &PrivateEndpointConnection{
		ProviderName: "AWS",
		Region:       "us-east-1",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint", groupID), func(w http.ResponseWriter, r *http.Request) {
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

	privateEndpointConnection, _, err := client.PrivateEndpoints.Create(ctx, groupID, createRequest)
	if err != nil {
		t.Errorf("PrivateEndpoints.Create returned error: %v", err)
		return
	}

	expected := &PrivateEndpointConnection{
		EndpointServiceName: "com.amazonaws.vpce.us-east-1.vpce-svc-0aee615d3fe32c14e",
		ID:                  "5df264b8f10fab7d2cad2f0d",
		InterfaceEndpoints:  []string{"vpce-08fb7e9319909ec7b"},
		Status:              "WAITING_FOR_USER",
	}

	if !reflect.DeepEqual(privateEndpointConnection, expected) {
		t.Errorf("PrivateEndpoints.Create\n got=%#v\nwant=%#v", privateEndpointConnection, expected)
	}
}

func TestPrivateEndpoints_Get(t *testing.T) {
	setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s", groupID, privateLinkID), func(w http.ResponseWriter, r *http.Request) {
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

	privateEndpointConnection, _, err := client.PrivateEndpoints.Get(ctx, groupID, privateLinkID)
	if err != nil {
		t.Errorf("PrivateEndpoints.Get returned error: %v", err)
	}

	expected := &PrivateEndpointConnection{
		EndpointServiceName: "com.amazonaws.vpce.us-east-1.vpce-svc-0aee615d3fe32c14e",
		ID:                  "5df264b8f10fab7d2cad2f0d",
		InterfaceEndpoints:  []string{"vpce-08fb7e9319909ec7b"},
		Status:              "WAITING_FOR_USER",
	}

	if !reflect.DeepEqual(privateEndpointConnection, expected) {
		t.Errorf("PrivateEndpoints.Get\n got=%#v\nwant=%#v", privateEndpointConnection, expected)
	}
}

func TestPrivateEndpoints_List(t *testing.T) {
	setup()
	defer teardown()

	groupID := "1"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint", groupID), func(w http.ResponseWriter, r *http.Request) {
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

	privateEndpoints, _, err := client.PrivateEndpoints.List(ctx, groupID, nil)
	if err != nil {
		t.Errorf("PrivateEndpoints.List returned error: %v", err)
	}

	expected := []PrivateEndpointConnection{
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
		t.Errorf("PrivateEndpoints.List\n got=%#v\nwant=%#v", privateEndpoints, expected)
	}
}

func TestPrivateEndpoints_Delete(t *testing.T) {
	setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s", groupID, privateLinkID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PrivateEndpoints.Delete(ctx, groupID, privateLinkID)
	if err != nil {
		t.Errorf("PrivateEndpoints.Delete returned error: %v", err)
	}
}

func TestPrivateEndpoint_AddOneInterfaceEndpoint(t *testing.T) {
	setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	interfaceEndpointID := "vpce-0b9c5701325cb15dd"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/interfaceEndpoints", groupID, privateLinkID), func(w http.ResponseWriter, r *http.Request) {
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

	interfaceEndpoint, _, err := client.PrivateEndpoints.AddOneInterfaceEndpoint(ctx, groupID, privateLinkID, interfaceEndpointID)
	if err != nil {
		t.Errorf("PrivateEndpoints.AddOneInterfaceEndpoint returned error: %v", err)
		return
	}

	expected := &InterfaceEndpointConnection{
		ID:               "vpce-08fb7e9319909ec7b",
		ConnectionStatus: "PENDING",
		DeleteRequested:  pointy.Bool(false),
	}

	if !reflect.DeepEqual(interfaceEndpoint, expected) {
		t.Errorf("PrivateEndpoints.AddOneInterfaceEndpoint\n got=%#v\nwant=%#v", interfaceEndpoint, expected)
	}
}

func TestPrivateEndpoints_GetOneInterfaceEndpoint(t *testing.T) {
	setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	interfaceEndpointID := "vpce-0b9c5701325cb15dd"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/interfaceEndpoints/%s", groupID, privateLinkID, interfaceEndpointID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"interfaceEndpointId": "vpce-08fb7e9319909ec7b",
			"connectionStatus": "PENDING",
			"deleteRequested": false
		}`)
	})

	interfaceEndpoint, _, err := client.PrivateEndpoints.GetOneInterfaceEndpoint(ctx, groupID, privateLinkID, interfaceEndpointID)
	if err != nil {
		t.Errorf("PrivateEndpoints.GetOneInterfaceEndpoint returned error: %v", err)
	}

	expected := &InterfaceEndpointConnection{
		ID:               "vpce-08fb7e9319909ec7b",
		ConnectionStatus: "PENDING",
		DeleteRequested:  pointy.Bool(false),
	}

	if !reflect.DeepEqual(interfaceEndpoint, expected) {
		t.Errorf("PrivateEndpoints.GetOneInterfaceEndpoint\n got=%#v\nwant=%#v", interfaceEndpoint, expected)
	}
}

func TestPrivateEndpoints_DeleteOneInterfaceEndpoint(t *testing.T) {
	setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	interfaceEndpointID := "vpce-0b9c5701325cb15dd"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/interfaceEndpoints/%s", groupID, privateLinkID, interfaceEndpointID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PrivateEndpoints.DeleteOneInterfaceEndpoint(ctx, groupID, privateLinkID, interfaceEndpointID)
	if err != nil {
		t.Errorf("PrivateEndpoints.DeleteOneInterfaceEndpoint returned error: %v", err)
	}
}
