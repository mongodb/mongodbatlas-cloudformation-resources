# MongoDB::Atlas::X509AuthenticationDatabaseUser

Generates one X.509 certificate for the specified MongoDB user. Atlas manages the certificate and MongoDB user that belong to one project. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

To get MongoDB Cloud to generate a managed certificate for a database user, set `"x509Type" : "MANAGED"` on the desired MongoDB Database User.

If you are managing your own Certificate Authority (CA) in Self-Managed X.509 mode, you must generate certificates for database users using your own CA.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::X509AuthenticationDatabaseUser",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#customerx509" title="CustomerX509">CustomerX509</a>" : <i><a href="customerx509.md">customerX509</a></i>,
        "<a href="#username" title="UserName">UserName</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::X509AuthenticationDatabaseUser
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#customerx509" title="CustomerX509">CustomerX509</a>: <i><a href="customerx509.md">customerX509</a></i>
    <a href="#username" title="UserName">UserName</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
</pre>

## Properties

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CustomerX509

_Required_: No

_Type_: <a href="customerx509.md">customerX509</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### UserName

Username of the database user to create a certificate for.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

The unique identifier for the project .

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Links

One or more links to sub-resources and/or related resources.

#### Results

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.

#### TotalCount

Total number of unexpired certificates returned in this response.

#### MonthsUntilExpiration

A number of months that the created certificate is valid for before expiry, up to 24 months.default 3.

#### CustomerX509

Returns the <code>CustomerX509</code> value.

