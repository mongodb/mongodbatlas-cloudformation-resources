# MongoDB Atlas Cloudformation Resources

This is the repository for the Terraform MongoDB Atlas CloudFormation, which allows one to use Cloudformation with MongoDB's Database as a Service offering, Atlas.
Learn more about Atlas at  [https://www.mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas)

For general information about CloudFormation, visit the [official website](https://aws.amazon.com/cloudformation) and the [GitHub project page](https://github.com/aws-cloudformation/).

# Support, Bugs, Feature Requests

Support for the Terraform MongoDB Atlas CloudFormation is provided under MongoDB Atlas support plans.   Please submit support questions within the Atlas UI.  Support questions submitted under the Issues section of this repo will be handled on a "best effort" basis.  

Bugs should be filed under the Issues section of this repo.

Feature requests can be submitted at https://feedback.mongodb.com/forums/924145-atlas - just select the Terraform plugin as the category or vote for an already suggested feature.

# Requirements
- [AWS CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli) 0.12+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

# Developing the cloudformation

If you wish to work on the cloudformation, you'll first need [Go](https://golang.org/doc/install) installed on your machine (please check the [requirements](#Requirements) before proceeding).

Note: This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](https://golang.org/doc/code.html#GOPATH). The instructions that follow assume a directory in your home directory outside of the standard GOPATH (i.e $HOME/development/).

Clone repository to: `$HOME/development/`

```
$ cd $HOME/development/
$ git clone https://github.com/mongodb/mongodbatlas-cloudformation-resources.git
...
```

To develop a resource provider, you need to follow steps from [AWS CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli), you can read more info about developing a resource provider [here](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).

To create a resource, you need to create a folder first then generate files with `cfn init`

```
$ cd $HOME/development/mongodbatlas-cloudformation-resources
$ mkdir resource
$ cd resource
$ cfn init
...
```

# Using the cloudformation resource

To use a resource provider, you need a template, templates of this project are available on `/examples` , then you need to run commands, you can read more commands available for [AWS Cloudformation commands](https://docs.aws.amazon.com/cli/latest/reference/cloudformation/index.html)

To create a stack of a resource provider

```
$ aws cloudformation create-stack --stack-name myTestProject --template-body file://../examples/project/project.json --parameters ParameterKey=Name,ParameterValue=test-project
```
If create is successful, it should return a stack id

To verify if it's working you can check on [AWS Cloudformation](https://console.aws.amazon.com/cloudformation)

# Testing the Provider

In order to test a provider resource, there two ways to do that.

1. Using AWS website

   After creating or updating or deleting a stack of a provider resource, you need to check on [AWS Cloudformation](https://console.aws.amazon.com/cloudformation) for events which shows a status of a resource provider.
   
   To see logs, if create is successful automatically creates a log for [CloudWatch](https://console.aws.amazon.com/cloudwatch) in `Log groups`.
   
   To print a log you need to add it in code `log.Print("")`.
    
   **Advantages:**
   
   Making a good example with various resources, it can create all of it with one command.
   
   **Disadvantages:** 
   
   You cannot test locally and it takes more time for results.
   
2. Using AWS SAM for local (WIP)

