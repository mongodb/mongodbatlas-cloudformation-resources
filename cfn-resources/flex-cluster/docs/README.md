# MongoDB::Atlas::FlexCluster

The flex cluster resource provides access to your flex cluster configurations. The resource lets you create, edit and delete flex clusters. The resource requires your Project ID.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::FlexCluster",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#providersettings" title="ProviderSettings">ProviderSettings</a>" : <i><a href="providersettings.md">ProviderSettings</a></i>,
        "<a href="#terminationprotectionenabled" title="TerminationProtectionEnabled">TerminationProtectionEnabled</a>" : <i>Boolean</i>,
        "<a href="#tags" title="Tags">Tags</a>" : <i>[ <a href="tag.md">tag</a>, ... ]</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::FlexCluster
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#providersettings" title="ProviderSettings">ProviderSettings</a>: <i><a href="providersettings.md">ProviderSettings</a></i>
    <a href="#terminationprotectionenabled" title="TerminationProtectionEnabled">TerminationProtectionEnabled</a>: <i>Boolean</i>
    <a href="#tags" title="Tags">Tags</a>: <i>
      - <a href="tag.md">tag</a></i>
</pre>

## Properties

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique identifier of the project the cluster belongs to.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Name

Human-readable label that identifies the flex cluster.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProviderSettings

Group of cloud provider settings that configure the provisioned MongoDB flex cluster.

_Required_: Yes

_Type_: <a href="providersettings.md">ProviderSettings</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TerminationProtectionEnabled

Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Tags

Map that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the flex cluster.

_Required_: No

_Type_: List of <a href="tag.md">tag</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

Unique 24-hexadecimal digit string that identifies the flex cluster.

#### StateName

Human-readable label that indicates the current operating condition of this flex cluster.

#### BackupSettings

Flex backup configuration

#### ClusterType

Flex cluster topology.

#### ConnectionStrings

Collection of Uniform Resource Locators that point to the MongoDB database.

#### CreateDate

Date and time when MongoDB Cloud created this flex cluster. This parameter expresses its value in ISO 8601 format in UTC.

#### MongoDBVersion

Version of MongoDB that the flex cluster runs.

#### VersionReleaseSystem

Method by which the cluster maintains the MongoDB versions.

#### DiskSizeGB

Returns the <code>DiskSizeGB</code> value.

#### ProviderName

Returns the <code>ProviderName</code> value.

