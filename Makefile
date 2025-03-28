tags="logging callback metrics scheduler"
cgo=0 
goos=linux
goarch=amd64

# Set next value at compile time for fallback log-level
ldXflags=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=info
ldXflagsD=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug

MOCKERY_VERSION=v2.42.1
GOLANGCI_VERSION=v2.0.2 # Also update golangci-lint GH action in code-health.yml when updating this version

.PHONY: submit
submit:
	(cd cfn-resources && ./cfn-submit-helper.sh $(filter-out $@,$(MAKECMDGOALS)))

.PHONY: test
test:
	(cd cfn-resources && ./cfn-testing-helper.sh $(filter-out $@,$(MAKECMDGOALS)))

.PHONY: fmt
fmt: ## Format changed go and sh
	@scripts/fmt.sh

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
	go install github.com/vektra/mockery/v2@$(MOCKERY_VERSION)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_VERSION)

.PHONY: link-git-hooks
link-git-hooks: ## Install git hooks
	@echo "==> Installing all git hooks..."
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

.PHONY: lint
lint: ## Run linter
	@scripts/lint.sh

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

