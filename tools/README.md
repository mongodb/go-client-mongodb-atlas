# SDK OpenAPI Tools

Folder contains scripts, generators and other files required to generate Golang client

## Running

### Fetching OpenAPI file

```
make fetch_openapi
```

### Generating Go SDK

```
make generate_client
```

### OpenAPI generator

OpenAPI generator version is set in `openapitools.json` file.

### Transformer

Follow [transformer readme](./transformer/README.md) for more information

### Custom templates

By default `./tools/config/go-templates` folder have all available templates as part of the openapi generator.
This enables us to verify new templates changes

#### Modification of the custom templates

Some of the templates were modified to fit for our general needs.
We use mustache comments to mark templates as modified:

- `{{! X-XGEN-CUSTOM }}` - usually present at the top of the file - means that the whole file is currently customized.
- `{{! X-XGEN-MODIFIED }}` - is put into existing templates to note areas in the code that were customized.

#### Contributing

1. After making change validate correctness of your changes by running

```
(cd .. && make v2-lint v2-test)
```

> NOTE: We need to regenerate our SDK if changes in tools will affect the generated clients.
> PR jobs will verify if changes are up to date.
> To regenerate run:

```
make clean_and_generate
```
