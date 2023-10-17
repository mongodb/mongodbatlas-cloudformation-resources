#!/usr/bin/env bash

set -e

cd schema-gen
go build

aws s3 cp s3://mdb-cfn-publish/swagger.json swagger.json
cd -

#generate the schema files
./schema-gen/schema-gen

latestFile="schema-gen/swagger.latest.json"

#download latest openapi spec
curl https://mongodb-mms-prod-build-server.s3.amazonaws.com/openapi/bfb9c50bdcffe3100e37c3f7b73b52915701c98b.json >"${latestFile}"
#find the diff in schema files with latest open api spec
./schema-gen/schema-gen compare

aws s3 cp "${latestFile}" s3://mdb-cfn-publish/swagger.json
