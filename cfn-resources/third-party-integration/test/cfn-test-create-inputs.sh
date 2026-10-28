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
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

profile="default"
if [ ${MONGODB_ATLAS_PROFILE+x} ]; then
	echo "profile set to ${MONGODB_ATLAS_PROFILE}"
	profile=${MONGODB_ATLAS_PROFILE}
fi

rm -rf inputs
mkdir inputs

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
	# Atlas project lists are eventually consistent: when this script follows a project
	# creation (contract testing creates the project first), the lookup above can miss
	# the new project and the create below fails with HTTP 409. Retry the lookup with
	# a bounded delay because the list can still be stale right after the conflict.
	projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id') || projectId=""
	if [ -z "$projectId" ]; then
		for _ in $(seq 1 12); do
			projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id') || projectId=""
			if [ -n "$projectId" ]; then break; fi
			sleep 5
		done
	fi
	if [ -z "$projectId" ]; then
		echo "Failed to create or find project ${projectName}" >&2
		exit 1
	fi
	echo -e "Using project \"${projectName}\" with id: ${projectId}\n"
else
	echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

#WebHook

webHookCreateUrl="$WEBHOOK_CREATE_URL"
webHookUpdateUrl="$WEBHOOK_UPDATE_URL"
webHookUpdateSecret="$WEBHOOK_UPDATE_SECRET"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg web_hook_create_url "$webHookCreateUrl" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .Url?|=$web_hook_create_url' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg web_hook_create_url "$webHookUpdateUrl" \
	--arg web_hook_create_secret "$webHookUpdateSecret" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .Url?|=$web_hook_create_url | .Secret?|=$web_hook_create_secret' \
	"$(dirname "$0")/inputs_1_update.template.json" >"inputs/inputs_1_update.json"

#PROMETHEUS
prometheusUsrName="$PROMETHEUS_USER_NAME"
prometheusPassword="$PROMETHEUS_PASSWORD_NAME"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg prometheus_usr_name "$prometheusUsrName" \
	--arg prometheus_password "$prometheusPassword" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .UserName?|=$prometheus_usr_name | .Password?|=$prometheus_password' \
	"$(dirname "$0")/inputs_2_create.template.json" >"inputs/inputs_2_create.json"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg prometheus_usr_name "$prometheusUsrName" \
	--arg prometheus_password "$prometheusPassword" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .UserName?|=$prometheus_usr_name | .Password?|=$prometheus_password' \
	"$(dirname "$0")/inputs_2_update.template.json" >"inputs/inputs_2_update.json"

#PAGER_DUTY

pagerDutyCreateServiceKey="$PAGER_DUTY_CREATE_SERVICE_KEY"
pagerDutyUpdateServiceKey="$PAGER_DUTY_UPDATE_SERVICE_KEY"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg service_key "$pagerDutyCreateServiceKey" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .ServiceKey?|=$service_key' \
	"$(dirname "$0")/inputs_3_create.template.json" >"inputs/inputs_3_create.json"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg service_key "$pagerDutyUpdateServiceKey" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .ServiceKey?|=$service_key' \
	"$(dirname "$0")/inputs_3_update.template.json" >"inputs/inputs_3_update.json"

#DATA_DOG

dataDogCreateApiKey="$DATA_DOG_CREATE_API_KEY"
dataDogUpdateApiKey="$DATA_DOG_UPDATE_API_KEY"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg data_dog_create_api_key "$dataDogCreateApiKey" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .ApiKey?|=$data_dog_create_api_key' \
	"$(dirname "$0")/inputs_4_create.template.json" >"inputs/inputs_4_create.json"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg data_dog_update_api_key "$dataDogUpdateApiKey" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .ApiKey?|=$data_dog_update_api_key' \
	"$(dirname "$0")/inputs_4_update.template.json" >"inputs/inputs_4_update.json"

#OPS_GENIE

opsGenieApiKey="$OPS_GENIE_API_KEY"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg ops_genie_api_key "$opsGenieApiKey" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .ApiKey?|=$ops_genie_api_key' \
	"$(dirname "$0")/inputs_5_create.template.json" >"inputs/inputs_5_create.json"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg ops_genie_api_key "$opsGenieApiKey" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .ApiKey?|=$ops_genie_api_key' \
	"$(dirname "$0")/inputs_5_update.template.json" >"inputs/inputs_5_update.json"

#MICROSOFT_TEAMS

microsoftTeamsCreateWebHook="$MICROSOFT_TEAMS_WEBHOOK_CREATE_URL"
microsoftTeamsUpdateWebHook="$MICROSOFT_TEAMS_WEBHOOK_UPDATE_URL"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg microsoft_teams_create_web_hook "$microsoftTeamsCreateWebHook" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .MicrosoftTeamsWebhookUrl?|=$microsoft_teams_create_web_hook' \
	"$(dirname "$0")/inputs_6_create.template.json" >"inputs/inputs_6_create.json"

jq --arg project_id "$projectId" \
	--arg profile "$profile" \
	--arg microsoft_teams_update_web_hook "$microsoftTeamsUpdateWebHook" \
	'.ProjectId?|=$project_id | .Profile?|=$profile | .MicrosoftTeamsWebhookUrl?|=$microsoft_teams_update_web_hook' \
	"$(dirname "$0")/inputs_6_update.template.json" >"inputs/inputs_6_update.json"

cd ..
