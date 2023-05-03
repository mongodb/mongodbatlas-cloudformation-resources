# MongoDB::Atlas::EncryptionAtRest

## Description

Returns and edits the Encryption at Rest using Customer Key Management configuration.

## Attributes and Parameters

See the [Resource Docs](./docs/README.md).

## Installation

```
TAGS=logging make
cfn submit --verbose --set-default
```

## Cloudformation Examples

See the [CFN Template](./test/encryptionatrest.sample-template.yaml) for example resource.

## Integration Testing with AWS

The [`launch-quickstart.sh`](https://github.com/aws-quickstart/quickstart-mongodb-atlas/blob/8cd6e87ed36f467202432eed170a25ff10abf566/scripts/launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other necessary parameters.

You can use the `project.sample-template.yaml` to create a stack using the resource.
Similar to the local testing described above you can follow the logs for the deployed
lambda function which handles the request for the Resource Type.

In one shell session:

```
aws logs tail mongodb-atlas-project-logs --follow
```

And then you can create the stack with a helper script it insert the API Keys for you:


```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
${repo_root}/quickstart-mongodb-atlas/scripts/launch-quickstart.sh ${repo_root}/cfn-resources/encryption-at-rest/test/encryptionatrest.sample-template.yaml SampleAccessList1 ParameterKey=ProjectName,ParameterValue=<YOUR_PROJECT_ID> ParameterKey=CustomerMasterKeyID,ParameterValue=<CustomerMasterKeyID> ParameterKey=RoleID,ParameterValue=<RoleID> ParameterKey=Region,ParameterValue=<Region> ParameterKey=Enabled,ParameterValue=<true or false>
```

For more information see the MongoDB Atlas API documentation for ["Encryption At Rest"](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Encryption-at-Rest-using-Customer-Key-Management).
