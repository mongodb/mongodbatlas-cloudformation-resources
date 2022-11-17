# Mongodb::Atlas::OrgInvitation

Returns, adds, and edits organizational units in MongoDB Cloud.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "Mongodb::Atlas::OrgInvitation",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">ApiKeyDefinition</a></i>,
        "<a href="#includecount" title="IncludeCount">IncludeCount</a>" : <i>Boolean</i>,
        "<a href="#invitationid" title="InvitationId">InvitationId</a>" : <i>String</i>,
        "<a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>" : <i>Integer</i>,
        "<a href="#orgid" title="OrgId">OrgId</a>" : <i>String</i>,
        "<a href="#orgname" title="OrgName">OrgName</a>" : <i>String</i>,
        "<a href="#pagenum" title="PageNum">PageNum</a>" : <i>Integer</i>,
        "<a href="#roles" title="Roles">Roles</a>" : <i>[ String, ... ]</i>,
        "<a href="#teamids" title="TeamIds">TeamIds</a>" : <i>[ String, ... ]</i>,
        "<a href="#totalcount" title="TotalCount">TotalCount</a>" : <i>Double</i>,
        "<a href="#username" title="Username">Username</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: Mongodb::Atlas::OrgInvitation
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">ApiKeyDefinition</a></i>
    <a href="#includecount" title="IncludeCount">IncludeCount</a>: <i>Boolean</i>
    <a href="#invitationid" title="InvitationId">InvitationId</a>: <i>String</i>
    <a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>: <i>Integer</i>
    <a href="#orgid" title="OrgId">OrgId</a>: <i>String</i>
    <a href="#orgname" title="OrgName">OrgName</a>: <i>String</i>
    <a href="#pagenum" title="PageNum">PageNum</a>: <i>Integer</i>
    <a href="#roles" title="Roles">Roles</a>: <i>
      - String</i>
    <a href="#teamids" title="TeamIds">TeamIds</a>: <i>
      - String</i>
    <a href="#totalcount" title="TotalCount">TotalCount</a>: <i>Double</i>
    <a href="#username" title="Username">Username</a>: <i>String</i>
</pre>

## Properties

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">ApiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IncludeCount

Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### InvitationId

Unique 24-hexadecimal digit string that identifies the invitation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ItemsPerPage

Number of items that the response returns per page.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgId

Unique 24-hexadecimal digit string that identifies the organization that contains your projects.

_Required_: No

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgName

Human-readable label that identifies this organization.

_Required_: No

_Type_: String

_Pattern_: <code>^[\p{L}\p{N}\-_.(),:&@+']{1,64}$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PageNum

Number of the page that displays the current set of the total objects that the response returns.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Roles

One or more organization or project level roles to assign to the MongoDB Cloud user.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TeamIds

List of unique 24-hexadecimal digit strings that identifies each team.

_Required_: No

_Type_: List of String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### TotalCount

Number of documents returned in this response.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Username

Email address of the MongoDB Cloud user invited to join the organization.

_Required_: No

_Type_: String

_Pattern_: <code>^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Id.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### ExpiresAt

Date and time when the invitation from MongoDB Cloud expires. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.

#### Id

Unique 24-hexadecimal digit string that identifies this organization.

#### CreatedAt

Date and time when MongoDB Cloud sent the invitation. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.

#### InviterUsername

Email address of the MongoDB Cloud user who sent the invitation to join the organization.

