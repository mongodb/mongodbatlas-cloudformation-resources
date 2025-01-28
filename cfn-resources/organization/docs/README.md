# MongoDB::Atlas::Organization

Returns, adds, and edits organizational units in MongoDB Cloud.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::Organization",
    "Properties" : {
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#apikey" title="APIKey">APIKey</a>" : <i><a href="apikey.md">APIKey</a></i>,
        "<a href="#federatedsettingsid" title="FederatedSettingsId">FederatedSettingsId</a>" : <i>String</i>,
        "<a href="#orgownerid" title="OrgOwnerId">OrgOwnerId</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#awssecretname" title="AwsSecretName">AwsSecretName</a>" : <i>String</i>,
        "<a href="#skipdefaultalertssettings" title="SkipDefaultAlertsSettings">SkipDefaultAlertsSettings</a>" : <i>Boolean</i>,
        "<a href="#isdeleted" title="IsDeleted">IsDeleted</a>" : <i>Boolean</i>,
        "<a href="#apiaccesslistrequired" title="ApiAccessListRequired">ApiAccessListRequired</a>" : <i>Boolean</i>,
        "<a href="#multifactorauthrequired" title="MultiFactorAuthRequired">MultiFactorAuthRequired</a>" : <i>Boolean</i>,
        "<a href="#restrictemployeeaccess" title="RestrictEmployeeAccess">RestrictEmployeeAccess</a>" : <i>Boolean</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::Organization
Properties:
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#apikey" title="APIKey">APIKey</a>: <i><a href="apikey.md">APIKey</a></i>
    <a href="#federatedsettingsid" title="FederatedSettingsId">FederatedSettingsId</a>: <i>String</i>
    <a href="#orgownerid" title="OrgOwnerId">OrgOwnerId</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#awssecretname" title="AwsSecretName">AwsSecretName</a>: <i>String</i>
    <a href="#skipdefaultalertssettings" title="SkipDefaultAlertsSettings">SkipDefaultAlertsSettings</a>: <i>Boolean</i>
    <a href="#isdeleted" title="IsDeleted">IsDeleted</a>: <i>Boolean</i>
    <a href="#apiaccesslistrequired" title="ApiAccessListRequired">ApiAccessListRequired</a>: <i>Boolean</i>
    <a href="#multifactorauthrequired" title="MultiFactorAuthRequired">MultiFactorAuthRequired</a>: <i>Boolean</i>
    <a href="#restrictemployeeaccess" title="RestrictEmployeeAccess">RestrictEmployeeAccess</a>: <i>Boolean</i>
</pre>

## Properties

#### Name

Human-readable label that identifies the organization.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### APIKey

_Required_: No

_Type_: <a href="apikey.md">APIKey</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FederatedSettingsId

Unique 24-hexadecimal digit string that identifies the federation to link the newly created organization to. If specified, the proposed Organization Owner of the new organization must have the Organization Owner role in an organization associated with the federation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OrgOwnerId

Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user that you want to assign the Organization Owner role. This user must be a member of the same organization as the calling API key. If you provide federationSettingsId, this user must instead have the Organization Owner role on an organization in the specified federation. This parameter is required only when you authenticate with Programmatic API Keys.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### AwsSecretName

AwsSecretName used to set newly created Org credentials information.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### SkipDefaultAlertsSettings

Disables automatic alert creation. When set to `true`, Atlas doesn't automatically create organization-level alerts. Defaults to `true` for new Atlas Organizations created with the provider to prevent infrastructure drift caused by creation of new alerts.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IsDeleted

Flag that indicates whether this organization has been deleted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ApiAccessListRequired

Flag that indicates whether to require API operations to originate from an IP Address added to the API access list for the specified organization.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MultiFactorAuthRequired

Flag that indicates whether to require users to set up Multi-Factor Authentication (MFA) before accessing the specified organization. To learn more, see: https://www.mongodb.com/docs/atlas/security-multi-factor-authentication/.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### RestrictEmployeeAccess

Flag that indicates whether to block MongoDB Support from accessing Atlas infrastructure for any deployment in the specified organization without explicit permission. Once this setting is turned on, you can grant MongoDB Support a 24-hour bypass access to the Atlas deployment to resolve support issues. To learn more, see: https://www.mongodb.com/docs/atlas/security-restrict-support-access/.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### OrgId

Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the /orgs endpoint to retrieve all organizations to which the authenticated user has access.

