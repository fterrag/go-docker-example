# go-docker-example

[![Build Status](https://travis-ci.org/fterrag/go-docker-example.svg?branch=master)](https://travis-ci.org/fterrag/go-docker-example) [![Coverage Status](https://coveralls.io/repos/github/fterrag/go-docker-example/badge.svg?branch=master)](https://coveralls.io/github/fterrag/go-docker-example?branch=master)

A small example of building and running a Go app inside of a Docker container.

- Uses Go Modules
- Runs `go test` during build
- Uses Docker's multi-stage builds
- Builds and runs (as an example) on Travis CI
- Uses Coveralls for test coverage history

## Running Locally

Be sure to have Go 1.12+ installed and try running:

```bash
$ go run cmd/example/main.go
```

## Coveralls

When running builds under your own Travis CI account (or any CI service), the
[travis.sh](/travis.sh) file requires the `COVERALLS_TOKEN` variable to be set to your [Coveralls](https://coveralls.io/) repo token. In Travis CI, you can set this variable
in your project's settings under the Environment Variables section. If you do not
want to use Coveralls, you may safely remove lines 10-12 in Dockerfile and line 7 in travis.sh.
