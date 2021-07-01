#!/usr/bin/env bash
p=$(pwd)
cd $(dirname "$0")
cd ../..
bash <(cat cfn-resources/aws-regions | cut -d' ' -f1 | xargs -I {} echo "AWS_DEFAULT_REGION=\"{}\" ./cfn-resources/scripts/list-mongodb-types.sh")
cd "${p}"
