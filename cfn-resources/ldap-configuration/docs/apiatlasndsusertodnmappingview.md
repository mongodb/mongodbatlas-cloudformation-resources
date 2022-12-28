# MongoDB::Atlas::LDAPConfiguration ApiAtlasNDSUserToDNMappingView

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#ldapquery" title="LdapQuery">LdapQuery</a>" : <i>String</i>,
    "<a href="#match" title="Match">Match</a>" : <i>String</i>,
    "<a href="#substitution" title="Substitution">Substitution</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#ldapquery" title="LdapQuery">LdapQuery</a>: <i>String</i>
<a href="#match" title="Match">Match</a>: <i>String</i>
<a href="#substitution" title="Substitution">Substitution</a>: <i>String</i>
</pre>

## Properties

#### LdapQuery

Lightweight Directory Access Protocol (LDAP) query template that inserts the LDAP name that the regular expression matches into an LDAP query Uniform Resource Identifier (URI). The formatting for the query must conform to [RFC 4515](https://datatracker.ietf.org/doc/html/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Match

Regular expression that MongoDB Cloud uses to match against the provided Lightweight Directory Access Protocol (LDAP) username. Each parenthesis-enclosed section represents a regular expression capture group that the substitution or `ldapQuery` template uses.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Substitution

Lightweight Directory Access Protocol (LDAP) Distinguished Name (DN) template that converts the LDAP username that matches regular expression in the *match* parameter into an LDAP Distinguished Name (DN).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

