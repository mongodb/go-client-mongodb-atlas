#!/usr/bin/env bash
set -e

OPENAPI_FOLDER=${OPENAPI_FOLDER:-./openapi}
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-atlas-api.yaml}
SDK_FOLDER=${SDK_FOLDER:-./}

echo "# Running generation pipeline"

transformed_file="atlas-api-transformed.yaml"
client_package="v2"
openapiFileLocation="$OPENAPI_FOLDER/$transformed_file"

echo "# Creating new $transformed_file OpenAPI from $OPENAPI_FILE_NAME"
cp "$OPENAPI_FOLDER/$OPENAPI_FILE_NAME" "$openapiFileLocation"

npm install
npm run sdk:transform -- "$openapiFileLocation"

npm exec openapi-generator-cli -- validate -i "$openapiFileLocation"

echo "# Running Client Generation"

npm exec openapi-generator-cli -- generate \
    -c "./config/go.yaml" -i "$openapiFileLocation" -o "$SDK_FOLDER" \
    --package-name="$client_package" \
    --ignore-file-override=config/.go-ignore

