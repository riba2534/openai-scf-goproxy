#!/bin/bash

go mod tidy
rm -rf main main.zip

GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o main main.go
zip main.zip main scf_bootstrap
