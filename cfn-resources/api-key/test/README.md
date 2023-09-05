# api-key

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - api-key L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
- AWS Secret

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- A new api-key in Atlas Organization should be correctly configured:

[//]: # (TODO: Image to be updated)
![image]()

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Auditing)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/database-auditing/)

## Local Testing

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/api-key
./test/apikey.create-sample-cfn-request.sh YourProjectName > apikey.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.