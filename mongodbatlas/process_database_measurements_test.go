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

func TestProcessDatabaseMeasurements_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groups := "12345678"
	host := "shard-00-00.mongodb.net"
	port := 27017
	database := "database"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/processes/%s:%d/databases/%s/measurements", groups, host, port, database), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				  "end" : "2017-08-22T20:31:14Z",
				  "granularity" : "PT1M",
				  "groupId" : "12345678",
				  "hostId" : "shard-00-00.mongodb.net:27017",
				  "links" : [ {
					"href" : "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/dataabses/database/measurements?granularity=PT1M&period=PT1M",
					"rel" : "self"
				  }, {
					"href" : "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017",
					"rel" : "http://mms.mongodb.com/host"
				  } ],
				  "measurements" : [ {
					"dataPoints" : [ {
					  "timestamp" : "2017-08-22T20:31:12Z",
					  "value" : null
					}, {
					  "timestamp" : "2017-08-22T20:31:14Z",
					  "value" : null
					} ],
					"name" : "DISK_PARTITION_IOPS_READ",
					"units" : "SCALAR_PER_SECOND"
				  }],
                  "databaseName":"database",
				  "processId" : "shard-00-00.mongodb.net:27017",
				  "start" : "2017-08-22T20:30:45Z"
				}`)
	})

	opts := &ProcessMeasurementListOptions{
		Granularity: "PT1M",
		Period:      "PT1M",
	}

	measurements, _, err := client.ProcessDatabaseMeasurements.List(ctx, groups, host, port, database, opts)
	if err != nil {
		t.Fatalf("ProcessDatabaseMeasurements.List returned error: %v", err)
	}

	expected := &ProcessDatabaseMeasurements{
		ProcessMeasurements: &ProcessMeasurements{
			End:         "2017-08-22T20:31:14Z",
			Granularity: "PT1M",
			GroupID:     "12345678",
			HostID:      "shard-00-00.mongodb.net:27017",
			Links: []*Link{
				{
					Rel:  "self",
					Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/dataabses/database/measurements?granularity=PT1M&period=PT1M",
				},
				{
					Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017",
					Rel:  "http://mms.mongodb.com/host",
				},
			},
			Measurements: []*Measurements{
				{
					DataPoints: []*DataPoints{
						{
							Timestamp: "2017-08-22T20:31:12Z",
							Value:     nil,
						},
						{
							Timestamp: "2017-08-22T20:31:14Z",
							Value:     nil,
						},
					},
					Name:  "DISK_PARTITION_IOPS_READ",
					Units: "SCALAR_PER_SECOND",
				},
			},
			ProcessID: "shard-00-00.mongodb.net:27017",
			Start:     "2017-08-22T20:30:45Z",
		},
		DatabaseName: "database",
	}

	if diff := deep.Equal(measurements, expected); diff != nil {
		t.Error(diff)
	}
}
