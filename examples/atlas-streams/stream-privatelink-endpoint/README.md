# How to create a MongoDB::Atlas::StreamPrivatelinkEndpoint

## Step 1: Activate the stream privatelink endpoint resource in CloudFormation

Step a: Create Role using [execution-role.yaml](../../execution-role.yaml) in examples folder.

Step b: Search for MongoDB::Atlas::StreamPrivatelinkEndpoint resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )

Step c: Select and activate
Enter the RoleArn that is created in step 1.

Your StreamPrivatelinkEndpoint Resource is ready to use.

## Step 2: Choose a template based on your vendor

### Example 1: S3 Private Link Endpoint ([stream-privatelink-endpoint-s3.json](stream-privatelink-endpoint-s3.json))

Creates a PrivateLink endpoint for AWS S3 to enable secure, private connectivity between Atlas Stream Processing and S3 buckets.

**Parameters:**

1. **ProjectId** - Atlas Project Id (24 hexadecimal characters)
2. **Region** - AWS region where the S3 bucket is located (e.g., `us-east-1`, `eu-west-1`)
3. **ServiceEndpointId** - S3 service endpoint ID in the format `com.amazonaws.<region>.s3`
4. **Profile** - Secret Manager Profile for Atlas credentials (optional, default: `default`)

**Note:** Vendor is `S3`, ProviderName is `AWS`

### Example 2: MSK Private Link Endpoint ([stream-privatelink-endpoint-msk.json](stream-privatelink-endpoint-msk.json))

Creates a PrivateLink endpoint for Amazon MSK (Managed Streaming for Kafka) to enable secure connectivity to MSK clusters.

**Parameters:**

1. **ProjectId** - Atlas Project Id (24 hexadecimal characters)
2. **MskClusterArn** - Amazon Resource Name (ARN) of the MSK cluster
3. **Profile** - Secret Manager Profile for Atlas credentials (optional, default: `default`)

**Note:** Vendor is `MSK`, ProviderName is `AWS`, Region is auto-computed from the ARN

### Example 3: Confluent Cloud Private Link Endpoint ([stream-privatelink-endpoint-confluent.json](stream-privatelink-endpoint-confluent.json))

Creates a PrivateLink endpoint for Confluent Cloud Kafka to enable secure connectivity to Confluent clusters.

**Parameters:**

1. **ProjectId** - Atlas Project Id (24 hexadecimal characters)
2. **ConfluentRegion** - Domain name of the Confluent cluster (e.g., `confluent.cloud`)
3. **ServiceEndpointId** - VPC Endpoint service name from Confluent Cloud
4. **DnsDomain** - DNS domain for the Confluent cluster (e.g., `pkc-xxxxx.us-west-2.aws.confluent.cloud`)
5. **DnsSubDomain0**, **DnsSubDomain1**, **DnsSubDomain2** - Availability zone sub-domains (leave empty if not used)
6. **Profile** - Secret Manager Profile for Atlas credentials (optional, default: `default`)

**Note:** Vendor is `CONFLUENT`, ProviderName is `AWS`. Set DnsSubDomain parameters to empty strings if your cluster doesn't use subdomains.

## Vendor Comparison

| Vendor        | Required Fields                                            | Use Case                   | Region Format                  |
| ------------- | ---------------------------------------------------------- | -------------------------- | ------------------------------ |
| **S3**        | `Region`, `ServiceEndpointId`                              | Access S3 buckets          | AWS region (e.g., `us-east-1`) |
| **MSK**       | `Arn`                                                      | Connect to Amazon MSK      | Auto-computed from ARN         |
| **CONFLUENT** | `Region`, `ServiceEndpointId`, `DnsDomain`, `DnsSubDomain` | Connect to Confluent Cloud | Confluent domain name          |

## Important Notes

1. **Update Not Supported**: Private link endpoints cannot be updated after creation. Delete and recreate to modify the configuration.

2. **State Output**: Check the `State` output to verify the connection status (`PENDING`, `DONE`, or `FAILED`).

3. **AWS Prerequisites**:

   - For S3: VPC endpoint for S3 must exist in your VPC
   - For MSK: MSK cluster must be configured with private connectivity
   - For Confluent: PrivateLink must be enabled in your Confluent Cloud cluster

4. **Next Steps**: After the endpoint reaches `DONE` state, use the returned `InterfaceEndpointId` when creating Stream Connections.
