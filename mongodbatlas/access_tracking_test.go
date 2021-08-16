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
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestAccessTracking_ListByCluster(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/dbAccessHistory/clusters/%s", groupID, clusterName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			  "accessLogs": [
					{
					  "authResult": true,
					  "authSource": "admin",
					  "groupId": "1",
					  "hostname": "cluster0-shard-00-00-c01ab.mongodb.net:27017",
					  "clusterName": "globalCluster",
					  "ipAddress": "123.45.0.1",
					  "logLine": "2019-07-25T19:14:42.484+0000 I  ACCESS   [conn2167] Successfully authenticated as principal jon-snow on admin from client 123.45.0.1:8080",
					  "timestamp": "Sun Jul 25 9:14:42 EDT 2019",
					  "username": "jon-snow"
					},
					{
					  "authResult": true,
					  "authSource": "admin",
					  "failureReason": "UserNotFound: Could not find user \"jane-doe\" for db \"admin\"",
					  "groupId": "1",
					  "hostname": "cluster0-shard-00-00-c01ab.mongodb.net:27017",
					  "clusterName": "globalCluster",
					  "ipAddress": "123.45.2.2",
					  "logLine": "2019-07-25T19:13:39.316+0000 I  ACCESS   [conn1893] SASL SCRAM-SHA-1 authentication failed for jane-doe on admin from client 123.45.2.2:51842 ; UserNotFound: Could not find user \"jane-doe\" for db \"admin\"",
					  "timestamp": "Sun Jul 25 9:14:42 EDT 2019",
					  "username": "jane-doe"
					}]
		}`)
	})
	authResult := true
	opts := &AccessLogOptions{
		Start:      "1564064082000",
		End:        "1564064082000",
		NLogs:      1,
		IPAddress:  "123.45.2.2",
		AuthResult: &authResult,
	}

	results, _, err := client.AccessTracking.ListByCluster(ctx, groupID, clusterName, opts)
	if err != nil {
		t.Fatalf("AccessTracking.ListByCluster returned error: %v", err)
	}

	expected := &AccessLogSettings{
		AccessLogs: []*AccessLogs{
			{
				GroupID:     "1",
				Hostname:    "cluster0-shard-00-00-c01ab.mongodb.net:27017",
				ClusterName: "globalCluster",
				IPAddress:   "123.45.0.1",
				LogLine:     "2019-07-25T19:14:42.484+0000 I  ACCESS   [conn2167] Successfully authenticated as principal jon-snow on admin from client 123.45.0.1:8080",
				Timestamp:   "Sun Jul 25 9:14:42 EDT 2019",
				Username:    "jon-snow",
				AuthSource:  "admin",
				AuthResult:  &authResult,
			},
			{
				GroupID:       "1",
				Hostname:      "cluster0-shard-00-00-c01ab.mongodb.net:27017",
				ClusterName:   "globalCluster",
				IPAddress:     "123.45.2.2",
				LogLine:       "2019-07-25T19:13:39.316+0000 I  ACCESS   [conn1893] SASL SCRAM-SHA-1 authentication failed for jane-doe on admin from client 123.45.2.2:51842 ; UserNotFound: Could not find user \"jane-doe\" for db \"admin\"",
				Timestamp:     "Sun Jul 25 9:14:42 EDT 2019",
				Username:      "jane-doe",
				AuthSource:    "admin",
				AuthResult:    &authResult,
				FailureReason: "UserNotFound: Could not find user \"jane-doe\" for db \"admin\"",
			},
		},
	}

	if diff := deep.Equal(results, expected); diff != nil {
		t.Error(diff)
	}
}

func TestAccessTracking_ListByHostname(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	hostname := "cluster0-shard-00-00-c01ab.mongodb.net:27017"
	mux.HandleFunc(fmt.Sprintf("/api/atlas/v1.0/groups/%s/dbAccessHistory/processes/%s", groupID, hostname), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			  "accessLogs": [
					{
					  "authResult": true,
					  "authSource": "admin",
					  "groupId": "1",
					  "hostname": "cluster0-shard-00-00-c01ab.mongodb.net:27017",
					  "clusterName": "globalCluster",
					  "ipAddress": "123.45.0.1",
					  "logLine": "2019-07-25T19:14:42.484+0000 I  ACCESS   [conn2167] Successfully authenticated as principal jon-snow on admin from client 123.45.0.1:8080",
					  "timestamp": "Sun Jul 25 9:14:42 EDT 2019",
					  "username": "jon-snow"
					},
					{
					  "authResult": true,
					  "authSource": "admin",
					  "failureReason": "UserNotFound: Could not find user \"jane-doe\" for db \"admin\"",
					  "groupId": "1",
					  "hostname": "cluster0-shard-00-00-c01ab.mongodb.net:27017",
					  "clusterName": "globalCluster",
					  "ipAddress": "123.45.2.2",
					  "logLine": "2019-07-25T19:13:39.316+0000 I  ACCESS   [conn1893] SASL SCRAM-SHA-1 authentication failed for jane-doe on admin from client 123.45.2.2:51842 ; UserNotFound: Could not find user \"jane-doe\" for db \"admin\"",
					  "timestamp": "Sun Jul 25 9:14:42 EDT 2019",
					  "username": "jane-doe"
					}]
		}`)
	})

	authResult := true
	opts := &AccessLogOptions{
		Start:      "1564064082000",
		End:        "1564064082000",
		NLogs:      1,
		IPAddress:  "123.45.2.2",
		AuthResult: &authResult,
	}

	results, _, err := client.AccessTracking.ListByHostname(ctx, groupID, hostname, opts)
	if err != nil {
		t.Fatalf("AccessTracking.ListByHostname returned error: %v", err)
	}

	expected := &AccessLogSettings{
		AccessLogs: []*AccessLogs{
			{
				GroupID:     "1",
				Hostname:    "cluster0-shard-00-00-c01ab.mongodb.net:27017",
				ClusterName: "globalCluster",
				IPAddress:   "123.45.0.1",
				LogLine:     "2019-07-25T19:14:42.484+0000 I  ACCESS   [conn2167] Successfully authenticated as principal jon-snow on admin from client 123.45.0.1:8080",
				Timestamp:   "Sun Jul 25 9:14:42 EDT 2019",
				Username:    "jon-snow",
				AuthSource:  "admin",
				AuthResult:  &authResult,
			},
			{
				GroupID:       "1",
				Hostname:      "cluster0-shard-00-00-c01ab.mongodb.net:27017",
				ClusterName:   "globalCluster",
				IPAddress:     "123.45.2.2",
				LogLine:       "2019-07-25T19:13:39.316+0000 I  ACCESS   [conn1893] SASL SCRAM-SHA-1 authentication failed for jane-doe on admin from client 123.45.2.2:51842 ; UserNotFound: Could not find user \"jane-doe\" for db \"admin\"",
				Timestamp:     "Sun Jul 25 9:14:42 EDT 2019",
				Username:      "jane-doe",
				AuthSource:    "admin",
				AuthResult:    &authResult,
				FailureReason: "UserNotFound: Could not find user \"jane-doe\" for db \"admin\"",
			},
		},
	}

	if diff := deep.Equal(results, expected); diff != nil {
		t.Error(diff)
	}
}
