# Cloud Backup Snapshots


## CFN resource type used
- MongoDB::Atlas::CloudBackupSnapshot

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Cloud Backup Snapshot CFN resource](../../../../cfn-resources/cloud-backup-snapshot/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- Backup snapshot for the Atlas Cluster should be shown in the "Snapshots" page:
![image](https://user-images.githubusercontent.com/5663078/227233348-ea32d93a-bfc6-468a-b111-fb12bc0a50ec.png)


## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Cloud-Backups-Schedule)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/restore-from-snapshot/)