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

func TestTeams_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/orgs/1/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [{
				   "href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/7253g4s820thk69bfecd7bv1/teams?pageNum=1&itemsPerPage=100",
				   "rel": "self"
			}],
			"results": [{
					"id": "b5cfec387d9d63926b9189g",
					"name": "Finance"
			}, {
					"id": "6b610e4f80eef5366613e4df",
					"name": "Research and Development"
			}, {
					"id": "6b610e1087d9d66b272f0c86",
					"name": "Technical Documentation"
			}],
			"totalCount": 3
		}`)
	})

	teams, _, err := client.Teams.List(ctx, "1", nil)

	if err != nil {
		t.Fatalf("Teams.List returned error: %v", err)
	}

	expected := []Team{
		{
			ID:   "b5cfec387d9d63926b9189g",
			Name: "Finance",
		},
		{
			ID:   "6b610e4f80eef5366613e4df",
			Name: "Research and Development",
		},
		{
			ID:   "6b610e1087d9d66b272f0c86",
			Name: "Technical Documentation",
		},
	}

	if diff := deep.Equal(teams, expected); diff != nil {
		t.Error(diff)
	}
}

func TestTeams_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "1"
	teamID := "1"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/teams/%s", orgID, teamID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": "1",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/1/teams/1",
					"rel": "self"
				}
			],
			"name": "myNewTeam"
		   }`)
	})

	team, _, err := client.Teams.Get(ctx, orgID, teamID)
	if err != nil {
		t.Fatalf("Teams.Get returned error: %v", err)
	}

	expected := &Team{
		ID:   "1",
		Name: "myNewTeam",
	}

	if diff := deep.Equal(team, expected); diff != nil {
		t.Error(diff)
	}
}

func TestTeams_GetOneTeamByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "1"
	teamName := "myNewTeam"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/teams/byName/%s", orgID, teamName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": "1",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/1/teams/1",
					"rel": "self"
				}
			],
			"name": "myNewTeam"
		   }`)
	})

	team, _, err := client.Teams.GetOneTeamByName(ctx, orgID, teamName)
	if err != nil {
		t.Fatalf("Teams.Get returned error: %v", err)
	}

	expected := &Team{
		ID:   "1",
		Name: "myNewTeam",
	}

	if diff := deep.Equal(team, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProject_GetTeamUsersAssigned(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "5a0a1e7e0f2912c554080adc"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/teams/1/users", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/teams/{TEAM-ID}/users?pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				{
					"emailAddress": "AtlasUser@example.com",
					"firstName": "Atlas",
					"id": "{USER-ID}",
					"lastName": "User",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/{USER-ID}",
							"rel": "self"
						}
					],
					"roles": [
						{
							"groupId": "{GROUP-ID}",
							"roleName": "GROUP_OWNER"
						},
						{
							"orgId": "{ORG-ID}",
							"roleName": "ORG_OWNER"
						}
					],
					"teamIds": [
						"{TEAM-ID}"
					],
					"username": "AtlasUser@example.com"
				}
			],
			"totalCount": 1
		   }`)
	})

	teamsAssigned, _, err := client.Teams.GetTeamUsersAssigned(ctx, orgID, "1")
	if err != nil {
		t.Fatalf("Projects.GetTeamUsersAssigned returned error: %v", err)
	}

	expected := []AtlasUser{
		{
			EmailAddress: "AtlasUser@example.com",
			FirstName:    "Atlas",
			ID:           "{USER-ID}",
			LastName:     "User",

			Roles: []AtlasRole{
				{
					GroupID:  "{GROUP-ID}",
					RoleName: "GROUP_OWNER",
				},
				{
					OrgID:    "{ORG-ID}",
					RoleName: "ORG_OWNER",
				},
			},
			TeamIds: []string{
				"{TEAM-ID}",
			},
			Username: "AtlasUser@example.com",
		},
	}

	if diff := deep.Equal(teamsAssigned, expected); diff != nil {
		t.Error(diff)
	}
}

func TestTeams_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "1"

	createRequest := &Team{
		Name:      "myNewTeam",
		Usernames: []string{"user1", "user2", "user3"},
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/teams", orgID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"name":      "myNewTeam",
			"usernames": []interface{}{"user1", "user2", "user3"},
		}

		jsonBlob := `
		{
			"id": "1",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/teams/{TEAM-ID}",
					"rel": "self"
				}
			],
			"name": "myNewTeam",
			"usernames": ["user1", "user2", "user3"]
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

	team, _, err := client.Teams.Create(ctx, orgID, createRequest)
	if err != nil {
		t.Fatalf("Teams.Create returned error: %v", err)
	}

	if name := team.Name; name != "myNewTeam" {
		t.Errorf("expected name '%s', received '%s'", "myNewTeam", name)
	}

	if id := team.ID; id != "1" {
		t.Errorf("expected id '%s', received '%s'", "1", id)
	}

	if usernames := len(team.Usernames); usernames != 3 {
		t.Errorf("expected len(usernames) '%d', received '%d'", 3, usernames)
	}
}

func TestTeams_Rename(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "1"

	renameRequest := "newTeamName"
	teamID := "6b720e1087d9d66b272f1c86"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/teams/%s", orgID, teamID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"name": renameRequest,
		}

		jsonBlob := `
		{
			"id" : "6b720e1087d9d66b272f1c86",
			"links" : [ {
			  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5991f2c580eef55aedbc6aa0/teams/6b720e1087d9d66b272f1c86",
			  "rel" : "self"
			} ],
			"name" : "newTeamName"
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

	team, _, err := client.Teams.Rename(ctx, orgID, teamID, renameRequest)
	if err != nil {
		t.Fatalf("Teams.Rename returned error: %v", err)
	}

	if name := team.Name; name != renameRequest {
		t.Errorf("expected name '%s', received '%s'", renameRequest, name)
	}

	if id := team.ID; id != teamID {
		t.Errorf("expected id '%s', received '%s'", teamID, id)
	}
}

