.PHONY: build test clean
tags=logging callback metrics scheduler
cgo=0
goos=linux
goarch=amd64

DEBUG_FLAGS=all=-N -l
CFNREP_GIT_SHA?=$(shell git rev-parse HEAD)
LINKER_FLAGS=-s -w -X github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=info -X github.com/mongodb/mongodbatlas-cloudformation-resources/version.Version=${CFNREP_GIT_SHA}

#export E2E_RAND_SUFFIX := $(shell bash -c 'echo $$RANDOM')
#export RESOURCE_TYPE_NAME := MongoDB::Atlas::Project
#export RESOURCE_DIRECTORY_NAME := project
#export RESOURCE_TYPE_NAME_FOR_E2E= $(RESOURCE_TYPE_NAME)$(E2E_RAND_SUFFIX)

build:
	@echo "==> Building handler binary"
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="$(LINKER_FLAGS)" -tags="$(tags)" -o bin/handler cmd/main.go

debug:
	@echo "==> Building handler binary for debugging"
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -gcflags="$(DEBUG_FLAGS)" -ldflags="$(LINKER_FLAGS)" -tags="$(tags)" -o bin/handler cmd/main.go


clean:
	rm -rf bin

e2e-test-run:
	cd ../test && ./publish_cfn_to_registry.sh "MongoDB::Atlas::Project" "project" && go test e2e/project/project_test.go && ./cleanup_cfn.sh

release:
	@echo "==> Release the CFN resource to the Private Registry"
	make build && cfn submit --set-default

create-test-resources:
	@echo "==> Creating test files for contract testing"
	./test/contract-testing/cfn-test-create-inputs.sh

delete-test-resources:
	@echo "==> Delete test resources used for contract testing"
	./test/contract-testing/cfn-test-delete-inputs.sh

run-contract-testing:
	@echo "==> Run contract testing"
	make build
	sam local start-lambda &
	cfn test --function-name TestEntrypoint --verbose
