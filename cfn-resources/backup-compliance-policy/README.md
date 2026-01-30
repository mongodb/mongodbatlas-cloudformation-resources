# MongoDB::Atlas::BackupCompliancePolicy

## Description

Resource for managing [Backup Compliance Policy](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-cloud-backups/operation/updateCompliancePolicy). Backup Compliance Policy prevents any user, regardless of role, from modifying or deleting specific cluster settings, backups, and backup configurations. When enabled, the Backup Compliance Policy will be applied as the minimum policy for all clusters and backups in the project.

## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](/examples/backup-compliance-policy/README.md) for example resource.
