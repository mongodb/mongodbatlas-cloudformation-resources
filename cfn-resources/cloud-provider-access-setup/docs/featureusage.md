# MongoDB::Atlas::CloudProviderAccessSetup FeatureUsage

One Atlas feature that uses the Cloud Provider Access role. A role that a feature still uses cannot be deleted. Only the feature type is returned; the Atlas API's feature identifier is not exposed because its shape varies by feature type.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#featuretype" title="FeatureType">FeatureType</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#featuretype" title="FeatureType">FeatureType</a>: <i>String</i>
</pre>

## Properties

#### FeatureType

Human-readable label that describes the Atlas feature that uses the role, for example ENCRYPTION_AT_REST.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

