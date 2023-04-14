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
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestBackupCompliancePolicy_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/backupCompliancePolicy", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"authorizedEmail": "user@example.com",
			"copyProtectionEnabled": false,
			"encryptionAtRestEnabled": false,
			"onDemandPolicyItem": 
			  {
			    "frequencyInterval": 1,
			    "frequencyType": "daily",
			    "id": "32b6e34b3d91647abb20e7b9",
			    "retentionUnit": "days",
			    "retentionValue": 0
		      },
		"pitEnabled": false,
		"projectId": "32b6e34b3d91647abb20e7b8",
		"restoreWindowDays": 7,
		"scheduledPolicyItems": 
		[
		  {
			"frequencyInterval": 1,
			"frequencyType": "daily",
			"id": "32b6e34b3d91647abb20e7b9",
			"retentionUnit": "days",
			"retentionValue": 0
		  },
		  {
			"frequencyInterval": 1,
			"frequencyType": "hourly",
			"id": "5c95242c87d9d636e70c28f2",
			"retentionUnit": "days",
			"retentionValue": 0
		  }
		],
		"state": "ACTIVE",
		"updatedDate": "2019-08-24T14:15:22Z",
		"updatedUser": "user@example.com"
	  }`)
	})

	backupCompliancePolicy, _, err := client.BackupCompliancePolicy.Get(ctx, groupID)
	if err != nil {
		t.Fatalf("BackupCompliancePolicy.Get returned error: %v", err)
	}

	expected := &BackupCompliancePolicy{
		AuthorizedEmail:         "user@example.com",
		CopyProtectionEnabled:   pointer(false),
		EncryptionAtRestEnabled: pointer(false),
		ProjectID:               "32b6e34b3d91647abb20e7b8",
		RestoreWindowDays:       pointer(int64(7)),
		State:                   "ACTIVE",
		UpdatedDate:             "2019-08-24T14:15:22Z",
		UpdatedUser:             "user@example.com",
		PitEnabled:              pointer(false),
		OnDemandPolicyItem: PolicyItem{
			ID:                "32b6e34b3d91647abb20e7b9",
			FrequencyInterval: 1,
			FrequencyType:     "daily",
			RetentionUnit:     "days",
			RetentionValue:    0,
		},
		ScheduledPolicyItems: []ScheduledPolicyItem{
			{
				ID:                "32b6e34b3d91647abb20e7b9",
				FrequencyInterval: 1,
				FrequencyType:     "daily",
				RetentionUnit:     "days",
				RetentionValue:    0,
			},
			{
				ID:                "5c95242c87d9d636e70c28f2",
				FrequencyInterval: 1,
				FrequencyType:     "hourly",
				RetentionUnit:     "days",
				RetentionValue:    0,
			},
		},
	}

	if diff := deep.Equal(backupCompliancePolicy, expected); diff != nil {
		t.Error(diff)
	}
}

func TestBackupCompliancePolicy_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "5b6212af90dc76637950a2c6"

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/backupCompliancePolicy", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"authorizedEmail":         "user@example.com",
			"copyProtectionEnabled":   false,
			"encryptionAtRestEnabled": false,
			"updatedDate":             "2019-08-24T14:15:22Z",
			"updatedUser":             "user@example.com",
			"pitEnabled":              false,
			"projectId":               "32b6e34b3d91647abb20e7b8",
			"restoreWindowDays":       float64(7),
			"state":                   "ACTIVE",
			"onDemandPolicyItem": map[string]interface{}{
				"id":                "32b6e34b3d91647abb20e7b9",
				"frequencyType":     "daily",
				"frequencyInterval": float64(1),
				"retentionValue":    float64(7),
				"retentionUnit":     "days",
			},
			"scheduledPolicyItems": []interface{}{
				map[string]interface{}{
					"id":                "5c95242c87d9d636e70c28f0",
					"frequencyType":     "hourly",
					"frequencyInterval": float64(6),
					"retentionValue":    float64(2),
					"retentionUnit":     "days",
				},
				map[string]interface{}{
					"id":                "5c95242c87d9d636e70c28f2",
					"frequencyType":     "weekly",
					"frequencyInterval": float64(1),
					"retentionValue":    float64(3),
					"retentionUnit":     "weeks",
				},
			},
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
			"authorizedEmail": "user@example.com",
			"copyProtectionEnabled": false,
			"encryptionAtRestEnabled": false,
			"onDemandPolicyItem": 
			  {
				"id": "32b6e34b3d91647abb20e7b9",
			    "frequencyInterval": 1,
			    "frequencyType": "daily",
				"retentionUnit": "days",
				"retentionValue": 7
		      },
		"pitEnabled": false,
		"projectId": "32b6e34b3d91647abb20e7b8",
		"restoreWindowDays": 7,
		"scheduledPolicyItems": 
		[
		  {
			"id": "5c95242c87d9d636e70c28f0",
			"frequencyInterval": 6,
			"frequencyType": "hourly",
			"retentionUnit": "days",
			"retentionValue": 2
		  },
		  {
			"id": "5c95242c87d9d636e70c28f2",
			"frequencyInterval": 1,
			"frequencyType": "weekly",
			"retentionUnit": "weeks",
			"retentionValue": 3
		  }
		],
		"state": "ACTIVE",
		"updatedDate": "2019-08-24T14:15:22Z",
		"updatedUser": "user@example.com"
	  }`)
	})

	updateRequest := &BackupCompliancePolicy{
		AuthorizedEmail:         "user@example.com",
		CopyProtectionEnabled:   pointer(false),
		EncryptionAtRestEnabled: pointer(false),
		ProjectID:               "32b6e34b3d91647abb20e7b8",
		RestoreWindowDays:       pointer(int64(7)),
		State:                   "ACTIVE",
		UpdatedDate:             "2019-08-24T14:15:22Z",
		UpdatedUser:             "user@example.com",
		PitEnabled:              pointer(false),
		OnDemandPolicyItem: PolicyItem{
			ID:                "32b6e34b3d91647abb20e7b9",
			FrequencyInterval: 1,
			FrequencyType:     "daily",
			RetentionUnit:     "days",
			RetentionValue:    7,
		},
		ScheduledPolicyItems: []ScheduledPolicyItem{
			{
				ID:                "5c95242c87d9d636e70c28f0",
				FrequencyInterval: 6,
				FrequencyType:     "hourly",
				RetentionUnit:     "days",
				RetentionValue:    2,
			},
			{
				ID:                "5c95242c87d9d636e70c28f2",
				FrequencyInterval: 1,
				FrequencyType:     "weekly",
				RetentionUnit:     "weeks",
				RetentionValue:    3,
			},
		},
	}

	backupCompliancePolicy, _, err := client.BackupCompliancePolicy.Update(ctx, groupID, updateRequest)
	if err != nil {
		t.Fatalf("BackupCompliancePolicy.Update returned error: %v", err)
	}

	expected := &BackupCompliancePolicy{
		AuthorizedEmail:         "user@example.com",
		CopyProtectionEnabled:   pointer(false),
		EncryptionAtRestEnabled: pointer(false),
		ProjectID:               "32b6e34b3d91647abb20e7b8",
		RestoreWindowDays:       pointer(int64(7)),
		State:                   "ACTIVE",
		UpdatedDate:             "2019-08-24T14:15:22Z",
		UpdatedUser:             "user@example.com",
		PitEnabled:              pointer(false),
		OnDemandPolicyItem: PolicyItem{
			ID:                "32b6e34b3d91647abb20e7b9",
			FrequencyInterval: 1,
			FrequencyType:     "daily",
			RetentionUnit:     "days",
			RetentionValue:    7,
		},
		ScheduledPolicyItems: []ScheduledPolicyItem{
			{
				ID:                "5c95242c87d9d636e70c28f0",
				FrequencyInterval: 6,
				FrequencyType:     "hourly",
				RetentionUnit:     "days",
				RetentionValue:    2,
			},
			{
				ID:                "5c95242c87d9d636e70c28f2",
				FrequencyInterval: 1,
				FrequencyType:     "weekly",
				RetentionUnit:     "weeks",
				RetentionValue:    3,
			},
		},
	}

	if diff := deep.Equal(backupCompliancePolicy, expected); diff != nil {
		t.Error(diff)
	}
}
