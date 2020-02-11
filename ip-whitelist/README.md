# MongoDB::Atlas::ProjectIPWhitelist

Congratulations on starting development!

Next steps:

1. Populate the JSON schema describing your resource, `mongodb-atlas-projectipwhitelist.json`
2. The RPDK will automatically generate the correct resource model from the
   schema whenever the project is built via Make.
   You can also do this manually with the following command: `cfn-cli generate`
3. Implement your resource handlers by adding code to provision your resources in your resource handler's methods.

Please don't modify files `model.go and main.go`, as they will be automatically overwritten.

## Attributes
`ProjectId` *(required)* : The unique identifier for the project to which you want to add one or more whitelist entries.<br>
`Id` : The unique identifier for the Project API Whitelist rules.<br>
`Whitelist.Comment` *(optional)* : Comment associated with the whitelist entry.<br>
`Whitelist.IpAddress` *(optional)* : Whitelisted IP address. Mutually exclusive with cidrBlock and awsSecurityGroup.<br>
`Whitelist.CidrBlock` *(optional)* : Whitelist entry in Classless Inter-Domain Routing (CIDR) notation. Mutually exclusive with ipAddress and awsSecurityGroup.<br>
`Whitelist.AwsSecurityGroup` *(optional)* : ID of the AWS security group to whitelist. Mutually exclusive with cidrBlock and ipAddress and cidrBlock.<br>
`TotalCount` : Count of the total number of items in the result set.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas.<br>
