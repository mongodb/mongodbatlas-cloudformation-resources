# MongoDB::Atlas::StreamProcessor StreamsOptions

Optional configuration for the stream processor.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#dlq" title="Dlq">Dlq</a>" : <i><a href="streamsdlq.md">StreamsDLQ</a></i>
}
</pre>

### YAML

<pre>
<a href="#dlq" title="Dlq">Dlq</a>: <i><a href="streamsdlq.md">StreamsDLQ</a></i>
</pre>

## Properties

#### Dlq

Dead letter queue for the stream processor. Refer to the MongoDB Atlas Docs for more information.

_Required_: Yes

_Type_: <a href="streamsdlq.md">StreamsDLQ</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

