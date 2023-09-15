## MongoDB::Atlas::LDAPVerify

### Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- LDAP verify L1 CDK constructor


### Resources (and parameters for local tests) needed to manually QA:
These LDAP resources must be manually created.
- LDAP Bind password (LDAP_BIND_PASSWORD)
- LDAP Bind user name (LDAP_BIND_USER_NAME)
- LDAP host name (LDAP_HOST_NAME)
- Atlas Project (created by cfn-test-create-inputs.sh)

## Manual QA:

### Prerequisite steps:
1. You would need AD servers that can be used to test this resource.
2. Export environment variables LDAP_BIND_PASSWORD, LDAP_BIND_USER_NAME, LDAP_HOST_NAME.

### Steps to test:
1. Follow general [prerequisites](../../../TESTING.md#prerequisites) for testing CFN resources. Note that running ./cfn-testing-helper creates an Atlas Project with an M10 cluster to create test parameters.
2. Update LDAPVerify.json under cfn-resources/examples/ if required.
3. Follow [general steps](../../../TESTING.md#steps) to test CFN resources.

### Success criteria when testing the resource
1. In the AWS CloudFormation stack, Output parameters should be correct as per LDAP configuration in your Atlas Project.
   ![image](https://user-images.githubusercontent.com/122359335/227264049-b1e44366-553c-417a-b541-15589a636037.png)

2. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api/ldaps-configuration-request-verification/)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-ldaps/)