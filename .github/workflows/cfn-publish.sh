#!/usr/bin/env bash

# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -x
set -Eeou pipefail

AWS_SSM_Document_Name="CFN-MongoDB-Atlas-Resource-Register"
BuilderRole="DevOpsIntegrationsContractors-CodeBuild"
AssumeRole="arn:aws:iam::711489243244:role/DevOpsIntegrationsContractorsSSM"
LogDeliveryBucket="atlascfnpublishing"
Repository="https://github.com/mongodb/mongodbatlas-cloudformation-resources"
BranchName="master"
ExecutionRoleName="DevOpsIntegrationsContractorsSSM"
Document_Region="us-east-1"
AccountIds="711489243244"
TargetLocationsMaxConcurrency="30"
Document_Version="\$DEFAULT"

# declare OTHER_PARAMS will be used for CFN TEST input creation.
# This script will update these params for  Few Resources l .
OtherParams="'{\"param\":\"value\"}'"


if [ -z "${RESOURCES+x}" ];then
  echo "ATLAS_ORG_ID must be set"
  exit 1
fi

if [ -z "${REGIONS+x}" ];then
  echo "REGIONS must be set"
  exit 1
fi

echo "resources: ${RESOURCES}"
echo "regions: ${REGIONS}"

#  loop for RESOURCES
IFS=','
read -ra ResourceNames <<< "$RESOURCES"

