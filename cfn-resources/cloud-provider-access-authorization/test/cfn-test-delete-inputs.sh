#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the AWS resources created for `cfn test`. The Atlas role
# registration is removed by the resource's delete handler, and the project is
# removed by the caller.

set -o errexit
set -o nounset
set -o pipefail

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
