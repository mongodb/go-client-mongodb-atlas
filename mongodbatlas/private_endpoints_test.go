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

func TestPrivateEndpointAWS_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	createRequest := &PrivateEndpointConnection{
		ProviderName: "AWS",
		Region:       "us-east-1",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/endpointService", groupID), func(w http.ResponseWriter, r *http.Request) {
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

func TestPrivateEndpointAzure_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	createRequest := &PrivateEndpointConnection{
		ProviderName: "AZURE",
		Region:       "eastus2",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/endpointService", groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"providerName": "AZURE",
			"region":       "eastus2",
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
"privateLinkServiceName": "pls_5f7b96cd8c39f93d0e5f1911",
"id": "5f7c8d4702d59b7fc0b4b842",
"privateEndpoints": [
	"vpce-08fb7e9319909ec7b"
],
"status": "AVAILABLE"
}`)
	})

	privateEndpointConnection, _, err := client.PrivateEndpoints.Create(ctx, groupID, createRequest)
	if err != nil {
		t.Errorf("PrivateEndpoints.Create returned error: %v", err)
		return
	}

	expected := &PrivateEndpointConnection{
		PrivateLinkServiceName: "pls_5f7b96cd8c39f93d0e5f1911",
		ID:                     "5f7c8d4702d59b7fc0b4b842",
		PrivateEndpoints:       []string{"vpce-08fb7e9319909ec7b"},
		Status:                 "AVAILABLE",
	}

	if !reflect.DeepEqual(privateEndpointConnection, expected) {
		t.Errorf("PrivateEndpoints.Create\n got=%#v\nwant=%#v", privateEndpointConnection, expected)
	}
}

func TestPrivateEndpointsAWS_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService/%s", groupID, "AWS", privateLinkID), func(w http.ResponseWriter, r *http.Request) {
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

	privateEndpointConnection, _, err := client.PrivateEndpoints.Get(ctx, groupID, "AWS", privateLinkID)
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

func TestPrivateEndpointsAzure_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService/%s", groupID, "AZURE", privateLinkID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"privateLinkServiceName": "pls_5f7b96cd8c39f93d0e5f1911",
			"id": "5f7c8d4702d59b7fc0b4b842",
			"privateLinkServiceResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg_5f7b96cd8c39f93d0e5f1911_8nw06xue/providers/Microsoft.Network/privateLinkServices/pls_5f7b96cd8c39f93d0e5f1911",
			"privateEndpoints": [
				"vpce-08fb7e9319909ec7b"
			],
			"status": "AVAILABLE"
		}`)
	})

	privateEndpointConnection, _, err := client.PrivateEndpoints.Get(ctx, groupID, "AZURE", privateLinkID)
	if err != nil {
		t.Errorf("PrivateEndpoints.Get returned error: %v", err)
	}

	expected := &PrivateEndpointConnection{
		PrivateLinkServiceName:       "pls_5f7b96cd8c39f93d0e5f1911",
		ID:                           "5f7c8d4702d59b7fc0b4b842",
		PrivateEndpoints:             []string{"vpce-08fb7e9319909ec7b"},
		PrivateLinkServiceResourceID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg_5f7b96cd8c39f93d0e5f1911_8nw06xue/providers/Microsoft.Network/privateLinkServices/pls_5f7b96cd8c39f93d0e5f1911",
		Status:                       "AVAILABLE",
	}

	if !reflect.DeepEqual(privateEndpointConnection, expected) {
		t.Errorf("PrivateEndpoints.Get\n got=%#v\nwant=%#v", privateEndpointConnection, expected)
	}
}

func TestPrivateEndpointsAWS_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService", groupID, "AWS"), func(w http.ResponseWriter, r *http.Request) {
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

	privateEndpoints, _, err := client.PrivateEndpoints.List(ctx, groupID, "AWS", nil)
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

