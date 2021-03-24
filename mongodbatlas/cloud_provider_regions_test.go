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
	"github.com/go-test/deep"
	"net/http"
	"testing"
)

func TestCloudProviderRegions_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	groupID := "5b6212af90dc76637950a2c6"

	path := fmt.Sprintf("/groups/%s/clusters/provider/regions", groupID)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
						   "links":[
							  {
								 "href":"http://mongodb-cloud.com/api/atlas/v1.0/groups/60585bf86e35cb1e7fbe739f/clusters/provider/regions",
								 "rel":"self"
							  }
						   ],
						   "results":[
							  {
								 "instanceSizes":[
									{
									   "availableRegions":[
										  {
											 "default":true,
											 "name":"US_EAST_1"
										  },
										  {
											 "default":false,
											 "name":"US_EAST_2"
										  }
									   ],
									   "name":"R50"
									},
									{
									   "availableRegions":[
										  {
											 "default":true,
											 "name":"US_EAST_1"
										  }
									   ],
									   "name":"M0"
									}
								 ],
								 "provider":"AWS"
							  },
							  {
								 "instanceSizes":[
									{
									   "availableRegions":[
										  {
											 "default":true,
											 "name":"US_WEST_2"
										  },
										  {
											 "default":false,
											 "name":"US_CENTRAL"
										  }
									   ],
									   "name":"R50"
									},
									{
									   "availableRegions":[
										  {
											 "default":true,
											 "name":"US_WEST_2"
										  }
									   ],
									   "name":"M0"
									}
								 ],
								 "provider":"AZURE"
							  },
							  {
								 "instanceSizes":[
									{
									   "availableRegions":[
										  {
											 "default":true,
											 "name":"CENTRAL_US"
										  },
										  {
											 "default":false,
											 "name":"EASTERN_US"
										  }
									   ],
									   "name":"R50"
									},
									{
									   "availableRegions":[
										  {
											 "default":true,
											 "name":"CENTRAL_US"
										  }
									   ],
									   "name":"M0"
									}
								 ],
								 "provider":"GCP"
							  }
						   ],
						   "totalCount":3
					}`)
	})

	availableRegions, _, err := client.CloudProviderRegions.Get(ctx, groupID, nil)
	if err != nil {
		t.Fatalf("CloudProviderRegions.Get returned error: %v", err)
	}

	expected := &CloudProviders{
		Links:      []*Link{
			{
				Rel:  "self",
				Href: "http://mongodb-cloud.com/api/atlas/v1.0/groups/60585bf86e35cb1e7fbe739f/clusters/provider/regions",
			},
		},
		Results:    []*CloudProvider{
			{
				Provider:      "AWS",
				InstanceSizes: []*InstanceSize{
					{
						Name:             "R50",
						AvailableRegions: []*AvailableRegion{
							{
								Name:    "US_EAST_1",
								Default: true,
							},
							{
								Name:    "US_EAST_2",
								Default: false,
							},
						},
					},
					{
						Name:             "M0",
						AvailableRegions: []*AvailableRegion{
							{
								Name:    "US_EAST_1",
								Default: true,
							},
						},
					},
				},
			},
			{
				Provider:      "AZURE",
				InstanceSizes: []*InstanceSize{
					{
						Name:             "R50",
						AvailableRegions: []*AvailableRegion{
							{
								Name:    "US_WEST_2",
								Default: true,
							},
							{
								Name:    "US_CENTRAL",
								Default: false,
							},
						},
					},
					{
						Name:             "M0",
						AvailableRegions: []*AvailableRegion{
							{
								Name:    "US_WEST_2",
								Default: true,
							},
						},
					},
				},
			},
			{
				Provider:      "GCP",
				InstanceSizes: []*InstanceSize{
					{
						Name:             "R50",
						AvailableRegions: []*AvailableRegion{
							{
								Name:    "CENTRAL_US",
								Default: true,
							},
							{
								Name:    "EASTERN_US",
								Default: false,
							},
						},
					},
					{
						Name:             "M0",
						AvailableRegions: []*AvailableRegion{
							{
								Name:    "CENTRAL_US",
								Default: true,
							},
						},
					},
				},
			},

		},
		TotalCount: 3,
	}
	
	if diff := deep.Equal(availableRegions, expected); diff != nil {
		t.Error(diff)
	}
}