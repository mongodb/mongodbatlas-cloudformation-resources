# MongoDB::Atlas::NetworkPeering

## Description
This resource allows you to create, read, update and delete a network peering.

## Attributes
`Id` : Unique identifier of the Network Peer.<br>
`ConnectionId` : Unique identifier for the peering connection.<br>
`ErrorStateName` : Error state, if any.<br>
`StatusNames` : The VPC peering connection status<br>

## Parameters
`ProjectId` *(required)* : The unique identifier of the project.<br>
`ContainerId` *(required)* : Unique identifier of the Atlas VPC container for the AWS region.<br>
`AccepterRegionName` *(optional)* : AWS region where the peer VPC resides. Returns null if the region is the same region in which the Atlas VPC resides.<br>
`AwsAccountId` *(optional)* : AWS account ID of the owner of the peer VPC.<br>
`ProviderName` *(required)* : The name of the provider.<br>
`RouteTableCidrBlock` *(optional)* : Peer VPC CIDR block or subnet.<br>
`VpcId` *(optional)* : Unique identifier of the peer VPC.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit
    ...