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

The test input script creates the AWS IAM role, so export the Cloud Tag Policy variables before running it: `MONGODB_TAG_OWNER` and `MONGODB_TAG_ENV`. See [TESTING.md](../../../TESTING.md).

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

### Contract-test deviation

The CloudFormation v1 contract test `contract_delete_create` expects a deleted resource to be creatable again. Atlas has no operation that revokes an authorization while keeping the role registered, so after the delete the re-create fails with `CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND`. Exception approval remains to be verified.

The current input generator supplies one registration for the entire suite. The first delete invalidates
that input and causes later fixture setup errors; the [PR contract run](https://github.com/mongodb/mongodbatlas-cloudformation-resources/actions/runs/39738197633/job/119688830903)
recorded eight such errors. Provision a fresh registration and its matching IAM trust policy for
independent tests before interpreting their results.
The create/update input files also use the same ARN, so a passing no-op update does not validate changing
to a second assumable IAM role. That change needs a separate live test.

Local handler tests use a fake HTTP transport to cover duplicate creates, callback replay, propagation
retries, timeout, missing authorization on update, and deletion conflicts. They do not create AWS or Atlas resources.

## Important Links

- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-authorizegroupcloudprovideraccessrole)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security/set-up-unified-aws-access/)
