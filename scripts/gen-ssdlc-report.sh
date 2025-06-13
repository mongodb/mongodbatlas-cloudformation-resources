#!/usr/bin/env bash
set -euo pipefail

release_date=${DATE:-$(date -u '+%Y-%m-%d')}

if [ $# -ne 2 ]; then
	echo "Usage: $0 <resource-name> <version>"
	exit 1
fi

RESOURCE="$1"
VERSION="$2"

export DATE="${release_date}"

# Debug logging
echo "PWD: $(pwd)"
echo "Listing templates directory:"
ls -l templates || echo "templates directory not found"
echo "Listing cfn-resources directory:"
ls -l cfn-resources || echo "cfn-resources directory not found"

echo "AUTHOR: ${AUTHOR}"

if [ "${AUGMENTED_REPORT:-false}" = "true" ]; then
	target_dir="."
	file_name="ssdlc-compliance-${RESOURCE}-${VERSION}-${DATE}.md"
	SBOM_TEXT="  - See Augmented SBOM manifests (CycloneDX in JSON format):
      - This file has been provided along with this report under the name 'linux_amd64_augmented_sbom_v${VERSION}.json'
      - Please note that this file was generated on ${DATE} and may not reflect the latest security information of all third party dependencies."

else # If not augmented, generate the standard report
	target_dir="cfn-resources/${RESOURCE}/compliance/v${VERSION}"
	file_name="ssdlc-compliance-${RESOURCE}-${VERSION}.md"
	SBOM_TEXT="  - See SBOM Lite manifests (CycloneDX in JSON format):
      - https://github.com/mongodb/mongodbatlas-cloudformation-resources/cfn-resources/${RESOURCE}/compliance/v${VERSION}/sbom.json"
	# Ensure terraform-provider-mongodbatlas version directory exists
	mkdir -p "${target_dir}"
	echo "target_dir: ${target_dir}"
	echo "file_name: ${file_name}"
	echo "SBOM_TEXT: ${SBOM_TEXT}"
fi

export AUTHOR
export VERSION
export SBOM_TEXT

echo "Generating SSDLC report for CloudFormation MongoDB::Atlas::${RESOURCE} version ${VERSION}, author ${AUTHOR} and release date ${DATE}..."

envsubst <templates/ssdlc-compliance.template.md \
	>"${target_dir}/${file_name}"

echo "SSDLC compliance report ready. Files in ${target_dir}/:"
ls -l "${target_dir}/"

echo "Printing the generated report:"
cat "${target_dir}/${file_name}"
