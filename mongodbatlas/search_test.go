package mongodbatlas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/openlyinc/pointy"
)

func TestSearch_ListIndexes(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "1"
	clusterName := "test"
	collectionName := "movies"
	databaseName := "sample_mflix"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s/fts/indexes/%s/%s", groupID, clusterName, databaseName, collectionName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
		  {
			"collectionName": "movies",
			"database": "sample_mflix",
			"indexID": "5d114a3587d9d65de99e7371",
			"mappings": {
			  "dynamic": true
			},
			"name": "default"
		  },
		  {
			"collectionName": "movies",
			"database": "sample_mflix",
			"indexID": "5d1268a980eef518dac0cf41",
			"mappings": {
			  "dynamic": false,
			  "fields": {
				"genres" : {
				  "analyzer": "lucene.standard",
				  "type": "string"
				},
				"plot": {
				  "analyzer": "lucene.standard",
				  "type": "string"
				}
			  }
			},
		   "name": "SearchIndex1"
		  }
		]`)
	})

	indexes, _, err := client.Search.ListIndexes(ctx, groupID, clusterName, databaseName, collectionName, nil)
	if err != nil {
		t.Fatalf("Search.ListIndexes returned error: %v", err)
	}

	expected := []*SearchIndex{
		{
			CollectionName: "movies",
			Database:       "sample_mflix",
			IndexID:        "5d114a3587d9d65de99e7371",
			Mappings: &IndexMapping{
				Dynamic: true,
			},
			Name: "default",
		},
		{
			CollectionName: "movies",
			Database:       "sample_mflix",
			IndexID:        "5d1268a980eef518dac0cf41",
			Mappings: &IndexMapping{
				Dynamic: false,
				Fields: &map[string]IndexField{
					"genres": {
						Analyzer: "lucene.standard",
						Type:     "string",
					},
					"plot": {
						Analyzer: "lucene.standard",
						Type:     "string",
					},
				},
			},
			Name: "SearchIndex1",
		},
	}

	if diff := deep.Equal(indexes, expected); diff != nil {
		t.Error(diff)
	}
}

func TestSearch_GetIndex(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "5a0a1e7e0f2912c554080adc"
	clusterName := "test"
	indexID := "5d1268a980eef518dac0cf41"

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s/fts/indexes/%s", projectID, clusterName, indexID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
		  "collectionName": "movies",
		  "database": "sample_mflix",
		  "indexID": "5d1268a980eef518dac0cf41",
		  "mappings": {
			"dynamic": false,
			"fields": {
			  "genres": {
				"analyzer": "lucene.standard",
				"type": "string"
			  },
			  "plot": {
				"analyzer": "lucene.standard",
				"type": "string"
			  }
			}
		  },
		 "name": "SearchIndex1"
		}`)
	})

	index, _, err := client.Search.GetIndex(ctx, projectID, clusterName, indexID)
	if err != nil {
		t.Errorf("Search.GetIndex returned error: %v", err)
	}

	expected := &SearchIndex{
		CollectionName: "movies",
		Database:       "sample_mflix",
		IndexID:        "5d1268a980eef518dac0cf41",
		Mappings: &IndexMapping{
			Dynamic: false,
			Fields: &map[string]IndexField{
				"genres": {
					Analyzer: "lucene.standard",
					Type:     "string",
				},
				"plot": {
					Analyzer: "lucene.standard",
					Type:     "string",
				},
			},
		},
		Name: "SearchIndex1",
	}

	if diff := deep.Equal(index, expected); diff != nil {
		t.Error(diff)
	}
}

func TestSearchServiceOp_CreateIndex(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "5a0a1e7e0f2912c554080adc"
	clusterName := "test"

	createRequest := &SearchIndex{
		CollectionName: "orders",
		Database:       "fiscalYear2018",
		Mappings: &IndexMapping{
			Dynamic: true,
		},
		Name: "default",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s/fts/indexes", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"collectionName": "orders",
			"database":       "fiscalYear2018",
			"mappings": map[string]interface{}{
				"dynamic": true,
			},
			"name": "default",
		}

		var v map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			t.Fatalf("decode json: %v", err)
		}
		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Search.CreateIndex Request Body = %v", diff)
		}

		jsonBlob := `{
		  "collectionName" : "orders",
		  "database" : "fiscalYear2018",
		  "indexID" : "5d12990380eef5303341accd",
		  "mappings" : {
			"dynamic" : true
		  },
		  "name" : "default"
		}`
		fmt.Fprint(w, jsonBlob)
	})

	index, _, err := client.Search.CreateIndex(ctx, groupID, clusterName, createRequest)
	if err != nil {
		t.Fatalf("Search.CreateIndex returned error: %v", err)
	}

	expected := &SearchIndex{
		CollectionName: "orders",
		Database:       "fiscalYear2018",
		IndexID:        "5d12990380eef5303341accd",
		Mappings:       &IndexMapping{Dynamic: true},
		Name:           "default",
	}

	if diff := deep.Equal(index, expected); diff != nil {
		t.Error(diff)
	}
}

