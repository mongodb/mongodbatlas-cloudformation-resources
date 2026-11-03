# Cloud Provider Access Authorization examples

[cloud-provider-access-stack.yaml](./cloud-provider-access-stack.yaml) shows the complete AWS trust setup for project Encryption at Rest using a customer-managed KMS key. It creates the Atlas registration, an AWS IAM role, its Atlas authorization, and the project encryption configuration. The Atlas project and KMS key already exist.

Use a project whose Encryption at Rest configuration is not managed by another stack, Terraform, or another owner. This example manages that project-wide configuration and disables it when the stack is deleted. It uses public KMS connectivity and leaves dedicated Search Node encryption disabled. It does not create clusters; configure customer key management on each eligible cluster separately, following the [Atlas instructions](https://www.mongodb.com/docs/atlas/security/aws-kms-over-public-network/).

Before deploying, activate or privately register `MongoDB::Atlas::CloudProviderAccessSetup`, `MongoDB::Atlas::CloudProviderAccessAuthorization`, and `MongoDB::Atlas::EncryptionAtRest` in the stack's AWS account and Region. Configure an [Atlas credential profile](../../README.md#mongodb-atlas-api-keys-credential-management) with the project permissions required by these resources.

Supply these parameters:

| Parameter | Value |
| --- | --- |
| `ProjectId` | Existing Atlas project ID. |
| `Profile` | Secrets Manager Atlas credential profile; defaults to `default`. |
| `KmsKeyArn` | Full ARN of an enabled symmetric customer-managed KMS key in the stack's AWS account. |
| `KmsKeyRegion` | The key's Region in Atlas format, such as `EU_WEST_1` for `eu-west-1`. It must match the Region in `KmsKeyArn`. |

The IAM policy grants only `kms:DescribeKey`, `kms:Encrypt`, and `kms:Decrypt` on the supplied key, as required by [Atlas](https://www.mongodb.com/docs/atlas/security/aws-kms-over-public-network/). The existing KMS key policy must allow this account to delegate those permissions through IAM. An IAM policy alone cannot grant access if the key policy does not enable it; see [AWS KMS key policies](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html). This template does not modify the key policy. Key policy conditions, permission boundaries, and organization policies must also permit the access.

After setting the variables below to your values, deploy from this directory:

```bash
aws cloudformation deploy \
  --template-file cloud-provider-access-stack.yaml \
  --stack-name atlas-cloud-provider-access \
  --region "$AWS_REGION" \
  --capabilities CAPABILITY_IAM \
  --parameter-overrides \
    ProjectId="$ATLAS_PROJECT_ID" \
    Profile="$ATLAS_PROFILE" \
    KmsKeyArn="$KMS_KEY_ARN" \
    KmsKeyRegion="$ATLAS_KMS_REGION"
```

`CAPABILITY_IAM` acknowledges creation of the IAM role. CloudFormation generates its name, so `CAPABILITY_NAMED_IAM` is unnecessary. The stack Region uses AWS format, such as `eu-west-1`; `ATLAS_KMS_REGION` uses Atlas format, such as `EU_WEST_1`.

The references create this order: registration, IAM role, authorization, then encryption configuration. `EncryptionAtRest` explicitly depends on `CloudProviderAccessAuthorization`: getting `RoleId` from setup alone would allow encryption configuration to start before authorization finishes. The reverse order applies during deletion. Authorization deletion currently removes the Atlas registration; setup cleanup then encounters an already absent registration. See the [resource lifecycle documentation](../../cfn-resources/cloud-provider-access-authorization/README.md#delete-and-recovery) for deletion and rollback limits.

Before deleting this stack after using its configuration, move dependent clusters off customer key management and wait for their changes to finish. The project and KMS key are outside the stack and remain. Preserve the key and arrange access for any backups that still require it; [Atlas backups retain the key used when they were created](https://www.mongodb.com/docs/atlas/security-aws-kms/).

For a registration and IAM role that already exist, use the smaller [cloud-provider-access-authorization.json](./cloud-provider-access-authorization.json) example. Its trust policy must already use that registration's Atlas principal and external ID, and the registration must not already be authorized. Deleting that example's stack also removes the registration, including one created outside CloudFormation.

## Other consumers

The trust handshake is shared by these Terraform and CloudFormation resources. The table shows the AWS variants
supported by the CloudFormation schemas in this repository. The KMS permissions in the example above are specific
to Encryption at Rest; use the permissions required by each consuming feature.

| Consumer | Terraform resource and field | CloudFormation resource and field | Guide |
| --- | --- | --- | --- |
| Encryption at Rest | `mongodbatlas_encryption_at_rest.aws_kms_config.role_id` | `MongoDB::Atlas::EncryptionAtRest.AwsKmsConfig.RoleID` | [Complete KMS stack](./cloud-provider-access-stack.yaml) |
| S3 log export | `mongodbatlas_log_integration.iam_role_id` | `MongoDB::Atlas::LogIntegration.IamRoleId` | [Complete log-export stack](../log-integration/README.md) |
| Backup export destination | `mongodbatlas_cloud_backup_snapshot_export_bucket.iam_role_id` | `MongoDB::Atlas::CloudBackupSnapshotExportBucket.IamRoleID` | [Backup export setup](../../cfn-resources/cloud-backup-snapshot-export-bucket/README.md) |
| Data Federation AWS access | `mongodbatlas_federated_database_instance.cloud_provider_config.aws.role_id` | `MongoDB::Atlas::FederatedDatabaseInstance.CloudProviderConfig.RoleId` | [Configuration and schema limits](../../cfn-resources/federated-database-instance/README.md) |
| AWS Lambda stream connection | `mongodbatlas_stream_connection.aws.role_arn` | `MongoDB::Atlas::StreamConnection.Aws.RoleArn` | [Lambda connection](../atlas-streams/stream-connection/README.md) |

The first four consumers take the **Atlas registration ID**, obtained with
`!GetAtt CloudProviderAccessSetup.RoleId`. Lambda takes the **AWS IAM role ARN**, obtained from the IAM role that
Authorization authorized. Neither reference by itself depends on Authorization, so add
`DependsOn: CloudProviderAccessAuthorization` to each consuming resource.

Terraform's deprecated `mongodbatlas_push_based_log_export` also takes `iam_role_id`. Use the modern
`MongoDB::Atlas::LogIntegration` resource for the CloudFormation example; there is no separate legacy
PushBasedLogExport resource in this repository.

Backup export jobs and automatic export schedules are indirect consumers: they reference the export bucket ID,
not the Cloud Provider Access role ID. In CloudFormation, use the export bucket's `Id` for
`CloudBackupSchedule.Export.ExportBucketId`. Remove automatic export references before deleting the export bucket;
see the backup guide for the required separate update. There is no CloudFormation export-job resource here.

Terraform also supports Azure and GCP access configurations, including Stream Processing connections, that these
AWS-only Setup and Authorization resources cannot create. Do not copy those examples without checking the
CloudFormation schemas.

Use a separate registration and IAM role for features that need independent permissions or teardown. If features
share a registration, give all of them a dependency on the same Authorization resource and remove all consumers
before deleting it. Do not add a second Authorization resource for the same registration. Deleting either Atlas
access resource removes the registration for every consumer; exported S3 objects are separate from that registration.
