# Custom DNS Configuration Cluster for AWS

## CFN resource type used
- MongoDB::Atlas::CustomDnsConfigurationClusterAws

This CFN resource must be active in your AWS account while using this constructor.

## Manual QA
- Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
- Follow prerequisite steps for the corresponding [Custom DNS Configuration Cluster for AWS CFN resource](../../../../cfn-resources/custom-dns-configuration-cluster-aws/test/README.md).
- Set any additional required configuration options/parameters as per your needs.
- Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. AWS custom DNS should be enabled for the project. This can be validated via a [GET API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/AWS-Clusters-DNS/operation/getAWSCustomDNS) call as:
```
https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/<ATLAS_PROJECT_ID>/awsCustomDNS
```

![image](https://user-images.githubusercontent.com/122359335/227661815-d48398a9-aaa3-4978-9de4-736acab6ddf8.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/AWS-Clusters-DNS/operation/toggleAWSCustomDNS)
