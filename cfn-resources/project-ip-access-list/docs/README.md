# MongoDB::Atlas::ProjectIpAccessList

Returns, adds, edits, and removes network access limits to database deployments in MongoDB Cloud.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::ProjectIpAccessList",
    "Properties" : {
        "<a href="#accesslist" title="AccessList">AccessList</a>" : <i>[ <a href="accesslistdefinition.md">accessListDefinition</a>, ... ]</i>,
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#listoptions" title="ListOptions">ListOptions</a>" : <i><a href="listoptions.md">listOptions</a></i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::ProjectIpAccessList
Properties:
    <a href="#accesslist" title="AccessList">AccessList</a>: <i>
      - <a href="accesslistdefinition.md">accessListDefinition</a></i>
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#listoptions" title="ListOptions">ListOptions</a>: <i><a href="listoptions.md">listOptions</a></i>
</pre>

## Properties

#### AccessList

_Required_: Yes

_Type_: List of <a href="accesslistdefinition.md">accessListDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiKeys

_Required_: Yes

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ListOptions

_Required_: No

_Type_: <a href="listoptions.md">listOptions</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Id.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

The unique identifier for the Project API ip access list rules.

#### TotalCount

Number of documents returned in this response.

