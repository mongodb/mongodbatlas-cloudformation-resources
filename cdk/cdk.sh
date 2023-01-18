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
            mv cdk-resources/${path}/src/*.ts cdk-resources/${path}/src/index.ts
            cd cdk-resources/${path}
            npx projen new awscdk-construct --author-name "MongoDB"  --docgen true --sample-code false --name path
            rm -rf .git
            cd ../..
          fi
      done
done



