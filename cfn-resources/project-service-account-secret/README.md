# MongoDB::Atlas::ProjectServiceAccountSecret

## Description

The Project Service Account Secret resource provides a secret for a Service Account at the project level. This resource lets you create and delete secrets for Project Service Accounts. For more information, see [Create One Project Service Account Secret](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Service-Accounts/operation/createGroupServiceAccountSecret) in the MongoDB Atlas API documentation.

~> **IMPORTANT WARNING:** Managing Service Account Secrets with CloudFormation **exposes sensitive organizational secrets** in CloudFormation's outputs and logs. We suggest following [AWS Secrets Manager best practices](https://docs.aws.amazon.com/secretsmanager/latest/userguide/best-practices.html) for handling sensitive data.

-> **NOTE:** This resource does not support updates. Any property change will trigger a replacement (delete + create). To rotate secrets, simply replace the resource.

## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](/examples/project-service-account-secret/README.md) for example resource.

## Important Notes

- The `Secret` property contains the actual secret value and is only returned once during creation
- This resource does not support updates - any change will trigger a replacement
- Secrets have an expiration time controlled by `SecretExpiresAfterHours`
- The minimum and maximum expiration times are controlled by your organization's settings
