package test

import (
	"fmt"

	mongodbatlas "go.mongodb.org/atlas/mongodbatlasv2"
)

func ExampleNewClient() {
	apiKey := "test"
	apiSecret := "test"
	host := "https://cloud.mongodb.com"

	sdk, err := mongodbatlas.NewClient(
		mongodbatlas.UseDigestAuth(apiKey, apiSecret),
		mongodbatlas.UseBaseURL(host),
		mongodbatlas.UseDebug(false))

	if err != nil {
		panic(err)
	}

	serverURL := sdk.GetConfig().Servers[0].URL
	fmt.Println(serverURL)
	// Output:
	// https://cloud.mongodb.com
}
