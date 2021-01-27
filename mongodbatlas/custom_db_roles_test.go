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
)

func TestCustomDBRoles_ListCustomDBRoles(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/groups/1/customDBRoles/roles", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"actions":[{"action":"CREATE_INDEX","resources":[{"collection":"test-collection","db":"test-db"}]}],"inheritedRoles":[{"db":"test-db","role":"read"}],"roleName":"test-role-name"}]`)
	})

	customDBRoles, _, err := client.CustomDBRoles.List(ctx, "1", nil)
	if err != nil {
		t.Fatalf("CustomDBRoles.List returned error: %v", err)
	}

	expected := &[]CustomDBRole{{
		Actions: []Action{{
			Action: "CREATE_INDEX",
			Resources: []Resource{{
				Collection: "test-collection",
				Db:         "test-db",
			}},
		}},
		InheritedRoles: []InheritedRole{{
			Db:   "test-db",
			Role: "read",
		}},
		RoleName: "test-role-name",
	}}

	if diff := deep.Equal(customDBRoles, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCustomDBRoles_GetCustomDBRole(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/groups/1/customDBRoles/roles/test-role-name", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"actions":[{"action":"CREATE_INDEX","resources":[{"collection":"test-collection","db":"test-db"}]}],"inheritedRoles":[{"db":"test-db","role":"read"}],"roleName":"test-role-name"}`)
	})

	customDBRole, _, err := client.CustomDBRoles.Get(ctx, "1", "test-role-name")
	if err != nil {
		t.Fatalf("CustomDBRoles.Get returned error: %v", err)
	}

	expected := &CustomDBRole{
		Actions: []Action{{
			Action: "CREATE_INDEX",
			Resources: []Resource{{
				Collection: "test-collection",
				Db:         "test-db",
			}},
		}},
		InheritedRoles: []InheritedRole{{
			Db:   "test-db",
			Role: "read",
		}},
		RoleName: "test-role-name",
	}
	if diff := deep.Equal(customDBRole, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCustomDBRoles_CreateCustomDBRole(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &CustomDBRole{
		Actions: []Action{{
			Action: "CREATE_INDEX",
			Resources: []Resource{{
				Collection: "test-collection",
				Db:         "test-db",
			}},
		}},
		InheritedRoles: []InheritedRole{{
			Db:   "test-db",
			Role: "read",
		}},
		RoleName: "test-role-name",
	}

	mux.HandleFunc("/groups/1/customDBRoles/roles", func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"actions": []interface{}{map[string]interface{}{
				"action": "CREATE_INDEX",
				"resources": []interface{}{map[string]interface{}{
					"collection": "test-collection",
					"db":         "test-db",
				}},
			}},
			"inheritedRoles": []interface{}{map[string]interface{}{
				"db":   "test-db",
				"role": "read",
			}},
			"roleName": "test-role-name",
		}

		jsonBlob := `
		{
			"actions": [
				{
					"action": "CREATE_INDEX",
					"resources": [
						{
							"collection": "test-collection",
							"db": "test-db"
						}
					]
				}
			],
			"inheritedRoles": [
				{
					"db": "test-db",
					"role": "read"
				}
			],
			"roleName":"test-role-name"
		}
		`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, jsonBlob)
	})

	customDBRole, _, err := client.CustomDBRoles.Create(ctx, "1", createRequest)
	if err != nil {
		t.Fatalf("CustomDBRoles.Create returned error: %v", err)
	}

	if roleName := customDBRole.RoleName; roleName != "test-role-name" {
		t.Errorf("expected roleName '%s', received '%s'", "test-role-name", roleName)
	}
}

func TestCustomDBRoles_UpdateCustomDBRole(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &CustomDBRole{
		Actions: []Action{{
			Action: "CREATE_INDEX",
			Resources: []Resource{{
				Collection: "test-collection",
				Db:         "test-db",
			}},
		}},
		InheritedRoles: []InheritedRole{{
			Db:   "test-db",
			Role: "read",
		}},
		RoleName: "test-role-name",
	}

	mux.HandleFunc("/groups/1/customDBRoles/roles/test-role-name", func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"actions": []interface{}{map[string]interface{}{
				"action": "CREATE_INDEX",
				"resources": []interface{}{map[string]interface{}{
					"collection": "test-collection",
					"db":         "test-db",
				}},
			}},
			"inheritedRoles": []interface{}{map[string]interface{}{
				"db":   "test-db",
				"role": "read",
			}},
			"roleName": "test-role-name",
		}

		jsonBlob := `
		{
			"actions": [
				{
					"action": "CREATE_INDEX",
					"resources": [
						{
							"collection": "test-collection",
							"db": "test-db"
						}
					]
				}
			],
			"inheritedRoles": [
				{
					"db": "test-db",
					"role": "read"
				}
			],
			"roleName":"test-role-name"
		}
		`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Request body\n got=%#v\nwant=%#v", v, expected)
		}

		fmt.Fprint(w, jsonBlob)
	})

	customDBRole, _, err := client.CustomDBRoles.Update(ctx, "1", "test-role-name", updateRequest)
	if err != nil {
		t.Fatalf("CustomDBRoles.Update returned error: %v", err)
	}

	if roleName := customDBRole.RoleName; roleName != "test-role-name" {
		t.Errorf("expected roleName '%s', received '%s'", "test-role-name", roleName)
	}
}

func TestDatabaseUsers_DeleteCustomDBRole(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	roleName := "test-role-name"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/customDBRoles/roles/%s", groupID, roleName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.CustomDBRoles.Delete(ctx, groupID, roleName)
	if err != nil {
		t.Fatalf("CustomDBRole.Delete returned error: %v", err)
	}
}

func TestCustomDBRoles_DeleteInheritedRole(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &CustomDBRole{
		Actions: []Action{{
			Action: "CREATE_INDEX",
			Resources: []Resource{{
				Collection: "test-collection",
				Db:         "test-db",
			}},
		}},
		InheritedRoles: []InheritedRole{},
		RoleName:       "test-role-name",
	}

	mux.HandleFunc("/groups/1/customDBRoles/roles/test-role-name", func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"actions": []interface{}{map[string]interface{}{
				"action": "CREATE_INDEX",
				"resources": []interface{}{map[string]interface{}{
					"collection": "test-collection",
					"db":         "test-db",
				}},
			}},
			"inheritedRoles": []interface{}{},
			"roleName":       "test-role-name",
		}

		jsonBlob := `
		{
			"actions": [
				{
					"action": "CREATE_INDEX",
					"resources": [
						{
							"collection": "test-collection",
							"db": "test-db"
						}
					]
				}
			],
			"inheritedRoles": [],
			"roleName":"test-role-name"
		}
		`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Request body\n got=%#v\nwant=%#v", v, expected)
		}

		fmt.Fprint(w, jsonBlob)
	})

	customDBRole, _, err := client.CustomDBRoles.Update(ctx, "1", "test-role-name", updateRequest)
	if err != nil {
		t.Fatalf("CustomDBRoles.Update returned error: %v", err)
	}

	if roleName := customDBRole.RoleName; roleName != "test-role-name" {
		t.Errorf("expected roleName '%s', received '%s'", "test-role-name", roleName)
	}
}
