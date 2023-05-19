// Copyright 2023 MongoDB Inc
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

func TestDataLakesPipeline_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/pipelines", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			  {
				"_id": "32b6e34b3d91647abb20e7b8",
				"createdDate": "2019-08-24T14:15:22Z",
				"groupId": "6c7498dg87d9e6526801572b",
				"lastUpdatedDate": "2019-08-24T14:15:22Z",
				"name": "test",
				"sink": {
				  "type": "DLS",
				  "metadataProvider": "AWS",
				  "metadataRegion": "test"
				},
				"source": {
				  "type": "PERIODIC_CPS",
				  "clusterName": "test",
				  "collectionName": "test",
				  "databaseName": "test",
				  "groupId": "6c7498dg87d9e6526801572b"
				},
				"state": "ACTIVE",
				"transformations": [
				  {
					"field": "test",
					"type": "EXCLUDE"
				  }
				]
			  }
		]`)
	})

	pipelines, _, err := client.DataLakePipeline.List(ctx, groupID)
	if err != nil {
		t.Fatalf("DataLakePipeline.List returned error: %v", err)
	}

	expected := []*DataLakePipeline{
		{
			ID:              "32b6e34b3d91647abb20e7b8",
			GroupID:         groupID,
			Name:            "test",
			CreatedDate:     "2019-08-24T14:15:22Z",
			LastUpdatedDate: "2019-08-24T14:15:22Z",
			State:           "ACTIVE",
			Sink: &DataLakePipelineSink{
				Type:             "DLS",
				MetadataProvider: "AWS",
				MetadataRegion:   "test",
			},
			Source: &DataLakePipelineSource{
				Type:           "PERIODIC_CPS",
				ClusterName:    "test",
				CollectionName: "test",
				DatabaseName:   "test",
				GroupID:        groupID,
			},
			Transformations: []*DataLakePipelineTransformation{
				{
					Field: "test",
					Type:  "EXCLUDE",
				},
			},
		},
	}

	if diff := deep.Equal(pipelines, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLakesPipeline_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	name := "test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/pipelines/%s", groupID, name)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
			  {
				"_id": "32b6e34b3d91647abb20e7b8",
				"createdDate": "2019-08-24T14:15:22Z",
				"groupId": "6c7498dg87d9e6526801572b",
				"lastUpdatedDate": "2019-08-24T14:15:22Z",
				"name": "test",
				"sink": {
				  "type": "DLS",
				  "metadataProvider": "AWS",
				  "metadataRegion": "test"
				},
				"source": {
				  "type": "PERIODIC_CPS",
				  "clusterName": "test",
				  "collectionName": "test",
				  "databaseName": "test",
				  "groupId": "6c7498dg87d9e6526801572b"
				},
				"state": "ACTIVE",
				"transformations": [
				  {
					"field": "test",
					"type": "EXCLUDE"
				  }
				]
			  }`)
	})

	pipeline, _, err := client.DataLakePipeline.Get(ctx, groupID, name)
	if err != nil {
		t.Fatalf("DataLakePipeline.Get returned error: %v", err)
	}

	expected := &DataLakePipeline{
		ID:              "32b6e34b3d91647abb20e7b8",
		GroupID:         groupID,
		Name:            "test",
		CreatedDate:     "2019-08-24T14:15:22Z",
		LastUpdatedDate: "2019-08-24T14:15:22Z",
		State:           "ACTIVE",
		Sink: &DataLakePipelineSink{
			Type:             "DLS",
			MetadataProvider: "AWS",
			MetadataRegion:   "test",
		},
		Source: &DataLakePipelineSource{
			Type:           "PERIODIC_CPS",
			ClusterName:    "test",
			CollectionName: "test",
			DatabaseName:   "test",
			GroupID:        groupID,
		},
		Transformations: []*DataLakePipelineTransformation{
			{
				Field: "test",
				Type:  "EXCLUDE",
			},
		},
	}

	if diff := deep.Equal(pipeline, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLakesPipeline_GetRun(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	name := "test"
	id := "1"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/pipelines/%s/runs/%s", groupID, name, id)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			  "_id": "32b6e34b3d91647abb20e7b8",
			  "backupFrequencyType": "HOURLY",
			  "createdDate": "2019-08-24T14:15:22Z",
			  "datasetName": "test",
			  "groupId": "32b6e34b3d91647abb20e7b8",
			  "lastUpdatedDate": "2019-08-24T14:15:22Z",
			  "phase": "SNAPSHOT",
			  "pipelineId": "32b6e34b3d91647abb20e7b8",
			  "snapshotId": "32b6e34b3d91647abb20e7b8",
			  "state": "PENDING",
			  "stats": {
				"bytesExported": 0,
				"numDocs": 0
			  }
		}`)
	})

	pipeline, _, err := client.DataLakePipeline.GetRun(ctx, groupID, name, id)
	if err != nil {
		t.Fatalf("DataLakePipeline.GetRun returned error: %v", err)
	}

	expected := &DataLakePipelineRun{
		ID:                  "32b6e34b3d91647abb20e7b8",
		BackupFrequencyType: "HOURLY",
		CreatedDate:         "2019-08-24T14:15:22Z",
		DatasetName:         "test",
		GroupID:             "32b6e34b3d91647abb20e7b8",
		LastUpdatedDate:     "2019-08-24T14:15:22Z",
		Phase:               "SNAPSHOT",
		PipelineID:          "32b6e34b3d91647abb20e7b8",
		SnapshotID:          "32b6e34b3d91647abb20e7b8",
		State:               "PENDING",
		Stats: &DataLakePipelineRunStats{
			BytesExported: 0,
			NumDocs:       0,
		},
	}

	if diff := deep.Equal(pipeline, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLakesPipeline_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/pipelines", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
			  {
				"_id": "32b6e34b3d91647abb20e7b8",
				"createdDate": "2019-08-24T14:15:22Z",
				"groupId": "6c7498dg87d9e6526801572b",
				"lastUpdatedDate": "2019-08-24T14:15:22Z",
				"name": "test",
				"sink": {
				  "type": "DLS",
				  "metadataProvider": "AWS",
				  "metadataRegion": "test"
				},
				"source": {
				  "type": "PERIODIC_CPS",
				  "clusterName": "test",
				  "collectionName": "test",
				  "databaseName": "test",
				  "groupId": "6c7498dg87d9e6526801572b"
				},
				"state": "ACTIVE",
				"transformations": [
				  {
					"field": "test",
					"type": "EXCLUDE"
				  }
				]
			  }`)
	})

	requestBody := &DataLakePipeline{
		Name: "test",
		Sink: &DataLakePipelineSink{
			Type:             "DLS",
			MetadataProvider: "AWS",
			MetadataRegion:   "test",
		},
		Source: &DataLakePipelineSource{
			Type:           "PERIODIC_CPS",
			ClusterName:    "test",
			CollectionName: "test",
			DatabaseName:   "test",
		},
		Transformations: []*DataLakePipelineTransformation{
			{
				Field: "test",
				Type:  "EXCLUDE",
			},
		},
	}

	pipeline, _, err := client.DataLakePipeline.Create(ctx, groupID, requestBody)
	if err != nil {
		t.Fatalf("DataLakePipeline.Create returned error: %v", err)
	}

	expected := &DataLakePipeline{
		ID:              "32b6e34b3d91647abb20e7b8",
		GroupID:         groupID,
		Name:            "test",
		CreatedDate:     "2019-08-24T14:15:22Z",
		LastUpdatedDate: "2019-08-24T14:15:22Z",
		State:           "ACTIVE",
		Sink: &DataLakePipelineSink{
			Type:             "DLS",
			MetadataProvider: "AWS",
			MetadataRegion:   "test",
		},
		Source: &DataLakePipelineSource{
			Type:           "PERIODIC_CPS",
			ClusterName:    "test",
			CollectionName: "test",
			DatabaseName:   "test",
			GroupID:        groupID,
		},
		Transformations: []*DataLakePipelineTransformation{
			{
				Field: "test",
				Type:  "EXCLUDE",
			},
		},
	}

	if diff := deep.Equal(pipeline, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLakesPipeline_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	name := "test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/pipelines/%s", groupID, name)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `
			  {
				"_id": "32b6e34b3d91647abb20e7b8",
				"createdDate": "2019-08-24T14:15:22Z",
				"groupId": "6c7498dg87d9e6526801572b",
				"lastUpdatedDate": "2019-08-24T14:15:22Z",
				"name": "test",
				"sink": {
				  "type": "DLS",
				  "metadataProvider": "AWS",
				  "metadataRegion": "test"
				},
				"source": {
				  "type": "PERIODIC_CPS",
				  "clusterName": "test",
				  "collectionName": "test",
				  "databaseName": "test",
				  "groupId": "6c7498dg87d9e6526801572b"
				},
				"state": "ACTIVE",
				"transformations": [
				  {
					"field": "test",
					"type": "EXCLUDE"
				  }
				]
			  }`)
	})

	requestBody := &DataLakePipeline{
		Name: "test",
		Sink: &DataLakePipelineSink{
			Type:             "DLS",
			MetadataProvider: "AWS",
			MetadataRegion:   "test",
		},
		Source: &DataLakePipelineSource{
			Type:           "PERIODIC_CPS",
			ClusterName:    "test",
			CollectionName: "test",
			DatabaseName:   "test",
		},
		Transformations: []*DataLakePipelineTransformation{
			{
				Field: "test",
				Type:  "EXCLUDE",
			},
		},
	}

	pipeline, _, err := client.DataLakePipeline.Update(ctx, groupID, name, requestBody)
	if err != nil {
		t.Fatalf("DataLakePipeline.Update returned error: %v", err)
	}

	expected := &DataLakePipeline{
		ID:              "32b6e34b3d91647abb20e7b8",
		GroupID:         groupID,
		Name:            "test",
		CreatedDate:     "2019-08-24T14:15:22Z",
		LastUpdatedDate: "2019-08-24T14:15:22Z",
		State:           "ACTIVE",
		Sink: &DataLakePipelineSink{
			Type:             "DLS",
			MetadataProvider: "AWS",
			MetadataRegion:   "test",
		},
		Source: &DataLakePipelineSource{
			Type:           "PERIODIC_CPS",
			ClusterName:    "test",
			CollectionName: "test",
			DatabaseName:   "test",
			GroupID:        groupID,
		},
		Transformations: []*DataLakePipelineTransformation{
			{
				Field: "test",
				Type:  "EXCLUDE",
			},
		},
	}

	if diff := deep.Equal(pipeline, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLakesPipeline_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	name := "test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/pipelines/%s", groupID, name)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.DataLakePipeline.Delete(ctx, groupID, name)
	if err != nil {
		t.Errorf("DataLakePipeline.Delete returned error: %v", err)
	}
}

