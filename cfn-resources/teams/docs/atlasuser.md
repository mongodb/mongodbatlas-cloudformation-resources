# MongoDB::Atlas::Teams AtlasUser

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#country" title="Country">Country</a>" : <i>String</i>,
    "<a href="#emailaddress" title="EmailAddress">EmailAddress</a>" : <i>String</i>,
    "<a href="#firstname" title="FirstName">FirstName</a>" : <i>String</i>,
    "<a href="#id" title="Id">Id</a>" : <i>String</i>,
    "<a href="#lastname" title="LastName">LastName</a>" : <i>String</i>,
    "<a href="#links" title="Links">Links</a>" : <i>[ <a href="link.md">Link</a>, ... ]</i>,
    "<a href="#mobilenumber" title="MobileNumber">MobileNumber</a>" : <i>String</i>,
    "<a href="#password" title="Password">Password</a>" : <i>String</i>,
    "<a href="#roles" title="Roles">Roles</a>" : <i>[ <a href="atlasrole.md">AtlasRole</a>, ... ]</i>,
    "<a href="#teamids" title="TeamIds">TeamIds</a>" : <i>[ String, ... ]</i>,
    "<a href="#username" title="Username">Username</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#country" title="Country">Country</a>: <i>String</i>
<a href="#emailaddress" title="EmailAddress">EmailAddress</a>: <i>String</i>
<a href="#firstname" title="FirstName">FirstName</a>: <i>String</i>
<a href="#id" title="Id">Id</a>: <i>String</i>
<a href="#lastname" title="LastName">LastName</a>: <i>String</i>
<a href="#links" title="Links">Links</a>: <i>
      - <a href="link.md">Link</a></i>
<a href="#mobilenumber" title="MobileNumber">MobileNumber</a>: <i>String</i>
<a href="#password" title="Password">Password</a>: <i>String</i>
<a href="#roles" title="Roles">Roles</a>: <i>
      - <a href="atlasrole.md">AtlasRole</a></i>
<a href="#teamids" title="TeamIds">TeamIds</a>: <i>
      - String</i>
<a href="#username" title="Username">Username</a>: <i>String</i>
</pre>

## Properties

#### Country

Two alphabet characters that identifies MongoDB Cloud user's geographic location. This parameter uses the ISO 3166-1a2 code format.

_Required_: No

_Type_: String

_Pattern_: <code>^([A-Z]{2})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### EmailAddress

Email address that belongs to the MongoDB Cloud user.

_Required_: No

_Type_: String

_Pattern_: <code>^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### FirstName

First or given name that belongs to the MongoDB Cloud user.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Id

Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user.

_Required_: No

_Type_: String

_Minimum_: <code>24</code>

_Maximum_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### LastName

Last name, family name, or surname that belongs to the MongoDB Cloud user.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Links

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.

_Required_: No

_Type_: List of <a href="link.md">Link</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MobileNumber

Mobile phone number that belongs to the MongoDB Cloud user.

_Required_: No

_Type_: String

_Pattern_: <code>(?:(?:\\+?1\\s*(?:[.-]\\s*)?)?(?:(\\s*([2-9]1[02-9]|[2-9][02-8]1|[2-9][02-8][02-9])\\s*)|([2-9]1[02-9]|[2-9][02-8]1|[2-9][02-8][02-9]))\\s*(?:[.-]\\s*)?)([2-9]1[02-9]|[2-9][02-9]1|[2-9][02-9]{2})\\s*(?:[.-]\\s*)?([0-9]{4})$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Password

Password applied with the username to log in to MongoDB Cloud. MongoDB Cloud does not return this parameter except in response to creating a new MongoDB Cloud user. Only the MongoDB Cloud user can update their password after it has been set from the MongoDB Cloud console.

_Required_: No

_Type_: String

_Minimum_: <code>8</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Roles

List of objects that display the MongoDB Cloud user's roles and the corresponding organization or project to which that role applies. A role can apply to one organization or one project but not both.

_Required_: No

_Type_: List of <a href="atlasrole.md">AtlasRole</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TeamIds

List of unique 24-hexadecimal digit strings that identifies the teams to which this MongoDB Cloud user belongs.

_Required_: No

_Type_: List of String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Username

Email address that represents the username of the MongoDB Cloud user.

_Required_: No

_Type_: String

_Pattern_: <code>^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

