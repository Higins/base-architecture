#!/bin/sh
export CONFIG_FILE=".env"
export LOG_FILE_PATH="./log/log.json"
export LOG_LEVEL="debug"
go get github.com/99designs/gqlgen/graphql/executor
go run ./main.go
