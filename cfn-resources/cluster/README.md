# MongoDB::Atlas::Cluster CFN resource

## Description
Provides a resource for managing [Clusters](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-clusters) in AWS Cloud Formation. The resource 
lets you create, edit, and delete clusters. The resource requires your Project 
ID to perform these actions.

*Important:* `MongoDB::Atlas::Cluster` supports Flex Clusters, and all future
updates for creating and managing Flex Clusters will be exclusively available 
through this resource.

*Important:* We recommend using `MongoDB::Atlas::Cluster` to create and manage
Flex clusters instead of the `MongoDB::Atlas::FlexCluster` as future upgrades 
will only be available through this resource.

*Important:* Cluster upgrades to and from Flex Clusters are currently 
unavailable but are planned for future development.



## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).


## Attributes and Parameters
For futher information, see the [resource docs](docs/README.md) section.


## Cloudformation Examples

For examples, see the [CFN Template](/examples/cluster/cluster.json) example.