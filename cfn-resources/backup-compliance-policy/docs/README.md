# MongoDB::Atlas::BackupCompliancePolicy

Resource for managing MongoDB Atlas Backup Compliance Policy. Backup Compliance Policy prevents any user, regardless of role, from modifying or deleting specific cluster settings, backups, and backup configurations. When enabled, the Backup Compliance Policy will be applied as the minimum policy for all clusters and backups in the project.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::BackupCompliancePolicy",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#authorizedemail" title="AuthorizedEmail">AuthorizedEmail</a>" : <i>String</i>,
        "<a href="#authorizeduserfirstname" title="AuthorizedUserFirstName">AuthorizedUserFirstName</a>" : <i>String</i>,
        "<a href="#authorizeduserlastname" title="AuthorizedUserLastName">AuthorizedUserLastName</a>" : <i>String</i>,
        "<a href="#copyprotectionenabled" title="CopyProtectionEnabled">CopyProtectionEnabled</a>" : <i>Boolean</i>,
        "<a href="#encryptionatrestenabled" title="EncryptionAtRestEnabled">EncryptionAtRestEnabled</a>" : <i>Boolean</i>,
        "<a href="#restorewindowdays" title="RestoreWindowDays">RestoreWindowDays</a>" : <i>Integer</i>,
        "<a href="#ondemandpolicyitem" title="OnDemandPolicyItem">OnDemandPolicyItem</a>" : <i><a href="ondemandpolicyitem.md">OnDemandPolicyItem</a></i>,
        "<a href="#pitenabled" title="PitEnabled">PitEnabled</a>" : <i>Boolean</i>,
        "<a href="#policyitemhourly" title="PolicyItemHourly">PolicyItemHourly</a>" : <i><a href="scheduledpolicyitem.md">ScheduledPolicyItem</a></i>,
        "<a href="#policyitemdaily" title="PolicyItemDaily">PolicyItemDaily</a>" : <i><a href="scheduledpolicyitem.md">ScheduledPolicyItem</a></i>,
        "<a href="#policyitemweekly" title="PolicyItemWeekly">PolicyItemWeekly</a>" : <i>[ <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a>, ... ]</i>,
        "<a href="#policyitemmonthly" title="PolicyItemMonthly">PolicyItemMonthly</a>" : <i>[ <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a>, ... ]</i>,
        "<a href="#policyitemyearly" title="PolicyItemYearly">PolicyItemYearly</a>" : <i>[ <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a>, ... ]</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::BackupCompliancePolicy
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#authorizedemail" title="AuthorizedEmail">AuthorizedEmail</a>: <i>String</i>
    <a href="#authorizeduserfirstname" title="AuthorizedUserFirstName">AuthorizedUserFirstName</a>: <i>String</i>
    <a href="#authorizeduserlastname" title="AuthorizedUserLastName">AuthorizedUserLastName</a>: <i>String</i>
    <a href="#copyprotectionenabled" title="CopyProtectionEnabled">CopyProtectionEnabled</a>: <i>Boolean</i>
    <a href="#encryptionatrestenabled" title="EncryptionAtRestEnabled">EncryptionAtRestEnabled</a>: <i>Boolean</i>
    <a href="#restorewindowdays" title="RestoreWindowDays">RestoreWindowDays</a>: <i>Integer</i>
    <a href="#ondemandpolicyitem" title="OnDemandPolicyItem">OnDemandPolicyItem</a>: <i><a href="ondemandpolicyitem.md">OnDemandPolicyItem</a></i>
    <a href="#pitenabled" title="PitEnabled">PitEnabled</a>: <i>Boolean</i>
    <a href="#policyitemhourly" title="PolicyItemHourly">PolicyItemHourly</a>: <i><a href="scheduledpolicyitem.md">ScheduledPolicyItem</a></i>
    <a href="#policyitemdaily" title="PolicyItemDaily">PolicyItemDaily</a>: <i><a href="scheduledpolicyitem.md">ScheduledPolicyItem</a></i>
    <a href="#policyitemweekly" title="PolicyItemWeekly">PolicyItemWeekly</a>: <i>
      - <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a></i>
    <a href="#policyitemmonthly" title="PolicyItemMonthly">PolicyItemMonthly</a>: <i>
      - <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a></i>
    <a href="#policyitemyearly" title="PolicyItemYearly">PolicyItemYearly</a>: <i>
      - <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a></i>
