#!/usr/bin/env bash
# cloud-backup-snapshot-export-job.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:1 <project_name>"
	echo "usage:2 <cluster_name>"

}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi
projectId="${1}"
clusterName="${2}"
jq --arg org "$MONGODB_ATLAS_ORG_ID" \
	--arg projectId "$projectId" \
	--arg clusterName "$clusterName" \
	'.desiredResourceState.ClusterName?|=$clusterName  | .desiredResourceState.ProjectId?|=$projectId ' \
	"$(dirname "$0")/global-cluster-config.sample-cfn-request.json"
