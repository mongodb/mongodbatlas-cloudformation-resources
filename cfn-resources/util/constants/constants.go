// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package constants

type Event string

type CfnFunctions string

const (
	CREATE CfnFunctions = "CREATE"
	LIST   CfnFunctions = "LIST"
	READ   CfnFunctions = "READ"
	UPDATE CfnFunctions = "UPDATE"
	DELETE CfnFunctions = "DELETE"
)

const (
	ProfileNamePrefix       = "cfn/atlas/profile"
	Profile                 = "Profile"
	PubKey                  = "ApiKeys.PublicKey"
	PvtKey                  = "ApiKeys.PrivateKey"
	OrgID                   = "OrgId"
	Name                    = "Name"
	ID                      = "Id"
	ProjectID               = "ProjectId"
	AccepterRegionName      = "AccepterRegionName"
	AwsAccountID            = "AwsAccountId"
	RouteTableCIDRBlock     = "RouteTableCIDRBlock"
	AWS                     = "AWS"
	VPCID                   = "VpcId"
	SubnetID                = "SubnetId"
	GroupID                 = "GroupId"
	Region                  = "Region"
	HostName                = "HostName"
	Port                    = "Port"
	ContainerID             = "ContainerId"
	Sink                    = "Sink"
	Transformations         = "transformations"
	CloudProvider           = "CloudProvider"
	Specs                   = "Specs"
	DefaultListItemsPerPage = 100

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

	ErrorCreateMongoClient = "error - Create MongoDB Client- Details: %+v"
	ResourceNotFound       = "resource not found"

	SnapshotID                       = "SnapshotId"
	Automated                        = "automated"
	Download                         = "download"
	ClusterName                      = "ClusterName"
	InstanceName                     = "InstanceName"
	InstanceType                     = "InstanceType"
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

	Criteria     = "Criteria"
	ArchiveID    = "ArchiveId"
	CriteriaType = "Criteria.Type"

	EventTypeName = "EventTypeName"
	PagerDuty     = "PAGER_DUTY"
	OpsGenie      = "OPS_GENIE"
	VictorOps     = "VICTOR_OPS"

	TeamID = "TeamId"

	ErrorCreateCloudBackup = "Error - Create Cloud Backup snapshot for Project(%s) and Cluster(%s)- Details: %+v"

	DeliveryType = "DeliveryType"

	TenantName                = "TenantName"
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
	ProjID              = "ProjectID"
	Serverless          = "SERVERLESS"
	Duplicate           = "DUPLICATE"

	AppID                = "AppId"
	FederationSettingsID = "FederationSettingsId"

	ExportBucketID             = "ExportBucketId"
	ExportID                   = "ExportId"
	UnfinishedOnDemandSnapshot = "UNFINISHED_ON_DEMAND_SNAPSHOT"

	ExternalGroupName          = "ExternalGroupName"
	RoleAssignments            = "RoleAssignments"
	Description                = "Description"
	AwsSecretName              = "AwsSecretName"
	APIUserID                  = "APIUserId"
	DataFederationRoleID       = "CloudProviderConfig.RoleId"
	DataFederationTestS3Bucket = "CloudProviderConfig.TestS3Bucket"
	DataProcessRegion          = "DataProcessRegion.Region"
	SkipRoleValidation         = "SkipRoleValidation"
	LimitName                  = "LimitName"
	Value                      = "Value"

	AlreadyExist = "Already Exist"
	EmptyString  = ""

	OrgOwnerID        = "OrgOwnerId"
	OrgKeyRoles       = "APIKey.Roles"
	OrgKeyDescription = "APIKey.Description"

	ConnectionName = "ConnectionName"
	Type           = "Type"
	StreamConfig   = "StreamConfig"

	ProcessorName = "ProcessorName"
	Pipeline      = "Pipeline"
)
