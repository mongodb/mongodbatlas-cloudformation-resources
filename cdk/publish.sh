#yarn add projen

# run publish.sh <resource-dir>
# example ./publish.sh cdk-resources/private-endpoint
# ./publish.sh  l3-cdk-resources/atlas-basic
# ./publish.sh  l2-cdk-resources/
cd $1
npx projen release:INTMDB-548
# set these in env variable
#export NPM_ACCESS_LEVEL=public
#export NPM_TOKEN=<>
yarn add publib
npx publib-npm