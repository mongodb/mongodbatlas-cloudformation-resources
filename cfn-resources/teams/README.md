# MongoDB::Atlas::Teams

## Description
Returns, adds, edits, or removes teams to Atlas Organization and Atlas Project (if specified).

## Attributes and Parameters

See the [resource docs](./docs/README.md).

## Installation

```
TAGS=logging make
cfn submit --verbose --set-default
```

## CloudFormation Examples

See the examples [CFN Template](/examples/teams/teams.json) for example resource.

## Integration Testing with AWS

The [/quickstart-mongodb-atlas/scripts/launch-quickstart.sh](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/quickstart-mongodb-atlas/scripts/launch-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other necessary parameters.

You can use the teams.sample-template.yaml to create a stack using the resource.
Similar to the local testing described above you can follow the logs for the deployed
lambda function which handles the request for the Resource Type.

In one shell session:
```
aws logs tail mongodb-atlas-teams-logs --follow
```

And then you can create the stack with a helper script it insert the apikeys for you:


```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/teams/test/teams.sample-template.yaml SampleAccessList1 ParameterKey=Name,ParameterValue=<YOUR_TEAM_NAME>  ParameterKey=Usernames,ParameterValue=<email id of one user assigned to your organization>
```

For more information see the MongoDB Atlas API documentation for [Teams](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Teams).