func TestPrivateEndpointsAzure_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService", groupID, "AZURE"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
		  {
			"id": "5f7c8d4702d59b7fc0b4b842",
			"privateEndpoints": [],
			"privateLinkServiceName": "pls_5f7b96cd8c39f93d0e5f1911",
			"privateLinkServiceResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg_5f7b96cd8c39f93d0e5f1911_8nw06xue/providers/Microsoft.Network/privateLinkServices/pls_5f7b96cd8c39f93d0e5f1911",
			"status": "AVAILABLE"
		  },
		  {
			"id": "5f7c8d6d02d59b7fc0b4b899",
			"privateEndpoints": [],
			"privateLinkServiceName": "pls_5f7c8d6d02d59b7fc0b4b89a",
			"privateLinkServiceResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg_5f7c8d6d02d59b7fc0b4b89a_pvrxsft6/providers/Microsoft.Network/privateLinkServices/pls_5f7c8d6d02d59b7fc0b4b89a",
			"status": "AVAILABLE"
		  }
		]`)
	})

	privateEndpoints, _, err := client.PrivateEndpoints.List(ctx, groupID, "AZURE", nil)
	if err != nil {
		t.Errorf("PrivateEndpoints.List returned error: %v", err)
	}

	expected := []PrivateEndpointConnection{
		{
			PrivateLinkServiceName:       "pls_5f7b96cd8c39f93d0e5f1911",
			ID:                           "5f7c8d4702d59b7fc0b4b842",
			PrivateEndpoints:             []string{},
			PrivateLinkServiceResourceID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg_5f7b96cd8c39f93d0e5f1911_8nw06xue/providers/Microsoft.Network/privateLinkServices/pls_5f7b96cd8c39f93d0e5f1911",
			Status:                       "AVAILABLE",
		},
		{
			PrivateLinkServiceName:       "pls_5f7c8d6d02d59b7fc0b4b89a",
			ID:                           "5f7c8d6d02d59b7fc0b4b899",
			PrivateEndpoints:             []string{},
			PrivateLinkServiceResourceID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg_5f7c8d6d02d59b7fc0b4b89a_pvrxsft6/providers/Microsoft.Network/privateLinkServices/pls_5f7c8d6d02d59b7fc0b4b89a",
			Status:                       "AVAILABLE",
		},
	}

	if !reflect.DeepEqual(privateEndpoints, expected) {
		t.Errorf("PrivateEndpoints.List\n got=%#v\nwant=%#v", privateEndpoints, expected)
	}
}

func TestPrivateEndpointsAWS_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService/%s", groupID, "AWS", privateLinkID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PrivateEndpoints.Delete(ctx, groupID, "AWS", privateLinkID)
	if err != nil {
		t.Errorf("PrivateEndpoints.Delete returned error: %v", err)
	}
}

func TestPrivateEndpointsAzure_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService/%s", groupID, "AZURE", privateLinkID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PrivateEndpoints.Delete(ctx, groupID, "AZURE", privateLinkID)
	if err != nil {
		t.Errorf("PrivateEndpoints.Delete returned error: %v", err)
	}
}

func TestPrivateEndpoint_AddOneInterfaceEndpointAWS(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	request := &InterfaceEndpointConnection{
		InterfaceEndpointID: "vpce-0b9c5701325cb15dd",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService/%s/endpoint", groupID, "AWS", privateLinkID), func(w http.ResponseWriter, r *http.Request) {
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

	interfaceEndpoint, _, err := client.PrivateEndpoints.AddOnePrivateEndpoint(ctx, groupID, "AWS", privateLinkID, request)
	if err != nil {
		t.Errorf("PrivateEndpoints.AddOnePrivateEndpoint returned error: %v", err)
		return
	}

	expected := &InterfaceEndpointConnection{
		InterfaceEndpointID: "vpce-08fb7e9319909ec7b",
		ConnectionStatus:    "PENDING",
		DeleteRequested:     pointy.Bool(false),
	}

	if !reflect.DeepEqual(interfaceEndpoint, expected) {
		t.Errorf("PrivateEndpoints.AddOnePrivateEndpoint\n got=%#v\nwant=%#v", interfaceEndpoint, expected)
	}
}

func TestPrivateEndpoint_AddOneInterfaceEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"

	request := &InterfaceEndpointConnection{
		PrivateEndpointResourceID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/privatelink/providers/Microsoft.Network/privateEndpoints/test",
		PrivateEndpointIPAddress:  "10.0.0.4",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService/%s/endpoint", groupID, "AZURE", privateLinkID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"privateEndpointResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/privatelink/providers/Microsoft.Network/privateEndpoints/test",
			"privateEndpointIPAddress":  "10.0.0.4",
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
			"privateEndpointResourceId":  "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/privatelink/providers/Microsoft.Network/privateEndpoints/test",
			"privateEndpointIPAddress": "10.0.0.4",
			"connectionStatus": "INITIATING",
			"deleteRequested": false
		}`)
	})

	interfaceEndpoint, _, err := client.PrivateEndpoints.AddOnePrivateEndpoint(ctx, groupID, "AZURE", privateLinkID, request)
	if err != nil {
		t.Errorf("PrivateEndpoints.AddOnePrivateEndpoint returned error: %v", err)
		return
	}

	expected := &InterfaceEndpointConnection{
		PrivateEndpointIPAddress:  "10.0.0.4",
		PrivateEndpointResourceID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/privatelink/providers/Microsoft.Network/privateEndpoints/test",
		ConnectionStatus:          "INITIATING",
		DeleteRequested:           pointy.Bool(false),
	}

	if !reflect.DeepEqual(interfaceEndpoint, expected) {
		t.Errorf("PrivateEndpoints.AddOnePrivateEndpoint\n got=%#v\nwant=%#v", interfaceEndpoint, expected)
	}
}

