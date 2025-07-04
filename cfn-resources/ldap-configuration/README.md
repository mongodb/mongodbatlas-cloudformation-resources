# MongoDB::Atlas::LDAPConfiguration

## Description
Resource for managing [LDAP Configurations](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-ldap-configuration).

An LDAP configuration defines settings for Atlas to connect to your LDAP server over TLS for user authentication and authorization.
Your LDAP server must be visible to the internet or connected to your Atlas cluster with VPC Peering.
In addition, your LDAP server must use TLS.

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](./docs/README.md).

## CloudFormation Examples

See the examples [CFN Template](/examples/ldap-configuration/ldap-configuration.json) for example resource.
