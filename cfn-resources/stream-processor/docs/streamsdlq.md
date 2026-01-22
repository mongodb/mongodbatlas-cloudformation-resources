# MongoDB::Atlas::StreamProcessor StreamsDLQ

Dead letter queue for the stream processor. Refer to the MongoDB Atlas Docs for more information.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#coll" title="Coll">Coll</a>" : <i>String</i>,
    "<a href="#connectionname" title="ConnectionName">ConnectionName</a>" : <i>String</i>,
    "<a href="#db" title="Db">Db</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#coll" title="Coll">Coll</a>: <i>String</i>
<a href="#connectionname" title="ConnectionName">ConnectionName</a>: <i>String</i>
<a href="#db" title="Db">Db</a>: <i>String</i>
</pre>

## Properties

#### Coll

Name of the collection to use for the DLQ.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ConnectionName

Name of the connection to write DLQ messages to. Must be an Atlas connection.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Db

Name of the database to use for the DLQ.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

