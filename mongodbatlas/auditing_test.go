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
	"testing"

	"github.com/go-test/deep"
)

func TestConfigureAuditing(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6d2065c687d9d64ae7acdg41"

	mux.HandleFunc(fmt.Sprintf("/"+auditingsPath, groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"auditAuthorizationSuccess": false,
			"auditFilter":               "{\n  \"atype\": \"authenticate\",\n  \"param\": {\n    \"user\": \"auditAdmin\",\n    \"db\": \"admin\",\n    \"mechanism\": \"SCRAM-SHA-1\"\n  }\n}",
			"enabled":                   true,
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
			"auditAuthorizationSuccess": false,
			"auditFilter": "{\n  \"atype\": \"authenticate\",\n  \"param\": {\n    \"user\": \"auditAdmin\",\n    \"db\": \"admin\",\n    \"mechanism\": \"SCRAM-SHA-1\"\n  }\n}",
			"configurationType": "FILTER_JSON",
			"enabled": true
		}`)
	})

	auditingRequest := &Auditing{
		AuditAuthorizationSuccess: pointer(false),
		AuditFilter:               "{\n  \"atype\": \"authenticate\",\n  \"param\": {\n    \"user\": \"auditAdmin\",\n    \"db\": \"admin\",\n    \"mechanism\": \"SCRAM-SHA-1\"\n  }\n}",
		Enabled:                   pointer(true),
	}

	auditingResponse, _, err := client.Auditing.Configure(ctx, groupID, auditingRequest)
	if err != nil {
		t.Fatalf("Auditing.Configure returned error: %v", err)
	}

	expected := &Auditing{
		AuditAuthorizationSuccess: pointer(false),
		AuditFilter:               "{\n  \"atype\": \"authenticate\",\n  \"param\": {\n    \"user\": \"auditAdmin\",\n    \"db\": \"admin\",\n    \"mechanism\": \"SCRAM-SHA-1\"\n  }\n}",
		ConfigurationType:         "FILTER_JSON",
		Enabled:                   pointer(true),
	}

	if diff := deep.Equal(auditingResponse, expected); diff != nil {
		t.Errorf("Request body\n got=%#v\nwant=%#v \n\ndiff=%#v", auditingResponse, expected, diff)
	}
}

func TestAuditing_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6d2065c687d9d64ae7acdg41"

	mux.HandleFunc(fmt.Sprintf("/"+auditingsPath, groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"auditAuthorizationSuccess": false,
			"auditFilter": "{\n  \"atype\": \"authenticate\",\n  \"param\": {\n    \"user\": \"auditAdmin\",\n    \"db\": \"admin\",\n    \"mechanism\": \"SCRAM-SHA-1\"\n  }\n}",
			"configurationType": "FILTER_JSON",
			"enabled": true
		}`)
	})

	auditing, _, err := client.Auditing.Get(ctx, groupID)
	if err != nil {
		t.Fatalf("Auditing.Get returned error: %v", err)
	}

	expected := &Auditing{
		AuditAuthorizationSuccess: pointer(false),
		AuditFilter:               "{\n  \"atype\": \"authenticate\",\n  \"param\": {\n    \"user\": \"auditAdmin\",\n    \"db\": \"admin\",\n    \"mechanism\": \"SCRAM-SHA-1\"\n  }\n}",
		ConfigurationType:         "FILTER_JSON",
		Enabled:                   pointer(true),
	}

	if diff := deep.Equal(auditing, expected); diff != nil {
		t.Errorf("Request body\n got=%#v\nwant=%#v \n\ndiff=%#v", auditing, expected, diff)
	}
}
