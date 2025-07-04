## MongoDB::Atlas::ServerlessPrivateEndpoint

### Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- Serverless PrivateEndpoint L1 CDK constructor
- Atlas Basic Serverless PrivateEndpoint L3 CDK constructor

### Resources (and parameters for local tests) needed to manually QA:
These LDAP resources must be manually created.
- AWS VPC and Subnet

## Manual QA:

### Prerequisite steps:
1. Obtain the VPC id and subnet to create the private endpoint:
   example:
```bash
# Use AWS CLI to list VPCs and get the first VPC ID
vpc_id=$(aws ec2 describe-vpcs --query 'Vpcs[0].VpcId' --output text)

# Use AWS CLI to list subnets in the selected VPC and get the first Subnet ID
subnet_id=$(aws ec2 describe-subnets --filters "Name=vpc-id,Values=$vpc_id" --query 'Subnets[0].SubnetId' --output text)
```

### Steps to test:
1. Follow general [prerequisites](../../../TESTING.md#prerequisites) for testing CFN resources.
2. In the Atlas Project you plan to use for testing, create a Serverless instance, if not already present.
3. Update serverless-private-endpoint-with-aws-private-endpoint.json under cfn-resources/examples/serverless-private-endpoint if required.
4. Follow [general steps](../../../TESTING.md#steps) to test CFN resources.
5. Once the template with required parameters is used to create, update and delete a stack successfully, validate that success criteria is met.

### Success criteria when testing the resource
1. the serverless private endpoint must be created
   ![img.png](img.png)
2. the aws private endpoint must be created
![img_1.png](img_1.png)
3. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-serverless-private-endpoints)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-serverless-private-endpoint/)

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
./test/serverless-private-endpoint.create-sample-cfn-request.sh YourProjectName YourInstanceName VpcId SubnetId > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE and DELETE tests must pass.
