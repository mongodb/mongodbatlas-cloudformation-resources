# Alert Configurations 

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Alert Configuration L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- Alert Settings for the respective Project in Atlas should be correctly configured:
![image](https://user-images.githubusercontent.com/5663078/226870968-9ef8ae46-b0cf-462b-ac62-7229d2d79ac0.png)

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-listalertconfigurations)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/configure-alerts/#configure-an-alert)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/alert-configuration
./test/alert-configuration.create-sample-cfn-request.sh YourProjectID  > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.
