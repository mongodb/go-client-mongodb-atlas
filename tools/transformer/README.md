## Atlas SDK Transformation Engine

## Usage

1. Apply `x-xgen-go-transform` extension to OpenAPI schema or in Java Annotations:.

2. value of the extension should reflect one of the available transformations

```
x-xgen-go-transform: oneOfMerge
```

3. Run transformation engine

```
OPENAPI_FILE=openapi.yaml node ./index.js
```

> NOTE: provided file file will be overwritten
> NOTE: engine supports yaml openapi files

## Implemented transformations

1. oneOf transformation

> "x-xgen-go-transform": "oneOfMerge"

For parent model containing multiple children:

- Moves all child enum or property fields into parent properties.
- Removes redundant oneOf fields on parent

2. allOf transformation.

> "x-xgen-go-transform": "allOfMerge"

For parent model containing multiple children:

- Moves all parent property fields into children.
- Removes redundant allOf fields on children
