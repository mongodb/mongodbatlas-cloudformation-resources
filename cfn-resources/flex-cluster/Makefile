.PHONY: build test clean
tags=logging callback metrics scheduler
cgo=0
goos=linux
goarch=amd64
CFNREP_GIT_SHA?=$(shell git rev-parse HEAD)
ldXflags=-s -w -X github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=info -X github.com/mongodb/mongodbatlas-cloudformation-resources/version.Version=${CFNREP_GIT_SHA}
ldXflagsD=-X github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug -X github.com/mongodb/mongodbatlas-cloudformation-resources/version.Version=${CFNREP_GIT_SHA}

build:
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="$(ldXflags)" -tags="$(tags)" -o bin/bootstrap cmd/main.go

debug:
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="$(ldXflagsD)" -tags="$(tags)" -o bin/debug cmd/main.go

clean:
	rm -rf bin

submit: clean build # submit to private registry must use release build not debug build
	 cfn submit --set-default --region us-east-1
