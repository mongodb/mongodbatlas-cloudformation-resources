#!/usr/bin/env bash

# registerer.sh
#
# This script will register the Atlas custom cfn resources.
# It should be run in each region you wish to use the Atlas resources.

set -o errexit
set -o nounset
set -o pipefail
source ./log.sh

region="${AWS_REGION}"
version=$(git rev-parse --short HEAD)
bucket="${2:-"s3://mongodb-cloudformation-resources-beta"}"
b_prefix="mongodb-atlas-"
log_info "registerer.sh hello"
log_info "region: ${region}"
log_info "bucket: ${bucket}"
log_info "b_prefix: ${b_prefix}"
bucket_ls=$(aws s3 ls --recursive "${bucket}")
log_info "bucket_ls: ${bucket_ls}"
TYPES="${1:-$(ls **/mongodb-atlas-*.json | xargs -I {} jq -r '.typeName' {})}"
while IFS= read -r type
do
    type_l=$(echo "${type}" | rev | cut -d':' -f1 | rev | tr '[:upper:]' '[:lower:]')
    schema_handler_package="${bucket}/${b_prefix}${type_l}/${b_prefix}${type_l}-${version}.zip"
    log_info "type: ${type}"
    log_info "type_l: ${type_l}"
    log_info "schema_handler_package: ${schema_handler_package}"
    reg_resp=$(aws cloudformation register-type \
        --region "${region}" \
        --type RESOURCE \
        --type-name "${type}" \
        --schema-handler-package "${schema_handler_package}")
    log_info "reg_resp: ${reg_resp}"
    list_type_versions=$(aws cloudformation list-type-versions \
        --region "${region}" \
        --type RESOURCE \
        --type-name "${type}")
    log_info "list_type_versions: ${list_type_versions}"
    version_id=$(echo ${list_type_versions} | jq -r '.TypeVersionSummaries[] | .VersionId' | tail -1)
    log_info "version_id: ${version_id}"
    def_version=$(aws cloudformation set-type-default-version \
    --region "${region}" \
    --type RESOURCE \
    --version-id "${version_id}" \
    --type-name "${type}")
    log_info "def_version: ${def_version}"

done < <(printf '%s\n' "${TYPES}")


