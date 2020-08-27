#!/usr/bin/env bash

echo "cleanup-runs-hub.sh"
echo "About: Deletes all the runs from a Github action."

REPO=$1
hub api /repos/$REPO/actions/runs | \
jq '.workflow_runs[] | .id' | \
xargs -I {} hub api -X DELETE /repos/jasonmimick/quickstart-mongodbatlas/actions/runs/{}

