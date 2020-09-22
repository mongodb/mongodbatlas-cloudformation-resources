#!/usr/bin/env bash
REGION="${1:-us-east-1}"
STACKS=$(aws cloudformation describe-stacks --region ${REGION} --output text --query 'Stacks[*].{Stack:StackName}'  | grep mongodb-atlas | grep role-stack)
if [[ "$*" == *killall* ]]
then
    echo "*********** killall initiated ******************"
    STACKS=$(aws cloudformation describe-stacks --region ${REGION} --output text --query 'Stacks[*].{Stack:StackName}')
    echo "Region: ${REGION} Stacks: ${STACKS}"
fi
if [[ "$*" == *dry-run* ]]
then
    echo "dry-run"
    echo "Region: ${REGION}"
    echo "Stacks: ${STACKS}"
    exit 0
fi
if [[ -z "${STACKS}" ]]; then
    echo "Skys look clear, no stacks in sight, proceed."
    exit 0
fi

# Start output structured YAML 'log'
echo "\"atlas-cfn-stack-cleaner\":"
echo "-"
echo "  \"message\": \"WARNING: This will blow-away all the stacks in a region. Swim at your own risk.\","
echo "  \"ts\": \"$(date --iso-8601=seconds)\", "

if [[ -z "${STACKS}" ]]; then
    echo "Skys look clear, no stacks in sight, proceed."
else
    echo "  \"stacks\":"
    while IFS= read -r stack
    do
        term_resp=$(aws cloudformation update-termination-protection \
        --no-enable-termination-protection \
        --region "${REGION}" --stack-name "$stack")

        delete_resp=$(aws cloudformation delete-stack --region "${REGION}" --stack-name "$stack")

        echo "  -  \"stack\": \"${stack}\""
        echo "     \"update-termination-protection\": \"${term_resp}\""
        echo "     \"delete-stack\": \"${delete_resp}\""
    done < <(printf '%s\n' "${STACKS}")
fi

