# MongoDB::Atlas::FederatedDatabaseInstance ReadPreference

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#mode" title="Mode">Mode</a>" : <i>String</i>,
    "<a href="#maxstalenessseconds" title="MaxStalenessSeconds">MaxStalenessSeconds</a>" : <i>String</i>,
    "<a href="#tagsets" title="TagSets">TagSets</a>" : <i>[ [ <a href="tagset.md">TagSet</a>, ... ], ... ]</i>
}
</pre>

### YAML

<pre>
<a href="#mode" title="Mode">Mode</a>: <i>String</i>
<a href="#maxstalenessseconds" title="MaxStalenessSeconds">MaxStalenessSeconds</a>: <i>String</i>
<a href="#tagsets" title="TagSets">TagSets</a>: <i>
      - 
      - <a href="tagset.md">TagSet</a></i>
</pre>

## Properties

#### Mode

"primary" "primaryPreferred" "secondary" "secondaryPreferred" "nearest"
Read preference mode that specifies to which replica set member to route the read requests.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MaxStalenessSeconds

Maximum replication lag, or staleness, for reads from secondaries.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TagSets

List that contains tag sets or tag specification documents. If specified, Atlas Data Federation routes read requests to replica set member or members that are associated with the specified tags.

_Required_: No

_Type_: List of List of <a href="tagset.md">TagSet</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

