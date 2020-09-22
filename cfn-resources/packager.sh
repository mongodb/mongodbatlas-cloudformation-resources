#!/usr/bin/env bash

# packager.sh
#
# This script will create a zip file for each Atlas custom cfn resource
# and upload this to a target s3 bucket.
# This script does not register the custom resources into any AWS regions.
# See `registerer.sh` for that procedure.
#
# In general, users of the Atlas cfn resources should not need to run this script.
#
# MongoDB will run it and publish our pre-built assests to well known s3 buckets.
#
# This tool works by running `cfn submit` on each resource, but with the
# magic `--dry-run` flag set. This will cause the tooling to generate a local zipfile
# package for each resource which we then just `cp` up to the appropriate s3 bucket.
#

set -o errexit
set -o nounset
set -o pipefail
source ./log.sh

region="${AWS_REGION}"
version=$(git rev-parse --short HEAD)
# Default, find all the directory names with the json custom resource schema files.
resources="${1:-$(ls -F **/mongodb-atlas-*.json | cut -d/ -f1)}"
bucket="${2:-"s3://mongodb-cloudformation-resources-beta"}"
log_info "packager.sh hello"
log_info "region: ${region}"
log_info "version: ${version}"
log_info "bucket: ${bucket}"
log_info "resources: $(echo ${resources}| tr '\n' ' ')"
if [[ -z $(aws s3api head-bucket --bucket my-bucket) ]]; then
    log_info "mkbucket_response: bucket exists"
else
    mkbucket_resp=$(aws s3 mb "${bucket}")
    log_info "mkbucket_response: ${mkbucket_resp}"
fi

for resource in ${resources};
do
    cwd=$(pwd)
    cd "${resource}"
    log_info "resource: ${resource}"
    resp=$(make clean build)
    log_info "make_response: ${resp}"
    package=$(cfn submit --dry-run | cut -d":" -f2 | cut -d" " -f2)
    log_info "package: ${package}"
    package_dir="$(basename "${package}" | cut -d'.' -f1)"
    log_info "package_dir: ${package_dir}"
    versioned_pkg="${package_dir}-${version}.zip"
    log_info "versioned_pkg: ${versioned_pkg}"
    target="${bucket}/${package_dir}/${versioned_pkg}"
    log_info "target: ${target}"
    aws_s3_cp_response=$(aws s3 cp "${package}" "${target}" 2>&1)
    log_info "aws_s3_cp_response: ${aws_s3_cp_response}"
    cd "${cwd}"
done
aws_s3_ls_response=$(aws s3 ls --recursive "${bucket}" 2>&1)
log_info "aws_s3_ls_response: ${aws_s3_ls_response}"

