# MongoDB::Atlas::LDAPVerify

## Description
Resource for managing [Teams](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-ldap-configuration)
in Atlas Organization and Atlas Project.

Requests a verification of an LDAP configuration over TLS for an Atlas project.
Pass the requestId in the response object to the Verify LDAP Configuration endpoint to get the status of a verification request. Atlas retains only the most recent request for each project.

If the cloud formation stack gets deleted, the current resource executes a validation again, in other to replace the original validation

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](./docs/README.md).

## CloudFormation Examples

See the examples [CFN Template](/examples/ldap-verify/ldap-verify.json) for example resource.
