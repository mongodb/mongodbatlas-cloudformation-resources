# Serverless Instance

## CFN resource type used
- MongoDB::Atlas::ServerlessInstance

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Serverless instance CFN resource](../../../../cfn-resources/serverless-instance/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. A serverless instance should show correctly configured for the database:

![image](https://user-images.githubusercontent.com/122359335/228200365-6e5950d8-1284-426c-97c8-57a6b24181d6.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Serverless-Instances)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/tutorial/create-serverless-instance/)