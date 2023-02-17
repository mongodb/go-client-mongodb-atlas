SOURCE_FILES?=./...
PKG_NAME=mongodbatlas
GOLANGCI_VERSION=v1.50.1
COVERAGE=coverage.out
APIV2_FOLDER=./api

export GO111MODULE := on
export PATH := ./bin:$(PATH)

default: build

.PHONY: link-git-hooks
link-git-hooks:
	find .git/hooks -type l -exec rm {} \;
	find githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

.PHONY: build
build:
	go install ./$(PKG_NAME)

.PHONY: test
test:
	go test $(SOURCE_FILES) -coverprofile $(COVERAGE) -timeout=30s -parallel=4 -cover -race

.PHONY: fmt
fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)
	goimports -w ./$(PKG_NAME)

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

.PHONY: lint
lint:
	golangci-lint run $(SOURCE_FILES)

.PHONY: check
check: test lint-fix

.PHONY: tools
tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_VERSION)

.PHONY: addcopy
addcopy:
	@scripts/add-copy.sh

TAG=$(patsubst v%,%,$(shell git describe --tags --dirty --always))
.PHONY: check-version
check-version:
	scripts/check-version.sh "$(TAG)"

.PHONY: openapi-pipeline
openapi-pipeline: tools
	echo "Running OpenAPI Generation and Validation process"
	$(MAKE) -C tools clean_client
	echo "Running client generation"
	$(MAKE) -C tools generate_client
	echo "Validating generated SDK"
	$(MAKE) v2-lint
	$(MAKE) v2-test
	$(MAKE) v2-examples-build

.PHONY: v2-lint
v2-lint:
	golangci-lint run "$(APIV2_FOLDER)/$(SOURCE_FILES)"

.PHONY: v2-examples-build
v2-examples-build:
	go build "$(APIV2_FOLDER)/examples/example_cluster_aws.go"

.PHONY: v2-test
v2-test:
	go test "$(APIV2_FOLDER)/$(SOURCE_FILES)" -coverprofile $(COVERAGE) -timeout=30s -parallel=4 -cover -race
