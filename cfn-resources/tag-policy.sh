#!/usr/bin/env bash

# Cloud Tag Policy: values are required when the test scripts create AWS resources.
require_tag_policy_vars() {
	if [ -z "${MONGODB_TAG_OWNER:-}" ] || [ -z "${MONGODB_TAG_ENV:-}" ]; then
		echo "MONGODB_TAG_OWNER and MONGODB_TAG_ENV must be set (Cloud Tag Policy)" >&2
		exit 1
	fi
}
