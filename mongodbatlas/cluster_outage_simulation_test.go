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
	"github.com/openlyinc/pointy"
)

func TestClusterOutageSimulationServiceOp_GetOutageSimulation(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/outageSimulation", groupID, clusterName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
  "clusterName": "%s",
  "groupId": "%s",
  "id": "63599c987d469e386156afcb",
  "outageFilters": [
    {
      "cloudProvider": "AWS",
      "regionName": "string",
      "type": "REGION"
    }
  ],
  "startRequestDate": "2022-01-01T00:00:00Z",
  "state": "START_REQUESTED"
}`, clusterName, groupID)
	})

	simulation, _, err := client.ClusterOutageSimulation.GetOutageSimulation(ctx, groupID, clusterName)
	if err != nil {
		t.Fatalf("ClusterOutageSimulation.GetOutageSimulation returned error: %v", err)
	}

	expected := &ClusterOutageSimulation{
		ClusterName: pointy.String(clusterName),
		GroupID:     pointy.String(groupID),
		ID:          pointy.String("63599c987d469e386156afcb"),
		OutageFilters: []ClusterOutageSimulationOutageFilter{
			{
				CloudProvider: pointy.String("AWS"),
				RegionName:    pointy.String("string"),
				Type:          pointy.String("REGION"),
			},
		},
		StartRequestDate: pointy.String("2022-01-01T00:00:00Z"),
		State:            pointy.String("START_REQUESTED"),
	}

	if diff := deep.Equal(simulation, expected); diff != nil {
		t.Error(diff)
	}
}

func TestClusterOutageSimulationServiceOp_EndOutageSimulation(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/outageSimulation", groupID, clusterName)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprintf(w, `{
  "clusterName": "%s",
  "groupId": "%s",
  "id": "63599c987d469e386156afcb",
  "outageFilters": [
    {
      "cloudProvider": "AWS",
      "regionName": "string",
      "type": "REGION"
    }
  ],
  "startRequestDate": "2022-01-01T00:00:00Z",
  "state": "START_REQUESTED"
}`, clusterName, groupID)
	})

	simulation, _, err := client.ClusterOutageSimulation.EndOutageSimulation(ctx, groupID, clusterName)
	if err != nil {
		t.Fatalf("ClusterOutageSimulation.EndOutageSimulation returned error: %v", err)
	}

	expected := &ClusterOutageSimulation{
		ClusterName: pointy.String(clusterName),
		GroupID:     pointy.String(groupID),
		ID:          pointy.String("63599c987d469e386156afcb"),
		OutageFilters: []ClusterOutageSimulationOutageFilter{
			{
				CloudProvider: pointy.String("AWS"),
				RegionName:    pointy.String("string"),
				Type:          pointy.String("REGION"),
			},
		},
		StartRequestDate: pointy.String("2022-01-01T00:00:00Z"),
		State:            pointy.String("START_REQUESTED"),
	}

	if diff := deep.Equal(simulation, expected); diff != nil {
		t.Error(diff)
	}
}

func TestClusterOutageSimulationServiceOp_StartOutageSimulation(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	path := fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/outageSimulation", groupID, clusterName)

	createRequest := &ClusterOutageSimulationRequest{
		OutageFilters: []ClusterOutageSimulationOutageFilter{
			{
				CloudProvider: pointy.String("AWS"),
				RegionName:    pointy.String("string"),
				Type:          pointy.String("REGION"),
			},
		},
	}

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		expected := map[string]interface{}{
			"outageFilters": []interface{}{
				map[string]interface{}{
					"cloudProvider": "AWS",
					"regionName":    "string",
					"type":          "REGION",
				},
			},
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err.Error())
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}
		fmt.Fprintf(w, `{
  "clusterName": "%s",
  "groupId": "%s",
  "id": "63599c987d469e386156afcb",
  "outageFilters": [
    {
      "cloudProvider": "AWS",
      "regionName": "string",
      "type": "REGION"
    }
  ],
  "startRequestDate": "2022-01-01T00:00:00Z",
  "state": "START_REQUESTED"
}`, clusterName, groupID)
	})

	simulation, _, err := client.ClusterOutageSimulation.StartOutageSimulation(ctx, groupID, clusterName, createRequest)
	if err != nil {
		t.Fatalf("ClusterOutageSimulation.StartOutageSimulation returned error: %v", err)
	}

	expected := &ClusterOutageSimulation{
		ClusterName: pointy.String(clusterName),
		GroupID:     pointy.String(groupID),
		ID:          pointy.String("63599c987d469e386156afcb"),
		OutageFilters: []ClusterOutageSimulationOutageFilter{
			{
				CloudProvider: pointy.String("AWS"),
				RegionName:    pointy.String("string"),
				Type:          pointy.String("REGION"),
			},
		},
		StartRequestDate: pointy.String("2022-01-01T00:00:00Z"),
		State:            pointy.String("START_REQUESTED"),
	}

	if diff := deep.Equal(simulation, expected); diff != nil {
		t.Error(diff)
	}
}
