# MongoDB::Atlas::Trigger

## Description
Resource for managing [Triggers](https://www.mongodb.com/docs/atlas/app-services/admin/api/v3/#tag/triggers).

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](../../examples/trigger/trigger.json) for example resource.

## Development
```shell
uvx atlas-init@0.4.0
atlas-init init
cd cfn-resources/trigger
atlas-init apply
atlas-init cfn contract-test # use --help to see more options
atlas-init destroy
```
