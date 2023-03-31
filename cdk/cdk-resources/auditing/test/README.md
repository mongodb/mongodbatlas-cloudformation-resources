# Auditing


## CFN resource type used
- MongoDB::Atlas::Auditing

This CFN resource must be active in your AWS account while using this constructor.


## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Auditing CFN resource](../../../../cfn-resources/auditing/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- Database Auditing Setting for the respective Project in Atlas should be correctly configured:
![image](https://user-images.githubusercontent.com/5663078/227519864-2d147a0b-4e57-48f8-8de8-48370f1cd037.png)



## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Auditing)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/database-auditing/)