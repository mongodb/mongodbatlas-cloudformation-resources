# CFN publishing automation document for AWS System Manager

** *Documents in this directory intended for internal use only by MongoDB, Inc.*

cfn-3P-provider-register.yaml file contains the automation document that is used to publish
MongoDB's CFN resources to AWS public registry.

This automation document builds and registers 3rd party resource providers. This automation document
can register resources in a single AWS account or in multiple accounts and regions. However,
multi-account and regions require a role. Please refer to
the [SSM multi-account ](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-automation-multiple-accounts-and-regions.html)
documentation.

This document builds and registers 3rd party resource providers build using the following plugins:

* JAVA [link to another webpage](https://aws.amazon.com/)
* Go [link to another webpage](https://aws.amazon.com/)
* Python (Coming soon) [link to another webpage](https://aws.amazon.com/)

## Making updates to the automation:

- The .yaml script is a copy of "CFN-MongoDB-Atlas-Resource-Register" automation document that can
  be found in AWS Systems Manager (please reach out to team for access)

To update the script in AWS you will need to create a new version of the document. We aim to keep a version in source control at all times which matches the currently deployed version. Please follow
below instructions:

1. Create a new test document in AWS account with your updated script and test your changes.
2. Upon successful testing, raise a PR with your changes in this repo.
3. Go to AWS publishing account -> Systems Manager -> Shared Resources in left menu -> Documents.
   Then locate `CFN-MongoDB-Atlas-Resource-Register` document under
   and note the current latest version of the document in the dropdown. You will use the version in the next steps
4. On PR approval please ensure to use the following format for commit message when _(squash &)_
   merging your PR:
   `AWS automation update: v<newVersion#> [<brief message about changes>]`.
   - For example, `AWS automation update: v52 [fix build status in step 6]`
5. Version number in #4 would be the same as the new version that you will create of the existing
   main document _(i.e. current latest version+1)_ under Systems Manager.
6. After merge, come back to AWS Systems Manager -> open `CFN-MongoDB-Atlas-Resource-Register` ->
   click on Actions and select `Create a new version`.
7. Update the content of the new version and ensure to select `Set new version as default` before
   saving. Click
   on `Create new version`
