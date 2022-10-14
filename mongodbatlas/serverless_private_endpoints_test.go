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
	"reflect"
	"testing"
)

func TestPrivateEndpoints_GetOnePrivateServerlessEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	instanceName := "serverless"
	endpointID := "3658569708087079"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateEndpoint/serverless/instance/%s/endpoint/%s", groupID, instanceName, endpointID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"ID":"", "CloudProviderEndpointID":"vpce-12356", "Comment":"Test Serverless", "EndpointServiceName":"", "Status":"AVAILABLE", "ProviderName":"AWS"
		}`)
	})

	interfaceEndpoint, _, err := client.ServerlessPrivateEndpoints.Get(ctx, groupID, instanceName, endpointID)
	if err != nil {
		t.Errorf("PrivateEndpoints.Get returned error: %v", err)
	}

	expected := &ServerlessPrivateEndpointConnection{
		Comment:                 "Test Serverless",
		ProviderName:            "AWS",
		Status:                  "AVAILABLE",
		CloudProviderEndpointID: "vpce-12356",
	}

	if !reflect.DeepEqual(interfaceEndpoint, expected) {
		t.Errorf("PrivateEndpoints.Get \n got=%#v\nwant=%#v", interfaceEndpoint, expected)
	}
}

func TestPrivateEndpoints_UpdateOnePrivateServerlessEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	instanceName := "serverless"
	endpointID := "3658569708087079"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/privateEndpoint/serverless/instance/%s/endpoint/%s", groupID, instanceName, endpointID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
			"ID":"", "CloudProviderEndpointID":"vpce1234567", "Comment":"Test Serverless", "EndpointServiceName":"", "Status":"AVAILABLE", "ProviderName":"AWS"
		}`)
	})

	update := &ServerlessPrivateEndpointConnection{
		CloudProviderEndpointID: "vpce1234567",
		ProviderName:            "AWS",
	}

	interfaceEndpoint, _, err := client.ServerlessPrivateEndpoints.Update(ctx, groupID, instanceName, endpointID, update)
	if err != nil {
		t.Errorf("PrivateEndpoints.Update returned error: %v", err)
	}

	expected := &ServerlessPrivateEndpointConnection{
		Comment:                 "Test Serverless",
		ProviderName:            "AWS",
		Status:                  "AVAILABLE",
		CloudProviderEndpointID: "vpce1234567",
	}

	if !reflect.DeepEqual(interfaceEndpoint, expected) {
		t.Errorf("PrivateEndpoints.Update\n got=%#v\nwant=%#v", interfaceEndpoint, expected)
	}
}
