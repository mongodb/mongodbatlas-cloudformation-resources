## MongoDB::Atlas::ServerlessInstance

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Serverless L1 CDK constructor


### Resources (and parameters for local tests) needed to manually QA:
- Atlas Organization (MONGODB_ATLAS_ORG_ID)
- Atlas Project (created as part of `cfn-testing-helper.sh`)

## Manual QA:

### Prerequisite steps:
1. Create an Atlas Organization if you don’t already have one and note the OrgId from URL (https://cloud.mongodb.com/v2#/org/<MONGODB_ATLAS_ORG_ID>/projects)
2. Export MONGODB_ATLAS_ORG_ID environment variable with OrgId from #1.

### Steps to test:
1. Ensure prerequisites above for this resource and general [prerequisites](../../../TESTING.md#prerequisites) are complete.
2. Follow [general steps](../../../TESTING.md#steps) to test a CFN resource.


### Success criteria when testing the resource
1. A serverless instance should show correctly configured for the database:

![image](https://user-images.githubusercontent.com/122359335/228200365-6e5950d8-1284-426c-97c8-57a6b24181d6.png)

2. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-serverless-instances)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/tutorial/create-serverless-instance/)

## Local Testing
The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/project
./test/project.create-sample-cfn-request.sh YourProjectName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE and DELETE tests must pass.
