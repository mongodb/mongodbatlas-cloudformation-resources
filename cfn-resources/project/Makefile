.PHONY: build test clean
tags=logging callback metrics scheduler
cgo=0
goos=linux
goarch=amd64

DEBUG_FLAGS=all=-N -l
CFNREP_GIT_SHA?=$(shell git rev-parse HEAD)
LINKER_FLAGS=-s -w -X github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=info -X github.com/mongodb/mongodbatlas-cloudformation-resources/version.Version=${CFNREP_GIT_SHA}

build:
	@echo "==> Building handler binary"
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="$(LINKER_FLAGS)" -tags="$(tags)" -o bin/bootstrap cmd/main.go

debug:
	@echo "==> Building handler binary for debugging"
	cfn generate
	env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -gcflags="$(DEBUG_FLAGS)" -ldflags="$(LINKER_FLAGS)" -tags="$(tags)" -o bin/bootstrap cmd/main.go


clean:
	rm -rf bin

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
