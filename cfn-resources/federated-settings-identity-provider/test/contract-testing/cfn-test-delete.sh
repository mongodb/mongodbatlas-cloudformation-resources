#!/usr/bin/env bash

# This tool deletes the mongodb resources used for `cfn test` as inputs.
set -o errexit
set -o nounset
set -o pipefail

echo "No cleanup required for federated settings identity provider."
