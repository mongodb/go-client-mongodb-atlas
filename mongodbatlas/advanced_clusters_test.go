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
	"github.com/openlyinc/pointy"
)

const (
	groupID     = "1"
	clusterName = "globalCluster"
)

func TestAdvancedClusters_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.5/groups/%s/clusters", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				  "links": [
					{
					  "href": "https://cloud.mongodb.com/api/atlas/v1.5/groups/1/clusters?pageNum=1&itemsPerPage=100",
					  "rel": "self"
					}
				  ],
				  "results": [
					{
					  "backupEnabled": true,
					  "biConnector": {
						"enabled": false,
						"readPreference": "secondary"
					  },
					  "clusterType": "REPLICASET",
					  "createDate": "2021-03-02T22:25:18Z",
					  "diskSizeGB": 10.0,
					  "encryptionAtRestProvider": "NONE",
					  "groupId": "1",
					  "id": "1",
					  "mongoDBMajorVersion": "4.4",
					  "mongoDBVersion": "4.4.4",
					  "name": "basicReplicaSet",
					  "paused": false,
					  "pitEnabled": true,
					  "replicationSpecs": [
						{
						  "id": "1",
						  "numShards": 1,
						  "regionConfigs": [
							{
                              "analyticsAutoScaling": {
								"compute": {
									"enabled": true,
									"maxInstanceSize": "M20",
									"minInstanceSize": "M10",
									"scaleDownEnabled": false
								}, 
								"diskGB": {
									"enabled": true
								}
							  },
							  "analyticsSpecs": {
								"instanceSize": "M10",
								"diskIOPS": 100,
								"ebsVolumeType": "STANDARD",
								"nodeCount": 0
							  },
							  "electableSpecs": {
								"instanceSize": "M10",
								"diskIOPS": 100,
								"ebsVolumeType": "STANDARD",
								"nodeCount": 3
							  },
							  "priority": 7,
							  "providerName": "AWS",
							  "readOnlySpecs": {
								"instanceSize": "M10",
								"diskIOPS": 100,
								"ebsVolumeType": "STANDARD",
								"nodeCount": 0
							  },
							  "regionName": "US_EAST_1"
							}
						  ],
						  "zoneName": "Zone 1"
						}
					  ],
					  "rootCertType": "DST",
					  "stateName": "CREATING"
					},
					{
						  "backupEnabled": true,
						  "biConnector": {
							"enabled": false,
							"readPreference": "secondary"
						  },
						  "clusterType": "GEOSHARDED",
						  "createDate": "2021-03-02T22:27:46Z",
						  "diskSizeGB": 40.0,
						  "encryptionAtRestProvider": "NONE",
						  "groupId": "2",
						  "id": "2",
						  "mongoDBMajorVersion": "4.4",
						  "mongoDBVersion": "4.4.4",
						  "name": "globalCluster",
						  "paused": false,
						  "pitEnabled": true,
						  "replicationSpecs": [
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "NORTH_AMERICA_NORTHEAST_1"
								}
							  ],
							  "zoneName": "Zone 1"
							},
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "ASIA_NORTHEAST_2"
								}
							  ],
							  "zoneName": "Zone 2"
							},
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "EUROPE_WEST_6"
								}
							  ],
							  "zoneName": "Zone 3"
							}
						  ],
						  "rootCertType": "DST",
						  "stateName": "CREATING",
						  "versionReleaseSystem": "LTS"
						}
				  ],
				"totalCount": 1
}`)
	})

	clusters, _, err := client.AdvancedClusters.List(ctx, "1", nil)
	if err != nil {
		t.Fatalf("Clusters.List returned error: %v", err)
	}

	backupEnabled := true
	enabled := false
	paused := false
	pitEnabled := true
	diskSizeGB := 10.0
	diskSizeGBGeoSharded := 40.0
	priority := 7
	diskIOPS := int64(100)
	nodeCountZero := 0
	nodeCount := 3
	expected := &AdvancedClustersResponse{

		Links: []*Link{{
			Rel:  "self",
			Href: "https://cloud.mongodb.com/api/atlas/v1.5/groups/1/clusters?pageNum=1&itemsPerPage=100",
		}},

		Results: []*AdvancedCluster{
			{
				BackupEnabled: &backupEnabled,
				BiConnector: &BiConnector{
					Enabled:        &enabled,
					ReadPreference: "secondary",
				},
				ClusterType:              "REPLICASET",
				DiskSizeGB:               &diskSizeGB,
				EncryptionAtRestProvider: "NONE",
				GroupID:                  "1",
				ID:                       "1",
				MongoDBMajorVersion:      "4.4",
				MongoDBVersion:           "4.4.4",
				Name:                     "basicReplicaSet",
				Paused:                   &paused,
				PitEnabled:               &pitEnabled,
				StateName:                "CREATING",

				ReplicationSpecs: []*AdvancedReplicationSpec{
					{
						ID:        "1",
						NumShards: 1,
						ZoneName:  "Zone 1",
						RegionConfigs: []*AdvancedRegionConfig{
							{
								AnalyticsAutoScaling: &AdvancedAutoScaling{
									DiskGB: &DiskGB{
										Enabled: pointy.Bool(true),
									},
									Compute: &Compute{
										Enabled:          pointy.Bool(true),
										ScaleDownEnabled: pointy.Bool(false),
										MinInstanceSize:  "M10",
										MaxInstanceSize:  "M20",
									},
								},
								AnalyticsSpecs: &Specs{
									DiskIOPS:      &diskIOPS,
									EbsVolumeType: "STANDARD",
									InstanceSize:  "M10",
									NodeCount:     &nodeCountZero,
								},
								ElectableSpecs: &Specs{
									DiskIOPS:      &diskIOPS,
									EbsVolumeType: "STANDARD",
									InstanceSize:  "M10",
									NodeCount:     &nodeCount,
								},
								ReadOnlySpecs: &Specs{
									DiskIOPS:      &diskIOPS,
									EbsVolumeType: "STANDARD",
									InstanceSize:  "M10",
									NodeCount:     &nodeCountZero,
								},
								Priority:     &priority,
								ProviderName: "AWS",
								RegionName:   "US_EAST_1",
							},
						},
					},
				},
				CreateDate:   "2021-03-02T22:25:18Z",
				RootCertType: "DST",
			},
			{
				BackupEnabled: &backupEnabled,
				BiConnector: &BiConnector{
					Enabled:        &enabled,
					ReadPreference: "secondary",
				},
				ClusterType:              "GEOSHARDED",
				DiskSizeGB:               &diskSizeGBGeoSharded,
				EncryptionAtRestProvider: "NONE",
				GroupID:                  "2",
				ID:                       "2",
				MongoDBMajorVersion:      "4.4",
				MongoDBVersion:           "4.4.4",
				Name:                     "globalCluster",
				Paused:                   &paused,
				PitEnabled:               &pitEnabled,
				StateName:                "CREATING",
				ReplicationSpecs: []*AdvancedReplicationSpec{
					{
						ID:        "2",
						NumShards: 1,
						ZoneName:  "Zone 1",
						RegionConfigs: []*AdvancedRegionConfig{
							{
								AnalyticsSpecs: &Specs{
									InstanceSize: "M30",
									NodeCount:    &nodeCountZero,
								},
								ElectableSpecs: &Specs{
									InstanceSize: "M30",
									NodeCount:    &nodeCount,
								},
								ReadOnlySpecs: &Specs{
									InstanceSize: "M30",
									NodeCount:    &nodeCountZero,
								},
								Priority:     &priority,
								ProviderName: "GCP",
								RegionName:   "NORTH_AMERICA_NORTHEAST_1",
							},
						},
					},
					{
						ID:        "2",
						NumShards: 1,
						ZoneName:  "Zone 2",
						RegionConfigs: []*AdvancedRegionConfig{
							{
								AnalyticsSpecs: &Specs{
									InstanceSize: "M30",
									NodeCount:    &nodeCountZero,
								},
								ElectableSpecs: &Specs{
									InstanceSize: "M30",
									NodeCount:    &nodeCount,
								},
								ReadOnlySpecs: &Specs{
									InstanceSize: "M30",
									NodeCount:    &nodeCountZero,
								},
								Priority:     &priority,
								ProviderName: "GCP",
								RegionName:   "ASIA_NORTHEAST_2",
							},
						},
					},
					{
						ID:        "2",
						NumShards: 1,
						ZoneName:  "Zone 3",
						RegionConfigs: []*AdvancedRegionConfig{
							{
								AnalyticsSpecs: &Specs{
									InstanceSize: "M30",
									NodeCount:    &nodeCountZero,
								},
								ElectableSpecs: &Specs{
									InstanceSize: "M30",
									NodeCount:    &nodeCount,
								},
								ReadOnlySpecs: &Specs{
									InstanceSize: "M30",
									NodeCount:    &nodeCountZero,
								},
								Priority:     &priority,
								ProviderName: "GCP",
								RegionName:   "EUROPE_WEST_6",
							},
						},
					}},
				CreateDate:           "2021-03-02T22:27:46Z",
				RootCertType:         "DST",
				VersionReleaseSystem: "LTS",
			}},
		TotalCount: 1}

	if diff := deep.Equal(clusters, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAdvancedClusters_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.5/groups/%s/clusters/%s", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
						  "backupEnabled": true,
						  "biConnector": {
							"enabled": false,
							"readPreference": "secondary"
						  },
						  "clusterType": "GEOSHARDED",
						  "createDate": "2021-03-02T22:27:46Z",
						  "diskSizeGB": 40.0,
						  "encryptionAtRestProvider": "NONE",
						  "groupId": "2",
						  "id": "2",
						  "mongoDBMajorVersion": "4.4",
						  "mongoDBVersion": "4.4.4",
						  "name": "globalCluster",
						  "paused": false,
						  "pitEnabled": true,
						  "replicationSpecs": [
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "NORTH_AMERICA_NORTHEAST_1"
								}
							  ],
							  "zoneName": "Zone 1"
							},
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "ASIA_NORTHEAST_2"
								}
							  ],
							  "zoneName": "Zone 2"
							},
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "EUROPE_WEST_6"
								}
							  ],
							  "zoneName": "Zone 3"
							}
						  ],
						  "rootCertType": "DST",
						  "stateName": "CREATING",
						  "terminationProtectionEnabled": false,
						  "versionReleaseSystem": "LTS"
						
}`)
	})

	cluster, _, err := client.AdvancedClusters.Get(ctx, groupID, clusterName)
	if err != nil {
		t.Fatalf("Clusters.Get returned error: %v", err)
	}

	backupEnabled := true
	enabled := false
	paused := false
	pitEnabled := true
	diskSizeGBGeoSharded := 40.0
	priority := 7
	nodeCountZero := 0
	nodeCount := 3

	expected := &AdvancedCluster{
		BackupEnabled: &backupEnabled,
		BiConnector: &BiConnector{
			Enabled:        &enabled,
			ReadPreference: "secondary",
		},
		ClusterType:                  "GEOSHARDED",
		DiskSizeGB:                   &diskSizeGBGeoSharded,
		EncryptionAtRestProvider:     "NONE",
		GroupID:                      "2",
		ID:                           "2",
		MongoDBMajorVersion:          "4.4",
		MongoDBVersion:               "4.4.4",
		Name:                         clusterName,
		Paused:                       &paused,
		PitEnabled:                   &pitEnabled,
		StateName:                    "CREATING",
		TerminationProtectionEnabled: pointy.Bool(false),
		VersionReleaseSystem:         "LTS",
		ReplicationSpecs: []*AdvancedReplicationSpec{
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 1",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "NORTH_AMERICA_NORTHEAST_1",
					},
				},
			},
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 2",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "ASIA_NORTHEAST_2",
					},
				},
			},
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 3",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "EUROPE_WEST_6",
					},
				},
			}},
		CreateDate:   "2021-03-02T22:27:46Z",
		RootCertType: "DST",
	}

	if diff := deep.Equal(cluster, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAdvancedClusters_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.5/groups/%s/clusters", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
						  "backupEnabled": true,
						  "biConnector": {
							"enabled": false,
							"readPreference": "secondary"
						  },
						  "clusterType": "GEOSHARDED",
						  "createDate": "2021-03-02T22:27:46Z",
						  "diskSizeGB": 40.0,
						  "encryptionAtRestProvider": "NONE",
						  "groupId": "2",
						  "id": "2",
						  "mongoDBMajorVersion": "4.4",
						  "mongoDBVersion": "4.4.4",
						  "name": "globalCluster",
						  "paused": false,
						  "pitEnabled": true,
						  "replicationSpecs": [
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "NORTH_AMERICA_NORTHEAST_1"
								}
							  ],
							  "zoneName": "Zone 1"
							},
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "ASIA_NORTHEAST_2"
								}
							  ],
							  "zoneName": "Zone 2"
							},
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "EUROPE_WEST_6"
								}
							  ],
							  "zoneName": "Zone 3"
							}
						  ],
						  "rootCertType": "DST",
						  "stateName": "CREATING",
						  "versionReleaseSystem": "LTS"
						
}`)
	})

	backupEnabled := true
	enabled := false
	paused := false
	pitEnabled := true
	diskSizeGBGeoSharded := 40.0
	priority := 7
	nodeCountZero := 0
	nodeCount := 3

	requestCluster := &AdvancedCluster{
		BackupEnabled: &backupEnabled,
		BiConnector: &BiConnector{
			Enabled:        &enabled,
			ReadPreference: "secondary",
		},
		ClusterType:              "GEOSHARDED",
		DiskSizeGB:               &diskSizeGBGeoSharded,
		EncryptionAtRestProvider: "NONE",
		GroupID:                  "2",
		ID:                       "2",
		MongoDBMajorVersion:      "4.4",
		MongoDBVersion:           "4.4.4",
		Name:                     clusterName,
		Paused:                   &paused,
		PitEnabled:               &pitEnabled,
		StateName:                "CREATING",
		ReplicationSpecs: []*AdvancedReplicationSpec{
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 1",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "NORTH_AMERICA_NORTHEAST_1",
					},
				},
			},
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 2",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "ASIA_NORTHEAST_2",
					},
				},
			},
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 3",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "EUROPE_WEST_6",
					},
				},
			}},
		CreateDate:           "2021-03-02T22:27:46Z",
		RootCertType:         "DST",
		VersionReleaseSystem: "LTS",
	}

	cluster, _, err := client.AdvancedClusters.Create(ctx, groupID, requestCluster)
	if err != nil {
		t.Fatalf("Clusters.Create returned error: %v", err)
	}

	expected := &AdvancedCluster{
		BackupEnabled: &backupEnabled,
		BiConnector: &BiConnector{
			Enabled:        &enabled,
			ReadPreference: "secondary",
		},
		ClusterType:              "GEOSHARDED",
		DiskSizeGB:               &diskSizeGBGeoSharded,
		EncryptionAtRestProvider: "NONE",
		GroupID:                  "2",
		ID:                       "2",
		MongoDBMajorVersion:      "4.4",
		MongoDBVersion:           "4.4.4",
		Name:                     clusterName,
		Paused:                   &paused,
		PitEnabled:               &pitEnabled,
		StateName:                "CREATING",
		ReplicationSpecs: []*AdvancedReplicationSpec{
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 1",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "NORTH_AMERICA_NORTHEAST_1",
					},
				},
			},
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 2",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "ASIA_NORTHEAST_2",
					},
				},
			},
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 3",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "EUROPE_WEST_6",
					},
				},
			}},
		CreateDate:           "2021-03-02T22:27:46Z",
		RootCertType:         "DST",
		VersionReleaseSystem: "LTS",
	}

	if diff := deep.Equal(cluster, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAdvancedClusters_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.5/groups/%s/clusters/%s", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
						  "backupEnabled": true,
						  "biConnector": {
							"enabled": false,
							"readPreference": "secondary"
						  },
						  "clusterType": "GEOSHARDED",
						  "createDate": "2021-03-02T22:27:46Z",
						  "diskSizeGB": 40.0,
						  "encryptionAtRestProvider": "NONE",
						  "groupId": "2",
						  "id": "2",
						  "mongoDBMajorVersion": "4.4",
						  "mongoDBVersion": "4.4.4",
						  "name": "globalCluster",
						  "paused": false,
						  "pitEnabled": true,
						  "replicationSpecs": [
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "NORTH_AMERICA_NORTHEAST_1"
								}
							  ],
							  "zoneName": "Zone 1"
							},
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "ASIA_NORTHEAST_2"
								}
							  ],
							  "zoneName": "Zone 2"
							},
							{
							  "id": "2",
							  "numShards": 1,
							  "regionConfigs": [
								{
								  "analyticsSpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "electableSpecs": {
									"instanceSize": "M30",
									"nodeCount": 3
								  },
								  "priority": 7,
								  "providerName": "GCP",
								  "readOnlySpecs": {
									"instanceSize": "M30",
									"nodeCount": 0
								  },
								  "regionName": "EUROPE_WEST_6"
								}
							  ],
							  "zoneName": "Zone 3"
							}
						  ],
						  "rootCertType": "DST",
						  "stateName": "CREATING"
						
}`)
	})

	backupEnabled := true
	enabled := false
	paused := false
	pitEnabled := true
	diskSizeGBGeoSharded := 40.0
	priority := 7
	nodeCountZero := 0
	nodeCount := 3

	requestCluster := &AdvancedCluster{
		BackupEnabled: &backupEnabled,
		BiConnector: &BiConnector{
			Enabled:        &enabled,
			ReadPreference: "secondary",
		},
		ClusterType:              "GEOSHARDED",
		DiskSizeGB:               &diskSizeGBGeoSharded,
		EncryptionAtRestProvider: "NONE",
		GroupID:                  "2",
		ID:                       "2",
		MongoDBMajorVersion:      "4.4",
		MongoDBVersion:           "4.4.4",
		Name:                     clusterName,
		Paused:                   &paused,
		PitEnabled:               &pitEnabled,
		StateName:                "CREATING",
		ReplicationSpecs: []*AdvancedReplicationSpec{
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 1",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "NORTH_AMERICA_NORTHEAST_1",
					},
				},
			},
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 2",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "ASIA_NORTHEAST_2",
					},
				},
			},
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 3",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "EUROPE_WEST_6",
					},
				},
			}},
		CreateDate:   "2021-03-02T22:27:46Z",
		RootCertType: "DST",
	}

	cluster, _, err := client.AdvancedClusters.Update(ctx, groupID, clusterName, requestCluster)
	if err != nil {
		t.Fatalf("Clusters.Update returned error: %v", err)
	}

	expected := &AdvancedCluster{
		BackupEnabled: &backupEnabled,
		BiConnector: &BiConnector{
			Enabled:        &enabled,
			ReadPreference: "secondary",
		},
		ClusterType:              "GEOSHARDED",
		DiskSizeGB:               &diskSizeGBGeoSharded,
		EncryptionAtRestProvider: "NONE",
		GroupID:                  "2",
		ID:                       "2",
		MongoDBMajorVersion:      "4.4",
		MongoDBVersion:           "4.4.4",
		Name:                     clusterName,
		Paused:                   &paused,
		PitEnabled:               &pitEnabled,
		StateName:                "CREATING",
		ReplicationSpecs: []*AdvancedReplicationSpec{
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 1",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "NORTH_AMERICA_NORTHEAST_1",
					},
				},
			},
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 2",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "ASIA_NORTHEAST_2",
					},
				},
			},
			{
				ID:        "2",
				NumShards: 1,
				ZoneName:  "Zone 3",
				RegionConfigs: []*AdvancedRegionConfig{
					{
						AnalyticsSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						ElectableSpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCount,
						},
						ReadOnlySpecs: &Specs{
							InstanceSize: "M30",
							NodeCount:    &nodeCountZero,
						},
						Priority:     &priority,
						ProviderName: "GCP",
						RegionName:   "EUROPE_WEST_6",
					},
				},
			}},
		CreateDate:   "2021-03-02T22:27:46Z",
		RootCertType: "DST",
	}

	if diff := deep.Equal(cluster, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAdvancedClusters_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.5/groups/%s/clusters/%s", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.AdvancedClusters.Delete(ctx, groupID, clusterName)
	if err != nil {
		t.Fatalf("Cluster.Delete returned error: %v", err)
	}
}

func TestFailover(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.5/groups/%s/clusters/%s/restartPrimaries", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	_, err := client.AdvancedClusters.TestFailover(ctx, groupID, clusterName)
	if err != nil {
		t.Fatalf("Cluster.TestFailover returned error: %v", err)
	}
}
