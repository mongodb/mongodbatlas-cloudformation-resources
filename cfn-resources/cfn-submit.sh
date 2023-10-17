#!/usr/bin/env bash

set -eo

export CLOUD_PUBLISH=true

resources="${1:-project}"
regions="${2:-us-east-1 }"
#resources=${1:project database-user network-peering network-container project-ip-access-list cloud-backup-snapshot cloud-backup-restore-jobs encryption-at-rest cluster private-endoint}
#regions="${1:-ap-northeast-2 us-east-1 us-west-2 ca-central-1 us-east-2 us-west-1 sa-east-1}"

echo "$(basename "$0") running for the following resources: ${resources}"

# Deploy in given regions
for resource in ${resources}; do
	for region in ${regions}; do

		export AWS_DEFAULT_REGION="${region}"

		echo "Step 1: cfn test"
		./cfn-testing-helper.sh "${resource}"

		echo "step 2: cfn submit for ${resource}"
		./cfn-submit-helper.sh "${resource}"

		echo " step 3: update default version for ${resource}"

		cd "${resource}"
		pwd
		# shellcheck disable=SC2001
		jsonschema="mongodb-atlas-$(echo "${resource}" | sed s/-//g).json"
		res_type=$(jq -r '.typeName' "${jsonschema}")
		echo "${res_type}"
		cd -

		echo "******** Successfully submitted ${resource} *************"

	done
done
