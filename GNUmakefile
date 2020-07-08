SOURCE_FILES?=./...
PKG_NAME=mongodbatlas
GOLANGCI_VERSION=v1.27.0
COVERAGE=coverage.out

export GO111MODULE := on
export PATH := ./bin:$(PATH)

default: build

build:
	go install ./$(PKG_NAME)

test:
	go test $(SOURCE_FILES) -coverprofile $(COVERAGE) -timeout=30s -parallel=4 -cover -race

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)
	goimports -w ./$(PKG_NAME)

lint-fix:
	golangci-lint run --fix

lint:
	golangci-lint run $(SOURCE_FILES)

check: test lint-fix

tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_VERSION)

.PHONY: build test fmt lint lint-fix check tools
