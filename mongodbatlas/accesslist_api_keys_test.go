package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestAccessListAPIKeys_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "ORG-ID"
	apiKeyID := "API-KEY-ID"

	mux.HandleFunc(fmt.Sprintf("/"+accessListAPIKeysPath, orgID, apiKeyID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/accesslist/?pretty=true&pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				{
					"cidrBlock": "147.58.184.16/32",
					"count": 0,
					"created": "2019-01-24T16:34:57Z",
					"ipAddress": "147.58.184.16",
					"lastUsed": "2019-01-24T20:18:25Z",
					"lastUsedAddress": "147.58.184.16",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist/147.58.184.16",
							"rel": "self"
						}
					]
				},
				{
					"cidrBlock": "84.255.48.125/32",
					"count": 0,
					"created": "2019-01-24T16:26:37Z",
					"ipAddress": "84.255.48.125",
					"lastUsed": "2019-01-24T20:18:25Z",
					"lastUsedAddress": "84.255.48.125",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist/206.252.195.126",
							"rel": "self"
						}
					]
				}
			],
			"totalCount": 2
		}`)
	})

	accessListAPIKeys, _, err := client.AccessListAPIKeys.List(ctx, orgID, apiKeyID, nil)
	if err != nil {
		t.Fatalf("AccessListAPIKeys.List returned error: %v", err)
	}

	expected := &AccessListAPIKeys{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/accesslist/?pretty=true&pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*AccessListAPIKey{
			{
				CidrBlock:       "147.58.184.16/32",
				Count:           0,
				Created:         "2019-01-24T16:34:57Z",
				IPAddress:       "147.58.184.16",
				LastUsed:        "2019-01-24T20:18:25Z",
				LastUsedAddress: "147.58.184.16",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist/147.58.184.16",
						Rel:  "self",
					},
				},
			},
			{
				CidrBlock:       "84.255.48.125/32",
				Count:           0,
				Created:         "2019-01-24T16:26:37Z",
				IPAddress:       "84.255.48.125",
				LastUsed:        "2019-01-24T20:18:25Z",
				LastUsedAddress: "84.255.48.125",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist/206.252.195.126",
						Rel:  "self",
					},
				},
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(accessListAPIKeys, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAccessListAPIKeys_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "ORG-ID"
	apiKeyID := "API-KEY-ID"
	ipAddress := "IP-ADDRESS"

	mux.HandleFunc(fmt.Sprintf("/"+accessListAPIKeysPath+"/%s", orgID, apiKeyID, ipAddress), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"cidrBlock": "147.58.184.16/32",
			"count": 0,
			"created": "2019-01-24T16:34:57Z",
			"ipAddress": "147.58.184.16",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist/147.58.184.16",
					"rel": "self"
				}
			]
		}`)
	})

	accesslistAPIKey, _, err := client.AccessListAPIKeys.Get(ctx, orgID, apiKeyID, ipAddress)
	if err != nil {
		t.Fatalf("AccessListAPIKeys.Get returned error: %v", err)
	}

	expected := &AccessListAPIKey{
		CidrBlock: "147.58.184.16/32",
		Count:     0,
		Created:   "2019-01-24T16:34:57Z",
		IPAddress: "147.58.184.16",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist/147.58.184.16",
				Rel:  "self",
			},
		},
	}

	if diff := deep.Equal(accesslistAPIKey, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAccessListAPIKeys_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "ORG-ID"
	apiKeyID := "API-KEY-ID"

	createRequest := []*AccessListAPIKeysReq{
		{
			IPAddress: "77.54.32.11",
			CidrBlock: "77.54.32.11/32",
		},
	}

	mux.HandleFunc(fmt.Sprintf("/"+accessListAPIKeysPath, orgID, apiKeyID), func(w http.ResponseWriter, r *http.Request) {
		expected := []map[string]interface{}{
			{
				"ipAddress": "77.54.32.11",
				"cidrBlock": "77.54.32.11/32",
			},
		}

		var v []map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("Decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/accesslist/?pretty=true&pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				{
					"cidrBlock": "147.58.184.16/32",
					"count": 0,
					"created": "2019-01-24T16:34:57Z",
					"ipAddress": "147.58.184.16",
					"lastUsed": "2019-01-24T20:18:25Z",
					"lastUsedAddress": "147.58.184.16",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist/147.58.184.16",
							"rel": "self"
						}
					]
				},
				{
					"cidrBlock": "77.54.32.11/32",
					"count": 0,
					"created": "2019-01-24T16:26:37Z",
					"ipAddress": "77.54.32.11",
					"lastUsed": "2019-01-24T20:18:25Z",
					"lastUsedAddress": "77.54.32.11",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist/77.54.32.11",
							"rel": "self"
						}
					]
				}
			],
			"totalCount": 2
		}`)
	})

	accesslistAPIKey, _, err := client.AccessListAPIKeys.Create(ctx, orgID, apiKeyID, createRequest)
	if err != nil {
		t.Fatalf("AccessListAPIKeys.Create returned error: %v", err)
	}

	expected := &AccessListAPIKeys{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/accesslist/?pretty=true&pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*AccessListAPIKey{
			{
				CidrBlock:       "147.58.184.16/32",
				Count:           0,
				Created:         "2019-01-24T16:34:57Z",
				IPAddress:       "147.58.184.16",
				LastUsed:        "2019-01-24T20:18:25Z",
				LastUsedAddress: "147.58.184.16",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist/147.58.184.16",
						Rel:  "self",
					},
				},
			},
			{
				CidrBlock:       "77.54.32.11/32",
				Count:           0,
				Created:         "2019-01-24T16:26:37Z",
				IPAddress:       "77.54.32.11",
				LastUsed:        "2019-01-24T20:18:25Z",
				LastUsedAddress: "77.54.32.11",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/accesslist/77.54.32.11",
						Rel:  "self",
					},
				},
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(accesslistAPIKey, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAccessListAPIKeys_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "ORG-ID"
	apiKeyID := "API-KEY-ID"
	ipAddress := "IP-ADDRESS"

	mux.HandleFunc(fmt.Sprintf("/"+accessListAPIKeysPath+"/%s", orgID, apiKeyID, ipAddress), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.AccessListAPIKeys.Delete(ctx, orgID, apiKeyID, ipAddress)
	if err != nil {
		t.Fatalf("AccessListAPIKeys.Delete returned error: %v", err)
	}
}
