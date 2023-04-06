#!/usr/bin/env bash
# federated-setting-org-configs-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

function usage {
	echo "usage:$0 <FederationSettingsId> <ConnectedOrganizationId>"
	exit 1
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

federationSettingsId="${1}"
connectedOrganizationId="${2}"

jq --arg org "$connectedOrganizationId" \
	--arg FederationSettingsId "$federationSettingsId" \
	'.desiredResourceState.FederationSettingsId?|=$FederationSettingsId | .desiredResourceState.OrgId?|=$org ' \
	"$(dirname "$0")/federated-setting-org-configs.sample-cfn-request.json"
