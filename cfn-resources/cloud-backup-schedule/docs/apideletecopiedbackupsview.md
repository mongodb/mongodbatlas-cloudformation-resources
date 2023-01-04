# MongoDB::Atlas::CloudBackupSchedule ApiDeleteCopiedBackupsView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#cloudprovider" title="CloudProvider">CloudProvider</a>" : <i>String</i>,
    "<a href="#regionname" title="RegionName">RegionName</a>" : <i>String</i>,
    "<a href="#replicationspecid" title="ReplicationSpecId">ReplicationSpecId</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#cloudprovider" title="CloudProvider">CloudProvider</a>: <i>String</i>
<a href="#regionname" title="RegionName">RegionName</a>: <i>String</i>
<a href="#replicationspecid" title="ReplicationSpecId">ReplicationSpecId</a>: <i>String</i>
</pre>

## Properties

#### CloudProvider

A label that identifies the cloud provider for the deleted copy setting whose backup copies you want to delete

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RegionName

Target region for the deleted copy setting whose backup copies you want to delete.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ReplicationSpecId

Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

