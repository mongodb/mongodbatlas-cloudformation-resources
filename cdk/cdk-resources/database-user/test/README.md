# Database User

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.

- Atlas basis L3 CDK constructor
- Encryption at rest L3 CDK constructor
- Atlas Quickstart
- Atlas Quickstart Fargate


## CFN resource type used
- MongoDB::Atlas::DatabaseUser

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Database User CFN resource](../../../../cfn-resources/database-user/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- The Database User should be visible in the "Database Users" page:
![image](https://user-images.githubusercontent.com/5663078/227314604-d15f10a4-5e3b-4010-b94f-621ec55eceb3.png)