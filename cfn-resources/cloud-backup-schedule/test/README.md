# Cloud backup schedule

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Cloud backup schedule L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
- Cluster with backup enabled
- PolicyId



All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- Backup policy should be updated in the "Backup" page of your cluster:
![image](https://user-images.githubusercontent.com/5663078/227544843-152b52ee-2c23-40db-b8bd-1391ef64aebc.png)
## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-cloud-backups)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/overview/#std-label-backup-cloud-provider)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```

then in another shell:

```
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/cloud-backup-schedule
./test/cloud-backup-schedule.create-sample-cfn-request.sh YourProjectID ClusterName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke resource CREATE test.request.json 
cfn invoke resource DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.
