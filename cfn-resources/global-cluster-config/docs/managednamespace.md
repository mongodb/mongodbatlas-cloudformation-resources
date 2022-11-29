# MongoDB::Atlas::GlobalClusterConfig managedNamespace

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#collection" title="Collection">Collection</a>" : <i>String</i>,
    "<a href="#customshardkey" title="CustomShardKey">CustomShardKey</a>" : <i>String</i>,
    "<a href="#db" title="Db">Db</a>" : <i>String</i>,
    "<a href="#iscustomshardkeyhashed" title="IsCustomShardKeyHashed">IsCustomShardKeyHashed</a>" : <i>Boolean</i>,
    "<a href="#isshardkeyunique" title="IsShardKeyUnique">IsShardKeyUnique</a>" : <i>Boolean</i>,
    "<a href="#numinitialchunks" title="NumInitialChunks">NumInitialChunks</a>" : <i>Double</i>
}
</pre>

### YAML

<pre>
<a href="#collection" title="Collection">Collection</a>: <i>String</i>
<a href="#customshardkey" title="CustomShardKey">CustomShardKey</a>: <i>String</i>
<a href="#db" title="Db">Db</a>: <i>String</i>
<a href="#iscustomshardkeyhashed" title="IsCustomShardKeyHashed">IsCustomShardKeyHashed</a>: <i>Boolean</i>
<a href="#isshardkeyunique" title="IsShardKeyUnique">IsShardKeyUnique</a>: <i>Boolean</i>
<a href="#numinitialchunks" title="NumInitialChunks">NumInitialChunks</a>: <i>Double</i>
</pre>

## Properties

#### Collection

Human-readable label of the collection to manage for this Global Cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CustomShardKey

Database parameter used to divide the *collection* into shards. Global clusters require a compound shard key. This compound shard key combines the location parameter and the user-selected custom key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Db

Human-readable label of the database to manage for this Global Cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IsCustomShardKeyHashed

Flag that indicates whether someone hashed the custom shard key for the specified collection. If you set this value to `false`, MongoDB Cloud uses ranged sharding.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IsShardKeyUnique

Flag that indicates whether someone [hashed](https://www.mongodb.com/docs/manual/reference/method/sh.shardCollection/#hashed-shard-keys) the custom shard key. If this parameter returns `false`, this cluster uses [ranged sharding](https://www.mongodb.com/docs/manual/core/ranged-sharding/).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NumInitialChunks

Minimum number of chunks to create initially when sharding an empty collection with a [hashed shard key](https://www.mongodb.com/docs/manual/core/hashed-sharding/).

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

