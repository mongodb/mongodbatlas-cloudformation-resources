#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail
set -x

function usage {
	echo "usage:$0 <project_name>"
}

export AWS_DEFAULT_REGION="us-east-1"
if [ -z "${AWS_DEFAULT_REGION+x}" ];then
  echo "AWS_DEFAULT_REGION must be set"
  exit 1
fi

rm -rf inputs
mkdir inputs

#project_id
projectName="pes-online-archive-$((1 + RANDOM % 10000))"

./test/cfn-test-create-inputs.sh $projectName