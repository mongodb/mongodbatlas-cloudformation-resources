# Cloud Backup Snapshots Export Bucket


## CFN resource type used
- MongoDB::Atlas::CloudBackupSnapshotExportBucket

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Cloud Backup Snapshot Export Bucket CFN resource](../../../../cfn-resources/cloud-backup-snapshot-export-bucket/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
-  The endpoint `api/atlas/v1.0/groups/<ProjectId>/backup/exportBuckets` should return your bucket:

```bash
curl --user "<PublicKey>:<PrivateKey>" -X GET --digest \
     --header "Accept: application/json" \
     --header "Content-Type: application/json" \
     "https://cloud.mongodb.com/api/atlas/v1.0/groups/<ProjectId>/backup/exportBuckets?pretty=true"

{
  "links" : [ {
    "href" : "https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/6414908c207f4d22f4d8f232/backup/exportBuckets?pretty=true&pageNum=1&itemsPerPage=100",
    "rel" : "self"
  } ],
  "results" : [ {
    "_id" : "641dddc051ed5c6792399422",
    "bucketName" : "andrea-angiolillo-mongocli",
    "cloudProvider" : "AWS",
    "iamRoleId" : "641dd86151ed5c67923984f1"
  } ],
  "totalCount" : 1
}


## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Cloud-Backups/operation/listExportBuckets)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/data-federation/config/config-aws-s3/)