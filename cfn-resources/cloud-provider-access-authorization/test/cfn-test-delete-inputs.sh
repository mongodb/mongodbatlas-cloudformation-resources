#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the AWS resources created for `cfn test`. The Atlas role
# registration is removed by the resource's delete handler, and the project is
# removed by the caller.

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_name>"
}

if [ "$#" -ne 1 ]; then usage; fi

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
echo "Check if a project is created ${projectId}"

if [ -f "$(dirname "$0")/role-name.txt" ]; then
	roleName=$(cat "$(dirname "$0")/role-name.txt")
	echo "Deleting AWS IAM role: ${roleName}"
	if ! deleteOutput=$(aws iam delete-role --role-name "${roleName}" 2>&1); then
		if [[ "$deleteOutput" != *"(NoSuchEntity)"* ]]; then
			printf '%s\n' "$deleteOutput" >&2
			exit 1
		fi
	fi
	rm -f "$(dirname "$0")/role-name.txt"
else
	echo "No role-name.txt found, nothing to delete"
fi

rm -f "$(dirname "$0")/trust-policy.json"
