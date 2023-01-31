## Run examples

.PHONY: example
example:
	go run ./examples/example_cluster_aws.go

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
	OPENAPI_FOLDER="../openapi" SDK_ROOT="../" \
	./scripts/generate.sh)

.PHONY: compile_client
compile_client:
	go build ./examples/example_cluster_aws.go

.PHONY: verify_client
verify_client: clean_client generate_client


