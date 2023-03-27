# Custom DB Roles


## CFN resource type used
- MongoDB::Atlas::CustomDBRole

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Custom DB Roles CFN resource](../../../../cfn-resources/custom-db-role/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- Custom role should be available in the "Database Access" page:
![image](https://user-images.githubusercontent.com/5663078/227566882-b6bb8a83-988a-402e-9211-ffc0073c5aed.png)