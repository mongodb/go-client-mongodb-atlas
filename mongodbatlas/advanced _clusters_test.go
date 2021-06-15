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
	"github.com/openlyinc/pointy"
)

func TestAdvancedClusters_ListClusters(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/groups/1/clusters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"results": [
				{
					"autoScaling": {
						"diskGBEnabled": true,
						"compute": {
						  "enabled": true,
						  "scaleDownEnabled": true
						}
					},
					"backupEnabled": true,
					"biConnector": {
						"enabled": false,
						"readPreference": "secondary"
					},
					"clusterType": "REPLICASET",
					"connectionStrings": {
						"standard": "mongodb://cluster0-shard-00-00-auylw.mongodb.net:27017,cluster0-shard-00-01-auylw.mongodb.net:27017,cluster0-shard-00-02-auylw.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0",
						"standardSrv": "mongodb+srv://cluster0-auylw.mongodb.net",
						"awsPrivateLink": {
							"vpce-0d00c26273372c6ef": "mongodb://pl-0-us-east-1-auylw.mongodb.net:1024,pl-0-us-east-1-auylw.mongodb.net:1025,pl-0-us-east-1-auylw.mongodb.net:1026/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0-shard-0"
						},
						"awsPrivateLinkSrv": {
							"vpce-0d00c26273372c6ef": "mongodb+srv://cluster0-pl-0-auylw.mongodb.net"
						},
						"private": "mongodb://cluster0-shard-00-00-pri.auylw.mongodb.net:27017,cluster0-shard-00-01-pri.auylw.mongodb.net:27017,cluster0-shard-00-02-pri.auylw.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0",
						"privateSrv": "mongodb+srv://cluster0-pri.auylw.mongodb.net",
						"privateEndpoint": [
						 {
						  "connectionString": "mongodb://pl-0-us-east-1-auylw.mongodb.net:1024,pl-0-us-east-1-auylw.mongodb.net:1025,pl-0-us-east-1-auylw.mongodb.net:1026/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0-shard-0",
						  "endpoints": [
						   {
							"endpointId": "vpce-0d00c26273372c6ef",
							"providerName": "AWS",
							"region": "US_EAST_1"
						   }
						  ],
						  "srvConnectionString": "mongodb+srv://cluster0-pl-0-auylw.mongodb.net",
						  "type": "MONGOD"
						 }
						]
					},
					"diskSizeGB": 160,
					"encryptionAtRestProvider": "AWS",
					"groupId": "5356823b3794de37132bb7b",
					"mongoDBVersion": "3.4.9",
					"mongoURI": "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
					"mongoURIUpdated": "2017-10-23T21:26:17Z",
					"mongoURIWithOptions": "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
					"name": "AppData",
					"numShards": 1,
					"paused": false,
					"providerSettings": {
						"providerName": "AWS",
						"diskIOPS": 1320,
						"encryptEBSVolume": false,
						"instanceSizeName": "M40",
						"regionName": "US_WEST_2",
						"autoScaling": {
							"compute": {
							  "maxInstanceSize": "M60",
							  "minInstanceSize": "M10"
							}
						}
					},
					"replicationFactor": 3,
					"replicationSpec": {
						"US_EAST_1": {
							"electableNodes": 3,
							"priority": 7,
							"readOnlyNodes": 0
						}
					},
					"srvAddress": "mongodb+srv://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
					"stateName": "IDLE"
				},
				{
					"autoScaling": {
						"diskGBEnabled": true,
						"compute": {
						  "enabled": true,
						  "scaleDownEnabled": true
						}
					},
					"backupEnabled": true,
					"biConnector": {
						"enabled": false,
						"readPreference": "secondary"
					},
					"clusterType": "REPLICASET",
					"connectionStrings": {
						"standard": "mongodb://cluster0-shard-00-00-auylw.mongodb.net:27017,cluster0-shard-00-01-auylw.mongodb.net:27017,cluster0-shard-00-02-auylw.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0",
						"standardSrv": "mongodb+srv://cluster0-auylw.mongodb.net",
						"awsPrivateLink": {
							"vpce-0d00c26273372c6ef": "mongodb://pl-0-us-east-1-auylw.mongodb.net:1024,pl-0-us-east-1-auylw.mongodb.net:1025,pl-0-us-east-1-auylw.mongodb.net:1026/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0-shard-0"
						},
						"awsPrivateLinkSrv": {
							"vpce-0d00c26273372c6ef": "mongodb+srv://cluster0-pl-0-auylw.mongodb.net"
						},
						"private": "mongodb://cluster0-shard-00-00-pri.auylw.mongodb.net:27017,cluster0-shard-00-01-pri.auylw.mongodb.net:27017,cluster0-shard-00-02-pri.auylw.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0",
						"privateSrv": "mongodb+srv://cluster0-pri.auylw.mongodb.net",
						"privateEndpoint": [
						 {
						  "connectionString": "mongodb://pl-0-us-east-1-auylw.mongodb.net:1024,pl-0-us-east-1-auylw.mongodb.net:1025,pl-0-us-east-1-auylw.mongodb.net:1026/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0-shard-0",
						  "endpoints": [
						   {
							"endpointId": "vpce-0d00c26273372c6ef",
							"providerName": "AWS",
							"region": "US_EAST_1"
						   }
						  ],
						  "srvConnectionString": "mongodb+srv://cluster0-pl-0-auylw.mongodb.net",
						  "type": "MONGOD"
						 }
						]
					},
					"diskSizeGB": 160,
					"encryptionAtRestProvider": "AWS",
					"groupId": "5356823b3794de37132bb7b",
					"mongoDBVersion": "3.4.9",
					"mongoURI": "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
					"mongoURIUpdated": "2017-10-23T21:26:17Z",
					"mongoURIWithOptions": "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
					"name": "AppData",
					"numShards": 1,
					"paused": false,
					"providerSettings": {
						"providerName": "AWS",
						"diskIOPS": 1320,
						"encryptEBSVolume": false,
						"instanceSizeName": "M40",
						"regionName": "US_WEST_2",
						"autoScaling": {
							"compute": {
							  "maxInstanceSize": "M60",
							  "minInstanceSize": "M10"
							}
						}
					},
					"replicationFactor": 3,
					"replicationSpec": {
						"US_EAST_1": {
							"electableNodes": 3,
							"priority": 7,
							"readOnlyNodes": 0
						}
					},
					"srvAddress": "mongodb+srv://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
					"stateName": "IDLE"
				}
			],
			"totalCount": 2
		}`)
	})

	clusters, _, err := client.Clusters.List(ctx, "1", nil)
	if err != nil {
		t.Fatalf("Clusters.List returned error: %v", err)
	}

	cluster1 := Cluster{
		AutoScaling: &AutoScaling{
			DiskGBEnabled: pointy.Bool(true),
			Compute: &Compute{
				Enabled:          pointy.Bool(true),
				ScaleDownEnabled: pointy.Bool(true),
			},
		},
		BaseCluster: &BaseCluster{
			BackupEnabled: pointy.Bool(true),
			BiConnector:   &BiConnector{Enabled: pointy.Bool(false), ReadPreference: "secondary"},
			ClusterType:   "REPLICASET",
			ConnectionStrings: &ConnectionStrings{
				Standard:          "mongodb://cluster0-shard-00-00-auylw.mongodb.net:27017,cluster0-shard-00-01-auylw.mongodb.net:27017,cluster0-shard-00-02-auylw.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0",
				StandardSrv:       "mongodb+srv://cluster0-auylw.mongodb.net",
				AwsPrivateLink:    map[string]string{"vpce-0d00c26273372c6ef": "mongodb://pl-0-us-east-1-auylw.mongodb.net:1024,pl-0-us-east-1-auylw.mongodb.net:1025,pl-0-us-east-1-auylw.mongodb.net:1026/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0-shard-0"},
				AwsPrivateLinkSrv: map[string]string{"vpce-0d00c26273372c6ef": "mongodb+srv://cluster0-pl-0-auylw.mongodb.net"},
				Private:           "mongodb://cluster0-shard-00-00-pri.auylw.mongodb.net:27017,cluster0-shard-00-01-pri.auylw.mongodb.net:27017,cluster0-shard-00-02-pri.auylw.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0",
				PrivateSrv:        "mongodb+srv://cluster0-pri.auylw.mongodb.net",
				PrivateEndpoint: []PrivateEndpoint{
					{
						ConnectionString:    "mongodb://pl-0-us-east-1-auylw.mongodb.net:1024,pl-0-us-east-1-auylw.mongodb.net:1025,pl-0-us-east-1-auylw.mongodb.net:1026/?ssl=true&authSource=admin&replicaSet=Cluster0-shard-0-shard-0",
						SRVConnectionString: "mongodb+srv://cluster0-pl-0-auylw.mongodb.net",
						Type:                "MONGOD",
						Endpoints: []Endpoint{
							{
								EndpointID:   "vpce-0d00c26273372c6ef",
								Region:       "US_EAST_1",
								ProviderName: "AWS",
							},
						},
					},
				},
			},
			DiskSizeGB:               pointy.Float64(160),
			EncryptionAtRestProvider: "AWS",
			GroupID:                  "5356823b3794de37132bb7b",
			MongoDBVersion:           "3.4.9",
			Name:                     "AppData",
			Paused:                   pointy.Bool(false),
			StateName:                "IDLE",
		},

		MongoURI:            "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
		MongoURIUpdated:     "2017-10-23T21:26:17Z",
		MongoURIWithOptions: "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
		NumShards:           pointy.Int64(1),
		ProviderSettings: &ProviderSettings{
			ProviderName:     "AWS",
			DiskIOPS:         pointy.Int64(1320),
			EncryptEBSVolume: pointy.Bool(false),
			InstanceSizeName: "M40",
			RegionName:       "US_WEST_2",
			AutoScaling: &AutoScaling{
				Compute: &Compute{
					MaxInstanceSize: "M60",
					MinInstanceSize: "M10",
				},
			},
		},
		ReplicationFactor: pointy.Int64(3),

		ReplicationSpec: map[string]RegionsConfig{
			"US_EAST_1": {
				ElectableNodes: pointy.Int64(3),
				Priority:       pointy.Int64(7),
				ReadOnlyNodes:  pointy.Int64(0),
			},
		},
		SrvAddress: "mongodb+srv://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
	}

	expected := []Cluster{cluster1, cluster1}
	if diff := deep.Equal(clusters, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAdvancedClusters_ListClustersMultiplePages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/groups/1/clusters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		dr := clustersResponse{
			Results: []Cluster{
				{BaseCluster: &BaseCluster{GroupID: "1", Name: "test-one"}},
				{BaseCluster: &BaseCluster{GroupID: "1", Name: "test-two"}},
			},
			Links: []*Link{
				{Href: "http://example.com/api/atlas/v1.0/groups/1/clusters?pageNum=2&itemsPerPage=2", Rel: "self"},
				{Href: "http://example.com/api/atlas/v1.0/groups/1/clusters?pageNum=2&itemsPerPage=2", Rel: "previous"},
			},
		}

		b, err := json.Marshal(dr)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprint(w, string(b))
	})

	_, resp, err := client.Clusters.List(ctx, "1", nil)
	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 2)
}

func TestAdvancedClusters_RetrievePageByNumber(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	jBlob := `
	{
		"links": [
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/clusters?pageNum=1&itemsPerPage=1",
				"rel": "previous"
			},
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/clusters?pageNum=2&itemsPerPage=1",
				"rel": "self"
			},
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/clusters?itemsPerPage=3&pageNum=2",
				"rel": "next"
			}
		],
		"results": [
			{
				"groupId": "5356823b3794de37132bb7b",
				"mongoDBVersion": "3.4.9",
				"name": "AppData"
			}
		],
		"totalCount": 3
	}`

	mux.HandleFunc("/groups/1/clusters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, jBlob)
	})

	opt := &ListOptions{PageNum: 2}
	_, resp, err := client.Clusters.List(ctx, "1", opt)

	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 2)
}

func TestAdvancedClusters_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	createRequest := &Cluster{
		BaseCluster: &BaseCluster{
			ID:                       "1",
			BackupEnabled:            pointy.Bool(true),
			BiConnector:              &BiConnector{Enabled: pointy.Bool(false), ReadPreference: "secondary"},
			ClusterType:              "REPLICASET",
			DiskSizeGB:               pointy.Float64(160),
			EncryptionAtRestProvider: "AWS",
			GroupID:                  groupID,
			MongoDBVersion:           "3.4.9",
			Name:                     "AppData",
			Paused:                   pointy.Bool(false),
			StateName:                "IDLE",
		},

		AutoScaling: &AutoScaling{DiskGBEnabled: pointy.Bool(true),
			Compute: &Compute{Enabled: pointy.Bool(true), ScaleDownEnabled: pointy.Bool(true)}},
		MongoURI:            "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
		MongoURIUpdated:     "2017-10-23T21:26:17Z",
		MongoURIWithOptions: "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
		NumShards:           pointy.Int64(1),
		ProviderSettings: &ProviderSettings{
			ProviderName:     "AWS",
			DiskIOPS:         pointy.Int64(1320),
			EncryptEBSVolume: pointy.Bool(false),
			InstanceSizeName: "M40",
			RegionName:       "US_WEST_2",
			AutoScaling:      &AutoScaling{Compute: &Compute{MinInstanceSize: "M10", MaxInstanceSize: "M60"}},
		},
		ReplicationFactor: pointy.Int64(3),

		ReplicationSpec: map[string]RegionsConfig{
			"US_EAST_1": {
				ElectableNodes: pointy.Int64(3),
				Priority:       pointy.Int64(7),
				ReadOnlyNodes:  pointy.Int64(0),
			},
		},
		SrvAddress: "mongodb+srv://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters", groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"id": "1",
			"autoScaling": map[string]interface{}{
				"diskGBEnabled": true,
				"compute": map[string]interface{}{
					"enabled":          true,
					"scaleDownEnabled": true,
				},
			},
			"backupEnabled": true,
			"biConnector": map[string]interface{}{
				"enabled":        false,
				"readPreference": "secondary",
			},
			"clusterType":              "REPLICASET",
			"diskSizeGB":               float64(160),
			"encryptionAtRestProvider": "AWS",
			"groupId":                  groupID,
			"mongoDBVersion":           "3.4.9",
			"mongoURI":                 "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
			"mongoURIUpdated":          "2017-10-23T21:26:17Z",
			"mongoURIWithOptions":      "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
			"name":                     "AppData",
			"numShards":                float64(1),
			"paused":                   false,
			"providerSettings": map[string]interface{}{
				"providerName":     "AWS",
				"diskIOPS":         float64(1320),
				"encryptEBSVolume": false,
				"instanceSizeName": "M40",
				"regionName":       "US_WEST_2",
				"autoScaling": map[string]interface{}{
					"compute": map[string]interface{}{
						"minInstanceSize": "M10",
						"maxInstanceSize": "M60",
					},
				},
			},
			"replicationFactor": float64(3),
			"replicationSpec": map[string]interface{}{
				"US_EAST_1": map[string]interface{}{
					"electableNodes": float64(3),
					"priority":       float64(7),
					"readOnlyNodes":  float64(0),
				},
			},
			"srvAddress": "mongodb+srv://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
			"stateName":  "IDLE",
		}

		jsonBlob := `
		{	
			"id":"1",
			"autoScaling": {
                "diskGBEnabled": true,
				"compute": {
				  "enabled": true,
				  "scaleDownEnabled": true
				}
            },
            "backupEnabled": true,
            "biConnector": {
                "enabled": false,
                "readPreference": "secondary"
            },
            "clusterType": "REPLICASET",
            "diskSizeGB": 160,
            "encryptionAtRestProvider": "AWS",
            "groupId": "1",
            "mongoDBVersion": "3.4.9",
            "mongoURI": "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
            "mongoURIUpdated": "2017-10-23T21:26:17Z",
            "mongoURIWithOptions": "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
            "name": "AppData",
            "numShards": 1,
			"paused": false,
			"pitEnabled": false,
            "providerSettings": {
                "providerName": "AWS",
                "diskIOPS": 1320,
                "encryptEBSVolume": false,
                "instanceSizeName": "M40",
                "regionName": "US_WEST_2",
				"autoScaling": {
					"compute": {
					  "minInstanceSize": "M10",
					  "maxInstanceSize": "M60"
					}
				}
            },
            "replicationFactor": 3,
            "replicationSpec": {
                "US_EAST_1": {
                    "electableNodes": 3,
                    "priority": 7,
                    "readOnlyNodes": 0
                }
            },
            "srvAddress": "mongodb+srv://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
            "stateName": "IDLE"
		}
		`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}
		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Clusters.Create Request Body = %v", diff)
		}

		fmt.Fprint(w, jsonBlob)
	})

	cluster, _, err := client.Clusters.Create(ctx, groupID, createRequest)
	if err != nil {
		t.Fatalf("Clusters.Create returned error: %v", err)
	}

	expectedName := "AppData"

	if clusterName := cluster.Name; clusterName != expectedName {
		t.Errorf("expected name '%s', received '%s'", expectedName, clusterName)
	}

	if id := cluster.GroupID; id != groupID {
		t.Errorf("expected groupId '%s', received '%s'", groupID, id)
	}

	if pitEnabled := cluster.PitEnabled; *pitEnabled {
		t.Errorf("expected pitEnabled 'false', received '%t'", *pitEnabled)
	}
}

func TestAdvancedClusters_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	clusterName := "AppData"

	updateRequest := &Cluster{
		BaseCluster: &BaseCluster{
			ID:                       "1",
			BackupEnabled:            pointy.Bool(true),
			BiConnector:              &BiConnector{Enabled: pointy.Bool(false), ReadPreference: "secondary"},
			ClusterType:              "REPLICASET",
			DiskSizeGB:               pointy.Float64(160),
			EncryptionAtRestProvider: "AWS",
			GroupID:                  groupID,
			MongoDBVersion:           "3.4.9",
			Name:                     clusterName,
			Paused:                   pointy.Bool(false),
			StateName:                "IDLE",
		},
		AutoScaling: &AutoScaling{DiskGBEnabled: pointy.Bool(true),
			Compute: &Compute{Enabled: pointy.Bool(true), ScaleDownEnabled: pointy.Bool(true)}},
		MongoURI:            "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
		MongoURIUpdated:     "2017-10-23T21:26:17Z",
		MongoURIWithOptions: "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
		NumShards:           pointy.Int64(1),
		ProviderSettings: &ProviderSettings{
			ProviderName:     "AWS",
			DiskIOPS:         pointy.Int64(1320),
			EncryptEBSVolume: pointy.Bool(false),
			InstanceSizeName: "M40",
			RegionName:       "US_WEST_2",
			AutoScaling:      &AutoScaling{Compute: &Compute{MinInstanceSize: "M20", MaxInstanceSize: "M80"}},
		},
		ReplicationFactor: pointy.Int64(3),

		ReplicationSpec: map[string]RegionsConfig{
			"US_EAST_1": {
				ElectableNodes: pointy.Int64(3),
				Priority:       pointy.Int64(7),
				ReadOnlyNodes:  pointy.Int64(0),
			},
		},
		SrvAddress: "mongodb+srv://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"id": "1",
			"autoScaling": map[string]interface{}{
				"diskGBEnabled": true,
				"compute": map[string]interface{}{
					"enabled":          true,
					"scaleDownEnabled": true,
				},
			},
			"backupEnabled": true,
			"biConnector": map[string]interface{}{
				"enabled":        false,
				"readPreference": "secondary",
			},
			"clusterType":              "REPLICASET",
			"diskSizeGB":               float64(160),
			"encryptionAtRestProvider": "AWS",
			"groupId":                  groupID,
			"mongoDBVersion":           "3.4.9",
			"mongoURI":                 "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
			"mongoURIUpdated":          "2017-10-23T21:26:17Z",
			"mongoURIWithOptions":      "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
			"name":                     "AppData",
			"numShards":                float64(1),
			"paused":                   false,
			"providerSettings": map[string]interface{}{
				"providerName":     "AWS",
				"diskIOPS":         float64(1320),
				"encryptEBSVolume": false,
				"instanceSizeName": "M40",
				"regionName":       "US_WEST_2",
				"autoScaling": map[string]interface{}{
					"compute": map[string]interface{}{
						"minInstanceSize": "M20",
						"maxInstanceSize": "M80",
					},
				},
			},
			"replicationFactor": float64(3),
			"replicationSpec": map[string]interface{}{
				"US_EAST_1": map[string]interface{}{
					"electableNodes": float64(3),
					"priority":       float64(7),
					"readOnlyNodes":  float64(0),
				},
			},
			"srvAddress": "mongodb+srv://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
			"stateName":  "IDLE",
		}

		jsonBlob := `
		{
			"autoScaling": {
                "diskGBEnabled": true,
				"compute": {
				  "enabled": true,
				  "scaleDownEnabled": true
				}
            },
            "backupEnabled": true,
            "biConnector": {
                "enabled": false,
                "readPreference": "secondary"
            },
            "clusterType": "REPLICASET",
            "diskSizeGB": 160,
            "encryptionAtRestProvider": "AWS",
            "groupId": "1",
            "mongoDBVersion": "3.4.9",
            "mongoURI": "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
            "mongoURIUpdated": "2017-10-23T21:26:17Z",
            "mongoURIWithOptions": "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
            "name": "AppData",
            "numShards": 1,
			"paused": false,
			"pitEnabled": false,
            "providerSettings": {
                "providerName": "AWS",
                "diskIOPS": 1320,
                "encryptEBSVolume": false,
                "instanceSizeName": "M40",
                "regionName": "US_WEST_2",
				"autoScaling": {
					"compute": {
					  "minInstanceSize": "M10",
					  "maxInstanceSize": "M60"
					}
				}
            },
            "replicationFactor": 3,
            "replicationSpec": {
                "US_EAST_1": {
                    "electableNodes": 3,
                    "priority": 7,
                    "readOnlyNodes": 0
                }
            },
            "srvAddress": "mongodb+srv://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
            "stateName": "IDLE"
		}
		`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Clusters.Update Request Body = %v", diff)
		}

		fmt.Fprint(w, jsonBlob)
	})

	dbUser, _, err := client.Clusters.Update(ctx, groupID, clusterName, updateRequest)
	if err != nil {
		t.Fatalf("Clusters.Update returned error: %v", err)
	}

	if name := dbUser.Name; name != clusterName {
		t.Errorf("expected name '%s', received '%s'", clusterName, name)
	}

	if id := dbUser.GroupID; id != groupID {
		t.Errorf("expected groupId '%s', received '%s'", groupID, id)
	}
}
