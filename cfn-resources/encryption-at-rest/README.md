# MongoDB::Atlas::EncryptionAtRest

## Description
This resource allows administrators to enable, disable, configure, and retrieve the configuration for Encryption at Rest.

## Parameters
`ProjectId` *(required)* : The unique identifier of the project.<br>
`AwsKms.RoleId` *(required)* : ID of an AWS IAM role authorized to manage an AWS customer master key.<br>
`AwsKms.CustomerMasterKeyID` *(required)* : The AWS customer master key used to encrypt and decrypt the MongoDB master keys.<br>
`AwsKms.Enabled` *(required)* : The IAM secret access key with permissions to access the customer master key specified by customerMasterKeyID.<br>
`AwsKms.Region` *(required)* : The AWS region in which the AWS customer master key exists.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit
    ...