package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestIndexesServiceOp_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "1"
	clusterName := "appData"

	createRequest := &IndexConfiguration{
		DB:         "test",
		Collection: "test",
		Keys: []map[string]string{
			{
				"name": "1",
			},
		},
		Options: &IndexOptions{
			Unique: true,
		},
	}
	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s/index", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		expected := map[string]interface{}{
			"db":         "test",
			"collection": "test",
			"keys": []interface{}{map[string]interface{}{
				"name": "1",
			}},
			"options": map[string]interface{}{
				"unique": true,
			},
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Clusters.Update Request Body = %v", diff)
		}
		fmt.Fprint(w, "{}")
	})

	_, err := client.Indexes.Create(ctx, groupID, clusterName, createRequest)

	if err != nil {
		t.Fatalf("Indexes.Create returned error: %v", err)
	}
}
