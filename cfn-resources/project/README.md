# MongoDB::Atlas::Project

## Description
This resource allows you to create a project, get one or a list of projects, and delete a project. Atlas provides a hierarchy based on organizations and projects to facilitate the management of your Atlas clusters. Projects allow you to isolate different environments (for instance, development/qa/prod environments) from each other, associate different users or teams with different environments, maintain separate cluster security configurations, and create different alert settings.

## Attributes
`Id` : The unique identifier of the project.<br>
`Created` : The ISO-8601-formatted timestamp of when Atlas created the project.<br>
`ClusterCount` : The number of Atlas clusters deployed in the project.<br>

## Parameters
`Name` *(required)* : Name of the project to create.<br>
`OrgId` *(required)* : Unique identifier of the organization within which to create the project.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>


## Local Testing

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/project
./test/project.create-sample-cfn-request.sh YourProjectName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.

## Installation
TAGS=logging make
cfn submit --verbose --set-default

## Usage

The [/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh](launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other neccessary parameters.

You can use the project.sample-template.yaml to create a stack using the resource.

```bash
repo_root=$(git rev-parse --show-toplevel)
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/project/test/project.sample-template.yaml SampleProject1 ParameterKey=OrgId,ParameterValue=${ATLAS_ORG_ID}
```

