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

func TestProjectIPAccessListServiceOp_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups/1/accessList", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"results": [{"groupId":"1", "cidrBlock":"0.0.0.0/12", "ipAddress":"0.0.0.0"},{"groupId":"1", "cidrBlock":"0.0.0.1/12", "ipAddress":"0.0.0.1"}], "totalCount":2}`)
	})

	projectIPWhitelists, _, err := client.ProjectIPAccessList.List(ctx, "1", nil)
	if err != nil {
		t.Fatalf("ProjectIPAccessList.List returned error: %v", err)
	}

	expected := &ProjectIPAccessLists{
		Results: []ProjectIPAccessList{
			{GroupID: "1", CIDRBlock: "0.0.0.0/12", IPAddress: "0.0.0.0"},
			{GroupID: "1", CIDRBlock: "0.0.0.1/12", IPAddress: "0.0.0.1"},
		},
		TotalCount: 2,
	}
	if diff := deep.Equal(projectIPWhitelists, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProjectIPAccessListServiceOp_List_WithPages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/atlas/v1.0/groups/1/accessList", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		dr := ProjectIPAccessLists{
			Results: []ProjectIPAccessList{
				{GroupID: "1", CIDRBlock: "0.0.0.1/12"},
				{GroupID: "1", CIDRBlock: "0.0.0.0/12"},
			},
			Links: []*Link{
				{Href: "http://example.com/api/atlas/v1.0/groups/1/accessList?pageNum=2&itemsPerPage=2", Rel: "self"},
				{Href: "http://example.com/api/atlas/v1.0/groups/1/accessList?pageNum=1&itemsPerPage=2", Rel: "previous"},
			},
		}

		b, err := json.Marshal(dr)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprint(w, string(b))
	})

	_, resp, err := client.ProjectIPAccessList.List(ctx, "1", nil)
	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 2)
}

func TestProjectIPAccessListServiceOp_List_ByPages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	jBlob := `
	{
		"links": [
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/accessList?pageNum=1&itemsPerPage=1",
				"rel": "previous"
			},
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/accessList?pageNum=2&itemsPerPage=1",
				"rel": "self"
			},
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/accessList?itemsPerPage=3&pageNum=2",
				"rel": "next"
			}
		],
		"results": [
			{
				"groupId":"1",
				"cidrBlock":"0.0.0.0/12",
				"ipAddress":"0.0.0.0",
				"links": [
					{
						"href": "http://example.com/api/atlas/v1.0/groups/1/accessList/0.0.0.0",
						"rel": "self"
					}
				]
			}
		],
		"totalCount": 3
	}`

	mux.HandleFunc("/api/atlas/v1.0/groups/1/accessList", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, jBlob)
	})

	opt := &ListOptions{PageNum: 2}
	_, resp, err := client.ProjectIPAccessList.List(ctx, "1", opt)
	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 2)
}

func TestProjectIPAccessListServiceOp_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	createRequest := []*ProjectIPAccessList{
		{
			GroupID:   groupID,
			CIDRBlock: "0.0.0.1/12",
		},
	}

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/accessList", groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := []map[string]interface{}{
			{
				"cidrBlock": "0.0.0.1/12",
				"groupId":   groupID,
			},
		}

		jsonBlob := `
		{
			"results": [{
				"groupId":"1",
				"cidrBlock":"0.0.0.1/12"
			}]
		}
		`

		var v []map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Error(diff)
		}

		fmt.Fprint(w, jsonBlob)
	})

	projectIPWhitelist, _, err := client.ProjectIPAccessList.Create(ctx, groupID, createRequest)
	if err != nil {
		t.Fatalf("ProjectIPAccessList.Create returned error: %v", err)
	}

	if len(projectIPWhitelist.Results) > 1 {
		t.Error("expected projectIPWhitelist response a list of one element")
	}

	if cidrBlock := projectIPWhitelist.Results[0].CIDRBlock; cidrBlock != "0.0.0.1/12" {
		t.Errorf("expected cidrBlock '%s', received '%s'", "0.0.0.1/12", cidrBlock)
	}

	if id := projectIPWhitelist.Results[0].GroupID; id != groupID {
		t.Errorf("expected groupId '%s', received '%s'", groupID, id)
	}
}

func TestProjectIPAccessListServiceOp_Get_byIP(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	ipAddress := "0.0.0.0"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/1/accessList/%s", ipAddress), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"ipAddress":"%s"}`, ipAddress)
	})

	projectIPWhitelists, _, err := client.ProjectIPAccessList.Get(ctx, "1", ipAddress)
	if err != nil {
		t.Fatalf("ProjectIPAccessList.Get returned error: %v", err)
	}

	expected := &ProjectIPAccessList{IPAddress: ipAddress}

	if diff := deep.Equal(projectIPWhitelists, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProjectIPAccessListServiceOp_Get_ByCIDR(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	cidr := "0.0.0.0/32"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/1/accessList/%s", cidr), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"cidrBlock":"%s"}`, cidr)
	})

	projectIPWhitelists, _, err := client.ProjectIPAccessList.Get(ctx, "1", cidr)
	if err != nil {
		t.Fatalf("ProjectIPAccessList.Get returned error: %v", err)
	}

	expected := &ProjectIPAccessList{CIDRBlock: cidr}
	if diff := deep.Equal(projectIPWhitelists, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProjectIPAccessListServiceOp_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	ipAddress := "0.0.0.1"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/accessList/%s", groupID, ipAddress), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.ProjectIPAccessList.Delete(ctx, groupID, ipAddress)
	if err != nil {
		t.Fatalf("ProjectIPAccessList.Delete returned error: %v", err)
	}
}
