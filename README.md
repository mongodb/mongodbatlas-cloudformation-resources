# MongoDB Atlas AWS CloudFormation Resources & Quickstart

***This file will be accurate post GA of the MongoDB Atlas Resource Provider for CloudFormation***

### Status: pre-BETA (actively looking for [feedback](https://feedback.mongodb.com/forums/924145-atlas/category/392596-atlas-cloudformation-resources))

Use AWS CloudFormation to manage [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).

The quickest way to get started:

1. Clone this repository.
2. Create Github Secrets for your AWS and MongoDB accounts. See [Securing your Quickstart](#securing-your-quickstart).
3. Run the [Launch Quickstart](https://github.com/mongodb/mongodbatlas-cloudformation-resources/actions) Github Action.

If you prefer to run things locally, see [Using the Atlas CFN Resources](#using-the-atlas-cfn-resources).

# Support, Bugs, Feature Requests

Feature requests can be submitted at [feedback.mongodb.com](https://feedback.mongodb.com/forums/924145-atlas/category/392596-atlas-cloudformation-resources) - just select "Atlas CloudFormation Resources" as the category or vote for an already suggested feature.

CURRENT BETA INFO: Support questions submitted under the Issues section of this repo will be handled on a "best effort" basis.
Bugs should be filed under the Issues section of this repo.

POST GA: Support for the MongoDB Atlas Resource Provider for CloudFormation is provided under MongoDB Atlas support plans, starting with Developer. Please submit support questions within the Atlas UI. 

# Securing your Quickstart

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

# Requirements

- [AWS CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli) 0.1.3
- (Optional - only need if building from source) [AWS CloudFormation CLI Go Plugin](https://github.com/aws-cloudformation/cloudformation-cli-go-plugin/) 0.1.6
- (Optional - only need if building from source) [Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)

# Using the Atlas CFN Resources 

This project contains 2 main items:

1. [quickstart-mongodbatlas](quickstart-mongodbatlas)tA sample quickstart AWS CloudFormation template to launch a full MongoDB Atlas deployment. This template uses the Atlas custom CFN resources. It is an example and starting point for your own CloudFormation projects using MongoDB Atlas.

2. [cfn-resources](cfn-resources) A set of AWS CloudFormation custom resource providers for MongoDB Atlas Resources. Currently, AWS requires users to manually deploy these resources manually in each AWS region of need. We support this workflow through the standard AWS `cfn submit` tooling, and there are scripts and Github actions which demonstrate automating this process.   

## Running the Github workflows locally

At this time, the [act]() tool doesn't support the `ubuntu-20.04` image as a local runner, so our actions won't running easily out of the box locally yet. 

You can build and run the action to deploy like this:

```bash
docker run -v mongodbatlas-cloudformation-resources/cfn-resources:/atlas-cfn/cfn-resources --env-file local.env -t jmimick/atlas-cfn-deploy
```
## Registering resources 

These are the detailed steps which are automated in the atlas-cfn-deploy Github workflow found in this repository.
You can use these steps, or leverage the [cfn-resources/utils/atlas-cfn-deploy](cfn-resources/utils/atlas-cfn-deploy) tool.

1. Please check that you satisfy all the [requirements](#Requirements) before proceeding.
2. Clone this repo, or head over to [releases](https://github.com/mongodb/mongodbatlas-cloudformation-resources/releases) and download the binary for the most recent release, `mongodbatlas-cloudformation-resources_<version>_Linux_amd64.tar.gz`
3. If needed, extract the tarball with
```
tar -xvzf mongodbatlas-cloudformation-resources_<version>_Linux_amd64.tar.gz
```
3. Once extracted, navigate to the resource that you are trying to build, eg `./project`
4. Run the following command to register the resource provider with CloudFormation in the specified region: 
  ```
  cfn submit -v --region <region>
  ```
  - This may take a few minutes.
  - Additional details about each resource can be found in their respective READMEs.
  - See [AWS docs](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-cli-submit.html) for additional options for `cfn submit` if needed.
  - You may need additional IAM permissions to register a resource provider, please see [AWS docs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html#registry-register-permissions). 
4. When the registration is successful, you will see `Registration complete` in your terminal. You will also be able to see it in the CloudFormation Stacks console.
5. Repeat these steps for any resource you are trying to build.

## Creating stacks
To use a resource provider, you need a template, templates of this project are available in `./examples`, then you need to run commands, you can read about available commands in [CloudFormation commands](https://docs.aws.amazon.com/cli/latest/reference/cloudformation/index.html)

To create a stack of a resource
```
$ aws cloudformation create-stack --stack-name myTestProject --template-body file://../examples/project/project.json --parameters ParameterKey=Name,ParameterValue=test-project
```
If errors are not shown, it should return a stack id

To verify if it's working you can check in the [Cloudformation console](https://console.aws.amazon.com/cloudformation)

# Developing the CloudFormation Resource Provider

If you wish to work on the CloudFormation resource provider, you'll first need [Go](https://golang.org/doc/install) installed on your machine (please check the [requirements](#Requirements) before proceeding).

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

# Testing the Provider

In order to test a provider resource, there are three ways to do that.

1. ### Using the AWS CloudFormation Console

   After creating, updating or deleting a stack of a provider resource in the AWS CloudFormation console, you can check the [AWS CloudFormation](https://console.aws.amazon.com/cloudformation) for events which will show the status of a resource provider.

   After uploading a stack logs are automatically created and may be viewed in [AWS CloudWatch](https://console.aws.amazon.com/cloudwatch) `Log groups`.

   To print a log you need to add it in code `log.Printf("")`.

2.  ### Using the AWS CloudFormation CLI

    After creating, updating or deleting a stack of a provider resource with the AWS CLI, you can use commands to see events, describe stacks and more. See [CloudFormation CLI commands](https://docs.aws.amazon.com/cli/latest/reference/cloudformation/index.html) for more.

    After uploading a stack logs are automatically created and may be viewed in [AWS CloudWatch](https://aws.amazon.com/cloudwatch/) `Log groups`.

    To retrieve logs you need the name of the `log group` and in some situations you need the name of the `log stream`. You can see more [here](https://aws.amazon.com/cloudwatch/) in `Log groups`, once there use [logs commads](https://docs.aws.amazon.com/cli/latest/reference/logs/index.html)


3.  ### Using the Contract Test

    The resources of this repository was created using CloudFormation best practices. Therefore, it will pass all of the CloudFormation CLI contract tests.
    To run the contrast test, follow the directions in the
    [documentation](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-cli-test.html).
