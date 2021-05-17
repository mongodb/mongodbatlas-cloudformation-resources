.PHONY: submit test
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

submit:
	cd cfn-resources && ./cfn-submit-helper.sh $(filter-out $@,$(MAKECMDGOALS))
test:
	cd cfn-resources && ./cfn-testing-helper.sh $(filter-out $@,$(MAKECMDGOALS))
