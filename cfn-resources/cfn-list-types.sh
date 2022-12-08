#!/usr/bin/env bash

resources="${1:-project}"
regions="us-east-1 us-west-2 ca-central-1 us-east-2 us-west-1 sa-east-1 ap-southeast-1 ap-southeast-2 ap-southeast-3 ap-south-1 ap-east-1 ap-northeast-1 ap-northeast-2 ap-northeast-3 eu-west-1 eu-central-1 eu-north-1 eu-west-2 eu-west-3 eu-south-1 me-south-1 af-south-1"

for resource in ${resources};
do
  echo "Step: Running 'list-type' on ${resource}"
  for region in ${regions};
  do
    cd "${resource}"
    jsonschema="mongodb-atlas-$(echo ${resource}| sed s/-//g).json"
    type_name=$(cat ${jsonschema}| jq -r '.typeName')
    type_info=$(aws cloudformation --region "${region}"  list-types --visibility PUBLIC --output=json | jq --arg typeName "${type_name}" '.TypeSummaries[] | select(.TypeName==$typeName)')
    if [ -z "${type_info}" ]; then
      echo "*********** ${resource} type NOT found in region : ${region} *******************"
    else
#      lastUpdated=$(type_info | jq -r '.LastUpdated')
#      if [ "${lastUpdated}" -le '2022-12-02' ];then
         echo "${resource} found in region : ${region}"
#      fi
    fi
    cd ../
  done
done





