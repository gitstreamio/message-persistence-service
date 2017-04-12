#!/bin/sh

go generate ./...

CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"'  -o message-persistence-service .

docker build -t message-persistence-service .
