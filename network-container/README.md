# MongoDB::Atlas::NetworkContainer

Congratulations on starting development!

Next steps:

1. Populate the JSON schema describing your resource, `mongodb-atlas-networkcontainer.json`
2. The RPDK will automatically generate the correct resource model from the
   schema whenever the project is built via Make.
   You can also do this manually with the following command: `cfn-cli generate`
3. Implement your resource handlers by adding code to provision your resources in your resource handler's methods.

Please don't modify files `model.go` and `main.go`, as they will be automatically overwritten.

## Description
This resource allows you to create, read, update and delete a network container

## Attributes
`Id` : Unique identifier of the Network Peer.<br>
`Provisioned` : Flag that indicates if the project has clusters deployed in the Network Peering container.<br>
`VpcId` : Unique identifier of the peer VPC.<br>

## Parameters
`ProjectId` *(required)* : The unique identifier of the project.<br>
`RegionName` *(optional)* : Name of region<br>
`ProviderName` *(optional)* : The name of the provider.<br>
`AtlasCidrBlock` *(required)* : CIDR block that Atlas uses for your clusters.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit
    ...