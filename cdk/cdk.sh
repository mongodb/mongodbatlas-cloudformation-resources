#!/bin/bash

# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#!/bin/bash
# This shell script can be use to generate one L1 | L2 | L3 CDK Construct.
# In case of L1 constructor, the script uses the CFN template to generate the CDK resource.
# How to execute:
#   1. CDK L1:  ./cdk.sh "<RESOURCE NAME>" L1
#   1. CDK L1:  ./cdk.sh "<RESOURCE NAME>" L1
#   2. CDK L2:  ./cdk/sh "<RESOURCE NAME>" L2
#   4. CDK L3:  ./cdk/sh "<RESOURCE NAME>" L3

set -euo pipefail


_print_usage() {
	echo
  echo 'Usage:'
	echo './cdk.sh   "<RESOURCE NAME>" "<L1|L2|L3>"'
  echo
  echo 'Example:'
  echo './cdk.sh cluster L1'
  echo
}


if [ "$#" -ne 2 ]; then
    echo "Error: please provide the resource name and the type of CDK constructor"
    _print_usage
    exit 1
fi

resource=$1
cdk_type=$(echo "$2" | tr '[:lower:]' '[:upper:]')


if [[ -z "${cdk_type}" || "${cdk_type}" = "L1" ]]; then
  echo "Generating L1 CDK resource"
  dir="../cfn-resources/$resource"
  for file in "$dir"/mongodb-atlas-*.json;
      do
          if [[ -f $file ]]; then
            src=$( jq -r '.typeName' "$file" )
            echo "generating for $src"
            path=$(basename "$dir")
            index_exists=false
            if [ -f cdk-resources/"${path}"/src/index.ts ]; then
                rm -rf cdk-resources/"${path}"/src/*.ts
                index_exists=true
            fi

            cdk-import cfn -l typescript -s "$file" -o cdk-resources/"${path}"/src "$src"
            # need rename resource file to index.ts file
            mv "cdk-resources"/"${path}"/"src"/"mongodb-atlas-""${path//-}".ts cdk-resources/"${path}"/src/index.ts
            if [ "$index_exists" = true ] ; then
              continue
            fi
            
            pushd "cdk-resources/${path}" || exit
            npx projen new awscdk-construct --npm-access "public" --author "MongoDBAtlas" --author-name "MongoDBAtlas" --docgen true --sample-code false --name '@mongodbatlas-awscdk/'"${path}" --author-address 'https://mongodb.com' --cdk-version '2.1.0' --default-release-branch 'master' --major-version 1 --release-to-npm true --repository-url 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git' --description 'Retrieves or creates '"${path}"' in any given Atlas organization' --keywords {'cdk','awscdk','aws-cdk','cloudformation','cfn','extensions','constructs','cfn-resources','cloudformation-registry','l1','mongodb','atlas',"$path"}
            rm -rf .git
            popd
          fi
      done
fi

if [ "$cdk_type" = "L2" ]; then
  echo "Generating L2 CDK resource"
  mkdir "l2-cdk-resources/${resource}"
  pushd "l2-cdk-resources/${resource}"

  npx projen new awscdk-construct --npm-access "public" --author "MongoDBAtlas" --author-name "MongoDBAtlas" --docgen true --sample-code false --name "@mongodbatlas-awscdk-l2/${resource}" --author-address 'https://mongodb.com' --cdk-version '2.1.0' --default-release-branch 'master' --major-version 1 --release-to-npm true --repository-url 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git' --description "Retrieves or creates ${resource} in any given Atlas organization" --keywords {'cdk','awscdk','aws-cdk','cloudformation','cfn','extensions','constructs','cfn-resources','cloudformation-registry','l2','mongodb','atlas',"$resource"}
  rm -rf .git

  popd
fi

if [ "$cdk_type" = "L3" ]; then
  echo "Generating L3 CDK resource"
  mkdir "l3-cdk-resources/${resource}"
  pushd "l3-cdk-resources/${resource}"

  npx projen new awscdk-construct --npm-access "public" --author "MongoDBAtlas" --author-name "MongoDBAtlas" --docgen true --sample-code false --name "@mongodbatlas-awscdk-l3/${resource}" --author-address 'https://mongodb.com' --cdk-version '2.1.0' --default-release-branch 'master' --major-version 1 --release-to-npm true --repository-url 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git' --description "Retrieves or creates ${resource} in any given Atlas organization" --keywords {'cdk','awscdk','aws-cdk','cloudformation','cfn','extensions','constructs','cfn-resources','cloudformation-registry','l3','mongodb','atlas',"$resource"}
  rm -rf .git

  popd
fi


