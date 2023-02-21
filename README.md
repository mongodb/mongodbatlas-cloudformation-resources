# MongoDB Atlas AWS CloudFormation Resources & AWS Partner Solution Deployments 

Use AWS CloudFormation to manage [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).

Partner Solutions (formally AWS Quick Starts) are automated reference deployments built by Amazon Web Services (AWS) solutions architects and AWS Partners. Partner Solutions help you deploy popular technologies to AWS according to AWS best practices. The quickest way to get started is to launch the official [MongoDB Atlas on AWS](https://aws.amazon.com/quickstart/architecture/mongodb-atlas/) Partner Solution Deployment directly from the AWS Management Console.

# Prerequisites

To use CloudFormation MongoDB Atlas public extensions (AWS Third Party), you have to activate the Public extension from your AWS console. You have to do this in each AWS Account and in each AWS Region. To activate Public extension, we have to create an execution role and pass the ARN of the role as an input. Use [this template](examples/execution-role.yaml) to create execution role.

# Support, Bugs, Feature Requests

Feature requests can be submitted at [feedback.mongodb.com](https://feedback.mongodb.com/forums/924145-atlas/category/392596-atlas-cloudformation-resources) - just select "Atlas CloudFormation Resources" as the category or vote for an already suggested feature.

Support for the MongoDB Atlas Resource Provider for CloudFormation is provided under MongoDB Atlas support plans, starting with Developer. Please submit support questions within the Atlas UI. In addition, support questions submitted under the Issues section of this repo are also being monitored. Bugs should be filed under the Issues section of this repo.

# MongoDB Atlas Programmatic API Key
It's necessary to generate and configure an API key for your organization for the acceptance test to succeed. To grant programmatic access to an organization or project using only the API you need to know:

The programmatic API key has two parts: a Public Key and a Private Key. To see more details on how to create a programmatic API key visit https://docs.atlas.mongodb.com/configure-api-access/#programmatic-api-keys.

The programmatic API key must be granted roles sufficient for the acceptance test to succeed. The Organization Owner and Project Owner roles should be sufficient. You can see the available roles at https://docs.atlas.mongodb.com/reference/user-roles.

You must configure Atlas API Access for your programmatic API key. You should allow API access for the IP address from which the acceptance test runs.

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

## Autoclose stale issues and PRs

- After 30 days of no activity (no comments or commits are on an issue or PR) we automatically tag it as “stale” and add a message: "This issue has gone 30 days without any activity and meets the project’s definition of ‘stale’. This will be auto-closed if there is no new activity over the next 30 days. If the issue is still relevant and active, you can simply comment with a “bump” to keep it open, or add the “[Status] Not Stale” label. Thanks for keeping our repository healthy!"

- After 30 more days of no activity we automatically close the issue / PR.