func TestTeams_UpdateTeamRoles(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "1"

	updateTeamRolesRequest := &TeamUpdateRoles{
		RoleNames: []string{"GROUP_OWNER"},
	}

	teamID := "6b720e1087d9d66b272f1c86"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/teams/%s", orgID, teamID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"roleNames": []interface{}{"GROUP_OWNER"},
		}

		jsonBlob := `
		{
			"links": [{
			  "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{PROJECT-ID}/teams/{TEAM-ID3}?pretty=true&pageNum=1&itemsPerPage=100",
			  "rel": "self"
			}],
			"results": [{
			  "links": [{
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{PROJECT-ID}/teams/{TEAM-ID1}",
				"rel": "self"
			  }],
			  "roleNames": ["GROUP_OWNER", "GROUP_DATA_ACCESS_READ_ONLY", "GROUP_DATA_ACCESS_ADMIN", "GROUP_DATA_ACCESS_READ_WRITE", "GROUP_READ_ONLY"],
			  "teamId": "6b720e1087d9d66b272f1c86"
			}, {
			  "links": [{
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{PROJECT-ID}/teams/{TEAM-ID2}",
				"rel": "self"
			  }],
			  "roleNames": ["GROUP_DATA_ACCESS_ADMIN", "GROUP_READ_ONLY"],
			  "teamId": "6b720e1087d9d66b272f1c86"
			}, {
			  "links": [{
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{PROJECT-ID}/teams/{TEAM-ID3}",
				"rel": "self"
			  }],
			  "roleNames": ["GROUP_OWNER"],
			  "teamId": "6b720e1087d9d66b272f1c86"
			}],
			"totalCount": 3
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

	teams, _, err := client.Teams.UpdateTeamRoles(ctx, orgID, teamID, updateTeamRolesRequest)
	if err != nil {
		t.Fatalf("Teams.UpdateTeamRoles returned error: %v", err)
	}

	if teamCount := len(teams); teamCount != 3 {
		t.Errorf("expected teamCount '%d', received '%d'", 3, teamCount)
	}

	if id := teams[0].TeamID; id != teamID {
		t.Errorf("expected id '%s', received '%s'", teamID, id)
	}
}

func TestTeams_AddUsersToTeam(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "1"

	usersID := []string{"1", "2"}
	teamID := "6b720e1087d9d66b272f1c86"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/teams/%s/users", orgID, teamID), func(w http.ResponseWriter, r *http.Request) {
		expected := []map[string]interface{}{
			{
				"id": usersID[0],
			},
			{
				"id": usersID[1],
			},
		}

		var v []map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/teams/{TEAM-ID}/users?pretty=true",
					"rel": "self"
				}
			],
			"results": [
				{
					"country": "US",
					"emailAddress": "atlasUser@example.com",
					"firstName": "Atlas",
					"id": "1",
					"lastName": "user",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/1",
							"rel": "self"
						}
					],
					"mobileNumber": "5555550100",
					"roles": [
						{
							"orgId": "{ORG-ID}",
							"roleName": "ORG_MEMBER"
						}
					],
					"teamIds": [
						"{TEAM-ID}"
					],
					"username": "atlasUser@example.com"
				},
				{
					"country": "US",
					"emailAddress": "atlasUser@example.com",
					"firstName": "Atlas",
					"id": "2",
					"lastName": "user",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/users/1",
							"rel": "self"
						}
					],
					"mobileNumber": "5555550100",
					"roles": [
						{
							"orgId": "{ORG-ID}",
							"roleName": "ORG_MEMBER"
						}
					],
					"teamIds": [
						"{TEAM-ID}"
					],
					"username": "atlasUser@example.com"
				}
			],
			"totalCount": 2
		}`)
	})

	atlasUsers, _, err := client.Teams.AddUsersToTeam(ctx, orgID, teamID, usersID)
	if err != nil {
		t.Fatalf("Teams.AddUsersToTeam returned error: %v", err)
	}

	if userCount := len(atlasUsers); userCount != 2 {
		t.Errorf("expected userCount '%d', received '%d'", 2, userCount)
	}
}

func TestTeams_RemoveUserToTeam(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "1"

	userID := "1"
	teamID := "6b720e1087d9d66b272f1c86"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/teams/%s/users/%s", orgID, teamID, userID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Teams.RemoveUserToTeam(ctx, orgID, teamID, userID)
	if err != nil {
		t.Fatalf("Teams.RemoveUserToTeam returned error: %v", err)
	}
}

func TestTeams_RemoveTeamFromOrganization(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "1"
	teamID := "6b720e1087d9d66b272f1c86"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/teams/%s", orgID, teamID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Teams.RemoveTeamFromOrganization(ctx, orgID, teamID)
	if err != nil {
		t.Fatalf("Teams.RemoveTeamFromOrganization returned error: %v", err)
	}
}

func TestTeams_RemoveTeamFromProject(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	teamID := "6b720e1087d9d66b272f1c86"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/teams/%s", groupID, teamID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Teams.RemoveTeamFromProject(ctx, groupID, teamID)
	if err != nil {
		t.Fatalf("Teams.RemoveTeamFromProject returned error: %v", err)
	}
}
