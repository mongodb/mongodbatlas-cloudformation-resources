#!/usr/bin/env bash

# This tool deletes the mongodb and AWS resources used for `cfn test` as inputs.
set -o errexit
set -o nounset
set -o pipefail

# Run from resource root so ./inputs/ and ./test/ exist
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
resourceRoot="$(dirname "$(dirname "$scriptDir")")"
cd "$resourceRoot"

./test/cfn-test-delete-inputs.sh
