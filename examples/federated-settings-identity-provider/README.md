# How to create a MongoDB::Atlas::FederatedSettingsIdentityProvider

## Step 1: Activate the resource in CloudFormation

Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

Step b: Search for MongoDB::Atlas::FederatedSettingsIdentityProvider resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )

Step c: Select and activate
Enter the RoleArn that is created in step 1.

Your FederatedSettingsIdentityProvider Resource is ready to use.

## Step 2: Create template using [federated-settings-identity-provider-oidc.json](federated-settings-identity-provider-oidc.json)

    Note: Make sure you are providing appropriate values for:
    1. FederationSettingsId (24-character hexadecimal string)
    2. IdentityProviderName
    3. IssuerUri
    4. Audience
    5. ClientId
    6. GroupsClaim (optional)
    7. UserClaim (optional)
    8. RequestedScopes (optional)
    9. AssociatedDomains (optional)
    10. Description (optional)
    11. AuthorizationType (optional)
    12. IdpType (optional, required for OIDC)
    13. Profile (optional)

## Important Notes:

- **CREATE operation only supports OIDC protocol**. SAML identity providers must be imported.
