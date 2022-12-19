# MongoDB::Atlas::Auditing

Returns and edits database auditing settings for MongoDB Cloud projects.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::Auditing",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
        "<a href="#auditauthorizationsuccess" title="AuditAuthorizationSuccess">AuditAuthorizationSuccess</a>" : <i>Boolean</i>,
        "<a href="#auditfilter" title="AuditFilter">AuditFilter</a>" : <i>String</i>,
        "<a href="#configurationtype" title="ConfigurationType">ConfigurationType</a>" : <i>String</i>,
        "<a href="#groupid" title="GroupId">GroupId</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::Auditing
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    <a href="#auditauthorizationsuccess" title="AuditAuthorizationSuccess">AuditAuthorizationSuccess</a>: <i>Boolean</i>
    <a href="#auditfilter" title="AuditFilter">AuditFilter</a>: <i>String</i>
    <a href="#configurationtype" title="ConfigurationType">ConfigurationType</a>: <i>String</i>
    <a href="#groupid" title="GroupId">GroupId</a>: <i>String</i>
</pre>

## Properties

#### ApiKeys

_Required_: Yes

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AuditAuthorizationSuccess

Flag that indicates whether someone set auditing to track successful authentications. This only applies to the `"atype" : "authCheck"` audit filter. Setting this parameter to `true` degrades cluster performance.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AuditFilter

JSON document that specifies which events to record. Escape any characters that may prevent parsing, such as single or double quotes, using a backslash (`\`), for more information about audit filters refer to https://www.mongodb.com/docs/manual/tutorial/configure-audit-filters/.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ConfigurationType

Human-readable label that displays how to configure the audit filter.

_Required_: No

_Type_: String

_Allowed Values_: <code>FILTER_BUILDER</code> | <code>FILTER_JSON</code> | <code>NONE</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### GroupId

Unique 24-hexadecimal digit string that identifies your project.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the GroupId.