func TestDataLakesPipeline_ListSnapshots(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	name := "test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/pipelines/%s/availableSnapshots", groupID, name)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			  "links": [
				{
				  "href": "https://cloud.mongodb.com/api/atlas",
				  "rel": "self"
				}
			  ],
			  "results": [
				{
				  "createdAt": "2019-08-24T14:15:22Z",
				  "description": "string",
				  "expiresAt": "2019-08-24T14:15:22Z",
				  "frequencyType": "hourly",
				  "id": "32b6e34b3d91647abb20e7b8",
				  "links": [
					{
					  "href": "https://cloud.mongodb.com/api/atlas",
					  "rel": "self"
					}
				  ],
				  "masterKeyUUID": "72659f08-8b3c-4913-bb4e-a8a68e036502",
				  "mongodVersion": "string",
				  "policyItems": [
					"32b6e34b3d91647abb20e7b8"
				  ],
				  "snapshotType": "onDemand",
				  "status": "queued",
				  "storageSizeBytes": 0,
				  "type": "REPLICA_SET"
				}
			  ],
			  "totalCount": 1
			}`)
	})

	pipelines, _, err := client.DataLakePipeline.ListSnapshots(ctx, groupID, name, nil)
	if err != nil {
		t.Fatalf("DataLakePipeline.ListSnapshots returned error: %v", err)
	}

	expected := &DataLakePipelineSnapshotsResponse{
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas",
			},
		},
		Results: []*DataLakePipelineSnapshot{
			{
				ID:               "32b6e34b3d91647abb20e7b8",
				CreatedAt:        "2019-08-24T14:15:22Z",
				Description:      "string",
				ExpiresAt:        "2019-08-24T14:15:22Z",
				FrequencyType:    "hourly",
				MasterKeyUUID:    "72659f08-8b3c-4913-bb4e-a8a68e036502",
				MongodVersion:    "string",
				SnapshotType:     "onDemand",
				Status:           "queued",
				Type:             "REPLICA_SET",
				StorageSizeBytes: 0,
				PolicyItems:      []string{"32b6e34b3d91647abb20e7b8"},
				Links: []*Link{
					{
						Rel:  "self",
						Href: "https://cloud.mongodb.com/api/atlas",
					},
				},
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(pipelines, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLakesPipeline_ListIngestionSchedules(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	name := "test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/pipelines/%s/availableSchedules", groupID, name)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"frequencyInterval": 1,
				"frequencyType": "daily",
				"id": "id",
				"retentionUnit": "days",
				"retentionValue": 0
			}
		]`)
	})

	pipelines, _, err := client.DataLakePipeline.ListIngestionSchedules(ctx, groupID, name)
	if err != nil {
		t.Fatalf("DataLakePipeline.ListIngestionSchedules returned error: %v", err)
	}

	expected := []*DataLakePipelineIngestionSchedule{
		{
			ID:                "id",
			FrequencyType:     "daily",
			RetentionUnit:     "days",
			FrequencyInterval: 1,
			RetentionValue:    0,
		},
	}

	if diff := deep.Equal(pipelines, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDataLakesPipeline_ListRuns(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6c7498dg87d9e6526801572b"
	name := "test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/pipelines/%s/runs", groupID, name)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				  "links": [
					{
					  "href": "https://cloud.mongodb.com/api/atlas",
					  "rel": "self"
					}
				  ],
				  "results": [
					{
					  "_id": "32b6e34b3d91647abb20e7b8",
					  "backupFrequencyType": "HOURLY",
					  "createdDate": "2019-08-24T14:15:22Z",
					  "datasetName": "test",
					  "groupId": "32b6e34b3d91647abb20e7b8",
					  "lastUpdatedDate": "2019-08-24T14:15:22Z",
					  "phase": "SNAPSHOT",
					  "pipelineId": "32b6e34b3d91647abb20e7b8",
					  "snapshotId": "32b6e34b3d91647abb20e7b8",
					  "state": "PENDING",
					  "stats": {
						"bytesExported": 0,
						"numDocs": 0
					  }
					}
				  ],
				  "totalCount": 1
				}`)
	})

	pipelines, _, err := client.DataLakePipeline.ListRuns(ctx, groupID, name)
	if err != nil {
		t.Fatalf("DataLakePipeline.ListRuns returned error: %v", err)
	}

	expected := &DataLakePipelineRunsResponse{
		Links: []*Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas"},
		},
		Results: []*DataLakePipelineRun{
			{
				ID:                  "32b6e34b3d91647abb20e7b8",
				BackupFrequencyType: "HOURLY",
				CreatedDate:         "2019-08-24T14:15:22Z",
				DatasetName:         "test",
				GroupID:             "32b6e34b3d91647abb20e7b8",
				LastUpdatedDate:     "2019-08-24T14:15:22Z",
				Phase:               "SNAPSHOT",
				PipelineID:          "32b6e34b3d91647abb20e7b8",
				SnapshotID:          "32b6e34b3d91647abb20e7b8",
				State:               "PENDING",
				Stats: &DataLakePipelineRunStats{
					BytesExported: 0,
					NumDocs:       0,
				},
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(pipelines, expected); diff != nil {
		t.Error(diff)
	}
}
