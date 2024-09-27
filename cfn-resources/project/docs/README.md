# MongoDB::Atlas::Project

Retrieves or creates projects in any given Atlas organization.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::Project",
    "Properties" : {
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#orgid" title="OrgId">OrgId</a>" : <i>String</i>,
        "<a href="#projectownerid" title="ProjectOwnerId">ProjectOwnerId</a>" : <i>String</i>,
        "<a href="#withdefaultalertssettings" title="WithDefaultAlertsSettings">WithDefaultAlertsSettings</a>" : <i>Boolean</i>,
        "<a href="#projectsettings" title="ProjectSettings">ProjectSettings</a>" : <i><a href="projectsettings.md">projectSettings</a></i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectteams" title="ProjectTeams">ProjectTeams</a>" : <i>[ <a href="projectteam.md">projectTeam</a>, ... ]</i>,
        "<a href="#projectapikeys" title="ProjectApiKeys">ProjectApiKeys</a>" : <i>[ <a href="projectapikey.md">projectApiKey</a>, ... ]</i>,
        "<a href="#regionusagerestrictions" title="RegionUsageRestrictions">RegionUsageRestrictions</a>" : <i>String</i>,
        "<a href="#tags" title="Tags">Tags</a>" : <i><a href="tags.md">Tags</a></i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::Project
Properties:
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#orgid" title="OrgId">OrgId</a>: <i>String</i>
    <a href="#projectownerid" title="ProjectOwnerId">ProjectOwnerId</a>: <i>String</i>
    <a href="#withdefaultalertssettings" title="WithDefaultAlertsSettings">WithDefaultAlertsSettings</a>: <i>Boolean</i>
    <a href="#projectsettings" title="ProjectSettings">ProjectSettings</a>: <i><a href="projectsettings.md">projectSettings</a></i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectteams" title="ProjectTeams">ProjectTeams</a>: <i>
      - <a href="projectteam.md">projectTeam</a></i>
    <a href="#projectapikeys" title="ProjectApiKeys">ProjectApiKeys</a>: <i>
      - <a href="projectapikey.md">projectApiKey</a></i>
    <a href="#regionusagerestrictions" title="RegionUsageRestrictions">RegionUsageRestrictions</a>: <i>String</i>
    <a href="#tags" title="Tags">Tags</a>: <i><a href="tags.md">Tags</a></i>
</pre>

## Properties

#### Name

Name of the project to create.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgId

Unique identifier of the organization within which to create the project.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectOwnerId

Unique identifier of the organization within which to create the project.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### WithDefaultAlertsSettings

Flag that indicates whether to create the project with default alert settings.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectSettings

_Required_: No

_Type_: <a href="projectsettings.md">projectSettings</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectTeams

Teams to which the authenticated user has access in the project specified using its unique 24-hexadecimal digit identifier.

_Required_: No

_Type_: List of <a href="projectteam.md">projectTeam</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectApiKeys

API keys that you assigned to the specified project.

_Required_: No

_Type_: List of <a href="projectapikey.md">projectApiKey</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RegionUsageRestrictions

Region usage restrictions that designate the project's AWS region.Enum: "GOV_REGIONS_ONLY" "COMMERCIAL_FEDRAMP_REGIONS_ONLY" "NONE"

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Tags

Map that contains key values between 1 to 255 characters in length for tagging and categorizing the project. To learn more, see https://www.mongodb.com/docs/atlas/tags/

_Required_: No

_Type_: <a href="tags.md">Tags</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

The unique identifier of the project.

#### Created

The ISO-8601-formatted timestamp of when Atlas created the project.

#### ClusterCount

The number of Atlas clusters deployed in the project.

