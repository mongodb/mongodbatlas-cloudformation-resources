#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail
set -x

function usage {
    echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs
name="${1}"
projectName="${1}"
if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

projectId=$(mongocli iam projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(mongocli iam projects create "${projectName}" --output=json | jq -r '.id')
    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Created project \"${projectName}\" with id: ${projectId}"
echo $(mongocli iam user invite --country IN --email testUser@gmail.com --firstName testUser --lastName testUser --password testUser --username testUser@gmail.com)
 userName=$(mongocli iam organizations users list  --orgId 6177d6c90e0c7f15f342a05d --output json | jq '.results[0]|.emailAddress')
userName=$(echo ${userName} | sed 's/"//g')
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg userName "$userName" \
   '.Usernames?|=[$userName]|.OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg userName "$userName" \
   '.Usernames?|=[$userName]|.OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
      "$(dirname "$0")/inputs_1_update.template.json" > "inputs/inputs_1_update.json"
name="${name}- more B@d chars !@(!(@====*** ;;::"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg org "$ATLAS_ORG_ID" \
   --arg userName "$userName" \
   '.Usernames?|=[$userName]|.OrgId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey' \
     "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"


ls -l inputs
