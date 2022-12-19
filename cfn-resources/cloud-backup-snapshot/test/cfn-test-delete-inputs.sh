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


projectId=$(jq -r '.GroupId' ./inputs/inputs_1_create.json)
clusterName=$(jq -r '.ClusterName' ./inputs/inputs_1_create.json)

#delete Cluster
if atlas clusters delete "$clusterName" --force
then
    echo "$clusterName cluster deletion OK"
else
    (echo "Failed cleaning cluster:$clusterName" && exit 1)
fi


#delete project
if atlas projects delete "$projectId" --force
then
    echo "$projectId project deletion OK"
else
    (echo "Failed cleaning project:$projectId" && exit 1)
fi

