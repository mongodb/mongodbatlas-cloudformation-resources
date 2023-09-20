#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 "
}


awsSecretName=$(jq -r '.AwsSecretName' ./inputs/inputs_1_create.json)
if aws secretsmanager delete-secret --secret-id "${awsSecretName}" --force-delete-without-recovery;then
  echo "aws secret deleted with name : ${awsSecretName}"
else
  echo "aws secret delete failed with name : ${awsSecretName}"
  exit 1
fi
