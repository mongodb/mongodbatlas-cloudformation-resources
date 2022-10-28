#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

#set -x

function usage {
    echo "usage:$0 <project_name>"
    echo "Creates a new project and an AccessList for testing"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi
rm -rf inputs
mkdir inputs

if [[   -v ATLAS_EXISTING_PROJECT_NAME ]]; then
      projectName=$ATLAS_EXISTING_PROJECT_NAME
else
      projectName="${1}"
fi
projectId=$(mongocli iam projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(mongocli iam projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi
echo -e "=====\nrun this command to clean up\n=====\nmongocli iam projects delete ${projectId} --force\n====="

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg projectId "$projectId" \
   --arg KMS_KEY "$KMS_KEY" \
   --arg KMS_ROLE "$KMS_ROLE" \
   --arg region "US_EAST_1" \
   '.AwsKms.CustomerMasterKeyID?|=$KMS_KEY | .AwsKms.RoleID?|=$KMS_ROLE | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .ProjectId?|=$projectId | .AwsKms.Region?|=$region ' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg KMS_KEY "$KMS_KEY" \
   --arg KMS_ROLE "$KMS_ROLE" \
   --arg projectId "$projectId" \
   --arg region "US_EAST_1" \
    '.AwsKms.CustomerMasterKeyID?|=$KMS_KEY | .AwsKms.RoleID?|=$KMS_ROLE | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .ProjectId?|=$projectId | .AwsKms.Region?|=$region' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"

ls -l inputs
#mongocli iam projects delete "${projectId}" --force
