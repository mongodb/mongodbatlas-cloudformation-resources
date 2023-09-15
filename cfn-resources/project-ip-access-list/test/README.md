# Project IP Access List 

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- Project IP access list L1 CDK constructor
- Atlas Basic L3 CDK constructor
- Encryption at rest L3 CDK constructor
- Atlas Quickstart


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
All resources are created as part of `cfn-testing-helper.sh`


## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- A new entry should be added to the "Network Access" page:
![image](https://user-images.githubusercontent.com/5663078/227484402-9189af3d-a3f0-4bde-a288-9ee847e6eeab.png)
## Important Links
- [API Documentation](https://docs-atlas-staging.mongodb.com/cloud-docs/docsworker-xlarge/openapi-docs-test/reference/api-resources-spec/#operation/createProjectIpAccessList)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security/ip-access-list/)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/project-ip-access-list
./test/projectipaccesslist.create-sample-cfn-request.sh YourProjectID > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE and DELETE tests must pass.
