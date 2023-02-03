package test

import (
	"testing"

	"go.mongodb.org/atlas/api/test/regression"
)

func TestIntegrationSDK(t *testing.T) {
	regression.TestProjectIPAcessList_ModelRegression()
	regression.TestMultiCloudClusters_ModelRegression()
	regression.TestDatabaseUsers_ModelRegression()
}
