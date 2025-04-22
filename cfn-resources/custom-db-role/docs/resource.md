# MongoDB::Atlas::CustomDBRole Resource

List of resources on which you grant the action.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#collection" title="Collection">Collection</a>" : <i>String</i>,
    "<a href="#db" title="DB">DB</a>" : <i>String</i>,
    "<a href="#cluster" title="Cluster">Cluster</a>" : <i>Boolean</i>
}
</pre>

### YAML

<pre>
<a href="#collection" title="Collection">Collection</a>: <i>String</i>
<a href="#db" title="DB">DB</a>: <i>String</i>
<a href="#cluster" title="Cluster">Cluster</a>: <i>Boolean</i>
</pre>

## Properties

#### Collection

Human-readable label that identifies the collection on which you grant the action to one MongoDB user. If you don't set this parameter, you grant the action to all collections in the database specified in the actions.resources.db parameter. If you set "actions.resources.cluster" : true, MongoDB Cloud ignores this parameter. Use the empty string ("") to allow an action on all collections.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DB

Human-readable label that identifies the database on which you grant the action to one MongoDB user. If you set "actions.resources.cluster" : true, MongoDB Cloud ignores this parameter. Use the empty string ("") to allow an action on all databases.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Cluster

Flag that indicates whether to grant the action on the cluster resource. If true, MongoDB Cloud ignores the actions.resources.collection and actions.resources.db parameters.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

