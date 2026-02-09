.PHONY: create-test-resources delete-test-resources test

PROJECT_NAME ?= cfn-test-project-$(shell date +%s)

create-test-resources:
	@echo "==> Creating test resources for access role and IAM authorization"
	@chmod +x ./cfn-test-create-inputs.sh
	./cfn-test-create-inputs.sh $(PROJECT_NAME)

delete-test-resources:
	@echo "==> Deleting test resources"
	@chmod +x ./cfn-test-delete-inputs.sh
	./cfn-test-delete-inputs.sh $(PROJECT_NAME)

test: create-test-resources
	@echo "==> Test resources created successfully"
	@echo "==> Run 'make delete-test-resources PROJECT_NAME=<your-project-name>' to clean up"

clean:
	@echo "==> Cleaning up generated files"
	@rm -f trust-policy.json
