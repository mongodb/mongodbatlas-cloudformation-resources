# MongoDB::Atlas::ServerlessInstance

## Description
Returns, adds, edits, and removes serverless instances.

## Profile Setup
This resource requires to set up a profile with the SecretManager.
For an example, see [Secret Manager Profile setup](../../examples/profile-secret.yaml)

## Attributes and Parameters
See the [resource docs](./docs/README.md)

## Installation

```
TAGS=logging make
cfn submit --verbose --set-default
```

## Usage

The [launch-x-quickstart.sh](https://github.com/aws-quickstart/quickstart-mongodb-atlas/blob/master/scripts/launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other necessary parameters.

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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/project/test/project.sample-template.yaml SampleProject1 ParameterKey=OrgId,ParameterValue=${ATLAS_ORG_ID}
```

## For More Information
See the MongoDB Atlas API [Serverless Instance](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Serverless-Instances) documentation.
