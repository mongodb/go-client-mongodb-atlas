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

func TestEvents_ListOrganizationEvents(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "5b478b3afc4625789ce616a3"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/events", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
						"links": [
							{
								"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5b478b3afc4625789ce616a3/events?pageNum=1&itemsPerPage=100",
								"rel": "self"
							}
						],
						"results": [
							{
								"created": "2018-07-12T16:30:05Z",
								"eventTypeName": "JOINED_TEAM",
								"id": "b3ad04e680eef540be141abe",
								"isGlobalAdmin": true,
								"links": [
									{
										"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5b478b3afc4625789ce616a3/events/5b47820d80eef5264e12514a",
										"rel": "self"
									}
								],
								"orgId": "5b478b3afc4625789ce616a3",
								"remoteAddress": "203.0.113.22",
								"targetUsername": "b.doe@example.com",
								"userId": "5898b79080eef53b3ad04e68",
								"username": "j.doe@example.com"
							},
							{
								"created": "2018-07-09T21:14:40Z",
								"eventTypeName": "GROUP_CREATED",
								"groupId": "5b43d04087d9d6357de591a2",
								"id": "5b478b3afc49d6357de591af",
								"isGlobalAdmin": false,
								"links": [
									{
										"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5b478b3afc4625789ce616a3/events/5b43d04087d9d6357de591af",
										"rel": "self"
									}
								],
								"orgId": "5b478b3afc4625789ce616a3",
								"remoteAddress": "192.0.2.88",
								"userId": "5898b79080eef53b3ad04e68",
								"username": "j.doe@example.com"
							}
						],
						"totalCount": 2
					}`)
	})

	opts := &EventListOptions{}

	cluster, _, err := client.Events.ListOrganizationEvents(ctx, orgID, opts)
	if err != nil {
		t.Fatalf("Events.ListOrganizationEvents returned error: %v", err)
	}

	expected := &EventResponse{
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5b478b3afc4625789ce616a3/events?pageNum=1&itemsPerPage=100",
			},
		},
		Results: []*Event{
			{
				Created:       "2018-07-12T16:30:05Z",
				EventTypeName: "JOINED_TEAM",
				ID:            "b3ad04e680eef540be141abe",
				IsGlobalAdmin: true,
				Links: []*Link{
					{
						Rel:  "self",
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5b478b3afc4625789ce616a3/events/5b47820d80eef5264e12514a",
					},
				},
				OrgID:          "5b478b3afc4625789ce616a3",
				RemoteAddress:  "203.0.113.22",
				TargetUsername: "b.doe@example.com",
				UserID:         "5898b79080eef53b3ad04e68",
				Username:       "j.doe@example.com",
			},
			{
				Created:       "2018-07-09T21:14:40Z",
				EventTypeName: "GROUP_CREATED",
				GroupID:       "5b43d04087d9d6357de591a2",
				ID:            "5b478b3afc49d6357de591af",
				IsGlobalAdmin: false,
				Links: []*Link{
					{
						Rel:  "self",
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5b478b3afc4625789ce616a3/events/5b43d04087d9d6357de591af",
					},
				},
				OrgID:         "5b478b3afc4625789ce616a3",
				RemoteAddress: "192.0.2.88",
				UserID:        "5898b79080eef53b3ad04e68",
				Username:      "j.doe@example.com",
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(cluster, expected); diff != nil {
		t.Error(diff)
	}
}

func TestEvents_GetOrganizationEvent(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "5b478b3afc4625789ce616a3"
	eventID := "b3ad04e680eef540be141abe"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/events/%s", orgID, eventID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
								"created": "2018-07-12T16:30:05Z",
								"eventTypeName": "JOINED_TEAM",
								"id": "b3ad04e680eef540be141abe",
								"isGlobalAdmin": true,
								"links": [
									{
										"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5b478b3afc4625789ce616a3/events/5b47820d80eef5264e12514a",
										"rel": "self"
									}
								],
								"orgId": "5b478b3afc4625789ce616a3",
								"remoteAddress": "203.0.113.22",
								"targetUsername": "b.doe@example.com",
								"userId": "5898b79080eef53b3ad04e68",
								"username": "j.doe@example.com"
					}`)
	})

	cluster, _, err := client.Events.GetOrganizationEvent(ctx, orgID, eventID)
	if err != nil {
		t.Fatalf("Events.GetOrganizationEvent returned error: %v", err)
	}

	expected := &Event{
		Created:       "2018-07-12T16:30:05Z",
		EventTypeName: "JOINED_TEAM",
		ID:            "b3ad04e680eef540be141abe",
		IsGlobalAdmin: true,
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5b478b3afc4625789ce616a3/events/5b47820d80eef5264e12514a",
			},
		},
		OrgID:          "5b478b3afc4625789ce616a3",
		RemoteAddress:  "203.0.113.22",
		TargetUsername: "b.doe@example.com",
		UserID:         "5898b79080eef53b3ad04e68",
		Username:       "j.doe@example.com",
	}

	if diff := deep.Equal(cluster, expected); diff != nil {
		t.Error(diff)
	}
}

