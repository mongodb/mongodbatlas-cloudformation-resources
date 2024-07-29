#!/usr/bin/env bash
# encryptionatrest.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project id>"
	echo "This test just creates a static access list, you can edit the test/ inputs if needed."

}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectID="${1}"
jq --arg projectID "$projectID" \
	--arg KMS_KEY "$KMS_KEY" \
	--arg KMS_ROLE "$KMS_ROLE" \
	--arg region "$KMS_REGION" \
	'.desiredResourceState.AwsKmsConfig.CustomerMasterKeyID?|=$KMS_KEY | .desiredResourceState.AwsKmsConfig.RoleID?|=$KMS_ROLE | .desiredResourceState.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.ProjectId?|=$projectID | .desiredResourceState.AwsKmsConfig.Region?|=$region ' \
	"$(dirname "$0")/encryptionatrest.sample-cfn-request.json"
