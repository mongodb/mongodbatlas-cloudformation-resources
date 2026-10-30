# MongoDB::Atlas::FederatedDatabaseInstance

Resource for creating and managing [Federated Database endpoint](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-data-federation).

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](/examples/federated-database-instance/federatedDatabaseInstance.json) for example resource.

## Cloud Provider Access

For AWS access, `CloudProviderConfig.RoleId` is the Atlas registration ID returned by
`MongoDB::Atlas::CloudProviderAccessSetup`, not an AWS IAM role ARN. Authorize the IAM role with
`MongoDB::Atlas::CloudProviderAccessAuthorization` first. The existing JSON example expects that authorization to
have been completed outside the template.

When composing the resources in one stack, add an explicit dependency on Authorization. This fragment assumes
the Setup, IAM role, and Authorization resources from the [shared access guide](/examples/cloud-provider-access-authorization/README.md)
and parameters for the existing project, credential profile, instance name, and S3 bucket:

```yaml
FederatedDatabaseInstance:
  Type: MongoDB::Atlas::FederatedDatabaseInstance
  DependsOn: CloudProviderAccessAuthorization
  Properties:
    ProjectId: !Ref ProjectId
    Profile: !Ref Profile
    TenantName: !Ref TenantName
    CloudProviderConfig:
      RoleId: !GetAtt CloudProviderAccessSetup.RoleId
      TestS3Bucket: !Ref TestS3Bucket
    SkipRoleValidation: false
```

This shows the access configuration, not a complete S3 data-store mapping. Supply the IAM permissions required
for the buckets and operations used by the instance; the KMS policy in the Encryption at Rest example does not
grant S3 access. Keep `TestS3Bucket` with `RoleId`: the current handler sends the AWS cloud configuration only
when both are supplied.

The current `Storage.Stores` schema does not expose the S3 bucket, prefix, delimiter, or region fields present
in the Terraform resource. A complete Terraform S3 storage example therefore cannot be copied into this
CloudFormation type as-is. See the [resource schema](./mongodb-atlas-federateddatabaseinstance.json) before adding
storage mappings.

Remove or reconfigure every instance using the registration before deleting Authorization or Setup. With the
dependency above, CloudFormation deletes this instance before Authorization during full-stack teardown. Deleting
the access registration does not delete the S3 bucket or its objects. Shared registrations also require cleanup
of their other consumers; see the [shared lifecycle guidance](../cloud-provider-access-authorization/README.md#delete-and-recovery).
