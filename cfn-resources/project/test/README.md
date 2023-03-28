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
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- The project should be visible in the project list page:
![image](https://user-images.githubusercontent.com/5663078/227225795-0f1b6650-95fe-40ca-942d-99902b747aa2.png)
- The api keys should be visible in the project API Keys page:
![image](https://user-images.githubusercontent.com/5663078/227303503-14e7a53b-92a0-46f3-9f4a-6ea9fbf2a20d.png)
- The team should be visible in the project Team page:
![image](https://user-images.githubusercontent.com/5663078/227303779-16069213-4fe7-49c8-a840-66afdb88cb6e.png)
## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Projects)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/tutorial/manage-projects/)
