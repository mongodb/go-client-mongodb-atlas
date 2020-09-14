package mongodbatlas

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestPerformanceAdvisor_GetNamespaces(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "ORG-ID"
	processName := "test:27017"

	mux.HandleFunc(fmt.Sprintf("/"+performanceAdvisorNamespacesPath, projectID, processName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				  "namespaces" : [ {
					"namespace" : "data.zips",
					"type" : "COLLECTION"
				  }, {
					"namespace" : "data.stocks",
					"type" : "COLLECTION"
				  } ]
		}`)
	})

	body := &NamespaceOptions{
		Since:    2,
		Duration: 2,
	}

	namespaces, _, err := client.PerformanceAdvisor.GetNamespaces(ctx, projectID, processName, body)
	if err != nil {
		t.Fatalf("PerformanceAdvisor.GetNamespaces returned error: %v", err)
	}

	expected := &Namespaces{
		Namespaces: []*Namespace{
			{
				Namespace: "data.zips",
				Type:      "COLLECTION",
			},
			{
				Namespace: "data.stocks",
				Type:      "COLLECTION",
			},
		},
	}

	if diff := deep.Equal(namespaces, expected); diff != nil {
		t.Error(diff)
	}
}

func TestPerformanceAdvisor_GetSlowQueries(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "ORG-ID"
	processName := "test:27017"

	mux.HandleFunc(fmt.Sprintf("/"+performanceAdvisorSlowQueryLogsPath, projectID, processName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			  "slowQueries" : [ {
				"line" : "2018-08-16T22:53:43.447+0000 I COMMAND  [conn10614] command myDb.users appName: \"MongoDB Shell\" command: find { find: \"users\", filter: { emails: \"tocde@fijoow.to\" }, lsid: { id: UUID(\"832b4b0e-085a-480e-b470-16a0994dc7cb\") }, $clusterTime: { clusterTime: Timestamp(1534460016, 1)...",
			   "namespace" : "myDb.users"
			  }, {
				"line" : "2018-08-16T22:54:32.705+0000 I COMMAND  [conn10614] command myDb.users appName: \"MongoDB Shell\" command: find { find: \"users\", filter: { emails: \"la@sa.kp\" }, lsid: { id: UUID(\"832b4b0e-085a-480e-b470-16a0994dc7cb\") }, $clusterTime: { clusterTime: Timestamp(1534460056, 1), ...",
				"namespace" : "myDb.users"
			  } ]
		}`)
	})

	body := &SlowQueryOptions{
		NamespaceOptions: NamespaceOptions{Since: 2, Duration: 2},
		Namespaces:       "test",
		NLogs:            2,
	}

	queries, _, err := client.PerformanceAdvisor.GetSlowQueries(ctx, projectID, processName, body)
	if err != nil {
		t.Fatalf("PerformanceAdvisor.GetSlowQueries returned error: %v", err)
	}

	expected := &SlowQueries{
		SlowQuery: []*SlowQuery{
			{
				Namespace: "myDb.users",
				Line:      "2018-08-16T22:53:43.447+0000 I COMMAND  [conn10614] command myDb.users appName: \"MongoDB Shell\" command: find { find: \"users\", filter: { emails: \"tocde@fijoow.to\" }, lsid: { id: UUID(\"832b4b0e-085a-480e-b470-16a0994dc7cb\") }, $clusterTime: { clusterTime: Timestamp(1534460016, 1)...",
			},
			{
				Namespace: "myDb.users",
				Line:      "2018-08-16T22:54:32.705+0000 I COMMAND  [conn10614] command myDb.users appName: \"MongoDB Shell\" command: find { find: \"users\", filter: { emails: \"la@sa.kp\" }, lsid: { id: UUID(\"832b4b0e-085a-480e-b470-16a0994dc7cb\") }, $clusterTime: { clusterTime: Timestamp(1534460056, 1), ...",
			},
		},
	}

	if diff := deep.Equal(queries, expected); diff != nil {
		t.Error(diff)
	}
}

