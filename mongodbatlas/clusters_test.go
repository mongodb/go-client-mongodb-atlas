package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/openlyinc/pointy"
)

func TestClusters_ListClusters(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/groups/1/clusters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"results": [
				{
					"autoScaling": {
						"diskGBEnabled": true
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
						"privateSrv": "mongodb+srv://cluster0-pri.auylw.mongodb.net"
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
						"regionName": "US_WEST_2"
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
						"diskGBEnabled": true
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
						"privateSrv": "mongodb+srv://cluster0-pri.auylw.mongodb.net"
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
						"regionName": "US_WEST_2"
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
		AutoScaling:   &AutoScaling{DiskGBEnabled: pointy.Bool(true)},
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
		},
		DiskSizeGB:               pointy.Float64(160),
		EncryptionAtRestProvider: "AWS",
		GroupID:                  "5356823b3794de37132bb7b",
		MongoDBVersion:           "3.4.9",
		MongoURI:                 "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
		MongoURIUpdated:          "2017-10-23T21:26:17Z",
		MongoURIWithOptions:      "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
		Name:                     "AppData",
		NumShards:                pointy.Int64(1),
		Paused:                   pointy.Bool(false),
		ProviderSettings: &ProviderSettings{
			ProviderName:     "AWS",
			DiskIOPS:         pointy.Int64(1320),
			EncryptEBSVolume: pointy.Bool(false),
			InstanceSizeName: "M40",
			RegionName:       "US_WEST_2",
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
		StateName:  "IDLE",
	}

	expected := []Cluster{cluster1, cluster1}
	if diff := deep.Equal(clusters, expected); diff != nil {
		t.Error(diff)
	}
}

func TestClusters_ListClustersMultiplePages(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/groups/1/clusters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		dr := clustersResponse{
			Results: []Cluster{
				{GroupID: "1", Name: "test-one"},
				{GroupID: "1", Name: "test-two"},
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

func TestClusters_RetrievePageByNumber(t *testing.T) {
	client, mux, _, teardown := setup()
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

func TestClusters_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	groupID := "1"

	createRequest := &Cluster{
		ID:                       "1",
		AutoScaling:              &AutoScaling{DiskGBEnabled: pointy.Bool(true)},
		BackupEnabled:            pointy.Bool(true),
		BiConnector:              &BiConnector{Enabled: pointy.Bool(false), ReadPreference: "secondary"},
		ClusterType:              "REPLICASET",
		DiskSizeGB:               pointy.Float64(160),
		EncryptionAtRestProvider: "AWS",
		GroupID:                  groupID,
		MongoDBVersion:           "3.4.9",
		MongoURI:                 "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
		MongoURIUpdated:          "2017-10-23T21:26:17Z",
		MongoURIWithOptions:      "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
		Name:                     "AppData",
		NumShards:                pointy.Int64(1),
		Paused:                   pointy.Bool(false),
		ProviderSettings: &ProviderSettings{
			ProviderName:     "AWS",
			DiskIOPS:         pointy.Int64(1320),
			EncryptEBSVolume: pointy.Bool(false),
			InstanceSizeName: "M40",
			RegionName:       "US_WEST_2",
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
		StateName:  "IDLE",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters", groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"id": "1",
			"autoScaling": map[string]interface{}{
				"diskGBEnabled": true,
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
                "diskGBEnabled": true
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
                "regionName": "US_WEST_2"
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

func TestClusters_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	groupID := "1"
	clusterName := "AppData"

	updateRequest := &Cluster{
		ID:                       "1",
		AutoScaling:              &AutoScaling{DiskGBEnabled: pointy.Bool(true)},
		BackupEnabled:            pointy.Bool(true),
		BiConnector:              &BiConnector{Enabled: pointy.Bool(false), ReadPreference: "secondary"},
		ClusterType:              "REPLICASET",
		DiskSizeGB:               pointy.Float64(160),
		EncryptionAtRestProvider: "AWS",
		GroupID:                  groupID,
		MongoDBVersion:           "3.4.9",
		MongoURI:                 "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
		MongoURIUpdated:          "2017-10-23T21:26:17Z",
		MongoURIWithOptions:      "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
		Name:                     clusterName,
		NumShards:                pointy.Int64(1),
		Paused:                   pointy.Bool(false),
		ProviderSettings: &ProviderSettings{
			ProviderName:     "AWS",
			DiskIOPS:         pointy.Int64(1320),
			EncryptEBSVolume: pointy.Bool(false),
			InstanceSizeName: "M40",
			RegionName:       "US_WEST_2",
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
		StateName:  "IDLE",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"id": "1",
			"autoScaling": map[string]interface{}{
				"diskGBEnabled": true,
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
                "diskGBEnabled": true
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
                "regionName": "US_WEST_2"
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

func TestClusters_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	groupID := "1"
	name := "test-name"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s", groupID, name), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Clusters.Delete(ctx, groupID, name)
	if err != nil {
		t.Fatalf("Cluster.Delete returned error: %v", err)
	}
}

func TestClusters_UpdateProcessArgs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	groupID := "1"
	clusterName := "AppData"
	tlsProtocol := "TLS1_2"

	updateRequest := &ProcessArgs{
		FailIndexKeyTooLong:              pointy.Bool(false),
		JavascriptEnabled:                pointy.Bool(false),
		MinimumEnabledTLSProtocol:        tlsProtocol,
		NoTableScan:                      pointy.Bool(true),
		OplogSizeMB:                      pointy.Int64(2000),
		SampleSizeBIConnector:            pointy.Int64(5000),
		SampleRefreshIntervalBIConnector: pointy.Int64(300),
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s/processArgs", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"failIndexKeyTooLong":              false,
			"javascriptEnabled":                false,
			"minimumEnabledTlsProtocol":        tlsProtocol,
			"noTableScan":                      true,
			"oplogSizeMB":                      float64(2000),
			"sampleSizeBIConnector":            float64(5000),
			"sampleRefreshIntervalBIConnector": float64(300),
		}

		jsonBlob := `
		{
			"failIndexKeyTooLong": false,
			"javascriptEnabled": false,
			"minimumEnabledTlsProtocol": "TLS1_2",
			"noTableScan": true,
			"oplogSizeMB": 2000,
			"sampleSizeBIConnector": 5000,
			"sampleRefreshIntervalBIConnector": 300
		}
		`

		var v map[string]interface{}
		d := json.NewDecoder(r.Body)

		err := d.Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, jsonBlob)
	})

	processArgs, _, err := client.Clusters.UpdateProcessArgs(ctx, groupID, clusterName, updateRequest)
	if err != nil {
		t.Fatalf("Clusters.UpdateProcessArgs returned error: %v", err)
	}

	if tls := processArgs.MinimumEnabledTLSProtocol; tls != tlsProtocol {
		t.Errorf("expected tlsProtocol '%s', received '%s'", tlsProtocol, tls)
	}

	if jsEnabled := processArgs.JavascriptEnabled; pointy.BoolValue(jsEnabled, false) != false {
		t.Errorf("expected javascriptEnabled '%t', received '%t'", pointy.BoolValue(jsEnabled, false), false)
	}

}

func TestClusters_GetProcessArgs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	groupID := "1"
	clusterName := "test-cluster"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s/processArgs", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"failIndexKeyTooLong": false,
			"javascriptEnabled": false,
			"minimumEnabledTlsProtocol": "TLS1_2",
			"noTableScan": true,
			"oplogSizeMB": 2000,
			"sampleSizeBIConnector": 5000,
			"sampleRefreshIntervalBIConnector": 300
		}`)
	})

	processArgs, _, err := client.Clusters.GetProcessArgs(ctx, groupID, clusterName)
	if err != nil {
		t.Fatalf("Clusters.GetProcessArgs returned error: %v", err)
	}

	expected := &ProcessArgs{
		FailIndexKeyTooLong:              pointy.Bool(false),
		JavascriptEnabled:                pointy.Bool(false),
		MinimumEnabledTLSProtocol:        "TLS1_2",
		NoTableScan:                      pointy.Bool(true),
		OplogSizeMB:                      pointy.Int64(2000),
		SampleSizeBIConnector:            pointy.Int64(5000),
		SampleRefreshIntervalBIConnector: pointy.Int64(300),
	}

	if diff := deep.Equal(processArgs, expected); diff != nil {
		t.Error(diff)
	}
}

func TestClusters_checkClusterNameParam(t *testing.T) {
	if err := checkClusterNameParam(""); err == nil {
		t.Fatalf("checkClusterNameParam didn't return error")
	}
}

func TestClusters_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	groupID := "1"
	clusterName := "appData"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{	
			"id":"1",
			"autoScaling": {
                "diskGBEnabled": true
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
				"privateSrv": "mongodb+srv://cluster0-pri.auylw.mongodb.net"
			},
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
                "regionName": "US_WEST_2"
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
		}`)
	})

	cluster, _, err := client.Clusters.Get(ctx, groupID, clusterName)
	if err != nil {
		t.Fatalf("Clusters.Get returned error: %v", err)
	}

	expected := &Cluster{
		ID:            "1",
		AutoScaling:   &AutoScaling{DiskGBEnabled: pointy.Bool(true)},
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
		},
		DiskSizeGB:               pointy.Float64(160),
		EncryptionAtRestProvider: "AWS",
		GroupID:                  groupID,
		MongoDBVersion:           "3.4.9",
		MongoURI:                 "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
		MongoURIUpdated:          "2017-10-23T21:26:17Z",
		MongoURIWithOptions:      "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
		Name:                     "AppData",
		NumShards:                pointy.Int64(1),
		Paused:                   pointy.Bool(false),
		PitEnabled:               pointy.Bool(false),
		ProviderSettings: &ProviderSettings{
			ProviderName:     "AWS",
			DiskIOPS:         pointy.Int64(1320),
			EncryptEBSVolume: pointy.Bool(false),
			InstanceSizeName: "M40",
			RegionName:       "US_WEST_2",
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
		StateName:  "IDLE",
	}

	if diff := deep.Equal(cluster, expected); diff != nil {
		t.Error(diff)
	}
}
