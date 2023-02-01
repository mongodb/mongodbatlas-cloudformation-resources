#yarn add projen

# run publish.sh <resource-dir>
# example ./publish.sh private-endpoint
cd cdk-resources/$1
npx projen release:INTMDB-548
#export NPM_ACCESS_LEVEL=public
#export NPM_TOKEN=<>
yarn add publib
npx publib-npm