#!/usr/bin/env bash

# This tool deletes the mongodb resources used for `cfn test` as inputs.
# Delegates to the main deletion script in the test directory.
set -o errexit
set -o nounset
set -o pipefail

# Call the main deletion script
./test/cfn-test-delete-inputs.sh
