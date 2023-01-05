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

	IntegrationType = "Type"

	UserID  = "UserName"
	Success = 200

	UserName                         = "UserName"
	ExportBucketID                   = "ExportBucketId"
	ExportID                         = "ExportId"
	ErrorExportJobCreate             = "error creating Export Job for the project(%s) : %s"
	ErrorExportJobRead               = "error reading export job for the projects(%s) : Job Id : %s with error :%+v"

)
