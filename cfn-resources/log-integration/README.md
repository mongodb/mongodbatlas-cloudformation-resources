# MongoDB::Atlas::LogIntegration

## Description

The log integration resource provides access to push-based log export configurations for MongoDB Atlas. The resource allows you to create, edit and delete log export integrations to AWS S3 buckets. Push-based log export enables you to automatically export MongoDB Atlas logs to your AWS S3 bucket with 1-minute frequency.

For more information, see [Push Logs to AWS S3 bucket](https://www.mongodb.com/docs/atlas/push-logs/) and the [Push-Based Log Export API](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-push-based-log-export).

## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## CloudFormation Examples

See the example [CFN Template](/examples/log-integration/README.md) for example resource.
