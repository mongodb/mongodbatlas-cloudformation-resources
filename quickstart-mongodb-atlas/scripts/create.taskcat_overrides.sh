#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail
cat << EOF
PublicKey: "${ATLAS_PUBLIC_KEY}"
PrivateKey: "${ATLAS_PRIVATE_KEY}"
OrgId: "${ATLAS_ORG_ID}"
EOF
