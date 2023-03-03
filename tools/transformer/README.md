## Atlas SDK Transformation Engine

## Usage

1. Apply `x-xgen-go-transform` extension to OpenAPI schema or in Java Annotations:.

2. value of the extension should reflect one of the available transformations

```
x-xgen-go-transform: oneOfMerge
```

3. Run transformation engine

```
OPENAPI_FILE=openapi.yaml node ./transform.js
```

> NOTE: provided file file will be overwritten
> NOTE: engine supports yaml openapi files

## Implemented transformations

1. oneOf transformation
   Applied to all objects that meet either of the following criteria:

> Have extension "x-xgen-go-transform": "oneOfMerge"
> Have "oneOf" field and all the objects referenced are enums

For parent model containing multiple children:

- Moves all child enum or property fields into parent properties.
- Removes redundant oneOf fields on parent

2. allOf transformation.

Applied to all objects that meet the following criteria:

> Have "properties" filed
> Have "oneOf" field

For parent model containing multiple children:

- Moves all parent property fields into children.
- Removes redundant allOf fields on children

3. Discriminator transformation

Applied to all objects that meet the following criteria:

> Have "discriminator" field
> Missing "oneOf" field

For parent model containing discriminator

- Fails if discriminator is missing mapping field (invalid case)
- Uses discriminator mapping in order to add oneOf field

## Adjusting logging

`XGEN_LOGGING_LEVEL` env variable can be used to specify logging levels:

- info
- warn
- debug
- error

Default: 'warn"
