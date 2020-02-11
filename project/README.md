# MongoDB::Atlas::Project

Congratulations on starting development!

Next steps:

1. Populate the JSON schema describing your resource, `mongodb-atlas-project.json`
2. The RPDK will automatically generate the correct resource model from the
   schema whenever the project is built via Make.
   You can also do this manually with the following command: `cfn-cli generate`
3. Implement your resource handlers by adding code to provision your resources in your resource handler's methods.

Please don't modify files `model.go` and `main.go`, as they will be automatically overwritten.

## Attributes
`Name` *(required)* : Name of the project to create.<br>
`OrgId` *(required)* : Unique identifier of the organization within which to create the project.<br>
`Id` : The unique identifier of the project.<br>
`Created` : The ISO-8601-formatted timestamp of when Atlas created the project.<br>
`ClusterCount` : The number of Atlas clusters deployed in the project.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas.<br>
