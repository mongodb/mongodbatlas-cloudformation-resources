# Project

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Project  L1 CDK constructor
 - Atlas basis L3 CDK constructor
 - Encryption at rest L3 CDK constructor
 - Atlas Quickstart
 - Atlas Quickstart Fargate


## Prerequisites 
### Resources needed to run the manual QA
- Atlas organization
- Atlas Team 
- Atlas user


All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- The project should be visible in the project list page:
![image](https://user-images.githubusercontent.com/5663078/227225795-0f1b6650-95fe-40ca-942d-99902b747aa2.png)
- The api keys should be visible in the project API Keys page:
![image](https://user-images.githubusercontent.com/5663078/227303503-14e7a53b-92a0-46f3-9f4a-6ea9fbf2a20d.png)
- The team should be visible in the project Team page:
![image](https://user-images.githubusercontent.com/5663078/227303779-16069213-4fe7-49c8-a840-66afdb88cb6e.png)

## Automated E2E Testing:
- E2E tests are located under `cfn-resources/test/e2e/project`. Please ensure to run/update those accordingly.
  
To run the automated E2E tests use following commands:
```
cd cfn-resources/test/e2e/project
go test -v project_test.go
```

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-projects)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/tutorial/manage-projects/)

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
    },
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
    },
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