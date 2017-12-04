#!/bin/bash

docker build -t app-test --target builder . || exit 1

# Run the app test container and publish test coverage to Coveralls.
CONTAINER_ID=$(docker run -d app-test tail -f /dev/null) || exit 1
docker exec $CONTAINER_ID go test -v ./... || exit 1
docker exec $CONTAINER_ID goveralls -repotoken $COVERALLS_TOKEN || exit 1

docker stop $CONTAINER_ID

# Build the app container.
docker build -t app . || exit 1

# Run the app container for fun.
docker run -it app
