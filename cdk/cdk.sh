for dir in ../cfn-resources/*;
do
  # echo $dir
  for file in $dir/mongodb-atlas-*.json;
      do
          if [[ -f $file ]]; then
            #echo $file
            src=$( jq -r '.typeName' $file )
            echo "generating for $src"
            #mkdir -p cdk/${dir} && cd cdk/${dir}
            path=$(basename $dir)
            echo $path
            rm -rf cdk-resources/${path}/src/*.ts
            cdk-import cfn -l typescript -s $file -o cdk-resources/${path}/src $src
            cd cdk-resources/${path}
            npx projen new awscdk-construct --author "MongoDBAtlas" --author-name "MongoDBAtlas" --docgen true --sample-code false --name '@mongodbatlas-awscdk/'${path} --author-address 'https://mongodb.com' --cdk-version '2.1.0' --default-release-branch 'INTMDB-548' --major-version 1 --release-to-npm true --repository-url 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git' --description 'Retrieves or creates '${path}' in any given Atlas organization' --keywords {'cdk','awscdk','aws-cdk','cloudformation','cfn','extensions','constructs','cfn-resources','cloudformation-registry','l1','mongodb','atlas',$path}
            rm -rf .git
            cd ../..
          fi
      done
done


