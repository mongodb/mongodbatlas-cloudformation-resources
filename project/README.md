# MongoDB::Atlas::Project

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