# @mongodbatlas-awscdk/atlas-encryption-at-rest

## CFN resource type used
- MongoDB::Atlas::EncryptionAtRest

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Encryption At Rest CFN resource](../../../../cfn-resources/encryption-at-rest/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- You should see the option "Encryption at Rest using your Key Management" enabled in the "Advanced" page:
![image](https://user-images.githubusercontent.com/5663078/227896265-7e489e9e-2666-4faa-8d10-5c8b3ee77620.png)

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Encryption-at-Rest-using-Customer-Key-Management/operation/updateEncryptionAtRest)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-kms-encryption/)
