# Encryption At Rest

## Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.

- Encryption at rest L2 CDK constructor
- Encryption at rest L3 CDK constructor

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