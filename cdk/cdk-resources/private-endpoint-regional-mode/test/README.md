# Private Endpoint Regional Mode

## CFN resource type used
- MongoDB::Atlas::PrivateEndPointRegionalMode

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [PrivateEndPointRegionalMode CFN resource](../../../../cfn-resources/private-endpoint-regional-mode/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Regionalized private endpoints setting should be enabled under Project Settings:

   ![image](https://user-images.githubusercontent.com/122359335/227656275-fd32b882-8b7d-4427-af6c-c4dc68fefd76.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Private-Endpoint-Services/operation/returnRegionalizedPrivateEndpointStatus)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-private-endpoint/#enable-regionalized-private-endpoints-1)
