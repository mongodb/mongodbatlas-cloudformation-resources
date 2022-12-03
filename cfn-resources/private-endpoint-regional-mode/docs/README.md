# MongoDB::Atlas::PrivateEndPointRegionalMode

An example resource schema demonstrating some basic constructs and validation rules.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::PrivateEndPointRegionalMode",
    "Properties" : {
        "<a href="#enabled" title="Enabled">Enabled</a>" : <i>Boolean</i>,
        "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>,
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikey.md">ApiKey</a></i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::PrivateEndPointRegionalMode
Properties:
    <a href="#enabled" title="Enabled">Enabled</a>: <i>Boolean</i>
    <a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikey.md">ApiKey</a></i>
</pre>

## Properties

#### Enabled

Name of the AWS PrivateLink endpoint service. Atlas returns null while it is creating the endpoint service.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### GroupId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ApiKeys

_Required_: No

_Type_: <a href="apikey.md">ApiKey</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the GroupId.
