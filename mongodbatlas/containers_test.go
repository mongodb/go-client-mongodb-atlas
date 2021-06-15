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
	"testing"

	"github.com/go-test/deep"
	"github.com/openlyinc/pointy"
)

func TestContainers_List(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()
		mux.HandleFunc("/api/atlas/v1.0/groups/1/containers", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `{
			"results": [
				{
					"atlasCidrBlock": "10.8.0.0/18",
					"id": "1112269b3bf99403840e8934",
					"gcpProjectId": "my-sample-project-191923",
					"networkName": "test1",
					"providerName": "GCP",
					"provisioned": true
				},
				{
					"atlasCidrBlock" : "10.8.0.0/21",
					"id" : "1112269b3bf99403840e8934",
					"providerName" : "AWS",
					"provisioned" : true,
					"regionName" : "US_EAST_1",
					"vpcId" : "vpc-zz0zzzzz"
				}
			],
			"totalCount": 2
		}`)
		})

		containers, _, err := client.Containers.List(ctx, "1", nil)
		if err != nil {
			t.Fatalf("Containers.List returned error: %v", err)
		}

		GCPContainer := Container{
			AtlasCIDRBlock: "10.8.0.0/18",
			ID:             "1112269b3bf99403840e8934",
			GCPProjectID:   "my-sample-project-191923",
			NetworkName:    "test1",
			ProviderName:   "GCP",
			Provisioned:    pointy.Bool(true),
		}

		AWSContainer := Container{
			AtlasCIDRBlock: "10.8.0.0/21",
			ID:             "1112269b3bf99403840e8934",
			ProviderName:   "AWS",
			Provisioned:    pointy.Bool(true),
			RegionName:     "US_EAST_1",
			VPCID:          "vpc-zz0zzzzz",
		}

		expected := []Container{GCPContainer, AWSContainer}

		if diff := deep.Equal(containers, expected); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("multiple pages", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()
		mux.HandleFunc("/api/atlas/v1.0/groups/1/containers", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)

			dr := containersResponse{
				Results: []Container{
					{
						AtlasCIDRBlock: "10.8.0.0/18",
						ID:             "1112269b3bf99403840e8934",
						GCPProjectID:   "my-sample-project-191923",
						NetworkName:    "test1",
						ProviderName:   "GCP",
						Provisioned:    pointy.Bool(true),
					},
					{
						AtlasCIDRBlock: "10.8.0.0/21",
						ID:             "1112269b3bf99403840e8934",
						ProviderName:   "AWS",
						Provisioned:    pointy.Bool(true),
						RegionName:     "US_EAST_1",
						VPCID:          "vpc-zz0zzzzz",
					},
				},
				Links: []*Link{
					{Href: "http://example.com/api/atlas/v1.0/groups/1/containers?pageNum=2&itemsPerPage=2", Rel: "self"},
					{Href: "http://example.com/api/atlas/v1.0/groups/1/containers?pageNum=2&itemsPerPage=2", Rel: "previous"},
				},
			}

			b, err := json.Marshal(dr)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Fprint(w, string(b))
		})

		_, resp, err := client.Containers.List(ctx, "1", nil)
		if err != nil {
			t.Fatal(err)
		}

		checkCurrentPage(t, resp, 2)
	})
	t.Run("by page number", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()
		jBlob := `
	{
		"links": [
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/containers?pageNum=1&itemsPerPage=1",
				"rel": "previous"
			},
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/containers?pageNum=2&itemsPerPage=1",
				"rel": "self"
			},
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/containers?itemsPerPage=3&pageNum=2",
				"rel": "next"
			}
		],
		"results": [
			{
				"atlasCidrBlock": "10.8.0.0/18",
				"id": "1112269b3bf99403840e8934",
				"gcpProjectId": "my-sample-project-191923",
				"networkName": "test1",
				"providerName": "GCP",
				"provisioned": true
			}
		],
		"totalCount": 3
	}`

		mux.HandleFunc("/api/atlas/v1.0/groups/1/containers", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)

			fmt.Fprint(w, jBlob)
		})

		opt := &ContainersListOptions{ListOptions: ListOptions{PageNum: 2}, ProviderName: "GCP"}
		_, resp, err := client.Containers.List(ctx, "1", opt)

		if err != nil {
			t.Fatal(err)
		}

		checkCurrentPage(t, resp, 2)
	})
}

