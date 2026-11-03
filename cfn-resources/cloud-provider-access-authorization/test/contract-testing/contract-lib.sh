#!/usr/bin/env bash
# contract-lib.sh
#
# Shared setup for the isolated contract-test scripts: builds the resource,
# starts the local Lambda service, and cleans up the fixtures on exit.

# start_contract_environment prepares the build, the local Lambda service and
# the result directory used by the callers.
start_contract_environment() {
	resultDir=$(mktemp -d)
	samLog="${resultDir}/sam.log"
	samPid=""

	echo "==> Building the resource"
	make build

	echo "==> Starting the local Lambda service"
	sam local start-lambda >"${samLog}" 2>&1 &
	samPid=$!
	for _ in {1..60}; do
		if curl -s -o /dev/null http://127.0.0.1:3001/; then
			break
		fi
		sleep 2
	done
	if ! curl -s -o /dev/null http://127.0.0.1:3001/; then
		echo "The local Lambda service did not start" >&2
		tail -n 40 "${samLog}" >&2
		exit 1
	fi
}

# stop_contract_environment stops the Lambda service and removes the project,
# the registrations and the IAM roles created by the callers.
stop_contract_environment() {
	if [ -n "${samPid}" ]; then
		kill "${samPid}" 2>/dev/null || true
	fi
	./test/contract-testing/cfn-test-delete.sh || true
	rm -rf "${resultDir}"
}

# run_contract_cases runs the selected contract cases and writes the output to
# the given log file and the JUnit report to the given path.
run_contract_cases() {
	local selector="$1" log="$2" junit="$3"
	cfn test --function-name TestEntrypoint --verbose -- -k "${selector}" --junitxml="${junit}" >"${log}" 2>&1
}
