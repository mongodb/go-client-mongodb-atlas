// Copyright 2022 MongoDB Inc
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

func TestProjects_GetProjectSettings(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/settings", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
               "isCollectDatabaseSpecificsStatisticsEnabled": true,
			   "isDataExplorerEnabled": true,
			   "isPerformanceAdvisorEnabled": true,
			   "isRealtimePerformancePanelEnabled": true,
			   "isSchemaAdvisorEnabled": true
}`)
	})

	projectSettings, _, err := client.Projects.GetProjectSettings(ctx, groupID)
	if err != nil {
		t.Fatalf("Projects.GetProjectSettings returned error: %v", err)
	}

	expected := &ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: pointy.Bool(true),
		IsDataExplorerEnabled:                       pointy.Bool(true),
		IsPerformanceAdvisorEnabled:                 pointy.Bool(true),
		IsRealtimePerformancePanelEnabled:           pointy.Bool(true),
		IsSchemaAdvisorEnabled:                      pointy.Bool(true),
	}

	if diff := deep.Equal(projectSettings, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProjects_UpdateProjectSettings(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/settings", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		_, _ = fmt.Fprint(w, `{
               "isCollectDatabaseSpecificsStatisticsEnabled": true,
			   "isDataExplorerEnabled": true,
			   "isPerformanceAdvisorEnabled": true,
			   "isRealtimePerformancePanelEnabled": true,
			   "isSchemaAdvisorEnabled": true
}`)
	})

	body := &ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: pointy.Bool(true),
		IsDataExplorerEnabled:                       pointy.Bool(true),
		IsPerformanceAdvisorEnabled:                 pointy.Bool(true),
		IsRealtimePerformancePanelEnabled:           pointy.Bool(true),
		IsSchemaAdvisorEnabled:                      pointy.Bool(true),
	}

	response, _, err := client.Projects.UpdateProjectSettings(ctx, groupID, body)
	if err != nil {
		t.Fatalf("Projects.UpdateProjectSettings returned error: %v", err)
	}

	expected := &ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: pointy.Bool(true),
		IsDataExplorerEnabled:                       pointy.Bool(true),
		IsPerformanceAdvisorEnabled:                 pointy.Bool(true),
		IsRealtimePerformancePanelEnabled:           pointy.Bool(true),
		IsSchemaAdvisorEnabled:                      pointy.Bool(true),
	}

	if diff := deep.Equal(response, expected); diff != nil {
		t.Error(diff)
	}
}
