# CFN publishing automation document for AWS System Manager

cfn-3P-provider-register.yaml file contains the automation document that is used to publish
MongoDB's CFN resources to AWS public registry.

This automation document builds and registers 3rd party resource providers. This automation document
can register resources in a single AWS account or in multiple accounts and regions. However,
multi-account and regions require a role. Please refer to
the [SSM multi-account ](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-automation-multiple-accounts-and-regions.html)
documentation.

##    

This document builds and registers 3rd party resource providers build using the following plugins:

* JAVA [link to another webpage](https://aws.amazon.com/)
* Go [link to another webpage](https://aws.amazon.com/)
* Python (Coming soon) [link to another webpage](https://aws.amazon.com/)

Documents in this directory intended for use only by MongoDB, Inc.
