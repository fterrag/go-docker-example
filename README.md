# go-docker-example

[![Build Status](https://travis-ci.org/fterrag/go-docker-example.svg?branch=master)](https://travis-ci.org/fterrag/go-docker-example) [![Coverage Status](https://coveralls.io/repos/github/fterrag/go-docker-example/badge.svg?branch=master)](https://coveralls.io/github/fterrag/go-docker-example?branch=master)

A small example of building and running a Go app inside of a Docker container.

- Runs `go test` during build
- Uses Docker's multi-stage builds
- Builds and runs (as an example) on Travis CI
- Uses Coveralls for test coverage history

## Running Locally

Be sure to have Go installed and try running:

```bash
go run main.go
```