func TestPrivateEndpoints_GetOneInterfaceEndpointAWS(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	interfaceEndpointID := "vpce-0b9c5701325cb15dd"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService/%s/endpoint/%s", groupID, "AWS", privateLinkID, interfaceEndpointID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"interfaceEndpointId": "vpce-08fb7e9319909ec7b",
			"connectionStatus": "PENDING",
			"deleteRequested": false
		}`)
	})

	interfaceEndpoint, _, err := client.PrivateEndpoints.GetOnePrivateEndpoint(ctx, groupID, "AWS", privateLinkID, interfaceEndpointID)
	if err != nil {
		t.Errorf("PrivateEndpoints.GetOnePrivateEndpoint returned error: %v", err)
	}

	expected := &InterfaceEndpointConnection{
		InterfaceEndpointID: "vpce-08fb7e9319909ec7b",
		ConnectionStatus:    "PENDING",
		DeleteRequested:     pointy.Bool(false),
	}

	if !reflect.DeepEqual(interfaceEndpoint, expected) {
		t.Errorf("PrivateEndpoints.GetOnePrivateEndpoint\n got=%#v\nwant=%#v", interfaceEndpoint, expected)
	}
}

func TestPrivateEndpoints_GetOneInterfaceEndpointAzure(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	interfaceEndpointID := "/subscriptions/19265c27-b60e-4c3b-9426-ae3f507300b5/resourceGroups/test/providers/Microsoft.Network/privateEndpoints/test"

	path := fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService/%s/endpoint%s", groupID, "AZURE", privateLinkID, interfaceEndpointID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"privateEndpointResourceId": "/subscriptions/19265c27-b60e-4c3b-9426-ae3f507300b5/resourceGroups/test/providers/Microsoft.Network/privateEndpoints/test",
			"privateEndpointIPAddress": "10.0.0.4",
			"connectionStatus": "INITIATING",
			"deleteRequested": false
		}`)
	})

	interfaceEndpoint, _, err := client.PrivateEndpoints.GetOnePrivateEndpoint(ctx, groupID, "AZURE", privateLinkID, interfaceEndpointID)
	if err != nil {
		t.Errorf("PrivateEndpoints.GetOnePrivateEndpoint returned error: %v", err)
	}

	expected := &InterfaceEndpointConnection{
		PrivateEndpointIPAddress:  "10.0.0.4",
		PrivateEndpointResourceID: interfaceEndpointID,
		ConnectionStatus:          "INITIATING",
		DeleteRequested:           pointy.Bool(false),
	}

	if !reflect.DeepEqual(interfaceEndpoint, expected) {
		t.Errorf("PrivateEndpoints.GetOnePrivateEndpoint\n got=%#v\nwant=%#v", interfaceEndpoint, expected)
	}
}

func TestPrivateEndpoints_DeleteOneInterfaceEndpointAWS(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	interfaceEndpointID := "vpce-0b9c5701325cb15dd"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService/%s/endpoint/%s", groupID, "AWS", privateLinkID, interfaceEndpointID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PrivateEndpoints.DeleteOnePrivateEndpoint(ctx, groupID, "AWS", privateLinkID, interfaceEndpointID)
	if err != nil {
		t.Errorf("PrivateEndpoints.DeleteOnePrivateEndpoint returned error: %v", err)
	}
}

func TestPrivateEndpoints_DeleteOneInterfaceEndpointAzure(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	privateLinkID := "5df264b8f10fab7d2cad2f0d"
	interfaceEndpointID := "vpce-0b9c5701325cb15dd"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/privateEndpoint/%s/endpointService/%s/endpoint/%s", groupID, "AZURE", privateLinkID, interfaceEndpointID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PrivateEndpoints.DeleteOnePrivateEndpoint(ctx, groupID, "AZURE", privateLinkID, interfaceEndpointID)
	if err != nil {
		t.Errorf("PrivateEndpoints.DeleteOnePrivateEndpoint returned error: %v", err)
	}
}
