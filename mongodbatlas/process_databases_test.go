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

func TestProcessDatabasesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/databases", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
		   "links":[
			  {
				 "href":"https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/databases?pageNum=1&itemsPerPage=100",
				 "rel":"self"
			  }
		   ],
		   "results":[
			  {
				 "links":[
					{
					   "href":"https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/databases/local",
					   "rel":"self"
					}
				 ],
				 "databaseName":"local"
			  }
		   ],
		   "totalCount":1
		}`)
	})

	dbs, _, err := client.ProcessDatabases.List(ctx, "12345678", "shard-00-00.mongodb.net", 27017, nil)
	if err != nil {
		t.Fatalf("ProcessDatabases.List returned error: %v", err)
	}

	expected := &ProcessDatabasesResponse{
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/databases?pageNum=1&itemsPerPage=100",
			},
		},
		Results: []*ProcessDatabase{
			{
				Links: []*Link{
					{
						Rel:  "self",
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/databases/local",
					},
				},
				DatabaseName: "local",
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(dbs, expected); diff != nil {
		t.Error(diff)
	}
}
