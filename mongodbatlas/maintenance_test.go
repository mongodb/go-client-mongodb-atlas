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

func TestMaintenanceWindows_UpdateWithSheduleTime(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6d2065c687d9d64ae7acdg41"

	mux.HandleFunc(fmt.Sprintf("/"+maintenanceWindowsPath, groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"dayOfWeek":         float64(2),
			"hourOfDay":         float64(3),
			"numberOfDeferrals": float64(4),
			"startASAP":         false,
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, `{}`)
	})

	maintenanceRequest := &MaintenanceWindow{
		DayOfWeek:         2,
		HourOfDay:         pointy.Int(3),
		NumberOfDeferrals: 4,
		StartASAP:         pointy.Bool(false),
	}

	_, err := client.MaintenanceWindows.Update(ctx, groupID, maintenanceRequest)
	if err != nil {
		t.Fatalf("MaintenanceWindow.Update returned error: %v", err)
		return
	}
}

func TestMaintenanceWindows_UpdateWithStartNow(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6d2065c687d9d64ae7acdg41"

	mux.HandleFunc(fmt.Sprintf("/"+maintenanceWindowsPath, groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"startASAP": true,
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, `{}`)
	})

	maintenanceRequest := &MaintenanceWindow{
		StartASAP: pointy.Bool(true),
	}

	_, err := client.MaintenanceWindows.Update(ctx, groupID, maintenanceRequest)
	if err != nil {
		t.Errorf("MaintenanceWindow.Update returned error: %v", err)
		return
	}
}

func TestMaintenanceWindows_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6d2065c687d9d64ae7acdg41"

	mux.HandleFunc(fmt.Sprintf("/"+maintenanceWindowsPath, groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				"dayOfWeek": 2,
				"hourOfDay": 3,
				"numberOfDeferrals": 4
		}`)
	})

	maintenanceWindow, _, err := client.MaintenanceWindows.Get(ctx, groupID)
	if err != nil {
		t.Errorf("MaintenanceWindows.Get returned error: %v", err)
	}

	expected := &MaintenanceWindow{
		DayOfWeek:         2,
		HourOfDay:         pointy.Int(3),
		NumberOfDeferrals: 4,
	}

	if diff := deep.Equal(maintenanceWindow, expected); diff != nil {
		t.Error(diff)
	}
}

func TestMaintenanceWindows_Defer(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6d2065c687d9d64ae7acdg41"

	mux.HandleFunc(fmt.Sprintf("/"+maintenanceWindowsPath+"/defer", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{}`)
	})

	_, err := client.MaintenanceWindows.Defer(ctx, groupID)
	if err != nil {
		t.Errorf("MaintenanceWindows.Defer returned error: %v", err)
	}
}

func TestMaintenanceWindows_AutoDefer(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6d2065c687d9d64ae7acdg41"

	mux.HandleFunc(fmt.Sprintf("/"+maintenanceWindowsPath+"/autoDefer", groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{}`)
	})

	_, err := client.MaintenanceWindows.AutoDefer(ctx, groupID)
	if err != nil {
		t.Errorf("MaintenanceWindows.AutoDefer returned error: %v", err)
	}
}

func TestMaintenanceWindows_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "6d2065c687d9d64ae7acdg41"

	mux.HandleFunc(fmt.Sprintf("/"+maintenanceWindowsPath, groupID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{}`)
	})

	_, err := client.MaintenanceWindows.Reset(ctx, groupID)
	if err != nil {
		t.Errorf("MaintenanceWindows.Get returned error: %v", err)
	}
}
