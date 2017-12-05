#!/bin/bash

docker build -t app-test --target builder . || exit 1

# Run unit tests and publish test coverage to Coveralls.
docker run app-test go test -v ./... || exit 1
docker run app-test goveralls -repotoken $COVERALLS_TOKEN || exit 1

# Build the app container.
docker build -t app . || exit 1

# Run the app container for fun.
docker run -it app
