# Teams

## CFN resource type used
- MongoDB::Atlas::Teams

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Teams CFN resource](../../../../cfn-resources/teams/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Atlas Team should show correctly configured in respective Organization's Access Manager (and in Project Access Manager, if applicable):

![image](https://user-images.githubusercontent.com/122359335/227534552-a338f068-2e60-4179-91cd-7a634a2dc9b3.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Teams)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/access/manage-teams-in-orgs/)