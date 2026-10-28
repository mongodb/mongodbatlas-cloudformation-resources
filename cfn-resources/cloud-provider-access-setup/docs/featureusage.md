# MongoDB::Atlas::CloudProviderAccessSetup FeatureUsage

One Atlas feature that uses the Cloud Provider Access role. A role that a feature still uses cannot be deleted.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#featuretype" title="FeatureType">FeatureType</a>" : <i>String</i>,
    "<a href="#featureid" title="FeatureId">FeatureId</a>" : <i><a href="featureid.md">FeatureId</a></i>
}
</pre>

### YAML

<pre>
<a href="#featuretype" title="FeatureType">FeatureType</a>: <i>String</i>
<a href="#featureid" title="FeatureId">FeatureId</a>: <i><a href="featureid.md">FeatureId</a></i>
</pre>

## Properties

#### FeatureType

Human-readable label that describes the Atlas feature that uses the role, for example ENCRYPTION_AT_REST.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FeatureId

Identifying characteristics of the feature that uses the role, when Atlas returns them, for example the S3 bucket name for a push-based log export. Not set for features that Atlas does not identify, for example ENCRYPTION_AT_REST.

_Required_: No

_Type_: <a href="featureid.md">FeatureId</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

