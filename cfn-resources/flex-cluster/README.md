# MongoDB::Atlas::FlexCluster

## Description

The flex cluster resource provides access to your [Flex Clusters](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-flex-clusters) configurations and enables you to create, edit, and delete flex clusters. For more information, see [The MongoDB Atlas Flex Tier](https://www.mongodb.com/company/blog/product-release-announcements/dynamic-workloads-predictable-costs-mongodb-atlas-flex-tier).

*Note:* Upgrades to or from flex clusters are currently unavailable. We expect to support upgrades to or from flex clusters in the forseeable future.

*Important:* Use the `MongoDB::Atlas::Cluster` resource instead of the `MongoDB::Atlas::FlexCluster` resource to create and manage flex clusters. `MongoDB::Atlas::Cluster` supports flex clusters and future upgrades will only be available through this resource. For more information, see [`MongoDB::Atlas::Cluster` README](../cluster/README.md).

## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

See the examples [CFN Template](/examples/flex-cluster/README.md) for example resource.
