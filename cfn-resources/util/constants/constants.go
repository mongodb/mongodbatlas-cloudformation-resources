package constants

type Event string

const (
	PubKey                 = "ApiKeys.PublicKey"
	PvtKey                 = "ApiKeys.PrivateKey"
	OrgID                  = "OrgId"
	Name                   = "Name"
	ID                     = "Id"
	ProjectID              = "ProjectId"
	AccepterRegionName     = "AccepterRegionName"
	AwsAccountID           = "AwsAccountId"
	RouteTableCIDRBlock    = "RouteTableCIDRBlock"
	AWS                    = "AWS"
	VPCID                  = "VpcId"
	CustomMasterKey        = "AwsKms.CustomerMasterKeyID"
	RoleID                 = "AwsKms.RoleID"
	Region                 = "AwsKms.Region"
	ErrorCreateMongoClient = "error - Create MongoDB Client- Details: %+v"
	ResourceNotFound       = "resource not found"
)
