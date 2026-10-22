# MongoDB::Atlas::CloudProviderAccessSetup

Resource for registering the Atlas side of the AWS Cloud Provider Access trust handshake.

This resource is the first half of the two-resource split that lets a customer create an Atlas project with
Encryption at Rest using a customer managed KMS key in a single deploy. It registers the Atlas role and returns
the values that the customer's `AWS::IAM::Role` trust policy needs. The customer then creates the IAM role and
authorizes it with the separate `MongoDB::Atlas::CloudProviderAccessAuthorization` resource.
Two resources are required because the trust policy needs a value that only exists after the registration, while
Atlas needs the role ARN from the customer.

The handler never creates or manages the customer IAM role, so the resource execution role needs no `iam:*`
permissions.

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](/examples/cloud-provider-access-setup/cloud-provider-access-setup.json) for example resource.

## Delete

Deleting this resource removes the whole Cloud Provider Access role registration, not just an authorization.
Atlas refuses the delete with `AWS_IAM_ROLE_IN_USE` while a feature still uses the role, for example Encryption
at Rest. Remove the dependent feature first, then delete the resource. When the registration is already gone,
delete returns `NotFound`, which CloudFormation treats as success.

Re-creating a deleted registration is not possible for the same role id. A new deploy registers a new role with a
new role id and a new external ID, so the customer's IAM role trust policy has to be updated too.

## Import

The resource has no update handler, so it registers as `IMMUTABLE`, which supports import in principle. Automatic
import is not available because `RoleId` is a read-only part of the primary identifier, so the template cannot
supply it. Manual import is not verified for this resource yet.

## Important Links

- [Atlas Cloud Provider Access endpoints](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-cloud-provider-access)
- [Set up unified AWS access](https://www.mongodb.com/docs/atlas/security/set-up-unified-aws-access/)
