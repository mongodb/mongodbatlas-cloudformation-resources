## MongoDB::Atlas::FederatedSettingsOrgRoleMapping

### Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- Federated settings org role mapping L1 CDK constructor


### Resources (and parameters for local tests) needed to manually QA:
These Atlas federated setting ID must be manually created.
- Atlas Organization (ATLAS_ORG_ID)
- Atlas federated settings id (ATLAS_FEDERATED_SETTINGS_ID)
- Atlas Project (created by cfn-test-create-inputs.sh)


## Manual QA:

### Prerequisite steps:
1. Create an Atlas Organization if you don’t already have one and get the OrgId from URL (https://cloud.mongodb.com/v2#/org/<ATLAS_ORG_ID>/projects)
2. Go to your organization settings and click on “Visit Federation Management App” under “Manage Federation Settings”
3. Note the federationSettingsId from the URL (https://cloud.mongodb.com/v2#/federation/<ATLAS_FEDERATED_SETTINGS_ID>/overview).
4. Configure your federation by configuring domains and Identity Providers.
5. Export ATLAS_ORG_ID and ATLAS_FEDERATED_SETTINGS_ID environment variables.

### Steps to test:
1. Ensure prerequisites above for this resource and general [prerequisites](../../../TESTING.md.md#prerequisites) are complete.
2. Follow [general steps](../../../TESTING.md.md#steps) to test a CFN resource.
3. Use this URL to view RoleMappings configured for your organization: https://cloud.mongodb.com/v2#/federation/<ATLAS_FEDERATED_SETTINGS_ID>/organizations/<ATLAS_ORG_ID>/roleMappings

### Success criteria when testing the resource
1. Role mappings should be correctly configured for the Organization under Federation settings:  

![image](https://user-images.githubusercontent.com/122359335/227274727-bee557f4-8def-467e-ad37-adcae1887911.png)

2. General [CFN resource success criteria](../../../TESTING.md.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Federated-Authentication/operation/createRoleMapping)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security/manage-role-mapping/#role-mapping-process)