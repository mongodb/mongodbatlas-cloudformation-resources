# MongoDB Atlas CloudFormation templates
MongoDB Atlas CloudFormation simplifies provisioning and management of Atlas features on AWS. You can create templates for the service or application architectures you want and have AWS CloudFormation use those templates for quick and reliable provisioning of the services or applications (called “stacks”). You can also easily update or replicate the stacks as needed.

This collection of sample templates will help you get started with MongoDB Atlas CloudFormation and quickly build your own templates

## Prerequisites
### MongoDB Atlas
#### Programmatic API Key
You must [configure API keys](https://www.mongodb.com/docs/atlas/configure-api-access/#std-label-atlas-admin-api-access) to authenticate with your MongoDB Atlas organization.

### AWS
#### Activate the MongoDB Atlas CloudFormation public extensions
To activate a public extension, create an execution role and pass the ARN of the role as an input. Use [this template](execution-role.yaml) to create a [new CloudFormation stack](https://console.aws.amazon.com/cloudformation/home#/stacks/create) to create the execution role.

You must then activate the Public extension from your AWS console. You have to do this in each AWS Account and in each AWS Region. Use [this link](https://us-east-1.console.aws.amazon.com/cloudformation/home#/registry/public-extensions?visibility=PUBLIC&type=RESOURCE&category=AWS_TYPES) to register extensions on CloudFormation.

#### CloudFormation Profile
A profile should be created in the AWS Secrets Manager, containing the MongoDB Atlas Programmatic API Key.

Use [this template](profile-secret.yaml) to create a [new CloudFormation stack](https://console.aws.amazon.com/cloudformation/home#/stacks/create) for the default profile that all resources will attempt to use unless a different override is specified.

## Using the examples
Once your prerequisites are configured, use the examples in this folder as a starting template for a resource to quickly create a CloudFormation stack. 

For example, the [cluster example](cluster/cluster.json) creates a project & cluster in your MongoDB Atlas organization. The template requires the following properties to be configured:
* OrgId - The name of the MongoDB Atlas Organization created earlier
* Profile - The name of the profile that contains your API key information
* ProjectName - The name of your new project
* ClusterName - The name of your new cluster
