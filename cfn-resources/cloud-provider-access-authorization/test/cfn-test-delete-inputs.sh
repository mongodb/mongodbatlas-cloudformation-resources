#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the AWS and Atlas resources created for `cfn test`. The
# Atlas role registration is usually already removed by the resource's delete
# handler, so its deauthorization tolerates a missing role. The project is
# removed by the caller.

set -o errexit
set -o nounset
set -o pipefail

deleteRole() {
	local roleFile="$1" roleName deleteOutput
	if [ ! -f "${roleFile}" ]; then
		echo "No $(basename "${roleFile}") found, nothing to delete"
		return
	fi
	roleName=$(cat "${roleFile}")
	echo "Deleting AWS IAM role: ${roleName}"
	if ! deleteOutput=$(aws iam delete-role --role-name "${roleName}" 2>&1); then
		if [[ "$deleteOutput" != *"(NoSuchEntity)"* ]]; then
			printf '%s\n' "$deleteOutput" >&2
			exit 1
		fi
	fi
	rm -f "${roleFile}"
}

deleteRole "$(dirname "$0")/role-name.txt"
deleteRole "$(dirname "$0")/update-role-name.txt"

if [ -f inputs/inputs_1_create.json ]; then
	projectId=$(jq -r '.ProjectId // empty' inputs/inputs_1_create.json)
	roleId=$(jq -r '.RoleId // empty' inputs/inputs_1_create.json)
	if [ -n "${projectId}" ] && [ -n "${roleId}" ]; then
		if ! deauthorizeOutput=$(atlas cloudProviders accessRoles aws deauthorize "${roleId}" --projectId "${projectId}" --force 2>&1); then
			printf 'Registration %s was already removed: %s\n' "${roleId}" "${deauthorizeOutput}" >&2
		fi
	fi
fi

rm -f "$(dirname "$0")/trust-policy.json"
