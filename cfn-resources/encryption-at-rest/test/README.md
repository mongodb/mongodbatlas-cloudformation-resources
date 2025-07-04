# Encryption at Rest

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Encryption at rest L1 CDK constructor
 - Encryption at rest L2 CDK constructor
 - Encryption at rest L3 CDK constructor



## Prerequisites 
### Resources needed to run the manual QA
- Atlas project
- AWS role
- AWS role policy



All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- You should see the option "Encryption at Rest using your Key Management" enabled in the "Advanced" page:
![image](https://user-images.githubusercontent.com/5663078/227896265-7e489e9e-2666-4faa-8d10-5c8b3ee77620.png)
## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-updateencryptionatrest)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-kms-encryption/)
