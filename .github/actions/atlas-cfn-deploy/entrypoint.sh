#!/bin/bash
# Exit on error. Append "|| true" if you expect an error.
set -o errexit  # same as -e
# Exit on error inside any functions or subshells.
set -o errtrace
# Do not allow use of undefined vars. Use ${VAR:-} to use an undefined VAR
set -o nounset
# Catch if the pipe fucntion fails
set -o pipefail
set -x

echo "#############################################################"
env
cd cfn-resources
BUILD_ONLY=1 CFN_FLAGS="--verbose --set-default --region ${AWS_DEFAULT_REGION}" ./cfn-submit-helper.sh

