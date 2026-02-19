# MongoDB::Atlas::ServiceAccountSecret

Creates a secret for the specified Service Account at the organization level.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::ServiceAccountSecret",
    "Properties" : {
        "<a href="#orgid" title="OrgId">OrgId</a>" : <i>String</i>,
        "<a href="#clientid" title="ClientId">ClientId</a>" : <i>String</i>,
        "<a href="#secretexpiresafterhours" title="SecretExpiresAfterHours">SecretExpiresAfterHours</a>" : <i>Integer</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::ServiceAccountSecret
Properties:
    <a href="#orgid" title="OrgId">OrgId</a>: <i>String</i>
    <a href="#clientid" title="ClientId">ClientId</a>: <i>String</i>
    <a href="#secretexpiresafterhours" title="SecretExpiresAfterHours">SecretExpiresAfterHours</a>: <i>Integer</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
</pre>

## Properties

#### OrgId

Unique 24-hexadecimal digit string that identifies the organization that contains your projects.

_Required_: Yes

_Type_: String

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ClientId

The Client ID of the Service Account.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### SecretExpiresAfterHours

The expiration time of the new Service Account secret, provided in hours. The minimum and maximum allowed expiration times are subject to change and are controlled by the organization's settings.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### SecretId

Unique 24-hexadecimal digit string that identifies the secret.

#### Secret

The secret value for the Service Account. It will be returned only the first time after creation.

#### MaskedSecretValue

The masked Service Account secret.

#### CreatedAt

The date that the secret was created on. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### ExpiresAt

The date for the expiration of the secret. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

#### LastUsedAt

The last time the secret was used. This parameter expresses its value in the ISO 8601 timestamp format in UTC.

