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

func TestSearch_ListIndexes(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "1"
	clusterName := "test"
	collectionName := "movies"
	databaseName := "sample_mflix"

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/fts/indexes/%s/%s", groupID, clusterName, databaseName, collectionName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"collectionName": "movies",
				"database": "sample_mflix",
				"indexID": "5d114a3587d9d65de99e7371",
				"mappings": {
					"dynamic": true
				},
				"name": "default",
				"synonyms": [
					{
						"analyzer": "lucene.english",
						"name": "mySynonyms",
						"source": {
							"collection": "synonyms"
						}
					}
				]
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
						},
						"title": [
							{
								"analyzer": "lucene.keyword",
								"searchAnalyzer": "lucene.keyword",
								"type": "string"
							},
							{
								"type": "autocomplete"
							}
						]
					}
				},
				"name": "SearchIndex1",
				"synonyms": [
					{
						"analyzer": "lucene.english",
						"name": "mySynonyms",
						"source": {
							"collection": "synonyms"
						}
					}
				]
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
			Synonyms: []map[string]interface{}{
				{
					"analyzer": "lucene.english",
					"name":     "mySynonyms",
					"source": map[string]interface{}{
						"collection": "synonyms",
					},
				},
			},
		},
		{
			CollectionName: "movies",
			Database:       "sample_mflix",
			IndexID:        "5d1268a980eef518dac0cf41",
			Mappings: &IndexMapping{
				Dynamic: false,
				Fields: &map[string]interface{}{
					"genres": map[string]interface{}{
						"analyzer": "lucene.standard",
						"type":     "string",
					},
					"plot": map[string]interface{}{
						"analyzer": "lucene.standard",
						"type":     "string",
					},
					"title": []interface{}{
						map[string]interface{}{
							"analyzer":       "lucene.keyword",
							"searchAnalyzer": "lucene.keyword",
							"type":           "string",
						},
						map[string]interface{}{
							"type": "autocomplete",
						},
					},
				},
			},
			Name: "SearchIndex1",
			Synonyms: []map[string]interface{}{
				{
					"analyzer": "lucene.english",
					"name":     "mySynonyms",
					"source": map[string]interface{}{
						"collection": "synonyms",
					},
				},
			},
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

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/fts/indexes/%s", projectID, clusterName, indexID), func(w http.ResponseWriter, r *http.Request) {
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
			"name": "SearchIndex1",
			"synonyms": [
				{
					"analyzer": "lucene.english",
					"name": "mySynonyms",
					"source": {
						"collection": "synonyms"
					}
				}
			]
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
			Fields: &map[string]interface{}{
				"genres": map[string]interface{}{
					"analyzer": "lucene.standard",
					"type":     "string",
				},
				"plot": map[string]interface{}{
					"analyzer": "lucene.standard",
					"type":     "string",
				},
			},
		},
		Name: "SearchIndex1",
		Synonyms: []map[string]interface{}{
			{
				"analyzer": "lucene.english",
				"name":     "mySynonyms",
				"source": map[string]interface{}{
					"collection": "synonyms",
				},
			},
		},
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

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/fts/indexes", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
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
			"name" : "default",
			"synonyms": [
				{
					"analyzer": "lucene.english",
					"name": "mySynonyms",
					"source": {
						"collection": "synonyms"
					}
				}
			]
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
		Synonyms: []map[string]interface{}{
			{
				"analyzer": "lucene.english",
				"name":     "mySynonyms",
				"source": map[string]interface{}{
					"collection": "synonyms",
				},
			},
		},
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

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/fts/indexes/%s", groupID, clusterName, indexID), func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/fts/indexes/%s", projectID, clusterName, indexID), func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/fts/analyzers", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
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
			MaxTokenLength: pointy.Int(32),
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

func TestSearch_UpdateAllAnalyzers(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	groupID := "1"
	clusterName := "test"
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/clusters/%s/fts/analyzers", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `[
			{
				"baseAnalyzer" : "lucene.standard",
				"maxTokenLength" : 32,
				"name" : "my_new_analyzer",
				"ignoreCase": true
			},
			{
				"baseAnalyzer" : "lucene.english",
				"name" : "my_new_analyzer2",
				"stopwords" : [ "first", "second", "third", "etc" ]
			}
		]`)
	})

	request := []*SearchAnalyzer{
		{
			BaseAnalyzer:   "lucene.standard",
			MaxTokenLength: pointy.Int(32),
			Name:           "my_new_analyzer",
			IgnoreCase:     pointy.Bool(true),
		},
		{
			BaseAnalyzer: "lucene.english",
			Name:         "my_new_analyzer2",
			Stopwords:    []string{"first", "second", "third", "etc"},
		},
	}

	indexes, _, err := client.Search.UpdateAllAnalyzers(ctx, groupID, clusterName, request)
	if err != nil {
		t.Fatalf("Search.GetAllAnalyzers returned error: %v", err)
	}

	expected := []*SearchAnalyzer{
		{
			BaseAnalyzer:   "lucene.standard",
			MaxTokenLength: pointy.Int(32),
			Name:           "my_new_analyzer",
			IgnoreCase:     pointy.Bool(true),
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
