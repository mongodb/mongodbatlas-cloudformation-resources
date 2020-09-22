#!/usr/bin/env bash

# stacker.sh
#
# This script knows how to run the "unit-test" example stack for
# each custom resource.
# Each resource has a file called 'stack.yaml' which contains
# the canonical example of how to use the given Atlas resource
# in a CloudFormation stack.
#

set -o errexit
set -o nounset
set -o pipefail
source ./log.sh

log_info "stacker.sh hello"
region="${AWS_REGION}"
log_info "region: ${region}"

# Default, find all the directory names with the json custom resource schema files.
# Or you pass them in.
resources="${1:-$(ls -F **/mongodb-atlas-*.json | cut -d/ -f1)}"
log_info "resources: ${resources}"

tag=$(openssl rand -hex 5)
log_info "tag: ${tag}"

for resource in ${resources};
do
    log_info "resource: ${resource}"
    validate=$(aws cloudformation validate-template --template-body=file://./${resource}/stack.yaml 2>&1)
    log_info "validate: ${validate}"
    cs_resp=$(aws cloudformation create-stack \
        --template-body=file://./${resource}/stack.yaml \
        --disable-rollback \
        --region ${region} \
        --stack-name ${resource}-test-${tag})
    log_info "create-stack-response: ${cs_resp}"
done
