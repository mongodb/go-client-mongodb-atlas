#!/usr/bin/env bash
set -e

OPENAPI_FOLDER=${OPENAPI_FOLDER:-"./openapi"}
OPENAPI_FILE="$OPENAPI_FOLDER/atlas-api.yaml"
TRANSFORMED_FILE="$OPENAPI_FOLDER/atlas-api-transformed.yaml"
PACKAGE="v2"

SDK_ROOT=${SDK_ROOT:-"./"}

echo "Running generation pipeline"

echo "Creating new $TRANSFORMED_FILE OpenAPI file from $OPENAPI_FILE"
cp "$OPENAPI_FILE" "$TRANSFORMED_FILE"

npm install
npm run sdk:transform -- "$TRANSFORMED_FILE"

npm exec openapi-generator-cli -- generate \
    -c "./config/go.yaml" -i "$TRANSFORMED_FILE" -o "$SDK_ROOT" \
    --package-name="$PACKAGE" \
    --ignore-file-override=config/.go-ignore

