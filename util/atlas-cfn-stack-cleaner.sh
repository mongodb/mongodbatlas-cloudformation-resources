#!/usr/bin/env bash
REGION="${1:-us-east-1}"
STACKS=$(aws cloudformation describe-stacks --region ${REGION} --output text --query 'Stacks[*].{Stack:StackName}')
if [[ "$*" == *dry-run* ]]
then
    echo "dry-run"
    echo "Region: ${REGION} Stacks: ${STACKS}"
    exit 0
fi
# Start output structured YAML 'log'
echo "\"atlas-cfn-stack-cleaner\":"
echo "-"
echo "  \"message\": \"WARNING: This will blow-away all the stacks in a region. Swim at your own risk.\","
echo "  \"ts\": \"$(date --iso-8601=seconds)\", "
echo "  \"stacks\":"

while IFS= read -r stack
do
    term_resp=$(aws cloudformation update-termination-protection \
    --no-enable-termination-protection \
    --region "${REGION}" --stack-name "$stack")

    delete_resp=$(aws cloudformation delete-stack --region "${REGION}" --stack-name "$stack")

    echo "  -  \"stack\": "${stack}""
    echo "     \"update-termination-protection\": "${tr}""
    echo "     \"delete-stack\": "${ds}""
done < <(printf '%s\n' "${STACKS}")
