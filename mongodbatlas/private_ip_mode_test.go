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
	"github.com/openlyinc/pointy"
)

func TestPrivateIPMode_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6d2065c687d9d64ae7acdg41"

	mux.HandleFunc(fmt.Sprintf("/"+privateIPModePath, groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"enabled": true
		}`)
	})

	privateIPMode, _, err := client.PrivateIPMode.Get(ctx, groupID)
	if err != nil {
		t.Fatalf("PrivateIPMode.Get returned error: %v", err)
	}

	expected := &PrivateIPMode{
		Enabled: pointy.Bool(true),
	}

	if diff := deep.Equal(privateIPMode, expected); diff != nil {
		t.Error(diff)
	}
}

func TestPrivateIPMode_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	updateRequest := &PrivateIPMode{
		Enabled: pointy.Bool(true),
	}

	mux.HandleFunc(fmt.Sprintf("/"+privateIPModePath, groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"enabled": true,
		}

		jsonBlob := `
		{
			"enabled":  true
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

	privateIPMode, _, err := client.PrivateIPMode.Update(ctx, groupID, updateRequest)
	if err != nil {
		t.Fatalf("PrivateIPMode.Update returned error: %v", err)
	}

	if enabled := pointy.BoolValue(privateIPMode.Enabled, false); !enabled {
		t.Errorf("expected privateIPMode '%t', received '%t'", true, enabled)
	}
}
