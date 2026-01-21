# MongoDB::Atlas::DatabaseUser

Returns, adds, edits, and removes database users.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::DatabaseUser",
    "Properties" : {
        "<a href="#deleteafterdate" title="DeleteAfterDate">DeleteAfterDate</a>" : <i>String</i>,
        "<a href="#awsiamtype" title="AWSIAMType">AWSIAMType</a>" : <i>String</i>,
        "<a href="#databasename" title="DatabaseName">DatabaseName</a>" : <i>String</i>,
        "<a href="#description" title="Description">Description</a>" : <i>String</i>,
        "<a href="#labels" title="Labels">Labels</a>" : <i>[ <a href="labeldefinition.md">labelDefinition</a>, ... ]</i>,
        "<a href="#ldapauthtype" title="LdapAuthType">LdapAuthType</a>" : <i>String</i>,
        "<a href="#x509type" title="X509Type">X509Type</a>" : <i>String</i>,
        "<a href="#oidcauthtype" title="OIDCAuthType">OIDCAuthType</a>" : <i>String</i>,
        "<a href="#password" title="Password">Password</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#roles" title="Roles">Roles</a>" : <i>[ <a href="roledefinition.md">roleDefinition</a>, ... ]</i>,
        "<a href="#scopes" title="Scopes">Scopes</a>" : <i>[ <a href="scopedefinition.md">scopeDefinition</a>, ... ]</i>,
        "<a href="#username" title="Username">Username</a>" : <i>String</i>,
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::DatabaseUser
Properties:
    <a href="#deleteafterdate" title="DeleteAfterDate">DeleteAfterDate</a>: <i>String</i>
    <a href="#awsiamtype" title="AWSIAMType">AWSIAMType</a>: <i>String</i>
    <a href="#databasename" title="DatabaseName">DatabaseName</a>: <i>String</i>
    <a href="#description" title="Description">Description</a>: <i>String</i>
    <a href="#labels" title="Labels">Labels</a>: <i>
      - <a href="labeldefinition.md">labelDefinition</a></i>
    <a href="#ldapauthtype" title="LdapAuthType">LdapAuthType</a>: <i>String</i>
    <a href="#x509type" title="X509Type">X509Type</a>: <i>String</i>
    <a href="#oidcauthtype" title="OIDCAuthType">OIDCAuthType</a>: <i>String</i>
    <a href="#password" title="Password">Password</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#roles" title="Roles">Roles</a>: <i>
      - <a href="roledefinition.md">roleDefinition</a></i>
    <a href="#scopes" title="Scopes">Scopes</a>: <i>
      - <a href="scopedefinition.md">scopeDefinition</a></i>
    <a href="#username" title="Username">Username</a>: <i>String</i>
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
</pre>

## Properties

#### DeleteAfterDate

Date and time when MongoDB Cloud deletes the user. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AWSIAMType

Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user's role. Default value is `NONE`.

_Required_: No

_Type_: String

_Allowed Values_: <code>NONE</code> | <code>USER</code> | <code>ROLE</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DatabaseName

MongoDB database against which the MongoDB database user authenticates. MongoDB database users must provide both a username and authentication database to log into MongoDB.  Default value is `admin`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Description

Description of this database user.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Labels

List that contains the key-value pairs for tagging and categorizing the MongoDB database user. The labels that you define do not appear in the console.

_Required_: No

_Type_: List of <a href="labeldefinition.md">labelDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### LdapAuthType

Method by which the provided username is authenticated. Default value is `NONE`.

_Required_: No

_Type_: String

_Allowed Values_: <code>NONE</code> | <code>USER</code> | <code>GROUP</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### X509Type

Method that briefs who owns the certificate provided. Default value is `NONE`.

_Required_: No

_Type_: String

_Allowed Values_: <code>NONE</code> | <code>MANAGED</code> | <code>CUSTOMER</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### OIDCAuthType

Human-readable label that indicates whether the new database user or group authenticates with OIDC federated authentication. To create a federated authentication user, specify the value of USER in this field. To create a federated authentication group, specify the value of IDP_GROUP in this field. Default value is `NONE`.

_Required_: No

_Type_: String

_Allowed Values_: <code>NONE</code> | <code>USER</code> | <code>IDP_GROUP</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Password

The user’s password. This field is not included in the entity returned from the server.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

Unique 24-hexadecimal digit string that identifies your Atlas Project.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Roles

List that provides the pairings of one role with one applicable database.

_Required_: Yes

_Type_: List of <a href="roledefinition.md">roleDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Scopes

List that contains clusters and MongoDB Atlas Data Federation that this database user can access. If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Federation in the project.

_Required_: No

_Type_: List of <a href="scopedefinition.md">scopeDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Username

Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication. This will be USER_ARN or ROLE_ARN if AWSIAMType is USER or ROLE. Refer https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Database-Users/operation/createDatabaseUser for details.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided `default` is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### UserCFNIdentifier

A unique identifier comprised of the Atlas Project ID and Username.

