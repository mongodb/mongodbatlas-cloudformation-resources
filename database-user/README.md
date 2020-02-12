# MongoDB::Atlas::DatabaseUser

Congratulations on starting development!

Next steps:

1. Populate the JSON schema describing your resource, `mongodb-atlas-databaseuser.json`
2. The RPDK will automatically generate the correct resource model from the
   schema whenever the project is built via Make.
   You can also do this manually with the following command: `cfn-cli generate`
3. Implement your resource handlers by adding code to provision your resources in your resource handler's methods.

Please don't modify files `model.go and main.go`, as they will be automatically overwritten.

## Description
The databaseUsers resource lets you retrieve, create and modify the MongoDB users in your cluster. Each user has a set of roles that provide access to the project’s databases. A user’s roles apply to all the clusters in the project: if two clusters have a products database and a user has a role granting read access on the products database, the user has that access on both clusters.

## Parameters
`ProjectId` *(required)* : Unique identifier of the Atlas project to which the user belongs.<br>
`DatabaseName` *(required)* : The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is always the admin database.<br>
`Labels` *(optional)* : Array containing key-value pairs that tag and categorize the database user.<br>
`LdapAuthType` *(optional)* : Method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE.<br>
`Roles` *(optional)* : Array of this user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well.<br>
`Username` *(required)* : Username for authenticating to MongoDB.<br>
`Password'` *(optional)* : The user’s password. This field is not included in the entity returned from the server.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas.<br>

## Installation
    $ make
    $ cfn submit
    ...