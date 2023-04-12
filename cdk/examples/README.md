#  MongoDB Atlas AWS CDK construct examples
This folder contains examples of using MongoDB Atlas AWS CDK constructs that will help you get started and quickly build your own applications.

## Prerequisites
### MongoDB Atlas
#### Programmatic API Key
You must [configure API keys](https://www.mongodb.com/docs/atlas/configure-api-access/#std-label-atlas-admin-api-access) to authenticate with your MongoDB Atlas organization.

### AWS
#### Activate the MongoDB Atlas CloudFormation public extensions
To activate a public extension, create an execution role and pass the ARN of the role as an input. Use [this template](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) to create a [new CloudFormation stack](https://console.aws.amazon.com/cloudformation/home#/stacks/create) to create the execution role.

You must then activate the Public extension from your AWS console. You have to do this in each AWS Account and in each AWS Region. Use [this link](https://us-east-1.console.aws.amazon.com/cloudformation/home#/registry/public-extensions?visibility=PUBLIC&type=RESOURCE&category=AWS_TYPES) to register extensions on CloudFormation.

#### CDK Application
Consult the [AWS CDK | Getting Started](https://docs.aws.amazon.com/cdk/v2/guide/getting_started.html) guide for full details. In general, you will need to:
* Install & configure [awscli](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html) 
  * Configure AWS credentials
  * Set your default region
* Configure your programming environment. 
  * Examples in this folder assume Typescript
* Install the AWS CDK
  * `npm install -g aws-cdk`
* [Bootstrap](https://docs.aws.amazon.com/cdk/v2/guide/bootstrapping.html) your AWS account to enable CDK deployments
  * `cdk bootstrap aws://ACCOUNT-NUMBER/REGION`

#### CloudFormation Profile
A profile should be created in the AWS Secrets Manager, containing the MongoDB Atlas Programmatic API Key.

Use [this template](../../examples/profile-secret.yaml) to create a [new CloudFormation stack](https://console.aws.amazon.com/cloudformation/home#/stacks/create) for the default profile that all resources will attempt to use unless a different override is specified.

## Using the examples
Once your prerequisites are configured, use the examples in this folder as a starting template for a resource to quickly create a new AWS cdk application.

For example, the [createCluster example](createCluster.ts) creates a project & cluster in your MongoDB Atlas organization, using a Typescript application. The example requires the following to be configured:
* A new CDK application, in a folder named `cdk-testing`
  * `cdk init app --language typescript`
* MongoDB Atlas AWS CDK libraries for project & cluster constructs
  * `npm install @mongodbatlas-awscdk/cluster`
  * `npm install @mongodbatlas-awscdk/project`
* Replace the generated `lib/cdk-testing-stack.ts` file with the [createCluster example](createCluster.ts) content
* Your application to be built successfully
  * `npm run build`
* You can deploy and monitor the progress using the cdk tooling, with parameters configured through the CDK [runtime context](https://docs.aws.amazon.com/cdk/v2/guide/context.html)
`* `cdk deploy --context orgId=<ORG_ID>`
