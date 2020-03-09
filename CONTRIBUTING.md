# Contributing to MongoDB CLI

Thanks for your interest in contributing to `go-client-mongodb-atlas`, 
this document describe the necessary steps to get a development environment going and the best way to contribute back to the project

## Development setup

### Prerequisite Tools 
- [Git](https://git-scm.com/)
- [Go (at least Go 1.12)](https://golang.org/dl/)

### Environment
- Fork the repository.
- Clone your forked repository locally.
- We use Go Modules to manage dependencies, so you can develop outside of your `$GOPATH`.
- We use [golangci-lint](https://github.com/golangci/golangci-lint) to lint our code, you can install it locally via `make tools`.

## Building and testing

The following is a short list of commands that can be run in the root of the project directory

- Run `make test` to run all unit tests.
- Run `make lint` to validate against our linting rules.
