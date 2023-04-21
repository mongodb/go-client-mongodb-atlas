package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"context"

	mongodbatlas "go.mongodb.org/atlas/mongodbatlasv2"
	utils "go.mongodb.org/atlas/test/generators"
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

	sdk, err := mongodbatlas.NewClient(
		mongodbatlas.UseDigestAuth(apiKey, apiSecret),
		mongodbatlas.UseBaseURL("https://cloud.mongodb.com"),
		mongodbatlas.UseDebug(false))
	handleErr(err, nil)

	// -- 1. Get first project using fluent API
	_, response, err := sdk.ProjectsApi.ListProjects(ctx).
		IncludeCount(false).ItemsPerPage(1).Execute()
	handleErr(err, response)

	// -- 1.1 Get first project using struct based API
	listParams := &mongodbatlas.ListProjectsApiParams{ItemsPerPage: mongodbatlas.PtrInt32(1)}
	projects, response, err := sdk.ProjectsApi.ListProjects(ctx).ExecuteWithParams(listParams)
	handleErr(err, response)

	if projects.GetTotalCount() == 0 {
		log.Fatal("account should have at least single project")
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
	ipAddress := getIpAddress()
	_, resp, err = sdk.ProjectIPAccessListApi.
		CreateProjectIpAccessList(context.Background(), projectId).
		NetworkPermissionEntry([]mongodbatlas.NetworkPermissionEntry{{IpAddress: &ipAddress}}).
		Execute()
	handleErr(err, resp)

	// -- Print connection details to user
	fmt.Printf("SDK: Created new MongoDB cluster: %v \n", createdCluster.GetId())
	fmt.Println("Please wait up to 10 minutes for cluster to provision.")
}

func createDatabaseUserRequest(sdk *mongodbatlas.APIClient, groupId string) *mongodbatlas.DatabaseUser {
	username := "sdk-example"
	databaseName := "admin"

	currentUser, _, _ := sdk.DatabaseUsersApi.GetDatabaseUser(context.Background(),
		groupId, databaseName, username).Execute()

	if currentUser != nil {
		return nil
	}

	collectionName := "sdk-test"
	password, err := utils.RandomASCIIString(10)
	if err != nil {
		log.Fatal(err)
	}

	return &mongodbatlas.DatabaseUser{
		Username:     username,
		Password:     &password,
		DatabaseName: databaseName,
		Roles: []mongodbatlas.Role{
			{
				DatabaseName:   databaseName,
				CollectionName: collectionName,
				RoleName:       "read",
			},
		},
	}
}

func createClusterRequest(projectId string) *mongodbatlas.ClusterDescriptionV15 {
	// Input arguments used for creation of the cluster
	clusterName, _ := utils.UniqueName("example-aws-cluster")

	// Location
	providerName := "AWS"
	clusterType := "SHARDED"
	regionName := "US_EAST_1"

	// Size
	numShards := int32(1)
	priority := int32(7)
	nodeCount := int32(3)
	instanceSize := "M10"

	return &mongodbatlas.ClusterDescriptionV15{
		Name:        &clusterName,
		ClusterType: &clusterType,
		ReplicationSpecs: []mongodbatlas.ReplicationSpec{
			{
				NumShards: &numShards,
				RegionConfigs: []mongodbatlas.RegionConfig{
					{
						AWSRegionConfig: &mongodbatlas.AWSRegionConfig{
							ProviderName: &providerName,
							Priority:     &priority,
							RegionName:   &regionName,
							ElectableSpecs: &mongodbatlas.HardwareSpec{
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
		// Printing generic message
		fmt.Println(err.Error())
	}
	apiErr, _ := mongodbatlas.AsError(err)
	log.Fatalf("Error when performing SDK request: %v", apiErr.GetDetail())

}

func getIpAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
