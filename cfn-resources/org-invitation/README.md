# Mongodb::Atlas::OrgInvitation

Add, delete and update invitation for MongoDB organizations.

## Attributes and Parameters

See the [Resource Docs](./docs/README.md).

## Installation

```
TAGS=logging make
cfn submit --verbose --set-default
```

## CloudFormation Examples

Please see the [CFN Template](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/org-invitation/org-invitation-sample.json) for example resource.

## Usage

The [launch-x-quickstart.sh](https://github.com/aws-quickstart/quickstart-mongodb-atlas/blob/8cd6e87ed36f467202432eed170a25ff10abf566/scripts/launch-x-quickstart.sh) script
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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/project/test/project.sample-template.yaml SampleProject1 ParameterKey=OrgId,ParameterValue=${ATLAS_ORG_ID}
```

## For More Information

See the MongoDB Atlas API [Organization](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Organizations) documentation.
