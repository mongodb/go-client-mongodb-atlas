// Copyright 2026 MongoDB Inc
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

package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteCallbackPageRendersTemplate(t *testing.T) {
	tests := []struct {
		name           string
		status         int
		title          string
		message        string
		expectedAccent string
	}{
		{
			name:           "success page",
			status:         http.StatusOK,
			title:          "Connected to Atlas",
			message:        "The Atlas CLI is now connected.",
			expectedAccent: "#00ed64",
		},
		{
			name:           "error page",
			status:         http.StatusBadRequest,
			title:          "Connection failed",
			message:        "The connection attempt failed.",
			expectedAccent: "#970606",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			writeCallbackPage(rec, tt.status, tt.title, tt.message)

			if rec.Code != tt.status {
				t.Errorf("status = %d, want %d", rec.Code, tt.status)
			}
			body := rec.Body.String()
			for _, want := range []string{
				"<title>Atlas CLI — " + tt.title + "</title>",
				tt.expectedAccent,
				tt.message,
			} {
				if !strings.Contains(body, want) {
					t.Errorf("body missing %q", want)
				}
			}
			if strings.Contains(body, "%!") {
				t.Error("unconsumed format verb in body")
			}
		})
	}
}
