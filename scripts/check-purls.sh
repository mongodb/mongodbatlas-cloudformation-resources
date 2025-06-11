#!/usr/bin/env bash
set -euo pipefail

if ! git diff --quiet --exit-code cfn-resources/**/compliance/purls.txt; then
	echo "cfn-resources/**/compliance/purls.txt is out of date. Please run 'make gen-purls' and commit the result."
	git --no-pager diff cfn-resources/**/compliance/purls.txt
	exit 1
fi
