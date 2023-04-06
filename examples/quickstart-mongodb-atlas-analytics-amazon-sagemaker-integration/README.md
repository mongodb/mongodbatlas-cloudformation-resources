# quickstart-mongodb-atlas-analytics-amazon-sagemaker-integration



## Overview

![simple-quickstart-arch](https://user-images.githubusercontent.com/5663078/229119386-0dbc6e30-a060-465e-86dd-f89712b0fc49.png)

This Partner Solutions template enables you to begin working with your machine learning models using your MongoDB Atlas Cluster and Amazon SageMaker Endpoint as tools. With this template, you can utilize MongoDB as a data source and Amazon SageMaker for data analysis, streamlining the process of building and deploying machine learning models.



## MongoDB Atlas CFN Resources used by the templates

- [MongoDB::Atlas::Trigger](../../cfn-resources/trigger/)


## Environment Configured by the Partner Solutions template
The Partner Solutions template will generate and configure the following resources:
 - a [MongoDB Partner Event Bus](http://mongodb.com/docs/atlas/app-services/triggers/aws-eventbridge/#std-label-aws-eventbridge)
 - a [database trigger](https://www.mongodb.com/docs/atlas/app-services/triggers/database-triggers/) with your Atlas Cluster
 - lambda functions to run the machine learning model and send the classification results to your MongoDB Atlas Cluster. (See [iris_classifier](sagemaker-example/iris_classifier/) for an example of machine learning model to use with this template. See [lambda_functions](sagemaker-example/lambda_functions/) for an example of lambda functions to use to read and write data to your MongoDB Atlas cluster.)


## Permissions required to run the template
See [execution-role.yaml](execution-role.yaml) for the list of permissions needed to run the template. 


If you want to create an IAM (Identity and Access Management) role to run the quickstart template on a CFN (CloudFormation) stack, you can use the [execution-role.yaml](execution-role.yaml) file. Note that passing an IAM role to CloudFormation when creating a stack is optional. If you don't supply one, the user permissions are assumed. 
See the [IAM permissions section in the General information guide](https://aws-ia.github.io/content/qs_info.html#_technical_requirements) for more information.


## Additional Information

- Repository: [quickstart-mongodb-atlas-analytics-amazon-sagemaker-integration](https://github.com/aws-quickstart/quickstart-mongodb-atlas-analytics-amazon-sagemaker-integration)
- Template: [mongodb-sagemaker-analytics-main.template.yaml](https://github.com/aws-quickstart/quickstart-mongodb-atlas-analytics-amazon-sagemaker-integration/blob/main/templates/mongodb-sagemaker-analytics-main.template.yaml)
- [What Is Amazon EventBridge? - Amazon EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-what-is.html)
- [What is an event bus](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventBus.html): An event bus receives events from a source, uses rules to evaluate them, applies any configured input transformation, and routes them to the appropriate target(s). Your account's default event bus receives events from AWS services. A custom event bus can receive events from your custom applications and services.

