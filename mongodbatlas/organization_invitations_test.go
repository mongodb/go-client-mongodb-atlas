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

const invitationID = "1"

func TestOrganizations_Invitations(t *testing.T) {
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

	invitation, _, err := client.Organizations.Invitations(ctx, orgID, nil)
	if err != nil {
		t.Fatalf("Organizations.Invitations returned error: %v", err)
	}

	expected := []*Invitation{
		{
			ID:              "5a0a1e7e0f2912c554080adc",
			OrgID:           "5df7a168f10fab3a149357fb",
			OrgName:         "jww-12-16",
			CreatedAt:       "2021-02-18T21:05:40Z",
			ExpiresAt:       "2021-03-20T21:05:40Z",
			InviterUsername: "admin@example.com",
			Username:        "wyatt.smith@example.com",
			Roles:           []string{"ORG_OWNER"},
		},
		{
			ID:              "5a0a1e7e0f2912c554080adc",
			OrgID:           "5df7a168f10fab3a149357fb",
			OrgName:         "jww-12-16",
			CreatedAt:       "2021-02-18T21:05:40Z",
			ExpiresAt:       "2021-03-20T21:05:40Z",
			InviterUsername: "admin@example.com",
			Username:        "wyatt.smith@example.com",
			Roles:           []string{"ORG_OWNER"},
			TeamIDs:         []string{"2"},
		},
	}

	if diff := deep.Equal(invitation, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizations_Invitation(t *testing.T) {
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

	invitation, _, err := client.Organizations.Invitation(ctx, orgID, invitationID)
	if err != nil {
		t.Fatalf("Organizations.Invitation returned error: %v", err)
	}

	expected := &Invitation{
		ID:              "5a0a1e7e0f2912c554080adc",
		OrgID:           "5df7a168f10fab3a149357fb",
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUsername: "admin@example.com",
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
		InviterUsername: "admin@example.com",
		Username:        "wyatt.smith@example.com",
		Roles:           []string{"ORG_OWNER"},
	}

	invitation, _, err := client.Organizations.InviteUser(ctx, orgID, body)
	if err != nil {
		t.Fatalf("Organizations.InviteUser returned error: %v", err)
	}

	expected := &Invitation{
		ID:              "5a0a1e7e0f2912c554080adc",
		OrgID:           "5df7a168f10fab3a149357fb",
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUsername: "admin@example.com",
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
		InviterUsername: "admin@example.com",
		Username:        "wyatt.smith@example.com",
		Roles:           []string{"ORG_OWNER"},
	}

	invitation, _, err := client.Organizations.UpdateInvitation(ctx, orgID, body)
	if err != nil {
		t.Fatalf("Organizations.UpdateInvitation returned error: %v", err)
	}

	expected := &Invitation{
		ID:              "5a0a1e7e0f2912c554080adc",
		OrgID:           "5df7a168f10fab3a149357fb",
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUsername: "admin@example.com",
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
		ID:              invitationID,
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUsername: "admin@example.com",
		Username:        "wyatt.smith@example.com",
		Roles:           []string{"ORG_OWNER"},
	}

	invitation, _, err := client.Organizations.UpdateInvitationByID(ctx, orgID, invitationID, body)
	if err != nil {
		t.Fatalf("Organizations.UpdateInvitationByID returned error: %v", err)
	}

	expected := &Invitation{
		ID:              "5a0a1e7e0f2912c554080adc",
		OrgID:           "5df7a168f10fab3a149357fb",
		OrgName:         "jww-12-16",
		CreatedAt:       "2021-02-18T21:05:40Z",
		ExpiresAt:       "2021-03-20T21:05:40Z",
		InviterUsername: "admin@example.com",
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
