# MongoDB::Atlas::DataLakes

Returns, adds, edits, and removes Federated Database Instances.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::DataLakes",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
        "<a href="#cloudproviderconfig" title="CloudProviderConfig">CloudProviderConfig</a>" : <i><a href="datalakecloudproviderconfigview.md">DataLakeCloudProviderConfigView</a></i>,
        "<a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>" : <i><a href="datalakedataprocessregionview.md">DataLakeDataProcessRegionView</a></i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::DataLakes
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    <a href="#cloudproviderconfig" title="CloudProviderConfig">CloudProviderConfig</a>: <i><a href="datalakecloudproviderconfigview.md">DataLakeCloudProviderConfigView</a></i>
    <a href="#dataprocessregion" title="DataProcessRegion">DataProcessRegion</a>: <i><a href="datalakedataprocessregionview.md">DataLakeDataProcessRegionView</a></i>
</pre>

## Properties

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CloudProviderConfig

_Required_: No

_Type_: <a href="datalakecloudproviderconfigview.md">DataLakeCloudProviderConfigView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DataProcessRegion

_Required_: No

_Type_: <a href="datalakedataprocessregionview.md">DataLakeDataProcessRegionView</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the GroupId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### GroupId

Unique 24-hexadecimal digit string that identifies your project.

#### DataProcessRegion

Returns the <code>DataProcessRegion</code> value.

#### TenantName

Human-readable label that identifies the Federated Database to remove.

#### Storage

Returns the <code>Storage</code> value.

#### SkipRoleValidation

Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.

#### StartDate

Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.

#### Hostnames

Human-readable label that identifies the Federated Database to update.

#### State

Human-readable label that identifies the Federated Database to update.

#### EndDate

Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.

#### TestS3Bucket

Returns the <code>TestS3Bucket</code> value.

#### NDSDataLakeDatabaseCollectionView

Returns the <code>NDSDataLakeDatabaseCollectionView</code> value.

#### DataLakeDatabaseDataSourceView

Returns the <code>DataLakeDatabaseDataSourceView</code> value.

#### DataLakeDatabaseView

Returns the <code>DataLakeDatabaseView</code> value.

#### NDSDataLakeDatabaseView

Returns the <code>NDSDataLakeDatabaseView</code> value.

#### DataLakeAWSCloudProviderConfigView

Returns the <code>DataLakeAWSCloudProviderConfigView</code> value.

#### DataLakeStorageView

Returns the <code>DataLakeStorageView</code> value.

#### ApiKeys

Returns the <code>ApiKeys</code> value.

