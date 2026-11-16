# MongoDB::Atlas::LogIntegration

## Description

The log integration resource creates, updates, and deletes AWS S3 log export integrations for MongoDB Atlas. It uses the `/groups/{groupId}/logIntegrations` API to export selected log types at one-minute intervals.

This resource manages the current log integration feature, not the deprecated `/pushBasedLogExport` configuration. See [Export Logs to AWS S3 Buckets](https://www.mongodb.com/docs/atlas/export-logs-external-sinks/) and the [legacy migration guidance](https://www.mongodb.com/docs/atlas/push-logs/#migrate-to-external-log-sinks).

## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).

For S3 export, provide an existing bucket and an authorized Atlas Cloud Provider Access role with access to that bucket. `IamRoleId` is the 24-character Atlas registration ID returned by `MongoDB::Atlas::CloudProviderAccessSetup`; it is not an AWS IAM role ARN or AWS role ID. The AWS IAM role and bucket must belong to the same AWS account. See the [Cloud Provider Access guide](/examples/cloud-provider-access-authorization/README.md) for registration, trust, and authorization.

When creating these resources in one stack, set `IamRoleId` from `CloudProviderAccessSetup.RoleId` and add `DependsOn: CloudProviderAccessAuthorization` to the log integration. The explicit dependency ensures authorization finishes before Atlas configures export and removes the log integration before authorization during deletion. A role still used by any other feature can prevent authorization or setup deletion. Deleting the log integration stops that integration; it does not delete the S3 bucket or exported objects.

## Attributes and Parameters

See the [resource docs](docs/README.md).

## CloudFormation Examples

See the [examples](/examples/log-integration/README.md) for an existing-role template and a complete setup, IAM role, authorization, and S3 log integration template.