func TestContainers_ListAll(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/atlas/v1.0/groups/1/containers/all", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"results": [
				{
					"atlasCidrBlock": "10.8.0.0/18",
					"id": "1112269b3bf99403840e8934",
					"gcpProjectId": "my-sample-project-191923",
					"networkName": "test1",
					"providerName": "GCP",
					"provisioned": true
				},
				{
					"atlasCidrBlock" : "10.8.0.0/21",
					"id" : "1112269b3bf99403840e8934",
					"providerName" : "AWS",
					"provisioned" : true,
					"regionName" : "US_EAST_1",
					"vpcId" : "vpc-zz0zzzzz"
				}
			],
			"totalCount": 2
		}`)
	})

	containers, _, err := client.Containers.ListAll(ctx, "1", nil)
	if err != nil {
		t.Fatalf("Containers.List returned error: %v", err)
	}

	GCPContainer := Container{
		AtlasCIDRBlock: "10.8.0.0/18",
		ID:             "1112269b3bf99403840e8934",
		GCPProjectID:   "my-sample-project-191923",
		NetworkName:    "test1",
		ProviderName:   "GCP",
		Provisioned:    pointy.Bool(true),
	}

	AWSContainer := Container{
		AtlasCIDRBlock: "10.8.0.0/21",
		ID:             "1112269b3bf99403840e8934",
		ProviderName:   "AWS",
		Provisioned:    pointy.Bool(true),
		RegionName:     "US_EAST_1",
		VPCID:          "vpc-zz0zzzzz",
	}

	expected := []Container{GCPContainer, AWSContainer}

	if diff := deep.Equal(containers, expected); diff != nil {
		t.Error(diff)
	}
}

func TestContainers_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	createRequest := &Container{
		AtlasCIDRBlock: "10.8.0.0/18",
		ProviderName:   "GCP",
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/containers", groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"atlasCidrBlock": "10.8.0.0/18",
			"providerName":   "GCP",
		}

		jsonBlob := `
		{
			"atlasCidrBlock": "10.8.0.0/18",
			"id": "1112269b3bf99403840e8934",
			"gcpProjectId": "my-sample-project-191923",
			"networkName": "test1",
			"providerName": "GCP",
			"provisioned": true
		}
		`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}
		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, jsonBlob)
	})

	container, _, err := client.Containers.Create(ctx, groupID, createRequest)
	if err != nil {
		t.Fatalf("Containers.Create returned error: %v", err)
	}

	expectedCID := "1112269b3bf99403840e8934"
	expectedGCPID := "my-sample-project-191923"

	if cID := container.ID; cID != expectedCID {
		t.Errorf("expected containerId '%s', received '%s'", expectedCID, cID)
	}

	if id := container.GCPProjectID; id != expectedGCPID {
		t.Errorf("expected gcpProjectId '%s', received '%s'", expectedGCPID, id)
	}

	if isProvisioned := container.Provisioned; !*isProvisioned {
		t.Errorf("expected provisioned '%t', received '%t'", true, *isProvisioned)
	}
}

func TestContainers_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	containerID := "1112269b3bf99403840e8934"

	updateRequest := &Container{
		AtlasCIDRBlock: "10.8.0.0/18",
		GCPProjectID:   "my-sample-project-191923",
		ProviderName:   "GCP",
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/containers/%s", groupID, containerID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"atlasCidrBlock": "10.8.0.0/18",
			"gcpProjectId":   "my-sample-project-191923",
			"providerName":   "GCP",
		}

		jsonBlob := `
		{
			"atlasCidrBlock": "10.8.0.0/18",
			"id": "1112269b3bf99403840e8934",
			"gcpProjectId": "my-sample-project-191923",
			"networkName": "test1",
			"providerName": "GCP",
			"provisioned": true
		}
		`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, jsonBlob)
	})

	container, _, err := client.Containers.Update(ctx, groupID, containerID, updateRequest)
	if err != nil {
		t.Fatalf("Containers.Update returned error: %v", err)
	}

	if cID := container.ID; cID != containerID {
		t.Errorf("expected containerID '%s', received '%s'", containerID, cID)
	}

	expectedGCPID := "my-sample-project-191923"

	if id := container.GCPProjectID; id != expectedGCPID {
		t.Errorf("expected gcpProjectId '%s', received '%s'", expectedGCPID, id)
	}

	if isProvisioned := container.Provisioned; !*isProvisioned {
		t.Errorf("expected provisioned '%t', received '%t'", true, *isProvisioned)
	}
}

func TestContainers_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	id := "1112222b3bf99403840e8934"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/containers/%s", groupID, id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Containers.Delete(ctx, groupID, id)
	if err != nil {
		t.Fatalf("Container.Delete returned error: %v", err)
	}
}

func TestContainersServiceOp_Create(t *testing.T) {
	type args struct {
		request          *Container
		expectedRequest  map[string]interface{}
		response         string
		expectedResponse *Container
	}
	groupID := "1"
	tests := map[string]args{
		"AWS": {
			request: &Container{
				AtlasCIDRBlock: "10.8.0.0/21",
				RegionName:     "US_EAST_1",
				ProviderName:   "AWS",
			},
			expectedRequest: map[string]interface{}{
				"atlasCidrBlock": "10.8.0.0/21",
				"regionName":     "US_EAST_1",
				"providerName":   "AWS",
			},
			response: `{
  "atlasCidrBlock" : "10.8.0.0/21",
  "id" : "1112269b3bf99403840e8934",
  "provisioned" : true,
  "regionName" : "US_EAST_1",
  "vpcId" : "vpc-zz0zzzzz"
}`,
			expectedResponse: &Container{
				AtlasCIDRBlock: "10.8.0.0/21",
				ID:             "1112269b3bf99403840e8934",
				Provisioned:    pointy.Bool(true),
				RegionName:     "US_EAST_1",
				VPCID:          "vpc-zz0zzzzz",
			},
		},
		"AZURE": {
			request: &Container{
				AtlasCIDRBlock: "10.8.0.0/21",
				Region:         "US_EAST_2",
				ProviderName:   "AZURE",
			},
			expectedRequest: map[string]interface{}{
				"atlasCidrBlock": "10.8.0.0/21",
				"region":         "US_EAST_2",
				"providerName":   "AZURE",
			},
			response: `{
  "atlasCidrBlock":"10.8.0.0/21",
  "azureSubscriptionId":"602336d43d098d433845971g",
  "id":"5cbf563d87d9d67253be590a",
  "providerName":"AZURE",
  "provisioned":false,
  "region":"US_EAST_2",
  "vnetName":null
}`,
			expectedResponse: &Container{
				AtlasCIDRBlock:      "10.8.0.0/21",
				AzureSubscriptionID: "602336d43d098d433845971g",
				ID:                  "5cbf563d87d9d67253be590a",
				Provisioned:         pointy.Bool(false),
				ProviderName:        "AZURE",
				Region:              "US_EAST_2",
			},
		},
		"GCP": {
			request: &Container{
				AtlasCIDRBlock: "10.8.0.0/18",
				ProviderName:   "GCP",
			},
			expectedRequest: map[string]interface{}{
				"atlasCidrBlock": "10.8.0.0/18",
				"providerName":   "GCP",
			},
			response: `{
  "atlasCidrBlock" : "10.8.0.0/18",
  "id" : "1112269b3bf99403840e8934",
  "gcpProjectId" : "null",
  "networkName" : "null",
  "provisioned" : true
}`,
			expectedResponse: &Container{
				AtlasCIDRBlock: "10.8.0.0/18",
				ID:             "1112269b3bf99403840e8934",
				Provisioned:    pointy.Bool(true),
				NetworkName:    "null",
				GCPProjectID:   "null",
			},
		},
	}
	for test, args := range tests {
		request := args.request
		expectedRequest := args.expectedRequest
		response := args.response
		expectedResponse := args.expectedResponse
		t.Run(test, func(t *testing.T) {
			client, mux, teardown := setup()
			defer teardown()
			mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/containers", groupID), func(w http.ResponseWriter, r *http.Request) {
				expected := expectedRequest
				jsonBlob := response

				var v map[string]interface{}
				err := json.NewDecoder(r.Body).Decode(&v)
				if err != nil {
					t.Fatalf("decode json: %v", err)
				}

				if diff := deep.Equal(v, expected); diff != nil {
					t.Error(diff)
				}

				fmt.Fprint(w, jsonBlob)
			})
			container, _, err := client.Containers.Create(ctx, groupID, request)
			if err != nil {
				t.Fatalf("Containers.Update returned error: %v", err)
			}
			if diff := deep.Equal(container, expectedResponse); diff != nil {
				t.Error(diff)
			}
		})
	}
}
