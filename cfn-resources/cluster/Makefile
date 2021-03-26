.PHONY: build test clean
tags=logging callback metrics scheduler
#tags=logging
cgo=0 
goos=linux
goarch=amd64
#export TAGS
#export CGO_ENABLED
#export GOARCH

build:
	#make -f makebuild  # this runs build steps required by the cfn cli
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="-s -w" -tags="$(tags)" -o bin/handler cmd/main.go


clean:
	rm -rf bin
