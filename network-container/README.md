# MongoDB::Atlas::NetworkContainer

## Description
This resource allows you to create, read, update and delete a network container. A network container must be created before one can enable network peering. You need one container per AWS region that will be peered with.
## Attributes
`Id` : Unique identifier of the Network Peer.<br>
`Provisioned` : Flag that indicates if the project has clusters deployed in the Network Peering container.<br>
`VpcId` : Unique identifier of the peer VPC.<br>

## Parameters
`ProjectId` *(required)* : The unique identifier of the project.<br>
`RegionName` *(optional)* : Name of region<br>
`ProviderName` *(optional)* : The name of the provider. Default is `AWS`<br>
`AtlasCidrBlock` *(required)* : CIDR block that Atlas uses for your clusters.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit
    ...