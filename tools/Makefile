export OPENAPI_FOLDER=../openapi
export OPENAPI_FILE_NAME=atlas-api.yaml
export SDK_FOLDER=../api/v2

.DEFAULT_GOAL := generate_client

.PHONY: clean_client
clean_client:
	rm -Rf ../api/v2

.PHONY: fetch_openapi
fetch_openapi:
	./scripts/fetch.sh

.PHONY: generate_client
generate_client:
	./scripts/generate.sh

.PHONY: verify_client
verify_client: clean_client generate_client


