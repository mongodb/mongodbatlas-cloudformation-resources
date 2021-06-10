#!/usr/bin/env bash

echo "# DO NOT RUN THIS UNLESS YOU REALLY REALLY REALLY KNOW WHAT YOU ARE DOING"
echo "# Output script to remove all versions of ALL Resource Types"
echo "# ---start---scrub-resource-types-script---start---"
aws cloudformation list-types | jq -r '.TypeSummaries[] | .TypeName' | xargs -I {} echo "./scripts/aws-cfn-resource-type-cleaner.sh {}"
echo "# ---end---scrub-resource-types-script---end---"

