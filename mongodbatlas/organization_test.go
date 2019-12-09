package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestOrganization_GetAllOrganizations(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links" : [
				 {
					 "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs",
					 "rel" : "self"
				 }
			],
			"results" : [
				 {
					 "id" : "5a2add1cfd9f3cfb17053317",
					 "links": [{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053317",
							"rel": "self"
				 }],
					 "name" : "Internal Application Support"
				 },
				 {
					 "id" : "5a2add1cfd9f3cfb17053318",
					 "links": [{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053318",
							"rel": "self"
				 }],
					 "name" : "Customer Application Support"
				 },
				 {
					 "id" : "5a2add1cfd9f3cfb17053319",
					 "links": [{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053319",
							"rel": "self"
				 }],
					 "name" : "Research and Development"
				 }
			],
			"totalCount" : 3
		}`)
	})

	organizations, _, err := client.Organizations.GetAllOrganizations(ctx)
	if err != nil {
		t.Errorf("Organization.GetAllOrganizations returned error: %v", err)
	}

	expected := &Organizations{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs",
				Rel:  "self",
			},
		},
		Results: []*Organization{
			{
				ID:   "5a2add1cfd9f3cfb17053317",
				Name: "Internal Application Support",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053317",
						Rel:  "self",
					},
				},
			},
			{
				ID:   "5a2add1cfd9f3cfb17053318",
				Name: "Customer Application Support",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053318",
						Rel:  "self",
					},
				},
			},
			{
				ID:   "5a2add1cfd9f3cfb17053319",
				Name: "Research and Development",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053319",
						Rel:  "self",
					},
				},
			},
		},
		TotalCount: 3,
	}

	if diff := deep.Equal(organizations, expected); diff != nil {
		t.Errorf("Organization.GetAllOrganizations\n got=%#v\nwant=%#v \n\ndiff=%#v", organizations, expected, diff)
	}
}

func TestOrganization_GetOneOrganization(t *testing.T) {
	setup()
	defer teardown()

	organizationID := "5a2add1cfd9f3cfb17053318"

	mux.HandleFunc(fmt.Sprintf("/%s/%s", organizationBasePath, organizationID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": "5a2add1cfd9f3cfb17053318",
			"links": [{
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053318",
				"rel": "self"
			}, {
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053318/groups",
				"rel": "http://mms.mongodb.com/groups"
			}, {
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053318/teams",
				"rel": "http://mms.mongodb.com/teams"
			}, {
				"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053318/users",
				"rel": "http://mms.mongodb.com/users"
			}],
			"name": "Customer Application Support"
		}`)
	})

	organization, _, err := client.Organizations.GetOneOrganization(ctx, organizationID)
	if err != nil {
		t.Errorf("Organization.GetOneOrganization returned error: %v", err)
	}

	expected := &Organization{
		ID:   "5a2add1cfd9f3cfb17053318",
		Name: "Customer Application Support",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053318",
				Rel:  "self",
			}, {
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053318/groups",
				Rel:  "http://mms.mongodb.com/groups",
			}, {
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053318/teams",
				Rel:  "http://mms.mongodb.com/teams",
			}, {
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053318/users",
				Rel:  "http://mms.mongodb.com/users",
			},
		},
	}

	if diff := deep.Equal(organization, expected); diff != nil {
		t.Errorf("Organization.GetOneOrganization\n got=%#v\nwant=%#v \n\ndiff=%#v", organization, expected, diff)
	}
}

