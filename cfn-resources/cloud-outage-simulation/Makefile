.PHONY: build test clean
tags=logging callback metrics scheduler
cgo=0
goos=linux
goarch=amd64
CFNREP_GIT_SHA?=$(shell git rev-parse HEAD)
ldXflags=-s -w -X github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=info -X github.com/mongodb/mongodbatlas-cloudformation-resources/version.Version=2
ldXflagsD=-s -w -X github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug -X github.com/mongodb/mongodbatlas-cloudformation-resources/version.Version=2

build:
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="$(ldXflags)" -tags="$(tags)" -o bin/handler cmd/main.go

debug:
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="$(ldXflagsD)" -tags="$(tags)" -o bin/handler cmd/main.go

clean:
	rm -rf bin
