#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes test input files and cleans up resources.
#

set -o errexit
set -o nounset
set -o pipefail

rm -rf inputs
echo "Deleted inputs directory"
