# MongoDB::Atlas::MongoDbEmployeeAccessGrant

## Description

Resource for managing [MongoDB Employee Access Grants](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-grantgroupclustermongodbemployeeaccess) for Atlas clusters. This resource grants temporary access to MongoDB employees for a specific cluster, allowing MongoDB support engineers to access cluster infrastructure, database logs, or app services sync data for troubleshooting purposes.

## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](/examples/mongodb-employee-access-grant/README.md) for example resource.
