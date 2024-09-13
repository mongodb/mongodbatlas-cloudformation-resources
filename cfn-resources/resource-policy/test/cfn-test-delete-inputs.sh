#!/usr/bin/env bash
# cfn-test-delete-inputs.sh
#
# This tool deletes the mongodb resources used for `cfn test` as inputs.
#

set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 "
}
