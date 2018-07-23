#!/bin/bash

docker build -t example-builder --target builder . || exit 1

# Run unit tests and publish test coverage to Coveralls.
docker run example-builder go test -v ./... || exit 1
docker run example-builder goveralls -repotoken $COVERALLS_TOKEN || exit 1

# Build the app container.
docker build -t example . || exit 1

# Run the app container for fun.
docker run -it example
