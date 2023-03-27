# MongoDB::Atlas::CustomDnsConfigurationClusterAws

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Custom DNS configuration L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
All resources are created as part of `cfn-testing-helper.sh`:

- Atlas project

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
1. AWS custom DNS should be enabled for the project. This can be validated via a [GET API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/AWS-Clusters-DNS/operation/getAWSCustomDNS) call as:
```
https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/<ATLAS_PROJECT_ID>/awsCustomDNS
```

![image](https://user-images.githubusercontent.com/122359335/227661815-d48398a9-aaa3-4978-9de4-736acab6ddf8.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.


## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/AWS-Clusters-DNS/operation/toggleAWSCustomDNS)

