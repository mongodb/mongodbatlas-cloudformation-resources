# LDAP Configuration

## CFN resource type used
- MongoDB::Atlas::LDAPConfiguration

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [LDAPConfiguration CFN resource](../../../../cfn-resources/ldap-configuration/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. LDAP Authentication (under Advanced section) should be correctly set up in your Atlas Project as per configuration specified in the inputs/example:
   ![image](https://user-images.githubusercontent.com/122359335/227264049-b1e44366-553c-417a-b541-15589a636037.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/LDAP-Configuration)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-ldaps/)