</pre>

## Properties

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

The unique identifier of the project for the Backup Compliance Policy.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### AuthorizedEmail

Email address of the user authorized to update the Backup Compliance Policy settings.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AuthorizedUserFirstName

First name of the user authorized to update the Backup Compliance Policy settings.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AuthorizedUserLastName

Last name of the user authorized to update the Backup Compliance Policy settings.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CopyProtectionEnabled

Flag that indicates whether to enable additional copy protection for the cluster. If enabled, cloud backup snapshots cannot be deleted until the retention period expires. Defaults to false if not specified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EncryptionAtRestEnabled

Flag that indicates whether Encryption at Rest using Customer Key Management is required for all clusters with a Backup Compliance Policy. Defaults to false if not specified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RestoreWindowDays

Number of days back in time you can restore to with Continuous Cloud Backup accuracy. Must be a positive, non-zero integer. This field is optional and computed from the API.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OnDemandPolicyItem

On-demand backup policy item configuration. When provided, FrequencyInterval, RetentionUnit, and RetentionValue are required.

_Required_: No

_Type_: <a href="ondemandpolicyitem.md">OnDemandPolicyItem</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PitEnabled

Flag that indicates whether the cluster uses Continuous Cloud Backup. If enabled, cloud backup snapshots are taken every 6 hours. Defaults to false if not specified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PolicyItemHourly

Scheduled backup policy item configuration (hourly, daily, weekly, monthly, or yearly). When provided, FrequencyInterval, RetentionUnit, and RetentionValue are required.

_Required_: No

_Type_: <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PolicyItemDaily

_Required_: No

_Type_: <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PolicyItemWeekly

Weekly backup policy item configuration. Multiple weekly policy items are allowed.

_Required_: No

_Type_: List of <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PolicyItemMonthly

Monthly backup policy item configuration. Multiple monthly policy items are allowed.

_Required_: No

_Type_: List of <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PolicyItemYearly

Yearly backup policy item configuration. Multiple yearly policy items are allowed.

_Required_: No

_Type_: List of <a href="scheduledpolicyitem.md">ScheduledPolicyItem</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ProjectId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### State

Current state of the Backup Compliance Policy settings. Possible values are: ACTIVE, INACTIVE.

#### UpdatedDate

ISO 8601 timestamp in UTC of when the Backup Compliance Policy settings were last updated.

#### UpdatedUser

Email address of the user who last updated the Backup Compliance Policy settings.

#### OnDemandPolicyItem.Id

Unique identifier of the on-demand backup policy item.

#### OnDemandPolicyItem.FrequencyType

Frequency type of the on-demand backup policy item. Always returns "ondemand".

#### PolicyItemHourly.Id

Unique identifier of the hourly backup policy item.

#### PolicyItemHourly.FrequencyType

Frequency type of the hourly backup policy item. Always returns "hourly".

#### PolicyItemDaily.Id

Unique identifier of the daily backup policy item.

#### PolicyItemDaily.FrequencyType

Frequency type of the daily backup policy item. Always returns "daily".

#### PolicyItemWeekly.Id

Unique identifier of the weekly backup policy item(s). For arrays, use index notation (e.g., PolicyItemWeekly[0].Id).

#### PolicyItemWeekly.FrequencyType

Frequency type of the weekly backup policy item(s). Always returns "weekly". For arrays, use index notation (e.g., PolicyItemWeekly[0].FrequencyType).

#### PolicyItemMonthly.Id

Unique identifier of the monthly backup policy item(s). For arrays, use index notation (e.g., PolicyItemMonthly[0].Id).

#### PolicyItemMonthly.FrequencyType

Frequency type of the monthly backup policy item(s). Always returns "monthly". For arrays, use index notation (e.g., PolicyItemMonthly[0].FrequencyType).

#### PolicyItemYearly.Id

Unique identifier of the yearly backup policy item(s). For arrays, use index notation (e.g., PolicyItemYearly[0].Id).

#### PolicyItemYearly.FrequencyType

Frequency type of the yearly backup policy item(s). Always returns "yearly". For arrays, use index notation (e.g., PolicyItemYearly[0].FrequencyType).

