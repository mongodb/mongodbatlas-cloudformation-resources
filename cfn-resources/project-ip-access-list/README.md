# MongoDB::Atlas::ProjectIpAccessList


## Description
The Access List endpoint manages a Atlas project's IP Access List.

The Access List endpoint supports creating temporary Access List entries that automatically expire within a user-configurable 7-day period.


## Attributes

## Parameters
`ProjectId` *(required)* : Unique identifier of the project within which to create the project ip access list.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>


## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/project-ip-access-list
./test/projectipaccesslist.create-sample-cfn-request.sh YourProjectID > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.

## Installation
TAGS=logging make
cfn submit --verbose --set-default

## Integration Testing w/ AWS

The [/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh](launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other neccessary parameters.

You can use the project.sample-template.yaml to create a stack using the resource.
Similar to the local testing described above you can follow the logs for the deployed
lambda function which handles the request for the Resource Type.

In one shell session:
```
aws logs tail mongodb-atlas-project-logs --follow
```

And then you can create the stack with a helper script it insert the apikeys for you:


```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/projectipaccesslist/test/projectipaccesslist.sample-template.yaml SampleAccessList1 ParameterKey=ProjectId,ParameterValue=<YOUR_PROJECT_ID>
```

