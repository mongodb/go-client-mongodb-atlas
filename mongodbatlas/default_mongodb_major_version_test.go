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
)

func TestDefaultMongoDBMajorVersionServiceOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/private/unauth/nds/defaultMongoDBMajorVersion", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, "4.2")
	})

	result, _, err := client.DefaultMongoDBMajorVersion.Get(ctx)
	if err != nil {
		t.Fatalf("DefaultMongoDBMajorVersion.Get returned error: %v", err)
	}
	const expected = "4.2"
	if result != expected {
		t.Errorf("Expected %s, Got %s", expected, result)
	}
}
