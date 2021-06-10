#!/usr/bin/env bash

RESOURCE_TYPES="${1}"
REGIONS="${2:-us-east-1}"

echo "cleaning resource types..."
for r in $REGIONS
do
  for t in $RESOURCE_TYPES
  do
    for v in $(aws cloudformation list-type-versions --type-name $t --type RESOURCE --query TypeVersionSummaries[].Arn --output text --region $r 2>/dev/null || true)
    do
      echo "Deregistering: region=${r} type=${t} version=${v}"
      aws cloudformation deregister-type --arn $v --region $r || \
        aws cloudformation deregister-type --type RESOURCE --type-name $t  --region $r
    done
  done
done

