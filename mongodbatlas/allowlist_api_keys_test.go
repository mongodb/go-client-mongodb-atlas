package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestAllowlistAPIKeys_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "ORG-ID"
	apiKeyID := "API-KEY-ID"

	mux.HandleFunc(fmt.Sprintf("/"+allowlistAPIKeysPath, orgID, apiKeyID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/whitelist/?pretty=true&pageNum=1&itemsPerPage=100",
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
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/whitelist/147.58.184.16",
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
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/whitelist/206.252.195.126",
							"rel": "self"
						}
					]
				}
			],
			"totalCount": 2
		}`)
	})

	allowlistAPIKeys, _, err := client.AllowlistAPIKeys.List(ctx, orgID, apiKeyID, nil)
	if err != nil {
		t.Fatalf("AllowlistAPIKeys.List returned error: %v", err)
	}

	expected := &AllowlistAPIKeys{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/whitelist/?pretty=true&pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*AllowlistAPIKey{
			{
				CidrBlock:       "147.58.184.16/32",
				Count:           0,
				Created:         "2019-01-24T16:34:57Z",
				IPAddress:       "147.58.184.16",
				LastUsed:        "2019-01-24T20:18:25Z",
				LastUsedAddress: "147.58.184.16",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/whitelist/147.58.184.16",
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
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/whitelist/206.252.195.126",
						Rel:  "self",
					},
				},
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(allowlistAPIKeys, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAllowlistAPIKeys_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "ORG-ID"
	apiKeyID := "API-KEY-ID"
	ipAddress := "IP-ADDRESS"

	mux.HandleFunc(fmt.Sprintf("/"+allowlistAPIKeysPath+"/%s", orgID, apiKeyID, ipAddress), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"cidrBlock": "147.58.184.16/32",
			"count": 0,
			"created": "2019-01-24T16:34:57Z",
			"ipAddress": "147.58.184.16",
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/whitelist/147.58.184.16",
					"rel": "self"
				}
			]
		}`)
	})

	allowlistAPIKey, _, err := client.AllowlistAPIKeys.Get(ctx, orgID, apiKeyID, ipAddress)
	if err != nil {
		t.Fatalf("AllowlistAPIKeys.Get returned error: %v", err)
	}

	expected := &AllowlistAPIKey{
		CidrBlock: "147.58.184.16/32",
		Count:     0,
		Created:   "2019-01-24T16:34:57Z",
		IPAddress: "147.58.184.16",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/whitelist/147.58.184.16",
				Rel:  "self",
			},
		},
	}

	if diff := deep.Equal(allowlistAPIKey, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAllowlistAPIKeys_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "ORG-ID"
	apiKeyID := "API-KEY-ID"

	createRequest := []*AllowlistAPIKeysReq{
		{
			IPAddress: "77.54.32.11",
			CidrBlock: "77.54.32.11/32",
		},
	}

	mux.HandleFunc(fmt.Sprintf("/"+allowlistAPIKeysPath, orgID, apiKeyID), func(w http.ResponseWriter, r *http.Request) {
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
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/whitelist/?pretty=true&pageNum=1&itemsPerPage=100",
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
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/whitelist/147.58.184.16",
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
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/whitelist/77.54.32.11",
							"rel": "self"
						}
					]
				}
			],
			"totalCount": 2
		}`)
	})

	allowlistAPIKey, _, err := client.AllowlistAPIKeys.Create(ctx, orgID, apiKeyID, createRequest)
	if err != nil {
		t.Fatalf("AllowlistAPIKeys.Create returned error: %v", err)
	}

	expected := &AllowlistAPIKeys{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/whitelist/?pretty=true&pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*AllowlistAPIKey{
			{
				CidrBlock:       "147.58.184.16/32",
				Count:           0,
				Created:         "2019-01-24T16:34:57Z",
				IPAddress:       "147.58.184.16",
				LastUsed:        "2019-01-24T20:18:25Z",
				LastUsedAddress: "147.58.184.16",
				Links: []*Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/whitelist/147.58.184.16",
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
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/{ORG-ID}/apiKeys/{API-KEY-ID}/whitelist/77.54.32.11",
						Rel:  "self",
					},
				},
			},
		},
		TotalCount: 2,
	}

	if diff := deep.Equal(allowlistAPIKey, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAllowlistAPIKeys_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	orgID := "ORG-ID"
	apiKeyID := "API-KEY-ID"
	ipAddress := "IP-ADDRESS"

	mux.HandleFunc(fmt.Sprintf("/"+allowlistAPIKeysPath+"/%s", orgID, apiKeyID, ipAddress), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.AllowlistAPIKeys.Delete(ctx, orgID, apiKeyID, ipAddress)
	if err != nil {
		t.Fatalf("AllowlistAPIKeys.Delete returned error: %v", err)
	}
}
