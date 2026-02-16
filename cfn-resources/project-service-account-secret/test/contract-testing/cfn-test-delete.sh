#!/usr/bin/env bash

# This tool cleans up resources created for `cfn test`.
set -o errexit
set -o nounset
set -o pipefail

./test/cfn-test-delete-inputs.sh
