# Contributing

MongoDB welcomes community contributions! If you’re interested in making a contribution, please follow the steps below before you start writing any code:

1. Reach out by filing an [issue](https://github.com/mongodb/mongodbatlas-cloudformation-resources/issues) to discuss your proposed contribution, be it a bug fix or feature/other improvements.
1. Sign the [contributor's agreement](http://www.mongodb.com/contributor). This will allow us to review and accept contributions.

After the above 2 steps are completed and we've agreed on a path forward:
1. Fork the repository on GitHub
2. Create a branch with a name that briefly describes your submission
3. Add comments around your new code that explain what's happening
4. Commit and push your changes to your branch then submit a pull request against the current release branch, not master. The naming scheme of the branch is `release-staging-v#.#.#`. **Note**: There will only be one release branch at a time.  
5. A repo maintainer will review the your pull request, and may either request additional changes or merge the pull request.

## Requirements

- Resources configurations, specifcally `.rpdk-config` and `Makefile` build commands, are compatible with cloudformation-cli (`cfn`) versions 0.2.34 and above.


## Code and Test Best Practices

- Each resource is implemented in a seperate directory within `./cfn-resources`. Under each resource directory operations will be implemented in `./cmd/resource/resource.go`, having a separate file `./cmd/resource/mappings.go` for defining conversion logic with respective SDK and CFN models. Associated unit testing files must defined for conversion logic, and can also exist for other business logic such as handling state transitions.
- [Testify Mock](https://pkg.go.dev/github.com/stretchr/testify/mock) and [Mockery](https://github.com/vektra/mockery) are used for test doubles in unit tests. Mocked interfaces are generated in folder `cfn-resources/testutil/mocksvc`.
- We have a `/test/README.md` for every resource in `cfn-resources`. You will also find [TESTING.md](./TESTING.md) which provides testing practices common to all resources.

Please follow below guidelines for testing to ensure quality:
**When adding a new feature:**
- `/test/README.md` file must be created and must include detailed pre-requisites (if any) and steps to test a resource.
- The file should include screenshots of how the resource is tested and how the changes reflect in Atlas UI.
- See for example, [/test/README.md for LDAP Configuration]( https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/cfn-resources/ldap-configuration/test/README.md) and [/test/README.md for Cloud Backup Snapshot](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/cfn-resources/cloud-backup-snapshot/test/README.md).

**When creating pull request:**
- Please include screenshots of any testing performed in the description. Following screenshots should be included:
  - Stack creation with the resource in AWS console showing successful create, update and delete operations.
  - Corresponding change reflected in Atlas UI.
  - Successful contract testing when run locally.
  - See [#669](https://github.com/mongodb/mongodbatlas-cloudformation-resources/pull/669) for example pull request.


## PR Title Format
We use [*Conventional Commits*](https://www.conventionalcommits.org/):
- `fix: description of the PR`: a commit of the type fix patches a bug in your codebase (this correlates with PATCH in Semantic Versioning).
- `chore: description of the PR`: the commit includes a technical or preventative maintenance task that is necessary for managing the product or the repository, but it is not tied to any specific feature or user story (this correlates with PATCH in Semantic Versioning).
- `doc: description of the PR`: The commit adds, updates, or revises documentation that is stored in the repository (this correlates with PATCH in Semantic Versioning).
- `test: description of the PR`: The commit enhances, adds to, revised, or otherwise changes the suite of automated tests for the product (this correlates with PATCH in Semantic Versioning).
- `security: description of the PR`: The commit improves the security of the product or resolves a security issue that has been reported (this correlates with PATCH in Semantic Versioning).
- `refactor: description of the PR`: The commit refactors existing code in the product, but does not alter or change existing behavior in the product (this correlates with Minor in Semantic Versioning).
- `perf: description of the PR`: The commit improves the performance of algorithms or general execution time of the product, but does not fundamentally change an existing feature (this correlates with Minor in Semantic Versioning).
- `ci: description of the PR`: The commit makes changes to continuous integration or continuous delivery scripts or configuration files (this correlates with Minor in Semantic Versioning).
- `revert: description of the PR`: The commit reverts one or more commits that were previously included in the product, but were accidentally merged or serious issues were discovered that required their removal from the main branch (this correlates with Minor in Semantic Versioning).
- `style: description of the PR`: The commit updates or reformats the style of the source code, but does not otherwise change the product implementation (this correlates with Minor in Semantic Versioning).
- `feat: description of the PR`: a commit of the type feat introduces a new feature to the codebase (this correlates with MINOR in Semantic Versioning).
- `deprecate: description of the PR`: The commit deprecates existing functionality, but does not remove it from the product (this correlates with MINOR in Semantic Versioning).
- `BREAKING CHANGE`: a commit that has a footer BREAKING CHANGE:, or appends a ! after the type/scope, introduces a breaking API change (correlating with MAJOR in Semantic Versioning). A BREAKING CHANGE can be part of commits of any type.
Examples:
  - `fix!: description of the ticket`
  - If the PR has `BREAKING CHANGE`: in its description is a breaking change
- `remove!: description of the PR`: The commit removes a feature from the product. Typically features are deprecated first for a period of time before being removed. Removing a feature is a breaking change (correlating with MAJOR in Semantic Versioning).

## Discovering New API features

Most of the new features of the provider are using [atlas-sdk](https://github.com/mongodb/atlas-sdk-go)
SDK is updated automatically, tracking all new Atlas features.

### Updating Atlas SDK

To update Atlas SDK run:

```bash
cd cfn-resources/
make update-atlas-sdk
```

> NOTE: Update mechanism is only needed for major releases. Any other releases will be supported by dependabot.

> NOTE: Command can make import changes to +500 files. Please make sure that you perform update on main branch without any uncommited changes.

### SDK Major Release Update Procedure

1. If the SDK update doesn’t cause any compilation issues create a new SDK update PR
  1. Review [API Changelog](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/changelog) for any deprecated fields and breaking changes.
2. For SDK updates introducing compilation issues without graceful workaround
  1. Use the previous major version of the SDK (including the old client) for the affected resource
  1. Create an issue to identify the root cause and mitigation paths based on changelog information
  2. If applicable: Make required notice/update to the end users based on the plan.

## Autoclose stale issues and PRs
- After 30 days of no activity (no comments or commits are on an issue or PR) we automatically tag it as “stale” and add a message: "This issue has gone 30 days without any activity and meets the project’s definition of ‘stale’. This will be auto-closed if there is no new activity over the next 30 days. If the issue is still relevant and active, you can simply comment with a “bump” to keep it open, or add the “[Status] Not Stale” label. Thanks for keeping our repository healthy!"
- After 30 more days of no activity we automatically close the issue / PR.

## Prerequisites
- [AWS CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli)
- (Optional - only need if building from source) [AWS CloudFormation CLI Go Plugin](https://github.com/aws-cloudformation/cloudformation-cli-go-plugin/) > v1.0
- (Optional - only need if building from source) [Go](https://golang.org/doc/install) v1.20

## Testing the Provider
Please see README for each resource for details on unit and integrated AWS testing.

## Security - Setup
1. Create and note your [MongoDB Atlas API Key](https://www.mongodb.com/docs/atlas/configure-api-access/#std-label-atlas-admin-api-access).
2. Create and note your AWS Access Key and AWS Secret Key ID.
3. Follow the Github docs on how to [create a Secret](https://docs.github.com/en/actions/configuring-and-managing-workflows/creating-and-storing-encrypted-secrets#creating-encrypted-secrets-for-a-repository) in your clone of this repository.

## Securing your Deployment
This repository contains utilities which use both the AWS and MongoDB Atlas APIs to progamatically create, destory, or otherwise manage resources in the AWS and MongoDB clouds. This can have costs or other impact so please use these tools with diligence and follow best practices.

The automated workflows require the following secrets to be placed into your Github fork's Settings/Secrets:

```
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_REGION
MONGODB_ATLAS_PUBLIC_KEY
MONGODB_ATLAS_PRIVATE_KEY
MONGODB_ATLAS_ORG_ID
```

## Using the MongoDB Atlas CFN Resources

The fastest way to use the resources is with the official [MongoDB Atlas on AWS Partner Solutions](https://aws.amazon.com/solutions/partners/mongodb-atlas/) CloudFormation templates. The [templates](https://github.com/aws-quickstart/quickstart-mongodb-atlas/tree/main/templates) folder contains concrete CloudFormation templates you can use to start your own projects.

There are two main parts of this project:

1. [quickstart-mongodbatlas](quickstart-mongodbatlas) This is a mirror of https://github.com/aws-quickstart/quickstart-mongodb-atlas.

2. [cfn-resources](cfn-resources) A set of AWS CloudFormation custom resource providers for MongoDB Atlas Resources. Currently, AWS requires users to manually deploy these resources in each AWS region that they want to use them in. We support this workflow through the standard AWS `cfn submit` tooling. Scripts and Github actions are contained in this repository that demonstrate automating this deployment process.


## Registering resources to run locally

1. Please check that you satisfy all the [requirements](#Requirements) before proceeding.
2. Clone this repo.
3. Use the helper script to `[cfn-resources/cfn-submit-helper.sh](cfn-submit-helper.sh)` will build and submit each resource for you. You can also run the `cfn submit` tool yourself. Note- this step will soon not be required when AWS launch a public registry.

The following command builds and registers all the types used in the quickstart:

```
cd mongodbatlas-cloudformation-resources\cfn-resources
./cfn-submit-helper.sh project cluster database-user project-ip-access-list network-peering
```

## IAM Access Error When Previsioning Resources
If you are having difficulty with IAM access, try registering first with the following IAM role [here](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/cfn-resources/execute-role.template.yml). This activates the public registry extensions by first using the private registry extensions.

The naming scheme for a MongoDB Atlas resource on the AWS CloudFormation Third-Party Public Registry is `MongoDB::Atlas::[RESOURCE-NAME]`.
