package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDefaultMongoDBMajorVersionServiceOp_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/private/unauth/nds/defaultMongoDBMajorVersion", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, "4.2")
	})

	result, _, err := client.DefaultMongoDBMajorVersion.Get(ctx)
	if err != nil {
		t.Fatalf("DefaultMongoDBMajorVersion.Get returned error: %v", err)
	}
	expected := "4.2"
	if result != expected {
		t.Errorf("Expected %s, Got %s", expected, result)
	}
}
