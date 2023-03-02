#!/usr/bin/env bash
set -o errexit
set -o nounset

#########################################################
# Fetch openapi from remote file
# Environment variables:
#   CURRENT_REVISION - current revision of the versioned API
#   OPENAPI_FILE_NAME - openapi file name to use
#   OPENAPI_FOLDER - folder for saving openapi file
#########################################################

## Input variables with defaults

## Locked version of the versioned API
CURRENT_REVISION=${CURRENT_REVISION:-"2023-02-01"}

## OpenAPI file (latest)
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-"atlas-api.yaml"}

## Base URL
API_BASE_URL=${API_BASE_URL:-"https://cloud.mongodb.com/api/openapi"}

## Folder used for fetching files
OPENAPI_FOLDER=${OPENAPI_FOLDER:-"../openapi"}

openapi_url="$API_BASE_URL/spec/2.0?version=$CURRENT_REVISION"
versions_url="$API_BASE_URL/versions"

pushd "$OPENAPI_FOLDER"

echo "Fetching versions from $versions_url"

curl --show-error --fail --silent -o "versions.yaml" \
     -H "Accept: application/yaml" "$versions_url"

echo "Fetching api from $openapi_url to $OPENAPI_FILE_NAME"

curl --show-error --fail --silent -o "$OPENAPI_FILE_NAME" \
     -H "Accept: application/yaml" "$openapi_url"

popd -0 
