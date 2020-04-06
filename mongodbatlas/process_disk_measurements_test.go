package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestDiskMeasurements_List(t *testing.T) {
	setup()
	defer teardown()

	groups := "12345678"
	host := "shard-00-00.mongodb.net"
	port := 27017
	disk := "disk"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/processes/%s:%d/disks/%s/measurements", groups, host, port, disk), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				  "end" : "2017-08-22T20:31:14Z",
				  "granularity" : "PT1M",
				  "groupId" : "12345678",
				  "hostId" : "shard-00-00.mongodb.net:27017",
				  "links" : [ {
					"href" : "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/measurements?granularity=PT1M&period=PT1M",
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
                  "partitionName":"xvdb",
				  "processId" : "shard-00-00.mongodb.net:27017",
				  "start" : "2017-08-22T20:30:45Z"
				}`)
	})

	opts := &ProcessMeasurementListOptions{
		Granularity: "PT1M",
		Period:      "PT1M",
	}

	measurements, _, err := client.DiskMeasurements.Get(ctx, groups, host, port, disk, opts)
	if err != nil {
		t.Fatalf("Teams.Get returned error: %v", err)
	}

	expected := &ProcessDiskMeasurements{
		ProcessMeasurements: &ProcessMeasurements{
			End:         "2017-08-22T20:31:14Z",
			Granularity: "PT1M",
			GroupID:     "12345678",
			HostID:      "shard-00-00.mongodb.net:27017",
			Links: []*Link{
				{
					Rel:  "self",
					Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/measurements?granularity=PT1M&period=PT1M",
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
		PartitionName: "xvdb",
	}

	if diff := deep.Equal(measurements, expected); diff != nil {
		t.Error(diff)
	}
}
