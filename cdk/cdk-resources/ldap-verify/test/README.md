# LDAP Verify

## CFN resource type used
- MongoDB::Atlas::LDAPVerify

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [LDAPVerify CFN resource](../../../../cfn-resources/ldap-verify/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. In the AWS CloudFormation stack, Output parameters should be correct as per LDAP configuration in your Atlas Project.
   ![image](https://user-images.githubusercontent.com/122359335/227264049-b1e44366-553c-417a-b541-15589a636037.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api/ldaps-configuration-request-verification/)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-ldaps/)