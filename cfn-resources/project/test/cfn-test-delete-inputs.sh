#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
    echo "usage:$0 "
}


apikeyId=$(jq -r '.ProjectApiKeys[0] | .Key' ./inputs/inputs_1_create.json)
teamId=$(jq -r '.ProjectTeams[0] | .TeamId' ./inputs/inputs_1_create.json)

#delete apikey
if mongocli iam project apikey delete "$apikeyId" --force
then
    echo "$apikeyId apikey deletion OK"
else
    (echo "Failed cleaning apikey:$apikeyId" && exit 1)
fi

#delete team
if mongocli iam team delete "$teamId" --force
then
    echo "$teamId team deletion OK"
else
    (echo "Failed cleaning team:$teamId" && exit 1)
fi

