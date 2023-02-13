# MongoDB Atlas v2 API client

A Go HTTP client for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/)

> NOTE: Atlas API v2 client is still under development. 
> Please consider using v1 client located in `go.mongodb.org/atlas/mongodbatlas` for production usage.


## Usage

### Adding Dependency

```
go install go.mongodb.org/atlas
```

### Using in the code

```go
mongodbatlas import "go.mongodb.org/atlas/mongodbatlas/api/v1alpha"
```

Construct a new Atlas client, then use the various services on the client to
access different parts of the Atlas API. For example:

```go
	mongodbatlas import "go.mongodb.org/atlas/mongodbatlas/api/v1alpha"

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
