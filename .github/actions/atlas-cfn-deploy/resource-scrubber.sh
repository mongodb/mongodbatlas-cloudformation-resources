#!/usr/bin/env bash
RESOURCE=${1:-${MONGODB_ATLAS_RESOURCE}}
REGION=${2:-${AWS_REGION}}
echo "scrubbing ${RESOURCE} from ${REGION}"
aws cloudformation list-type-versions \
    --type RESOURCE \
    --type-name ${RESOURCE} \
    --region ${REGION} | \
    jq -r '.TypeVersionSummaries[] | .Arn' | \
    xargs -I {} aws cloudformation deregister-type \
    --region ${REGION} \
    --arn {}

aws cloudformation deregister-type \
    --type RESOURCE \
    --type-name ${RESOURCE} \
    --region ${REGION}


