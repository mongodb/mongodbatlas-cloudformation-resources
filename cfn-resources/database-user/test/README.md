# Database User

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Database user L1 CDK constructor
 - Atlas basis L3 CDK constructor
 - Encryption at rest L3 CDK constructor
 - Atlas Quickstart
 - Atlas Quickstart Fargate


## Prerequisites 
### Resources needed to run the manual QA
- Atlas project


All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- The Database User should be visible in the "Database Users" page:
![image](https://user-images.githubusercontent.com/5663078/227314604-d15f10a4-5e3b-4010-b94f-621ec55eceb3.png)
## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-database-users)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-add-mongodb-users/)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/database-user
./test/databaseuser.create-sample-cfn-request.sh YourProjectID SomeIAMRole > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.