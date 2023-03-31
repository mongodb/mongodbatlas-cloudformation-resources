# Org Invitation

## CFN resource type used
- MongoDB::Atlas::OrgInvitation

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [OrgInvitation CFN resource](../../../../cfn-resources/org-invitation/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Under Access Manager for your Atlas Organization, you should see a new entry under “Users” tab with status “PENDING INVITE”

![image](https://user-images.githubusercontent.com/122359335/227275914-4af66737-fa72-49f8-8713-9d298606bc4f.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api/organization-create-one-invitation/#invite-one-user-to-an-service-organization)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/invitations/#invitations-to-organizations-and-projects)