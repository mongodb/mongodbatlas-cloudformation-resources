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
pipelineName=$(jq -r '.Name' ./inputs/inputs_1_create.json)
clusterName=$(jq -r '.Source/ClusterName' ./inputs/inputs_1_create.json)
#delete project
if  atlas dataLakePipelines delete "$pipelineName" --force; then
	echo "$pipelineName project deletion OK"
else
	(echo "Failed cleaning datalake pipeline:$pipelineName" && exit 1)
fi

atlas clusters delete "$clusterName" --projectId "${projectId}" --force
atlas clusters watch "${clusterName}" --projectId "${projectId}"
echo -e "Cluster Deleted \"${clusterName}\""

if atlas projects delete "$projectId"  -- force; then
  echo "$projectId is deleted"
else
  (echo "Failed cleaning project:$projectId" && exit 1)
fi





