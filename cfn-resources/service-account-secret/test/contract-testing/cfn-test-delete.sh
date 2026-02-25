#!/usr/bin/env bash

# This tool deletes the MongoDB resources used for `cfn test` as inputs.
set -o errexit
set -o nounset
set -o pipefail

./test/cfn-test-delete-inputs.sh
