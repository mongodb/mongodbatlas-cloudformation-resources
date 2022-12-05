# MongoDB::Atlas::DataLakes

## Description
Returns, adds, edits, and removes Federated Database Instances.

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
cd ${repo_root}/cfn-resources/datalakes
./test/datalakes.create-sample-cfn-request.sh YourProjectID ExternalId IamAssumedRoleARN IamUserARN RoleId TestS3BucketId > test.request.json 
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

Please see the [CFN Template](test/datalakes.sample-template.yaml) for example resource.

## Integration Testing w/ AWS

The [../../quickstart-mongodb-atlas/scripts/launch-quickstart.sh]( ../../quickstart-mongodb-atlas/scripts/launch-quickstart.sh)  script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other necessary parameters.

You can use the datalakes.sample-template.yaml to create a stack using the resource.
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
${repo_root}/quickstart-mongodb-atlas/scripts/launch-quickstart.sh ${repo_root}/cfn-resources/datalakes/test/datalakes.sample-template.yaml SampleAccessList1 ParameterKey=ProjectId,ParameterValue=<YOUR_PROJECT_ID> ParameterKey=TenantName,ParameterValue=<TenantName> ParameterKey=ExternalId,ParameterValue=<ExternalId> ParameterKey=IamAssumedRoleARN,ParameterValue=<IamAssumedRoleARN> ParameterKey=IamUserARN,ParameterValue=<IamUserARN> ParameterKey=RoleID,ParameterValue=<RoleId> ParameterKey=TestS3Bucket,ParameterValue=<TestS3Bucket> 

```

For more information see: MongoDB Atlas API [DataLakes Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Data-Federation) .
