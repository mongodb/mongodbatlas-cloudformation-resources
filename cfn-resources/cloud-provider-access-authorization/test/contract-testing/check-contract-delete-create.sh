#!/usr/bin/env bash

# check-contract-delete-create.sh
#
# Runs contract_delete_create on its own fixtures and reports it separately.
# The case cannot pass: Atlas has no operation that revokes an authorization
# while keeping the role registered, so re-creating after a delete fails with
# NotFound. This step tolerates exactly that expected outcome, so the check
# stays green for the known deviation, and fails on any other outcome: a
# fixture setup error, a different failure, or another failing case. The AWS
# exception for the publishing test step covers the same deviation.

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

caseName="contract_delete_create"
log="${resultDir}/contract.log"
junit="${resultDir}/junit.xml"

echo "==> Preparing a fresh registration and IAM roles"
./test/cfn-test-create-inputs.sh "${projectName}"

run_contract_cases "${caseName}" "${log}" "${junit}" || true

outcome=$(./test/contract-testing/contract_junit.py --outcome "${caseName}" "${junit}" 2>/dev/null || echo missing)
case "${outcome}" in
passed)
	echo "==> contract_delete_create passed; the known deviation no longer reproduces and this check can be removed"
	exit 0
	;;
failed)
	# The signature is a test failure (not a fixture setup error) plus the
	# registration-not-found error the re-create produces. The helper rejects a
	# setup error, an unexpected pass and any additional failing case, and the
	# handler log check proves the failure came from the missing registration.
	if ./test/contract-testing/contract_junit.py --expect-failure "${caseName}" "${junit}" >/dev/null 2>&1 &&
		grep -q "CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND" "${log}" "${samLog}"; then
		echo "==> contract_delete_create reproduced the known deviation and is tolerated: NotFound on re-create after a successful create and delete"
		exit 0
	fi
	echo "FAIL: contract_delete_create failed in an unexpected form; expected NotFound on re-create after a successful create and delete" >&2
	;;
*)
	echo "FAIL: contract_delete_create ended with an unexpected outcome: ${outcome}" >&2
	;;
esac

echo "==> outcome: ${outcome}" >&2
cat "${log}" >&2
tail -n 60 "${samLog}" >&2
exit 1
