# MongoDB::Atlas::APIKey

## Description
Resource for managing [APIKeys](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-programmatic-api-keys).
The API keys created as result, will be securely stored in AWS Secret Manager, ensuring the highest level of data protection and access control.
## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](/examples/api-key/api-key.json) for example resource.
