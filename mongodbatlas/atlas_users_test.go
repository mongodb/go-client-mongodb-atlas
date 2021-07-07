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
)

func TestAtlasUsers_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups/1/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/1/users?pretty=true&pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				{
					"emailAddress": "joe.bloggs@example.com",
					"firstName": "Joe",
					"id": "1",
					"lastName": "Bloggs",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/1",
							"rel": "self"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/1/whitelist",
							"rel": "http://mms.mongodb.com/whitelist"
						}
					],
					"roles": [
						{
							"groupId": "1",
							"roleName": "GROUP_OWNER"
						},
						{
							"groupId": "2",
							"roleName": "GROUP_OWNER"
						}
					],
					"username": "joe.bloggs"
				},
				{
					"emailAddress": "jim.bloggs@example.com",
					"firstName": "Jim",
					"id": "2",
					"lastName": "Bloggs",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/2",
							"rel": "self"
						},
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/2/whitelist",
							"rel": "http://mms.mongodb.com/whitelist"
						}
					],
					"roles": [
						{
							"roleName": "GLOBAL_READ_ONLY"
						},
						{
							"groupId": "1",
							"roleName": "GROUP_OWNER"
						}
					],
					"username": "jim.bloggs"
				}
			],
			"totalCount": 2
		}`)
	})

	teams, _, err := client.AtlasUsers.List(ctx, "1", nil)

	if err != nil {
		t.Fatalf("AtlasUsers.List returned error: %v", err)
	}

	expected := []AtlasUser{
		{
			ID:           "1",
			EmailAddress: "joe.bloggs@example.com",
			FirstName:    "Joe",
			LastName:     "Bloggs",
			Roles: []AtlasRole{
				{
					GroupID:  "1",
					RoleName: "GROUP_OWNER",
				},
				{
					GroupID:  "2",
					RoleName: "GROUP_OWNER",
				},
			},
			Username: "joe.bloggs",
		},
		{
			EmailAddress: "jim.bloggs@example.com",
			FirstName:    "Jim",
			ID:           "2",
			LastName:     "Bloggs",
			Roles: []AtlasRole{
				{
					RoleName: "GLOBAL_READ_ONLY",
				},
				{
					GroupID:  "1",
					RoleName: "GROUP_OWNER",
				},
			},
			Username: "jim.bloggs",
		},
	}
	if diff := deep.Equal(teams, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAtlasUsers_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	userID := "5af1c27a0a7fa48c76d3a761"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/users/%s", userID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"emailAddress": "john.doe@example.com",
			"firstName": "John",
			"id": "5af1c27a0a7fa48c76d3a761",
			"lastName": "Doe",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/5af1c27a0a7fa48c76d3a761",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/5af1c27a0a7fa48c76d3a761/whitelist",
					"rel": "http://mms.mongodb.com/whitelist"
				}
			],
			"mobileNumber" : "2125550198",
			"roles": [
				{
					"orgId": "5af1c27a0a7fa48c76d3a762",
					"roleName": "ORG_OWNER"
				},
				{
					"groupId": "5af1c27a0a7fa48c76d3a763",
					"roleName": "GROUP_OWNER"
				}
			],
			"teamIds": [
				"5af1c27a0a7fa48c76d3a764"
			],
			"username": "john.doe@example.com"
		}`)
	})

	team, _, err := client.AtlasUsers.Get(ctx, userID)
	if err != nil {
		t.Fatalf("AtlasUsers.Get returned error: %v", err)
	}

	expected := &AtlasUser{
		ID:           "5af1c27a0a7fa48c76d3a761",
		EmailAddress: "john.doe@example.com",
		FirstName:    "John",
		LastName:     "Doe",
		MobileNumber: "2125550198",
		Roles: []AtlasRole{
			{
				OrgID:    "5af1c27a0a7fa48c76d3a762",
				RoleName: "ORG_OWNER",
			},
			{
				GroupID:  "5af1c27a0a7fa48c76d3a763",
				RoleName: "GROUP_OWNER",
			},
		},
		TeamIds: []string{
			"5af1c27a0a7fa48c76d3a764",
		},
		Username: "john.doe@example.com",
	}

	if diff := deep.Equal(team, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAtlasUsers_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	username := "john.doe@example.com"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/users/byName/%s", username), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"emailAddress": "john.doe@example.com",
			"firstName": "John",
			"id": "5af1c27a0a7fa48c76d3a761",
			"lastName": "Doe",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/5af1c27a0a7fa48c76d3a761",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/5af1c27a0a7fa48c76d3a761/whitelist",
					"rel": "http://mms.mongodb.com/whitelist"
				}
			],
			"mobileNumber" : "2125550198",
			"roles": [
				{
					"orgId": "5af1c27a0a7fa48c76d3a762",
					"roleName": "ORG_OWNER"
				},
				{
					"groupId": "5af1c27a0a7fa48c76d3a763",
					"roleName": "GROUP_OWNER"
				}
			],
			"teamIds": [
				"5af1c27a0a7fa48c76d3a764"
			],
			"username": "john.doe@example.com"
		}`)
	})

	team, _, err := client.AtlasUsers.GetByName(ctx, username)
	if err != nil {
		t.Fatalf("AtlasUsers.GetByName returned error: %v", err)
	}

	expected := &AtlasUser{
		ID:           "5af1c27a0a7fa48c76d3a761",
		EmailAddress: "john.doe@example.com",
		FirstName:    "John",
		LastName:     "Doe",
		MobileNumber: "2125550198",
		Roles: []AtlasRole{
			{
				OrgID:    "5af1c27a0a7fa48c76d3a762",
				RoleName: "ORG_OWNER",
			},
			{
				GroupID:  "5af1c27a0a7fa48c76d3a763",
				RoleName: "GROUP_OWNER",
			},
		},
		TeamIds: []string{
			"5af1c27a0a7fa48c76d3a764",
		},
		Username: "john.doe@example.com",
	}

	if diff := deep.Equal(team, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAtlasUsers_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &AtlasUser{

		Username:     "john.doe@example.com",
		Password:     "myPassword1@",
		EmailAddress: "john.doe@example.com",
		MobileNumber: "2125550198",
		FirstName:    "John",
		LastName:     "Doe",
		Roles: []AtlasRole{
			{
				OrgID:    "8dbbe4570bd55b23f25444db",
				RoleName: "ORG_MEMBER",
			},
			{
				GroupID:  "2ddoa1233ef88z75f64578ff",
				RoleName: "GROUP_READ_ONLY",
			},
		},
		Country: "US",
	}

	mux.HandleFunc("/api/atlas/v1.0/users", func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"username":     "john.doe@example.com",
			"password":     "myPassword1@",
			"emailAddress": "john.doe@example.com",
			"mobileNumber": "2125550198",
			"firstName":    "John",
			"lastName":     "Doe",
			"roles": []interface{}{
				map[string]interface{}{
					"orgId":    "8dbbe4570bd55b23f25444db",
					"roleName": "ORG_MEMBER",
				},
				map[string]interface{}{
					"groupId":  "2ddoa1233ef88z75f64578ff",
					"roleName": "GROUP_READ_ONLY",
				},
			},
			"country": "US",
		}

		jsonBlob := `
		{
			"emailAddress": "john.doe@example.com",
			"firstName": "John",
			"id": "5b06ed7083fb5a40df86e93b",
			"lastName": "Doe",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/5b06ed7083fb5a40df86e93b",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/5b06ed7083fb5a40df86e93b/whitelist",
					"rel": "http://mms.mongodb.com/whitelist"
				}
			],
			"mobileNumber" : "2125550198",
			"roles": [
			  {
				"orgId" : "8dbbe4570bd55b23f25444db",
				"roleName" : "ORG_MEMBER"
		
			  },
			  {
				"groupId" : "2ddoa1233ef88z75f64578ff",
				"roleName" : "GROUP_READ_ONLY"
		
			  }
			],
			"teamIds": [],
			"username": "john.doe@example.com"
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

	atlasUser, _, err := client.AtlasUsers.Create(ctx, createRequest)
	if err != nil {
		t.Fatalf("AtlasUsers.Create returned error: %v", err)
	}

	if name := atlasUser.Username; name != "john.doe@example.com" {
		t.Errorf("expected name '%s', received '%s'", "john.doe@example.com", name)
	}

	if id := atlasUser.ID; id != "5b06ed7083fb5a40df86e93b" {
		t.Errorf("expected id '%s', received '%s'", "5b06ed7083fb5a40df86e93b", id)
	}

	if roles := len(atlasUser.Roles); roles != 2 {
		t.Errorf("expected len(roles) '%d', received '%d'", 2, roles)
	}
}
