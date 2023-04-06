# MongoDB::Atlas::CloudBackupSnapshotExportJob ExportStatus

State of the export job for the collections on the replica set only.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#exportedcollections" title="ExportedCollections">ExportedCollections</a>" : <i>Integer</i>,
    "<a href="#totalcollections" title="TotalCollections">TotalCollections</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#exportedcollections" title="ExportedCollections">ExportedCollections</a>: <i>Integer</i>
<a href="#totalcollections" title="TotalCollections">TotalCollections</a>: <i>Integer</i>
</pre>

## Properties

#### ExportedCollections

Number of collections on the replica set that MongoDB Cloud exported.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TotalCollections

Total number of collections on the replica set to export.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

