#!/usr/bin/env bash

# run-contract-cases.sh
#
# Runs the contract suite in isolated groups. Each group gets a fresh Cloud
# Provider Access registration and matching IAM roles, because this resource's
# delete removes the registration and the suite's module fixtures share one
# registration across their cases. contract_delete_create is not run here; see
# check-contract-delete-create.sh.

set -o errexit
set -o nounset
set -o pipefail

scriptDir=$(cd "$(dirname "$0")" && pwd)
cd "${scriptDir}/../.."
# shellcheck source=/dev/null
source "${scriptDir}/contract-lib.sh"

# Assigned by contract-lib.sh.
resultDir=""
samLog=""

projectName="cfn-test-bot-$(date +%s)-${RANDOM}"

trap stop_contract_environment EXIT
start_contract_environment

# Cases in a group are the ones the suite's own module fixture shares one
# resource between; a group never spans modules, so a fresh registration per
# group keeps every group independent.
groups=(
	"contract_create_delete"
	"contract_create_create|contract_create_read"
	"contract_delete_read|contract_delete_update|contract_delete_delete"
	"contract_update_read"
	"contract_update_without_create|contract_check_asserts_work"
)

failures=0
for group in "${groups[@]}"; do
	selector="${group//|/ or }"
	log="${resultDir}/contract.log"
	junit="${resultDir}/junit.xml"

	echo "==> Preparing a fresh registration and IAM roles for: ${group}"
	./test/cfn-test-create-inputs.sh "${projectName}"

	if run_contract_cases "${selector}" "${log}" "${junit}" && ./test/contract-testing/contract_junit.py "${junit}"; then
		echo "PASS: ${group}"
	else
		echo "FAIL: ${group}"
		cat "${log}"
		tail -n 60 "${samLog}"
		failures=$((failures + 1))
	fi

	./test/cfn-test-delete-inputs.sh
done

if [ "${failures}" -ne 0 ]; then
	echo "==> ${failures} contract group(s) failed"
	exit 1
fi

echo "==> All contract groups passed"
