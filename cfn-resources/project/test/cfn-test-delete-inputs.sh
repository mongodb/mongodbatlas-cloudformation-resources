#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

function usage {
	echo "usage:$0 "
}

apikeyId=$(jq -r '.ProjectApiKeys[0] | .Key' ./inputs/inputs_1_create.json)
teamId=$(jq -r '.ProjectTeams[0] | .TeamId' ./inputs/inputs_1_create.json)

#delete apikey
if atlas organizations apikey delete "$apikeyId" --force; then
	echo "$apikeyId apikey deletion OK"
else
	(echo "Failed cleaning apikey:$apikeyId" && exit 1)
fi

#delete team
if atlas teams delete "$teamId" --force; then
	echo "$teamId team deletion OK"
else
	(echo "Failed cleaning team:$teamId" && exit 1)
fi
