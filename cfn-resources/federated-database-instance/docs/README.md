# MongoDB::Atlas::FederatedDatabaseInstance

Returns, adds, edits, and removes Federated Database Instances.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::FederatedDatabaseInstance",
    "Properties" : {
        "<a href="#cloudproviderconfig" title="CloudProviderConfig">CloudProviderConfig</a>" : <i><a href="cloudproviderconfig.md">CloudProviderConfig</a></i>,
        "<a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>" : <i><a href="dataprocessregion.md">DataProcessRegion</a></i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#tenantname" title="TenantName">TenantName</a>" : <i>String</i>,
        "<a href="#skiprolevalidation" title="SkipRoleValidation">SkipRoleValidation</a>" : <i>Boolean</i>,
        "<a href="#storage" title="Storage">Storage</a>" : <i><a href="storage.md">Storage</a></i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::FederatedDatabaseInstance
Properties:
    <a href="#cloudproviderconfig" title="CloudProviderConfig">CloudProviderConfig</a>: <i><a href="cloudproviderconfig.md">CloudProviderConfig</a></i>
    <a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>: <i><a href="dataprocessregion.md">DataProcessRegion</a></i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#tenantname" title="TenantName">TenantName</a>: <i>String</i>
    <a href="#skiprolevalidation" title="SkipRoleValidation">SkipRoleValidation</a>: <i>Boolean</i>
    <a href="#storage" title="Storage">Storage</a>: <i><a href="storage.md">Storage</a></i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
</pre>

## Properties

#### CloudProviderConfig

Cloud provider linked to this Atlas Data Federation.

_Required_: No

_Type_: <a href="cloudproviderconfig.md">CloudProviderConfig</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DataProcessRegion

Information about the cloud provider region to which the Atlas Data Federation routes client connections. MongoDB Cloud supports AWS only.

_Required_: No

_Type_: <a href="dataprocessregion.md">DataProcessRegion</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### TenantName

Human-readable label that identifies the data federation.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### SkipRoleValidation

Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Storage

Configuration information for each data store and its mapping to MongoDB Cloud databases.

_Required_: No

_Type_: <a href="storage.md">Storage</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### ExternalId

Returns the <code>ExternalId</code> value.

#### IamAssumedRoleARN

Returns the <code>IamAssumedRoleARN</code> value.

#### IamUserARN

Returns the <code>IamUserARN</code> value.

#### HostNames

Type of Federated Database Instances to return.

#### State

Type of Federated Database Instances to return.

