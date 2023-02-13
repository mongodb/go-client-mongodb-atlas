package regression

import (
	mongodbatlas "go.mongodb.org/atlas/api/v1alpha"
)

func TestProjectIPAcessListModelRegression() {
	_ = mongodbatlas.PaginatedNetworkAccessView{
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
