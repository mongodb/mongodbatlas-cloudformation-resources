# MongoDB Atlas AWS CloudFormation Resources & AWS Partner Solution Deployments 

Use AWS CloudFormation to manage [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).

Partner Solutions (formally AWS Quick Starts) are automated reference deployments built by Amazon Web Services (AWS) solutions architects and AWS Partners. Partner Solutions help you deploy popular technologies to AWS according to AWS best practices. The quickest way to get started is to launch the official [MongoDB Atlas on AWS](https://aws.amazon.com/quickstart/architecture/mongodb-atlas/) Partner Solution Deployment directly from the AWS Management Console.

# Prerequisites

To use CloudFormation MongoDB Atlas public extensions (AWS Third Party), you have to activate the Public extension from your AWS console. You have to do this in each AWS Account and in each AWS Region. To activate Public extension, we have to create an execution role and pass the ARN of the role as an input. Use [this template](examples/execution-role.yaml) to create execution role.

# Support, Bugs, Feature Requests

Feature requests can be submitted at [feedback.mongodb.com](https://feedback.mongodb.com/forums/924145-atlas/category/392596-atlas-cloudformation-resources) - just select "Atlas CloudFormation Resources" as the category or vote for an already suggested feature.

Support for the MongoDB Atlas Resource Provider for CloudFormation is provided under MongoDB Atlas support plans, starting with Developer. Please submit support questions within the Atlas UI. In addition, support questions submitted under the Issues section of this repo are also being monitored. Bugs should be filed under the Issues section of this repo.

# MongoDB Atlas API Keys Credential Management
Atlas API keys Configuration are required for both CloudFormation and CDK resources, and this Atlas API key pair are provided as input by the use of a Profile

to configure the API keys and a secret profile follow the next steps:
## 1. Configure your MongoDB Atlas API Keys 
You'll need to generate an API key pair (public and private keys) for your Atlas organization and configure them to grant CloudFormation access to your Atlas project.
Refer to the [Atlas documentation](https://www.mongodb.com/docs/atlas/configure-api-access/#manage-programmatic-access-to-an-organization) for detailed instructions.

## 2. Configure your Profile
To use Atlas CloudFormation resources, you must configure a "profile" with your API keys using [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/).

The secret should follow this format:
```
SecretName: cfn/atlas/profile/{ProfileName}
SecretValue: {PublicKey: {YourPublicKey}, PrivateKey: {YourPrivateKey}}
```

To create a new secret for a default profile, use the [PROFILE SECRET TEMPLATE](/examples/profile-secret.yaml) file provided in this repository.

Here are some examples of how to use this template:

### example 1:
```
  ProfileName: default
  SecretName: cfn/atlas/profile/default
  SecretValue = {PublicKey: xxxxxxx , PrivateKey: yyyyyyyy}
```
### example 2:
```
  ProfileName: tetProfile1
  SecretName: cfn/atlas/profile/tetProfile1
  SecretValue = {PublicKey: zzzzzzzzzz , PrivateKey:jjjjjjjjj}
```

## 3 Provide the profile to your Cloud Formation template

All Atlas CloudFormation resources include a "Profile" property that specifies which profile to use. You'll need to provide the profile you created in the previous step to the CloudFormation template.

Note that if you don't provide a profile, the resource will use a default profile (will try to get a secret named cfn/atlas/profile/default). We recommend always specifying the profile to avoid any unexpected behavior.

Once you've provided the profile, you can deploy the CloudFormation stack using the AWS Console or the AWS CLI. Refer to the AWS documentation for instructions on how to deploy CloudFormation stacks.

IMPORTANT: when specifying the profile in your CloudFormation template, you must specify the Profile Name, NOT the Secret Name

Right:
```
  "Profile" : "ProfileName"
```
Wrong:
```
  "Profile" : "cfn/atlas/profile/ProfileName"
```

# Logging 

Logging for AWS CloudFormation Public extensions is currently disabled. AWS is evaluating if logging is useful for consumers of third party extensions, if this is something you need or would like to request please open a ticket directly with AWS Support. 

# Testing the Provider

Please see README for each resource for details on unit and integrated AWS testing.

## Securing your Deployment 

This repository contains utilities which use both the AWS and MongoDB Atlas APIs to progamatically create, destory, or otherwise manage resources in the AWS and MongoDB clouds. This can have costs or other impact so please use these tools with diligence and follow best practices. 

The automated workflows require the following secrets to be placed into your Github forks' Settings/Secrets:

```
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_REGION
ATLAS_PUBLIC_KEY
ATLAS_PRIVATE_KEY
ATLAS_ORG_ID
```

## Security - Setup 

Step 1) Create and note your MongoDB Atlas API Key.
Step 2) Create and note your AWS Access Key and AWS Secret Key ID.
Step 3) Follow the Github docs on how to [create a Secret](https://docs.github.com/en/actions/configuring-and-managing-workflows/creating-and-storing-encrypted-secrets#creating-encrypted-secrets-for-a-repository) in your clone of this repository.

## Requirements

- [AWS CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli) 
- (Optional - only need if building from source) [AWS CloudFormation CLI Go Plugin](https://github.com/aws-cloudformation/cloudformation-cli-go-plugin/) > v1.0
- (Optional - only need if building from source) [Go](https://golang.org/doc/install) > v1.14 


## Using the MongoDB Atlas CFN Resources 

The fastest way to use the resources is with the official [MongoDB Atlas on AWS Partner Solutions](https://github.com/aws-quickstart/quickstart-mongodb-atlas) CloudFormation templates. The [templates](https://github.com/aws-quickstart/quickstart-mongodb-atlas/templates) folder contains concrete CloudFormation templates you can use to start your own projects.

There are two main parts of this project:

1. [quickstart-mongodbatlas](quickstart-mongodbatlas) This is a mirror of https://github.com/aws-quickstart/quickstart-mongodb-atlas. 

2. [cfn-resources](cfn-resources) A set of AWS CloudFormation custom resource providers for MongoDB Atlas Resources. Currently, AWS requires users to manually deploy these resources in each AWS region one one desires to use them in. We support this workflow through the standard AWS cfn submit tooling. Scripts and Github actions are contained in this repository which demonstrate automating this deployment process.


## Registering resources to run locally 

1. Please check that you satisfy all the [requirements](#Requirements) before proceeding.
2. Clone this repo. 
3. A helper script `[cfn-resources/cfn-submit-helper.sh](cfn-submit-helper.sh)` will build and submit each resource for you. You can also run the `cfn submit` tool yourself. Note- this step will soon not be required when AWS launch a public registry. 

The following command builds and registers all the types used in the quickstart:

```
cd mongodbatlas-cloudformation-resources\cfn-resources
./cfn-submit-helper.sh project cluster database-user project-ip-access-list network-peering
```

## IAM Access Error When Previsioning Resources 
If you are having difficulty with IAM access, suggest try registering first with the following IAM role [here](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/cfn-resources/execute-role.template.yml). This activates the public registry extensions by first using the private registry extensions.

The naming scheme  for a MongoDB Atlas resource on the AWS CloudFormation Third-Party Public Registry is "MongoDB::Atlas::[RESOURCE-NAME]". 

# Common Troubleshooting when using AWS CloudFormation/CDK with MongoDB Atlas Resources  
1. Activate the 3rd party extension for each resource (i.e. MONGODB::ATLAS::[RESOURCE-NAME]) in each AWS region and from each AWS account that you wish to deploy from
2. Ensure you have sufficiently strong AWS IAM Activation Role attached to each 3rd party extension. For sample IAM Role see [here](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/cfn-resources/execute-role.template.yml)
3. Ensure your activated 3rd party public extension matches name exactly to MONGODB::ATLAS::[RESOURCE-NAME] (you may need to delete private extension if this namespace is already occupied)
4. Ensure your MongoDB Atlas Programmatic API Keys (PAKs) being used with CloudFormation have sufficiently strong permissions (Organization Project Creator or Organization Owner)
5. Ensure your MongoDB Atlas PAKs have correct IP Address / CIDR range access. For testing purposes with caution you can open keys to all access by adding “0.0.0.0/1” and “128.0.0.0/1” (do not use for production workloads). 
6. How to determine which IP address AWS CloudFormation uses to deploy MongoDB Atlas resouces with my Atlas Programmatic API Keys (PAK)?
  When you deploy MongoDB Atlas using CloudFormation with your Atlas PAK, CloudFormation will default to use the IP address of the machine from which you are making the API call. The "machine" however making the API call to 3rd party MongoDB Atlas API would be various AWS servers hosting Lambda functions and won't be static. Review additional details at https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html and contact AWS Support directly who can help confirm CIDR range to be used in your Atlas PAK IP Whitelist.  

## Autoclose stale issues and PRs

- After 30 days of no activity (no comments or commits are on an issue or PR) we automatically tag it as “stale” and add a message: "This issue has gone 30 days without any activity and meets the project’s definition of ‘stale’. This will be auto-closed if there is no new activity over the next 30 days. If the issue is still relevant and active, you can simply comment with a “bump” to keep it open, or add the “[Status] Not Stale” label. Thanks for keeping our repository healthy!"

- After 30 more days of no activity we automatically close the issue / PR.

