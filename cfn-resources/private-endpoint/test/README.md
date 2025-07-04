## MongoDB::Atlas::PrivateEndpoint

### Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- Private endpoint L1 CDK constructor
- Atlas Quickstart



### Resources (and parameters for local tests) needed to manually QA:
The VPD ID and subnet ID is to be manually provided.
- Atlas project (PROJECT_NAME(only required if you want to use an existing project))
- AWS VPC ID (AWS_VPC_ID)
- Subnet ID (AWS_SUBNET_ID)
- AWS Region (AWS_DEFAULT_REGION, will be read from AWS config if not provided)

## Manual QA:

### Prerequisite steps:
1. In the AWS VPC console, create a VPC with a Subnet. Note the VPC ID and the Subnet ID.
2. Export VPC ID, Subnet ID and AWS Region with environment variables AWS_VPC_ID, AWS_SUBNET_ID, AWS_DEFAULT_REGION respectively.

### Steps to test:
1. Follow general [prerequisites](../../../TESTING.md#prerequisites) for testing CFN resources.
2. Follow [general steps](../../../TESTING.md#steps) to test CFN resources.

### Success criteria when testing the resource
1. Private Endpoint should be correctly set up in your Atlas Project as per configuration specified in the inputs/example:   

![image](https://user-images.githubusercontent.com/122359335/227300711-ca08e118-8718-4285-a975-8ec4e01899f9.png)

2. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-private-endpoint-services)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-cluster-private-endpoint/#set-up-a-private-endpoint-for-a-dedicated-cluster)

## Local Testing

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/project
./test/project.create-sample-cfn-request.sh YourProjectName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.