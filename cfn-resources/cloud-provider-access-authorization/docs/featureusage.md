# MongoDB::Atlas::CloudProviderAccessAuthorization FeatureUsage

MongoDB Cloud feature associated with the authorized AWS IAM role.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#featuretype" title="FeatureType">FeatureType</a>" : <i>String</i>,
    "<a href="#featureid" title="FeatureId">FeatureId</a>" : <i><a href="featureusageid.md">FeatureUsageId</a></i>
}
</pre>

### YAML

<pre>
<a href="#featuretype" title="FeatureType">FeatureType</a>: <i>String</i>
<a href="#featureid" title="FeatureId">FeatureId</a>: <i><a href="featureusageid.md">FeatureUsageId</a></i>
</pre>

## Properties

#### FeatureType

Human-readable label that describes one MongoDB Cloud feature linked to this AWS IAM role.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FeatureId

Identifying characteristics about the AWS Simple Storage Service (S3) export bucket linked to this AWS IAM role.

_Required_: No

_Type_: <a href="featureusageid.md">FeatureUsageId</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

