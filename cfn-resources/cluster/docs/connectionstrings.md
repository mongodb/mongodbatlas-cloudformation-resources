# MongoDB::Atlas::Cluster connectionStrings

Collection of Uniform Resource Locators that point to the MongoDB database.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#standard" title="Standard">Standard</a>" : <i>String</i>,
    "<a href="#standardsrv" title="StandardSrv">StandardSrv</a>" : <i>String</i>,
    "<a href="#private" title="Private">Private</a>" : <i>String</i>,
    "<a href="#privatesrv" title="PrivateSrv">PrivateSrv</a>" : <i>String</i>,
    "<a href="#privateendpoint" title="PrivateEndpoint">PrivateEndpoint</a>" : <i>[ <a href="privateendpoint.md">privateEndpoint</a>, ... ]</i>,
    "<a href="#awsprivatelinksrv" title="AwsPrivateLinkSrv">AwsPrivateLinkSrv</a>" : <i>String</i>,
    "<a href="#awsprivatelink" title="AwsPrivateLink">AwsPrivateLink</a>" : <i>String</i>
}
</pre>

### YAML

<pre>
<a href="#standard" title="Standard">Standard</a>: <i>String</i>
<a href="#standardsrv" title="StandardSrv">StandardSrv</a>: <i>String</i>
<a href="#private" title="Private">Private</a>: <i>String</i>
<a href="#privatesrv" title="PrivateSrv">PrivateSrv</a>: <i>String</i>
<a href="#privateendpoint" title="PrivateEndpoint">PrivateEndpoint</a>: <i>
      - <a href="privateendpoint.md">privateEndpoint</a></i>
<a href="#awsprivatelinksrv" title="AwsPrivateLinkSrv">AwsPrivateLinkSrv</a>: <i>String</i>
<a href="#awsprivatelink" title="AwsPrivateLink">AwsPrivateLink</a>: <i>String</i>
</pre>

## Properties

#### Standard

Public connection string that you can use to connect to this cluster. This connection string uses the mongodb:// protocol.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### StandardSrv

Public connection string that you can use to connect to this cluster. This connection string uses the mongodb+srv:// protocol.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Private

Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster. This connection string uses the mongodb+srv:// protocol. The resource returns this parameter once someone creates a network peering connection to this cluster. This protocol tells the application to look up the host seed list in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn't, use connectionStrings.private. For Amazon Web Services (AWS) clusters, this resource returns this parameter only if you enable custom DNS.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PrivateSrv

Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster. This connection string uses the mongodb+srv:// protocol. The resource returns this parameter when someone creates a network peering connection to this cluster. This protocol tells the application to look up the host seed list in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the Uniform Resource Identifier (URI) if the nodes change. Use this Uniform Resource Identifier (URI) format if your driver supports it. If it doesn't, use connectionStrings.private. For Amazon Web Services (AWS) clusters, this parameter returns only if you enable custom DNS.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### PrivateEndpoint

List of private endpoint connection strings that you can use to connect to this cluster through a private endpoint. This parameter returns only if you deployed a private endpoint to all regions to which you deployed this clusters' nodes.

_Required_: No

_Type_: List of <a href="privateendpoint.md">privateEndpoint</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AwsPrivateLinkSrv

Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink. Each key identifies an Amazon Web Services (AWS) interface endpoint. Each value identifies the related mongodb:// connection string that you use to connect to Atlas through the interface endpoint that the key names.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### AwsPrivateLink

Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink. Each key identifies an Amazon Web Services (AWS) interface endpoint. Each value identifies the related mongodb:// connection string that you use to connect to MongoDB Cloud through the interface endpoint that the key names.

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

