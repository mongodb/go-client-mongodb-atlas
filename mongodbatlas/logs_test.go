package mongodbatlas

import (
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
	})

	_, err := client.Logs.Get(ctx, groupID, cluster, log, nil, nil)
	if err != nil {
		t.Fatalf("DatabaseUser.Delete returned error: %v", err)
	}
}
