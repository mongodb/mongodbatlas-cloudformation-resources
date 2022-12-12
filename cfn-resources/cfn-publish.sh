#!/usr/bin/env bash


#trap "exit" INT TERM ERR
#set -x
#set -o errexit
set -o nounset
#set -o pipefail


resources="${1:-project}"


# Handling other parameters if there are any.
IFS=","
if [ -n "${2}" ]; then
for param in ${2};
do
echo "$param"
export "$param"
echo
done
fi

cloud_publish=${3:-true}

export CLOUD_PUBLISH="${cloud_publish}"

echo "CLOUD_PUBLISH : ${CLOUD_PUBLISH}"


for resource in ${resources};
do
    . ./cfn-testing-helper.config
    env | grep CFN_TEST_

    PROJECT_NAME="${CFN_TEST_NEW_PROJECT_NAME}"
    echo "PROJECT_NAME:${PROJECT_NAME}"

    echo " Started Publishing ${resource} resource"
    echo "Step 1: cfn test"

    if ! ./cfn-testing-helper.sh "${resource}" "${PROJECT_NAME}";
      then
       	echo "Error in Testing phase"
        exit 1
    fi
    exit 0

    echo "step 2: cfn submit for ${resource}"

     if ! ./cfn-submit-helper.sh "${resource}";
      then
        echo "Error in Submit phase"
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

    if ! ./cfn-publishing-helper.sh "${resource}" "${latestVersion}";
                       then
                        	echo "Error in Publishing phase"
                         exit 1
    fi

    #delete the input params
    cd "${resource}"
    ./test/cfn-test-delete-inputs.sh "${PROJECT_NAME}-${resource}" && echo "resource:${resource} inputs delete OK" || echo "resource:${resource} input delete FAILED"

#    Deleting the projects
    echo "Looking up Atlas project id for resource:${resource} project name:${PROJECT_NAME}-${resource}"
    p_id=$(atlas project list --output=json | jq --arg name "${PROJECT_NAME}-${resource}" -r '.results[] | select(.name==$name) | .id')
    [ -z "$p_id" ] && echo "No project found" && continue
    p_name=$(atlas project list --output=json | jq --arg name "${PROJECT_NAME}-${resource}" -r '.results[] | select(.name==$name) | .name')
    echo "Cleaning up for resource:${resource}, project:${p_name} id:${p_id}"
    atlas project delete "${p_id}" --force && echo "Cleaned up project:${p_name} id:${p_id}" || (echo "Failed cleaning up project:${p_id}" && exit 1)

    echo "******** Successfully published ${resource} *************"
 done
