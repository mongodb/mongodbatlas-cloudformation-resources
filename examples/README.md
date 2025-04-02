# MongoDB Atlas CloudFormation templates
MongoDB Atlas CloudFormation simplifies provisioning and management of Atlas features on AWS. You can create YAML/JSON based templates for the service or application architectures you want and have AWS CloudFormation use those templates for quick and reliable provisioning of the services or applications (called “[stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacks.html)”). You can also easily update or replicate the stacks as needed.

This collection of sample templates will help you get started with MongoDB Atlas CloudFormation and quickly build your own templates.

## Prerequisites
### MongoDB Atlas
#### Programmatic API Key
You must [configure API keys](https://www.mongodb.com/docs/atlas/configure-api-access/#std-label-atlas-admin-api-access) to authenticate with your MongoDB Atlas organization.

### AWS

### CloudFormation Profile
You should create a profile in the AWS Secrets Manager that contains the MongoDB Atlas Programmatic API Key.

The secret must be named `cfn/atlas/profile/{ProfileName}` and exist in the same AWS account and AWS region where the Cloudformation stack is run.

Use [this template](profile-secret.yaml) to create a [new CloudFormation stack](https://console.aws.amazon.com/cloudformation/home#/stacks/create) for the default profile that all resources attempt to use unless you specify a different profile.

### Configure the IAM Execution Role
CloudFormation extensions use an *execution role* ([IAM Role](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)) to access AWS resources.
You must configure this role must with at least a policy allowing `secretsmanager:GetSecretValue` to access the configured [profile](../README.md#2-configure-your-profile).

To create the execution role, do one of the following steps:
- Use the [execution-role.yaml](execution-role.yaml) template
- Create the execution-role as part of the [next step](#activate-the-mongodb-atlas-cloudformation-public-extensions)
- Use the [AWS Console](https://us-east-1.console.aws.amazon.com/iam/home?region=us-east-1#/roles) or another method of your choice.

### Activate the MongoDB Atlas CloudFormation public extensions
You have the following options for activating the extensions, for example `MongoDB::Atlas::Cluster`:

- Use the [CloudFormation service in the AWS Console](https://us-east-1.console.aws.amazon.com/cloudformation/home?region=us-east-1#/registry/public-extensions?visibility=PUBLIC&type=RESOURCE&category=THIRD_PARTY):
   1. Ensure you are in the correct AWS region.
   2. Select Publisher=`ThirdParty`.
   3. Extension name prefix = `MongoDB::Atlas`.
   4. Select the resource type and click `Activate`.
   5. In the *Execution role ARN* specify the role arn, for example `arn:aws:iam::123456789012:role/cfn-execution-role`, from the previous step.
- Use the provided [CFN template](activate-mongodb-atlas-resources.template.yaml) to create an IAM execution role and activate all MongoDB Atlas extensions:
   1. Specify the Region.
   2. [Create the stack](https://us-east-1.console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks?filteringText=&filteringStatus=active&viewNested=true). If the stack has been run before in the same region, you will need to delete the previous stack and re-create the stack to re-activate the extensions.

### Configure your KMS Key Policy
If your profile secret is encrypted with a KMS Customer Managed Key (CMK), you must configure the key to allow access.
To give the `execution role` access to the CMK, you must configure a [KMS Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html).
The recommended options include:
- [No `kms` permissions on the `execution role`, allow access via Secrets Manager Service](https://docs.aws.amazon.com/secretsmanager/latest/userguide/security-encryption.html#security-encryption-policies):
  ```json
  {
    "Statement": [
        // other statements
        {
            "Sid": "Enable full access to IAM Admin Users",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::{CFN_ACCOUNT}:root" // can limit to a specific user/group/role for more restrictive access
            },
            "Action": "kms:*",
            "Resource": "*"
        },
        {
            "Sid": "Allow access through AWS Secrets Manager for the CFN Execution Role",
            "Action": [
                "kms:Decrypt"
            ],
            "Effect": "Allow",
            "Principal": {
                "AWS": "*"
            },
            "Resource": "*",
            "Condition": {
                "StringEquals": {
                    "kms:CallerAccount": "{CFN_ACCOUNT}",
                    "kms:ViaService": "secretsmanager.{REGION}.amazonaws.com",
                    "aws:PrincipalArn": "arn:aws:iam::{CFN_ACCOUNT}:role/CFN_EXECUTION_ROLE" // optional clause to limit to your execution role only
                }
            }
        }
    ],
    "Version": "2012-10-17"
  }
  ```
  - replace `{CFN_ACCOUNT}` with your [AWS account ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html)
  - replace `{REGION}` with the CFN region, e.g., `us-east-1`
- Use `kms:decrypt` on the `execution role` and [enable IAM policies](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#key-policy-default-allow-root-enable-iam) on the KMS key:
  ```json
  {
    "Statement": [
        // other statements
        {
            "Sid": "Enable full access to IAM Admin Users",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::{CFN_ACCOUNT}:root" // can limit to a specific user/group/role for more restrictive access
            },
            "Action": "kms:*",
            "Resource": "*"
        },
        {
            "Sid": "Enable Decryption for the CFN Execution Role",
            "Action": "kms:Decrypt",
            "Effect": "Allow",
            "Principal": {
                "AWS": "*"
            },
            "Resource": "*",
            "Condition": { // see comment below
                "StringEquals": {
                    "aws:PrincipalArn": "{EXECUTION_ROLE_ARN}" 
                }
        }
    }
    // role specific access
    ],
    "Version": "2012-10-17"
    }
  ```
  - replace `CFN_ACCOUNT` with your [AWS account ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html)
  - `Condition` can contain more open permissions, for example: `aws:PrincipalOrgID`, `aws:SourceOrgID`, and `aws:SourceAccount`, to learn more see the [AWS User Guide.](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html)
  - ensure the IAM Execution Role has `kms:decrypt` permission:
  ```json
  {
    "Statement": [
        // other statements, e.g., `secretsmanager:GetSecretValue`
        {
            "Action": "kms:Decrypt",
            "Effect": "Allow",
            "Resource": "{KEY_ARN}"
        }
    ],
    "Version": "2012-10-17"
  }
  ```
  - replace `{KEY_ARN}` with your kms key arn

## Upgrade the MongoDB Atlas CloudFormation public extensions
To upgrade an existing `MongoDB::Atlas` resource type you need to: (see [Activate the MongoDB Atlas CloudFormation public extensions](#activate-the-mongodb-atlas-cloudformation-public-extensions) to activate a new resource type):

1. Navigate to the [Cloudformation -> Registry -> Activated extensions](https://eu-north-1.console.aws.amazon.com/cloudformation/home?region=eu-north-1#/registry/activated-extensions/resource-types?category=ACTIVATED&type=RESOURCE).
2. Select the resource type you want to update.
3. Click Actions -> Update (if you enabled Automatic Updates when activating the resource type, any minor, for example `1.0.0` to `1.1.0` will be automaticlly updated).
4. Re run the [stacks](https://eu-north-1.console.aws.amazon.com/cloudformation/home?region=eu-north-1#/stacks?filteringText=&filteringStatus=active&viewNested=true) using the upgraded resource type.

## Using the examples
Once your prerequisites are configured, use the examples in this folder as a starting template for a resource to quickly create a CloudFormation stack. 

For example, the [cluster example](cluster/project-cluster.json) creates a project & cluster in your MongoDB Atlas organization. The template requires the following properties to be configured:
- OrgId - The name of the MongoDB Atlas Organization created earlier
- Profile - The name of the profile that contains your API key information
- ProjectName - The name of your new project
- ClusterName - The name of your new cluster
