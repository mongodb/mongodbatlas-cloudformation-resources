## MongoDB::Atlas::LDAPConfiguration

### Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- LDAP configuration L1 CDK constructor


### Resources (and parameters for local tests) needed to manually QA:
These LDAP resources must be manually created.
- LDAP Bind password (LDAP_BIND_PASSWORD)
- LDAP Bind user name (LDAP_BIND_USER_NAME)
- LDAP host name (LDAP_HOST_NAME)
- Port (defaults to 636)
- Atlas Project (created by cfn-test-create-inputs.sh)

## Manual QA:

### Prerequisite steps:
1. You would need AD servers that can be used to test this resource.
2. Export environment variables LDAP_BIND_PASSWORD,LDAP_BIND_USER_NAME, LDAP_HOST_NAME.

### Steps to test:
1. Follow general [prerequisites](../../../TESTING.md.md#prerequisites) for testing CFN resources.
2. In the Atlas Project you plan to use for testing, create an M10 Atlas Cluster or higher, if not already present.
3. Update LDAPConfiguration.json under cfn-resources/examples/ if required.
4. Follow [general steps](../../../TESTING.md.md#steps) to test CFN resources.
5. Once the template with required parameters is used to create, update and delete a stack successfully, validate that success criteria is met.

### Success criteria when testing the resource
1. LDAP Authentication (under Advanced section) should be correctly set up in your Atlas Project as per configuration specified in the inputs/example:
   ![image](https://user-images.githubusercontent.com/122359335/227264049-b1e44366-553c-417a-b541-15589a636037.png)

2. General [CFN resource success criteria](../../../TESTING.md.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/LDAP-Configuration)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-ldaps/)