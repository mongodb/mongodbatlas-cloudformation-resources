# MongoDB::Atlas::DatabaseUser

## Description
The databaseUsers resource lets you retrieve, create and modify the MongoDB users in your cluster. Each user has a set of roles that provides access to the project’s databases. A user’s roles apply to all the clusters in the project: if two clusters have a products database and a user has a role granting read access on the products database, the user has that access on both clusters.

## Parameters
`ProjectId` *(required)* : Unique identifier of the Atlas project to which the user belongs.<br>
`DatabaseName` *(required)* : The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is always the admin database.<br>
`Labels` *(optional)* : Array containing key-value pairs that tag and categorize the database user.<br>
`LdapAuthType` *(optional)* : Method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE.<br>
`Roles` *(optional)* : Array of this user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well.<br>
`Username` *(required)* : Username for authenticating to MongoDB.<br>
`Password'` *(optional)* : The user’s password. This field is not included in the entity returned from the server.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit
    ...
