# Private Endpoint

## CFN resource type used
- MongoDB::Atlas::PrivateEndpoint

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [PrivateEndpoint CFN resource](../../../../cfn-resources/private-endpoint/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Private Endpoint should be correctly set up in your Atlas Project as per configuration specified in the inputs/example:

![image](https://user-images.githubusercontent.com/122359335/227300711-ca08e118-8718-4285-a975-8ec4e01899f9.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Private-Endpoint-Services)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-cluster-private-endpoint/#set-up-a-private-endpoint-for-a-dedicated-cluster)