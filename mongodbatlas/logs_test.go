package mongodbatlas

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestLogs_Get(t *testing.T) {
	setup()
	defer teardown()

	groupID := "1"
	cluster := "test-username"
	log := "log"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s/logs/%s", groupID, cluster, log), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, "test")
	})

	buf := new(bytes.Buffer)
	_, err := client.Logs.Get(ctx, groupID, cluster, log, buf, nil)

	if buf.String() != "test" {
		t.Fatalf("Logs.Get returned error: %v", err)
	}
}
