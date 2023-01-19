resources="${1:-project}"
for resource in ${resources};
do
  # echo $dir
  dir="../cfn-resources/$resource"
#  cd $dir || exit 1
  for file in ./$dir/mongodb-atlas-*.json;
      do
          if [[ -f $file ]]; then
            src=$( jq -r '.typeName' $file )
            echo $file
            echo $src
            if [ -n "$src" ]; then
              cd "./cdk-resources/$resource"
              pwd
              cp ../../.projenrc.js .
              IFS='::'
              #Read the split words into an array based on comma delimiter
              read -a strarr <<< "$src"
              echo ${strarr[4]}
              sed -i -e "s/Project/${strarr[4]}/g" ".projenrc.js"
              sed -i -e "s/project/$resource/g" ".projenrc.js"
              cd -
            fi
          fi
      done
done