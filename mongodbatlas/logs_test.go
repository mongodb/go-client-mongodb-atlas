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
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestLogs_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	cluster := "test-username"
	log := "log"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/logs/%s", groupID, cluster, log), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, "test")
	})

	buf := new(bytes.Buffer)
	_, err := client.Logs.Get(ctx, groupID, cluster, log, buf, nil)

	if buf.String() != "test" {
		t.Fatalf("Logs.Get returned error: %v", err)
	}
}
