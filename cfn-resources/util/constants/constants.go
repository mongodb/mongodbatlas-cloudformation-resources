package constants

type Event any

const (
	PubKey              = "ApiKeys.PublicKey"
	PvtKey              = "ApiKeys.PrivateKey"
	OrgID               = "OrgId"
	Name                = "Name"
	ID                  = "Id"
	AwsProviderName     = "AWS"
	TypeName            = "AWS::IAM::Role"
	AwsRolePrefix       = "mongodbatlas-role-"
	RolePolicyJsonGoSdk = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"$ATLAS_AWS_ACCOUNT_ARN\"}" +
		",\"Action\":\"sts:AssumeRole\",\"Condition\":{\"StringEquals\":{\"sts:ExternalId\":\"$ATLAS_ASSUMEDROLE_EXTERNAL_ID\"}}}]}\n"
	RolePolicyJson = "{\"AssumeRolePolicyDocument\":{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"$ATLAS_AWS_ACCOUNT_ARN\"},\"Action\":\"sts:AssumeRole\",\"Condition\":{\"StringEquals\":{\"sts:ExternalId\":\"$ATLAS_ASSUMEDROLE_EXTERNAL_ID\"}}}]},\"RoleName\":\"$ROLE_NAME\"}"
)
