#!/usr/bin/env bash

# The Compleat Published
# One master script to generate a publish package

trap "exit" INT TERM ERR
trap "kill 0" EXIT
#set -x
set -o errexit
set -o nounset
set -o pipefail

publish_bucket="../publishing"

tag=$(git rev-parse --short HEAD)

for region in $(cat regions); do

	bucket="${publish_bucket}/${region}-${tag}"
	rt_cleaner=$(AWS_DEFAULT_REGION=us-west-2 ./scripts/scrub-all-resource-types.sh)  
 
	cat <<- HEADER >> ${region}-publish-${tag}.sh
		#!/usr/bin/env bash
		echo "+++++++ COMPLEAT PUBLISH START +++++++"
		export ATLAS_PUBLIC_KEY=$(cat ~/.config/mongocli.toml | grep 'public_api_key' | cut -d\" -f2)
		export ATLAS_PRIVATE_KEY=$(cat ~/.config/mongocli.toml | grep 'private_api_key' | cut -d\" -f2)
		export ATLAS_ORG_ID=$(cat ~/.config/mongocli.toml | grep 'org_id' | cut -d\" -f2)
		AWS_DEFAULT_REGION=$region 
		env | grep ATLAS_ 
		env | grep AWS_DEFAULT
		echo "+++++++ COMPLEAT PUBLISH RUNTIME FOLDER +++++++"
		rmdir --ignore-fail-on-non-empty "${bucket}"
		mkdir -p "${bucket}"
		cp -r . "${bucket}"
		cd  "${bucket}"
		rm **/*rpdk.log
		echo "+++++++ COMPLEAT PUBLISH LIST-TYPES-1 +++++++"
		AWS_DEFAULT_REGION=${region} ./list-mongodb-types.sh
		${rt_cleaner}
		(. ./scripts/aws-cfn-stack-cleaner.sh killall $region)  
	HEADER

	for resource in project cluster database-user; do
		cat <<- EOF2 >> "${region}-publish-${tag}.sh"
			(. AWS_DEFAULT_REGION=${region} ./cfn-testing-helper.sh $resource)
			(. SUBMIT_ONLY=true AWS_DEFAULT_REGION=${region} ./cfn-submit-helper.sh $resource)
			(. AWS_DEFAULT_REGION=${region} ./cfn-publishing-helper.sh $resource)

		EOF2
    
		done
  
	cat <<- EOF3 >> "${region}-publish-${tag}.sh"
		echo "+++++++ COMPLEAT PUBLISH LIST-TYPES-2 +++++++"
		(. AWS_DEFAULT_REGION=${region} ./list-mongodb-types.sh)
		cd -
		echo "+++++++ COMPLEAT PUBLISH COMPLETE +++++++"
	EOF3

	chmod +x "${region}-publish-${tag}.sh"
	ls -l ${region}-publish-${tag}.sh
done


