tags="logging callback metrics scheduler"
cgo=0
goos=linux
goarch=amd64

# Set next value at compile time for fallback log-level
ldXflags=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=info
ldXflagsD=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug

MOCKERY_VERSION=v3.5.3
GOLANGCI_VERSION=v2.11.4

export PATH := $(shell go env GOPATH)/bin:$(PATH)
export SHELL := env PATH=$(PATH) /bin/bash

default: fix

.PHONY: submit
submit:
	(cd cfn-resources && ./cfn-submit-helper.sh $(filter-out $@,$(MAKECMDGOALS)))

.PHONY: test
test:
	(cd cfn-resources && ./cfn-testing-helper.sh $(filter-out $@,$(MAKECMDGOALS)))

.PHONY: fix
fix: ## Format, lint-fix, tidy, and apply go fix
	(cd cfn-resources && gofmt -s -w .)
	(cd cfn-resources && goimports -w .)
	(cd cfn-resources && golangci-lint run --fix --timeout 5m)
	(cd cfn-resources && go mod tidy)
	(cd cfn-resources && go fix ./...)

.PHONY: verify
verify: ## Verify Go code without modifying files. Usage: make verify [files="file1.go file2.go"]
ifdef files
	$(eval files_rel := $(patsubst cfn-resources/%,%,$(files)))
	@bad_fmt=$$(cd cfn-resources && gofmt -l -s $(files_rel)); \
	if [ -n "$$bad_fmt" ]; then echo "ERROR: gofmt issues:"; echo "$$bad_fmt"; exit 1; fi
	@bad_imports=$$(cd cfn-resources && goimports -l $(files_rel)); \
	if [ -n "$$bad_imports" ]; then echo "ERROR: goimports issues:"; echo "$$bad_imports"; exit 1; fi
	(cd cfn-resources && golangci-lint run $(addsuffix ...,$(sort $(dir $(files_rel)))))
	(cd cfn-resources && go fix -diff $(addprefix ./,$(addsuffix ...,$(sort $(dir $(files_rel))))))
else
	@bad_fmt=$$(cd cfn-resources && gofmt -l -s .); \
	if [ -n "$$bad_fmt" ]; then echo "ERROR: gofmt issues:"; echo "$$bad_fmt"; exit 1; fi
	@bad_imports=$$(cd cfn-resources && goimports -l .); \
	if [ -n "$$bad_imports" ]; then echo "ERROR: goimports issues:"; echo "$$bad_imports"; exit 1; fi
	(cd cfn-resources && golangci-lint run --timeout 5m)
	(cd cfn-resources && go mod tidy -diff)
	(cd cfn-resources && go fix -diff ./...)
endif

.PHONY: tools
tools:  ## Install dev tools
	@echo "==> Installing dev tools..."
	go install github.com/icholy/gomajor@latest
	go install github.com/google/addlicense@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/google/go-licenses@latest
	go install mvdan.cc/sh/v3/cmd/shfmt@latest
	go install github.com/rhysd/actionlint/cmd/actionlint@latest
	go install go.uber.org/mock/mockgen@latest
	go install github.com/vektra/mockery/v3@$(MOCKERY_VERSION)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_VERSION)

.PHONY: link-git-hooks
link-git-hooks: ## Install git hooks
	@echo "==> Installing all git hooks..."
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

.PHONY: unit-test
unit-test:
	(cd cfn-resources && go test $$(go list ./... | grep -v /e2e))

.PHONY: update-atlas-sdk
update-atlas-sdk: ## Update the atlas-sdk dependency
	(cd cfn-resources && ./scripts/update-sdk.sh)

.PHONY: generate-mocks
generate-mocks: # uses mockery to generate mocks in folder `cfn-resources/testutil/mocksvc`
	(cd cfn-resources && mockery)

# resulting file placed in cfn-resources/resource-versions.md
# aws regions must defined by using AWS_REGIONS env variable, example: `export AWS_REGIONS=af-south-1,ap-east-1`
.PHONY: generate-resource-versions-markdown
generate-resource-versions-markdown:
	(cd cfn-resources && go run tool/markdown-generator/*.go)

.PHONY: gen-sbom-and-ssdlc-report
gen-sbom-and-ssdlc-report:
	./scripts/gen-purl.sh $(RESOURCE)
	./scripts/generate-sbom.sh $(RESOURCE) $(VERSION)
	./scripts/gen-ssdlc-report.sh $(RESOURCE) $(VERSION)
	./scripts/upload-sbom.sh $(RESOURCE) $(VERSION)

.PHONY: augment-sbom
augment-sbom:
	./scripts/augment-sbom.sh $(RESOURCE) $(VERSION)
