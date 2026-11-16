# MongoDB::Atlas::CloudBackupSnapshotExportBucket

## Description

Registers an existing AWS S3 bucket as an Atlas destination for [Cloud Backup snapshot exports](https://www.mongodb.com/docs/atlas/backup/cloud-backup/export/aws-s3/). This resource manages the Atlas export-bucket configuration; it does not create the S3 bucket or start a snapshot export.

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

Register and authorize an AWS IAM role with [`CloudProviderAccessSetup`](../cloud-provider-access-setup/README.md) and [`CloudProviderAccessAuthorization`](../cloud-provider-access-authorization/README.md). Give that IAM role `s3:GetBucketLocation` on the bucket ARN and `s3:PutObject` on the bucket's objects, following the [Atlas bucket access policy](https://www.mongodb.com/docs/atlas/backup/cloud-backup/export/aws-s3/?interface=atlas-cli).

Set `IamRoleID` to the 24-character Atlas registration ID returned by `!GetAtt CloudProviderAccessSetup.RoleId`. The AWS IAM role ARN belongs in Authorization's `IamAssumedRoleArn`; `!Ref CloudProviderAccessSetup` returns a composite identifier. When using Setup's role ID here, add `DependsOn: CloudProviderAccessAuthorization` so bucket registration waits for authorization. Also depend on any separately attached IAM policy that grants bucket access.

## Attributes & Parameters

See the [resource docs](docs/README.md).

## CloudFormation Examples

- [Setup, authorization, bucket permissions, and backup schedule wiring](../../examples/cloud-backup-snapshot-export-bucket/README.md).
- [Standalone template for an existing authorized role and bucket](../../examples/cloud-backup-snapshot-export-bucket/CloudBackupSnapshotExportBucket.json).

## Delete

Before removing this destination, disable automatic export on every cluster whose backup schedule uses it. Complete that schedule update before deleting the export-bucket resource, as required by the [Atlas delete API](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-deletegroupbackupexportbucket). Remove the export-bucket configuration before deleting its Cloud Provider Access authorization or setup registration. The dependency described above provides that order within one stack.

Deleting this Atlas configuration leaves the S3 bucket and exported objects in place. See [Atlas snapshot export management](https://www.mongodb.com/docs/atlas/backup/cloud-backup/export/#delete-a-snapshot-export-bucket).