func TestEvents_ListProjectEvents(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "5b43d04087d9d6357de591a2"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/events", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
						"links": [
							{
								"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b43d04087d9d6357de591a2/events?pageNum=1&itemsPerPage=100",
								"rel": "self"
							}
						],
						"results": [
							{
								"created": "2018-07-12T16:30:05Z",
								"eventTypeName": "JOINED_TEAM",
								"id": "b3ad04e680eef540be141abe",
								"isGlobalAdmin": true,
								"links": [
									{
										"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b43d04087d9d6357de591a2/events/5b47820d80eef5264e12514a",
										"rel": "self"
									}
								],
								"orgId": "5b478b3afc4625789ce616a3",
								"remoteAddress": "203.0.113.22",
								"targetUsername": "b.doe@example.com",
								"userId": "5898b79080eef53b3ad04e68",
								"username": "j.doe@example.com"
							},
							{
								"created": "2018-07-09T21:14:40Z",
								"eventTypeName": "GROUP_CREATED",
								"groupId": "5b43d04087d9d6357de591a2",
								"id": "5b478b3afc49d6357de591af",
								"isGlobalAdmin": false,
								"links": [
									{
										"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b43d04087d9d6357de591a2/events/5b43d04087d9d6357de591af",
										"rel": "self"
									}
								],
								"orgId": "5b478b3afc4625789ce616a3",
								"remoteAddress": "192.0.2.88",
								"userId": "5898b79080eef53b3ad04e68",
								"username": "j.doe@example.com"
							}
						],
						"totalCount": 2
					}`)
	})

	cluster, _, err := client.Events.ListProjectEvents(ctx, groupID, nil)
	if err != nil {
		t.Fatalf("Events.ListOrganizationEvents returned error: %v", err)
	}

	expected := &EventResponse{
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b43d04087d9d6357de591a2/events?pageNum=1&itemsPerPage=100",
			},
		},
		Results: []*Event{
			{
				Created:       "2018-07-12T16:30:05Z",
				EventTypeName: "JOINED_TEAM",
				ID:            "b3ad04e680eef540be141abe",
				IsGlobalAdmin: true,
				Links: []*Link{
					{
						Rel:  "self",
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b43d04087d9d6357de591a2/events/5b47820d80eef5264e12514a",
					},
				},
				OrgID:          "5b478b3afc4625789ce616a3",
				RemoteAddress:  "203.0.113.22",
				TargetUsername: "b.doe@example.com",
				UserID:         "5898b79080eef53b3ad04e68",
				Username:       "j.doe@example.com",
			},
			{
				Created:       "2018-07-09T21:14:40Z",
				EventTypeName: "GROUP_CREATED",
				GroupID:       "5b43d04087d9d6357de591a2",
				ID:            "5b478b3afc49d6357de591af",
				IsGlobalAdmin: false,
				Links: []*Link{
					{
						Rel:  "self",
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b43d04087d9d6357de591a2/events/5b43d04087d9d6357de591af",
					},
				},
				OrgID:         "5b478b3afc4625789ce616a3",
				RemoteAddress: "192.0.2.88",
				UserID:        "5898b79080eef53b3ad04e68",
				Username:      "j.doe@example.com",
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(cluster, expected); diff != nil {
		t.Error(diff)
	}
}

func TestEvents_GetProjectEvent(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "5b478b3afc4625789ce616a3"
	eventID := "b3ad04e680eef540be141abe"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/events/%s", groupID, eventID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
								"created": "2018-07-12T16:30:05Z",
								"eventTypeName": "JOINED_TEAM",
								"id": "b3ad04e680eef540be141abe",
								"isGlobalAdmin": true,
								"links": [
									{
										"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b43d04087d9d6357de591a2/events/5b43d04087d9d6357de591af",
										"rel": "self"
									}
								],
								"orgId": "5b478b3afc4625789ce616a3",
								"remoteAddress": "203.0.113.22",
								"targetUsername": "b.doe@example.com",
								"userId": "5898b79080eef53b3ad04e68",
								"username": "j.doe@example.com"
					}`)
	})

	cluster, _, err := client.Events.GetProjectEvent(ctx, groupID, eventID)
	if err != nil {
		t.Fatalf("Events.GetOrganizationEvent returned error: %v", err)
	}

	expected := &Event{
		Created:       "2018-07-12T16:30:05Z",
		EventTypeName: "JOINED_TEAM",
		ID:            "b3ad04e680eef540be141abe",
		IsGlobalAdmin: true,
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5b43d04087d9d6357de591a2/events/5b43d04087d9d6357de591af",
			},
		},
		OrgID:          "5b478b3afc4625789ce616a3",
		RemoteAddress:  "203.0.113.22",
		TargetUsername: "b.doe@example.com",
		UserID:         "5898b79080eef53b3ad04e68",
		Username:       "j.doe@example.com",
	}

	if diff := deep.Equal(cluster, expected); diff != nil {
		t.Error(diff)
	}
}
