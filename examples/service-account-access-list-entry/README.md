# How to create a MongoDB::Atlas::ServiceAccountAccessListEntry

## Step 1: Activate the resource in CloudFormation

Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

Step b: Search for MongoDB::Atlas::ServiceAccountAccessListEntry resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )

Step c: Select and activate
Enter the RoleArn that is created in step 1.

Your ServiceAccountAccessListEntry Resource is ready to use.

## Step 2: Create template using [service-account-access-list-entry.json](service-account-access-list-entry.json)

    Note: Make sure you are providing appropriate values for:
    1. OrgId
    2. ClientId (of an existing Service Account)
    3. CIDRBlock or IPAddress
    4. Profile (optional)

## Important Notes

- **Existing Service Account Required**: You must have an existing Service Account. Create one using the [service-account example](../service-account/) or via the MongoDB Atlas UI.
- **CIDR or IP Address**: You must specify either a CIDR block or an IP address, but not both in the same resource.
- **Auto-generated CIDR**: When you provide an IP address, Atlas automatically generates a `/32` CIDR block.
- **No Updates**: This resource does not support updates. Any property change will trigger a replacement (delete + create).
- **List Operation**: Use the List handler to retrieve all access list entries for a service account.

## Example Use Cases

This example demonstrates both patterns:

1. Adding a CIDR block (e.g., `203.0.113.0/24`) to allow access from a subnet
2. Adding a single IP address (e.g., `198.51.100.10`) to allow access from a specific host
