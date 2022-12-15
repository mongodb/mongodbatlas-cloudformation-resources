.PHONY: build test clean
tags=logging callback metrics scheduler
cgo=0
goos=linux
goarch=amd64
ldXflags=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=warning
ldXflagsD=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug
build:
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="-s -w -X '$(ldXflags)'" -tags="$(tags)" -o bin/handler cmd/main.go

debug:
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="-s -w -X '$(ldXflagsD)'" -tags="$(tags)" -o bin/handler cmd/main.go

clean:
	rm -rf bin