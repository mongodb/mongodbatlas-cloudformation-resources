# MongoDB::Atlas::ProjectInvitation

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Project invitation L1 CDK constructor



## Prerequisites 
### Resources needed to manually QA
- Atlas organization (based on Atlas profile configured)

All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- Atlas Project Access Manager should show `PENDING INVITE` for invited user:

![image](https://user-images.githubusercontent.com/122359335/227505950-afc41fa7-abb5-478b-807d-c9510a40888c.png)

2. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-createprojectinvitation)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/invitations/)