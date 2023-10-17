#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.

set -o errexit
set -o nounset
set -o pipefail

function usage {
    echo "usage:$0 <project/cluster_name>"
    echo "Creates a new project and cluster by that name for the test"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$1" == help ]]; then usage; fi

# Set profile
profile="default"
if [ -n "${MONGODB_ATLAS_PROFILE:-}" ]; then
    echo "profile set to ${MONGODB_ATLAS_PROFILE}"
    profile=${MONGODB_ATLAS_PROFILE}
fi

# Initialize variables to store VPC ID, subnet ID 1, and subnet ID 2
vpc_id=""
subnet_id_1=""
subnet_id_2=""

# Get a list of VPCs using the AWS CLI
vpcs=$(aws ec2 describe-vpcs --query "Vpcs[*].[VpcId]" --output text)

# Loop through each VPC
for vpc_id in $vpcs; do
    # Get the number of subnets in the VPC
    subnet_count=$(aws ec2 describe-subnets --filters "Name=vpc-id,Values=$vpc_id" --query "length(Subnets)" --output text)

    # Check if the VPC has at least 2 subnets
    if [ "$subnet_count" -ge 2 ]; then
        # Get the IDs of the first two subnets
        subnet_ids=$(aws ec2 describe-subnets --filters "Name=vpc-id,Values=$vpc_id" --output json)

        # Assign VPC ID, subnet ID 1, and subnet ID 2 to variables
        subnet_id_1=$(echo "$subnet_ids" | jq -r '.Subnets[0].SubnetId')
        subnet_id_2=$(echo "$subnet_ids" | jq -r '.Subnets[1].SubnetId')

        # Exit the loop once we've found a matching VPC
        break
    fi
done

rm -rf inputs
mkdir inputs

if ! test -v AWS_DEFAULT_REGION; then
	region=$(aws configure get region)
else
	region=$AWS_DEFAULT_REGION
fi

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

check_status() {
  endpoint_id=$1
  project_id=$2
  status=$(atlas privateEndpoints describe "$endpoint_id" --projectId "$project_id" --output json | jq -r '.[0].status')
  echo "$status"
}

# Run the AWS command and capture the output
output=$(atlas privateEndpoints aws list --projectId "${projectId}" --output json)
private_endpoint_id=""
# Check if the output is empty
if [ "$(echo "$output" | jq -e '. | length == 0')" = true ]; then
  echo "Empty"
  # Execute the create command if the output is empty
  create_output=$(atlas privateEndpoints aws create --region "${region}" --projectId "${projectId}" --output json)
  private_endpoint_id=$(echo "$create_output" | jq -r '.id')
  echo "Created endpoint with ID: $private_endpoint_id"

  # Poll and wait for the status to become "AVAILABLE"
  while true; do
    status=$(atlas privateEndpoints aws describe "$private_endpoint_id" --projectId "$projectId" --output json | jq -r '.status')
    if [ "$status" = "AVAILABLE" ]; then
      echo "Status: $status"
      break
    fi
    echo "Status: $status (waiting for AVAILABLE)"
    sleep 5
  done
else
  # Use jq to extract the ID of the first result
    private_endpoint_id=$(echo "$output" | jq -r '.[0].id')
    echo "ID: $private_endpoint_id"
fi

endpoint_service_id=$(atlas privateEndpoints aws describe "$private_endpoint_id" --projectId "$projectId" --output json | jq -r '.endpointServiceName')

aws_private_endpoint_id=$(aws ec2 create-vpc-endpoint \
  --vpc-id "$vpc_id" \
  --service-name "$endpoint_service_id" \
  --region "$region" \
  --subnet-ids "$subnet_id_1" \
  --vpc-endpoint-type Interface \
  --output json | jq -r '.VpcEndpoint.VpcEndpointId')

atlas privateendpoints aws interfaces create "$private_endpoint_id" --privateEndpointId "$aws_private_endpoint_id" --projectId "$projectId"

aws_private_endpoint_id2=$(aws ec2 create-vpc-endpoint \
  --vpc-id "$vpc_id" \
  --service-name "$endpoint_service_id" \
  --region "$region" \
  --subnet-ids "$subnet_id_2" \
  --vpc-endpoint-type Interface \
  --output json | jq -r '.VpcEndpoint.VpcEndpointId')


jq --arg projectId "$projectId" \
  --arg endpointServiceId "$private_endpoint_id" \
   --arg interfaceEndpointId "$aws_private_endpoint_id2" \
   --arg profile "$profile" \
   '.Profile?|=$profile | .ProjectId?|=$projectId  | .EndpointServiceId?|=$endpointServiceId | .InterfaceEndpointId?|=$interfaceEndpointId ' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"
