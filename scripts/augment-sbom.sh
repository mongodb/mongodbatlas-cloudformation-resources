#!/usr/bin/env bash
set -euo pipefail

: "${RELEASE_VERSION:?RELEASE_VERSION environment variable not set}"
: "${RESOURCE:?RESOURCE environment variable not set}"
DATE="${DATE:-$(date +'%Y-%m-%d')}"

echo "Augmenting SBOM..."
docker run \
	--pull=always \
	--platform="linux/amd64" \
	--rm \
	-v "${PWD}:/pwd" \
	-e KONDUKTO_TOKEN \
	"$SILKBOMB_IMG" \
	augment \
	--sbom-in "/pwd/cfn-resources/${RESOURCE}/compliance/v${RELEASE_VERSION}/sbom.json" \
	--repo "$KONDUKTO_REPO" \
	--branch "$KONDUKTO_BRANCH_PREFIX-linux-arm64" \
	--sbom-out "/pwd/cfn-resources/${RESOURCE}/compliance/augmented-sbom-v${RELEASE_VERSION}-${DATE}.json"
