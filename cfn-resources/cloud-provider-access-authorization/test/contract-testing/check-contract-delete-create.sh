#!/usr/bin/env bash

# check-contract-delete-create.sh
#
# Runs contract_delete_create on its own fixtures and reports it separately.
# The case cannot pass: Atlas has no operation that revokes an authorization
# while keeping the role registered, so re-creating after a delete fails with
# NotFound. This script verifies that exact signature and stays red until AWS
# accepts the exception (CLOUDP-465505). Flip EXCEPTION_ACCEPTED only in the
# commit that records the acceptance.

set -o errexit
set -o nounset
set -o pipefail

scriptDir=$(cd "$(dirname "$0")" && pwd)
cd "${scriptDir}/../.."
# shellcheck source=/dev/null
source "${scriptDir}/contract-lib.sh"
trap stop_contract_environment EXIT
start_contract_environment

# The AWS exception for the delete-then-create deviation (CLOUDP-465505).
EXCEPTION_ACCEPTED="false"

caseName="contract_delete_create"
log="${resultDir}/contract.log"
junit="${resultDir}/junit.xml"

echo "==> Preparing a fresh registration and IAM roles"
./test/cfn-test-create-inputs.sh "${projectName}"

run_contract_cases "${caseName}" || true

# The signature is a test failure (not a fixture setup error) plus the
# registration-not-found error the re-create produces. The helper rejects a
# setup error, an unexpected pass and any additional failing case, and the
# handler log check proves the failure came from the missing registration.
junitResult=""
if junitResult=$(./test/contract-testing/contract_junit.py --expect-failure "${caseName}" "${junit}" 2>&1) &&
	grep -q "CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND" "${log}" "${samLog}"; then
	if [ "${EXCEPTION_ACCEPTED}" = "true" ]; then
		echo "==> contract_delete_create reproduced the accepted deviation: NotFound on re-create after a successful create and delete"
		exit 0
	fi
	echo "==> contract_delete_create reproduced the known deviation: NotFound on re-create after a successful create and delete"
	echo "==> The AWS exception is still pending (CLOUDP-465505); this step stays red until it is accepted"
	exit 1
fi

echo "FAIL: contract_delete_create did not reproduce the expected failure (NotFound on re-create after a successful create and delete)" >&2
printf '%s\n' "${junitResult}" >&2
cat "${log}" >&2
tail -n 60 "${samLog}" >&2
exit 1
