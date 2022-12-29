#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

# NOTE: You need to set Federation Settings Id and Orgnization Id, in order to execute this resource.
#       You can get the Federation Settings Id on Atlas UI under the 'Manage Federation Settings' cosole

set -o nounset
set -o pipefail
WORDTOREMOVE="template."
function usage {
    echo "usage:$0 <project_name>"
}

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs


cd "$(dirname "$0")" || exit
for inputFile in inputs_*;
do
  outputFile=${inputFile//$WORDTOREMOVE/};
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
       --arg org "$ATLAS_ORG_ID" \
     --arg FederationSettingsId "$ATLAS_FEDERATED_SETTINGS_ID" \
     --arg group_id "$projectId" \
     '.FederationSettingsId?|=$FederationSettingsId | .OrgId?|=$org | .RoleAssignments[0].GroupId?|=$group_id | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey ' \
     "$inputFile" > "../inputs/$outputFile"
done
cd ..

ls -l inputs