export OPENAPI_FOLDER=../openapi
export OPENAPI_FILE_NAME=atlas-api.yaml
export SDK_FOLDER=../mongodbatlasv2/latest

.DEFAULT_GOAL := generate_client

.PHONY: clean_client
clean_client:
	rm -Rf ../mongodbatlasv2/latest

.PHONY: fetch_openapi
fetch_openapi:
	./scripts/fetch.sh

.PHONY: transform_openapi
transform_openapi:
	./scripts/transform.sh

.PHONY: generate_client
generate_client:
	./scripts/generate.sh

.PHONY: clean_and_generate
clean_and_generate: clean_client generate_client
