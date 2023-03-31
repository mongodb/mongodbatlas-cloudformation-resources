# Cloud Backup Schedule


## CFN resource type used
- MongoDB::Atlas::CloudBackupSchedule

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Cloud Backup Schedule CFN resource](../../../../cfn-resources/cloud-backup-schedule/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- Backup policy should be updated in the "Backup" page of your cluster:
![image](https://user-images.githubusercontent.com/5663078/227544843-152b52ee-2c23-40db-b8bd-1391ef64aebc.png)


## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Cloud-Backups-Schedule)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/overview/#std-label-backup-cloud-provider)