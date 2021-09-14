#!/bin/sh

# Load Env File
set -a
[ -f .env ] && . .env
set +a

# Download all the dependencies that are required in your source files and update go.mod file with that dependency and
# remove all dependencies from the go.mod file which are not required in the source files.
go mod tidy

# Run app
go run main.go
