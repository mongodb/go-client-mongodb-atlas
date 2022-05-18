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

const orgConnectionID = "5a0a1e7e0f2912c554080adb"
const orgConnectionFederationSettingsID = "5a0a1e7e0f2912c554080adc"
const roleMappingID = "5a0a1e7e0f2912c554080add"

func TestFederatedSettingsOrganizationConnectionServiceOp_List(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()

		mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/federationSettings/%s/connectedOrgConfigs/%s/roleMappings", orgConnectionFederationSettingsID, orgConnectionID), func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			_, _ = fmt.Fprint(w, `{
				"links": [
				  {
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/federationSettings/5a0a1e7e0f2912c554080adc/connectedOrgConfigs/5a0a1e7e0f2912c554080adb/roleMappings",
					"rel": "self"
				  }
				],
				"results": [
				  {
				   "externalGroupName": "example",
				   "id": "61e89721b827b56c845ff44c",
				   "roleAssignments": [
					 {
					   "groupId": null,
					   "orgId": "5a0a1e7e0f2912c554080adc",
					   "role": "ORG_MEMBER"
					 }
				   ]
				  }
				],
				"totalCount": 1
			  }`)
		})

		orgs, _, err := client.FederatedSettingsOrganizationConnection.List(ctx, nil, orgConnectionFederationSettingsID, orgConnectionID)
		if err != nil {
			t.Fatalf("Organizations.List returned error: %v", err)
		}

		expected := &FederatedSettingsOrganizationConnections{
			Links: []*Link{
				{
					Href: "https://cloud.mongodb.com/api/atlas/v1.0/federationSettings/5a0a1e7e0f2912c554080adc/connectedOrgConfigs/5a0a1e7e0f2912c554080adb/roleMappings",
					Rel:  "self",
				},
			},
			Results: []*FederatedSettingsOrganizationConnection{
				{
					ExternalGroupName: "example",
					ID:                "61e89721b827b56c845ff44c",
					RoleAssignments: RoleAssignments{
						{
							GroupID: "",
							OrgID:   "5a0a1e7e0f2912c554080adc",
							Role:    "ORG_MEMBER",
						},
					},
				},
			},

			TotalCount: 1,
		}

		if diff := deep.Equal(orgs, expected); diff != nil {
			t.Error(diff)
		}
	})

}

func TestFederatedSettingsOrganizationConnectionServiceOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/federationSettings/%s/connectedOrgConfigs/%s/roleMappings/%s", orgConnectionFederationSettingsID, orgConnectionID, roleMappingID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
			"externalGroupName": "autocomplete-highlight",
			"id": "61d88e15e6cc044270a36fce",
			"roleAssignments": [
				{
					"groupId": null,
					"orgId": "5f86fb11e0079069c9ec3132",
					"role": "ORG_OWNER"
				},
				{
					"groupId": "5f86fb2ff9c4e56d39502559",
					"orgId": null,
					"role": "GROUP_OWNER"
				}
			]
		   }`)
	})

	response, _, err := client.FederatedSettingsOrganizationConnection.Get(ctx, orgConnectionFederationSettingsID, orgConnectionID, roleMappingID)
	if err != nil {
		t.Fatalf("Organizations.Get returned error: %v", err)
	}

	expected := &FederatedSettingsOrganizationConnection{
		ExternalGroupName: "autocomplete-highlight",
		ID:                "61d88e15e6cc044270a36fce",
		RoleAssignments: RoleAssignments{
			{
				GroupID: "",
				OrgID:   "5f86fb11e0079069c9ec3132",
				Role:    "ORG_OWNER",
			},
			{
				GroupID: "5f86fb2ff9c4e56d39502559",
				OrgID:   "",
				Role:    "GROUP_OWNER",
			},
		},
	}
	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}

func TestFederatedSettingsOrganizationConnectionServiceOp_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/federationSettings/%s/connectedOrgConfigs/%s/roleMappings/%s", orgConnectionFederationSettingsID, orgConnectionID, roleMappingID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.FederatedSettingsOrganizationConnection.Delete(ctx, orgConnectionFederationSettingsID, orgConnectionID, roleMappingID)
	if err != nil {
		t.Fatalf("Organizations.Delete returned error: %v", err)
	}
}
