#!/bin/sh
# cgo disabled avoid dynamic linking, to work with Docker scratch image
CGO_ENABLED=0 go build app.go
docker build -t acogoluegnes/whoami .