func TestSearchServiceOp_UpdateIndex(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "5a0a1e7e0f2912c554080adc"
	clusterName := "test"
	indexID := "5d129aef87d9d64f310bd79f"

	updateRequest := &SearchIndex{
		Analyzer:       "lucene.swedish",
		CollectionName: "orders",
		Database:       "fiscalYear2018",
		Mappings: &IndexMapping{
			Dynamic: true,
		},
		Name:           "2018ordersIndex",
		SearchAnalyzer: "lucene.whitespace",
	}

	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s/fts/indexes/%s", groupID, clusterName, indexID), func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"analyzer":       "lucene.swedish",
			"searchAnalyzer": "lucene.whitespace",
			"collectionName": "orders",
			"database":       "fiscalYear2018",
			"mappings": map[string]interface{}{
				"dynamic": true,
			},
			"name": "2018ordersIndex",
		}

		var v map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			t.Fatalf("decode json: %v", err)
		}
		if diff := deep.Equal(v, expected); diff != nil {
			t.Errorf("Search.UpdateIndex Request Body = %v", diff)
		}

		jsonBlob := `{
		  "analyzer" : "lucene.swedish",
		  "searchAnalyzer" : "lucene.whitespace",
		  "collectionName" : "orders",
		  "database" : "fiscalYear2018",
		  "indexID" : "5d129aef87d9d64f310bd79f",
		  "mappings" : {
			"dynamic" : true
		  },
		  "name" : "2018ordersIndex"
		}`
		fmt.Fprint(w, jsonBlob)
	})

	index, _, err := client.Search.UpdateIndex(ctx, groupID, clusterName, indexID, updateRequest)
	if err != nil {
		t.Fatalf("Search.UpdateIndex returned error: %v", err)
	}

	expected := &SearchIndex{
		Analyzer:       "lucene.swedish",
		CollectionName: "orders",
		Database:       "fiscalYear2018",
		IndexID:        "5d129aef87d9d64f310bd79f",
		Mappings:       &IndexMapping{Dynamic: true},
		Name:           "2018ordersIndex",
		SearchAnalyzer: "lucene.whitespace",
	}

	if diff := deep.Equal(index, expected); diff != nil {
		t.Error(diff)
	}
}

func TestSearchServiceOp_DeleteIndex(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "5a0a1e7e0f2912c554080adc"
	clusterName := "test"
	indexID := "5d1268a980eef518dac0cf41"
	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s/fts/indexes/%s", projectID, clusterName, indexID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Search.DeleteIndex(ctx, projectID, clusterName, indexID)
	if err != nil {
		t.Fatalf("Search.DeleteIndex returned error: %v", err)
	}
}

func TestSearch_ListAnalyzers(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "1"
	clusterName := "test"
	mux.HandleFunc(fmt.Sprintf("/groups/%s/clusters/%s/fts/analyzers", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
		  {
			"baseAnalyzer" : "lucene.standard",
			"maxTokenLength" : 32,
			"name" : "my_new_analyzer"
		  },
		  {
			"baseAnalyzer" : "lucene.english",
			"name" : "my_new_analyzer2",
			"stopwords" : [ "first", "second", "third", "etc" ]
		  }
		]`)
	})

	indexes, _, err := client.Search.ListAnalyzers(ctx, groupID, clusterName, nil)
	if err != nil {
		t.Fatalf("Search.ListAnalyzers returned error: %v", err)
	}

	expected := []*SearchAnalyzer{
		{
			BaseAnalyzer:   "lucene.standard",
			MaxTokenLength: pointy.Float64(32),
			Name:           "my_new_analyzer",
		},
		{
			BaseAnalyzer: "lucene.english",
			Name:         "my_new_analyzer2",
			Stopwords:    []string{"first", "second", "third", "etc"},
		},
	}

	if diff := deep.Equal(indexes, expected); diff != nil {
		t.Error(diff)
	}
}
