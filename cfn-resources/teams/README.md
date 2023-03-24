# MongoDB::Atlas::Teams

## Description
Returns, adds, edits, or removes teams to Atlas Organization and Atlas Project (if specified).

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/teams
./test/teams.create-sample-cfn-request.sh {TEAM_NAME},{ ORGANIZATION_USER_EMAIL} > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.

## Installation
TAGS=logging make
cfn submit --verbose --set-default

## CloudFormation Examples

Please see the [CFN Template](/examples/teams/teams.json) for example resource

## Integration Testing w/ AWS

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

For more information see: MongoDB Atlas API [Teams Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Teams) Documentation.
