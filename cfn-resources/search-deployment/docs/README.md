# MongoDB::Atlas::SearchDeployment

The resource lets you create, edit and delete dedicated search nodes in a cluster. For details on supported cloud providers and existing limitations you can visit the Search Node Documentation: https://www.mongodb.com/docs/atlas/cluster-config/multi-cloud-distribution/#search-nodes-for-workload-isolation. Only a single search deployment resource can be defined for each cluster.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "MongoDB::Atlas::SearchDeployment",
    "Properties" : {
        "<a href="#profile" title="Profile">Profile</a>" : <i>String</i>,
        "<a href="#clustername" title="ClusterName">ClusterName</a>" : <i>String</i>,
        "<a href="#projectid" title="ProjectId">ProjectId</a>" : <i>String</i>,
        "<a href="#specs" title="Specs">Specs</a>" : <i>[ <a href="apisearchdeploymentspec.md">ApiSearchDeploymentSpec</a>, ... ]</i>,
    }
}
</pre>

### YAML

<pre>
Type: MongoDB::Atlas::SearchDeployment
Properties:
    <a href="#profile" title="Profile">Profile</a>: <i>String</i>
    <a href="#clustername" title="ClusterName">ClusterName</a>: <i>String</i>
    <a href="#projectid" title="ProjectId">ProjectId</a>: <i>String</i>
    <a href="#specs" title="Specs">Specs</a>: <i>
      - <a href="apisearchdeploymentspec.md">ApiSearchDeploymentSpec</a></i>
</pre>

## Properties

#### Profile

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used

_Required_: No

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ClusterName

Label that identifies the cluster to return the search nodes for.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### ProjectId

Unique 24-hexadecimal character string that identifies the project.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>24</code>

_Maximum Length_: <code>24</code>

_Pattern_: <code>^([a-f0-9]{24})$</code>

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

#### Specs

List of settings that configure the search nodes for your cluster. This list is currently limited to defining a single element.

_Required_: Yes

_Type_: List of <a href="apisearchdeploymentspec.md">ApiSearchDeploymentSpec</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Id

Unique 24-hexadecimal digit string that identifies the search deployment.

#### StateName

Human-readable label that indicates the current operating condition of this search deployment.

#### EncryptionAtRestProvider

Cloud service provider that manages your customer keys to provide an additional layer of Encryption At Rest for the cluster.