# iterate over the array of resources
for ResourceName in "${ResourceNames[@]}"; do
      echo "resource: $ResourceName"

      # Currently OTHER_PARAMS is only used for
      # 1. federated-settings-org-role-mapping
      # 2. trigger
      # 3. ldap-configuration
      # 4. ldap-verify
      # 5. thirdpartyintegrations

     # if condition for string comparison
     if [ -n "${OTHER_PARAMS}" ];then
           OtherParams=${OTHER_PARAMS}
     elif [[ "$ResourceName" == "trigger" ]]; then
          echo "setting up other params for trigger"
          echo "Required: projectId,cluster, app , function and service"

     elif [[ "$ResourceName" == "federated-settings-org-role-mapping" ]]; then
          echo "setting up other params for federated-settings-org-role-mapping"
          cat "$(dirname "$0")"/templates/ldap-configuration.json
          # fill the LDAP_BIND_PASSWORD in templates/ldap-configuration.json file using jq
          jq --arg ATLAS_FEDERATED_SETTINGS_ID "${ATLAS_FEDERATED_SETTINGS_ID}" \
             '.ATLAS_FEDERATED_SETTINGS_ID |= $ATLAS_FEDERATED_SETTINGS_ID' \
              "$(dirname "$0")/templates/federated-settings-org-role-mapping.json" >tmp.$$.json && mv tmp.$$.json "$(dirname "$0")/templates/federated-settings-org-role-mapping.json"
           OtherParams=$(cat "$(dirname "$0")"/templates/federated-settings-org-role-mapping-temp.json)

     elif [ "$ResourceName" == "ldap-verify" ] || [ "$ResourceName" == "ldap-configuration" ]; then
          echo "setting up other params for ldap"
          cat "$(dirname "$0")"/templates/ldap-configuration.json
          # fill the LDAP_BIND_PASSWORD in templates/ldap-configuration.json file using jq
          jq --arg LDAP_BIND_PASSWORD "${LDAP_BIND_PASSWORD}" \
             --arg LDAP_BIND_USER_NAME "${LDAP_BIND_USER_NAME}" \
             --arg LDAP_HOST_NAME "${LDAP_HOST_NAME}" \
             '.LDAP_BIND_PASSWORD |= $LDAP_BIND_PASSWORD | .LDAP_BIND_USER_NAME |= $LDAP_BIND_USER_NAME | .LDAP_HOST_NAME |= $LDAP_HOST_NAME' \
              "$(dirname "$0")/templates/ldap-configuration.json" >tmp.$$.json && mv tmp.$$.json "$(dirname "$0")/templates/ldap-configuration-temp.json"

          OtherParams=$(cat "$(dirname "$0")"/templates/ldap-configuration-temp.json)
          # convert json to string
          #OtherParamsString=$(echo "${OTHER_PARAMS}" | jq -r '.[0]')
    elif [[ "$ResourceName" == "third-party-integration" ]];  then
          echo "setting up other params for third-party-integration"
          # use jq to setup the parameters in third-party-integration.json file
          jq --arg webhook_create_url "$WEBHOOK_CREATE_URL" \
             --arg webhook_update_url "$WEBHOOK_UPDATE_URL" \
             --arg webhook_update_secret "$WEBHOOK_UPDATE_SECRET" \
             --arg prometheus_user_name "$PROMETHEUS_USER_NAME" \
             --arg prometheus_password_name "$PROMETHEUS_PASSWORD_NAME" \
             --arg pager_duty_create_service_key "$PAGER_DUTY_CREATE_SERVICE_KEY" \
             --arg pager_duty_update_service_key "$PAGER_DUTY_UPDATE_SERVICE_KEY" \
             --arg data_dog_create_api_key "$DATA_DOG_CREATE_API_KEY" \
             --arg data_dog_update_api_key "$DATA_DOG_UPDATE_API_KEY" \
             --arg ops_genie_api_key "$OPS_GENIE_API_KEY" \
             --arg microsoft_teams_webhook_create_url "$MICROSOFT_TEAMS_WEBHOOK_CREATE_URL" \
             --arg microsoft_teams_webhook_update_url "$MICROSOFT_TEAMS_WEBHOOK_UPDATE_URL" \
             '.WEBHOOK_CREATE_URL = $webhook_create_url |
              .WEBHOOK_UPDATE_URL = $webhook_update_url |
              .WEBHOOK_UPDATE_SECRET = $webhook_update_secret |
              .PROMETHEUS_USER_NAME = $prometheus_user_name |
              .PROMETHEUS_PASSWORD_NAME = $prometheus_password_name |
              .PAGER_DUTY_CREATE_SERVICE_KEY = $pager_duty_create_service_key |
              .PAGER_DUTY_UPDATE_SERVICE_KEY = $pager_duty_update_service_key |
              .DATA_DOG_CREATE_API_KEY = $data_dog_create_api_key |
              .DATA_DOG_UPDATE_API_KEY = $data_dog_update_api_key |
              .OPS_GENIE_API_KEY = $ops_genie_api_key |
              .MICROSOFT_TEAMS_WEBHOOK_CREATE_URL = $microsoft_teams_webhook_create_url |
              .MICROSOFT_TEAMS_WEBHOOK_UPDATE_URL = $microsoft_teams_webhook_update_url' \
              "$(dirname "$0")/templates/third-party-integration.json" >tmp.$$.json && mv tmp.$$.json "$(dirname "$0")/templates/third-party-integration-temp.json"

          OtherParams=$(cat "$(dirname "$0")"/templates/third-party-integration-temp.json)
    fi


    Path="cfn-resources/${ResourceName}/"
    CodeBuild_Project_Name="${ResourceName}-project-$((1 + RANDOM % 1000))"

    jq --arg ExecutionRoleName "${ExecutionRoleName}" \
        --arg TargetLocationsMaxConcurrency "${TargetLocationsMaxConcurrency}" \
        --arg AccountIds "${AccountIds}" \
        --arg Regions "${REGIONS}" \
        '.[0].ExecutionRoleName?|=$ExecutionRoleName |
        .[0].TargetLocationMaxConcurrency?|=$TargetLocationsMaxConcurrency |
        .[0].Accounts[0]?|=$AccountIds |
        .[0].Regions?|=($Regions | gsub(" "; "") | split(",")) ' \
        "$(dirname "$0")/templates/locations.json" >tmp.$$.json && mv tmp.$$.json "$(dirname "$0")/locations-temp.json"


    jq --arg Repository "${Repository}" \
       --arg ResourceName "${ResourceName}" \
       --arg OrgID "${ATLAS_ORG_ID}" \
       --arg PubKey "${ATLAS_PUBLIC_KEY}" \
       --arg PvtKey "${ATLAS_PRIVATE_KEY}" \
       --arg BranchName "${BranchName}" \
       --arg ProjectName "${CodeBuild_Project_Name}" \
       --arg OtherParams "${OtherParams}" \
       --arg Path "${Path}" \
       --arg BuilderRole "${BuilderRole}" \
       --arg AssumeRole "${AssumeRole}" \
       --arg LogDeliveryBucket "${LogDeliveryBucket}" \
      '.Repository[0]?|=$Repository |
      .ResourceName[0]?|=$ResourceName |
      .OrgID[0]?|=$OrgID |
      .PubKey[0]?|=$PubKey |
      .PvtKey[0]?|=$PvtKey |
      .ProjectName[0]?|=$ProjectName |
      .OtherParams[0]?|=$OtherParams |
      .BranchName[0]?|=$BranchName |
      .Path[0]?|=$Path |
      .BuilderRole[0]?|=$BuilderRole |
      .AssumeRole[0]?|=$AssumeRole |
      .LogDeliveryBucket[0]?|=$LogDeliveryBucket ' \
      "$(dirname "$0")/templates/params.json" >tmp.$$.json && mv tmp.$$.json "$(dirname "$0")/params-temp.json"


    params_json_content=$(cat "$(dirname "$0")"/params-temp.json)
    locations_json_content=$(cat "$(dirname "$0")"/locations-temp.json)

    # use the aws cli to start the automation execution
    aws ssm start-automation-execution \
        --document-name  ${AWS_SSM_Document_Name}\
        --document-version ${Document_Version} \
        --parameters "${params_json_content}" \
        --target-locations "${locations_json_content}" \
        --region "${Document_Region}"
done

