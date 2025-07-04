## MongoDB::Atlas::Teams

### Impact 
The following components use this resource and are potentially impacted by any changes.
They should also be validated to ensure the changes do not cause a regression.
 
- Teams L1 CDK constructor

### Resources (and parameters for local tests) needed to manually QA:
The Atlas organization must be manually provided.
- Atlas Organization (`MONGODB_ATLAS_ORG_ID`)
- Atlas Project (created as part of `cfn-testing-helper.sh`)

## Manual QA:

### Prerequisites
1. Create an Atlas Organization if you don’t already have one.
2. Export environment variable `MONGODB_ATLAS_ORG_ID` with the value as your Organization Id.
   You can get this ID from your Atlas UI by clicking on your organization and note the Organization ID from the URL
   which should look like this https://cloud.mongodb.com/v2#/org/<MONGODB_ATLAS_ORG_ID>/projects.

### Steps to test:
1. Ensure prerequisites above for this resource and general [prerequisites](/TESTING.md#prerequisites) are complete.
2. Follow [general steps](/TESTING.md#steps) to test a CFN resource.

### Success criteria when testing the resource
1. Atlas Team should show correctly configured in respective Organization's Access Manager (and in Project Access Manager, if applicable):
   ![image](https://user-images.githubusercontent.com/122359335/227534552-a338f068-2e60-4179-91cd-7a634a2dc9b3.png)
2. General [CFN resource success criteria](/TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-teams)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/access/manage-teams-in-orgs/)
