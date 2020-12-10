# MongoDB::Atlas::ProjectIPWhitelist

## Description
This resource allows you to create, read, update and delete an IP whitelist. Atlas only allows client connections to the cluster from entries in the project’s whitelist. Each entry is either a single IP address, a CIDR-notated range of addresses, or an AWS security group.

## Attributes
`Id` : The unique identifier for the Project API Whitelist rules.<br>
`TotalCount` : Count of the total number of items in the result set.<br>

## Parameters
`ProjectId` *(required)* : The unique identifier for the project to which you want to add one or more whitelist entries.<br>
`Whitelist.Comment` *(optional)* : Comment associated with the whitelist entry.<br>
`Whitelist.IpAddress` *(optional)* : Whitelisted IP address. Mutually exclusive with cidrBlock and awsSecurityGroup.<br>
`Whitelist.CidrBlock` *(optional)* : Whitelist entry in Classless Inter-Domain Routing (CIDR) notation. Mutually exclusive with ipAddress and awsSecurityGroup.<br>
`Whitelist.AwsSecurityGroup` *(optional)* : ID of the AWS security group to whitelist. Mutually exclusive with cidrBlock and ipAddress and cidrBlock.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit
    ...