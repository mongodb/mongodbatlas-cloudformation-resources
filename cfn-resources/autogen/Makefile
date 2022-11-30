.PHONY: submit test
tags=logging callback metrics scheduler
cgo=0
goos=linux
goarch=amd64
#export TAGS
#export CGO_ENABLED
#export GOARCH
# Set next value at compile time for fallback log-level
ldXflags=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=warning
ldXflagsD=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug

# example for resource: MongoDB::Atlas::Teams
# example for go path: github.com/mongodb/mongodbatlas-cloudformation-resources/invites
# example for go path: github.com/mongodb/mongodbatlas-cloudformation-resources/indexes

create:
	echo "Enter the go import path as: github.com/mongodb/mongodbatlas-cloudformation-resources/$(dirName) "
	cd .. && \
	mkdir $(dirName) && cp autogen/configs/mongodb-atlas-$(resourceName).json $(dirName)/mongodb-atlas-$(resourceName).json && \
	cd $(dirName) && \
	cfn init -t MongoDB::Atlas::$(resourceName) -a r && cfn generate
	cd ../autogen && \
	go run main.go -res ../$(dirName) -schemaName $(resourceName)
	cd ../$(dirName) && rm go.mod && rm .gitignore

schema:
	cd schema-gen && go build
	./schema-gen/schema-gen
