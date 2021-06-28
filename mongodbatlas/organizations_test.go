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
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

const (
	orgID =  "5a0a1e7e0f2912c554080adc"
	invitationID = "1"
)

func TestOrganizationsServiceOp_List(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()

		mux.HandleFunc("/api/atlas/v1.0/orgs", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			_, _ = fmt.Fprint(w, `{
				"links": [{
					"href": "https://cloud.mongodb.com/api/public/v1.0/orgs",
					"rel": "self"
				}],
				"results": [{
					"id": "56a10a80e4b0fd3b9a9bb0c2",
					"links": [{
						"href": "https://cloud.mongodb.com/api/public/v1.0/orgs/56a10a80e4b0fd3b9a9bb0c2",
						"rel": "self"
					}],
					"name": "012i3091203jioawjioej"
				}, {
					"id": "56aa691ce4b0a0e8c4be51f7",
					"links": [{
						"href": "https://cloud.mongodb.com/api/public/v1.0/orgs/56aa691ce4b0a0e8c4be51f7",
						"rel": "self"
					}],
					"name": "1454008603036"
				}],
				"totalCount": 2
			}`)
		})

		orgs, _, err := client.Organizations.List(ctx, nil)
		if err != nil {
			t.Fatalf("Organizations.List returned error: %v", err)
		}

		expected := &Organizations{
			Links: []*Link{
				{
					Href: "https://cloud.mongodb.com/api/public/v1.0/orgs",
					Rel:  "self",
				},
			},
			Results: []*Organization{
				{
					ID: "56a10a80e4b0fd3b9a9bb0c2",
					Links: []*Link{
						{
							Href: "https://cloud.mongodb.com/api/public/v1.0/orgs/56a10a80e4b0fd3b9a9bb0c2",
							Rel:  "self",
						},
					},
					Name: "012i3091203jioawjioej",
				},
				{
					ID: "56aa691ce4b0a0e8c4be51f7",
					Links: []*Link{
						{
							Href: "https://cloud.mongodb.com/api/public/v1.0/orgs/56aa691ce4b0a0e8c4be51f7",
							Rel:  "self",
						},
					},
					Name: "1454008603036",
				},
			},
			TotalCount: 2,
		}

		if diff := deep.Equal(orgs, expected); diff != nil {
			t.Error(diff)
		}
	})
	t.Run("by page number", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()

		mux.HandleFunc("/api/atlas/v1.0/orgs", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			_, _ = fmt.Fprint(w, `{
				"links": [
					{
						"href": "https://cloud.mongodb.com/api/public/v1.0/orgs?pageNum=1&itemsPerPage=1",
						"rel": "previous"
					},
					{
						"href": "https://cloud.mongodb.com/api/public/v1.0/orgs?pageNum=2&itemsPerPage=1",
						"rel": "self"
					},
					{
						"href": "https://cloud.mongodb.com/api/public/v1.0/orgs?itemsPerPage=3&pageNum=2",
						"rel": "next"
					}
				],
				"results": [{
					"id": "56a10a80e4b0fd3b9a9bb0c2",
					"links": [{
						"href": "https://cloud.mongodb.com/api/public/v1.0/orgs/56a10a80e4b0fd3b9a9bb0c2",
						"rel": "self"
					}],
					"name": "FooBar"
				}],
				"totalCount": 3
			}`)
		})

		opt := &OrganizationsListOptions{ListOptions: ListOptions{PageNum: 2}, Name: "FooBar"}

		orgs, _, err := client.Organizations.List(ctx, opt)
		if err != nil {
			t.Fatalf("Organizations.List returned error: %v", err)
		}

		expected := &Organizations{
			Links: []*Link{
				{
					Href: "https://cloud.mongodb.com/api/public/v1.0/orgs?pageNum=1&itemsPerPage=1",
					Rel:  "previous",
				},
				{
					Href: "https://cloud.mongodb.com/api/public/v1.0/orgs?pageNum=2&itemsPerPage=1",
					Rel:  "self",
				},
				{
					Href: "https://cloud.mongodb.com/api/public/v1.0/orgs?itemsPerPage=3&pageNum=2",
					Rel:  "next",
				},
			},
			Results: []*Organization{
				{
					ID: "56a10a80e4b0fd3b9a9bb0c2",
					Links: []*Link{
						{
							Href: "https://cloud.mongodb.com/api/public/v1.0/orgs/56a10a80e4b0fd3b9a9bb0c2",
							Rel:  "self",
						},
					},
					Name: "FooBar",
				},
			},
			TotalCount: 3,
		}

		if diff := deep.Equal(orgs, expected); diff != nil {
			t.Error(diff)
		}
	})
}

func TestOrganizationsServiceOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s/%s", orgsBasePath, orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
		"id": "5a0a1e7e0f2912c554080adc",
		"lastActiveAgent": "2016-03-09T18:19:37Z",
		"links": [{
			"href": "https://cloud.mongodb.com/api/public/v1.0/orgs/5a0a1e7e0f2912c554080adc",
			"rel": "self"
		}],
		"name": "012i3091203jioawjioej"
	  }`)
	})

	response, _, err := client.Organizations.Get(ctx, orgID)
	if err != nil {
		t.Fatalf("Organizations.Get returned error: %v", err)
	}

	expected := &Organization{
		ID: "5a0a1e7e0f2912c554080adc",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/public/v1.0/orgs/5a0a1e7e0f2912c554080adc",
				Rel:  "self",
			},
		},
		Name: "012i3091203jioawjioej",
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizationsServiceOp_Projects(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s/%s/groups", orgsBasePath, orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
			"links": [{
				"href": "https://cloud.mongodb.com/api/public/v1.0/orgs/5a0a1e7e0f2912c554080adc/groups",
				"rel": "self"
			}],
			"results": [{
				"id": "56a10a80e4b0fd3b9a9bb0c2",
				"links": [{
					"href": "https://cloud.mongodb.com/api/public/v1.0/groups/56a10a80e4b0fd3b9a9bb0c2",
					"rel": "self"
				}],
				"name": "012i3091203jioawjioej",
				"orgId": "5a0a1e7e0f2912c554080adc"
			}],
			"totalCount": 1
		}`)
	})

	projects, _, err := client.Organizations.Projects(ctx, orgID, nil)
	if err != nil {
		t.Fatalf("Organizations.GetProjects returned error: %v", err)
	}

	expected := &Projects{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/public/v1.0/orgs/5a0a1e7e0f2912c554080adc/groups",
				Rel:  "self",
			},
		},
		Results: []*Project{
			{
				ID: "56a10a80e4b0fd3b9a9bb0c2",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/public/v1.0/groups/56a10a80e4b0fd3b9a9bb0c2",
						Rel:  "self",
					},
				},
				Name:  "012i3091203jioawjioej",
				OrgID: "5a0a1e7e0f2912c554080adc",
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(projects, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizationsServiceOp_Users(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%s/%s/users", orgsBasePath, orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
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

	users, _, err := client.Organizations.Users(ctx, orgID, nil)
	if err != nil {
		t.Fatalf("Organizations.Users returned error: %v", err)
	}

	expected := &AtlasUsersResponse{
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/1/users?pretty=true&pageNum=1&itemsPerPage=100",
			},
		},
		Results: []AtlasUser{
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
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(users, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizations_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Organizations.Delete(ctx, orgID)
	if err != nil {
		t.Fatalf("Organizations.Delete returned error: %v", err)
	}
}

func TestOrganizations_ListUnacceptedInvitations(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/invites", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `[
               {
               "createdAt": "2021-02-18T21:05:40Z",
			   "expiresAt": "2021-03-20T21:05:40Z",
			   "id": "5a0a1e7e0f2912c554080adc",
			   "inviterUsername": "admin@example.com",
			   "orgId": "5df7a168f10fab3a149357fb",
			   "orgName": "jww-12-16",
			   "roles": [
				   "ORG_OWNER"
			   ],
			   "username": "wyatt.smith@example.com"},
			   {"createdAt": "2021-02-18T21:05:40Z",
			   "expiresAt": "2021-03-20T21:05:40Z",
			   "id": "5a0a1e7e0f2912c554080adc",
			   "inviterUsername": "admin@example.com",
			   "orgId": "5df7a168f10fab3a149357fb",
			   "orgName": "jww-12-16",
			   "roles": [
				   "ORG_OWNER"
			   ],
			   "teamIds": ["2"],
			   "username": "wyatt.smith@example.com"}]`)
	})

	invitation, _, err := client.Organizations.ListUnacceptedInvitations(ctx, orgID, nil)
	if err != nil {
		t.Fatalf("Organizations.GetUnacceptedInvitation returned error: %v", err)
	}

	expected := &[]Invitation{
		{
			ID:              "5a0a1e7e0f2912c554080adc",
			OrgID:           "5df7a168f10fab3a149357fb",
			OrgName:         "jww-12-16",
			CreatedAt:       "2021-02-18T21:05:40Z",
			ExpiresAt:       "2021-03-20T21:05:40Z",
			InviterUserName: "admin@example.com",
			Username:        "wyatt.smith@example.com",
			Roles:           []string{"ORG_OWNER"},
		},
		{
			ID:              "5a0a1e7e0f2912c554080adc",
			OrgID:           "5df7a168f10fab3a149357fb",
			OrgName:         "jww-12-16",
			CreatedAt:       "2021-02-18T21:05:40Z",
			ExpiresAt:       "2021-03-20T21:05:40Z",
			InviterUserName: "admin@example.com",
			Username:        "wyatt.smith@example.com",
			Roles:           []string{"ORG_OWNER"},
			TeamIDs:[]string{"2"},
		},

	}

	if diff := deep.Equal(invitation, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizations_GetUnacceptedInvitation(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/invites/%s", orgID, invitationID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
			   "createdAt": "2021-02-18T21:05:40Z",
			   "expiresAt": "2021-03-20T21:05:40Z",
			   "id": "5a0a1e7e0f2912c554080adc",
			   "inviterUsername": "admin@example.com",
			   "orgId": "5df7a168f10fab3a149357fb",
			   "orgName": "jww-12-16",
			   "roles": [
				   "ORG_OWNER"
			   ],
			   "username": "wyatt.smith@example.com"
		}`)
	})

	invitation, _, err := client.Organizations.GetUnacceptedInvitation(ctx, orgID, invitationID)
	if err != nil {
		t.Fatalf("Organizations.GetUnacceptedInvitation returned error: %v", err)
	}

	expected := &Invitation{
		ID:              "5a0a1e7e0f2912c554080adc",
		OrgID:           "5df7a168f10fab3a149357fb",
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUserName: "admin@example.com",
		Username:        "wyatt.smith@example.com",
		Roles:           []string{"ORG_OWNER"},
	}

	if diff := deep.Equal(invitation, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizations_InviteUser(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/invites", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, _ = fmt.Fprint(w, `{
			   "createdAt": "2021-02-18T21:05:40Z",
			   "expiresAt": "2021-03-20T21:05:40Z",
			   "id": "5a0a1e7e0f2912c554080adc",
			   "inviterUsername": "admin@example.com",
			   "orgId": "5df7a168f10fab3a149357fb",
			   "orgName": "jww-12-16",
			   "roles": [
				   "ORG_OWNER"
			   ],
			   "username": "wyatt.smith@example.com"
		}`)
	})

	body := &Invitation{
		OrgID:           orgID,
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUserName: "admin@example.com",
		Username:        "wyatt.smith@example.com",
		Roles:           []string{"ORG_OWNER"},
	}

	invitation, _, err := client.Organizations.InviteUser(ctx, body)
	if err != nil {
		t.Fatalf("Organizations.InviteUser returned error: %v", err)
	}

	expected := &Invitation{
		ID:              "5a0a1e7e0f2912c554080adc",
		OrgID:           "5df7a168f10fab3a149357fb",
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUserName: "admin@example.com",
		Username:        "wyatt.smith@example.com",
		Roles:           []string{"ORG_OWNER"},
	}

	if diff := deep.Equal(invitation, expected); diff != nil {
		t.Error(diff)
	}
}


func TestOrganizations_UpdateInvitation(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/invites", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		_, _ = fmt.Fprint(w, `{
			   "createdAt": "2021-02-18T21:05:40Z",
			   "expiresAt": "2021-03-20T21:05:40Z",
			   "id": "5a0a1e7e0f2912c554080adc",
			   "inviterUsername": "admin@example.com",
			   "orgId": "5df7a168f10fab3a149357fb",
			   "orgName": "jww-12-16",
			   "roles": [
				   "ORG_OWNER"
			   ],
			   "username": "wyatt.smith@example.com"
		}`)
	})

	body := &Invitation{
		OrgID:           orgID,
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUserName: "admin@example.com",
		Username:        "wyatt.smith@example.com",
		Roles:           []string{"ORG_OWNER"},
	}

	invitation, _, err := client.Organizations.UpdateInvitation(ctx, body)
	if err != nil {
		t.Fatalf("Organizations.UpdateInvitation returned error: %v", err)
	}

	expected := &Invitation{
		ID:              "5a0a1e7e0f2912c554080adc",
		OrgID:           "5df7a168f10fab3a149357fb",
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUserName: "admin@example.com",
		Username:        "wyatt.smith@example.com",
		Roles:           []string{"ORG_OWNER"},
	}

	if diff := deep.Equal(invitation, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizations_UpdateInvitationByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/invites/%s", orgID, invitationID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		_, _ = fmt.Fprint(w, `{
			   "createdAt": "2021-02-18T21:05:40Z",
			   "expiresAt": "2021-03-20T21:05:40Z",
			   "id": "5a0a1e7e0f2912c554080adc",
			   "inviterUsername": "admin@example.com",
			   "orgId": "5df7a168f10fab3a149357fb",
			   "orgName": "jww-12-16",
			   "roles": [
				   "ORG_OWNER"
			   ],
			   "username": "wyatt.smith@example.com"
		}`)
	})

	body := &Invitation{
		OrgID:           orgID,
		ID: invitationID,
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUserName: "admin@example.com",
		Username:        "wyatt.smith@example.com",
		Roles:           []string{"ORG_OWNER"},
	}

	invitation, _, err := client.Organizations.UpdateInvitation(ctx, body)
	if err != nil {
		t.Fatalf("Organizations.UpdateInvitation returned error: %v", err)
	}

	expected := &Invitation{
		ID:              "5a0a1e7e0f2912c554080adc",
		OrgID:           "5df7a168f10fab3a149357fb",
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUserName: "admin@example.com",
		Username:        "wyatt.smith@example.com",
		Roles:           []string{"ORG_OWNER"},
	}

	if diff := deep.Equal(invitation, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizations_DeleteInvitation(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/invites/%s", orgID, invitationID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Organizations.DeleteInvitation(ctx, orgID, invitationID)
	if err != nil {
		t.Fatalf("Organizations.DeleteInvitation returned error: %v", err)
	}
}
