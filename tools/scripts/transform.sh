#!/usr/bin/env bash
set -o errexit
set -o nounset

#######################################
# Transform openapi from remote file
# Environment variables:
#   OPENAPI_FILE_NAME - openapi file name to use
#   OPENAPI_FOLDER - folder for saving openapi file
#######################################

OPENAPI_FOLDER=${OPENAPI_FOLDER:-./openapi}
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-atlas-api.yaml}

transformed_file="atlas-api-transformed.yaml"
openapiFileLocation="$OPENAPI_FOLDER/$transformed_file"

echo "# Running transformation from $transformed_file OpenAPI from $OPENAPI_FILE_NAME"
cp "$OPENAPI_FOLDER/$OPENAPI_FILE_NAME" "$openapiFileLocation"

npm install
npm run sdk:transform -- "$openapiFileLocation"
