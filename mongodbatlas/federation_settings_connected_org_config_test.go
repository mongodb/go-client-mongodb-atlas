// Copyright 2022 MongoDB Inc
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

const connectedOrgID = "627a9683eafda674de306d14"
const connectedOrgFederationSettingsID = "627af3988201970138fd5dbe"

func TestFederatedSettingsConnectedOrganizationOp_List(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()

		mux.HandleFunc("/api/atlas/v1.0/federationSettings/627af3988201970138fd5dbe/connectedOrgConfigs", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			_, _ = fmt.Fprint(w, `{
				"links": [
				  {
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/federationSettings/5a0a1e7e0f2912c554080adc/connectedOrgConfigs/5a0a1e7e0f2912c554080ada/roleMappings?pageNum=1&itemsPerPage=100",
					"rel": "self"
				  }
				],
				"results": [
					{
				"domainAllowList":[
				   "reorganizeyourworld.com"
				],
				"domainRestrictionEnabled":false,
				"identityProviderId":"0oad0iizxh30vuyr5297",
				"orgId":"627a9683eafda674de306f14",
				"postAuthRoleGrants":[
				],
				"roleMappings":[
				   {
					  "externalGroupName":"myGroup",
					  "id":"627b1f8f244ad705fd542d81",
					  "roleAssignments":[
						 {
							"groupId":null,
							"orgId":"627a9683eafda674de306d14",
							"role":"ORG_OWNER"
						 }
					  ]
				   }
				],
				"userConflicts":null
			 }],
			 "totalCount": 1
		   }`)
		})

		connectedOrganization, _, err := client.FederatedSettings.ListConnectedOrgs(ctx, connectedOrgFederationSettingsID, nil)
		if err != nil {
			t.Fatalf("Organizations.List returned error: %v", err)
		}

		DomainRestrictionEnabled := false

		expected := &FederatedSettingsConnectedOrganizations{
			Links: []*Link{
				{
					Href: "https://cloud.mongodb.com/api/atlas/v1.0/federationSettings/5a0a1e7e0f2912c554080adc/connectedOrgConfigs/5a0a1e7e0f2912c554080ada/roleMappings?pageNum=1&itemsPerPage=100",
					Rel:  "self",
				},
			},
			Results: []*FederatedSettingsConnectedOrganization{

				{
					DomainAllowList:          []string{"reorganizeyourworld.com"},
					DomainRestrictionEnabled: &DomainRestrictionEnabled,
					IdentityProviderID:       "0oad0iizxh30vuyr5297",
					OrgID:                    "627a9683eafda674de306f14",
					PostAuthRoleGrants:       []string{},
					RoleMappings: []*RoleMappings{
						{
							ExternalGroupName: "myGroup",
							ID:                "627b1f8f244ad705fd542d81",
							RoleAssignments: []*RoleAssignments{
								{
									OrgID: "627a9683eafda674de306d14",
									Role:  "ORG_OWNER",
								},
							},
						},
					},
				},
			},
			TotalCount: 1,
		}

		if diff := deep.Equal(connectedOrganization, expected); diff != nil {
			t.Error(diff)
		}
	})
}

func TestFederatedSettingsConnectedOrganizationOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/federationSettings/%s/connectedOrgConfigs/%s", connectedOrgFederationSettingsID, connectedOrgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
   "domainAllowList":[
      "reorganizeyourworld.com"
   ],
   "domainRestrictionEnabled":false,
   "identityProviderId":"0oad0iizxh30vuyr5297",
   "orgId":"627a9683eafda674de306f14",
   "postAuthRoleGrants":[
   ],
   "roleMappings":[
      {
         "externalGroupName":"myGroup",
         "id":"627b1f8f244ad705fd542d81",
         "roleAssignments":[
            {
               "groupId":null,
               "orgId":"627a9683eafda674de306d14",
               "role":"ORG_OWNER"
            }
         ]
      }
   ],
   "userConflicts":null
}`)
	})

	response, _, err := client.FederatedSettings.GetConnectedOrg(ctx, connectedOrgFederationSettingsID, connectedOrgID)
	if err != nil {
		t.Fatalf("FederatedSettingsConnectedOrganization.Get returned error: %v", err)
	}

	expected := &FederatedSettingsConnectedOrganization{
		DomainAllowList:          []string{"reorganizeyourworld.com"},
		DomainRestrictionEnabled: pointy.Bool(false),
		IdentityProviderID:       "0oad0iizxh30vuyr5297",
		OrgID:                    "627a9683eafda674de306f14",
		PostAuthRoleGrants:       []string{},
		RoleMappings: []*RoleMappings{
			{
				ExternalGroupName: "myGroup",
				ID:                "627b1f8f244ad705fd542d81",
				RoleAssignments: []*RoleAssignments{
					{
						OrgID: "627a9683eafda674de306d14",
						Role:  "ORG_OWNER",
					},
				},
			},
		},
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}

func TestFederatedSettingsConnectedOrganizationOp_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/federationSettings/%s/connectedOrgConfigs/%s", connectedOrgFederationSettingsID, connectedOrgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		_, _ = fmt.Fprint(w, `{
   "domainAllowList":[
      "reorganizeyourworld.com"
   ],
   "domainRestrictionEnabled":false,
   "identityProviderId":"0oad0iizxh30vuyr5297",
   "orgId":"627a9683eafda674de306f14",
   "postAuthRoleGrants":[
   ],
   "roleMappings":[
      {
         "externalGroupName":"myGroup1",
         "id":"627b1f8f244ad705fd542d81",
         "roleAssignments":[
            {
               "groupId":null,
               "orgId":"627a9683eafda674de306d14",
               "role":"ORG_OWNER"
            }
         ]
      }
   ],
   "userConflicts":null
}`)
	})

	updated := &FederatedSettingsConnectedOrganization{
		DomainAllowList:          []string{"reorganizeyourworld.com"},
		DomainRestrictionEnabled: pointy.Bool(false),
		IdentityProviderID:       "0oad0iizxh30vuyr5297",
		OrgID:                    "627a9683eafda674de306d14",
		PostAuthRoleGrants:       []string{},
		RoleMappings: []*RoleMappings{
			{
				ExternalGroupName: "myGroup1",
				ID:                "627b1f8f244ad705fd542d81",
				RoleAssignments: []*RoleAssignments{
					{
						OrgID: "627a9683eafda674de306d14",
						Role:  "ORG_OWNER",
					},
				},
			},
		},
	}

	response, _, err := client.FederatedSettings.UpdateConnectedOrg(ctx, connectedOrgFederationSettingsID, connectedOrgID, updated)
	if err != nil {
		t.Fatalf("FederatedSettingsConnectedOrganization.Get returned error: %v", err)
	}

	expected := &FederatedSettingsConnectedOrganization{
		DomainAllowList:          []string{"reorganizeyourworld.com"},
		DomainRestrictionEnabled: pointy.Bool(false),
		IdentityProviderID:       "0oad0iizxh30vuyr5297",
		OrgID:                    "627a9683eafda674de306f14",
		PostAuthRoleGrants:       []string{},
		RoleMappings: []*RoleMappings{
			{
				ExternalGroupName: "myGroup1",
				ID:                "627b1f8f244ad705fd542d81",
				RoleAssignments: []*RoleAssignments{
					{
						OrgID: "627a9683eafda674de306d14",
						Role:  "ORG_OWNER",
					},
				},
			},
		},
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}

func TestFederatedSettingsConnectedOrganization_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/federationSettings/%s/connectedOrgConfigs/%s", connectedOrgFederationSettingsID, connectedOrgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.FederatedSettings.DeleteConnectedOrg(ctx, connectedOrgFederationSettingsID, connectedOrgID)
	if err != nil {
		t.Fatalf("FederatedSettingsConnectedOrganization.Delete returned error: %v", err)
	}
}
