# Cloud Backup snapshot export bucket

[`CloudBackupSnapshotExportBucket.json`](./CloudBackupSnapshotExportBucket.json) registers an existing S3 bucket using an already authorized Atlas Cloud Provider Access role. Its `IamRoleID` parameter takes the 24-character Atlas role registration ID. `BucketName` takes the S3 bucket name.

For a template that also manages the role, follow the [Setup and Authorization guide](../cloud-provider-access-authorization/README.md) for the Atlas registration and IAM trust policy. Use a role dedicated to backup exports and grant the S3 permissions below. Authorization receives the AWS IAM role ARN; the export-bucket resource receives `!GetAtt CloudProviderAccessSetup.RoleId`.

## Bucket permissions and authorization dependency

The following entries belong under `Resources` in a template that already defines:

- `CloudProviderAccessSetup`, an Atlas registration for the project.
- `AtlasExportRole`, an `AWS::IAM::Role` whose trust policy uses that registration's `AtlasAWSAccountArn` and `AtlasAssumedRoleExternalId`.
- `CloudProviderAccessAuthorization`, which authorizes `!GetAtt AtlasExportRole.Arn` against `!GetAtt CloudProviderAccessSetup.RoleId`.
- String parameters `ProjectId`, `Profile`, and `BackupBucketName`; the S3 bucket already exists.

```yaml
BackupExportAccessPolicy:
  Type: AWS::IAM::Policy
  Properties:
    PolicyName: AtlasBackupExport
    Roles:
      - !Ref AtlasExportRole
    PolicyDocument:
      Version: "2012-10-17"
      Statement:
        - Effect: Allow
          Action: s3:GetBucketLocation
          Resource: !Sub arn:${AWS::Partition}:s3:::${BackupBucketName}
        - Effect: Allow
          Action: s3:PutObject
          Resource: !Sub arn:${AWS::Partition}:s3:::${BackupBucketName}/*

BackupExportBucket:
  Type: MongoDB::Atlas::CloudBackupSnapshotExportBucket
  DependsOn:
    - CloudProviderAccessAuthorization
    - BackupExportAccessPolicy
  Properties:
    ProjectId: !Ref ProjectId
    Profile: !Ref Profile
    IamRoleID: !GetAtt CloudProviderAccessSetup.RoleId
    BucketName: !Ref BackupBucketName
```

These are the bucket permissions in the [Atlas export prerequisites](https://www.mongodb.com/docs/atlas/backup/cloud-backup/export/aws-s3/?interface=atlas-cli) and the [Terraform AWS backup export example](https://github.com/mongodb/terraform-provider-mongodbatlas/tree/master/examples/mongodbatlas_cloud_backup_snapshot_export_bucket/aws). They are separate from the role's trust policy. The explicit dependencies wait for both authorization and policy attachment before Atlas checks bucket access, and preserve them until the export-bucket configuration is removed.

## Automatic and individual exports

Bucket registration alone does not export snapshots. To enable automatic exports, merge these properties into an existing [`MongoDB::Atlas::CloudBackupSchedule`](../../cfn-resources/cloud-backup-schedule/README.md), keeping its project, cluster, profile, and backup policy settings:

```yaml
AutoExportEnabled: true
Export:
  ExportBucketId: !GetAtt BackupExportBucket.Id
  FrequencyType: daily
```

`BackupExportBucket.Id` is the Atlas export-bucket configuration ID, distinct from the Cloud Provider Access role ID and S3 bucket name. The reference makes the schedule depend on bucket registration. Use a frequency that has a matching backup policy item; this example exports daily snapshots.

For one snapshot, use an [Atlas snapshot export job](https://www.mongodb.com/docs/atlas/cli/current/command/atlas-backups-exports-jobs-create/) with the export-bucket ID and snapshot ID. This repository does not provide a CloudFormation snapshot export job resource. Terraform provides a separate [`mongodbatlas_cloud_backup_snapshot_export_job` resource](https://github.com/mongodb/terraform-provider-mongodbatlas/tree/master/examples/mongodbatlas_cloud_backup_snapshot_export_job).

Terraform's export-job Delete only removes the job from state; it does not cancel an active export. Wait for
outstanding exports to finish before retiring their destination or IAM access, including jobs started outside
CloudFormation.

## Teardown

First update every backup schedule using this destination to disable automatic export, and let that update finish. Then remove `BackupExportBucket` before removing the authorization, IAM role, or setup registration. This includes schedules managed outside this stack; see the [Atlas export-bucket deletion requirement](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-deletegroupbackupexportbucket).

The Atlas export-bucket resource does not delete the S3 bucket or its exported data. The S3 bucket in this example is outside the stack. Authorization deletion currently removes the whole Cloud Provider Access registration; see the [shared role lifecycle and recovery limits](../../cfn-resources/cloud-provider-access-authorization/README.md#delete-and-recovery) before removing only part of the configuration.
