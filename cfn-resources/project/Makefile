.PHONY: build test clean
tags=logging callback metrics scheduler
cgo=0 
goos=linux
goarch=amd64
#export TAGS
#export CGO_ENABLED
#export GOARCH
# Set next value at compile time for fallback log-level
ldXflags=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=info
ldXflagsD=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug
build:
	#make -f makebuild  # this runs build steps required by the cfn cli
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="-s -w -X '$(ldXflags)'" -tags="$(tags)" -o bin/handler cmd/main.go

debug:
	#make -f makebuild  # this runs build steps required by the cfn cli
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="-s -w -X '$(ldXflagsD)'" -tags="$(tags)" -o bin/handler cmd/main.go


clean:
	rm -rf bin
