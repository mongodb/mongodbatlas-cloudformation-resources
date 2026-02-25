#!/usr/bin/env bash

# This tool generates the resources and json files in the inputs/ for `cfn test`.
set -o errexit
set -o nounset
set -o pipefail

if [ -z "${MONGODB_ATLAS_ORG_ID:-}" ]; then
	echo "Error: MONGODB_ATLAS_ORG_ID environment variable is not set"
	exit 1
fi

orgId="${MONGODB_ATLAS_ORG_ID}"
echo "Using OrgId: $orgId"

./test/cfn-test-create-inputs.sh "$orgId"
