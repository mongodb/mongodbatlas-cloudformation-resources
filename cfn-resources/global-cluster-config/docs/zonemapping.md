# MongoDB::Atlas::GlobalClusterConfig zoneMapping

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#location" title="Location">Location</a>" : <i>String</i>,
    "<a href="#zone" title="Zone">Zone</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#location" title="Location">Location</a>: <i>String</i>
<a href="#zone" title="Zone">Zone</a>: <i>String</i>
</pre>

## Properties

#### Location

Code that represents a location that maps to a zone in your global cluster. MongoDB Cloud represents this location with a ISO 3166-2 location and subdivision codes when possible.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Zone

Human-readable label that identifies the zone in your global cluster. This zone maps to a location code.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

