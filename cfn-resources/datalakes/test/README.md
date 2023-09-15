# Datalake

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Datalake L1 CDK constructor



## Prerequisites 
### Resources needed to run the manual QA
- Atlas project
- Atlas Cluster
- AWS role
- AWS role policy



All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- You should see your data federation in the "Data Federation" page:
![image](https://user-images.githubusercontent.com/5663078/227923171-b9aa0067-d8a0-41b3-94c5-9e5ce5ea222c.png)
## Important Links
- [API Documentation](https://docs-atlas-staging.mongodb.com/cloud-docs/docsworker-xlarge/openapi-docs-test/reference/api-resources-spec/#tag/Data-Federationt)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/data-federation/config/config-adl-datasets/)

## Unit Testing Locally
The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/datalakes
./test/datalakes.create-sample-cfn-request.sh YourProjectID ExternalId IamAssumedRoleARN IamUserARN RoleId TestS3BucketId > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.