#!/usr/bin/env bash
set -x
function usage {
	echo "usage:$0 <project_name>"
	echo "Creates a new encryption key for the the project "
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi
rm -rf inputs
mkdir inputs

jq --arg FederationSettingsId "$ATLAS_FEDERATED_SETTINGS_ID" \
	'.FederationSettingsId?|=$FederationSettingsId' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_create.json"

jq --arg FederationSettingsId "$ATLAS_FEDERATED_SETTINGS_ID" \
	'.FederationSettingsId?|=$FederationSettingsId ' \
	"$(dirname "$0")/inputs_1_create.template.json" >"inputs/inputs_1_update.json"

ls -l inputs
