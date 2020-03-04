package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestOrganizations_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links" : [ {
			  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs?pretty=true&pageNum=1&itemsPerPage=100",
			  "rel" : "self"
			} ],
			"results" : [ {
			  "id" : "9b1dcf89d42c0585cbceb464",
			  "links" : [ {
				"href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/9b1dcf89d42c0585cbceb464",
				"rel" : "self"
			  } ],
			  "name" : "FirstOrg"
			},
			{
			  "id" : "bcb8206e35f8619bfd6329a5",
			  "links" : [ {
				"href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5",
				"rel" : "self"
			  } ],
			  "name" : "SecondOrg"
			} ],
			"totalCount" : 2
		  }`)
	})

	orgs, _, err := client.Organizations.List(ctx, nil)

	if err != nil {
		t.Fatalf("Organizations.List returned error: %v", err)
	}

	expected := []Organization{
		{
			ID:   "9b1dcf89d42c0585cbceb464",
			Name: "FirstOrg",
			Links: []*Link{
				{
					Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/9b1dcf89d42c0585cbceb464",
					Rel:  "self",
				},
			},
		},
		{
			ID:   "bcb8206e35f8619bfd6329a5",
			Name: "SecondOrg",
			Links: []*Link{
				{
					Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5",
					Rel:  "self",
				},
			},
		},
	}
	if diff := deep.Equal(orgs, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizations_Get(t *testing.T) {
	setup()
	defer teardown()

	orgID := "bcb8206e35f8619bfd6329a5"

	mux.HandleFunc(fmt.Sprintf("/orgs/%s", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id" : "bcb8206e35f8619bfd6329a5",
			"links" : [ {
			  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5",
			  "rel" : "self"
			}, {
			  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5/groups",
			  "rel" : "http://mms.mongodb.com/groups"
			}, {
			  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5/teams",
			  "rel" : "http://mms.mongodb.com/teams"
			}, {
			  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5/users",
			  "rel" : "http://mms.mongodb.com/users"
			} ],
			"name" : "TestOrg"
		  }`)
	})

	org, _, err := client.Organizations.Get(ctx, orgID)
	if err != nil {
		t.Fatalf("Organizations.Get returned error: %v", err)
	}

	expected := &Organization{
		ID:   "bcb8206e35f8619bfd6329a5",
		Name: "TestOrg",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5",
				Rel:  "self",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5/groups",
				Rel:  "http://mms.mongodb.com/groups",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5/teams",
				Rel:  "http://mms.mongodb.com/teams",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5/users",
				Rel:  "http://mms.mongodb.com/users",
			},
		},
	}

	if diff := deep.Equal(org, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizations_ListUsers(t *testing.T) {
	setup()
	defer teardown()

	orgID := "bcb8206e35f8619bfd6329a5"

	mux.HandleFunc(fmt.Sprintf("/orgs/%s/users", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links" : [ {
			  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/bcb8206e35f8619bfd6329a5/users",
			  "rel" : "self"
			} ],
			"results" : [ {
			  "country" : "DE",
			  "emailAddress" : "joe@example.com",
			  "firstName" : "Joe",
			  "id" : "5e53effc0c3bc54dcb3b4874",
			  "lastName" : "Biden",
			  "links" : [ {
				"href" : "https://cloud.mongodb.com/api/atlas/v1.0/users/5e53effc0c3bc54dcb3b4874",
				"rel" : "self"
			  } ],
			  "roles" : [ {
				"orgId" : "5e21ce90cf09a22e3feb4646",
				"roleName" : "ORG_READ_ONLY"
			  }, {
				"orgId" : "5e21ce90cf09a22e3feb4646",
				"roleName" : "ORG_MEMBER"
			  } ],
			  "teamIds" : [ "5e4be36fc56c9837809f2f92" ],
			  "username" : "joe@example.com"
			}, {
			  "country" : "DE",
			  "emailAddress" : "sam@example.com",
			  "firstName" : "Sam",
			  "id" : "5e53f0ef2589b129f220570d",
			  "lastName" : "Valentin",
			  "links" : [ {
				"href" : "https://cloud.mongodb.com/api/atlas/v1.0/users/5e53f0ef2589b129f220570d",
				"rel" : "self"
			  } ],
			  "roles" : [ {
				"orgId" : "5e21ce90cf09a22e3feb4646",
				"roleName" : "ORG_READ_ONLY"
			  } ],
			  "teamIds" : [ ],
			  "username" : "sam@example.com"
			} ],
			"totalCount" : 2
		  }`)
	})

	orgs, _, err := client.Organizations.ListUsers(ctx, orgID, nil)

	if err != nil {
		t.Fatalf("Organizations.List returned error: %v", err)
	}

	expected := []AtlasUser{
		{
			ID:           "5e53effc0c3bc54dcb3b4874",
			EmailAddress: "joe@example.com",
			FirstName:    "Joe",
			LastName:     "Biden",
			Country:      "DE",
			Roles: []AtlasRole{
				{
					OrgID:    "5e21ce90cf09a22e3feb4646",
					RoleName: "ORG_READ_ONLY",
				},
				{
					OrgID:    "5e21ce90cf09a22e3feb4646",
					RoleName: "ORG_MEMBER",
				},
			},
			TeamIds: []string{
				"5e4be36fc56c9837809f2f92",
			},
			Username: "joe@example.com",
		},
		{
			ID:           "5e53f0ef2589b129f220570d",
			EmailAddress: "sam@example.com",
			FirstName:    "Sam",
			LastName:     "Valentin",
			Country:      "DE",
			Roles: []AtlasRole{
				{
					OrgID:    "5e21ce90cf09a22e3feb4646",
					RoleName: "ORG_READ_ONLY",
				},
			},
			TeamIds:  []string{},
			Username: "sam@example.com",
		},
	}
	if diff := deep.Equal(orgs, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizations_Rename(t *testing.T) {
	setup()
	defer teardown()

	orgID := "bcb8206e35f8619bfd6329a5"

	mux.HandleFunc(fmt.Sprintf("/orgs/%s", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
		"id" : "6001f2c580eef55aedbc6bb1",
		"links" : [ {
		  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1",
		  "rel" : "self"
		}, {
		  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1/groups",
		  "rel" : "http://mms.mongodb.com/groups"
		}, {
		  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1/teams",
		  "rel" : "http://mms.mongodb.com/teams"
		}, {
		  "href" : "https://cloud.mongodb.com/api/atlas/v1.0/orgs/6001f2c580eef55aedbc6bb1/users",
		  "rel" : "http://mms.mongodb.com/users"
		} ],
		"name" : "NewName"
	  }`)
	})

	org, _, err := client.Organizations.Rename(ctx, orgID, "NewName")
	if err != nil {
		t.Fatalf("Organizations.Rename returned error: %v", err)
	}

	expected := &Organization{
		ID:   "6001f2c580eef55aedbc6bb1",
		Name: "NewName",
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

	if diff := deep.Equal(org, expected); diff != nil {
		t.Error(diff)
	}
}

func TestOrganizations_GetAllProjects(t *testing.T) {
	setup()
	defer teardown()

	orgID := "5a0a1e7e0f2912c554080adc"

	mux.HandleFunc(fmt.Sprintf("/orgs/%s/groups", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a0a1e7e0f2912c554080adc/groups",
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

	orgs, _, err := client.Organizations.GetAllProjects(ctx, orgID, nil)

	if err != nil {
		t.Fatalf("Organizations.GetAllProjects returned error: %v", err)
	}

	expected := &Projects{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/5a0a1e7e0f2912c554080adc/groups",
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
	if diff := deep.Equal(orgs, expected); diff != nil {
		t.Error(diff)
	}
}
