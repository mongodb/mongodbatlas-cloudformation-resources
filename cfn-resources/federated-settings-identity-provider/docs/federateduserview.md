# MongoDB::Atlas::federatedsettingsidentityprovider FederatedUserView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">ApiKeyDefinition</a></i>,
    "<a href="#emailaddress" title="EmailAddress">EmailAddress</a>" : <i>String</i>,
    "<a href="#federationsettingsid" title="FederationSettingsId">FederationSettingsId</a>" : <i>String</i>,
    "<a href="#firstname" title="FirstName">FirstName</a>" : <i>String</i>,
    "<a href="#lastname" title="LastName">LastName</a>" : <i>String</i>,
    "<a href="#userid" title="UserId">UserId</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">ApiKeyDefinition</a></i>
<a href="#emailaddress" title="EmailAddress">EmailAddress</a>: <i>String</i>
<a href="#federationsettingsid" title="FederationSettingsId">FederationSettingsId</a>: <i>String</i>
<a href="#firstname" title="FirstName">FirstName</a>: <i>String</i>
<a href="#lastname" title="LastName">LastName</a>: <i>String</i>
<a href="#userid" title="UserId">UserId</a>: <i>String</i>
</pre>

## Properties

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">ApiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EmailAddress

Email address of the MongoDB Cloud user linked to the federated organization.

_Required_: No

_Type_: String

_Pattern_: <code>^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FederationSettingsId

Unique 24-hexadecimal digit string that identifies the federation to which this MongoDB Cloud user belongs.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FirstName

First or given name that belongs to the MongoDB Cloud user.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### LastName

Last name, family name, or surname that belongs to the MongoDB Cloud user.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### UserId

Unique 24-hexadecimal digit string that identifies this user.

_Required_: No

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

