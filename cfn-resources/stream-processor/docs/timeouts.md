# MongoDB::Atlas::StreamProcessor Timeouts

Configurable timeouts for stream processor operations.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#create" title="Create">Create</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#create" title="Create">Create</a>: <i>String</i>
</pre>

## Properties

#### Create

Timeout for create operation in Go duration format (e.g., '5m', '10s'). Default is 20 minutes.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

