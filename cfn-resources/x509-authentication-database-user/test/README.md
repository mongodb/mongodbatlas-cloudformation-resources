## MongoDB::Atlas::X509AuthenticationDatabaseUser

### Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- X509 authentication database user L1 CDK constructor




### Resources (and parameters for local tests) needed to manually QA:
CustomerX509 is to be manually provided.
- Atlas project (created by cfn-test-create-inputs.sh)
- CustomerX509



## Manual QA:

### Prerequisite steps:
1. Create a test certificate: One of the ways to create a test certificate is to install OpenSSL and run the following command (source):
```
openssl req -new -newkey rsa:1024 -days 365 -nodes -x509 -keyout test.key -out test.cert
```
A certificate for you will be generated in test.cert file.

2. Use cert from #1 and update cert in inputs/inputs_1_create.json before running cfn test and when creating the stack in AWS CloudFormation in the steps below.


### Steps to test:
1. Follow general [prerequisites](../../../TESTING.md#prerequisites) for testing CFN resources.
2. Follow [general steps](../../../TESTING.md#steps) to test CFN resources.

### Success criteria when testing the resource
1. X.509 Authentication should be enabled correctly set up in your Atlas Project as per configuration specified in the inputs/example. This can be found under Security section -> Advanced:

   ![image](https://user-images.githubusercontent.com/122359335/227374480-1afa48a4-5265-4a2a-ad92-067f5015eeca.png)

2. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-createdatabaseusercertificate)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-self-managed-x509/#set-up-self-managed-x.509-authentication)