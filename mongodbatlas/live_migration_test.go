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

func TestLiveMigration_CreateLinkToken(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/liveMigrations/linkTokens", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"linkToken": "test"
		}`)
	})

	body := &TokenCreateRequest{
		AccessListIPs: []string{"test"},
	}

	response, _, err := client.LiveMigration.CreateLinkToken(ctx, orgID, body)
	if err != nil {
		t.Fatalf("LiveMigration.CreateLinkToken returned error: %v", err)
	}

	expected := &LinkToken{
		LinkToken: "test",
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}

func TestLiveMigration_DeleteLinkToken(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/liveMigrations/linkTokens", orgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.LiveMigration.DeleteLinkToken(ctx, orgID)
	if err != nil {
		t.Fatalf("LiveMigration.DeleteLinkToken returned error: %v", err)
	}
}

func TestLiveMigration_CreateValidation(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/liveMigrations/validate", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			  "_id": "a659ce44c03ca1e348b1e992",
			  "groupId": "1",
			  "sourceGroupId": "9b43a5b329223c3a1591a678",
			  "status": "PENDING"
		}`)
	})

	body := &LiveMigration{
		Source: &Source{
			ClusterName:           "exampleClusterA",
			GroupID:               "9b43a5b329223c3a1591a678",
			Username:              "example",
			Password:              "string",
			SSL:                   pointy.Bool(true),
			CACertificatePath:     "/path/to/ca",
			ManagedAuthentication: pointy.Bool(false),
		},
		Destination:    &Destination{ClusterName: "exampleClusterB", GroupID: "e01c9427f054fe80745f3f6c"},
		MigrationHosts: []string{"vm001.example.com"},
		DropEnabled:    pointy.Bool(true),
	}

	response, _, err := client.LiveMigration.CreateValidation(ctx, groupID, body)
	if err != nil {
		t.Fatalf("LiveMigration.CreateLinkToken returned error: %v", err)
	}

	expected := &Validation{
		ID:            "a659ce44c03ca1e348b1e992",
		GroupID:       groupID,
		Status:        "PENDING",
		SourceGroupID: "9b43a5b329223c3a1591a678",
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}

func TestLiveMigration_GetValidationStatus(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/liveMigrations/validate/%s", groupID, id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			  "_id": "a659ce44c03ca1e348b1e992",
			  "groupId": "1",
			  "sourceGroupId": "9b43a5b329223c3a1591a678",
			  "status": "PENDING"
		}`)
	})

	response, _, err := client.LiveMigration.GetValidationStatus(ctx, groupID, id)
	if err != nil {
		t.Fatalf("LiveMigration.GetValidationStatus returned error: %v", err)
	}

	expected := &Validation{
		ID:            "a659ce44c03ca1e348b1e992",
		GroupID:       groupID,
		Status:        "PENDING",
		SourceGroupID: "9b43a5b329223c3a1591a678",
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}

func TestLiveMigration_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/liveMigrations", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			  "_id": "a659ce44c03ca1e348b1e992",
			  "migrationHosts": [ "vm001.example.com" ],
			  "status": "PENDING"
		}`)
	})

	body := &LiveMigration{
		Source: &Source{
			ClusterName:           "exampleClusterA",
			GroupID:               "9b43a5b329223c3a1591a678",
			Username:              "example",
			Password:              "string",
			SSL:                   pointy.Bool(true),
			CACertificatePath:     "/path/to/ca",
			ManagedAuthentication: pointy.Bool(false),
		},
		Destination:    &Destination{ClusterName: "exampleClusterB", GroupID: "e01c9427f054fe80745f3f6c"},
		MigrationHosts: []string{"vm001.example.com"},
		DropEnabled:    pointy.Bool(true),
	}

	response, _, err := client.LiveMigration.Create(ctx, groupID, body)
	if err != nil {
		t.Fatalf("LiveMigration.Create returned error: %v", err)
	}

	expected := &LiveMigration{
		ID:             "a659ce44c03ca1e348b1e992",
		Status:         "PENDING",
		MigrationHosts: []string{"vm001.example.com"},
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}

func TestLiveMigration_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/liveMigrations/%s", groupID, id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"_id": "a659ce44c03ca1e348b1e992",
			"migrationHosts": [ "vm001.example.com" ],
			"status": "PENDING",
			"readyForCutover": false
		}`)
	})

	response, _, err := client.LiveMigration.Get(ctx, groupID, id)
	if err != nil {
		t.Fatalf("LiveMigration.Get returned error: %v", err)
	}

	expected := &LiveMigration{
		ID:              "a659ce44c03ca1e348b1e992",
		Status:          "PENDING",
		MigrationHosts:  []string{"vm001.example.com"},
		ReadyForCutover: pointy.Bool(false),
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}

func TestLiveMigration_Start(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/liveMigrations/%s/cutover", groupID, id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{
		  "_id": "a659ce44c03ca1e348b1e992",
		  "groupId": "7e68e12770616e75e6df43d0",
		  "sourceGroupId": "7e68e12770616e75e6df43d0",
		  "status": "PENDING"
		}`)
	})

	response, _, err := client.LiveMigration.Start(ctx, groupID, id)
	if err != nil {
		t.Fatalf("LiveMigration.Start returned error: %v", err)
	}

	expected := &Validation{
		ID:            "a659ce44c03ca1e348b1e992",
		Status:        "PENDING",
		GroupID:       "7e68e12770616e75e6df43d0",
		SourceGroupID: "7e68e12770616e75e6df43d0",
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}
