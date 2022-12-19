package constants

type Event string

const (
	PubKey              = "ApiKeys.PublicKey"
	PvtKey              = "ApiKeys.PrivateKey"
	OrgID               = "OrgId"
	Name                = "Name"
	ID                  = "Id"
	ProjectID           = "ProjectId"
	AccepterRegionName  = "AccepterRegionName"
	AwsAccountID        = "AwsAccountId"
	RouteTableCIDRBlock = "RouteTableCIDRBlock"
	AWS                 = "AWS"
	VPCID               = "VpcId"
	SubnetID            = "SubnetId"
	GroupID             = "GroupId"
	Region              = "Region"

	RegionName     = "RegionName"
	AtlasCIDRBlock = "AtlasCidrBlock"
	RoleName       = "RoleName"

	DatabaseName = "DatabaseName"
	Username     = "Username"
	Roles        = "Roles"
	AccessList   = "AccessList"

	CreatingState = "CREATING"
	UpdateState   = "UPDATING"
	DeletingState = "DELETING"
	DeletedState  = "DELETED"
	IdleState     = "IDLE"

	Error            = "ERROR"
	DeleteInProgress = "Delete in progress"
	StateName        = "StateName"
	Complete         = "Complete"
	Pending          = "Pending"
	ReadComplete     = "Read Complete"

	CustomMasterKey        = "AwsKms.CustomerMasterKeyID"
	RoleID                 = "AwsKms.RoleID"
	ErrorCreateMongoClient = "error - Create MongoDB Client- Details: %+v"
	ResourceNotFound       = "resource not found"

	SnapshotID                       = "SnapshotId"
	Automated                        = "automated"
	Download                         = "download"
	ClusterName                      = "ClusterName"
	ErrorCreateCloudBackupRestoreJob = "Error - Create Cloud Backup Restore snapshot for Snapshot(%s)- Details: %+v"
	ErrorReadCloudBackUpRestoreJob   = "Error - Read Restore Job with id(%s)"

	EndpointID = "EndpointId"
	DataLake   = "DATA_LAKE"

	IntegrationType = "Type"

	IndexID        = "IndexId"
	CollectionName = "CollectionName"
	Database       = "Database"

	UserID  = "UserName"
	Success = 200

	EventTypeName = "EventTypeName"
	PagerDuty     = "PAGER_DUTY"
	OpsGenie      = "OPS_GENIE"
	VictorOps     = "VICTOR_OPS"

	TeamID = "TeamId"

	DeliveryType = "DeliveryType"

	TenantName = "TenantName"

	CloudProviderAccessRoleID = "RoleId"
	LocalSessionType          = "LOCAL_SESSION"
	RequestSessionType        = "REQUEST_SESSION"
	TypeName                  = "AWS::IAM::Role"
	AwsRolePrefix             = "mongodbatlas-role-"
	RolePolicyJSON            = "{\"AssumeRolePolicyDocument\":{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{" +
		"\"AWS\":\"$ATLAS_AWS_ACCOUNT_ARN\"},\"Action\":\"sts:AssumeRole\",\"Condition\":{\"StringEquals\":{\"sts:ExternalId\":" +
		"\"$ATLAS_ASSUMEDROLE_EXTERNAL_ID\"}}}]},\"RoleName\":\"$ROLE_NAME\"}"
	RoleCreatingMessage = "Aws Role Creating..."
	RoleDeletingMessage = "Aws Role Deleting..."
)
