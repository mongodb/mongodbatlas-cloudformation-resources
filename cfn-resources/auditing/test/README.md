# Auditing 

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Auditing L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- Database Auditing Setting for the respective Project in Atlas should be correctly configured:
![image](https://user-images.githubusercontent.com/5663078/227519864-2d147a0b-4e57-48f8-8de8-48370f1cd037.png)

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Auditing)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/database-auditing/)