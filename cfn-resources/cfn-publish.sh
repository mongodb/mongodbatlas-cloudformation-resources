#!/usr/bin/env bash

#set -x

trap "exit" INT TERM ERR
set -x
set -o errexit
set -o nounset
set -o pipefail


resources="${1:-project}"
cloud_publish=${2:-true}
export CLOUD_PUBLISH="${cloud_publish}"

echo "CLOUD_PUBLISH : ${CLOUD_PUBLISH}"

#regions="${2:-ap-northeast-2 }"
#resources=${1:project database-user network-peering network-container project-ip-access-list cloud-backup-snapshot cloud-backup-restore-jobs encryption-at-rest cluster private-endoint}
#regions="${1:-us-east-1 us-west-2 ca-central-1 us-east-2 us-west-1 sa-east-1 ap-southeast-1 ap-southeast-2 ap-southeast-3 ap-south-1 ap-east-1 ap-northeast-1 ap-northeast-2 ap-northeast-3
#              eu-west-1 eu-central-1 eu-north-1 eu-west-2 eu-west-3 eu-south-1 me-south-1 af-south-1}"

# Deploy in given regions
for resource in ${resources};
do

    echo " Started Publishing ${resource} resource"
    echo "Step 1: cfn test"
    ./cfn-testing-helper.sh "${resource}"

    echo "step 2: cfn submit for ${resource}"
    ./cfn-submit-helper.sh "${resource}"

    echo " step 3: update default version for ${resource}"

    cd "${resource}"
    pwd
    jsonschema="mongodb-atlas-$(echo ${resource}| sed s/-//g).json"
    res_type=$(cat ${jsonschema}| jq -r '.typeName')
    echo "${res_type}"
    cd -

    # keeping 5s sleep to get the updated submit version (may not be required)
    sleep 5

    # get latest submit version of the resource
    latestVersion=$(aws cloudformation list-type-versions --type RESOURCE --type-name "${res_type}" | jq -r '.TypeVersionSummaries[-1].VersionId')

    echo "Setting default version to ${latestVersion} "
    aws cloudformation set-type-default-version --type RESOURCE --type-name "${res_type}" --version-id "${latestVersion}"

    echo " step 4:  Publishing  ${resource}"
    ./cfn-publishing-helper.sh "${resource}" "${latestVersion}"

    echo "******** Successfully published ${resource} *************"
 done
#done
