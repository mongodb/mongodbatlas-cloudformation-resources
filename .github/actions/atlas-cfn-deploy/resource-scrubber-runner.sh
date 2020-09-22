#!/usr/bin/env bash
cd ~/work/mongodbatlas-cloudformation-resources
schemas=$(ls cfn-resources/**/mongodb*.json)
region="${1:-us-east-1}"
#echo "schemas=${schemas}"
while IFS= read -r schema
do
    type_name=$(cat $schema | jq -r '.typeName')
    response=$(.github/actions/atlas-cfn-deploy/resource-scrubber.sh ${type_name} ${region} 2>&1 >/dev/null)
    echo "  -  \"schema\": \"${schema}\""
    echo "     \"type_name\": \"${type_name}\""
    echo "     \"response\": \"${response}\""

done < <(printf '%s\n' "${schemas}")

