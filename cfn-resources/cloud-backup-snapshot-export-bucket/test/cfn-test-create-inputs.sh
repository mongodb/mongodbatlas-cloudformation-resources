#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

if [ -z ${CLOUD_BACKUP_GROUP_ID:-} ];
 then echo "For cloud provider snapshot you need to provide a valid group Id with an linked role ID and Bucket ID, check atlas docs"
fi

echo "$ATLAS_PUBLIC_KEY"
echo "$ATLAS_PRIVATE_KEY"
echo "$CLOUD_BACKUP_GROUP_ID"
echo "$CLOUD_BACKUP_ROLE_ID"
echo "$CLOUD_BACKUP_BUCKET_NAME"

function usage {
    echo "Creates a new private endpoint role for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg groupId "$CLOUD_BACKUP_GROUP_ID" \
   --arg iamRoleID "$CLOUD_BACKUP_ROLE_ID" \
   --arg bucketName "$CLOUD_BACKUP_BUCKET_NAME" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$groupId | .IamRoleID?|=$iamRoleID | .BucketName?|=$bucketName ' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
   --arg pvtkey "$ATLAS_PRIVATE_KEY" \
   --arg groupId "dsafasdgsdhsdfh" \
   --arg iamRoleID "$CLOUD_BACKUP_ROLE_ID" \
   --arg bucketName "$CLOUD_BACKUP_BUCKET_NAME" \
   '.ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey | .GroupId?|=$groupId | .IamRoleID?|=$iamRoleID | .BucketName?|=$bucketName ' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_invalid.json"

