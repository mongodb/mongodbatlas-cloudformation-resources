# MongoDB::Atlas::ProjectIpAccessList


## Description
Returns, adds, edits, and removes network access limits to database deployments in MongoDB Cloud.

## Attributes and Parameters

See the [Resource Docs](./docs/README.md).

## Installation

```
TAGS=logging make
cfn submit --verbose --set-default
```

## Integration Testing with AWS

The [launch-x-quickstart.sh](https://github.com/aws-quickstart/quickstart-mongodb-atlas/blob/master/scripts/launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other neccessary parameters.

You can use the [project.sample-template.yaml](test/projectipaccesslist.sample-template.yaml) to create a stack using the resource.
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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/projectipaccesslist/test/projectipaccesslist.sample-template.yaml ParameterKey=ProjectId,ParameterValue=<YOUR_PROJECT_ID>
```

For more information see the MongoDB Atlas API ["Project"](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Project-IP-Access-List) documentation.
