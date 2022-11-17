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

	DatabaseName = "DatabaseName"
	Username     = "Username"
	Roles        = "Roles"
	AccessList   = "AccessList"

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
)
