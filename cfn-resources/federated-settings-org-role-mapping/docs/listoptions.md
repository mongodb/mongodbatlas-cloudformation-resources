# MongoDB::Atlas::FederatedSettingsOrgRoleMapping ListOptions

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#pagenum" title="PageNum">PageNum</a>" : <i>Integer</i>,
    "<a href="#includecount" title="IncludeCount">IncludeCount</a>" : <i>Boolean</i>,
    "<a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#pagenum" title="PageNum">PageNum</a>: <i>Integer</i>
<a href="#includecount" title="IncludeCount">IncludeCount</a>: <i>Boolean</i>
<a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>: <i>Integer</i>
</pre>

## Properties

#### PageNum

Unique identifier of the AWS bucket to export the cloud backup snapshot to

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IncludeCount

Frequency associated with the export policy. Value can be daily, weekly, or monthly.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ItemsPerPage

Frequency associated with the export policy. Value can be daily, weekly, or monthly.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

