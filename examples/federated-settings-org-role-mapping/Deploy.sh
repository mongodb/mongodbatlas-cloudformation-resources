#!/usr/bin/env bash

set -x
export AWS_DEFAULT_REGION="ap-northeast-2"
federationSettingsId="63875a930d379520633ed4e8"
orgId="6387590e54d177787a239b4c"
externalGroupName="RoleMappingGroup-02"
groupId="638a46aef75a91632ce26596"
echo "FederatedSettingsOrgRoleMapping creation"
aws cloudformation deploy \
   --stack-name atlas-federated-settings-org-role-mapping \
   --template-file federatedSettingsOrgRoleMapping.json \
    --no-fail-on-empty-changeset \
    --parameter-overrides PrivateKey="${ATLAS_PRIVATE_KEY}" PublicKey="${ATLAS_PUBLIC_KEY}" ExternalGroupName="$externalGroupName" GroupId="$groupId" FederationSettingsId="$federationSettingsId" OrgId="$orgId" \
    "$@"