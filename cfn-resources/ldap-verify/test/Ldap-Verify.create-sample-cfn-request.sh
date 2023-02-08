#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
#!/usr/bin/env bash
# project.create-sample-cfn-request.sh
#
# This tool generates text for a `cfn invoke` request json message.
#

set -o errexit
set -o nounset
set -o pipefail

projectId="${1}"
bindPassword="$LDAP_BIND_PASSWORD"
bindUsername="$LDAP_BIND_USER_NAME"
hostname="$LDAP_HOST_NAME"
jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
	--arg pvtkey "$ATLAS_PRIVATE_KEY" \
	--arg groupID "$projectId" \
	--arg bindPassword "$bindPassword" \
	--arg bindUsername "$bindUsername" \
	--arg hostname "$hostname" \
	'.desiredResourceState.GroupId?|=$groupID | .desiredResourceState.ApiKeys.PublicKey?|=$pubkey | .desiredResourceState.ApiKeys.PrivateKey?|=$pvtkey | .desiredResourceState.BindPassword?|=$bindPassword | .desiredResourceState.BindUsername?|=$bindUsername | .desiredResourceState.Hostname?|=$hostname' \
	"$(dirname "$0")/Ldap-Verify.sample-cfn-request.json"
