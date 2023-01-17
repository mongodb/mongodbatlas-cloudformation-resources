for dir in ../cfn-resources/*;
do
  # echo $dir
  for file in $dir/mongodb-atlas-*.json;
      do
          if [[ -f $file ]]; then
            path=$(basename $dir)
            echo $path
            cd cdk-resources/${path}
            rm -rf .git
            cd ../..
          fi
      done
done



