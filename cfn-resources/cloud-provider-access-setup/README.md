# MongoDB::Atlas::CloudProviderAccessSetup

Resource for registering the Atlas side of the AWS Cloud Provider Access trust handshake.

This resource is the first half of the two-resource split used to give Atlas features access to AWS services,
including Encryption at Rest, S3 log export, backup export, Data Federation, and Lambda stream connections.
It registers the Atlas role and returns
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

## CloudFormation Examples

- [Setup only](/examples/cloud-provider-access-setup/cloud-provider-access-setup.json).
- [Combined setup and authorization guide](/examples/cloud-provider-access-authorization/README.md) and
  [stack template](/examples/cloud-provider-access-authorization/cloud-provider-access-stack.yaml).

The dependency chain is `Setup → AWS::IAM::Role → Authorization → consuming Atlas feature`. Reference Setup's
`AtlasAWSAccountArn` and `AtlasAssumedRoleExternalId` in the IAM trust policy, and pass the IAM role ARN and
`!GetAtt CloudProviderAccessSetup.RoleId` to Authorization. `!Ref CloudProviderAccessSetup` returns the composite
resource identifier, not the Atlas role id. If a consumer such as Encryption at Rest reads `RoleId` directly from Setup,
also give it `DependsOn: CloudProviderAccessAuthorization`. The role id alone does not establish a dependency
on authorization. These dependencies order creation and reverse that order during teardown.

Manage each registration with one Setup and one Authorization resource. Do not manage the same role concurrently
from multiple stacks or tools such as Terraform; the resources do not coordinate ownership across deployments.

See the [consumer and identifier table](/examples/cloud-provider-access-authorization/README.md#other-consumers)
for each feature's CloudFormation property. Most consumers take the Atlas role id; Lambda stream connections
take the authorized AWS IAM role ARN instead. Every consumer still needs to wait for Authorization.

## Create retries

Create registers the role with a single request and returns the registered role. Create is not idempotent: if
Atlas creates the registration but its response is lost before CloudFormation receives it, replaying the
original request creates an orphaned registration. The Atlas create endpoint accepts no idempotency key or
caller-supplied role identifier, and matching on project/provider cannot distinguish independently managed
registrations, so the handler cannot safely adopt an existing role as a retry. After an ambiguous failure,
check the project's Cloud Provider Access roles in Atlas and remove only registrations confirmed to be
unused duplicates. Making Create idempotent is separate from handling a registration already removed during Delete.

## Delete

Setup and Authorization track different stages of the same Atlas registration. Deleting either resource removes
that registration; Atlas has no operation that revokes authorization while retaining the registration.

During full stack teardown, the consuming feature is removed first, then Authorization deletes the registration.
Setup's later Delete returns `NotFound` because the registration is already gone. This is the required
[registry handler response](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-test-contract.html);
[CloudFormation ignores it during stack deletion](https://github.com/aws-cloudformation/cloudformation-cli/issues/456#issuecomment-643639639).
It does not by itself make the stack deletion fail.

Deleting only Authorization while keeping Setup leaves Setup pointing at a missing registration. Its Read then
returns `NotFound`, and a later rollback cannot restore authorization with the old role id. Treat the pair as one
lifecycle when planning removal or replacement.

Atlas refuses deletion with `AWS_IAM_ROLE_IN_USE` while a feature still uses the role, for example Encryption at
Rest. The handler reports `ResourceConflict`; `FeatureUsages` identifies current consumers. Remove or reconfigure
those consumers before retrying deletion.

Re-creating a deleted registration is not possible for the same role id: a replacement or a new deploy registers a
new role with a new role id and a new external ID, so the customer's IAM role trust policy has to be updated too.
That update requires `iam:UpdateAssumeRolePolicy` on the identity that manages the customer's IAM role.

To deliberately preserve a registration, review retention on **both** resources: retaining Setup alone does not
prevent Authorization from deleting it. See the [Authorization retention guidance](../cloud-provider-access-authorization/README.md#retaining-access).

## Import

Manual import has not been verified for this resource. It has no update handler and registers as `IMMUTABLE`,
which is eligible for registry resource import in principle. [Automatic import](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/import-resources-automatically.html)
cannot use its primary identifier because `RoleId` is read-only. Validate manual import in your account before
using it as a recovery or migration procedure; Create always registers a new role.

## Important Links

- [Atlas Cloud Provider Access endpoints](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-cloud-provider-access)
- [Set up unified AWS access](https://www.mongodb.com/docs/atlas/security/set-up-unified-aws-access/)
