# Custom DB Roles

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Custom DB Role L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project. This resource is created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- Custom role should be available in the "Database Access" page:
![image](https://user-images.githubusercontent.com/5663078/227566882-b6bb8a83-988a-402e-9211-ffc0073c5aed.png)

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Custom-Database-Roles)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-add-mongodb-roles/)
