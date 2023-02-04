#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

temp=$(mktemp)
aws sts get-session-token > "${temp}"
echo "export AWS_SESSION_TOKEN=$(cat "${temp}" | jq -r '.Credentials.SessionToken')"
echo "export AWS_SECRET_ACCESS_KEY=$(cat "${temp}" | jq -r '.Credentials.SecretAccessKey')"
echo "export AWS_ACCESS_KEY_ID=$(cat "${temp}" | jq -r '.Credentials.AccessKeyId')"
rm "${temp}"
