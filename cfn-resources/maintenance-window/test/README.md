# MongoDB::Atlas::MaintenanceWindow

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Maintenance window L1 CDK constructor



## Prerequisites 
### Resources needed to manually QA
- Atlas Project

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Atlas Project should show configured maintenance window correctly under Project Settings:

![image](https://user-images.githubusercontent.com/122359335/227540482-6f021ea1-7b7e-4fbf-b883-1d9e0e2eea9a.png)

2. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-maintenance-windows)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/tutorial/cluster-maintenance-window/)

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

Both CREATE & DELETE tests must pass.