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
)

func TestProcesses_ListProcesses(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups/1/processes", func(w http.ResponseWriter, r *http.Request) {
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
						"version": "3.6.7",
						"userAlias": "zuul"
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
						"version": "3.6.7",
						"userAlias": "zuul"
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
		UserAlias:      "zuul",
	}

	expected := []*Process{&process, &process}
	if diff := deep.Equal(peers, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProcesses_ListProcessesMultiplePages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups/1/processes", func(w http.ResponseWriter, r *http.Request) {
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
					UserAlias:      "zuul",
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
					UserAlias:      "zuul",
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
	client, mux, teardown := setup()
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
					"version": "3.6.7",
					"userAlias": "zuul"
				}
			],
			"totalCount": 1
		}
	`

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/processes", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		expectedQuery := "pageNum=2"
		if r.URL.RawQuery != expectedQuery {
			t.Fatalf("expected query string to be '%s', was '%s'", expectedQuery, r.URL.RawQuery)
		}
		fmt.Fprint(w, jBlob)
	})

	opt := &ProcessesListOptions{}
	opt.PageNum = 2
	_, resp, err := client.Processes.List(ctx, "1", opt)

	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 2)
}

func TestProcessesServiceOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	const hostname = "atlas-abcdef-shard-00-00.nta8e.mongodb.net"
	const port = 27017
	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/processes/%s:%d", groupID, hostname, port)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
		  "created" : "2020-08-25T18:44:13Z",
		  "groupId" : "1",
		  "hostname" : "atlas-abcdef-shard-00-00.nta8e.mongodb.net",
		  "id" : "atlas-abcdef-shard-00-00.nta8e.mongodb.net:27017",
		  "lastPing" : "2020-09-01T18:40:06Z",
		  "port" : 27017,
		  "replicaSetName" : "atlas-abcdef-shard-0",
		  "typeName" : "REPLICA_PRIMARY",
		  "userAlias" : "testcluster-shard-00-00.nta8e.mongodb.net",
		  "version" : "4.4.0"
		}`)
	})

	cluster, _, err := client.Processes.Get(ctx, groupID, hostname, port)
	if err != nil {
		t.Fatalf("Processes.Get returned error: %v", err)
	}

	expected := &Process{
		ID:             "atlas-abcdef-shard-00-00.nta8e.mongodb.net:27017",
		GroupID:        groupID,
		Hostname:       hostname,
		TypeName:       "REPLICA_PRIMARY",
		Created:        "2020-08-25T18:44:13Z",
		LastPing:       "2020-09-01T18:40:06Z",
		Port:           port,
		ReplicaSetName: "atlas-abcdef-shard-0",
		UserAlias:      "testcluster-shard-00-00.nta8e.mongodb.net",
		Version:        "4.4.0",
	}

	if diff := deep.Equal(cluster, expected); diff != nil {
		t.Error(diff)
	}
}
