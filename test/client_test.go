package test

import (
	"testing"

	mongodbatlas "go.mongodb.org/atlas/mongodbatlasv2"
)

func Test_Client_Config(t *testing.T) {
	apiKey := "test"
	apiSecret := "test"
	host := "https://cloud.mongodb.com"

	sdk, err := mongodbatlas.NewClient(
		mongodbatlas.UseDigestAuth(apiKey, apiSecret),
		mongodbatlas.UseBaseURL(host),
		mongodbatlas.UseDebug(false))

	if err != nil {
		t.Error(err)
	}

	serverURL := sdk.GetConfig().Servers[0].URL

	if serverURL != host {
		t.Errorf("Expected %s, got %s", host, serverURL)
	}
}
