#!/usr/bin/env bash
filter="${1}"
echo "filter: ${filter}"
PAGER="" aws s3api list-buckets \
 --query "Buckets[?starts_with(Name, \`$filter\`) == \`true\`].[Name]" \
 --output text



PAGER="" aws s3api list-buckets \
 --query "Buckets[?starts_with(Name, \`$filter\`) == \`true\`].[Name]" \
 --output text | xargs -I {} ./delete_all_object_versions.sh {}

#./delete_all_object_versions.sh {}

