# Mongodb::Atlas::cloudProviderAccess

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "Mongodb::Atlas::cloudProviderAccess",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#iamassumedrolearn" title="IamAssumedRoleArn">IamAssumedRoleArn</a>" : <i>String</i>,
        "<a href="#iamawsregion" title="IamAwsRegion">IamAwsRegion</a>" : <i>String</i>,
    }
}
</pre>

### YAML

<pre>
Type: Mongodb::Atlas::cloudProviderAccess
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#iamassumedrolearn" title="IamAssumedRoleArn">IamAssumedRoleArn</a>: <i>String</i>
    <a href="#iamawsregion" title="IamAwsRegion">IamAwsRegion</a>: <i>String</i>
</pre>

## Properties

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: No

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IamAssumedRoleArn

Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account.

_Required_: No

_Type_: String

_Minimum_: <code>20</code>

_Maximum_: <code>2048</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IamAwsRegion

Valid AWS Region in which IAM assumed role will be provisioned for cloudprovider access from Atlas

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the RoleId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### RoleId

Unique 24-hexadecimal digit string that identifies the role.

#### AtlasAWSAccountArn

Amazon Resource Name that identifies the Amazon Web Services (AWS) user account that MongoDB Cloud uses when it assumes the Identity and Access Management (IAM) role.

#### AuthorizedDate

Date and time when someone authorized this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### AtlasAssumedRoleExternalId

Unique external ID that MongoDB Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account.

#### CreatedDate

Date and time when someone created this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### FeatureUsages

Returns the <code>FeatureUsages</code> value.

