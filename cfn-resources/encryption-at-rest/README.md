# MongoDB::Atlas::EncryptionAtRest

## Description
Resource for managing [Encryption at Rest](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-encryption-at-rest-using-customer-key-management)
using Customer Key Management configuration.

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](./docs/README.md).

## CloudFormation Examples

See the examples [CFN Template](/examples/encryption-at-rest/encryption-at-rest.json) for example resource.

<!-- 
--------------

# MongoDB::Atlas::EncryptionAtRest

## Description

Returns and edits the Encryption at Rest using Customer Key Management configuration.
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
set or export below environment variables 
export KMS_KEY=<<CustomerMasterKeyID>>
export KMS_ROLE=<<RoleID>>
export KMS_REGION=<<key region>>
cd ${repo_root}/cfn-resources/encryption-at-rest
./test/encryptionatrest.create-sample-cfn-request.sh YourProjectID > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.

## Installation
TAGS=logging make
cfn submit --verbose --set-default

## Cloudformation Examples

Please see the [CFN Template](test/encryptionatrest.sample-template.yaml) for example resource.

## Integration Testing w/ AWS

The [../../quickstart-mongodb-atlas/scripts/launch-quickstart.sh]( ../../quickstart-mongodb-atlas/scripts/launch-quickstart.sh)  script
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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-quickstart.sh ${repo_root}/cfn-resources/encryption-at-rest/test/encryptionatrest.sample-template.yaml SampleAccessList1 ParameterKey=ProjectName,ParameterValue=<YOUR_PROJECT_ID> ParameterKey=CustomerMasterKeyID,ParameterValue=<CustomerMasterKeyID> ParameterKey=RoleID,ParameterValue=<RoleID> ParameterKey=Region,ParameterValue=<Region> ParameterKey=Enabled,ParameterValue=<true or false>
```

For more information see: MongoDB Atlas API [Encryption At Rest Endpoint](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-encryption-at-rest-using-customer-key-management) Documentation.
 -->
