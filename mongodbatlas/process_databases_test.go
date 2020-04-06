package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestProcessDatabasesService_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/groups/12345678/processes/shard-00-00.mongodb.net:27017/databases", func(w http.ResponseWriter, r *http.Request) {
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
