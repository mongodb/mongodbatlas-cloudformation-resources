# MongoDB::Atlas::ServerlessInstance

**WARNING:** This resource is deprecated and will be removed in January 2026. For more details, see [Migrate your programmatic tools from M2, M5, or Serverless Instances to Flex Clusters](https://www.mongodb.com/docs/atlas/flex-migration/). Returns, adds, edits, and removes serverless instances.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::ServerlessInstance",
    "Properties" : {
        "<a href="#connectionstrings" title="ConnectionStrings">ConnectionStrings</a>" : <i><a href="serverlessinstanceconnectionstrings.md">ServerlessInstanceConnectionStrings</a></i>,
        "<a href="#continuousbackupenabled" title="ContinuousBackupEnabled">ContinuousBackupEnabled</a>" : <i>Boolean</i>,
        "<a href="#includecount" title="IncludeCount">IncludeCount</a>" : <i>Boolean</i>,
        "<a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>" : <i>Integer</i>,
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#pagenum" title="PageNum">PageNum</a>" : <i>Integer</i>,
        "<a href="#projectid" title="ProjectID">ProjectID</a>" : <i>String</i>,
        "<a href="#providersettings" title="ProviderSettings">ProviderSettings</a>" : <i><a href="serverlessinstanceprovidersettings.md">ServerlessInstanceProviderSettings</a></i>,
        "<a href="#terminationprotectionenabled" title="TerminationProtectionEnabled">TerminationProtectionEnabled</a>" : <i>Boolean</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::ServerlessInstance
Properties:
    <a href="#connectionstrings" title="ConnectionStrings">ConnectionStrings</a>: <i><a href="serverlessinstanceconnectionstrings.md">ServerlessInstanceConnectionStrings</a></i>
    <a href="#continuousbackupenabled" title="ContinuousBackupEnabled">ContinuousBackupEnabled</a>: <i>Boolean</i>
    <a href="#includecount" title="IncludeCount">IncludeCount</a>: <i>Boolean</i>
    <a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>: <i>Integer</i>
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#pagenum" title="PageNum">PageNum</a>: <i>Integer</i>
    <a href="#projectid" title="ProjectID">ProjectID</a>: <i>String</i>
    <a href="#providersettings" title="ProviderSettings">ProviderSettings</a>: <i><a href="serverlessinstanceprovidersettings.md">ServerlessInstanceProviderSettings</a></i>
    <a href="#terminationprotectionenabled" title="TerminationProtectionEnabled">TerminationProtectionEnabled</a>: <i>Boolean</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
</pre>

## Properties

#### ConnectionStrings

_Required_: No

_Type_: <a href="serverlessinstanceconnectionstrings.md">ServerlessInstanceConnectionStrings</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ContinuousBackupEnabled

Flag that indicates whether the serverless instances uses Serverless Continuous Backup. If this parameter is false, the serverless instance uses Basic Backup. | Option | Description | |---|---| | Serverless Continuous Backup | Atlas takes incremental snapshots of the data in your serverless instance every six hours and lets you restore the data from a selected point in time within the last 72 hours. Atlas also takes daily snapshots and retains these daily snapshots for 35 days. To learn more, see Serverless Instance Costs. | | Basic Backup | Atlas takes incremental snapshots of the data in your serverless instance every six hours and retains only the two most recent snapshots. You can use this option for free. 

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IncludeCount

Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ItemsPerPage

Number of items that the response returns per page.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Name

Human-readable label that identifies the serverless instance.

_Required_: No

_Type_: String

_Minimum Length_: <code>1</code>

_Maximum Length_: <code>64</code>

_Pattern_: <code>^[a-zA-Z0-9][a-zA-Z0-9-]*$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### PageNum

Number of the page that displays the current set of the total objects that the response returns.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectID

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProviderSettings

_Required_: No

_Type_: <a href="serverlessinstanceprovidersettings.md">ServerlessInstanceProviderSettings</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TerminationProtectionEnabled

Flag that indicates whether termination protection is enabled on the serverless instance. If set to true, MongoDB Cloud won't delete the serverless instance. If set to false, MongoDB cloud will delete the serverless instance."

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### CreateDate

Date and time when MongoDB Cloud created this serverless instance. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.

#### Id

Unique 24-hexadecimal digit string that identifies the serverless instance.

#### TotalCount

Number of documents returned in this response.

#### ConnectionStrings

Returns the <code>ConnectionStrings</code> value.

#### StateName

Human-readable label that indicates the current operating condition of the serverless instance.

#### MongoDBVersion

Version of MongoDB that the serverless instance runs.

#### StandardSrv

Returns the <code>StandardSrv</code> value.

#### PrivateEndpoint

Returns the <code>PrivateEndpoint</code> value.

#### SrvConnectionString

Returns the <code>SrvConnectionString</code> value.

#### Type

Returns the <code>Type</code> value.

#### EndpointId

Returns the <code>EndpointId</code> value.

#### ProviderName

Returns the <code>ProviderName</code> value.

#### Region

Returns the <code>Region</code> value.

