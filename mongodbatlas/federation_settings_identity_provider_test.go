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
	"time"

	"github.com/go-test/deep"
	"github.com/openlyinc/pointy"
)

const ipOrgID = "1234567890abcdefghij"
const identityProviderfederationSettingsID = "5a0a1e7e0f2912c554080adc"

var ta = time.Date(2035, time.September, 29, 15, 3, 55, 0, time.UTC)
var tb = time.Date(2022, time.January, 20, 15, 3, 55, 0, time.UTC)

func TestFederatedSettingsIdentityProviderOp_List(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		client, mux, teardown := setup()
		defer teardown()
		mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/federationSettings/%s/identityProviders", identityProviderfederationSettingsID), func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			_, _ = fmt.Fprint(w, `{
						"links" : [ {
				"href" : "https://cloud.mongodb.com/api/public/v1.0/federationSettings/5a0a1e7e0f2912c554080adc/identityProviders?pretty=true&pageNum=1&itemsPerPage=100",
						"rel": "self"
					}
				],
				"results": [
					{
						"acsUrl" : "https://example.mongodb.com/sso/saml2/12345678901234567890",
						"associatedDomains" : [ ],
						"associatedOrgs" : [ ],
						"audienceUri" : "https://www.example.com/saml2/service-provider/abcdefghij1234567890",
						"displayName" : "Test",
						"issuerUri" : "urn:123456789000.us.provider.com",
						"oktaIdpId" : "1234567890abcdefghij",
						"pemFileInfo" : {
							"certificates" : [ {
							   "notAfter" : "2035-09-29T15:03:55Z",
							   "notBefore" : "2022-01-20T15:03:55Z"
							} ],
							"fileName" : "file.pem"
						},
						"requestBinding" : "HTTP-POST",
						"responseSignatureAlgorithm" : "SHA-256",
						"ssoDebugEnabled" : true,
						"ssoUrl" : "https://123456789000.us.provider.com/samlp/12345678901234567890123456789012",
						"status" : "INACTIVE"
					} ],
				"totalCount": 1
			   }`)
		})

		identityProviders, _, err := client.FederatedSettings.ListIdentityProviders(ctx, identityProviderfederationSettingsID, nil)
		if err != nil {
			t.Fatalf("FederatedSettingsIdentityProvider.List returned error: %v", err)
		}

		expected := []FederatedSettingsIdentityProvider{
			{
				AcsURL:            "https://example.mongodb.com/sso/saml2/12345678901234567890",
				AssociatedDomains: []string{},
				AssociatedOrgs:    []*AssociatedOrgs{},
				AudienceURI:       "https://www.example.com/saml2/service-provider/abcdefghij1234567890",
				DisplayName:       "Test",
				IssuerURI:         "urn:123456789000.us.provider.com",
				OktaIdpID:         "1234567890abcdefghij",
				PemFileInfo: &PemFileInfo{
					Certificates: []*Certificates{
						{
							NotAfter:  ta,
							NotBefore: tb,
						},
					},
					FileName: "file.pem",
				},
				RequestBinding:             "HTTP-POST",
				ResponseSignatureAlgorithm: "SHA-256",
				SsoDebugEnabled:            pointy.Bool(true),
				SsoURL:                     "https://123456789000.us.provider.com/samlp/12345678901234567890123456789012",
				Status:                     "INACTIVE"},
		}

		if diff := deep.Equal(identityProviders, expected); diff != nil {
			t.Error(diff)
		}
	})
}

func TestFederatedSettingsIdentityProviderOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/federationSettings/%s/identityProviders/%s", identityProviderfederationSettingsID, ipOrgID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
			"acsUrl" : "https://example.mongodb.com/sso/saml2/12345678901234567890",
			"associatedDomains" : [ ],
			"associatedOrgs" : [ ],
			"audienceUri" : "https://www.example.com/saml2/service-provider/abcdefghij1234567890",
			"displayName" : "Test",
			"issuerUri" : "urn:123456789000.us.provider.com",
			"oktaIdpId" : "1234567890abcdefghij",
			"pemFileInfo" : {
				"certificates" : [ {
				   "notAfter" : "2035-09-29T15:03:55Z",
				   "notBefore" : "2022-01-20T15:03:55Z"
				} ],
				"fileName" : "file.pem"
			},
			"requestBinding" : "HTTP-POST",
			"responseSignatureAlgorithm" : "SHA-256",
			"ssoDebugEnabled" : true,
			"ssoUrl" : "https://123456789000.us.provider.com/samlp/12345678901234567890123456789012",
			"status" : "INACTIVE"
		}`)
	})

	response, _, err := client.FederatedSettings.GetIdentityProvider(ctx, identityProviderfederationSettingsID, ipOrgID)
	if err != nil {
		t.Fatalf("FederatedSettingsIdentityProvider.Get returned error: %v", err)
	}

	expected := &FederatedSettingsIdentityProvider{
		AcsURL:            "https://example.mongodb.com/sso/saml2/12345678901234567890",
		AssociatedDomains: []string{},
		AssociatedOrgs:    []*AssociatedOrgs{},
		PemFileInfo: &PemFileInfo{
			Certificates: []*Certificates{
				{
					NotAfter:  ta,
					NotBefore: tb,
				},
			},
			FileName: "file.pem",
		},
		AudienceURI:                "https://www.example.com/saml2/service-provider/abcdefghij1234567890",
		DisplayName:                "Test",
		IssuerURI:                  "urn:123456789000.us.provider.com",
		OktaIdpID:                  "1234567890abcdefghij",
		RequestBinding:             "HTTP-POST",
		ResponseSignatureAlgorithm: "SHA-256",
		SsoDebugEnabled:            pointy.Bool(true),
		SsoURL:                     "https://123456789000.us.provider.com/samlp/12345678901234567890123456789012",
		Status:                     "INACTIVE",
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}
