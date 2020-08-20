# MongoDB::Atlas::EncryptionAtRest

## Description
This resource allows administrators to enable, disable, configure, and retrieve the configuration for Encryption at Rest.

## Parameters
`ProjectId` *(required)* : The unique identifier of the project.<br>
`AwsKms.AccessKeyID` *(required)* : The IAM access key ID with permissions to access the customer master key specified by customerMasterKeyID.<br>
`AwsKms.CustomerMasterKeyID` *(required)* : The AWS customer master key used to encrypt and decrypt the MongoDB master keys.<br>
`AwsKms.Enabled` *(required)* : The IAM secret access key with permissions to access the customer master key specified by customerMasterKeyID.<br>
`AwsKms.SecretAccessKey` *(required)* : Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.<br>
`AwsKms.Region` *(required)* : The AWS region in which the AWS customer master key exists.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas organization or project.<br>

## Installation
    $ make
    $ cfn submit
    ...