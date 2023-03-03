#!/usr/bin/env bash

# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -x
echo "--------------------------------delete key and key policy document policy document starts ----------------------------"


projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
clusterName=$(jq -r '.ClusterName' ./inputs/inputs_1_create.json)
export MCLI_PROJECT_ID=$projectId


keyRegion=$AWS_DEFAULT_REGION
if [ -z "$keyRegion" ]; then
	keyRegion=$(aws configure get region)
fi
# shellcheck disable=SC2001
keyRegion=$(echo "$keyRegion" | sed -e "s/-/_/g")
keyRegion=$(echo "$keyRegion" | tr '[:lower:]' '[:upper:]')
echo "$keyRegion"

roleName="mongodb-test-export-role-${CFN_TEST_TAG}-${keyRegion}"
policyName="atlas-bucket-role-policy-${keyRegion}"

echo "--------------------------------delete key and key policy document policy document ends ----------------------------"
pwd
trustPolicy=$(jq '.Statement[0].Condition.StringEquals["sts:ExternalId"]' "add-policy.json")
echo "$trustPolicy"
roleExternalID=$(${trustPolicy##*/})
# shellcheck disable=SC2001
atlasAssumedRoleExternalID=$(echo "${roleExternalID}" | sed 's/"//g')
echo "$atlasAssumedRoleExternalID"

roleId=$(atlas cloudProviders accessRoles list --output json | jq --arg roleID "${atlasAssumedRoleExternalID}" -r '.awsIamRoles[] |select(.atlasAssumedRoleExternalId |test( $roleID)) |.roleId')
echo "$roleId"

atlas cloudProviders accessRoles aws deauthorize "${roleId}" --force
echo "--------------------------------delete role starts ----------------------------"

aws iam delete-role-policy --role-name "$roleName" --policy-name "$policyName"
aws iam delete-role --role-name "$roleName"
echo "--------------------------------delete role ends ----------------------------"


# delete cluster
if atlas clusters delete "$clusterName" --projectId "${projectId}" --force
then
    echo "$clusterName cluster deletion OK"
else
    (echo "Failed cleaning cluster:$clusterName" && exit 1)
fi

echo "Waiting for cluster to get deleted"

status=$(atlas clusters describe "${clusterName}" --projectId "${projectId}" --output=json | jq -r '.stateName')
echo "status: ${status}"

while atlas clusters describe "${clusterName}" --projectId "${projectId}"; do
        sleep 30
        if atlas clusters describe "${clusterName}" --projectId "${projectId}"
        then
          status=$(atlas clusters describe "${clusterName}" --projectId "${projectId}"  --output=json | jq -r '.stateName')
        else
          status="DELETED"
        fi
        echo "status: ${status}"
done

#delete project
if atlas projects delete "$projectId" --force; then
	echo "$projectId project deletion OK"
else
	(echo "Failed cleaning project:$projectId" && exit 1)
fi
