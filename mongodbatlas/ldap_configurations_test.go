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

func TestLDAPConfigurations_Verify(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/userSecurity/ldap/verify", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"groupId" : "{PROJECT-ID}",
			  "request" : {
				"bindUsername" : "CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com",
				"hostname" : "atlas-ldaps-01.ldap.myteam.com",
				"port" : 636
			  },
			  "requestId" : "{REQUEST-ID}",
			  "status" : "PENDING"
		}`)
	})

	request := &LDAP{
		Hostname:     pointy.String("atlas-ldaps-01.ldap.myteam.com"),
		Port:         pointy.Int(636),
		BindUsername: pointy.String("CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com"),
		BindPassword: pointy.String("admin"),
	}
	ldap, _, err := client.LDAPConfigurations.Verify(ctx, groupID, request)
	if err != nil {
		t.Fatalf("LDAPConfigurations.Verify returned error: %v", err)
	}

	expected := &LDAPConfiguration{
		RequestID: "{REQUEST-ID}",
		GroupID:   "{PROJECT-ID}",
		Request: &LDAPRequest{
			Hostname:     "atlas-ldaps-01.ldap.myteam.com",
			Port:         636,
			BindUsername: "CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com",
		},
		Status: "PENDING",
	}

	if diff := deep.Equal(ldap, expected); diff != nil {
		t.Error(diff)
	}
}

func TestLDAPConfigurations_GetStatus(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"
	requestID := "22"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/userSecurity/ldap/verify/%s", groupID, requestID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"groupId" : "{PROJECT-ID}",
			  "request" : {
				"bindUsername" : "CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com",
				"hostname" : "atlas-ldaps-01.ldap.myteam.com",
				"port" : 636
			  },
			  "requestId" : "{REQUEST-ID}",
			  "status" : "SUCCESS",
			  "validations" : [{
				"status" : "OK",
				"validationType" : "PARSE_AUTHZ_QUERY_TEMPLATE"
			  }, {
				"status" : "OK",
				"validationType" : "QUERY_SERVER"
			  } ]
		}`)
	})

	ldap, _, err := client.LDAPConfigurations.GetStatus(ctx, groupID, requestID)
	if err != nil {
		t.Fatalf("LDAPConfigurations.GetStatus returned error: %v", err)
	}

	expected := &LDAPConfiguration{
		RequestID: "{REQUEST-ID}",
		GroupID:   "{PROJECT-ID}",
		Request: &LDAPRequest{
			Hostname:     "atlas-ldaps-01.ldap.myteam.com",
			Port:         636,
			BindUsername: "CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com",
		},
		Status: "SUCCESS",
		Validations: []*LDAPValidation{
			{
				Status:         "OK",
				ValidationType: "PARSE_AUTHZ_QUERY_TEMPLATE",
			},
			{
				Status:         "OK",
				ValidationType: "QUERY_SERVER",
			},
		},
	}

	if diff := deep.Equal(ldap, expected); diff != nil {
		t.Error(diff)
	}
}

func TestLDAPConfigurations_Save(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/userSecurity", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
			  "ldap" : {
				"authenticationEnabled" : true,
				"authorizationEnabled" : true,
				"authzQueryTemplate" : "{USER}?memberOf?base",
				"bindUsername" : "CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com",
				"hostname" : "atlas-ldaps-01.ldap.myteam.com",
				"port" : 636,
				"userToDNMapping" : [ {
				  "match" : "(.*)",
				  "substitution" : "CN={0},CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com"
				} ]
			  }
		}`)
	})

	request := &LDAPConfiguration{
		LDAP: &LDAP{
			AuthenticationEnabled: pointy.Bool(true),
			AuthorizationEnabled:  pointy.Bool(true),
			Hostname:              pointy.String("atlas-ldaps-01.ldap.myteam.com"),
			Port:                  pointy.Int(636),
			BindUsername:          pointy.String("CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com"),
			BindPassword:          pointy.String("admin"),
		},
	}
	ldap, _, err := client.LDAPConfigurations.Save(ctx, groupID, request)
	if err != nil {
		t.Fatalf("LDAPConfigurations.Save returned error: %v", err)
	}

	expected := &LDAPConfiguration{
		LDAP: &LDAP{
			AuthenticationEnabled: pointy.Bool(true),
			AuthorizationEnabled:  pointy.Bool(true),
			Hostname:              pointy.String("atlas-ldaps-01.ldap.myteam.com"),
			Port:                  pointy.Int(636),
			BindUsername:          pointy.String("CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com"),
			UserToDNMapping: []*UserToDNMapping{
				{
					Match:        "(.*)",
					Substitution: "CN={0},CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com",
				},
			},
			AuthzQueryTemplate: pointy.String("{USER}?memberOf?base"),
		},
	}

	if diff := deep.Equal(ldap, expected); diff != nil {
		t.Error(diff)
	}
}

func TestLDAPConfigurations_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/userSecurity", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			  "ldap" : {
				"authenticationEnabled" : true,
				"authorizationEnabled" : true,
				"authzQueryTemplate" : "{USER}?memberOf?base",
				"bindUsername" : "CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com",
				"hostname" : "atlas-ldaps-01.ldap.myteam.com",
				"port" : 636,
				"userToDNMapping" : [ {
				  "match" : "(.*)",
				  "substitution" : "CN={0},CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com"
				} ]
			  }
		}`)
	})

	ldap, _, err := client.LDAPConfigurations.Get(ctx, groupID)
	if err != nil {
		t.Fatalf("LDAPConfigurations.Get returned error: %v", err)
	}

	expected := &LDAPConfiguration{
		LDAP: &LDAP{
			AuthenticationEnabled: pointy.Bool(true),
			AuthorizationEnabled:  pointy.Bool(true),
			Hostname:              pointy.String("atlas-ldaps-01.ldap.myteam.com"),
			Port:                  pointy.Int(636),
			BindUsername:          pointy.String("CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com"),
			UserToDNMapping: []*UserToDNMapping{
				{
					Match:        "(.*)",
					Substitution: "CN={0},CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com",
				},
			},
			AuthzQueryTemplate: pointy.String("{USER}?memberOf?base"),
		},
	}

	if diff := deep.Equal(ldap, expected); diff != nil {
		t.Error(diff)
	}
}

func TestLDAPConfigurations_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "535683b3794d371327b"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/userSecurity/ldap/userToDNMapping", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{
			  "ldap" : {
				"authenticationEnabled" : true,
				"authorizationEnabled" : true,
				"authzQueryTemplate" : "{USER}?memberOf?base",
				"bindUsername" : "CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com",
				"hostname" : "atlas-ldaps-01.ldap.myteam.com",
				"port" : 636
			  }
		}`)
	})

	ldap, _, err := client.LDAPConfigurations.Delete(ctx, groupID)
	if err != nil {
		t.Fatalf("LDAPConfigurations.Delete returned error: %v", err)
	}

	expected := &LDAPConfiguration{
		LDAP: &LDAP{
			AuthenticationEnabled: pointy.Bool(true),
			AuthorizationEnabled:  pointy.Bool(true),
			Hostname:              pointy.String("atlas-ldaps-01.ldap.myteam.com"),
			Port:                  pointy.Int(636),
			BindUsername:          pointy.String("CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com"),
			AuthzQueryTemplate:    pointy.String("{USER}?memberOf?base"),
		},
	}

	if diff := deep.Equal(ldap, expected); diff != nil {
		t.Error(diff)
	}
}
