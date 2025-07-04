# MongoDB::Atlas::PrivateEndpointAWS

Resource for creating and managing [Private Endpoint Services](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-private-endpoint-services).

# V2 Migration Guideline

For migrating from the unified Private endpoint (V1) to the divided (V2) follow the [V2 Upgrade Guide](../private-endpoint/upgradeguidev2/V2-UpgradeGuide.md)

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](/examples/private-endpoint/privateEndpointV2.json) for example resource.
