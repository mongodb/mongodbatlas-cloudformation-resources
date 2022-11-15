package constants

type Event string

const (
	PubKey                           = "ApiKeys.PublicKey"
	PvtKey                           = "ApiKeys.PrivateKey"
	OrgID                            = "OrgId"
	Name                             = "Name"
	ID                               = "Id"
	ProjectID                        = "ProjectId"
	AccepterRegionName               = "AccepterRegionName"
	AwsAccountID                     = "AwsAccountId"
	RouteTableCIDRBlock              = "RouteTableCIDRBlock"
	AWS                              = "AWS"
	VPCID                            = "VpcId"

	RegionName     = "RegionName"
	AtlasCIDRBlock = "AtlasCidrBlock"

	DatabaseName = "DatabaseName"
	Username     = "Username"
	Roles        = "Roles"
	AccessList   = "AccessList"


	SnapshotID                       = "SnapshotId"
	Automated                        = "automated"
	Download                         = "download"
	ClusterName                      = "ClusterName"
	ErrorCreateMongoClient           = "Error - Create MongoDB Client- Details: %+v"
	ErrorCreateCloudBackupRestoreJob = "Error - Create Cloud Backup Restore snapshot for Snapshot(%s)- Details: %+v"
	ErrorReadCloudBackUpRestoreJob   = "Error - Read Restore Job with id(%s)"
	ResourceNotFound                 = "resource not found"
)
