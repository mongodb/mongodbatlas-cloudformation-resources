# MongoDB::Atlas::CloudBackupSnapshotExportJob

Returns, adds, and removes Cloud Backup snapshot export buckets. Also returns and adds Cloud Backup export jobs.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::CloudBackupSnapshotExportJob",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
        "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
        "<a href="#customdataset" title="CustomDataSet">CustomDataSet</a>" : <i>[ <a href="customdata.md">customData</a>, ... ]</i>,
        "<a href="#exportstatus" title="ExportStatus">ExportStatus</a>" : <i><a href="apiexportstatusview.md">ApiExportStatusView</a></i>,
        "<a href="#includecount" title="IncludeCount">IncludeCount</a>" : <i>Boolean</i>,
        "<a href="#snapshotid" title="SnapshotId">SnapshotId</a>" : <i>String</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::CloudBackupSnapshotExportJob
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    <a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
    <a href="#customdataset" title="CustomDataSet">CustomDataSet</a>: <i>
      - <a href="customdata.md">customData</a></i>
    <a href="#exportstatus" title="ExportStatus">ExportStatus</a>: <i><a href="apiexportstatusview.md">ApiExportStatusView</a></i>
    <a href="#includecount" title="IncludeCount">IncludeCount</a>: <i>Boolean</i>
    <a href="#snapshotid" title="SnapshotId">SnapshotId</a>: <i>String</i>
</pre>

## Properties

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClusterName

Human-readable label that identifies the cluster.

_Required_: No

_Type_: String

_Minimum_: <code>1</code>

_Maximum_: <code>64</code>

_Pattern_: <code>^([a-zA-Z0-9]([a-zA-Z0-9-]){0,21}(?<!-)([\w]{0,42}))$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CustomDataSet

Collection of key-value pairs that represent custom data for the metadata file that MongoDB Cloud uploads to the bucket when the export job finishes.

_Required_: No

_Type_: List of <a href="customdata.md">customData</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ExportStatus

_Required_: No

_Type_: <a href="apiexportstatusview.md">ApiExportStatusView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IncludeCount

Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SnapshotId

Unique 24-hexadecimal character string that identifies the snapshot.

_Required_: No

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ExportId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### ExportId

Unique string that identifies the AWS S3 bucket to which you export your snapshots.

#### FinishedAt

Date and time when this export job completed. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.

#### Components

Information on the export job for each replica set in the sharded cluster.

#### Links

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.

#### PageNum

Number of the page that displays the current set of the total objects that the response returns.

#### GroupId

Unique 24-hexadecimal digit string that identifies your project.

#### ExportStatus

Returns the <code>ExportStatus</code> value.

#### DeliveryUrl

One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when `"deliveryType" : "download"`.

#### ItemsPerPage

Number of items that the response returns per page.

#### DeliveryUrl

One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when `"deliveryType" : "download"`.

#### ExportBucketId

Unique 24-hexadecimal character string that identifies the AWS bucket to which MongoDB Cloud exports the Cloud Backup snapshot.

#### Id

Unique 24-hexadecimal character string that identifies the restore job.

#### Prefix

Full path on the cloud provider bucket to the folder where the snapshot is exported.

#### CreatedAt

Date and time when someone created this export job. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.

#### State

State of the export job.

#### ReplicaSetName

Returns the <code>ReplicaSetName</code> value.

#### ExportedCollections

Returns the <code>ExportedCollections</code> value.

#### TotalCollections

Returns the <code>TotalCollections</code> value.

