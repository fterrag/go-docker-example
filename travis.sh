#!/bin/bash

docker build -t app-test --target builder . || exit 1
docker run -it app-test go test ./... || exit 1

docker build -t app . || exit 1
docker run -it app
