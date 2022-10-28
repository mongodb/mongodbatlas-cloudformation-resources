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
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectID "$projectID" \
   '.desiredResourceState.properties.AwsKms.CustomerMasterKeyID?|=$KMS_KEY | .desiredResourceState.properties.AwsKms.RoleID?|=$KMS_ROLE | .desiredResourceState.properties.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.properties.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.properties.ProjectId?|=$projectId | .desiredResourceState.properties.AwsKms.Region?|=$region ' \
   "$(dirname "$0")/encryptionatrest.sample-cfn-request.json"
