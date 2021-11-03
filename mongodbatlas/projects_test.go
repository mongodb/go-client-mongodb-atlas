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
	"github.com/openlyinc/pointy"
)

func TestProject_GetAllProjects(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links" : [ {
				"href" : "https://cloud.mongodb.com/api/atlas/v1.0/groups",
				"rel" : "self"
			} ],
			"results" : [ {
				"clusterCount" : 2,
				"created" : "2016-07-14T14:19:33Z",
				"id" : "5a0a1e7e0f2912c554080ae6",
				"links" : [ {
					"href" : "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
					"rel" : "self"
				} ],
				"name" : "ProjectBar",
				"orgId" : "5a0a1e7e0f2912c554080adc"
			}, {
				"clusterCount" : 0,
				"created" : "2017-10-16T15:24:01Z",
				"id" : "5a0a1e7e0f2912c554080ae7",
				"links" : [ {
					"href" : "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae7",
					"rel" : "self"
				} ],
				"name" : "Project Foo",
				"orgId" : "5a0a1e7e0f2912c554080adc"
			} ],
			"totalCount" : 2
		}`)
	})

	projects, _, err := client.Projects.GetAllProjects(ctx, nil)
	if err != nil {
		t.Fatalf("Projects.GetAllProjects returned error: %v", err)
	}

	expected := &Projects{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups",
				Rel:  "self",
			},
		},
		Results: []*Project{
			{
				ClusterCount: 2,
				Created:      "2016-07-14T14:19:33Z",
				ID:           "5a0a1e7e0f2912c554080ae6",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
						Rel:  "self",
					},
				},
				Name:  "ProjectBar",
				OrgID: "5a0a1e7e0f2912c554080adc",
			},
			{
				ClusterCount: 0,
				Created:      "2017-10-16T15:24:01Z",
				ID:           "5a0a1e7e0f2912c554080ae7",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae7",
						Rel:  "self",
					},
				},
				Name:  "Project Foo",
				OrgID: "5a0a1e7e0f2912c554080adc",
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(projects, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProject_GetOneProject(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "5a0a1e7e0f2912c554080adc"

	mux.HandleFunc(fmt.Sprintf("/%s/%s", projectBasePath, projectID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id" : "5a0a1e7e0f2912c554080ae6",
			"orgId" : "5a0a1e7e0f2912c554080adc",
			"name" : "DocsFeedbackGroup",
			"clusterCount" : 2,
			"created" : "2016-07-14T14:19:33Z",
			"links" : [ {
				"href" : "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
				"rel" : "self"
			} ]
		}`)
	})

	projectResponse, _, err := client.Projects.GetOneProject(ctx, projectID)
	if err != nil {
		t.Errorf("Projects.GetOneProject returned error: %v", err)
	}

	expected := &Project{
		ID:           "5a0a1e7e0f2912c554080ae6",
		OrgID:        "5a0a1e7e0f2912c554080adc",
		Name:         "DocsFeedbackGroup",
		ClusterCount: 2,
		Created:      "2016-07-14T14:19:33Z",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
				Rel:  "self",
			},
		},
	}

	if diff := deep.Equal(projectResponse, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProject_GetOneProjectByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectName := "5a0a1e7e0f2912c554080adc"

	mux.HandleFunc(fmt.Sprintf("/%s/byName/%s", projectBasePath, projectName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id" : "5a0a1e7e0f2912c554080ae6",
			"orgId" : "5a0a1e7e0f2912c554080adc",
			"name" : "ProjectBar",
			"clusterCount" : 2,
			"created" : "2016-07-14T14:19:33Z",
			"links" : [ {
				"href" : "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
				"rel" : "self"
			} ]
		}`)
	})

	projectResponse, _, err := client.Projects.GetOneProjectByName(ctx, projectName)
	if err != nil {
		t.Fatalf("Projects.GetOneProject returned error: %v", err)
	}

	expected := &Project{
		ID:           "5a0a1e7e0f2912c554080ae6",
		OrgID:        "5a0a1e7e0f2912c554080adc",
		Name:         "ProjectBar",
		ClusterCount: 2,
		Created:      "2016-07-14T14:19:33Z",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
				Rel:  "self",
			},
		},
	}

	if diff := deep.Equal(projectResponse, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProject_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &Project{
		OrgID: "5a0a1e7e0f2912c554080adc",
		Name:  "ProjectFoobar",
	}

	mux.HandleFunc("/api/atlas/v1.0/groups", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"clusterCount": 2,
			"created": "2016-07-14T14:19:33Z",
			"id": "5a0a1e7e0f2912c554080ae6",
			"links": [{
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
				"rel": "self"
			}],
			"name": "ProjectFoobar",
			"orgId": "5a0a1e7e0f2912c554080adc",
			"withDefaultAlertsSettings": true
		}`)
	})

	opts := &CreateProjectOptions{ProjectOwnerID: "1"}

	project, _, err := client.Projects.Create(ctx, createRequest, opts)
	if err != nil {
		t.Fatalf("Projects.Create returned error: %v", err)
	}

	expected := &Project{
		ClusterCount: 2,
		Created:      "2016-07-14T14:19:33Z",
		ID:           "5a0a1e7e0f2912c554080ae6",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
				Rel:  "self",
			},
		},
		Name:                      "ProjectFoobar",
		OrgID:                     "5a0a1e7e0f2912c554080adc",
		WithDefaultAlertsSettings: pointy.Bool(true),
	}

	if diff := deep.Equal(project, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProject_Create_without_opts(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &Project{
		OrgID: "5a0a1e7e0f2912c554080adc",
		Name:  "ProjectFoobar",
	}

	mux.HandleFunc("/api/atlas/v1.0/groups", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"clusterCount": 2,
			"created": "2016-07-14T14:19:33Z",
			"id": "5a0a1e7e0f2912c554080ae6",
			"links": [{
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
				"rel": "self"
			}],
			"name": "ProjectFoobar",
			"orgId": "5a0a1e7e0f2912c554080adc"
		}`)
	})

	project, _, err := client.Projects.Create(ctx, createRequest, nil)
	if err != nil {
		t.Fatalf("Projects.Create returned error: %v", err)
	}

	expected := &Project{
		ClusterCount: 2,
		Created:      "2016-07-14T14:19:33Z",
		ID:           "5a0a1e7e0f2912c554080ae6",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
				Rel:  "self",
			},
		},
		Name:  "ProjectFoobar",
		OrgID: "5a0a1e7e0f2912c554080adc",
	}

	if diff := deep.Equal(project, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProject_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "5a0a1e7e0f2912c554080adc"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s", projectID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Projects.Delete(ctx, projectID)
	if err != nil {
		t.Fatalf("Projects.Delete returned error: %v", err)
	}
}

func TestProject_GetProjectTeamsAssigned(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "5a0a1e7e0f2912c554080adc"

	mux.HandleFunc(fmt.Sprintf("/%s/%s/teams", projectBasePath, projectID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/teams",
					"rel": "self"
				}
			],
			"results": [
				{
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/teams/{TEAM-ID}",
							"rel": "self"
						}
					],
					"roleNames": [
						"GROUP_READ_ONLY"
					],
					"teamId": "{TEAM-ID}"
				}
			],
			"totalCount": 1
		}`)
	})

	teamsAssigned, _, err := client.Projects.GetProjectTeamsAssigned(ctx, projectID)
	if err != nil {
		t.Fatalf("Projects.GetProjectTeamsAssigned returned error: %v", err)
	}

	expected := &TeamsAssigned{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/teams",
				Rel:  "self",
			},
		},
		Results: []*Result{
			{
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/teams/{TEAM-ID}",
						Rel:  "self",
					},
				},
				RoleNames: []string{"GROUP_READ_ONLY"},
				TeamID:    "{TEAM-ID}",
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(teamsAssigned, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProject_AddTeamsToProject(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "5a0a1e7e0f2912c554080adc"

	createRequest := []*ProjectTeam{{
		TeamID:    "{TEAM-ID}",
		RoleNames: []string{"GROUP_OWNER", "GROUP_READ_ONLY"},
	}}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/teams", projectID), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"links": [{
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/teams",
				"rel": "self"
			}],
			"results": [{
				"links": [{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/teams/{TEAM-ID}",
					"rel": "self"
				}],
				"roleNames": ["GROUP_OWNER"],
				"teamId": "{TEAM-ID}"
			}],
			"totalCount": 1
		}`)
	})

	team, _, err := client.Projects.AddTeamsToProject(ctx, projectID, createRequest)
	if err != nil {
		t.Errorf("Projects.AddTeamsToProject returned error: %v", err)
	}

	expected := &TeamsAssigned{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/teams",
				Rel:  "self",
			},
		},
		Results: []*Result{
			{
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/teams/{TEAM-ID}",
						Rel:  "self",
					},
				},
				RoleNames: []string{"GROUP_OWNER"},
				TeamID:    "{TEAM-ID}",
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(team, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProject_RemoveUserFromProject(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "5a0a1e7e0f2912c554080adc"
	userID := "1213232233243434"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/users/%s", projectID, userID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Projects.RemoveUserFromProject(ctx, projectID, userID)
	if err != nil {
		t.Fatalf("Projects.RemoveUserFromProject returned error: %v", err)
	}
}
