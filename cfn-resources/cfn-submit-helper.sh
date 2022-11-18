#!/usr/bin/env bash

# cfn-submit-helper.sh
#
#
# This tool works by running `cfn submit` on each resource
# By default it will build and submit every resource found
# in this directory.
# There are some options.
#
# BUILD_ONLY=true|false
# SUBMIT_ONLY=true|false
# LOG_LEVEL=logrus valid string loglevel
#
# Example with DEBUG logging enabled by default for set of resources:
# LOG_LEVEL=debug ./cfn-submit-helper.sh project database-user project-ip-access-list cluster network-peering
#
set -x
set -o errexit
set -o nounset
set -o pipefail

. ./cfn-submit-helper.config
env | grep CFN_SUBMIT_
echo "AWS_DEFAULT_PROFILE=${AWS_DEFAULT_PROFILE}"

_DRY_RUN=${DRY_RUN:-false}
_BUILD_ONLY=${BUILD_ONLY:-false}
_SUBMIT_ONLY=${SUBMIT_ONLY:-false}

[[ "${_DRY_RUN}" == "true" ]] && echo "*************** DRY_RUN mode enabled **************"

# By default, submit the entire library of active custom resources, but usually
# this is run one resource at a time.
resources="${1:-project database-user project-ip-access-list network-peering cluster}"
echo "$(basename "$0") running for the following resources: ${resources}"



echo "Step 1/2: Building"
if [[ "${_SUBMIT_ONLY}" == "true" ]]; then
    echo "SUBMIT_ONLY, skipping build."
else
    for resource in ${resources};
    do
        [[ "${_DRY_RUN}" == "true" ]] && echo "[dry-run] would have run make on:${resource}" && continue
        echo "Working on resource:${resource}"
        cwd=$(pwd)
        cd "${resource}"
        echo "resource: ${resource}"
        if [[ "${CFN_SUBMIT_LOG_LEVEL}" == "debug" ]]; then
            make debug
        else
            make
        fi
        cat rpdk.log
        cd -
    done
fi
if [[ "${_BUILD_ONLY}" == "true" ]]; then
    echo "BUILD_ONLY true, skipping submit to CloudFormation"
    exit 0
fi
echo "Step 2/2: Submit resource type to CloudFormation Private Registry"
for resource in ${resources};
do
    echo "Working on resource:${resource}"
    [[ "${_DRY_RUN}" == "true" ]] && echo "[dry-run] would have run 'cfn submit' on:${resource}" && continue
    cwd=$(pwd)
    cd "${resource}"
    echo "resource: ${resource}"
    echo "Submiting to CloudFormation with flags: ${CFN_SUBMIT_CFN_FLAGS}"
    cfn submit ${CFN_SUBMIT_CFN_FLAGS}||true
    cat rpdk.log
    cd -
done



