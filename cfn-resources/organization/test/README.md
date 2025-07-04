# Organization 

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Auditing L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Organization

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-organizations)

## Local Testing

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/organization
./test/org.create-sample-cfn-request.sh YourOrgName > org.sample-cfn-request.json
echo "Sample request:"
cat test.request.json
cfn invoke resource CREATE /test/org.sample-cfn-request.json
cfn invoke resource DELETE /test/org.sample-cfn-request.json 
```
Once tested you can clean up by running [org.delete-sample-cfn-request.sh](org.delete-sample-cfn-request.sh)

```
./test/org.delete-sample-cfn-request.sh 
```

Both CREATE & DELETE tests must pass.