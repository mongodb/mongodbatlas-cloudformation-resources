# MongoDB Atlas AWS CloudFormation Resources & AWS Partner Solution Deployments 
Use AWS CloudFormation to manage [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).

Partner Solutions (formally AWS Quick Starts) are automated reference deployments built by Amazon Web Services (AWS) solutions architects and AWS Partners. Partner Solutions help you deploy popular technologies to AWS according to AWS best practices. The quickest way to get started is to launch the official [MongoDB Atlas on AWS](https://aws.amazon.com/quickstart/architecture/mongodb-atlas/) Partner Solution Deployment directly from the AWS Management Console.

## Getting Started
See the [examples](examples/README.md) for how to setup prerequisites & get started with your first cluster, using our sample CloudFormation Stack templates.

## Support, Bugs, Feature Requests
Feature requests can be submitted at [feedback.mongodb.com](https://feedback.mongodb.com/forums/924145-atlas/category/392596-atlas-cloudformation-resources) - just select "Atlas CloudFormation Resources" as the category or vote for an already suggested feature.

Support for the MongoDB Atlas Resource Provider for CloudFormation is provided under MongoDB Atlas support plans, starting with Developer. Please submit support questions within the Atlas UI. In addition, support questions submitted under the Issues section of this repo are also being monitored. Bugs should be filed under the Issues section of this repo.

## MongoDB Atlas API Keys Credential Management
Atlas API keys Configuration are required for both CloudFormation and CDK resources, and this Atlas API key pair are provided as input by the use of a Profile

AWS CloudFormation limits Third Parties from using non-AWS API Keys as either hardcoded secrets in CloudFormation templates or via CDK, hence we now require all the users store MongoDB Atlas API Keys via [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/).   

`NOTE: the process for configuring the PROFILE is the same and is required both for CloudFormation and CDK`

### 1. Configure your MongoDB Atlas API Keys 
You'll need to generate an API key pair (public and private keys) for your Atlas organization and configure them to grant CloudFormation access to your Atlas project.
Refer to the [Atlas documentation](https://www.mongodb.com/docs/atlas/configure-api-access/#manage-programmatic-access-to-an-organization) for detailed instructions.

### 2. Configure your Profile
To use Atlas CloudFormation resources, you must configure a "profile" with your API keys using [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/).

The secret should follow this format:
```
SecretName: cfn/atlas/profile/{ProfileName}
SecretValue: {PublicKey: {YourPublicKey}, PrivateKey: {YourPrivateKey}}
```

To create a new secret for a default profile, use the [PROFILE SECRET TEMPLATE](/examples/profile-secret.yaml) file provided in this repository.

Here are some examples of how to use this template:

#### example 1:
```
  ProfileName: default
  SecretName: cfn/atlas/profile/default
  SecretValue = {PublicKey: xxxxxxx , PrivateKey: yyyyyyyy}
```
#### example 2:
```
  ProfileName: tetProfile1
  SecretName: cfn/atlas/profile/tetProfile1
  SecretValue = {PublicKey: zzzzzzzzzz , PrivateKey:jjjjjjjjj}
```

### 3. Provide the profile to your CloudFormation template

All Atlas CloudFormation resources include a "Profile" property that specifies which profile to use. You'll need to provide the profile you created in the previous step to the CloudFormation template.

Note that if you don't provide a profile, the resource will use a default profile (will try to get a secret named cfn/atlas/profile/default). We recommend always specifying the profile to avoid any unexpected behavior.

Once you've provided the profile, you can deploy the CloudFormation stack using the AWS Console or the AWS CLI. Refer to the AWS documentation for instructions on how to deploy CloudFormation stacks.

IMPORTANT: when specifying the profile in your CloudFormation template, you must specify the Profile Name, NOT the Secret Name

Correct usage:
```
  "Profile" : "ProfileName"
```
Incorrect usage:
```
  "Profile" : "cfn/atlas/profile/ProfileName"
```

## Logging 

Logging for AWS CloudFormation Public extensions is currently disabled. AWS is evaluating if logging is useful for consumers of third party extensions, if this is something you need or would like to request please open a ticket directly with AWS Support.

## Contributing

See our [CONTRIBUTING.md](CONTRIBUTING.md) guide.


## Troubleshooting
The following are common issues encountered when using AWS CloudFormation/CDK with MongoDB Atlas Resources  
1. Activate the 3rd party extension for each resource (i.e. MONGODB::ATLAS::[RESOURCE-NAME]) in each AWS region and from each AWS account that you wish to deploy from
2. Ensure you have sufficiently strong AWS IAM Activation Role attached to each 3rd party extension. For sample IAM Role see [here](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/cfn-resources/execute-role.template.yml)
3. Ensure your activated 3rd party public extension matches name exactly to MONGODB::ATLAS::[RESOURCE-NAME] (you may need to delete private extension if this namespace is already occupied)
4. Ensure your MongoDB Atlas Programmatic API Keys (PAKs) being used with CloudFormation have sufficiently strong permissions (Organization Project Creator or Organization Owner)
5. Ensure your MongoDB Atlas PAKs have correct IP Address / CIDR range access. For testing purposes with caution you can open keys to all access by adding “0.0.0.0/1” and “128.0.0.0/1” (do not use for production workloads). 
6. How to determine which IP address AWS CloudFormation uses to deploy MongoDB Atlas resouces with my Atlas Programmatic API Keys (PAK)?
  When you deploy MongoDB Atlas using CloudFormation with your Atlas PAK, CloudFormation will default to use the IP address of the machine from which you are making the API call. The "machine" however making the API call to 3rd party MongoDB Atlas API would be various AWS servers hosting Lambda functions and won't be static. Review additional details at https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html and contact AWS Support directly who can help confirm CIDR range to be used in your Atlas PAK IP Whitelist.
