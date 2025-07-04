## MongoDB::Atlas::PrivateEndpointADL

### Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
-Private endpoint adl L1 CDK constructor



### Resources (and parameters for local tests) needed to manually QA:
The private endpoint id is to be manually provided.
- Atlas project (created by cfn-testing-helper.sh)
- Private endpoint id (AWS_VPC_ENDPOINT)

## Manual QA:

### Prerequisite steps:
1. You need a private VPC endpoint to test this resource:
   
   **a) [Recommended] Option1: Create a new test VPC with VPC endpoint in AWS:**
       
   i) Go to the AWS VPC console and click on Create VPC.

      ii) Refer to the [configuration here](https://user-images.githubusercontent.com/122359335/227306518-26eb8155-db09-4db1-8e7d-7a4a9eb1548d.png). This will have AWS quickly spin up a new VPC with private subnets and endpoints. Also select Enable DNS hostnames and Enable DNS resolution boxes.

      iii) Once VPC creation is done, navigate to the Endpoints section under the VPC console and search for the created endpoint with your VPC ID. Note the VPC endpoint ID with Service name as “com.amazonaws.vpce.us-east-1.*”

   **b) Option 2: Use an existing VPC with private subnet and create endpoint using AWS CLI:**

   i) In Atlas UI, navigate to your project -> Network Access -> Private Endpoint -> click on tab Federated Database Instance / Online Archive -> Create New Endpoint button.

   ii) Follow UI prompts to select AWS region and add your VPC and subnet IDs ([see screenshot](https://user-images.githubusercontent.com/122359335/227306584-3205de0c-a0a3-4d79-a20a-925630f10b85.png))

   iii) Copy and run the command mentioned in the UI prompt to create VPC Interface Endpoint.

   iv) Note your VPCEndpointId from the response.


2. Export AWS_VPC_ENDPOINT environment variable with value as the VPCEndpointId from #1.

### Steps to test:
1. Follow general [prerequisites](../../../TESTING.md#prerequisites) for testing CFN resources.
2. Follow [general steps](../../../TESTING.md#steps) to test CFN resources.

### Success criteria when testing the resource
1. Private Endpoint should be correctly set up in your Atlas Project as per configuration specified in the inputs/example:   

![image](https://user-images.githubusercontent.com/122359335/227305880-c6c70d20-7f38-4885-a3ed-1de7b4921aa3.png)

2. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-private-endpoint-services)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-cluster-private-endpoint/#set-up-a-private-endpoint-for-a-dedicated-cluster)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/private-endpoint-adl
./test/cluster.create-sample-cfn-request.sh YourProjectID YourClusterName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke resource CREATE test.request.json 
cfn invoke resource DELETE test.request.json 
cd -
```

Both CREATE & DELETE tests must pass.