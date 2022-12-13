# MongoDB::Atlas::X509AuthenticationDatabaseUser Certificate

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#username" title="UserName">UserName</a>" : <i>String</i>,
    "<a href="#createdat" title="CreatedAt">CreatedAt</a>" : <i>String</i>,
    "<a href="#monthsuntilexpiration" title="MonthsUntilExpiration">MonthsUntilExpiration</a>" : <i>Integer</i>,
    "<a href="#notafter" title="NotAfter">NotAfter</a>" : <i>String</i>,
    "<a href="#subject" title="Subject">Subject</a>" : <i>String</i>,
    "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>,
    "<a href="#id" title="Id">Id</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#username" title="UserName">UserName</a>: <i>String</i>
<a href="#createdat" title="CreatedAt">CreatedAt</a>: <i>String</i>
<a href="#monthsuntilexpiration" title="MonthsUntilExpiration">MonthsUntilExpiration</a>: <i>Integer</i>
<a href="#notafter" title="NotAfter">NotAfter</a>: <i>String</i>
<a href="#subject" title="Subject">Subject</a>: <i>String</i>
<a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
<a href="#id" title="Id">Id</a>: <i>String</i>
</pre>

## Properties

#### UserName

 Username of the database user to create a certificate for.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CreatedAt

Timestamp in ISO 8601 date and time format in UTC when Atlas created this X.509 certificate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### MonthsUntilExpiration

A number of months that the created certificate is valid for before expiry, up to 24 months.default 3.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### NotAfter

Timestamp in ISO 8601 date and time format in UTC when this certificate expires.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Subject

Fully distinguished name of the database user to which this certificate belongs. To learn more, see RFC 2253.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### GroupId

Unique identifier of the Atlas project to which this certificate belongs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Id

Serial number of this certificate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

