Contributing
---------------------------
## Workflow
MongoDB welcomes community contributions! If you’re interested in making a contribution, please follow the steps below before you start writing any code:

1. Reach out by filing an [issue](https://github.com/mongodb/mongodbatlas-cloudformation-resources/issues) to discuss your proposed contribution, be it a bug fix or feature/other improvements.
1. Sign the [contributor's agreement](http://www.mongodb.com/contributor). This will allow us to review and accept contributions.

After the above 2 steps are completed and we've agreed on a path forward:
1. Fork the repository on GitHub
2. Create a branch with a name that briefly describes your submission
3. Add comments around your new code that explain what's happening
4. Commit and push your changes to your branch then submit a pull request against the current release branch, not master.  The naming scheme of the branch is `release-staging-v#.#.#`. Note: There will only be one release branch at a time.  
5. A repo maintainer will review the your pull request, and may either request additional changes or merge the pull request.

## Autoclose stale issues and PRs
- After 30 days of no activity (no comments or commits are on an issue or PR) we automatically tag it as “stale” and add a message: "This issue has gone 30 days without any activity and meets the project’s definition of ‘stale’. This will be auto-closed if there is no new activity over the next 30 days. If the issue is still relevant and active, you can simply comment with a “bump” to keep it open, or add the “[Status] Not Stale” label. Thanks for keeping our repository healthy!"
- After 30 more days of no activity we automatically close the issue / PR.

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
