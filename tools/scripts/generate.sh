#!/usr/bin/env bash
set -e

OPENAPI_FOLDER=${OPENAPI_FOLDER:-"./openapi"}
OPENAPI_FILE="$OPENAPI_FOLDER/atlas-api.yaml"
TRANSFORMED_FILE="$OPENAPI_FOLDER/atlas-api-transformed.yaml"
PACKAGE="v1"
## Version of the versioned API
CLIENT_VERSION=${CLIENT_VERSION:-"1.20230201.0-dev1"}

SDK_ROOT=${SDK_ROOT:-"./"}

echo "Running generation pipeline"

echo "Creating new $TRANSFORMED_FILE OpenAPI file from $OPENAPI_FILE"
cp "$OPENAPI_FILE" "$TRANSFORMED_FILE"

npm install
npm run sdk:transform -- "$TRANSFORMED_FILE"

npm exec openapi-generator-cli -- generate \
    -c "./config/go.yaml" -i "$TRANSFORMED_FILE" -o "$SDK_ROOT" \
    --package-name="$PACKAGE" \
    --ignore-file-override=config/.go-ignore \
    --additional-properties=xgenClientVersion="$CLIENT_VERSION"

