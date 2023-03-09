package regression

import (
	mongodbatlas "go.mongodb.org/atlas/api/v1alpha"
)

func TestDatabaseUsersModelRegression() {
	_ = mongodbatlas.PaginatedApiAtlasDatabaseUser{
		TotalCount: nil,
		Results:    []mongodbatlas.DatabaseUser{},
		Links:      []mongodbatlas.Link{},
	}

	_ = mongodbatlas.DatabaseUser{
		AwsIAMType:      nil,
		DeleteAfterDate: nil,
		GroupId:         "",
		Username:        "",
		Password:        nil,
		DatabaseName:    "",
		LdapAuthType:    nil,
		X509Type:        nil,
		Scopes: []mongodbatlas.UserScope{
			{
				Name: "",
				Type: "",
			},
		},
		Roles: []mongodbatlas.Role{
			{
				CollectionName: "",
				DatabaseName:   "",
				RoleName:       "",
			},
		},
		Labels: []mongodbatlas.NDSLabel{
			{
				Key:   nil,
				Value: nil,
			},
		},
		Links: []mongodbatlas.Link{
			{
				Href: nil,
				Rel:  nil,
			},
		},
	}
}
