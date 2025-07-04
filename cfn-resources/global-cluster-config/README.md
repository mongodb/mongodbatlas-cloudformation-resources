# Mongodb::Atlas::GlobalClusterConfig

## Description
Resource for managing [Global Cluster](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-global-clusters) managed namespaces and custom zone mappings.
This resource can only be used with Atlas-managed clusters, see doc for `GlobalClusterSelfManagedSharding` attribute in `Mongodb::Atlas::Cluster` resource for more info.

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes & Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](test/global-cluster-config.sample-template.yaml) for example resource.
