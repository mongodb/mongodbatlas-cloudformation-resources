#!/usr/bin/env bash


trap "exit" INT TERM ERR
set -x
set -o errexit
set -o nounset
set -o pipefail


resources="${1:-project}"
cloud_publish=${2:-true}
export CLOUD_PUBLISH="${cloud_publish}"

echo "CLOUD_PUBLISH : ${CLOUD_PUBLISH}"

for resource in ${resources};
do

    echo " Started Publishing ${resource} resource"
    echo "Step 1: cfn test"
    ./cfn-testing-helper.sh "${resource}"
    if [ "$?" -ne 0 ]
            then
             	echo "Error in Testing phase" 1>&2
              exit 1
    fi

    echo "step 2: cfn submit for ${resource}"
    ./cfn-submit-helper.sh "${resource}"
    if [ "$?" -ne 0 ]
        then
         	echo "Error in Submit phase" 1>&2
          exit 1
    fi

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
    if [ "$?" -ne 0 ]
    then
     	echo "Error in Publishing phase" 1>&2
      exit 1
    fi

    echo "******** Successfully published ${resource} *************"
 done
