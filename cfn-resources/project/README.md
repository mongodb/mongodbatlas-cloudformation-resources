# MongoDB::Atlas::Project

## Description
Returns, adds, and edits collections of clusters and users in MongoDB Cloud.

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

## Local Testing

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

### Build Handler
```bash
make build
```
### Run the handler in a docker container
```bash
# Required the docker daemon running
sam local start-lambda --skip-pull-image
```

### Update the SAM template
Update the template file in `test/templates/sam/project.sample-cfn-request.json` and add the `Name` (project name), `OrgId` and `ApiKeys`.
Example:
```yaml
"desiredResourceState":{
    "Name": "YourProjectName",
    "OrgId": "60ddf55c27a5a20955a707d7",
    "ApiKeys": {
      "PublicKey": "wwdsirvb",
      "PrivateKey": "privateKey"
    }
```

### Test the handler operations CREATE and READ
```bash
cfn invoke --function-name TestEntrypoint resource CREATE test/templates/project.sample-cfn-request.json
cfn invoke --function-name TestEntrypoint resource READ test/templates/project.sample-cfn-request.json
```

### Update the SAM template to test the DELETE operation
In order to test DELETE, you need to add the property `Id` (projectId) in `test/templates/sam/project.sample-cfn-request.json`.
Example:
```yaml
"desiredResourceState":{
    "Name": "YourProjectName",
    "OrgId": "60ddf55c27a5a20955a707d7",
    "Id": "63dcc31db5a65b3c3500bc62",
    "ApiKeys": {
      "PublicKey": "wwdsirvb",
      "PrivateKey": "privateKey"
    }
```
You can retrieve the projectId to add to the sam template by running:
```yaml
cfn invoke --function-name TestEntrypoint resource READ test/templates/project.sample-cfn-request.json
```
### Test the handler operations DELETE
```bash
cfn invoke --function-name TestEntrypoint resource DELETE test/templates/project.sample-cfn-request.json
```
### Update the SAM template to test the UPDATE operation
In order to test UPDATE, you need to add the property `Id` (projectId) and `ProjectTeams` in `test/templates/sam/project.sample-cfn-request.json`.
Example:
```yaml
"desiredResourceState":{
    "Name": "YourProjectName",
    "OrgId": "60ddf55c27a5a20955a707d7",
    "Id": "63dcc31db5a65b3c3500bc62",
    "ApiKeys": {
      "PublicKey": "wwdsirvb",
      "PrivateKey": "privateKey"
    },
   "ProjectTeams": [
        {
            "TeamId": "63dccf0bb5a65b3c3500d5d7",
            "RoleNames": ["GROUP_OWNER"]
        }
    ]
```
You can retrieve the teams available in your organization with [AtlasCLI](https://github.com/mongodb/mongodb-atlas-cli):
```bash
atlas teams list
ID                         NAME
63dccf0bb5a65b3c3500d5d7   Test
```
### Test the handler operations UPDATE
```bash
cfn invoke --function-name TestEntrypoint resource UPDATE test/templates/project.sample-cfn-request.json
```

## Installation
TAGS=logging make
cfn submit --verbose --set-default

## Usage

The [launch-x-quickstart.sh](../../quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other neccessary parameters.

You can use the project.sample-template.yaml to create a stack using the resource.
Similar to the local testing described above you can follow the logs for the deployed
lambda function which handles the request for the Resource Type.

In one shell session:
```
aws logs tail mongodb-atlas-project-logs --follow
```

And then you can create the stack with a helper script it insert the apikeys for you:


```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/project/test/project.sample-template.yaml SampleProject1 ParameterKey=OrgId,ParameterValue=${ATLAS_ORG_ID}
```

## For More Information 
See the MongoDB Atlas API [Project Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Projects) Documentation.
