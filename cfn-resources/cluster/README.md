# MongoDB::Atlas::Cluster CFN resource

## Description
Provides a resource for managing [Clusters](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-clusters) in AWS Cloud Formation. The resource 
lets you create, edit, and delete clusters. The resource requires your Project 
ID to perform these actions.

*Note:* Upgrades to or from flex clusters are currently unavailable. We expect to support upgrades to or from flex clusters in the foreseeable future.

*Important:* Use the `MongoDB::Atlas::Cluster` resource instead of the `MongoDB::Atlas::FlexCluster` resource to create and manage flex clusters. `MongoDB::Atlas::Cluster` supports flex clusters and future upgrades will only be available through this resource.



## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).


## Attributes and Parameters
For futher information, see the [resource docs](docs/README.md) section.


## Cloudformation Examples

For examples, see the [CFN Template](/examples/cluster/cluster.json) example.