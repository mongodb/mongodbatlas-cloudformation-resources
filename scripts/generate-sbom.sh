#!/usr/bin/env bash
set -euo pipefail

if [ $# -ne 2 ]; then
	echo "Usage: $0 <resource-name> <version>"
	exit 1
fi

RESOURCE="$1"
VERSION="$2"
COMPLIANCE_DIR="cfn-resources/${RESOURCE}/compliance"

if [ ! -d "$COMPLIANCE_DIR" ]; then
	echo "Compliance directory not found: $COMPLIANCE_DIR"
	exit 1
fi

PURL_FILE="${COMPLIANCE_DIR}/purls.txt"
SBOM_FILE="${COMPLIANCE_DIR}/v${VERSION}/sbom.json"

# Ensure the output directory exists
mkdir -p "$(dirname "$SBOM_FILE")"

echo "Generating SBOM for resource: $RESOURCE..."
docker run --rm \
	-v "$PWD:/pwd" \
	"$SILKBOMB_IMG" \
	update \
	--purls "/pwd/${PURL_FILE}" \
	--sbom-out "/pwd/${SBOM_FILE}"