func TestPerformanceAdvisor_GetSuggestedIndexes(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	projectID := "ORG-ID"
	processName := "test:27017"

	mux.HandleFunc(fmt.Sprintf("/"+performanceAdvisorSuggestedIndexesLogsPath, projectID, processName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
 "shapes" : [ {
			"avgMs" : 42,
			"count" : 2,
			"id" : "5b74689a80eef53f3388897e",
			"inefficiencyScore" : 50000,
			"namespace" : "test.users",
			"operations" : [ {
			  "predicates" : [{ "find" : { "emails" : "la@sa.kp" } }],
			  "raw" : "2018-08-15T17:14:11.115+0000 I COMMAND  [conn4576] command test.users appName: \"MongoDB Shell\" command: find { find: \"users\", filter: { emails: \"la@sa.kp\" }, lsid: { id: UUID(\"1a4e71d3-9b67-4e9c-b078-9fdf3fae9091\") }, $clusterTime: { clusterTime: Timestamp(1534353241, 1), signature: { hash: BinData(0, AB91938B7CF7BC87994A2909A98D87F29101EFA0), keyId: 6589681559618453505 } }, $db: \"test\" } planSummary: COLLSCAN keysExamined:0 docsExamined:50000 cursorExhausted:1 numYields:391 nreturned:1 reslen:339 locks:{ Global: { acquireCount: { r: 784 } }, Database: { acquireCount: { r: 392 } }, Collection: { acquireCount: { r: 392 } } } protocol:op_msg 34ms",
			  "stats" : {
				"ms" : 34,
				"nReturned" : 1,
				"nScanned" : 50000,
				"ts" : 1534353251147
			  }
			}, {
			  "predicates" : [{ "find" : { "emails" : "tocde@fijoow.to" } }],
			  "raw" : "2018-08-15T17:14:18.665+0000 I COMMAND  [conn4576] command test.users appName: \"MongoDB Shell\" command: find { find: \"users\", filter: { emails: \"tocde@fijoow.to\" }, lsid: { id: UUID(\"1a4e71d3-9b67-4e9c-b078-9fdf3fae9091\") }, $clusterTime: { clusterTime: Timestamp(1534353241, 1), signature: { hash: BinData(0, AB91938B7CF7BC87994A2909A98D87F29101EFA0), keyId: 6589681559618453505 } }, $db: \"test\" } planSummary: COLLSCAN keysExamined:0 docsExamined:50000 cursorExhausted:1 numYields:390 nreturned:1 reslen:342 locks:{ Global: { acquireCount: { r: 782 } }, Database: { acquireCount: { r: 391 } }, Collection: { acquireCount: { r: 391 } } } protocol:op_msg 36ms",
			  "stats" : {
				"ms" : 36,
				"nReturned" : 1,
				"nScanned" : 50000,
				"ts" : 1534353258697
			  }
			} ]
		  } ],
		  "suggestedIndexes" : [ {
			"id" : "5b74689a80eef53f3388897f",
			"impact" : [ "5b74689a80eef53f3388897e" ],
			"index" : [ {
			  "emails" : 1
			} ],
			"namespace" : "test.users",
			"weight" : 37.220480901815623
		  }, {
			"id" : "5b74689a80eef53f33888980",
			"impact" : [ "5b74689a80eef53f3388897d" ],
			"index" : [ {
			  "emails" : 1
			} ],
			"namespace" : "test.inventory",
			"weight" : 19.037578309966488
		  } ]
		}`)
	})

	body := &SuggestedIndexOptions{
		NamespaceOptions: NamespaceOptions{Since: 2, Duration: 2},
		Namespaces:       "test",
		NIndexes:         55,
		NExamples:        4,
	}

	indexes, _, err := client.PerformanceAdvisor.GetSuggestedIndexes(ctx, projectID, processName, body)
	if err != nil {
		t.Fatalf("PerformanceAdvisor.GetSuggestedIndexes returned error: %v", err)
	}

	expected := &SuggestedIndexes{
		SuggestedIndexes: []*SuggestedIndex{
			{
				ID:        "5b74689a80eef53f3388897f",
				Impact:    []string{"5b74689a80eef53f3388897e"},
				Namespace: "test.users",
				Weight:    37.220480901815623,
				Index: []map[string]int{
					{"emails": 1},
				},
			},
			{
				ID:        "5b74689a80eef53f33888980",
				Impact:    []string{"5b74689a80eef53f3388897d"},
				Namespace: "test.inventory",
				Weight:    19.037578309966488,
				Index: []map[string]int{
					{"emails": 1},
				},
			},
		},
		Shapes: []*Shape{
			{
				AvgMs:             42,
				Count:             2,
				ID:                "5b74689a80eef53f3388897e",
				InefficiencyScore: 50000,
				Namespace:         "test.users",
				Operations: []*Operation{
					{
						Raw: "2018-08-15T17:14:11.115+0000 I COMMAND  [conn4576] command test.users appName: \"MongoDB Shell\" command: find { find: \"users\", filter: { emails: \"la@sa.kp\" }, lsid: { id: UUID(\"1a4e71d3-9b67-4e9c-b078-9fdf3fae9091\") }, $clusterTime: { clusterTime: Timestamp(1534353241, 1), signature: { hash: BinData(0, AB91938B7CF7BC87994A2909A98D87F29101EFA0), keyId: 6589681559618453505 } }, $db: \"test\" } planSummary: COLLSCAN keysExamined:0 docsExamined:50000 cursorExhausted:1 numYields:391 nreturned:1 reslen:339 locks:{ Global: { acquireCount: { r: 784 } }, Database: { acquireCount: { r: 392 } }, Collection: { acquireCount: { r: 392 } } } protocol:op_msg 34ms",
						Stats: Stats{
							MS:        34,
							NReturned: 1,
							NScanned:  50000,
							TS:        1534353251147,
						},
						Predicates: []map[string]map[string]string{
							{"find": {"emails": "la@sa.kp"}},
						},
					},
					{
						Raw: "2018-08-15T17:14:18.665+0000 I COMMAND  [conn4576] command test.users appName: \"MongoDB Shell\" command: find { find: \"users\", filter: { emails: \"tocde@fijoow.to\" }, lsid: { id: UUID(\"1a4e71d3-9b67-4e9c-b078-9fdf3fae9091\") }, $clusterTime: { clusterTime: Timestamp(1534353241, 1), signature: { hash: BinData(0, AB91938B7CF7BC87994A2909A98D87F29101EFA0), keyId: 6589681559618453505 } }, $db: \"test\" } planSummary: COLLSCAN keysExamined:0 docsExamined:50000 cursorExhausted:1 numYields:390 nreturned:1 reslen:342 locks:{ Global: { acquireCount: { r: 782 } }, Database: { acquireCount: { r: 391 } }, Collection: { acquireCount: { r: 391 } } } protocol:op_msg 36ms",
						Stats: Stats{
							MS:        36,
							NReturned: 1,
							NScanned:  50000,
							TS:        1534353258697,
						},
						Predicates: []map[string]map[string]string{
							{"find": {"emails": "tocde@fijoow.to"}},
						},
					},
				},
			},
		},
	}

	if diff := deep.Equal(indexes, expected); diff != nil {
		t.Error(diff)
	}
}
