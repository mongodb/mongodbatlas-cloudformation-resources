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
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- You should see the option "Encryption at Rest using your Key Management" enabled in the "Advanced" page:
![image](https://user-images.githubusercontent.com/5663078/227896265-7e489e9e-2666-4faa-8d10-5c8b3ee77620.png)
## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Encryption-at-Rest-using-Customer-Key-Management/operation/updateEncryptionAtRest)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-kms-encryption/)


## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
set or export below environment variables 
export KMS_KEY=<<CustomerMasterKeyID>>
export KMS_ROLE=<<RoleID>>
export KMS_REGION=<<key region>>
cd ${repo_root}/cfn-resources/encryption-at-rest
./test/encryptionatrest.create-sample-cfn-request.sh YourProjectID > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE and DELETE tests must pass.
