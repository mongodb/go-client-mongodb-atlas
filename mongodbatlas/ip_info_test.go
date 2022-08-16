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
)

func TestIPInfoOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/private/ipinfo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				"currentIpv4Address": "127.0.0.1"
		}`)
	})

	result, _, err := client.IPInfo.Get(ctx)
	if err != nil {
		t.Fatalf("IPInfo.Get returned error: %v", err)
	}
	expected := &IPInfo{CurrentIPv4Address: "127.0.0.1"}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}
