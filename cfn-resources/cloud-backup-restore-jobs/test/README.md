# Cloud backup restore jobs 

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Cloud backup restore jobs  L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
- Cluster with backup enabled
- SnapshotId


All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- Backup restore job for the Atlas Cluster should be shown in "Restores & Downloads" page:
![image](https://user-images.githubusercontent.com/5663078/227225795-0f1b6650-95fe-40ca-942d-99902b747aa2.png)


## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-cloud-backups)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/restore-overview/)

## Unit Testing Locally
The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/cloud-backup-restore-jobs
./test/cloud-backup-restore-job.create-sample-cfn-request.sh YourProjectID ClusterName SnapshotId > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.