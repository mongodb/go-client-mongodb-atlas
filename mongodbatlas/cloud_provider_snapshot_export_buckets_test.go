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

func TestCloudProviderSnapshotExportBuckets_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "test-project-id"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/backup/exportBuckets", projectID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
  "links": [
    {
      "href": "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/backup/exportBuckets?pageNum=1&itemsPerPage=100",
      "rel": "self"
    }
  ],
  "results": [
    {
      "_id": "{BUCKET-ID}",
      "bucketName": "example-bucket",
      "cloudProvider": "AWS",
      "iamRoleId": "12345678f901a234dbdb00ca"
    }
  ],
  "totalCount": 1
}`)
	})

	cloudProviderSnapshots, _, err := client.CloudProviderSnapshotExportBuckets.List(ctx, projectID, nil)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotExportBuckets.List returned error: %v", err)
	}

	expected := &CloudProviderSnapshotExportBuckets{
		Links: []*Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/{GROUP-ID}/backup/exportBuckets?pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*CloudProviderSnapshotExportBucket{
			{
				ID:            "{BUCKET-ID}",
				BucketName:    "example-bucket",
				CloudProvider: "AWS",
				IamRoleID:     "12345678f901a234dbdb00ca",
			},
		},
		TotalCount: 1,
	}

	if diff := deep.Equal(cloudProviderSnapshots, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotExportBuckets_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "test-project-id"
	bucketID := "bucket-id-test"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/backup/exportBuckets/%s", projectID, bucketID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
  "_id": "{BUCKET-ID}",
  "bucketName": "example-bucket",
  "cloudProvider": "AWS",
  "iamRoleId": "12345678f901a234dbdb00ca"
}`)
	})

	cloudProviderSnapshotBucket, _, err := client.CloudProviderSnapshotExportBuckets.Get(ctx, projectID, bucketID)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotExportBuckets.Get returned error: %v", err)
	}

	expected := &CloudProviderSnapshotExportBucket{
		ID:            "{BUCKET-ID}",
		BucketName:    "example-bucket",
		CloudProvider: "AWS",
		IamRoleID:     "12345678f901a234dbdb00ca",
	}

	if diff := deep.Equal(cloudProviderSnapshotBucket, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotExportBuckets_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "test-project-id"

	createRequest := &CloudProviderSnapshotExportBucket{
		ID:            "{BUCKET-ID}",
		BucketName:    "example-bucket",
		CloudProvider: "AWS",
		IamRoleID:     "12345678f901a234dbdb00ca",
		ProjectID:     "{PROJECT-ID}",
	}

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/backup/exportBuckets", projectID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"_id":           "{BUCKET-ID}",
			"bucketName":    "example-bucket",
			"cloudProvider": "AWS",
			"iamRoleId":     "12345678f901a234dbdb00ca",
			"projectId":     "{PROJECT-ID}",
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("Decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, `{
  "_id": "{BUCKET-ID}",
  "bucketName": "example-bucket",
  "cloudProvider": "AWS",
  "iamRoleId": "12345678f901a234dbdb00ca",
  "projectId": "{PROJECT-ID}"
}`)
	})

	cloudProviderSnapshot, _, err := client.CloudProviderSnapshotExportBuckets.Create(ctx, projectID, createRequest)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotExportBuckets.Create returned error: %v", err)
	}

	expected := &CloudProviderSnapshotExportBucket{
		ID:            "{BUCKET-ID}",
		BucketName:    "example-bucket",
		CloudProvider: "AWS",
		IamRoleID:     "12345678f901a234dbdb00ca",
		ProjectID:     "{PROJECT-ID}",
	}

	if diff := deep.Equal(cloudProviderSnapshot, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderSnapshotExportBuckets_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "project-id"
	bucketID := "bucket-id"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/backup/exportBuckets/%s", projectID, bucketID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.CloudProviderSnapshotExportBuckets.Delete(ctx, projectID, bucketID)
	if err != nil {
		t.Fatalf("CloudProviderSnapshotExportBuckets.Delete returned error: %v", err)
	}
}
