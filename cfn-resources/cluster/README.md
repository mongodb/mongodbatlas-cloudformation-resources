# MongoDB::Atlas::Cluster CFN resource

## Description
Provides a resource for managing [Clusters](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-clusters) in AWS Cloud Formation. The resource 
lets you create, edit, and delete clusters. The resource requires your Project 
ID to perform these actions.


## Requirements

To securely give CloudFormation access to your Atlas credentials, you must
set up an [AWS Profile](/README.md#mongodb-atlas-api-keys-credential-management).


## Attributes and Parameters

* `AdvancedSettings` -  processArgs
Advanced configuration details to add for one cluster in the specified project.

* `BackupEnabled` -  Boolean
Flag that indicates whether the cluster can perform backups. If set to true, the cluster can perform backups. You must set this value to true for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to false, the cluster doesn't use backups.


* `BiConnector` - BiConnector
Settings needed to configure the MongoDB Connector for Business Intelligence for this cluster.

* `ClusterType` - String
Configuration of nodes that comprise the cluster.

* `ConnectionStrings` - connectionStrings
Collection of Uniform Resource Locators that point to the MongoDB database.


* `DiskSizeGB` - Double
 Storage capacity that the host's root volume possesses expressed in gigabytes. Increase this number to add capacity. MongoDB Cloud requires this parameter if you set replicationSpecs. If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value. Storage charge calculations depend on whether you choose the default value or a custom value. The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier.

* `EncryptionAtRestProvider` - String
Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster. To enable customer key management for encryption at rest, the cluster replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize setting must be M10 or higher and "backupEnabled" : false or omitted entirely. - Allowed Values: `AWS` | `GCP` | `AZURE` | `NONE`

* `GlobalClusterSelfManagedSharding` - Optional - Boolean - Requires replacement
Flag that indicates if cluster uses Atlas-Managed Sharding (false, default) or Self-Managed Sharding (true). It can only be enabled for Global Clusters (GEOSHARDED). It cannot be changed once the cluster is created. Use this mode if you're an advanced user and the default configuration is too restrictive for your workload. If you select this option, you must manually configure the sharding strategy, more info here.

* `Profile` - String - Requires replacement
Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used


* `ProjectId` - String - Requires replacement
Unique identifier of the project the cluster belongs to.


* `Labels` - Array
Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn't display your labels.


* `MongoDBMajorVersion` - String 
Major MongoDB version of the cluster. MongoDB Cloud deploys the cluster with the latest stable release of the specified version.


* `Name` - String - Requires replacement
Human-readable label that identifies the advanced cluster.

* `Paused` - Boolean
Flag that indicates whether the cluster is paused or not.


* `PitEnabled` - Boolean
Flag that indicates whether the cluster uses continuous cloud backups.

* `ReplicationSpecs` - Array
List of settings that configure your cluster regions. For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global replica sets and sharded clusters, this array has one object representing where your clusters nodes deploy.

* `RootCertType` - String 
Root Certificate Authority that MongoDB Cloud cluster uses. MongoDB Cloud supports Internet Security Research Group.

* `VersionReleaseSystem` - String 
Method by which the cluster maintains the MongoDB versions. If value is CONTINUOUS, you must not specify mongoDBMajorVersion

* `TerminationProtectionEnabled` - Boolean 
Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.

* `Tags` - Array
List of settings that configure your cluster regions. For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global replica sets and sharded clusters, this array has one object representing where your clusters nodes deploy.


## Cloudformation Examples

See the examples [CFN Template](/examples/cluster/README.md) for example resource.