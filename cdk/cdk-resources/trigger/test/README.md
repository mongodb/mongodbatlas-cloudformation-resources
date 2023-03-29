# Trigger

## CFN resource type used
- MongoDB::Atlas::Trigger

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Trigger CFN resource](../../../../cfn-resources/trigger/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Trigger should be set up in your Atlas account as per configuration specified in the inputs/example.

   ![image](https://user-images.githubusercontent.com/122359335/227495196-59063691-c475-449c-b6b1-f206f4404715.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/app-services/admin/api/v3/#tag/triggers)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/triggers/#service-functions-provide-server-side-logic)