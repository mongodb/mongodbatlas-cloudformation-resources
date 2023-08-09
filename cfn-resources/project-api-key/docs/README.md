# MongoDB::Atlas::ProjectAPIKey

Creates one API key for the specified organization. An organization API key grants programmatic access to an organization.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::ProjectAPIKey",
    "Properties" : {
        "<a href="#description" title="Description">Description</a>" : <i>String</i>,
        "<a href="#orgid" title="OrgId">OrgId</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#roles" title="Roles">Roles</a>" : <i>[ String, ... ]</i>,
        "<a href="#listoptions" title="ListOptions">ListOptions</a>" : <i><a href="listoptions.md">ListOptions</a></i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::ProjectAPIKey
Properties:
    <a href="#description" title="Description">Description</a>: <i>String</i>
    <a href="#orgid" title="OrgId">OrgId</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#roles" title="Roles">Roles</a>: <i>
      - String</i>
    <a href="#listoptions" title="ListOptions">ListOptions</a>: <i><a href="listoptions.md">ListOptions</a></i>
</pre>

## Properties

#### Description

Purpose or explanation provided when someone created this project API key.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgId

Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the /orgs endpoint to retrieve all organizations to which the authenticated user has access.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project. Use the groups endpoint to retrieve all projects to which the authenticated user has access.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Profile

Redacted private key returned for this organization API key. This key displays unredacted when first created.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Roles

List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this organization.

_Required_: No

_Type_: List of String

_Allowed Values_: <code>ORG_OWNER</code> | <code>ORG_MEMBER</code> | <code>ORG_GROUP_CREATOR</code> | <code>ORG_BILLING_ADMIN</code> | <code>ORG_READ_ONLY</code> | <code>ORG_TEAM_MEMBERS_ADMIN</code> | <code>GROUP_ATLAS_ADMIN</code> | <code>GROUP_AUTOMATION_ADMIN</code> | <code>GROUP_BACKUP_ADMIN</code> | <code>GROUP_MONITORING_ADMIN</code> | <code>GROUP_OWNER</code> | <code>GROUP_READ_ONLY</code> | <code>GROUP_USER_ADMIN</code> | <code>GROUP_BILLING_ADMIN</code> | <code>GROUP_DATA_ACCESS_ADMIN</code> | <code>GROUP_DATA_ACCESS_READ_ONLY</code> | <code>GROUP_DATA_ACCESS_READ_WRITE</code> | <code>GROUP_CHARTS_ADMIN</code> | <code>GROUP_CLUSTER_MANAGER</code> | <code>GROUP_SEARCH_INDEX_EDITOR</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ListOptions

_Required_: No

_Type_: <a href="listoptions.md">ListOptions</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Links

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.

#### PrivateKey

Purpose or explanation provided when someone created this organization API key.

#### PublicKey

Public API key value set for the specified organization API key.

#### Id

Unique 24-hexadecimal digit string that identifies this organization API key assigned to this project.

