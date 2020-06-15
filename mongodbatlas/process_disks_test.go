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

	mux.HandleFunc("/groups/12345678/processes/shard-00-00.mongodb.net:27017/disks", func(w http.ResponseWriter, r *http.Request) {
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
