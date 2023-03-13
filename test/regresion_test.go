package test

import (
	"testing"

	"go.mongodb.org/atlas/api/test/regression"
)

func TestIntegrationSDK(t *testing.T) {
	regression.TestProjectIPAcessListModelRegression()
	regression.TestMultiCloudClustersModelRegression()
	regression.TestDatabaseUsersModelRegression()
	t.Log("Compilation only")
}
