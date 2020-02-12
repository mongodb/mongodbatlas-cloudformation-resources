# MongoDB::Atlas::Project

Congratulations on starting development!

Next steps:

1. Populate the JSON schema describing your resource, `mongodb-atlas-project.json`
2. The RPDK will automatically generate the correct resource model from the
   schema whenever the project is built via Make.
   You can also do this manually with the following command: `cfn-cli generate`
3. Implement your resource handlers by adding code to provision your resources in your resource handler's methods.

Please don't modify files `model.go` and `main.go`, as they will be automatically overwritten.

## Description
This resource allows you to create a project, get one or a list of projects, and delete a project. Atlas provides a hierarchy based on organizations and projects to facilitate the management of your Atlas clusters. Projects allow you to isolate different environments (for instance, development/qa/prod environments) from each other, associate different users or teams with different environments, maintain separate cluster security configurations, and create different alert settings.

## Attributes
`Id` : The unique identifier of the project.<br>
`Created` : The ISO-8601-formatted timestamp of when Atlas created the project.<br>
`ClusterCount` : The number of Atlas clusters deployed in the project.<br>

##Parameters
`Name` *(required)* : Name of the project to create.<br>
`OrgId` *(required)* : Unique identifier of the organization within which to create the project.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit
    ...