# MongoDB::Atlas::DataFederation

Returns, adds, edits, and removes Federated Database Instances.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::DataFederation",
    "Properties" : {
        "<a href="#cloudproviderconfig" title="CloudProviderConfig">CloudProviderConfig</a>" : <i><a href="atlasdatalakecloudproviderconfig.md">AtlasDataLakeCloudProviderConfig</a></i>,
        "<a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>" : <i><a href="atlasdatalakedataprocessregion.md">AtlasDataLakeDataProcessRegion</a></i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#skiprolevalidation" title="SkipRoleValidation">SkipRoleValidation</a>" : <i>Boolean</i>,
        "<a href="#storage" title="Storage">Storage</a>" : <i><a href="atlasdatalakestorage.md">AtlasDataLakeStorage</a></i>,
        "<a href="#type" title="Type">Type</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::DataFederation
Properties:
    <a href="#cloudproviderconfig" title="CloudProviderConfig">CloudProviderConfig</a>: <i><a href="atlasdatalakecloudproviderconfig.md">AtlasDataLakeCloudProviderConfig</a></i>
    <a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>: <i><a href="atlasdatalakedataprocessregion.md">AtlasDataLakeDataProcessRegion</a></i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#skiprolevalidation" title="SkipRoleValidation">SkipRoleValidation</a>: <i>Boolean</i>
    <a href="#storage" title="Storage">Storage</a>: <i><a href="atlasdatalakestorage.md">AtlasDataLakeStorage</a></i>
    <a href="#type" title="Type">Type</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
</pre>

## Properties

#### CloudProviderConfig

Name of the cloud service that hosts the data lake's data stores.

_Required_: No

_Type_: <a href="atlasdatalakecloudproviderconfig.md">AtlasDataLakeCloudProviderConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DataProcessRegion

Information about the cloud provider region to which the data lake routes client connections. MongoDB Cloud supports AWS only.

_Required_: No

_Type_: <a href="atlasdatalakedataprocessregion.md">AtlasDataLakeDataProcessRegion</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### SkipRoleValidation

Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Storage

_Required_: No

_Type_: <a href="atlasdatalakestorage.md">AtlasDataLakeStorage</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Type

Type of Federated Database Instances to return.

_Required_: No

_Type_: String

_Allowed Values_: <code>USER</code> | <code>ONLINE_ARCHIVE</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ProjectId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### ExternalId

Returns the <code>ExternalId</code> value.

#### IamAssumedRoleARN

Returns the <code>IamAssumedRoleARN</code> value.

#### IamUserARN

Returns the <code>IamUserARN</code> value.

#### Name

Human-readable label that identifies the data lake.

