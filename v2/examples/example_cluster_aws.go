package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"context"

	"github.com/mongodb-forks/digest"
	client "go.mongodb.org/atlas/mongodbatlas/v2"
	apiv2latest "go.mongodb.org/atlas/mongodbatlas/v2/api"

	utils "go.mongodb.org/atlas/mongodbatlas/v2/test/generators"
)

/*
* MongoDB Atlas Go SDK Example
*
* Example performs following actions:
* - Finds first project in the user organization
* - Creates new MongoDB cluster backed by AWS provider in that project
* - Creates database user for connection
* - Creates IP access to enable connection
 */
func main() {
	ctx := context.Background()
	// Values provided as part of env variables
	// See: https://www.mongodb.com/docs/atlas/app-services/authentication/api-key/
	apiKey := os.Getenv("MDB_API_KEY")
	apiSecret := os.Getenv("MDB_API_SECRET")

	transport := digest.NewTransport(apiKey, apiSecret)
	httpClient, err := transport.Client()
	if err != nil {
		log.Fatalf(err.Error())
	}
	sdk := client.NewClientWithUrl(httpClient, "https://cloud-dev.mongodb.com")
	sdk.GetConfig().Debug = true
	// -- 1. Get first project
	projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
	handleErr(err, response)

	fmt.Println(projects)

	if projects.GetTotalCount() == 0 {
		panic("account should have at least single project")
	}

	projectId := projects.GetResults()[0].GetId()

	// -- 2. Create Cluster
	cluster := createClusterRequest(projectId)
	createdCluster, resp, err := sdk.MultiCloudClustersApi.CreateCluster(ctx, projectId).
		ClusterDescriptionV15(*cluster).Execute()
	handleErr(err, resp)

	// -- 3. Create Database User
	user := createDatabaseUserRequest(sdk, projectId)
	if user != nil {
		_, resp, err = sdk.DatabaseUsersApi.
			CreateDatabaseUser(ctx, projectId).
			DatabaseUser(*user).
			Execute()
		handleErr(err, resp)
	}

	// -- 4. Enable IP access
	ipAddress := utils.GetIpAddress()
	_, resp, err = sdk.ProjectIPAccessListApi.
		CreateProjectIpAccessList(context.Background(), projectId).
		NetworkPermissionEntry([]apiv2latest.NetworkPermissionEntry{{IpAddress: &ipAddress}}).
		Execute()
	handleErr(err, resp)

	// -- Print connection details to user
	fmt.Printf("SDK: Created new MongoDB cluster: %v \n", createdCluster.GetId())
	fmt.Println("Please wait up to 10 minutes for cluster to provision.")
}

func createDatabaseUserRequest(sdk *apiv2latest.APIClient, groupId string) *apiv2latest.DatabaseUser {
	username := "sdk-example"
	databaseName := "admin"

	currentUser, _, _ := sdk.DatabaseUsersApi.GetDatabaseUser(context.Background(),
		groupId, databaseName, username).Execute()

	if currentUser != nil {
		return nil
	}

	collectionName := "sdk-test"
	password, err := utils.GenerateRandomASCIIString(10)
	if err != nil {
		panic(err)
	}

	return &apiv2latest.DatabaseUser{
		Username:     username,
		Password:     &password,
		DatabaseName: databaseName,
		Roles: []apiv2latest.Role{
			{
				DatabaseName:   databaseName,
				CollectionName: collectionName,
				RoleName:       "read",
			},
		},
	}
}

func createClusterRequest(projectId string) *apiv2latest.ClusterDescriptionV15 {
	// Input arguments used for creation of the cluster
	clusterName, _ := utils.GenerateUniqueName("example-aws-cluster")

	// Location
	providerName := "AWS"
	clusterType := "SHARDED"
	regionName := "US_EAST_1"

	// Size
	numShards := int32(1)
	priority := int32(7)
	nodeCount := int32(3)
	instanceSize := "M10"

	return &apiv2latest.ClusterDescriptionV15{
		Name:        &clusterName,
		ClusterType: &clusterType,
		ReplicationSpecs: []apiv2latest.ReplicationSpec{
			{
				NumShards: &numShards,
				RegionConfigs: []apiv2latest.RegionConfig{
					{
						AWSRegionConfig: &apiv2latest.AWSRegionConfig{
							ProviderName: &providerName,
							Priority:     &priority,
							RegionName:   &regionName,
							ElectableSpecs: &apiv2latest.HardwareSpec{
								InstanceSize: &instanceSize,
								NodeCount:    &nodeCount,
							},
						},
					},
				},
			},
		},
	}

}

func handleErr(err error, resp *http.Response) {
	if err == nil {
		return
	}

	if resp != nil {
		fmt.Println(resp.Body)
	}
	log.Fatalf("Error when performing SDK request: %v", err.Error())

}
