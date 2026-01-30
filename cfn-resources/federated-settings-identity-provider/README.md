# MongoDB::Atlas::FederatedSettingsIdentityProvider

## Description

The federated settings identity provider resource provides access to your Atlas
federated authentication identity providers (SAML and OIDC). It lets you
create, edit, and delete identity providers within an Atlas federation.

## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

Examples for this resource will be added in `/examples/`.

## Contract Testing

Contract testing requires a valid Federation Settings ID
export MONGODB_ATLAS_FEDERATION_SETTINGS_ID="your-federation-settings-id"

# Run contract tests

make create-test-resources
cfn test -- -k contract_create_delete
make delete-test-resources
