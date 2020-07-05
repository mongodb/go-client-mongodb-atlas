package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestProjectIPAllowlist_ListProjectIPAllowlist(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/groups/1/whitelist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"results": [{"groupId":"1", "cidrBlock":"0.0.0.0/12", "ipAddress":"0.0.0.0"},{"groupId":"1", "cidrBlock":"0.0.0.1/12", "ipAddress":"0.0.0.1"}], "totalCount":2}`)
	})

	projectIPAllowlists, _, err := client.ProjectIPAllowlist.List(ctx, "1", nil)
	if err != nil {
		t.Fatalf("ProjectIPAllowlist.List returned error: %v", err)
	}

	expected := []ProjectIPAllowlist{{GroupID: "1", CIDRBlock: "0.0.0.0/12", IPAddress: "0.0.0.0"}, {GroupID: "1", CIDRBlock: "0.0.0.1/12", IPAddress: "0.0.0.1"}}
	if diff := deep.Equal(projectIPAllowlists, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProjectIPAllowlist_ListProjectIPAllowlistMultiplePages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/groups/1/whitelist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		dr := projectIPAllowlistsResponse{
			Results: []ProjectIPAllowlist{
				{GroupID: "1", CIDRBlock: "0.0.0.1/12"},
				{GroupID: "1", CIDRBlock: "0.0.0.0/12"},
			},
			Links: []*Link{
				{Href: "http://example.com/api/atlas/v1.0/groups/1/whitelist?pageNum=2&itemsPerPage=2", Rel: "self"},
				{Href: "http://example.com/api/atlas/v1.0/groups/1/whitelist?pageNum=1&itemsPerPage=2", Rel: "previous"},
			},
		}

		b, err := json.Marshal(dr)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprint(w, string(b))
	})

	_, resp, err := client.ProjectIPAllowlist.List(ctx, "1", nil)
	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 2)
}

func TestProjectIPAllowlist_RetrievePageByNumber(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	jBlob := `
	{
		"links": [
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/whitelist?pageNum=1&itemsPerPage=1",
				"rel": "previous"
			},
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/whitelist?pageNum=2&itemsPerPage=1",
				"rel": "self"
			},
			{
				"href": "http://example.com/api/atlas/v1.0/groups/1/whitelist?itemsPerPage=3&pageNum=2",
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
						"href": "http://example.com/api/atlas/v1.0/groups/1/whitelist/0.0.0.0",
						"rel": "self"
					}
				]
			}
		],
		"totalCount": 3
	}`

	mux.HandleFunc("/groups/1/whitelist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, jBlob)
	})

	opt := &ListOptions{PageNum: 2}
	_, resp, err := client.ProjectIPAllowlist.List(ctx, "1", opt)
	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 2)
}

func TestProjectIPAllowlist_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"

	createRequest := []*ProjectIPAllowlist{
		{
			GroupID:   groupID,
			CIDRBlock: "0.0.0.1/12",
		},
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/whitelist", groupID), func(w http.ResponseWriter, r *http.Request) {
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

	projectIPAllowlist, _, err := client.ProjectIPAllowlist.Create(ctx, groupID, createRequest)
	if err != nil {
		t.Fatalf("ProjectIPAllowlist.Create returned error: %v", err)
	}

	if len(projectIPAllowlist) > 1 {
		t.Error("expected projectIPAllowlist response a list of one element")
	}

	if cidrBlock := projectIPAllowlist[0].CIDRBlock; cidrBlock != "0.0.0.1/12" {
		t.Errorf("expected cidrBlock '%s', received '%s'", "0.0.0.1/12", cidrBlock)
	}

	if id := projectIPAllowlist[0].GroupID; id != groupID {
		t.Errorf("expected groupId '%s', received '%s'", groupID, id)
	}
}

func TestProjectIPAllowlist_GetProjectIPAllowlist(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	ipAddress := "0.0.0.0"

	mux.HandleFunc(fmt.Sprintf("/groups/1/whitelist/%s", ipAddress), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"ipAddress":"%s"}`, ipAddress)
	})

	projectIPAllowlists, _, err := client.ProjectIPAllowlist.Get(ctx, "1", ipAddress)
	if err != nil {
		t.Fatalf("ProjectIPAllowlist.Get returned error: %v", err)
	}

	expected := &ProjectIPAllowlist{IPAddress: ipAddress}

	if diff := deep.Equal(projectIPAllowlists, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProjectIPAllowlist_GetProjectIPAllowlistByCIDR(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	cidr := "0.0.0.0/32"

	mux.HandleFunc(fmt.Sprintf("/groups/1/whitelist/%s", cidr), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"cidrBlock":"%s"}`, cidr)
	})

	projectIPAllowlists, _, err := client.ProjectIPAllowlist.Get(ctx, "1", cidr)
	if err != nil {
		t.Fatalf("ProjectIPAllowlist.Get returned error: %v", err)
	}

	expected := &ProjectIPAllowlist{CIDRBlock: cidr}
	if diff := deep.Equal(projectIPAllowlists, expected); diff != nil {
		t.Error(diff)
	}
}

func TestProjectIPAllowlist_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	ipAddress := "0.0.0.1"

	createRequest := []*ProjectIPAllowlist{{
		GroupID:   groupID,
		IPAddress: ipAddress,
	}}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/whitelist", groupID), func(w http.ResponseWriter, r *http.Request) {
		expected := []map[string]interface{}{{
			"ipAddress": ipAddress,
			"groupId":   groupID,
		}}

		jsonBlob := `
		{
			"results": [
				{
					"ipAddress": "0.0.0.1",
					"groupId": "1"
				}
			]
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

	projectIPAllowlist, _, err := client.ProjectIPAllowlist.Update(ctx, groupID, createRequest)
	if err != nil {
		t.Fatalf("ProjectIPAllowlist.Update returned error: %v", err)
	}

	if ip := projectIPAllowlist[0].IPAddress; ip != ipAddress {
		t.Errorf("expected ipAddress '%s', received '%s'", ipAddress, ipAddress)
	}

	if id := projectIPAllowlist[0].GroupID; id != groupID {
		t.Errorf("expected groupId '%s', received '%s'", groupID, id)
	}
}

func TestProjectIPAllowlist_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	ipAddress := "0.0.0.1"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/whitelist/%s", groupID, ipAddress), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.ProjectIPAllowlist.Delete(ctx, groupID, ipAddress)
	if err != nil {
		t.Fatalf("ProjectIPAllowlist.Delete returned error: %v", err)
	}
}
