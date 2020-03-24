package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestProcesses_ListProcesses(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/1/processes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
			{
				"results": [
					{
						"created": "2017-08-22T15:14:06Z",
						"groupId": "1",
						"hostname": "shard-00-00.mongodb.net",
						"id": "shard-00-00.mongodb.net:27017",
						"lastPing": "2017-08-22T21:36:23Z",
						"links": [
							{
								"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/1/processes/shard-00-00.mongodb.net:27017",
								"rel": "self"
							}
						],
						"port": 27017,
						"replicaSetName": "replica-set-0",
						"shardName": "shard-0",
						"typeName": "REPLICA_PRIMARY",
						"version": "3.6.7"
					},
					{
						"created": "2017-08-22T15:14:06Z",
						"groupId": "1",
						"hostname": "shard-00-00.mongodb.net",
						"id": "shard-00-00.mongodb.net:27017",
						"lastPing": "2017-08-22T21:36:23Z",
						"links": [
							{
								"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/1/processes/shard-00-00.mongodb.net:27017",
								"rel": "self"
							}
						],
						"port": 27017,
						"replicaSetName": "replica-set-0",
						"shardName": "shard-0",
						"typeName": "REPLICA_PRIMARY",
						"version": "3.6.7"
					}
				],
				"totalCount": 2
			}
		`)
	})

	peers, _, err := client.Processes.List(ctx, "1", nil)
	if err != nil {
		t.Fatalf("Processes.List returned error: %v", err)
	}

	process := Process{
		Created:  "2017-08-22T15:14:06Z",
		GroupID:  "1",
		Hostname: "shard-00-00.mongodb.net",
		ID:       "shard-00-00.mongodb.net:27017",
		LastPing: "2017-08-22T21:36:23Z",
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/1/processes/shard-00-00.mongodb.net:27017",
				Rel:  "self",
			},
		},
		Port:           27017,
		ShardName:      "shard-0",
		ReplicaSetName: "replica-set-0",
		TypeName:       "REPLICA_PRIMARY",
		Version:        "3.6.7",
	}

	expected := []*Process{&process, &process}
	if diff := deep.Equal(peers, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProcesses_ListProcessesMultiplePages(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/1/processes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		dr := processesResponse{
			Results: []*Process{
				{
					Created:  "2017-08-22T15:14:06Z",
					GroupID:  "1",
					Hostname: "shard-00-00.mongodb.net",
					ID:       "shard-00-00.mongodb.net:27017",
					LastPing: "2017-08-22T21:36:23Z",
					Links: []*Link{
						{
							Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/1/processes/shard-00-00.mongodb.net:27017",
							Rel:  "self",
						},
					},
					Port:           27017,
					ShardName:      "shard-0",
					ReplicaSetName: "replica-set-0",
					TypeName:       "REPLICA_PRIMARY",
					Version:        "3.6.7",
				},
				{
					Created:  "2017-08-22T15:14:06Z",
					GroupID:  "1",
					Hostname: "shard-00-00.mongodb.net",
					ID:       "shard-00-00.mongodb.net:27017",
					LastPing: "2017-08-22T21:36:23Z",
					Links: []*Link{
						{
							Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/1/processes/shard-00-00.mongodb.net:27017",
							Rel:  "self",
						},
					},
					Port:           27017,
					ShardName:      "shard-0",
					ReplicaSetName: "replica-set-0",
					TypeName:       "REPLICA_PRIMARY",
					Version:        "3.6.7",
				},
			},
			Links: []*Link{
				{Href: "http://example.com/api/atlas/v1.0/groups/1/processes?pageNum=2&itemsPerPage=2", Rel: "self"},
				{Href: "http://example.com/api/atlas/v1.0/groups/1/processes?pageNum=2&itemsPerPage=2", Rel: "previous"},
			},
		}

		b, err := json.Marshal(dr)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprint(w, string(b))
	})

	_, resp, err := client.Processes.List(ctx, "1", nil)
	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 2)
}

func TestProcesses_RetrievePageByNumber(t *testing.T) {
	setup()
	defer teardown()

	jBlob := `
		{
			"links": [
				{
					"href": "http://example.com/api/atlas/v1.0/groups/1/processes?pageNum=1&itemsPerPage=1",
					"rel": "previous"
				},
				{
					"href": "http://example.com/api/atlas/v1.0/groups/1/processes?pageNum=2&itemsPerPage=1",
					"rel": "self"
				},
				{
					"href": "http://example.com/api/atlas/v1.0/groups/1/processes?pageNum=3&itemsPerPage=1",
					"rel": "next"
				}
			],
			"results": [
				{
					"created": "2017-08-22T15:14:06Z",
					"groupId": "1",
					"hostname": "shard-00-00.mongodb.net",
					"id": "shard-00-00.mongodb.net:27017",
					"lastPing": "2017-08-22T21:36:23Z",
					"links": [
						{
							"href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/1/processes/shard-00-00.mongodb.net:27017",
							"rel": "self"
						}
					],
					"port": 27017,
					"replicaSetName": "replica-set-0",
					"shardName": "shard-0",
					"typeName": "REPLICA_PRIMARY",
					"version": "3.6.7"
				}
			],
			"totalCount": 1
		}
	`

	mux.HandleFunc("/groups/1/processes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		expectedQuery := "pageNum=2"
		if r.URL.RawQuery != expectedQuery {
			t.Fatalf("expected query string to be '%s', was '%s'", expectedQuery, r.URL.RawQuery)
		}
		fmt.Fprint(w, jBlob)
	})

	opt := &ListOptions{PageNum: 2}
	_, resp, err := client.Processes.List(ctx, "1", opt)

	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 2)
}
