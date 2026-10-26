# MongoDB::Atlas::CloudProviderAccessAuthorization

## Description

Authorizes an AWS IAM role against a [Cloud Provider Access](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-cloud-provider-access) role registration so Atlas can assume it.

Register the role first with the `MongoDB::Atlas::CloudProviderAccessSetup` resource, create the AWS IAM role whose trust policy uses the returned `AtlasAWSAccountArn` and `AtlasAssumedRoleExternalId`, then use this resource to authorize it. The existing [`MongoDB::Atlas::EncryptionAtRest`](../encryption-at-rest/README.md) resource can then use the role id.

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](./docs/README.md).

## CloudFormation Examples

See the examples [CFN Template](/examples/cloud-provider-access-authorization/cloud-provider-access-authorization.json) for an example resource.

## Behaviour to be aware of

- Deleting this resource removes the whole Cloud Provider Access role registration, not only the authorization. Atlas has no operation that revokes an authorization while keeping the role registered, so the paired `MongoDB::Atlas::CloudProviderAccessSetup` resource is left orphaned after a delete.
- Re-creating after a delete is not supported for the same role. It means registering the role again with `MongoDB::Atlas::CloudProviderAccessSetup`, which produces a new role id and a new external id, and the IAM trust policy has to be updated too.
- Deleting a role that is registered but not authorized reports `NotFound` and leaves the registration in place, because the registration belongs to `MongoDB::Atlas::CloudProviderAccessSetup`. A stack rollback removes it through that resource; a registration created outside the stack remains in Atlas until it is removed there.
- A role that a feature uses cannot be removed. The delete fails with `ResourceConflict` for `AWS_IAM_ROLE_IN_USE` and the dependent feature has to be removed first, for example by disabling Encryption at Rest. When Encryption at Rest gets its role id from the setup resource, add `DependsOn: CloudProviderAccessAuthorization` to the Encryption at Rest resource. That explicit dependency ensures authorization completes before enabling encryption and that encryption is disabled before removing authorization during teardown.
- Changing `IamAssumedRoleArn` re-authorizes the same role in place. The resource is not replaced, the role id does not change, and the paired setup resource keeps its registration.
- Creating this resource for a role that is already authorized fails with `AlreadyExists`, even if the requested ARN differs. The role registration identifies the authorization; a second resource must not overwrite it.
- Create checkpoints its callback state before changing Atlas. A repeated callback that finds its requested ARN already authorized succeeds without another write. A fresh create still returns `AlreadyExists`.
- `CANNOT_ASSUME_ROLE` is retried through callbacks every ten seconds for up to five minutes. No handler sleeps between attempts. If the trust policy never becomes usable, the operation fails with `NotStabilized` and guidance to check the policy and external ID.
- Update returns `NotFound` when the registration has no authorization. It does not create an authorization that Read considers absent.
- `Profile` is part of the primary identifier so Read and Delete can use the original credentials even when CloudFormation supplies only identifying properties.

Deleting this resource while retaining the paired `MongoDB::Atlas::CloudProviderAccessSetup` resource also affects stack rollback: a later stack failure cannot restore the deleted authorization using the old role id.

## CloudFormation contract tests

The v1 contract test `contract_delete_create` expects a deleted resource to be creatable again. Because the delete removes the whole role registration, the re-create fails with `CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND`. This deviation is expected and is pending an exception request to AWS.
