#!/usr/bin/env bash

# list all our types in a region - deal with paging api

page=$(aws cloudformation list-types --visibility PUBLIC)
token=$(echo ${page} | jq -r ".NextToken")
echo "${page}" | jq --arg filter MongoDB '.TypeSummaries[] | select(.TypeName? | match($filter))'
while [[ ! -z ${token} ]]; 
do
	page=$(aws cloudformation list-types --visibility PUBLIC --next-token ${token})
	token=$(echo ${page} | jq -r ".NextToken")
	echo "${page}" | jq --arg filter MongoDB '.TypeSummaries[] | select(.TypeName? | match($filter))'
done
