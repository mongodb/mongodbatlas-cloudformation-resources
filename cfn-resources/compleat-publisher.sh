#!/usr/bin/env bash

# The Compleat Published
# One master script to generate a publish package

#trap "exit" INT TERM ERR
#trap "kill 0" EXIT
#set -x
#set -o errexit
#set -o nounset
#set -o pipefail

publish_bucket="/home/ubuntu/publishing"
tag=$(git rev-parse --short HEAD)
publish_pkg="${publish_bucket}/mongodbatlas-cloudformation-resources.${tag}.tar.gz"
rm --force "${publish_pkg}"
cd ..
tar czvf "${publish_pkg}" cfn-resources
cd -
ls -l "${publish_bucket}"


for region in $(cat regions); do

	bucket="${publish_bucket}/${region}-${tag}"
	script="${publish_bucket}/${region}-publish-${tag}.sh" 
	cat <<- HEADER >> "${script}"
		#!/usr/bin/env bash
		echo "+++++++ COMPLEAT PUBLISH START +++++++"
		export ATLAS_PUBLIC_KEY=$(cat ~/.config/mongocli.toml | grep 'public_api_key' | cut -d\" -f2)
		export ATLAS_PRIVATE_KEY=$(cat ~/.config/mongocli.toml | grep 'private_api_key' | cut -d\" -f2)
		export ATLAS_ORG_ID=$(cat ~/.config/mongocli.toml | grep 'org_id' | cut -d\" -f2)
		export AWS_DEFAULT_REGION=$region 
		env | grep ATLAS_ 
		env | grep AWS_DEFAULT
		echo "+++++++ COMPLEAT PUBLISH RUNTIME FOLDER +++++++"
		rmdir --ignore-fail-on-non-empty "${bucket}"
		mkdir -p "${bucket}"
		cp "${publish_pkg}" "${bucket}"
		cd  "${bucket}"
		tar xzvf "mongodbatlas-cloudformation-resources.${tag}.tar.gz"
		cd cfn-resources
		rm **/*rpdk.log
		echo "+++++++ COMPLEAT PUBLISH LIST-TYPES-1 +++++++"
		. ./scripts/list-mongodb-types.sh
	HEADER
	for resource in $(aws cloudformation list-types | jq -r '.TypeSummaries[] | .TypeName'); do
		cat <<- HEADER2 >> "${script}"
			. ./scripts/aws-cfn-resource-type-cleaner.sh "${resource}"
		HEADER2
	done
	cat <<- HEADER3 >> "${script}"
		./scripts/aws-cfn-stack-cleaner.sh killall $region
	HEADER3
	cfnmui=$(aws cloudformation describe-stacks --region $region --stack-name "CloudFormationManagedUploadInfrastructure" --query "Stacks[0].Outputs[1].OutputValue")
	if [ -z "${cfnmui}" ]; do
		cat <<- HEADER4 >> "${script}"
			aws cloudformation update-termination-protection --no-enable-termination-protection--stack-name "CloudFormationManagedUploadInfrastructure"
			. ./scripts/delete-s3-bucket.py ${cfnmui}
			aws cloudformation delete-stack--stack-name "CloudFormationManagedUploadInfrastructure"
		HEADER4
	done
			
	for resource in project cluster database-user; do
		cat <<- EOF2 >> "${script}"
			export AWS_DEFAULT_REGION=$region 
			export SUBMIT_ONLY=true  
			. ./cfn-testing-helper.sh $resource
			. ./cfn-submit-helper.sh $resource
			. ./cfn-publishing-helper.sh $resource
		EOF2
	done
  
	cat <<- EOF3 >> "${script}"
		echo "+++++++ COMPLEAT PUBLISH LIST-TYPES-2 +++++++"
		export AWS_DEFAULT_REGION=$region 
		. ./scripts/list-mongodb-types.sh
		cd "${publish_bucket}"
		echo "+++++++ COMPLEAT PUBLISH COMPLETE +++++++"
	EOF3

  chmod +x "${script}"
  ls -l "${script}"
done


