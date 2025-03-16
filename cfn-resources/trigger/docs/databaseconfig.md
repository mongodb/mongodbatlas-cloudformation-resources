# MongoDB::Atlas::Trigger DatabaseConfig

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#serviceid" title="ServiceId">ServiceId</a>" : <i>String</i>,
    "<a href="#database" title="Database">Database</a>" : <i>String</i>,
    "<a href="#collection" title="Collection">Collection</a>" : <i>String</i>,
    "<a href="#operationtypes" title="OperationTypes">OperationTypes</a>" : <i>[ String, ... ]</i>,
    "<a href="#match" title="Match">Match</a>" : <i>String</i>,
    "<a href="#project" title="Project">Project</a>" : <i>String</i>,
    "<a href="#fulldocument" title="FullDocument">FullDocument</a>" : <i>Boolean</i>,
    "<a href="#fulldocumentbeforechange" title="FullDocumentBeforeChange">FullDocumentBeforeChange</a>" : <i>Boolean</i>,
    "<a href="#skipcatchupevents" title="SkipCatchupEvents">SkipCatchupEvents</a>" : <i>Boolean</i>,
    "<a href="#tolerateresumeerrors" title="TolerateResumeErrors">TolerateResumeErrors</a>" : <i>Boolean</i>,
    "<a href="#maximumthroughput" title="MaximumThroughput">MaximumThroughput</a>" : <i>Boolean</i>,
    "<a href="#unordered" title="Unordered">Unordered</a>" : <i>Boolean</i>
}
</pre>

### YAML

<pre>
<a href="#serviceid" title="ServiceId">ServiceId</a>: <i>String</i>
<a href="#database" title="Database">Database</a>: <i>String</i>
<a href="#collection" title="Collection">Collection</a>: <i>String</i>
<a href="#operationtypes" title="OperationTypes">OperationTypes</a>: <i>
      - String</i>
<a href="#match" title="Match">Match</a>: <i>String</i>
<a href="#project" title="Project">Project</a>: <i>String</i>
<a href="#fulldocument" title="FullDocument">FullDocument</a>: <i>Boolean</i>
<a href="#fulldocumentbeforechange" title="FullDocumentBeforeChange">FullDocumentBeforeChange</a>: <i>Boolean</i>
<a href="#skipcatchupevents" title="SkipCatchupEvents">SkipCatchupEvents</a>: <i>Boolean</i>
<a href="#tolerateresumeerrors" title="TolerateResumeErrors">TolerateResumeErrors</a>: <i>Boolean</i>
<a href="#maximumthroughput" title="MaximumThroughput">MaximumThroughput</a>: <i>Boolean</i>
<a href="#unordered" title="Unordered">Unordered</a>: <i>Boolean</i>
</pre>

## Properties

#### ServiceId

The _id value of a linked MongoDB data source.

See [Get a Data Source](#operation/adminGetService).


_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Database

The name of a database in the linked data source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Collection

The name of a collection in the specified database. The
trigger listens to events from this collection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OperationTypes

The type(s) of MongoDB change event that the trigger listens for.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Match

stringify version of a [$match](https://www.mongodb.com/docs/manual/reference/operator/aggregation/match) expression filters change events. The trigger will only fire if the expression evaluates to true for a given change event.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Project

stringify version of a [$project](https://www.mongodb.com/docs/manual/reference/operator/aggregation/project/) expressions to limit the data included in each event.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FullDocument

If `true`, indicates that `UPDATE` change events should
include the most current
[majority-committed](https://www.mongodb.com/docs/manual/reference/read-concern-majority/)
version of the modified document in the `fullDocument`
field.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FullDocumentBeforeChange

If true, indicates that `UPDATE` change events should
include a snapshot of the modified document from
immediately before the update was applied.

You must enable [document
preimages](https://www.mongodb.com/docs/atlas/app-services/mongodb/preimages/)
for your cluster to include these snapshots.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SkipCatchupEvents

If `true`, enabling the Trigger after it was disabled
will not invoke events that occurred while the Trigger
was disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TolerateResumeErrors

If `true`, when this Trigger's resume token
cannot be found in the cluster's oplog, the Trigger automatically resumes
processing events at the next relevant change stream event.
All change stream events from when the Trigger was suspended until the Trigger
resumes execution do not have the Trigger fire for them.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MaximumThroughput

If `true`, the trigger will use the maximize throughput option (https://www.mongodb.com/docs/atlas/app-services/triggers/database-triggers/#std-label-triggers-maximum-throughput).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Unordered

If `true`, event ordering is disabled and this Trigger
can process events in parallel. If `false`, event
ordering is enabled and the Trigger executes events
serially.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

