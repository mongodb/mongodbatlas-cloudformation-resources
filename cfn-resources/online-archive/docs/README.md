# MongoDB::Atlas::OnlineArchive

Returns, adds, edits, or removes an online archive.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::OnlineArchive",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
        "<a href="#collname" title="CollName">CollName</a>" : <i>String</i>,
        "<a href="#collectiontype" title="CollectionType">CollectionType</a>" : <i>String</i>,
        "<a href="#criteria" title="Criteria">Criteria</a>" : <i><a href="criteriaview.md">CriteriaView</a></i>,
        "<a href="#dbname" title="DbName">DbName</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#includecount" title="IncludeCount">IncludeCount</a>" : <i>Boolean</i>,
        "<a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>" : <i>Integer</i>,
        "<a href="#pagenum" title="PageNum">PageNum</a>" : <i>Integer</i>,
        "<a href="#partitionfields" title="PartitionFields">PartitionFields</a>" : <i>[ <a href="partitionfieldview.md">PartitionFieldView</a>, ... ]</i>,
        "<a href="#schedule" title="Schedule">Schedule</a>" : <i><a href="scheduleview.md">ScheduleView</a></i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::OnlineArchive
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
    <a href="#collname" title="CollName">CollName</a>: <i>String</i>
    <a href="#collectiontype" title="CollectionType">CollectionType</a>: <i>String</i>
    <a href="#criteria" title="Criteria">Criteria</a>: <i><a href="criteriaview.md">CriteriaView</a></i>
    <a href="#dbname" title="DbName">DbName</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#includecount" title="IncludeCount">IncludeCount</a>: <i>Boolean</i>
    <a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>: <i>Integer</i>
    <a href="#pagenum" title="PageNum">PageNum</a>: <i>Integer</i>
    <a href="#partitionfields" title="PartitionFields">PartitionFields</a>: <i>
      - <a href="partitionfieldview.md">PartitionFieldView</a></i>
    <a href="#schedule" title="Schedule">Schedule</a>: <i><a href="scheduleview.md">ScheduleView</a></i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ClusterName

Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>1</code>

_Maximum Length_: <code>64</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### CollName

Human-readable label that identifies the collection for which you created the online archive.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CollectionType

Classification of MongoDB database collection that you want to return.

If you set this parameter to `TIMESERIES`, set `"criteria.type" : "date"` and `"criteria.dateFormat" : "ISODATE"`.

_Required_: No

_Type_: String

_Allowed Values_: <code>STANDARD</code> | <code>TIMESERIES</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Criteria

_Required_: Yes

_Type_: <a href="criteriaview.md">CriteriaView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DbName

Human-readable label of the database that contains the collection that contains the online archive.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### IncludeCount

Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ItemsPerPage

Number of items that the response returns per page.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PageNum

Number of the page that displays the current set of the total objects that the response returns.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PartitionFields

List that contains document parameters to use to logically divide data within a collection. Partitions provide a coarse level of filtering of the underlying collection data. To divide your data, specify up to two parameters that you frequently query. Any queries that don't use these parameters result in a full collection scan of all archived documents. This takes more time and increase your costs.

_Required_: No

_Type_: List of <a href="partitionfieldview.md">PartitionFieldView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Schedule

_Required_: No

_Type_: <a href="scheduleview.md">ScheduleView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### State

Phase of the process to create this online archive when you made this request.

| State       | Indication |
|-------------|------------|
| `PENDING`   | MongoDB Cloud has queued documents for archive. Archiving hasn't started. |
| `ARCHIVING` | MongoDB Cloud started archiving documents that meet the archival criteria. |
| `IDLE`      | MongoDB Cloud waits to start the next archival job. |
| `PAUSING`   | Someone chose to stop archiving. MongoDB Cloud finishes the running archival job then changes the state to `PAUSED` when that job completes. |
| `PAUSED`    | MongoDB Cloud has stopped archiving. Archived documents can be queried. The specified archiving operation on the active cluster cannot archive additional documents. You can resume archiving for paused archives at any time. |
| `ORPHANED`  | Someone has deleted the collection associated with an active or paused archive. MongoDB Cloud doesn't delete the archived data. You must manually delete the online archives associated with the deleted collection. |
| `DELETED`   | Someone has deleted the archive was deleted. When someone deletes an online archive, MongoDB Cloud removes all associated archived documents from the cloud object storage. |

#### TotalCount

Number of documents returned in this response.

#### ArchiveId

Unique 24-hexadecimal digit string that identifies the online archive to delete.