func TestOrganization_GetAllOrganizationUsers(t *testing.T) {
	setup()
	defer teardown()

	organizationID := "59db8d1d87d9d6420df0613f"

	mux.HandleFunc(fmt.Sprintf("/%s/%s/users", organizationBasePath, organizationID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053317",
					"rel": "self"
				}
			],
			"results": [
				{
					"id": "59db8d1d87d9d6420df0613a",
					"teamIds": [
						"5aeeed020bd6ef9d00033291",
						"5ac2aeadcabceef96172be31"
					],
					"username": "someone@example.com",
					"country": "US",
					"emailAddress": "someone@example.com",
					"firstName": "John",
					"lastName": "Smith",
					"mobileNumber": "123-456-7890",
					"roles": [
						{
							"groupId": "59ea02e087d9d636b587a967",
							"roleName": "GROUP_OWNER"
						},
						{
							"groupId": "59db8d1d87d9d6420df70902",
							"roleName": "GROUP_OWNER"
						},
						{
							"orgId": "59db8d1d87d9d6420df0613f",
							"roleName": "ORG_OWNER"
						}
					],
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053317",
							"rel": "self"
						}
					]
				}
			],
			"totalCount": 1
		}`)
	})

	organizationUsers, _, err := client.Organizations.GetAllOrganizationUsers(ctx, organizationID)
	if err != nil {
		t.Errorf("Organization.GetAllOrganizationUsers returned error: %v", err)
	}

	expected := &OrganizationUsers{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053317",
				Rel:  "self",
			},
		},
		Results: []*OrganizationUser{
			{
				ID:           "59db8d1d87d9d6420df0613a",
				TeamIDS:      []string{"5aeeed020bd6ef9d00033291", "5ac2aeadcabceef96172be31"},
				Username:     "someone@example.com",
				Country:      "US",
				EmailAddress: "someone@example.com",
				FirstName:    "John",
				LastName:     "Smith",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a2add1cfd9f3cfb17053317",
						Rel:  "self",
					},
				},
				MobileNumber: "123-456-7890",
				Roles: []*OrganizationRole{
					{
						GroupID:  "59ea02e087d9d636b587a967",
						RoleName: "GROUP_OWNER",
					},
					{
						GroupID:  "59db8d1d87d9d6420df70902",
						RoleName: "GROUP_OWNER",
					},
					{
						OrgID:    "59db8d1d87d9d6420df0613f",
						RoleName: "ORG_OWNER",
					},
				},
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(organizationUsers, expected); diff != nil {
		t.Errorf("Organization.GetAllOrganizationUsers\n got=%#v\nwant=%#v \n\ndiff=%#v", organizationUsers, expected, diff)
	}
}

func TestOrganization_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs", func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"name": "myNewOrganization",
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("Decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Request body\n got=%#v\nwant=%#v \n\ndiff=%#v", v, expected, diff)
		}

		fmt.Fprintf(w, `{
			"id":"5a38289880eef5072af2f1f1",
			"name":"myNewOrganization"
			}`)
	})

	// Creating a new organization
	organization, _, err := client.Organizations.Create(ctx, "myNewOrganization")
	if err != nil {
		t.Errorf("Organization.Create returned error: %v", err)
	}

	expected := &Organization{
		ID:   "5a38289880eef5072af2f1f1",
		Name: "myNewOrganization",
	}

	if diff := deep.Equal(organization, expected); diff != nil {
		t.Errorf("Organization.Get\n got=%#v\nwant=%#v \n\ndiff=%#v", organization, expected, diff)
	}
}

func TestOrganization_UpdateOrganizationName(t *testing.T) {
	setup()
	defer teardown()

	orgID := "6001f2c580eef55aedbc6bb1"

	updateRequst := &Organization{
		ID:   orgID,
		Name: "MyAtlasCluster",
	}

	mux.HandleFunc(fmt.Sprintf("/orgs/%s", orgID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"id":   orgID,
			"name": "MyAtlasCluster",
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("Decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Request body\n got=%#v\nwant=%#v \n\ndiff=%#v", v, expected, diff)
		}

		fmt.Fprintf(w, `{
			"id": "6001f2c580eef55aedbc6bb1",
			"name": "MyAtlasCluster",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1",
					"rel": "self"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1/groups",
					"rel": "http://mms.mongodb.com/groups"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1/teams",
					"rel": "http://mms.mongodb.com/teams"
				},
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1/users",
					"rel": "http://mms.mongodb.com/users"
				}
			]
		}`)
	})

	organization, _, err := client.Organizations.UpdateOrganizationName(ctx, updateRequst)
	if err != nil {
		t.Errorf("Organization.UpdateOrganizationName returned error: %v", err)
	}

	expected := &Organization{
		ID:   "6001f2c580eef55aedbc6bb1",
		Name: "MyAtlasCluster",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1",
				Rel:  "self",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1/groups",
				Rel:  "http://mms.mongodb.com/groups",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1/teams",
				Rel:  "http://mms.mongodb.com/teams",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1/users",
				Rel:  "http://mms.mongodb.com/users",
			},
		},
	}

	if diff := deep.Equal(organization, expected); diff != nil {
		t.Errorf("Organization.Get\n got=%#v\nwant=%#v \n\ndiff=%#v", organization, expected, diff)
	}
}

func TestOrganization_Delete(t *testing.T) {
	setup()
	defer teardown()

	organizationID := "5a38289880eef5072af2f1f1"

	mux.HandleFunc(fmt.Sprintf("/orgs/%s", organizationID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Organizations.Delete(ctx, organizationID)
	if err != nil {
		t.Errorf("Organization.Delete returned error: %v", err)
	}
}

func TestOrganization_GetAllOrganizationProjects(t *testing.T) {
	setup()
	defer teardown()

	organizationID := "5a2add1cfd9f3cfb17053318"

	mux.HandleFunc(fmt.Sprintf("/%s/%s/groups", organizationBasePath, organizationID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a0a1e7e0f2912c554080adc/groups?pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				{
					"clusterCount": 1,
					"created": "2018-02-02T00:32:31Z",
					"id": "5ad645d4806d0119d6b35638",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5ad645d4806d0119d6b35638",
							"rel": "self"
						}
					],
					"name": "Production",
					"orgId": "5a0a1e7e0f2912c554080adc"
				},
				{
					"clusterCount": 1,
					"created": "2017-02-06T17:59:05Z",
					"id": "5ad645d4806d0119d6b35639",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/5ad645d4806d0119d6b35639",
							"rel": "self"
						}
					],
					"name": "Staging",
					"orgId": "5a0a1e7e0f2912c554080adc"
				}
			],
			"totalCount": 2
		}`)
	})

	projects, _, err := client.Organizations.GetAllOrganizationProjects(ctx, organizationID)
	if err != nil {
		t.Errorf("Organization.GetAllOrganizationProjects returned error: %v", err)
	}

	expected := &Projects{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a0a1e7e0f2912c554080adc/groups?pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*Project{
			{
				ClusterCount: 1,
				Created:      "2018-02-02T00:32:31Z",
				ID:           "5ad645d4806d0119d6b35638",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5ad645d4806d0119d6b35638",
						Rel:  "self",
					},
				},
				Name:  "Production",
				OrgID: "5a0a1e7e0f2912c554080adc",
			},
			{
				ClusterCount: 1,
				Created:      "2017-02-06T17:59:05Z",
				ID:           "5ad645d4806d0119d6b35639",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5ad645d4806d0119d6b35639",
						Rel:  "self",
					},
				},
				Name:  "Staging",
				OrgID: "5a0a1e7e0f2912c554080adc",
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(projects, expected); diff != nil {
		t.Errorf("Organization.GetAllOrganizationProjects\n got=%#v\nwant=%#v \n\ndiff=%#v", projects, expected, diff)
	}
}
