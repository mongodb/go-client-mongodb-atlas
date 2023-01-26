## Run examples

.PHONY: example
example:
	go run ./examples/example_cluster_aws.go

# Tools 
.PHONY: clean
clean:
	rm -Rf ./api/** ./openapi/**

.PHONY: fetch
fetch:
	./tools/scripts/fetch.sh

.PHONY: generate
generate:
	(cd ./tools && \
	OPENAPI_FOLDER="../openapi" SDK_ROOT="../" \
	./scripts/generate.sh)


