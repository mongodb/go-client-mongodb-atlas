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
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/openlyinc/pointy"
)

func TestX509AuthDBUsers_CreateUserCertificate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	username := "username_test"
	monthsUntilExpiration := 4

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/databaseUsers/%s/certs", groupID, username), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"monthsUntilExpiration": float64(4),
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Request body\n got=%#v\nwant=%#v \n\ndiff=%#v", v, expected, diff)
		}

		fmt.Fprint(w, "-BEGIN CERTIFICATE--MIIFCTCCAvGgAwIBAgIIb--END CERTIFICATE---BEGIN PRIVATE KEY--MIIJQgIBADANBgkqhkiG==--END PRIVATE KEY--")
	})

	x509AuthDBUser, _, err := client.X509AuthDBUsers.CreateUserCertificate(ctx, groupID, username, monthsUntilExpiration)
	if err != nil {
		t.Errorf("X509AuthDBUsers.CreateUserCertificate returned error: %v", err)
		return
	}

	expected := &UserCertificate{
		Username:              "username_test",
		MonthsUntilExpiration: 4,
		Certificate:           "-BEGIN CERTIFICATE--MIIFCTCCAvGgAwIBAgIIb--END CERTIFICATE---BEGIN PRIVATE KEY--MIIJQgIBADANBgkqhkiG==--END PRIVATE KEY--",
	}

	if !reflect.DeepEqual(x509AuthDBUser, expected) {
		t.Errorf("X509AuthDBUsers.CreateUserCertificate\n got=%#v\nwant=%#v", x509AuthDBUser, expected)
	}
}

func TestX509AuthDBUsers_GetUserCertificates(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	username := "username_test"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/databaseUsers/%s/certs", groupID, username), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"links": [
				{
					"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/ec13d5f09d521a5206e5d726/databaseUsers/myX509User/certs?pretty=true&pageNum=1&itemsPerPage=100",
					"rel": "self"
				}
			],
			"results": [
				{
					"_id": 5433027191242719574,
					"createdAt": "2019-12-03T21:00:42Z",
					"groupId": "ec13d5f09d521a5206e5d726",
					"notAfter": "2020-06-03T22:00:42Z",
					"subject": "CN=myX509User"
				},
				{
					"_id": 946028664465903704,
					"createdAt": "2019-12-05T13:50:49Z",
					"groupId": "ec13d5f09d521a5206e5d726",
					"notAfter": "2020-03-05T14:50:49Z",
					"subject": "CN=myX509User"
				}
			],
			"totalCount": 2
		}`)
	})

	x509AuthDBUserCertificates, _, err := client.X509AuthDBUsers.GetUserCertificates(ctx, groupID, username, nil)
	if err != nil {
		t.Errorf("X509AuthDBUsers.GetUserCertificates returned error: %v", err)
	}

	expected := []UserCertificate{
		{
			ID:        pointy.Int64(5433027191242719574),
			CreatedAt: "2019-12-03T21:00:42Z",
			GroupID:   "ec13d5f09d521a5206e5d726",
			NotAfter:  "2020-06-03T22:00:42Z",
			Subject:   "CN=myX509User",
		},
		{
			ID:        pointy.Int64(946028664465903704),
			CreatedAt: "2019-12-05T13:50:49Z",
			GroupID:   "ec13d5f09d521a5206e5d726",
			NotAfter:  "2020-03-05T14:50:49Z",
			Subject:   "CN=myX509User",
		},
	}

	if !reflect.DeepEqual(x509AuthDBUserCertificates, expected) {
		t.Errorf("X509AuthDBUsers.GetUserCertificates\n got=%#v\nwant=%#v", x509AuthDBUserCertificates, expected)
	}
}

func TestX509AuthDBUsers_SaveConfiguration(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	customerX509Req := &CustomerX509{
		Cas: "-BEGIN CERTIFICATE--MIIFCTCCAvGgAwIBAgIIb--END CERTIFICATE---BEGIN PRIVATE KEY--MIIJQgIBADANBgkqhkiG==--END PRIVATE KEY--",
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/userSecurity", groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"customerX509": map[string]interface{}{
				"cas": "-BEGIN CERTIFICATE--MIIFCTCCAvGgAwIBAgIIb--END CERTIFICATE---BEGIN PRIVATE KEY--MIIJQgIBADANBgkqhkiG==--END PRIVATE KEY--",
			},
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Request body\n got=%#v\nwant=%#v \n\ndiff=%#v", v, expected, diff)
		}

		fmt.Fprint(w, `{
			"customerX509": {
				"cas": "-BEGIN CERTIFICATE--MIIFCTCCAvGgAwIBAgIIb--END CERTIFICATE---BEGIN PRIVATE KEY--MIIJQgIBADANBgkqhkiG==--END PRIVATE KEY--"
			}
		}`)
	})

	customerX509Conf, _, err := client.X509AuthDBUsers.SaveConfiguration(ctx, groupID, customerX509Req)
	if err != nil {
		t.Errorf("X509AuthDBUsers.SaveConfiguration returned error: %v", err)
		return
	}

	if !reflect.DeepEqual(customerX509Conf, customerX509Req) {
		t.Errorf("X509AuthDBUsers.SaveConfiguration\n got=%#v\nwant=%#v", customerX509Conf, customerX509Req)
	}
}

func TestX509AuthDBUsers_GetCurrentX509Conf(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/userSecurity", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"ldap": {},
			"customerX509": {
				"cas": "-BEGIN CERTIFICATE--MIIFCTCCAvGgAwIBAgIIb--END CERTIFICATE---BEGIN PRIVATE KEY--MIIJQgIBADANBgkqhkiG==--END PRIVATE KEY--"
			}
		}`)
	})

	currentX509Conf, _, err := client.X509AuthDBUsers.GetCurrentX509Conf(ctx, groupID)
	if err != nil {
		t.Errorf("X509AuthDBUsers.GetCurrentX509Conf returned error: %v", err)
	}

	expected := &CustomerX509{
		Cas: "-BEGIN CERTIFICATE--MIIFCTCCAvGgAwIBAgIIb--END CERTIFICATE---BEGIN PRIVATE KEY--MIIJQgIBADANBgkqhkiG==--END PRIVATE KEY--",
	}

	if !reflect.DeepEqual(currentX509Conf, expected) {
		t.Errorf("X509AuthDBUsers.GetCurrentX509Conf\n got=%#v\nwant=%#v", currentX509Conf, expected)
	}
}

func TestX509AuthDBUsers_DisableCustomerX509(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/userSecurity/customerX509", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.X509AuthDBUsers.DisableCustomerX509(ctx, groupID)
	if err != nil {
		t.Errorf("X509AuthDBUsers.DisableCustomerX509 returned error: %v", err)
	}
}
