# MongoDB::Atlas::ServerlessPrivateEndpoint

**WARNING:** This resource is deprecated and will be removed in January 2026. For more details, see [Migrate your programmatic tools from M2, M5, or Serverless Instances to Flex Clusters](https://www.mongodb.com/docs/atlas/flex-migration/).

Resource for creating and managing [Private Endpoint Services](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Serverless-Private-Endpoints).

If the CreateAndAssignAWSPrivateEndpoint Only the Atlas private endpoint will be deleted, the AWS private endpoint will remain and might need to be deleted manually

This version of private endpoint serverless, has 2 flows:
- Without AWS private endpoint
- With AWS private endpoint

## Without AWS Private Endpoint Flow:
In this configuration, you have the flexibility to establish an Atlas private endpoint independently. 
By opting not to immediately connect it, you can manually set up the linkage with an AWS private endpoint or Azure.
Simply set the CreateAndAssignAWSPrivateEndpoint attribute to false during resource creation.

## With AWS Private Endpoint Flow:
For a streamlined experience, you can enable the automatic generation of both Atlas private and AWS private endpoints,
along with their interconnection, all within a single resource.
Achieve this by configuring the CreateAndAssignAWSPrivateEndpoint attribute as true and supplying the following configuration structure:

``` json
"AwsPrivateEndpointConfigurationProperties": {
    "VpcId": "vpc-zxxxxxx",
    "SubnetIds": ["subnet-xxxxxx", "subnet-yyyyy"],
    "Region": "us_east_1"
}
```

## Requirements

Set up an AWS profile to securely give CloudFormation access to your Atlas credentials.
For instructions on setting up a profile, [see here](/README.md#mongodb-atlas-api-keys-credential-management).

## Attributes and Parameters

See the [resource docs](docs/README.md).

## Cloudformation Examples

### With AWS private endpoint
See the examples [CFN Template](/examples/serverless-private-endpoint/serverless-private-endpoint-with-aws-private-endpoint.json) for example resource.

### Without AWS private endpoint
See the examples [CFN Template](/examples/serverless-private-endpoint/serverless-private-endpoint-without-aws-private-endpoint.json) for example resource.
