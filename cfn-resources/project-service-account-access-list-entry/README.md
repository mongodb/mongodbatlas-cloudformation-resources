# MongoDB::Atlas::ProjectServiceAccountAccessListEntry

## Description

The Project Service Account Access List Entry resource manages IP access list entries for MongoDB Atlas Project Service Accounts. This resource lets you create, read, delete, and list IP access list entries at the project level. For more information, see [Create One Project Service Account Access List Entry](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Service-Accounts/operation/createAccessList) in the MongoDB Atlas API documentation.

-> **NOTE:** This resource does not support updates. Any property change will trigger a replacement (delete + create).

## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the example [CFN Template](/examples/project-service-account-access-list-entry/README.md) for example resource.

## Important Notes

- You must specify either `CIDRBlock` or `IPAddress`, but not both
- When you specify an IP address, Atlas automatically generates a `/32` CIDR block
- This resource does not support updates - any change will trigger a replacement
- Access list entries are identified by the combination of `ProjectId`, `ClientId`, and `CIDRBlock`
- The List operation returns all access list entries for a given project service account
