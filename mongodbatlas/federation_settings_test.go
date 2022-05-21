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
)

const federationSettingsID = "5e8cc670a16506712e0b1e95"

const federationSettingsOrgID = "5a0a1e7e0f2912c554080adc"

func TestFederatedSettingsServiceOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/orgs/%s/federationSettings", federationSettingsOrgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
			"federatedDomains": [
			  "example.com"
			],
			"hasRoleMappings": false,
			"id": "5e8cc670a16506712e0b1e95",
			"identityProviderId": "0oa8i0grsgbwDiIyw453",
			"identityProviderStatus": "INACTIVE"
		  }`)
	})

	response, _, err := client.FederatedSettings.Get(ctx, federationSettingsOrgID)
	if err != nil {
		t.Fatalf("FederatedSettings.Get returned error: %v", err)
	}

	HasRoleMappings := false

	expected := &FederatedSettings{
		FederatedDomains:       []string{"example.com"},
		ID:                     "5e8cc670a16506712e0b1e95",
		HasRoleMappings:        &HasRoleMappings,
		IdentityProviderID:     "0oa8i0grsgbwDiIyw453",
		IdentityProviderStatus: "INACTIVE",
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}

func TestFederatedSettingsServiceOp_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/federationSettings/%s", federationSettingsID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.FederatedSettings.Delete(ctx, federationSettingsID)
	if err != nil {
		t.Fatalf("federationSettings.Delete returned error: %v", err)
	}
}
