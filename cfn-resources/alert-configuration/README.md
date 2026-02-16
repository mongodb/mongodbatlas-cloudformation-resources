# MongoDB::Atlas::AlertConfiguration

## Description

Resource for managing [Alert](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-alerts) conditions.

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes & Parameters

See the [resource docs](docs/README.md).

## CloudFormation Examples

See the example CFN templates in the [examples directory](/examples/alert-configuration/):

- [alert-configuration-email.json](/examples/alert-configuration/alert-configuration-email.json) - Email alert for host metrics
- [alert-configuration-microsoft-teams.json](/examples/alert-configuration/alert-configuration-microsoft-teams.json) - Microsoft Teams alert for host metrics
- [alert-configuration-stream-processor.json](/examples/alert-configuration/alert-configuration-stream-processor.json) - Email alert for Stream Processor metrics
