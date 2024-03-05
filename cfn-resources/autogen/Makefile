.PHONY: submit test
tags=logging callback metrics scheduler
cgo=0
goos=linux
goarch=amd64
ldXflags=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=warning
ldXflagsD=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug


create:
	echo "Enter the go import path as: github.com/mongodb/mongodbatlas-cloudformation-resources/$(dirName) "
	cd .. && \
	mkdir $(dirName) && cp autogen/schemas/mongodb-atlas-$(resourceName).json $(dirName)/mongodb-atlas-$(resourceName).json && \
	cd $(dirName) && \
	cfn init -t MongoDB::Atlas::$(typeName) -a r && cfn generate
	cd ../autogen && \
	go run main.go -res ../$(dirName) -schemaName $(resourceName)
	cp -fr templateMakeFile ../$(dirName)/Makefile
	cd ../$(dirName) && rm go.mod && rm .gitignore
schema:
	cd schema-gen && go build
	./schema-gen/schema-gen
