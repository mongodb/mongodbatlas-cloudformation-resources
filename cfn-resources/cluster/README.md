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


## Operation timeout

CloudFormation limits how long a single resource operation can run. This resource does not set that limit, so the CloudFormation default of 120 minutes applies and you cannot override it. Long Atlas changes, such as a shard removal or a storage scale-down, can outlast that limit. When they do, CloudFormation fails the operation and rolls back the stack, which can leave the cluster mid-change.

Set `SkipWaitForIdle` to `true` to report success as soon as Atlas accepts the change, instead of polling until the cluster reaches `IDLE`. Leave `SkipWaitForIdle` unset to keep the default behavior of waiting for `IDLE`.

### SkipWaitForIdle limitations

- The stack reports success before the change converges. An Atlas failure after that point is not reported to CloudFormation, and anything triggered by the stack completing can reach a cluster that is not ready yet. To confirm the change landed, check the cluster in Atlas. Drift detection only catches a failed change when you run it, and only for properties set in your template.
- Read-only properties reflect the cluster before the change converges. `Id` and `CreatedDate` are final, but `StateName` reports `CREATING` or `UPDATING`, and on update `ConnectionStrings` and `MongoDBVersion` can be stale if the change alters them.
- On create, Atlas has not generated connection strings yet, so a template that reads `ConnectionStrings` with `Fn::GetAtt` fails and the stack rolls back.
- `SkipWaitForIdle` cannot be combined with a change to `Paused`. Apply the pause in a separate update.
- `SkipWaitForIdle` has no effect on flex clusters. Flex clusters support no long-running changes, such as shard removal or a storage scale-down, so the CloudFormation operation timeout is not a risk for them.


## Cloudformation Examples

For examples, see the [CFN Template](/examples/cluster/cluster.json) example.