#!/usr/bin/env bash

set -x


export CLOUD_PUBLISH=true


resources="${1:-project}"
regions="${1:-ap-northeast-2 }"
#regions="${1:-ap-northeast-2 us-east-1 us-west-2 ca-central-1 us-east-2 us-west-1 sa-east-1}"
echo "$(basename "$0") running for the following resources: ${resources}"

#create a dir for submit outputs

echo "Step 1: cfn test"
# Deploy in given regions
for region in ${regions}
do
  # Deploy given resources
  for resource in ${resources};
  do
    AWS_DEFAULT_REGION=$region

    ./cfn-testing-helper.sh "${resource}"

    AWS_DEFAULT_REGION=$region

    echo "cfn submit for ${resource}"
    ./cfn-submit-helper.sh "${resource}"

    echo "cfn publish for ${resource}"
    cd "${resource}"
    pwd
    jsonschema="mongodb-atlas-$(echo ${resource}| sed s/-//g).json"
    res_type=$(cat ${jsonschema}| jq -r '.typeName')
    echo "${res_type}"
    cd -

    sleep 5
    aws cloudformation list-type-versions --type RESOURCE --type-name res_type
    latestVersion=$(aws cloudformation list-type-versions --type RESOURCE --type-name "${res_type}" | jq -r '.TypeVersionSummaries[-1].VersionId')

    echo "Set default version to ${latestVersion} "
    aws cloudformation set-type-default-version --type RESOURCE --type-name "${res_type}" --version-id "${latestVersion}"
    # Publish latest version
    ./cfn-publishing-helper.sh "${resource}" "${latestVersion}"
  done
done
