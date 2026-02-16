#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# Cleans up the service account created for testing
#

set -o errexit
set -o nounset
set -o pipefail

if [ ! -f ./inputs/inputs_1_create.json ]; then
	echo "No inputs file found, nothing to clean up"
	exit 0
fi

orgId=$(jq -r '.OrgId' ./inputs/inputs_1_create.json)
clientId=$(jq -r '.ClientId' ./inputs/inputs_1_create.json)

if atlas api serviceAccounts deleteOrgServiceAccount --orgId "$orgId" --clientId "$clientId" --version "2024-08-05" --output json 2>/dev/null; then
	echo "Service account $clientId deletion OK"
else
	exitCode=$?
	if [ $exitCode -eq 0 ]; then
		echo "Service account $clientId deletion OK"
	else
		echo "Warning: Failed cleaning service account: $clientId (may already be deleted)"
	fi
fi
