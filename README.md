# go-client-mongodb-atlas [![Build Status](https://travis-ci.org/mongodb/go-client-mongodb-atlas.svg?branch=master)](https://travis-ci.org/mongodb/go-client-mongodb-atlas)

A Go HTTP client for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/).

You can view the Official API docs here: https://docs.atlas.mongodb.com/api/

## Installation

To get the latest version run this command:

```sh
go get github.com/mongodb/go-client-mongodb-atlas
```

## Usage

```go
import "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
```

## Authentication

The Atlas API uses [HTTP Digest Authentication](https://docs.atlas.mongodb.com/api/#api-authentication). Provide your Atlas PUBLIC_KEY as the username and PRIVATE_KEY as the password as part of the HTTP request. See Programmatic API Keys docs for more detailed information: https://docs.atlas.mongodb.com/configure-api-access/#atlas-prog-api-key.

We use the following library to get HTTP Digest Auth:

https://github.com/Sectorbob/mlab-ns2/gae/ns/digest

## Example Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func newClient(publicKey, privateKey string) (*mongodbatlas.Client, error) {

	//Setup a transport to handle digest
	transport := digest.NewTransport(publicKey, privateKey)

	//Initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	//Initialize the MongoDB Atlas API Client.
	return mongodbatlas.NewClient(client), nil
}

func main() {
	publicKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	privateKey := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
	projectID := os.Getenv("MONGODB_ATLAS_PROJECT_ID")

	if publicKey == "" || privateKey == "" || projectID == "" {
		log.Fatalln("MONGODB_ATLAS_PROJECT_ID, MONGODB_ATLAS_PUBLIC_KEY and MONGODB_ATLAS_PRIVATE_KEY must be set to run this example")
	}

	client, err := newClient(publicKey, privateKey)
	if err != nil {
		log.Fatalf(err.Error())
	}

	clusters, _, err := client.Clusters.List(context.Background(), projectID, nil)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%+v \n", clusters)

}
```

## Versioning

Each version of the client is tagged and the version is updated accordingly.

To see the list of past versions, run `git tag`.

## Roadmap

This library is being initially developed for [mongocli](https://github.com/mongodb/mongocli),
[Atlas Terraform Provider](https://github.com/terraform-providers/terraform-provider-mongodbatlas), [Atlas Vault Plugin](https://github.com/hashicorp/vault-plugin-secrets-mongodbatlas), and [Atlas Cloudformation Provider](https://github.com/mongodb/mongodbatlas-cloudformation-resources)
so API methods will likely be implemented in the order that they are
needed by those those projects.

## Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) Guide.

## License

`go-client-mongodb-atlas` is released under the Apache 2.0 license. See [LICENSE](LICENSE)
