package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestThirdPartyIntegration_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "ORG-ID"

	mux.HandleFunc(fmt.Sprintf("/"+integrationBasePath, projectID), func(w http.ResponseWriter, r *http.Request) {
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
             "apiKey": "112233",
             "region": "US",
             "type": "DATADOG"
         },
         {
             "apiToken": "112233",
             "channelName": "My Channel",
             "teamName": "My Team",
             "type": "SLACK"
         },
         {
             "MicrosoftTeamsWebhookURL": "https://teams.microsoft.com/webhook",
             "Name": "MyTeam",
             "type": "MICROSOFT_TEAMS"
         },
         {
			 "enabled": true,
			 "scheme": "http",
             "serviceDiscovery": "112233",
             "password": "myPassword",
             "username": "myUser",
             "type": "PROMETHEUS"
         }
			],
			"totalCount": 4
		}`)
	})

	integrationAPIKeys, _, err := client.Integrations.List(ctx, projectID)
	if err != nil {
		t.Fatalf("Integrations.List returned error: %v", err)
	}

	expected := &ThirdPartyIntegrations{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/whitelist/?pretty=true&pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*ThirdPartyIntegration{
			{
				Type:   "DATADOG",
				APIKey: "112233",
				Region: "US",
			},
			{
				Type:        "SLACK",
				APIToken:    "112233",
				TeamName:    "My Team",
				ChannelName: "My Channel",
			},
			{
				Type:                     "MICROSOFT_TEAMS",
				Name:                     "MyTeam",
				MicrosoftTeamsWebhookURL: "https://teams.microsoft.com/webhook",
			},
			{
				Type:             "PROMETHEUS",
				UserName:         "myUser",
				Password:         "myPassword",
				ServiceDiscovery: "112233",
				Scheme:           "http",
				Enabled:          true,
			},
		},
		TotalCount: 4,
	}

	if diff := deep.Equal(integrationAPIKeys, expected); diff != nil {
		t.Error(diff)
	}
}

func TestThirdPartyIntegration_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "ORG-ID"
	integrationType := "DATADOG"

	mux.HandleFunc(fmt.Sprintf("/"+integrationBasePath+"/%s", projectID, integrationType), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			
    "apiKey": "112233",
    "region": "US",
    "type": "DATADOG"

		}`)
	})

	integrationAPIKeys, _, err := client.Integrations.Get(ctx, projectID, integrationType)
	if err != nil {
		t.Fatalf("Integrations.Get returned error: %v", err)
	}

	expected := &ThirdPartyIntegration{
		Type:   "DATADOG",
		APIKey: "112233",
		Region: "US",
	}

	if diff := deep.Equal(integrationAPIKeys, expected); diff != nil {
		t.Error(diff)
	}
}

func TestThirdPartyIntegration_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "ORG-ID"
	integrationType := "DATADOG"

	mux.HandleFunc(fmt.Sprintf("/"+integrationBasePath+"/%s", projectID, integrationType), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Integrations.Delete(ctx, projectID, integrationType)
	if err != nil {
		t.Fatalf("WhitelistAPIKeys.Delete returned error: %v", err)
	}
}

func TestThirdPartyIntegration_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "ORG-ID"
	integrationType := "DATADOG"

	mux.HandleFunc(fmt.Sprintf("/"+integrationBasePath+"/%s", projectID, integrationType), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/whitelist/?pretty=true&pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				         {
             "apiKey": "112233",
             "region": "US",
             "type": "DATADOG"
         },
         {
             "apiToken": "112233",
             "channelName": "My Channel",
             "teamName": "My Team",
             "type": "SLACK"
         },
		 {
			"MicrosoftTeamsWebhookURL": "https://teams.microsoft.com/webhook",
			"Name": "MyTeam",
			"type": "MICROSOFT_TEAMS"
		},
		{
			"enabled": true,
			"scheme": "http",
			"serviceDiscovery": "112233",
			"password": "myPassword",
			"username": "myUser",
			"type": "PROMETHEUS"
		}
			],
			"totalCount": 4
		}`)
	})

	service := &ThirdPartyIntegration{
		Type:   "DATADOG",
		APIKey: "112233",
		Region: "US",
	}

	integrationAPIKeys, _, err := client.Integrations.Create(ctx, projectID, integrationType, service)
	if err != nil {
		t.Fatalf("Integrations.Create returned error: %v", err)
	}

	expected := &ThirdPartyIntegrations{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/whitelist/?pretty=true&pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*ThirdPartyIntegration{
			{
				Type:   "DATADOG",
				APIKey: "112233",
				Region: "US",
			},
			{
				Type:        "SLACK",
				APIToken:    "112233",
				TeamName:    "My Team",
				ChannelName: "My Channel",
			},
			{
				Type:                     "MICROSOFT_TEAMS",
				Name:                     "MyTeam",
				MicrosoftTeamsWebhookURL: "https://teams.microsoft.com/webhook",
			},
			{
				Type:             "PROMETHEUS",
				UserName:         "myUser",
				Password:         "myPassword",
				ServiceDiscovery: "112233",
				Scheme:           "http",
				Enabled:          true,
			},
		},
		TotalCount: 4,
	}

	if diff := deep.Equal(integrationAPIKeys, expected); diff != nil {
		t.Error(diff)
	}
}

func TestThirdPartyIntegration_Replace(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "ORG-ID"
	integrationType := "DATADOG"

	mux.HandleFunc(fmt.Sprintf("/"+integrationBasePath+"/%s", projectID, integrationType), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/whitelist/?pretty=true&pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				         {
             "apiKey": "112233",
             "region": "US",
             "type": "DATADOG"
         },
         {
             "apiToken": "112233",
             "channelName": "My Channel",
             "teamName": "My Team",
             "type": "SLACK"
         },
		 {
			"MicrosoftTeamsWebhookURL": "https://teams.microsoft.com/webhook",
			"Name": "MyTeam",
			"type": "MICROSOFT_TEAMS"
		},
		{
			"enabled": true,
			"scheme": "http",
			"serviceDiscovery": "112233",
			"password": "myPassword",
			"username": "myUser",
			"type": "PROMETHEUS"
		}
			],
			"totalCount": 4
		}`)
	})

	service := &ThirdPartyIntegration{
		Type:   "DATADOG",
		APIKey: "112233",
		Region: "US",
	}

	integrationAPIKeys, _, err := client.Integrations.Replace(ctx, projectID, integrationType, service)
	if err != nil {
		t.Fatalf("Integrations.Replace returned error: %v", err)
	}

	expected := &ThirdPartyIntegrations{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs/599c510c80eef518f3b63fe1/apiKeys/5c49e72980eef544a218f8f8/whitelist/?pretty=true&pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*ThirdPartyIntegration{
			{
				Type:   "DATADOG",
				APIKey: "112233",
				Region: "US",
			},
			{
				Type:        "SLACK",
				APIToken:    "112233",
				TeamName:    "My Team",
				ChannelName: "My Channel",
			},
			{
				Type:                     "MICROSOFT_TEAMS",
				Name:                     "MyTeam",
				MicrosoftTeamsWebhookURL: "https://teams.microsoft.com/webhook",
			},
			{
				Type:             "PROMETHEUS",
				UserName:         "myUser",
				Password:         "myPassword",
				ServiceDiscovery: "112233",
				Scheme:           "http",
				Enabled:          true,
			},
		},
		TotalCount: 4,
	}

	if diff := deep.Equal(integrationAPIKeys, expected); diff != nil {
		t.Error(diff)
	}
}
