# Federated Settings Org Role Mapping

## CFN resource type used
- MongoDB::Atlas::FederatedSettingsOrgRoleMapping

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [FederatedSettingsOrgRoleMapping CFN resource](../../../../cfn-resources/federated-settings-org-role-mapping/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).
- Use this URL to view RoleMappings configured for your organization: https://cloud.mongodb.com/v2#/federation/<ATLAS_FEDERATED_SETTINGS_ID>/organizations/<ATLAS_ORG_ID>/roleMappings


### Success criteria when testing the resource
1. Role mappings should be correctly configured for the Organization under Federation settings:

![image](https://user-images.githubusercontent.com/122359335/227274727-bee557f4-8def-467e-ad37-adcae1887911.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Federated-Authentication/operation/createRoleMapping)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security/manage-role-mapping/#role-mapping-process)