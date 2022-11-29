# MongoDB::Atlas::GlobalClusterConfig

Returns, adds, and removes Global Cluster managed namespaces and custom zone mappings.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::GlobalClusterConfig",
    "Properties" : {
        "<a href="#apikeys" title="ApiKeys">ApiKeys</a>" : <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>,
        "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
        "<a href="#managednamespaces" title="ManagedNamespaces">ManagedNamespaces</a>" : <i>[ <a href="managednamespace.md">managedNamespace</a>, ... ]</i>,
        "<a href="#id" title="Id">Id</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::GlobalClusterConfig
Properties:
    <a href="#apikeys" title="ApiKeys">ApiKeys</a>: <i><a href="apikeydefinition.md">apiKeyDefinition</a></i>
    <a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
    <a href="#managednamespaces" title="ManagedNamespaces">ManagedNamespaces</a>: <i>
      - <a href="managednamespace.md">managedNamespace</a></i>
    <a href="#id" title="Id">Id</a>: <i>String</i>
</pre>

## Properties

#### ApiKeys

_Required_: No

_Type_: <a href="apikeydefinition.md">apiKeyDefinition</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ClusterName

The name of the Atlas cluster that contains the snapshots you want to retrieve.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ManagedNamespaces

List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to the human-readable label for the desired custom zone. MongoDB Cloud maps the ISO 3166-1a2 code to the nearest geographical zone by default. Include this parameter to override the default mappings.

This parameter returns an empty object if no custom zones exist.

_Required_: No

_Type_: List of <a href="managednamespace.md">managedNamespace</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Id

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ProjectId.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### RemoveAllZoneMapping

Flag that indicates whether all custom zone mapping to be deleted during delete.

#### CustomShardKey

Returns the <code>CustomShardKey</code> value.

#### ProjectId

The unique identifier of the project for the Atlas cluster.

#### CustomZoneMappings

List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to the human-readable label for the desired custom zone. MongoDB Cloud maps the ISO 3166-1a2 code to the nearest geographical zone by default. Include this parameter to override the default mappings.

This parameter returns an empty object if no custom zones exist.

