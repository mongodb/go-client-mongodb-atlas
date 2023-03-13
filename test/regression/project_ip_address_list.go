package regression

import (
	mongodbatlas "go.mongodb.org/atlas/mongodbatlasv2"
)

func TestProjectIPAcessListModelRegression() {
	_ = mongodbatlas.PaginatedNetworkAccess{
		TotalCount: nil,
		Results:    []mongodbatlas.NetworkPermissionEntry{},
		Links: []mongodbatlas.Link{
			{
				Href: nil,
				Rel:  nil,
			},
		},
	}

	_ = mongodbatlas.NetworkPermissionEntry{
		AwsSecurityGroup: nil,
		CidrBlock:        nil,
		Comment:          nil,
		DeleteAfterDate:  nil,
		GroupId:          nil,
		IpAddress:        nil,
		Links:            []mongodbatlas.Link{},
	}
}
