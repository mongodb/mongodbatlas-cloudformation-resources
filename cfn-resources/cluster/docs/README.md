# MongoDB::Atlas::Cluster

The cluster resource provides access to your cluster configurations. The resource lets you create, edit and delete clusters. The resource requires your Project ID.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::Cluster",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
        "<a href="#autoscaling" title="AutoScaling">AutoScaling</a>" : <i><a href="autoscaling.md">autoScaling</a></i>,
        "<a href="#backupenabled" title="BackupEnabled">BackupEnabled</a>" : <i>Boolean</i>,
        "<a href="#biconnector" title="BiConnector">BiConnector</a>" : <i><a href="biconnector.md">BiConnector</a></i>,
        "<a href="#clustertype" title="ClusterType">ClusterType</a>" : <i>String</i>,
        "<a href="#connectionstrings" title="ConnectionStrings">ConnectionStrings</a>" : <i><a href="connectionstrings.md">connectionStrings</a></i>,
        "<a href="#disksizegb" title="DiskSizeGB">DiskSizeGB</a>" : <i>Double</i>,
        "<a href="#encryptionatrestprovider" title="EncryptionAtRestProvider">EncryptionAtRestProvider</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#labels" title="Labels">Labels</a>" : <i>[ [ <a href="labels.md">Labels</a>, ... ], ... ]</i>,
        "<a href="#mongodbmajorversion" title="MongoDBMajorVersion">MongoDBMajorVersion</a>" : <i>String</i>,
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#numshards" title="NumShards">NumShards</a>" : <i>Integer</i>,
        "<a href="#pitenabled" title="PitEnabled">PitEnabled</a>" : <i>Boolean</i>,
        "<a href="#providerbackupenabled" title="ProviderBackupEnabled">ProviderBackupEnabled</a>" : <i>Boolean</i>,
        "<a href="#providersettings" title="ProviderSettings">ProviderSettings</a>" : <i><a href="providersettings.md">ProviderSettings</a></i>,
        "<a href="#replicationfactor" title="ReplicationFactor">ReplicationFactor</a>" : <i>Integer</i>,
        "<a href="#replicationspecs" title="ReplicationSpecs">ReplicationSpecs</a>" : <i>[ <a href="replicationspec.md">ReplicationSpec</a>, ... ]</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::Cluster
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    <a href="#autoscaling" title="AutoScaling">AutoScaling</a>: <i><a href="autoscaling.md">autoScaling</a></i>
    <a href="#backupenabled" title="BackupEnabled">BackupEnabled</a>: <i>Boolean</i>
    <a href="#biconnector" title="BiConnector">BiConnector</a>: <i><a href="biconnector.md">BiConnector</a></i>
    <a href="#clustertype" title="ClusterType">ClusterType</a>: <i>String</i>
    <a href="#connectionstrings" title="ConnectionStrings">ConnectionStrings</a>: <i><a href="connectionstrings.md">connectionStrings</a></i>
    <a href="#disksizegb" title="DiskSizeGB">DiskSizeGB</a>: <i>Double</i>
    <a href="#encryptionatrestprovider" title="EncryptionAtRestProvider">EncryptionAtRestProvider</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#labels" title="Labels">Labels</a>: <i>
      - 
      - <a href="labels.md">Labels</a></i>
    <a href="#mongodbmajorversion" title="MongoDBMajorVersion">MongoDBMajorVersion</a>: <i>String</i>
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#numshards" title="NumShards">NumShards</a>: <i>Integer</i>
    <a href="#pitenabled" title="PitEnabled">PitEnabled</a>: <i>Boolean</i>
    <a href="#providerbackupenabled" title="ProviderBackupEnabled">ProviderBackupEnabled</a>: <i>Boolean</i>
    <a href="#providersettings" title="ProviderSettings">ProviderSettings</a>: <i><a href="providersettings.md">ProviderSettings</a></i>
    <a href="#replicationfactor" title="ReplicationFactor">ReplicationFactor</a>: <i>Integer</i>
    <a href="#replicationspecs" title="ReplicationSpecs">ReplicationSpecs</a>: <i>
      - <a href="replicationspec.md">ReplicationSpec</a></i>
</pre>

## Properties

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AutoScaling

_Required_: No

_Type_: <a href="autoscaling.md">autoScaling</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### BackupEnabled

Applicable only for M10+ clusters. Set to true to enable Atlas continuous backups for the cluster. Set to false to disable continuous backups for the cluster. Atlas deletes any stored snapshots. See the continuous backup Snapshot Schedule for more information. You cannot enable continuous backups if you have an existing cluster in the project with Cloud Provider Snapshots enabled. The default value is false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### BiConnector

_Required_: No

_Type_: <a href="biconnector.md">BiConnector</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClusterType

Type of the cluster that you want to create.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ConnectionStrings

_Required_: No

_Type_: <a href="connectionstrings.md">connectionStrings</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DiskSizeGB

Capacity, in gigabytes, of the host’s root volume. Increase this number to add capacity, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive integer.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EncryptionAtRestProvider

Set the Encryption at Rest parameter.

_Required_: No

_Type_: String

_Allowed Values_: <code>AWS</code> | <code>GCP</code> | <code>AZURE</code> | <code>NONE</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique identifier of the project the cluster belongs to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Labels

Array containing key-value pairs that tag and categorize the cluster.

_Required_: No

_Type_: List of List of <a href="labels.md">Labels</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MongoDBMajorVersion

Major version of the cluster to deploy.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

Name of the cluster. Once the cluster is created, its name cannot be changed.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NumShards

Positive integer that specifies the number of shards to deploy for a sharded cluster.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PitEnabled

Flag that indicates if the cluster uses Point-in-Time backups. If set to true, providerBackupEnabled must also be set to true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProviderBackupEnabled

Applicable only for M10+ clusters. Set to true to enable Atlas Cloud Provider Snapshots backups for the cluster. Set to false to disable Cloud Provider Snapshots backups for the cluster. You cannot enable Cloud Provider Snapshots if you have an existing cluster in the project with continuous backups enabled. Note that you must set this value to true for NVMe clusters. The default value is false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProviderSettings

_Required_: No

_Type_: <a href="providersettings.md">ProviderSettings</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ReplicationFactor

ReplicationFactor is deprecated. Use replicationSpecs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ReplicationSpecs

Configuration for cluster regions.

_Required_: No

_Type_: List of <a href="replicationspec.md">ReplicationSpec</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Id.

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

#### StateName

Current state of the cluster.

#### SrvAddress

Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS. The mongoURI parameter lists additional options.

#### Paused

Flag that indicates whether the cluster is paused or not.

#### MongoDBVersion

Version of MongoDB the cluster runs, in <major version>.<minor version> format.

#### MongoURI

Base connection string for the cluster.

#### MongoURIUpdated

Timestamp in ISO 8601 date and time format in UTC when the connection string was last updated. The connection string changes if you update any of the other values.

#### MongoURIWithOptions

connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster.

#### Id

Unique identifier of the cluster.

