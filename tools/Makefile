# Tools 

export OPENAPI_FOLDER=../openapi
export SDK_ROOT=../api/v1

.PHONY: clean_client
clean_client:
	rm -Rf ../api/v1

.PHONY: fetch_openapi
fetch_openapi:
	./scripts/fetch.sh

.PHONY: generate_client
generate_client:
	./scripts/generate.sh

.PHONY: verify_client
verify_client: clean_client generate_client


