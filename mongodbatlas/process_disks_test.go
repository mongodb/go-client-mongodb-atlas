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

func TestProcessDisksService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/disks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
		   "links":[
			  {
				 "href":"https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/disks?pageNum=1&itemsPerPage=100",
				 "rel":"self"
			  }
		   ],
		   "results":[
			  {
				 "links":[
					{
					   "href":"https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/disks/xvdb",
					   "rel":"self"
					}
				 ],
				 "partitionName":"xvdb"
			  }
		   ],
		   "totalCount":1
		}`)
	})

	disks, _, err := client.ProcessDisks.List(ctx, "12345678", "shard-00-00.mongodb.net", 27017, nil)
	if err != nil {
		t.Fatalf("ProcessDisks.List returned error: %v", err)
	}

	expected := &ProcessDisksResponse{
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/disks?pageNum=1&itemsPerPage=100",
			},
		},
		Results: []*ProcessDisk{
			{
				Links: []*Link{
					{
						Rel:  "self",
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/disks/xvdb",
					},
				},
				PartitionName: "xvdb",
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(disks, expected); diff != nil {
		t.Error(diff)
	}
}
