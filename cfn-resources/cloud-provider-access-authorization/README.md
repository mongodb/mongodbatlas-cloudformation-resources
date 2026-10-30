# MongoDB::Atlas::CloudProviderAccessAuthorization

## Description

Authorizes an AWS IAM role against a [Cloud Provider Access](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-cloud-provider-access) role registration so Atlas can assume it.

Register the role first with the `MongoDB::Atlas::CloudProviderAccessSetup` resource, create the AWS IAM role whose trust policy uses the returned `AtlasAWSAccountArn` and `AtlasAssumedRoleExternalId`, then use this resource to authorize it. Encryption at Rest, log integration, backup export buckets, Data Federation, and AWS Lambda stream connections can then use that access. See the [consumer and identifier table](/examples/cloud-provider-access-authorization/README.md#other-consumers) for their property names and examples.

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](./docs/README.md).

## CloudFormation Examples

- [Authorize an existing registration](/examples/cloud-provider-access-authorization/cloud-provider-access-authorization.json).
- [Combined setup and authorization guide](/examples/cloud-provider-access-authorization/README.md) and
  [stack template](/examples/cloud-provider-access-authorization/cloud-provider-access-stack.yaml).

## Dependencies

Use the chain `Setup → AWS::IAM::Role → Authorization → consuming Atlas feature`. Setup supplies the Atlas principal
and external ID for the IAM trust policy; Authorization uses `!GetAtt CloudProviderAccessSetup.RoleId` and the IAM
role ARN. `!Ref CloudProviderAccessSetup` returns a composite identifier, not the Atlas role id. If Encryption
at Rest takes its role id directly from Setup, add `DependsOn: CloudProviderAccessAuthorization` to Encryption at
Rest. This ensures authorization finishes before encryption is configured, and encryption is disabled before
authorization is removed.

Manage each registration with one Setup and one Authorization resource. Do not manage the same role concurrently
from multiple stacks or tools. The duplicate-create check is not an atomic ownership claim: concurrent creates
can both observe an unauthorized registration before either authorizes it.

If several features share a registration, make each consumer depend on the same Authorization resource and
remove every consumer before deleting the pair. Removing log integration, for example, does not release a role
still used by a backup export bucket. Separate registrations and IAM roles let independently managed features
have separate permissions and deletion lifecycles.

## Create and update

- A new Create for an already authorized role returns `AlreadyExists`, even for the same ARN. It does not adopt
  or overwrite an existing authorization. A retry carrying this operation's saved callback context can recognize
  its completed authorization and succeed without another PATCH.
- `CANNOT_ASSUME_ROLE`, which Atlas can return while IAM changes propagate, is retried every ten seconds for up
  to five minutes. A timeout returns `NotStabilized`; check the trust policy, Atlas principal and external ID.
- Changing `IamAssumedRoleArn` re-authorizes the same registration in place. `ProjectId`, `RoleId` and `Profile`
  require replacement. Changing only `Profile` to another secret for the same registration cannot bypass the
  existing authorization: the replacement Create still returns `AlreadyExists`.
- Read and Update return `NotFound` when the registration is missing or unauthorized. Update does not repair
  a missing authorization by creating it.

Setup has a separate [Create retry limitation](../cloud-provider-access-setup/README.md#create-retries): a lost
registration response can leave a duplicate registration. Authorization's callback handling does not fix that.

## Delete and recovery

Setup and Authorization share one Atlas registration. Delete removes the **whole registration**, because Atlas
has no revoke-only operation. This differs from Terraform, where Authorization Delete only clears Terraform state
and Setup Delete removes the registration.

With the dependency chain above, full stack teardown removes the consumer, then Authorization, then the IAM role
and Setup. Setup finds the registration already absent and returns `NotFound`.
[CloudFormation ignores that Delete result during stack workflows](https://github.com/aws-cloudformation/cloudformation-cli/issues/456#issuecomment-643639639),
so this expected second deletion does not itself fail teardown.

Removing only Authorization leaves a surviving Setup resource pointing at a missing registration. Re-adding
Authorization or rolling back to its old role id cannot restore it. Reprovision Setup to obtain a new role id and
external ID, update the IAM trust policy and resource references, then authorize the new registration. Do not
reuse another registration merely because it belongs to the same project.

If the registration exists but was never authorized, Delete returns `NotFound` and leaves it for Setup to remove.
This matters when Create fails before authorization: rollback can clean up Setup if it was created in the same
operation; an existing or externally managed Setup remains.

If an Atlas feature still uses the role, Delete returns `ResourceConflict` for `AWS_IAM_ROLE_IN_USE`.
Use `FeatureUsages` to identify consumers, remove or reconfigure them, and retry. For Encryption at Rest, disable
that configuration before deleting the registration.

## Retaining access

To deliberately keep authorization when removing this resource, set
[`DeletionPolicy: Retain`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-deletionpolicy.html).
Use [`UpdateReplacePolicy: Retain`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-updatereplacepolicy.html)
for the old instance during replacement. These policies skip Delete; they do not revoke access or make a
replacement Create succeed for an already authorized role.

To preserve the registration through full stack deletion, retain **both** Setup and Authorization. Retaining
either one alone does not prevent the other from deleting their shared registration. Consider the IAM role and
consuming resources separately: access can continue while the IAM role's trust and permissions remain in place.
Retained resources leave CloudFormation's management and need an explicit cleanup or adoption plan.

The handler itself does not implement Terraform's no-op Delete: the
[CloudFormation contract](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-test-contract.html)
requires subsequent Read, Update and Delete calls to report `NotFound` after a successful Delete.

## Import

Manual import has not been verified for this resource. Do not use Create to adopt a retained or externally
authorized registration; it returns `AlreadyExists`. Validate an import procedure in your account before relying
on it for recovery or migration.
