# MongoDB::Atlas::ResourcePolicy

Atlas Resource Policies

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::ResourcePolicy",
    "Properties" : {
        "<a href="#createdbyuser" title="CreatedByUser">CreatedByUser</a>" : <i><a href="apiatlasusermetadata.md">ApiAtlasUserMetadata</a></i>,
        "<a href="#lastupdatedbyuser" title="LastUpdatedByUser">LastUpdatedByUser</a>" : <i><a href="apiatlasusermetadata.md">ApiAtlasUserMetadata</a></i>,
        "<a href="#orgid" title="OrgId">OrgId</a>" : <i>String</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::ResourcePolicy
Properties:
    <a href="#createdbyuser" title="CreatedByUser">CreatedByUser</a>: <i><a href="apiatlasusermetadata.md">ApiAtlasUserMetadata</a></i>
    <a href="#lastupdatedbyuser" title="LastUpdatedByUser">LastUpdatedByUser</a>: <i><a href="apiatlasusermetadata.md">ApiAtlasUserMetadata</a></i>
    <a href="#orgid" title="OrgId">OrgId</a>: <i>String</i>
</pre>

## Properties

#### CreatedByUser

_Required_: No

_Type_: <a href="apiatlasusermetadata.md">ApiAtlasUserMetadata</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### LastUpdatedByUser

_Required_: No

_Type_: <a href="apiatlasusermetadata.md">ApiAtlasUserMetadata</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgId

Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Body

Returns the <code>Body</code> value.

#### Id

Returns the <code>Id</code> value.

#### Id

Returns the <code>Id</code> value.

#### Name

Returns the <code>Name</code> value.

#### CreatedByUser

Returns the <code>CreatedByUser</code> value.

#### CreatedDate

Date and time in UTC when the atlas resource policy was created.

#### Id

Unique 24-hexadecimal character string that identifies the atlas resource policy.

#### LastUpdatedByUser

Returns the <code>LastUpdatedByUser</code> value.

#### LastUpdatedDate

Date and time in UTC when the atlas resource policy was last updated.

#### Name

Human-readable label that describes the atlas resource policy.

#### Policies

List of policies that make up the atlas resource policy.

#### ResourcePolicyId

Unique 24-hexadecimal digit string that identifies an atlas resource policy.

#### Version

A string that identifies the version of the atlas resource policy.

