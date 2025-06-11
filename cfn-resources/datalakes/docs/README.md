# MongoDB::Atlas::DataLakes

Returns, adds, edits, and removes Federated Database Instances.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::DataLakes",
    "Properties" : {
        "<a href="#cloudproviderconfig" title="CloudProviderConfig">CloudProviderConfig</a>" : <i><a href="datalakecloudproviderconfigview.md">DataLakeCloudProviderConfigView</a></i>,
        "<a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>" : <i><a href="datalakedataprocessregionview.md">DataLakeDataProcessRegionView</a></i>,
        "<a href="#enddate" title="EndDate">EndDate</a>" : <i>Double</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#skiprolevalidation" title="SkipRoleValidation">SkipRoleValidation</a>" : <i>Boolean</i>,
        "<a href="#storage" title="Storage">Storage</a>" : <i><a href="datalakestorageview.md">DataLakeStorageView</a></i>,
        "<a href="#tenantname" title="TenantName">TenantName</a>" : <i>String</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::DataLakes
Properties:
    <a href="#cloudproviderconfig" title="CloudProviderConfig">CloudProviderConfig</a>: <i><a href="datalakecloudproviderconfigview.md">DataLakeCloudProviderConfigView</a></i>
    <a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>: <i><a href="datalakedataprocessregionview.md">DataLakeDataProcessRegionView</a></i>
    <a href="#enddate" title="EndDate">EndDate</a>: <i>Double</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#skiprolevalidation" title="SkipRoleValidation">SkipRoleValidation</a>: <i>Boolean</i>
    <a href="#storage" title="Storage">Storage</a>: <i><a href="datalakestorageview.md">DataLakeStorageView</a></i>
    <a href="#tenantname" title="TenantName">TenantName</a>: <i>String</i>
</pre>

## Properties

#### CloudProviderConfig

_Required_: No

_Type_: <a href="datalakecloudproviderconfigview.md">DataLakeCloudProviderConfigView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DataProcessRegion

_Required_: No

_Type_: <a href="datalakedataprocessregionview.md">DataLakeDataProcessRegionView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EndDate

Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SkipRoleValidation

Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Storage

_Required_: No

_Type_: <a href="datalakestorageview.md">DataLakeStorageView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TenantName

Human-readable label that identifies the Federated Database to remove.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Storage

Returns the <code>Storage</code> value.

#### StartDate

Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.

#### Hostnames

Human-readable label that identifies the Federated Database to update.

#### State

Human-readable label that identifies the Federated Database to update.

#### TestS3Bucket

Returns the <code>TestS3Bucket</code> value.

