# MongoDB::Atlas::GlobalClusterConfig

Returns, adds, and removes Global Cluster managed namespaces and custom zone mappings. This resource can only be used with Atlas-managed clusters, see doc for `GlobalClusterSelfManagedSharding` attribute in `Mongodb::Atlas::Cluster` resource for more info.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::GlobalClusterConfig",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
        "<a href="#managednamespaces" title="ManagedNamespaces">ManagedNamespaces</a>" : <i>[ <a href="managednamespace.md">managedNamespace</a>, ... ]</i>,
        "<a href="#customzonemappings" title="CustomZoneMappings">CustomZoneMappings</a>" : <i>[ <a href="zonemapping.md">zoneMapping</a>, ... ]</i>
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::GlobalClusterConfig
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
    <a href="#managednamespaces" title="ManagedNamespaces">ManagedNamespaces</a>: <i>
      - <a href="managednamespace.md">managedNamespace</a></i>
    <a href="#customzonemappings" title="CustomZoneMappings">CustomZoneMappings</a>: <i>
      - <a href="zonemapping.md">zoneMapping</a></i>
</pre>

## Properties

#### Profile

The profile is defined in AWS Secret manager. See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ProjectId

The unique identifier of the project for the Atlas cluster.

_Required_: No

_Type_: String

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

#### CustomZoneMappings

List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to the human-readable label for the desired custom zone. MongoDB Cloud maps the ISO 3166-1a2 code to the nearest geographical zone by default. Include this parameter to override the default mappings.

This parameter returns an empty object if no custom zones exist.

_Required_: No

_Type_: List of <a href="zonemapping.md">zoneMapping</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### RemoveAllZoneMapping

Flag that indicates whether all custom zone mapping to be deleted during delete.

