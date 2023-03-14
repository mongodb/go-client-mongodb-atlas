# go-client-mongodb-atlas
[![PkgGoDev](https://pkg.go.dev/badge/go.mongodb.org/atlas)](https://pkg.go.dev/go.mongodb.org/atlas)
![CI](https://github.com/mongodb/go-client-mongodb-atlas/workflows/CI/badge.svg)

A Go HTTP client for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/).

Note that `go-client-mongodb-atlas` only supports the two most recent major versions of Go.

> NOTE: Atlas API v2 client is still under development. 
> Please consider using v1 client: `go.mongodb.org/atlas/mongodbatlas` for production usage.


## Usage

### Adding Dependency

```
go install go.mongodb.org/atlas
```

### Using in the code

```go
mongodbatlas import "go.mongodb.org/atlas/mongodbatlasv2"
```

Construct a new Atlas client, then use the various services on the client to
access different parts of the Atlas API. For example:

```go
	mongodbatlas import "go.mongodb.org/atlas/mongodbatlasv2"

   	apiKey := os.Getenv("MDB_API_KEY")
	apiSecret := os.Getenv("MDB_API_SECRET")

	sdk := mongodbatlas.NewClient(mongodbatlas.UseDigestAuth(apiKey, apiSecret))
	projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
```

For documentation about obtaining private and public API token go to
https://docs.atlas.mongodb.com/configure-api-access.
The services of a client divide the API into logical chunks and correspond to
the structure of the Atlas API documentation at
https://www.mongodb.com/docs/atlas/reference/api-resources-spec/.

## Examples

Example for creating an dedicated MongoDB cluster on AWS infrastructure

```
go run ./examples/example_cluster_aws.go
```

## Usage for v1 client

```go
import "go.mongodb.org/atlas/mongodbatlas"
```

Construct a new Atlas client, then use the various services on the client to
access different parts of the Atlas API. For example:

```go
client := mongodbatlas.NewClient(nil)
```

The services of a client divide the API into logical chunks and correspond to
the structure of the Atlas API documentation at
https://docs.atlas.mongodb.com/api/.

**NOTE:** Using the [context](https://godoc.org/context) package, one can easily
pass cancellation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.

### Authentication

The `go-client-mongodb-atlas` library does not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle Digest Access authentication for
you. The easiest way to do this is using the [digest](https://github.com/mongodb-forks/digest)
library, but you can always use any other library that provides an `http.Client`.
If you have a private and public API token pair (https://docs.atlas.mongodb.com/configure-api-access),
you can use it with the digest library like:

```go
import (
    "context"
    "log"

    "github.com/mongodb-forks/digest"
    "go.mongodb.org/atlas/mongodbatlas"
)

func main() {
    t := digest.NewTransport("your public key", "your private key")
    tc, err := t.Client()
    if err != nil {
        log.Fatalf(err.Error())
    }

    client := mongodbatlas.NewClient(tc)
    projects, _, err := client.Projects.GetAllProjects(context.Background(), nil)
}
```

Note that when using an authenticated Client, all calls made by the client will
include the specified tokens. Therefore, authenticated clients should
almost never be shared between different users.

## Versioning

Each version of the client is tagged and the version is updated accordingly.

To see the list of past versions, run `git tag`.

To release a new version, first ensure that [Version](./mongodbatlas/mongodbatlas.go) is updated 
(i.e., before running `git push origin vx.y.z`, verify that `Version=x.y.z` should match the tag being pushed to GitHub)

## Roadmap

This library is being initially developed for [mongocli](https://github.com/mongodb/mongocli),
[Atlas Terraform Provider](https://github.com/mongodb/terraform-provider-mongodbatlas), 
[Atlas Vault Plugin](https://github.com/hashicorp/vault-plugin-secrets-mongodbatlas), and 
[Atlas Cloudformation Provider](https://github.com/mongodb/mongodbatlas-cloudformation-resources)
so API methods will likely be implemented in the order that they are
needed by those projects.

## Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) Guide.

## License

`go-client-mongodb-atlas` is released under the Apache 2.0 license. See [LICENSE](LICENSE)
