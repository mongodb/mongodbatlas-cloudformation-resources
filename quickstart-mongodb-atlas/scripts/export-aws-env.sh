#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

temp=$(mktemp)
aws sts get-session-token >"${temp}"

AWS_SESSION_TOKEN=$(jq -r '.Credentials.SessionToken' "$temp")
export AWS_SESSION_TOKEN

AWS_SECRET_ACCESS_KEY=$(jq -r '.Credentials.SecretAccessKey' "$temp")
export AWS_SECRET_ACCESS_KEY

AWS_ACCESS_KEY_ID=$(jq -r '.Credentials.AccessKeyId' "$temp")
export AWS_ACCESS_KEY_ID

rm "${temp}"
