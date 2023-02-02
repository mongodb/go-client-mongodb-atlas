# MongoDB Atlas Versioned SDK

A Go HTTP client for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/)

## Usage

### Adding Dependency 

```
go install go.mongodb.org/atlas
```

### Using in the code

```go
mongodbatlas import "go.mongodb.org/atlas/mongodbatlas/api/v1"
```

Construct a new Atlas client, then use the various services on the client to
access different parts of the Atlas API. For example:

```go
	mongodbatlas import "go.mongodb.org/atlas/mongodbatlas/api/v1"

   	apiKey := os.Getenv("MDB_API_KEY")
	apiSecret := os.Getenv("MDB_API_SECRET")

	// SDK client with authentication
	sdk := mongodbatlas.NewSDKClientWithCredentials(apiKey, apiSecret)
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

### Authentication

The `go-client-mongodb-atlas` library provides helper method for handling Digest Access authentication. Library using [digest](https://github.com/mongodb-forks/digest) for authentication, but you can always use any other library that provides an `http.Client`:

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
