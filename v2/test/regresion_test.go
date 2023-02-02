package test

import (
	"testing"

	"go.mongodb.org/atlas/mongodbatlas/v2/test/regression"
)

func TestIntegrationSDK(t *testing.T) {
	regression.TestProjectIPAcessList_ModelRegression()
	regression.TestMultiCloudClusters_ModelRegression()
	regression.TestDatabaseUsers_ModelRegression()
}
