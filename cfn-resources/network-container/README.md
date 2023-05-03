# MongoDB::Atlas::NetworkContainer

## Description
Returns, adds, edits, and removes network peering containers.

## Attributes and Parameters

See the [Resource Docs](./docs/README.md).

## Installation

Installation currently requires the follow steps to build and then submit/register the 
new MongoDB::Atlas::networkcontainer Resource Type into your AWS Region. Note, this command uses the
default AWS region.

```bash
TAGS=logging make
cfn submit --verbose --set-default
```

## Integration Testing with AWS

Once the resource is installed, you can do integrated testing from your shell to AWS.

Use the [launch-x-quickstart.sh](https://github.com/aws-quickstart/quickstart-mongodb-atlas/blob/8cd6e87ed36f467202432eed170a25ff10abf566/scripts/launch-x-quickstart.sh) script
to safely inject your MongoDB Cloud ApiKey environment variables into an example
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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/network-container/test/networkcontainer.sample-template.json SampleNetworkContainer-123 ParameterKey=ProjectId,ParameterValue=<YOUR_PROJECT_ID>
 
 
```

For more information see the Atlas API Documentation ["Network Peering Containers"](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Network-Peering-Containers).
