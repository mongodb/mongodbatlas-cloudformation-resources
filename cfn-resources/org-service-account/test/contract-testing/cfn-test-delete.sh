#!/usr/bin/env bash

# This tool deletes the mongodb resources used for `cfn test` as inputs.
# For org-service-account, no pre-created resources need cleanup.
# The CloudFormation test framework handles resource cleanup automatically.

set -o errexit
set -o nounset
set -o pipefail

echo "No test resources to delete"
