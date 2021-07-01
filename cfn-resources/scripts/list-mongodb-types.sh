#!/usr/bin/env bash

# list all our types in a region - deal with paging api

page=$(aws cloudformation list-types --visibility PUBLIC 2> /dev/null)
if [[ $? -ne 0 ]]; then
	echo "{  \"Region\": \"${AWS_DEFAULT_REGION}\", \"Error\": \"REGION ACCOUNT NOT VALID\"}" | jq '.'
	exit 0
else
	echo "{  \"Region\": \"${AWS_DEFAULT_REGION}\"}" | jq '.'
fi
echo "${page}" | jq --arg filter MongoDB '.TypeSummaries[] | select(.TypeName? | match($filter))'
token=$(echo ${page} | jq -r ".NextToken") 
while [[ ! -z ${token} ]]; 
do
  if [[ "${token}" == "null" ]]; then
		break
	fi
	page=$(aws cloudformation list-types --visibility PUBLIC --next-token ${token})
	token=$(echo ${page} | jq -r ".NextToken")
	echo "${page}" | jq --arg filter MongoDB '.TypeSummaries[] | select(.TypeName? | match($filter))'
  #echo "~~~~~~~~~~~~token=${token}"
  #echo "-------"
done
