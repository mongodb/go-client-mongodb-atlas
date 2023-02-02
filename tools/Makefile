# Tools 
.PHONY: clean_client
clean_client:
	rm -Rf ./api/**

.PHONY: fetch_openapi
fetch_openapi:
	./tools/scripts/fetch.sh

.PHONY: generate_client
generate_client:
	(cd ./tools && \
	OPENAPI_FOLDER="../openapi" SDK_ROOT="../api/v1" \
	./scripts/generate.sh)

.PHONY: verify_client
verify_client: clean_client generate_client


