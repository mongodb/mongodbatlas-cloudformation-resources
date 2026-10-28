#!/usr/bin/env bash

# project-resolution.sh
#
# Shared project resolution for the contract-test input scripts. Those scripts run
# right after a project creation (contract testing creates the project first) and
# Atlas project lists are eventually consistent, so a single lookup can miss the
# new project. The create then fails with HTTP 409, and the lookup is retried with
# a bounded delay before giving up.

# Prints the id of the project with the given name, creating the project when it
# does not exist. Logs the resolution to stderr, and prints an empty string when
# the project cannot be created or found.
resolve_project_id() {
	local name="$1" id
	id=$(atlas projects list --output json | jq --arg NAME "${name}" -r '.results[] | select(.name==$NAME) | .id') || id=""
	if [ -z "${id}" ]; then
		# The create can fail with HTTP 409 when the project already exists but the
		# lookup above missed it, so the id is looked up again with a bounded retry.
		id=$(atlas projects create "${name}" --output=json | jq -r '.id') || id=""
		for _ in $(seq 1 12); do
			if [ -n "${id}" ]; then break; fi
			id=$(atlas projects list --output json | jq --arg NAME "${name}" -r '.results[] | select(.name==$NAME) | .id') || id=""
			if [ -n "${id}" ]; then break; fi
			sleep 5
		done
	fi
	if [ -n "${id}" ]; then
		echo -e "Using project \"${name}\" with id: ${id}\n" >&2
	fi
	echo "${id}"
}
