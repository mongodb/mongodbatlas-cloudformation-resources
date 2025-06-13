#!/usr/bin/env bash
set -euo pipefail

if [ $# -ne 2 ]; then
	echo "Usage: $0 <resource-name> <version>"
	exit 1
fi

RESOURCE="$1"
echo "Uploading SBOMs for resource: $RESOURCE..."
docker run --rm \
	-v "$PWD:/pwd" \
	-e KONDUKTO_TOKEN \
	"$SILKBOMB_IMG" \
	upload \
	--sbom-in "/pwd/cfn-resources/${RESOURCE}/compliance/v${VERSION}/sbom.json" \
	--repo "$KONDUKTO_REPO" \
	--branch "$KONDUKTO_BRANCH_PREFIX-${RESOURCE}-linux-arm64"
