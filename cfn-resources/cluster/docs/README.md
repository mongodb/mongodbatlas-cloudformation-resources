# MongoDB::Atlas::Cluster

The cluster resource provides access to your cluster configurations. The resource lets you create, edit and delete clusters. The resource requires your Project ID.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::Cluster",
    "Properties" : {
        "<a href="#advancedsettings" title="AdvancedSettings">AdvancedSettings</a>" : <i><a href="processargs.md">processArgs</a></i>,
        "<a href="#backupenabled" title="BackupEnabled">BackupEnabled</a>" : <i>Boolean</i>,
        "<a href="#biconnector" title="BiConnector">BiConnector</a>" : <i><a href="biconnector.md">BiConnector</a></i>,
        "<a href="#clustertype" title="ClusterType">ClusterType</a>" : <i>String</i>,
        "<a href="#connectionstrings" title="ConnectionStrings">ConnectionStrings</a>" : <i><a href="connectionstrings.md">connectionStrings</a></i>,
        "<a href="#disksizegb" title="DiskSizeGB">DiskSizeGB</a>" : <i>Double</i>,
        "<a href="#encryptionatrestprovider" title="EncryptionAtRestProvider">EncryptionAtRestProvider</a>" : <i>String</i>,
        "<a href="#globalclusterselfmanagedsharding" title="GlobalClusterSelfManagedSharding">GlobalClusterSelfManagedSharding</a>" : <i>Boolean</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#labels" title="Labels">Labels</a>" : <i>[ [ <a href="labels.md">Labels</a>, ... ], ... ]</i>,
        "<a href="#mongodbmajorversion" title="MongoDBMajorVersion">MongoDBMajorVersion</a>" : <i>String</i>,
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#paused" title="Paused">Paused</a>" : <i>Boolean</i>,
        "<a href="#pitenabled" title="PitEnabled">PitEnabled</a>" : <i>Boolean</i>,
        "<a href="#replicationspecs" title="ReplicationSpecs">ReplicationSpecs</a>" : <i>[ <a href="advancedreplicationspec.md">advancedReplicationSpec</a>, ... ]</i>,
        "<a href="#rootcerttype" title="RootCertType">RootCertType</a>" : <i>String</i>,
        "<a href="#versionreleasesystem" title="VersionReleaseSystem">VersionReleaseSystem</a>" : <i>String</i>,
        "<a href="#terminationprotectionenabled" title="TerminationProtectionEnabled">TerminationProtectionEnabled</a>" : <i>Boolean</i>,
        "<a href="#tags" title="Tags">Tags</a>" : <i>[ <a href="tag.md">tag</a>, ... ]</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::Cluster
Properties:
    <a href="#advancedsettings" title="AdvancedSettings">AdvancedSettings</a>: <i><a href="processargs.md">processArgs</a></i>
    <a href="#backupenabled" title="BackupEnabled">BackupEnabled</a>: <i>Boolean</i>
    <a href="#biconnector" title="BiConnector">BiConnector</a>: <i><a href="biconnector.md">BiConnector</a></i>
    <a href="#clustertype" title="ClusterType">ClusterType</a>: <i>String</i>
    <a href="#connectionstrings" title="ConnectionStrings">ConnectionStrings</a>: <i><a href="connectionstrings.md">connectionStrings</a></i>
    <a href="#disksizegb" title="DiskSizeGB">DiskSizeGB</a>: <i>Double</i>
    <a href="#encryptionatrestprovider" title="EncryptionAtRestProvider">EncryptionAtRestProvider</a>: <i>String</i>
    <a href="#globalclusterselfmanagedsharding" title="GlobalClusterSelfManagedSharding">GlobalClusterSelfManagedSharding</a>: <i>Boolean</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#labels" title="Labels">Labels</a>: <i>
      - 
      - <a href="labels.md">Labels</a></i>
    <a href="#mongodbmajorversion" title="MongoDBMajorVersion">MongoDBMajorVersion</a>: <i>String</i>
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#paused" title="Paused">Paused</a>: <i>Boolean</i>
    <a href="#pitenabled" title="PitEnabled">PitEnabled</a>: <i>Boolean</i>
    <a href="#replicationspecs" title="ReplicationSpecs">ReplicationSpecs</a>: <i>
      - <a href="advancedreplicationspec.md">advancedReplicationSpec</a></i>
    <a href="#rootcerttype" title="RootCertType">RootCertType</a>: <i>String</i>
    <a href="#versionreleasesystem" title="VersionReleaseSystem">VersionReleaseSystem</a>: <i>String</i>
    <a href="#terminationprotectionenabled" title="TerminationProtectionEnabled">TerminationProtectionEnabled</a>: <i>Boolean</i>
    <a href="#tags" title="Tags">Tags</a>: <i>
      - <a href="tag.md">tag</a></i>
</pre>

## Properties

#### AdvancedSettings

Advanced configuration details to add for one cluster in the specified project.

_Required_: No

_Type_: <a href="processargs.md">processArgs</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### BackupEnabled

Flag that indicates whether the cluster can perform backups. If set to true, the cluster can perform backups. You must set this value to true for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to false, the cluster doesn't use backups.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### BiConnector

Settings needed to configure the MongoDB Connector for Business Intelligence for this cluster.

_Required_: No

_Type_: <a href="biconnector.md">BiConnector</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClusterType

Configuration of nodes that comprise the cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ConnectionStrings

Collection of Uniform Resource Locators that point to the MongoDB database.

_Required_: No

_Type_: <a href="connectionstrings.md">connectionStrings</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DiskSizeGB

Storage capacity that the host's root volume possesses expressed in gigabytes. Increase this number to add capacity. MongoDB Cloud requires this parameter if you set replicationSpecs. If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value. Storage charge calculations depend on whether you choose the default value or a custom value. The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EncryptionAtRestProvider

Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster. To enable customer key management for encryption at rest, the cluster replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize setting must be M10 or higher and "backupEnabled" : false or omitted entirely.

_Required_: No

_Type_: String

_Allowed Values_: <code>AWS</code> | <code>GCP</code> | <code>AZURE</code> | <code>NONE</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### GlobalClusterSelfManagedSharding

(Optional) Flag that indicates if cluster uses Atlas-Managed Sharding (false, default) or Self-Managed Sharding (true). It can only be enabled for Global Clusters (`GEOSHARDED`). It cannot be changed once the cluster is created. Use this mode if you're an advanced user and the default configuration is too restrictive for your workload. If you select this option, you must manually configure the sharding strategy, more info [here](https://www.mongodb.com/docs/atlas/tutorial/create-global-cluster/#select-your-sharding-configuration).

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique identifier of the project the cluster belongs to.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Labels

Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn't display your labels.

_Required_: No

_Type_: List of List of <a href="labels.md">Labels</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MongoDBMajorVersion

Major MongoDB version of the cluster. MongoDB Cloud deploys the cluster with the latest stable release of the specified version.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

Human-readable label that identifies the advanced cluster.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Paused

Flag that indicates whether the cluster is paused or not. When unpausing a cluster the update operaton will not perform any other changes to the cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PitEnabled

Flag that indicates whether the cluster uses continuous cloud backups.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ReplicationSpecs

List of settings that configure your cluster regions. For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global replica sets and sharded clusters, this array has one object representing where your clusters nodes deploy.

_Required_: No

_Type_: List of <a href="advancedreplicationspec.md">advancedReplicationSpec</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RootCertType

Root Certificate Authority that MongoDB Cloud cluster uses. MongoDB Cloud supports Internet Security Research Group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### VersionReleaseSystem

Method by which the cluster maintains the MongoDB versions. If value is CONTINUOUS, you must not specify mongoDBMajorVersion

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TerminationProtectionEnabled

Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Tags

List of settings that configure your cluster regions. For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global replica sets and sharded clusters, this array has one object representing where your clusters nodes deploy.

_Required_: No

_Type_: List of <a href="tag.md">tag</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Standard

Returns the <code>Standard</code> value.

#### StandardSrv

Returns the <code>StandardSrv</code> value.

#### Private

Returns the <code>Private</code> value.

#### PrivateSrv

Returns the <code>PrivateSrv</code> value.

#### PrivateEndpoints

Returns the <code>PrivateEndpoints</code> value.

#### PrivateEndpointsSrv

Returns the <code>PrivateEndpointsSrv</code> value.

#### SRVShardOptimizedConnectionString

Returns the <code>SRVShardOptimizedConnectionString</code> value.

#### StateName

Current state of the cluster.

#### MongoDBVersion

Version of MongoDB that the cluster runs.

#### CreatedDate

Date and time when MongoDB Cloud created this cluster. This parameter expresses its value in ISO 8601 format in UTC.

#### Id

Unique identifier of the cluster.

