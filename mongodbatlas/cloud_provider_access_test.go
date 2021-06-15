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

func TestCloudProviderAccessServiceOp_ListRoles(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups/1/cloudProviderAccess", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
		  "awsIamRoles": [{
			"atlasAWSAccountArn": "arn:aws:iam::123456789012:root",
			"atlasAssumedRoleExternalId": "3192be49-6e76-4b7d-a7b8-b486a8fc4483",
			"authorizedDate": "2020-08-03T20:42:49Z",
			"createdDate": "2020-07-30T20:20:36Z",
			"featureUsages": [],
			"iamAssumedRoleArn": "arn:aws:iam::772401394250:role/my-test-aws-role",
			"providerName": "AWS",
			"roleId": "5f232b94af0a6b41747bcc2d"
		  }]
		}`)
	})

	roles, _, err := client.CloudProviderAccess.ListRoles(ctx, "1")
	if err != nil {
		t.Fatalf("CloudProviderAccess.ListRoles returned error: %v", err)
	}

	expected := &CloudProviderAccessRoles{
		AWSIAMRoles: []AWSIAMRole{
			{
				AtlasAWSAccountARN:         "arn:aws:iam::123456789012:root",
				AtlasAssumedRoleExternalID: "3192be49-6e76-4b7d-a7b8-b486a8fc4483",
				AuthorizedDate:             "2020-08-03T20:42:49Z",
				CreatedDate:                "2020-07-30T20:20:36Z",
				FeatureUsages:              []*FeatureUsage{},
				IAMAssumedRoleARN:          "arn:aws:iam::772401394250:role/my-test-aws-role",
				ProviderName:               "AWS",
				RoleID:                     "5f232b94af0a6b41747bcc2d",
			},
		},
	}
	if diff := deep.Equal(roles, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderAccessServiceOp_CreateRole(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &CloudProviderAccessRoleRequest{
		ProviderName: "AWS",
	}

	mux.HandleFunc("/api/atlas/v1.0/groups/1/cloudProviderAccess", func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"providerName": "AWS",
		}

		jsonBlob := `{
		  "atlasAWSAccountArn": "arn:aws:iam::123456789012:root",
		  "atlasAssumedRoleExternalId": "3192be49-6e76-4b7d-a7b8-b486a8fc4483",
		  "authorizedDate": null,
		  "createdDate": "2020-07-30T20:20:36Z",
		  "featureUsages": [],
		  "iamAssumedRoleArn": null,
		  "providerName": "AWS",
		  "roleId": "5f232b94af0a6b41747bcc2d"
		}`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, jsonBlob)
	})

	role, _, err := client.CloudProviderAccess.CreateRole(ctx, "1", createRequest)
	if err != nil {
		t.Fatalf("CloudProviderAccess.CreateRole returned error: %v", err)
	}

	expected := &AWSIAMRole{
		AtlasAWSAccountARN:         "arn:aws:iam::123456789012:root",
		AtlasAssumedRoleExternalID: "3192be49-6e76-4b7d-a7b8-b486a8fc4483",
		CreatedDate:                "2020-07-30T20:20:36Z",
		FeatureUsages:              []*FeatureUsage{},
		ProviderName:               "AWS",
		RoleID:                     "5f232b94af0a6b41747bcc2d",
	}
	if diff := deep.Equal(role, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderAccessServiceOp_AuthorizeRole(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	roleID := "5f232b94af0a6b41747bcc2d"

	request := &CloudProviderAuthorizationRequest{
		ProviderName:      "AWS",
		IAMAssumedRoleARN: "arn:aws:iam::772401394250:role/test-user-role",
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/1/cloudProviderAccess/%s", roleID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"providerName":      "AWS",
			"iamAssumedRoleArn": "arn:aws:iam::772401394250:role/test-user-role",
		}

		jsonBlob := `{
		  "atlasAWSAccountArn": "arn:aws:iam::123456789012:user/test.user",
		  "atlasAssumedRoleExternalId": "3192be49-6e76-4b7d-a7b8-b486a8fc4483",
		  "authorizedDate": "2020-07-30T22:17:09Z",
		  "createdDate": "2020-07-30T20:20:36Z",
		  "featureUsages": [],
		  "iamAssumedRoleArn": "arn:aws:iam::772401394250:role/test-user-role",
		  "providerName": "AWS",
		  "roleId": "5f232b94af0a6b41747bcc2d"
		}`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, jsonBlob)
	})

	role, _, err := client.CloudProviderAccess.AuthorizeRole(ctx, "1", roleID, request)
	if err != nil {
		t.Fatalf("CloudProviderAccess.AuthorizeRole returned error: %v", err)
	}

	expected := &AWSIAMRole{
		AtlasAWSAccountARN:         "arn:aws:iam::123456789012:user/test.user",
		AtlasAssumedRoleExternalID: "3192be49-6e76-4b7d-a7b8-b486a8fc4483",
		AuthorizedDate:             "2020-07-30T22:17:09Z",
		CreatedDate:                "2020-07-30T20:20:36Z",
		FeatureUsages:              []*FeatureUsage{},
		IAMAssumedRoleARN:          "arn:aws:iam::772401394250:role/test-user-role",
		ProviderName:               "AWS",
		RoleID:                     "5f232b94af0a6b41747bcc2d",
	}
	if diff := deep.Equal(role, expected); diff != nil {
		t.Error(diff)
	}
}

func TestCloudProviderAccessServiceOp_DeauthorizeRole(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	roleID := "5f232b94af0a6b41747bcc2d"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/1/cloudProviderAccess/%s/%s", "AWS", roleID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	request := &CloudProviderDeauthorizationRequest{
		GroupID:      "1",
		RoleID:       roleID,
		ProviderName: "AWS",
	}

	if _, err := client.CloudProviderAccess.DeauthorizeRole(ctx, request); err != nil {
		t.Fatalf("CloudProviderAccess.DeauthorizeRole returned error: %v", err)
	}
}
