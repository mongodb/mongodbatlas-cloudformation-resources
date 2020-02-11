# MongoDB::Atlas::EncryptionAtRest

Congratulations on starting development!

Next steps:

1. Populate the JSON schema describing your resource, `mongodb-atlas-encryptionatrest.json`
2. The RPDK will automatically generate the correct resource model from the
   schema whenever the project is built via Make.
   You can also do this manually with the following command: `cfn-cli generate`
3. Implement your resource handlers by adding code to provision your resources in your resource handler's methods.

Please don't modify files `model.go` and `main.go`, as they will be automatically overwritten.

## Attributes
`ProjectId` *(required)* : The unique identifier of the project.<br>
`AwsKms.AccessKeyID` *(required)* : The IAM access key ID with permissions to access the customer master key specified by customerMasterKeyID.<br>
`AwsKms.CustomerMasterKeyID` *(required)* : he AWS customer master key used to encrypt and decrypt the MongoDB master keys.<br>
`AwsKms.Enabled` *(required)* : he IAM secret access key with permissions to access the customer master key specified by customerMasterKeyID.<br>
`AwsKms.SecretAccessKey` *(required)* : Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.<br>
`AwsKms.Region` *(required)* : The AWS region in which the AWS customer master key exists.<br>
`ApiKeys` *(required)* : The private and public keys of the MongoDB Atlas.<br>