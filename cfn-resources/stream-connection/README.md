# MongoDB::Atlas::StreamConnection

Resource for creating and managing [Connections for an Atlas Stream Instance](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-createstreamconnection).

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md). Also refer [AWS security best practices for CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds) to manage credentials.

## Cloudformation Examples

See the example [CFN Template](/examples/stream-connection/kafka-stream-connection.json) for example resource.
