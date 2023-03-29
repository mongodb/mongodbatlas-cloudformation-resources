# @mongodbatlas-awscdk/atlas-encryption-at-rest

## CFN resource type used
- MongoDB::Atlas::EncryptionAtRest
- MongoDB::Atlas::Cluster
- MongoDB::Atlas::DatabaseUser
- MongoDB::Atlas::ProjectIpAccessList

These CFN resources must be active in your AWS account while using this constructor.

## Manual QA
### Prerequisite steps:
1. Create an Atlas Project 
2. [Create CloudAccessProvider](https://www.mongodb.com/docs/atlas/reference/api/cloud-provider-access-create-one-role/) role for your project via postman:
  
`POST https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/<ATLAS_PROJECT_ID>/cloudProviderAccess`
 - Note `roleId`, `atlasAWSAccountArn` and `atlasAssumedRoleExternalId` from the response. You will be using these in the next steps.
3. Use `atlasAWSAccountArn`, `atlasAssumedRoleExternalId` from last step to create an IAM role for Atlas in your AWS account as follows:
 ```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "<atlasAWSAccountArn>"
            },
            "Action": "sts:AssumeRole",
            "Condition": {
                "StringEquals": {
                    "sts:ExternalId": "atlasAssumedRoleExternalId"
                }
            }
        }
    ]
}
```


4. Authorize your newly created IAM role for the `roleId` with Atlas via postman:
```
PATCH https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/<ATLAS_PROJECT_ID>/cloudProviderAccess/<roleId>

RequestBody:
{
"providerName": "AWS",
"iamAssumedRoleArn": "<AWS_IAM_ROLE_ARN>"
}
```

5. Create a Key in AWS KMS: Follow UI prompts and on `Define key administrative permissions` page under `Key administrators` select your IAM role created in last step.
6. Once KMS key is created in KMS, note your Key ID, this is your Customer Master Key.

### Steps to test:
1. Follow prerequisite steps for testing a CDK construct in [TESTING.md](../../../TESTING.md).
2. Use `roleID` from prerequisite steps above as role id param for this CDK construct
3. Set any additional required configuration options/parameters as per your needs.
4. Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Ensure all underlying resources are configured correctly as specified (Atlas Cluster, EncryptionAtRest, Network IPAccessList and DatabaseUser):

![image](https://user-images.githubusercontent.com/122359335/228581798-691ab912-4397-4fd5-aa7d-237b88a46465.png)

![image](https://user-images.githubusercontent.com/122359335/228581863-e04ee602-68b7-4e70-a88f-96da4ea2535d.png)

![image](https://user-images.githubusercontent.com/122359335/228581979-615b19cb-4d72-47e6-8aca-44a8c3013825.png)

![image](https://user-images.githubusercontent.com/122359335/228582032-14ef9f65-dc8d-4bb4-b7b0-1ab77a6257e6.png)

2. Ensure general [CDK resource success criteria](../../../TESTING.md#success-criteria-to-be-satisfied-when-testing-a-construct) for this resource is met.

## Important links
- [Resource Usage Doc](https://www.mongodb.com/docs/atlas/security-aws-kms/#std-label-security-aws-kms)