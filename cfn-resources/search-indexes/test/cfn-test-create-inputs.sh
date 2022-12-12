#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#
set -o errexit
set -o nounset
set -o pipefail

function usage {
    echo "usage:$1 <projectID> $2<Cluster_Name> $3<DB_Name> $4<Coll_Name>"
    echo "Creates a new Search Index"
}

if [ "$#" -ne 1 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

projectId="${2:-625454459c4e6108393d650d}"
db_name="${4:-testdb}"
coll_name="${5:-testcoll}"
cluster_name="${3:-CFNTest-Cluster-001}"
index_name="search-$RANDOM"
u_index_name="${index_name}"
WORDTOREMOVE="template."
cd "$(dirname "$0")" || exit
for inputFile in inputs_*;
do
  outputFile=${inputFile//$WORDTOREMOVE/};
  index_name="${u_index_name}"
  if [[ ${inputFile} == *"invalid"* ]]; then
    index_name="invalid_name"
  fi
 jq --arg pubkey "$ATLAS_PUBLIC_KEY" \
    --arg pvtkey "$ATLAS_PRIVATE_KEY" \
    --arg org "$projectId" \
    --arg cluster  "$cluster_name" \
    --arg name  "$index_name"\
    --arg db "$db_name" \
    --arg coll "$coll_name" \
    '.CollectionName?|=$coll |.Database?|=$db |.GroupId?|=$org | .ApiKeys.PublicKey?|=$pubkey | .ApiKeys.PrivateKey?|=$pvtkey |.ClusterName?|=$cluster |.Name?|=$name' \
     "$inputFile" > "../inputs/$outputFile"
done
cd ..
ls -l inputs