# MongoDB::Atlas::BackupCompliancePolicy OnDemandPolicyItem

On-demand backup policy item configuration. When provided, FrequencyInterval, RetentionUnit, and RetentionValue are required.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#frequencyinterval" title="FrequencyInterval">FrequencyInterval</a>" : <i>Integer</i>,
    "<a href="#retentionunit" title="RetentionUnit">RetentionUnit</a>" : <i>String</i>,
    "<a href="#retentionvalue" title="RetentionValue">RetentionValue</a>" : <i>Integer</i>
}
</pre>

### YAML

<pre>
<a href="#frequencyinterval" title="FrequencyInterval">FrequencyInterval</a>: <i>Integer</i>
<a href="#retentionunit" title="RetentionUnit">RetentionUnit</a>: <i>String</i>
<a href="#retentionvalue" title="RetentionValue">RetentionValue</a>: <i>Integer</i>
</pre>

## Properties

#### FrequencyInterval

Number that indicates the frequency interval for a set of snapshots. Required when OnDemandPolicyItem is provided.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RetentionUnit

Unit of time in which MongoDB Cloud measures snapshot retention. Required when OnDemandPolicyItem is provided.

_Required_: No

_Type_: String

_Allowed Values_: <code>days</code> | <code>weeks</code> | <code>months</code> | <code>years</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RetentionValue

Duration in days, weeks, months, or years that MongoDB Cloud retains the snapshot. Required when OnDemandPolicyItem is provided.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

