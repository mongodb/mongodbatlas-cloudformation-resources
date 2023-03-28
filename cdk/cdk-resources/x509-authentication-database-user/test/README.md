# X509 Authentication Database User

## CFN resource type used
- MongoDB::Atlas::X509AuthenticationDatabaseUser

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [X509AuthenticationDatabaseUser CFN resource](../../../../cfn-resources/x509-authentication-database-user/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. X.509 Authentication should be enabled correctly set up in your Atlas Project as per configuration specified in the inputs/example. This can be found under Security section -> Advanced:

   ![image](https://user-images.githubusercontent.com/122359335/227374480-1afa48a4-5265-4a2a-ad92-067f5015eeca.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/X.509-Authentication/operation/createDatabaseUserCertificate)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-self-managed-x509/#set-up-self-managed-x.509-authentication)