# quickstart-mongodb-atlas-mean-stack-aws-fargate-integration



## Overview

![simple-quickstart-arch](https://user-images.githubusercontent.com/5663078/229105149-59015114-1c14-44e3-ad5a-b48d9a487797.png)

This Partner Solutions template provides the architecture necessary to scale a [MEAN](https://www.mongodb.com/mean-stack) (MongoDB, Express, Angular, Node.js) stack application using a combination of CloudFormation, MongoDB Atlas, and AWS Fargate. The template leverages the MongoDB Atlas CFN resources to configure the MongoDB infrastructure and AWS CFN resources to set up an Application Load Balancer and a VPC. Additionally, the template employs AWS Fargate to run your Docker image.



## MongoDB Atlas CFN Resources used by the templates

- [MongoDB::Atlas::Cluster](../../cfn-resources/cluster/)
- [MongoDB::Atlas::ProjectIpAccessList](../../cfn-resources/project-ip-access-list/)
- [MongoDB::Atlas::DatabaseUser](../../cfn-resources/database-user/)
- [MongoDB::Atlas::Project](../../cfn-resources/project/)
- [MongoDB::Atlas::NetworkPeering](../../cfn-resources/network-peering/)
- [MongoDB::Atlas::NetworkContainer](../../cfn-resources/network-container/)
- [MongoDB::Atlas::PrivateEndpoint](../../cfn-resources/private-endpoint/)


## Environment Configured by the Partner Solutions template
The Partner Solutions template will generate the following resources:
 - A virtual private cloud (VPC) configured with public and private subnets, according to AWS best practices, to provide you with your own virtual network on AWS. The VPC provides Domain Name System (DNS) resolution. The template leverages the [official AWS quickstart template](https://github.com/aws-quickstart/quickstart-aws-vpc/blob/9dc47510f71f1fb6baf8c4e96b5330a6f51f540e/templates/aws-vpc.template.yaml) to build your VPC infrastructure. See [Deployment Guide](https://aws-quickstart.github.io/quickstart-aws-vpc/) for more information.
- An Atlas Project in the organization that was provided as input.
- An Atlas Cluster with authentication and authorization enabled, and not accessible through the public internet.
- A Database user with access to the Atlas Cluster.
- An Atlas IP access list, allowing the cluster to be accessed through the public internet.
- A VPC peering connection between the MongoDB Atlas VPC (where the cluster is located) and the VPC on AWS.
- An application Load Balancer 
- AWS Fargate to run your Docker image. See [fargate-example/](fargate-example/) for an example of docker images to use with Fargate.


## Permissions required to run the template
See [execution-role.yaml](execution-role.yaml) for the list of permissions needed to run the template. 


If you want to create an IAM (Identity and Access Management) role to run the Partner Solutions template on a CFN (CloudFormation) stack, you can use the [execution-role.yaml](execution-role.yaml) file. Note that passing an IAM role to CloudFormation when creating a stack is optional. If you don't supply one, the user permissions are assumed. 
See the [IAM permissions section in the General information guide](https://aws-ia.github.io/content/qs_info.html#_technical_requirements) for more information.


## Additional Information

- Repository: [quickstart-mongodb-atlas-mean-stack-aws-fargate-integration](https://github.com/aws-quickstart/quickstart-mongodb-atlas-mean-stack-aws-fargate-integration/tree/main)
- Template: [mongodb-atlas-mean-stack.template.yaml](https://github.com/aws-quickstart/quickstart-mongodb-atlas-mean-stack-aws-fargate-integration/blob/main/templates/mongodb-atlas-mean-stack.template.yaml)
- [What Is The MEAN Stack? Introduction & Examples | MongoDB](https://www.mongodb.com/mean-stack)
- [Serverless Compute Engine–AWS Fargate–Amazon Web Services](https://aws.amazon.com/fargate/)

