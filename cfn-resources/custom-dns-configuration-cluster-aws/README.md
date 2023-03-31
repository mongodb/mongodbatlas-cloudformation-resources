# MongoDB::Atlas::CustomDnsConfigurationClusterAws

## Description
Returns, adds and removes custom DNS configurations for MongoDB Atlas database deployments on AWS.

## Attributes & Parameters

Please consult the [Resource Docs](https://github.com/PeerIslands/mongodbatlas-cloudformation-resources/blob/feature-custom-dns-config-cluster-aws/cfn-resources/custom-dns-configuration-cluster-aws/docs/README.md)

## Local Testing

The local tests are integrated with the AWS sam local and cfn invoke tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
cd ${repo_root}/cfn-resources/custom-dns-configuration-cluster-aws
./test/custom-dns-config-cluster-aws.create-sample-cfn-request.sh > test.request.json
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json
cfn invoke DELETE test.request.json
```

Both CREATE & DELETE tests must pass.

## CloudFormation Example
Please see the [CFN Template](../../examples/custom-dns-configuration-cluster-aws/CustomDnsConfigurationClusterAws.json) for example resource

## Installation
```
TAGS=logging make
cfn submit --verbose --set-default
```

Usage
Examples aws cloudformation template is available here example template.


```bash
#Configure you AWS Credentials to create Cloudformation Stack
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export AWS_REGION=""
export AWS_DEFAULT_REGION=""

#Command to deploy the sample CustomeDNSConfigurationClusterAWS stack (Before this step "cfn submit" should have been executed successfully)
./examples/custom-dns-configuration-cluster-aws/Deploy.sh
```

| Operation | Flag       | Reference links                                                                                                                                                                                                                                                 |
|-----------|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| PATCH     | true/false | [From MongoDB](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Custom-DNS-for-Atlas-Clusters-Deployed-to-AWS/operation/toggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAws) |
| READ      | true/false | [From MongoDB](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Custom-DNS-for-Atlas-Clusters-Deployed-to-AWS)                                    |


For more information see: MongoDB Atlas API Endpoint [Custom DNS for Atlas Clusters Deployed to AWS](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Custom-DNS-for-Atlas-Clusters-Deployed-to-AWS) Documentation.