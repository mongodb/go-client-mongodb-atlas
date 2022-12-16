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

func TestRoot_ListRoot(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"apiKey": {
			  "accessList": [
				{
				  "cidrBlock": "0.0.0.0/1",
				  "ipAddress": null
				},
				{
				  "cidrBlock": "128.0.0.0/1",
				  "ipAddress": null
				},
				{
				  "cidrBlock": "47.225.212.178/32",
				  "ipAddress": "47.225.212.178"
				}
			  ],
			  "id": "5c47503320eef5699e1cce8f",
			  "publicKey": "ewmaqvde",
			  "roles": [
				{
				  "orgId": "5c47503320eef5699e1cce8f",
				  "roleName": "ORG_OWNER"
				},
				{
				  "orgId": "5c47503320eef5699e1cce8f",
				  "roleName": "ORG_MEMBER"
				},
				{
				  "orgId": "5c47503320eef5699e1cce8f",
				  "roleName": "ORG_GROUP_CREATOR"
				},
				{
				  "groupId": "5c47503320eef5699e1cce8f",
				  "roleName": "GROUP_OWNER"
				}
			  ]
			},
			"appName": "MongoDB Atlas",
			"build": "a646b0cf64fe70b213d206768900eb5e4982a476",
			"links": [
			  {
				"href": "https://cloud.mongodb.com/api/atlas/v1.0",
				"rel": "self"
			  },
			  {
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5c47503320eef5699e1cce8f/apiKeys/5c47503320eef5699e1cce8f",
				"rel": "https://cloud.mongodb.com/apiKeys"
			  }
			],
			"throttling": false
		  }`)
	})

	apiKeys, _, err := client.Root.List(ctx, nil)

	if err != nil {
		t.Fatalf("APIKeys.List returned error: %v", err)
	}

	expected := &Root{
		APIKey: struct {
			AccessList []struct {
				CIDRBlock string `json:"cidrBlock"`
				IPAddress string `json:"ipAddress"`
			} `json:"accessList"`
			ID        string      `json:"id"`
			PublicKey string      `json:"publicKey"`
			Roles     []AtlasRole `json:"roles,omitempty"`
		}{
			AccessList: []struct {
				CIDRBlock string `json:"cidrBlock"`
				IPAddress string `json:"ipAddress"`
			}{
				{IPAddress: "", CIDRBlock: "0.0.0.0/1"},
				{IPAddress: "", CIDRBlock: "128.0.0.0/1"},
				{IPAddress: "47.225.212.178", CIDRBlock: "47.225.212.178/32"},
			},
			ID:        "5c47503320eef5699e1cce8f",
			PublicKey: "ewmaqvde",
			Roles: []AtlasRole{
				{
					OrgID:    "5c47503320eef5699e1cce8f",
					RoleName: "ORG_OWNER",
				},
				{
					OrgID:    "5c47503320eef5699e1cce8f",
					RoleName: "ORG_MEMBER",
				},
				{
					OrgID:    "5c47503320eef5699e1cce8f",
					RoleName: "ORG_GROUP_CREATOR",
				},
				{
					GroupID:  "5c47503320eef5699e1cce8f",
					RoleName: "GROUP_OWNER",
				},
			},
		},
		AppName: "MongoDB Atlas",
		Build:   "a646b0cf64fe70b213d206768900eb5e4982a476",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0",
				Rel:  "self",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5c47503320eef5699e1cce8f/apiKeys/5c47503320eef5699e1cce8f",
				Rel:  "https://cloud.mongodb.com/apiKeys",
			},
		},
		Throttling: false,
	}

	if diff := deep.Equal(apiKeys, expected); diff != nil {
		t.Error(diff)
	}
}
