# MongoDB::Atlas::ProjectServiceAccount

## Description

Resource for managing [Project Service Accounts](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-createprojectserviceaccount) for a MongoDB Atlas project. Project service accounts provide programmatic access to project-level resources and are used for automation, CI/CD pipelines, and service-to-service authentication within a specific project scope.

**Note:** Deleting this resource only unassigns the Service Account from the project, but doesn't delete it from the organization.

## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](/examples/project-service-account/README.md) for example resource.
