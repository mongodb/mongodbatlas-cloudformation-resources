# MongoDB::Atlas::CloudBackupSchedule ApiAtlasDiskBackupCopySettingView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#cloudprovider" title="CloudProvider">CloudProvider</a>" : <i>String</i>,
    "<a href="#regionname" title="RegionName">RegionName</a>" : <i>String</i>,
    "<a href="#replicationspecid" title="ReplicationSpecId">ReplicationSpecId</a>" : <i>String</i>,
    "<a href="#shouldcopyoplogs" title="ShouldCopyOplogs">ShouldCopyOplogs</a>" : <i>Boolean</i>,
    "<a href="#frequencies" title="Frequencies">Frequencies</a>" : <i>[ String, ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#cloudprovider" title="CloudProvider">CloudProvider</a>: <i>String</i>
<a href="#regionname" title="RegionName">RegionName</a>: <i>String</i>
<a href="#replicationspecid" title="ReplicationSpecId">ReplicationSpecId</a>: <i>String</i>
<a href="#shouldcopyoplogs" title="ShouldCopyOplogs">ShouldCopyOplogs</a>: <i>Boolean</i>
<a href="#frequencies" title="Frequencies">Frequencies</a>: <i>
      - String</i>
</pre>

## Properties

#### CloudProvider

A label that identifies the cloud provider that stores the snapshot copy.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RegionName

Target region to copy snapshots belonging to replicationSpecId to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ReplicationSpecId

Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ShouldCopyOplogs

Flag that indicates whether to copy the oplogs to the target region. 

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Frequencies

List that describes which types of snapshots to copy.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

