# MongoDB::Atlas::ProjectIPWhitelist

Atlas only allows client connections to the cluster from entries in the project’s whitelist. Each entry is either a single IP address or a CIDR-notated range of addresses.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::ProjectIPWhitelist",
    "Properties" : {
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#whitelist" title="Whitelist">Whitelist</a>" : <i>[ <a href="whitelist.md">Whitelist</a>, ... ]</i>,
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeys.md">ApiKeys</a></i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::ProjectIPWhitelist
Properties:
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#whitelist" title="Whitelist">Whitelist</a>: <i>
      - <a href="whitelist.md">Whitelist</a></i>
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeys.md">ApiKeys</a></i>
</pre>

## Properties

#### ProjectId

The unique identifier for the project to which you want to add one or more whitelist entries.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Whitelist

_Required_: Yes

_Type_: List of <a href="whitelist.md">Whitelist</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiKeys

_Required_: Yes

_Type_: <a href="apikeys.md">ApiKeys</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Id.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

The unique identifier for the Project API Whitelist rules.

#### TotalCount

The unique identifier for the Project API Whitelist rules.

