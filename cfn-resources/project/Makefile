tags=logging callback metrics scheduler
cgo=0 
goos=linux
goarch=amd64

DEBUG_FLAGS=all=-N -l
VERSION?=$(shell git describe --always --tags | cut -d "v" -f 2)
LINKER_FLAGS=-s -w -X github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=info -X github.com/mongodb/mongodbatlas-cloudformation-resources/version.Version=${VERSION}

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

release:
	@echo "==> Release the CFN resource to the Private Registry"
	make build && cfn submit --set-default

e2e-test:
	@echo "==> Run E2E test for the Project Resource"
	./test/e2e/run-e2e-test.sh
