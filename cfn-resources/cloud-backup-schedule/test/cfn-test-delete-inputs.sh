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


projectId=$(jq -r '.ProjectId' ./inputs/inputs_1_create.json)
clusterName=$(jq -r '.ClusterName' ./inputs/inputs_1_create.json)


echo $projectId
echo $clusterName

#delete Cluster
if atlas clusters delete "$clusterName" --projectId "${projectId}" --force
then
    echo "$clusterName cluster deletion OK"
else
    (echo "Failed cleaning cluster:$clusterName" && exit 1)
fi

echo "Waiting for cluster to get deleted"
#sleep 900

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
if atlas projects delete "$projectId" --force
then
    echo "$projectId project deletion OK"
else
    (echo "Failed cleaning project:$projectId" && exit 1)
fi
