#!/bin/bash
# This shell script can be use to generate one or multiple L1 Constrcuts from CFN resources.
# How to execute:
#   1. For all resources without arguments :  ./cdk.sh
#   2. For specific resource with argument :  ./cdk/sh "<RESOURCE NAME>"
#                            For example   :  ./cdk/sh "auditing"  or  ./cdk/sh "cluster"
#
#   3. We can run it for multiple resources, let's say we are creating for "auditing", "cluster" and  "database-user"
#      then we can pass resource name with space separated.
#                                          : ./cdk/sh "<RESOURCE NAME> <RESOURCE NAME> <RESOURCE NAME>"
#                            For example   : ./cdk.sh "auditing cluster database-user"
#
#   In this way we can run it for one or multiple resources

resources=$1
if [[ -n "$resources" ]];
then
  source_dir=${resources};
else
  echo "Generating for all resources."
  source_dir='*';
  dir="../cfn-resources/$source_dir"
fi
for resource in ../cfn-resources/$source_dir;
do
  # echo $dir
  dir="../cfn-resources/$resource"
  for file in "$dir"/mongodb-atlas-*.json;
      do
          if [[ -f $file ]]; then
            #echo $file
            src=$( jq -r '.typeName' "$file" )
            echo "generating for $src"
            #mkdir -p cdk/${dir} && cd cdk/${dir}
            path=$(basename "$dir")
            echo "$path"
            index_exists=false
             if [ -f cdk-resources/"${path}"/src/index.ts ]; then
                rm -rf cdk-resources/"${path}"/src/*.ts
                index_exists=true
             fi

            cdk-import cfn -l typescript -s "$file" -o cdk-resources/"${path}"/src "$src"
            # need rename resource file to index.ts file
            mv cdk-resources/"${path}"/src/"mongodb-atlas-""${path//-}".ts cdk-resources/"${path}"/src/index.ts
            if [ "$index_exists" = true ] ; then
               continue
            fi
            cd cdk-resources/${path} || exit
            npx projen new awscdk-construct --npm-access "public" --author "MongoDBAtlas" --author-name "MongoDBAtlas" --docgen true --sample-code false --name '@mongodbatlas-awscdk/'"${path}" --author-address 'https://mongodb.com' --cdk-version '2.1.0' --default-release-branch 'INTMDB-548' --major-version 1 --release-to-npm true --repository-url 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git' --description 'Retrieves or creates '"${path}"' in any given Atlas organization' --keywords {'cdk','awscdk','aws-cdk','cloudformation','cfn','extensions','constructs','cfn-resources','cloudformation-registry','l1','mongodb','atlas',"$path"}
            rm -rf .git
            cd ../.. || exit
          fi
      done
done

