// Copyright 2019 MongoDB Inc
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

func TestCustomAWSDNSOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "5e2211c17a3e5a48f5497de3"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/awsCustomDNS", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
				"enabled": true
			}`)
	})

	archive, _, err := client.CustomAWSDNS.Get(ctx, groupID)
	if err != nil {
		t.Fatalf("CustomAWSDNS.Get returned error: %v", err)
	}

	expected := &AWSCustomDNSSetting{
		Enabled: true,
	}

	if diff := deep.Equal(archive, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCustomAWSDNSOp_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "5e2211c17a3e5a48f5497de3"

	updateRequest := &AWSCustomDNSSetting{
		Enabled: true,
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/awsCustomDNS", groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"enabled": true,
		}

		var v map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			t.Fatalf("decode json: %v", err)
		}
		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("CustomAWSDNS.Update Request Body = %v", diff)
		}

		jsonBlob := `{			
			"enabled": true
		}`
		fmt.Fprint(w, jsonBlob)
	})

	result, _, err := client.CustomAWSDNS.Update(ctx, groupID, updateRequest)
	if err != nil {
		t.Fatalf("CustomAWSDNS.Update returned error: %v", err)
	}

	if !result.Enabled {
		t.Errorf("expected to be enabled\n")
	}
}
