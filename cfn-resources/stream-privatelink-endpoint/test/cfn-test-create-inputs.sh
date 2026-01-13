#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <project_name>"
	echo "Generates test input files for stream private link endpoint"
	exit 0
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

#set profile - relevant for contract tests which define a custom profile
profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')
	echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

# Determine region (consistent with other resources)
region="${AWS_DEFAULT_REGION:-}"
if [ -z "$region" ]; then
	region=$(aws configure get region 2>/dev/null || echo "")
fi
if [ -z "$region" ]; then
	region="${AWS_REGION:-eu-west-1}"
fi
echo "Using region: ${region}"
echo "Note: S3 bucket creation is not required. Using standard AWS S3 service endpoint format: com.amazonaws.${region}.s3"

# Confluent Cloud configuration (from environment variables, similar to Terraform)
confluentRegion="${CONFLUENT_CLOUD_REGION:-us-east-1}"
confluentDnsDomain="${CONFLUENT_CLOUD_DNS_DOMAIN:-dom4gllez7g.us-east-1.aws.confluent.cloud}"
confluentServiceEndpointId="${CONFLUENT_CLOUD_SERVICE_ENDPOINT_ID:-com.amazonaws.vpce.us-east-1.vpce-svc-09f77bf9637bb0090}"
confluentDnsSubDomain="${CONFLUENT_CLOUD_DNS_SUB_DOMAIN:-use1-az1.dom4gllez7g.us-east-1.aws.confluent.cloud,use1-az2.dom4gllez7g.us-east-1.aws.confluent.cloud,use1-az4.dom4gllez7g.us-east-1.aws.confluent.cloud}"

if [ -n "$confluentDnsDomain" ]; then
	echo "Confluent Cloud configuration found:"
	echo "  Region: $confluentRegion"
	echo "  DNS Domain: $confluentDnsDomain"
	echo "  Service Endpoint ID: $confluentServiceEndpointId"
	echo "  DNS Sub Domain: $confluentDnsSubDomain"
fi

WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*; do
	outputFile=${inputFile//$WORDTOREMOVE/}

	# Check if this is a Confluent template
	if [[ "$inputFile" == *"confluent"* ]]; then
		# Convert comma-separated DNS subdomains to JSON array
		if [ -n "$confluentDnsSubDomain" ]; then
			dnsSubDomainArray=$(echo -n "$confluentDnsSubDomain" | tr ',' '\n' | jq -R . | jq -s .)
		else
			dnsSubDomainArray="[]"
		fi

		jq --arg projectId "$projectId" \
			--arg profile "$profile" \
			--arg region "$confluentRegion" \
			--arg dnsDomain "$confluentDnsDomain" \
			--arg serviceEndpointId "$confluentServiceEndpointId" \
			--argjson dnsSubDomain "$dnsSubDomainArray" \
			'.Profile?|=$profile | .ProjectId?|=$projectId | .Region?|=$region | .DnsDomain?|=$dnsDomain | .ServiceEndpointId?|=$serviceEndpointId | .DnsSubDomain?|=$dnsSubDomain' \
			"$inputFile" >"../inputs/$outputFile"
	else
		# S3 template
		jq --arg projectId "$projectId" \
			--arg profile "$profile" \
			--arg region "$region" \
			'.Profile?|=$profile | .ProjectId?|=$projectId | .Region?|=$region | .ServiceEndpointId?|="com.amazonaws." + $region + ".s3"' \
			"$inputFile" >"../inputs/$outputFile"
	fi
done
cd ..
ls -l inputs
