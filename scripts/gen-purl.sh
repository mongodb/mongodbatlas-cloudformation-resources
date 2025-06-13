#!/bin/bash

set -euo pipefail
: "${LINKER_FLAGS:=}"

if [ $# -ne 1 ]; then
	echo "Usage: $0 <resource-name>"
	exit 1
fi

RESOURCE="$1"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
EXTRACT_PURL_SCRIPT="${SCRIPT_DIR}/extract-purls.sh"

if [ ! -x "$EXTRACT_PURL_SCRIPT" ]; then
	echo "extract-purls.sh not found or not executable"
	exit 1
fi

echo "==> Generating purls"

# Define output and temp files
OUT_DIR="cfn-resources/${RESOURCE}/compliance"
BIN_DIR="${OUT_DIR}/bin"
PURL_ALL="${OUT_DIR}/purls.txt"

# Build and extract for Linux
pushd "cfn-resources/${RESOURCE}/cmd" >/dev/null
GOOS=linux GOARCH=amd64 go build -ldflags "${LINKER_FLAGS}" -o "../compliance/bin"
popd >/dev/null
"$EXTRACT_PURL_SCRIPT" "${BIN_DIR}" "${PURL_ALL}"

# Clean up temp files
rm -f "${BIN_DIR}"
