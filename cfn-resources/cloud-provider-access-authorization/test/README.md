# Cloud Provider Access Authorization

## Impact

The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.

- Cloud Provider Access authorization L1 CDK constructor

## Prerequisites

### Resources needed to run the manual QA

- An Atlas project.
- A Cloud Provider Access role registration in that project, created by the `MongoDB::Atlas::CloudProviderAccessSetup` resource or by `atlas cloudProviders accessRoles aws create`.
- An AWS IAM role whose trust policy allows the Atlas AWS account returned by the registration and requires the registration external id.

All of these resources are created as part of `cfn-testing-helper.sh`, except for the contract-test deviation described below.

The test input script creates the AWS IAM roles, so export the Cloud Tag Policy variables before running it: `MONGODB_TAG_OWNER` and `MONGODB_TAG_ENV`. See [TESTING.md](../../../TESTING.md).

## Manual QA

Please follow the steps in [TESTING.md](../../../TESTING.md).

### Success criteria when testing the resource

- Creating the resource authorizes the role, and the Atlas project's Cloud Provider Access page shows the role as authorized.
- Reading the resource returns the authorization date and the feature usages.
- Changing `IamAssumedRoleArn` re-authorizes the same role in place. The role id does not change and the resource is not replaced.
- Deleting the resource removes the whole Cloud Provider Access role registration. The paired `MongoDB::Atlas::CloudProviderAccessSetup` resource is then orphaned.
- Deleting a role that a feature uses fails with `AWS_IAM_ROLE_IN_USE`. Disable the feature first, for example by disabling Encryption at Rest.
- A second create for an existing authorization fails with `AlreadyExists` for either the same or a different ARN.
- Create returns a checkpoint before authorizing. Replaying its callback after successful authorization returns `SUCCESS` without another PATCH.
- `CANNOT_ASSUME_ROLE` returns `IN_PROGRESS` and retries across callbacks. A callback after the five-minute budget fails with `NotStabilized`.
- Updating a registered but unauthorized role returns `NotFound` without a PATCH.
- Read and Delete work with only `ProjectId`, `RoleId`, and a non-default `Profile`.

### Validate the two resources together

Use the [combined stack example](../../../examples/cloud-provider-access-authorization/README.md) in a test
project and follow its KMS and IAM prerequisites. Handler unit tests cannot verify CloudFormation's stack
ordering, IAM propagation, or the Atlas API's live behavior. Before release, check these workflows:

1. Create the stack. Setup must register the role before the IAM trust policy is created, Authorization must
   complete before Encryption at Rest is enabled, and the stack outputs must identify the same registration.
2. Change Authorization's `IamAssumedRoleArn` to a second assumable IAM role with the same Atlas principal and
   external ID, and the required KMS permissions. Verify the registration's role ID and external ID remain
   unchanged, Read returns the new ARN, and the consumer still works. Keep the old IAM role available until
   the update completes so rollback can restore its authorization.
3. Delete the complete stack after removing any cluster or other feature dependencies described in the
   example. Verify Encryption at Rest is disabled before Authorization deletes the registration. Setup's
   subsequent Delete must report `NotFound`, and the stack must still reach `DELETE_COMPLETE`.
4. In a separate disposable stack, remove the consumer and Authorization while keeping Setup. Verify the
   registration is gone and Setup Read returns `NotFound`. Re-adding Authorization with the old role ID must
   fail; recovery requires a new Setup registration and updated IAM trust policy. Do not interpret this as
   a failure of complete-stack deletion.
5. With `DeletionPolicy: Retain` and `UpdateReplacePolicy: Retain` deliberately set on both Atlas resources,
   verify the registration remains when they are removed. If the IAM role is also retained, confirm the
   authorization remains active. Clean up retained test resources explicitly after checking consumers.

Record the stack events and corresponding Atlas state for each live run. These are validation instructions,
not evidence that the live workflows have passed.

### Contract-test deviation

The CloudFormation v1 contract test `contract_delete_create` expects a deleted resource to be creatable again. Atlas has no operation that revokes an authorization while keeping the role registered, so after the delete the re-create fails with `CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND`. Exception approval remains to be verified.

CI runs the suite in isolated groups (`test/contract-testing/run-contract-cases.sh`), each with a fresh registration and matching IAM roles. A single shared registration would be invalidated by the first delete and later fixtures would fail at setup instead of validating their cases. The update group uses a second IAM role, so changing `IamAssumedRoleArn` exercises a real re-authorization.

`contract_delete_create` runs in its own step (`test/contract-testing/check-contract-delete-create.sh`). The step verifies the exact expected signature, a failed re-create with `CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND` after a successful create and delete with no other failure, and stays red until the AWS exception is accepted. Flip `EXCEPTION_ACCEPTED` in the script in the commit that records the acceptance.

The [current AWS contract-test documentation](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/contract-tests.html)
also describes v2 requirements for newly registered types. Its delete/create assertion differs from the v1
test above; assess the applicable publication tests separately. The [local v2 runner supports Java only](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-test.html),
so a successful local Go unit run or private registration does not establish public-registry contract compliance.

Local handler tests use a fake HTTP transport to cover duplicate creates, callback replay, propagation
retries, timeout, missing authorization on update, deletion conflicts, non-AWS registrations, and client
initialization failures after a Create checkpoint. A shared-state test runs Authorization Delete followed by
Setup Delete and checks that Setup returns `NotFound` without a second API delete. These tests do not create
AWS or Atlas resources or execute the CloudFormation stack engine.

## Important Links

- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-authorizegroupcloudprovideraccessrole)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security/set-up-unified-aws-access/)
