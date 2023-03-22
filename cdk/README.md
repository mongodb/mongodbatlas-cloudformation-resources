# MongoDB Atlas AWS CDK Resources
Use MongoDB Atlas [AWS CDK](https://aws.amazon.com/cdk/) (or Cloud Development Kit) constructs to manage [MongoDB Atlas](https://www.mongodb.com/cloud/atlas). The AWS CDK is a framework for defining infrastructure as code (IaC). It allows developers to write code in their preferred programming language, such as TypeScript for example, to define and deploy infrastructure. AWS CDK gets synthesized down into [AWS CloudFormation](https://aws.amazon.com/cloudformation/) templates at deployment so users no longer have to write or maintain YAML/JSON based CloudFormation templates. 

## Getting Started
See the [cdk examples](examples/README.md) for how to setup prerequisites & get started with your first cluster, using our AWS CDK sample code.

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

## Contributing
See our [CONTRIBUTING.md](CONTRIBUTING.md) guide.
