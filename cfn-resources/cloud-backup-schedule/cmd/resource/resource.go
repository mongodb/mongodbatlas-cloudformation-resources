package resource

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.ClusterName}
var ReadRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.ClusterName}
var UpdateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.ClusterName}
var DeleteRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.ClusterName}

// validateModel inputs based on the method
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// logger setup function
func setup() {
	util.SetupLogger("mongodb-atlas-cloud-backup-schedule")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// logger setup
	setup()
	_, _ = logger.Debugf("Create() currentModel:%+v", currentModel)
	// Validate required fields in the request
	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	// API call to Create cloud backup schedule
	return cloudBackupScheduleCreateOrUpdate(req, prevModel, currentModel, client)
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// logger setup
	setup()
	_, _ = logger.Debugf("Get the current snapshot schedule and retention settings for the cluster:%+v", currentModel.ClusterName)
	// Validate required fields in the request
	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName

	// API call to Get the cloud backup schedule
	backupPolicy, resp, err := client.CloudProviderSnapshotBackupPolicies.Get(context.Background(), projectID, clusterName)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error deleting cloud backup schedule : %s", err.Error()),
			resp.Response), nil
	}
	_, _ = logger.Debugf("Read() end currentModel:%+v", currentModel)
	// check the policy backup schedule is present for the cluster
	if !isPolicySchedulePresent(backupPolicy) {
		_, _ = logger.Warnf("Error - Read policy backup schedule for cluster(%s)", *currentModel.ClusterName)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   backupPolicyToModel(*currentModel, backupPolicy),
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// logger setup
	setup()
	// Validate required fields in the request
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	_, _ = logger.Debugf("Update() currentModel:%+v", currentModel)
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	// API call to Get the cloud backup schedule
	events, _ := Read(req, prevModel, currentModel)
	if events.HandlerErrorCode == cloudformation.HandlerErrorCodeNotFound {
		return handler.ProgressEvent{
			Message:          "Not Found",
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// API call to Update cloud backup schedule
	return cloudBackupScheduleCreateOrUpdate(req, prevModel, currentModel, client)
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// logger setup
	setup()
	_, _ = logger.Debugf("Delete all the snapshot schedules for the cluster with currentModel:%+v", currentModel)
	// Validate required fields in the request
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	projectID := currentModel.ProjectId
	clusterName := currentModel.ClusterName
	// Check if cloud backup policy already exist
	if !isPolicyItemsExists(currentModel, client) {
		return progressevents.GetFailedEventByCode("Not Found", cloudformation.HandlerErrorCodeNotFound), nil
	}
	_, _ = logger.Debugf("Deleting all snapshot backup schedules for cluster(%s) from project(%s)", *currentModel.ClusterName, *currentModel.ProjectId)

	// API call to delete cloud backup schedule
	_, resp, err := client.CloudProviderSnapshotBackupPolicies.Delete(context.Background(), *projectID, *clusterName)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error deleting cloud backup schedule : %s", err.Error()),
			resp.Response), nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   nil,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("not implemented")
}

// handles the Create/Update event from the Cloudformation service.
func cloudBackupScheduleCreateOrUpdate(req handler.Request, prevModel *Model, currentModel *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {
	projectID := currentModel.ProjectId
	clusterName := currentModel.ClusterName
	_, _ = logger.Debugf("Update cloud backup schedule for clusterName:%s", *clusterName)

	// Delete policies items
	_, _ = logger.Debugf("First deleting cloud backup schedule for clusterName:%s", *clusterName)
	_, resp, err := client.CloudProviderSnapshotBackupPolicies.Delete(context.Background(), *projectID, *clusterName)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error deleting cloud backup schedule : %s", err.Error()),
			resp.Response), nil
	}
	if *currentModel.AutoExportEnabled && currentModel.Export != nil {
		if (currentModel.Export.FrequencyType) == nil {
			err := errors.New("error updating cloud backup schedule: FrequencyType should be set when `Export` is set")
			_, _ = logger.Warnf("Update - error: %+v", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
	}
	if currentModel.Policies != nil {
		for _, policy := range currentModel.Policies {
			if policy.PolicyItems != nil {
				for _, policyItem := range policy.PolicyItems {
					if policyItem.FrequencyInterval == nil {
						err := errors.New("error updating cloud backup schedule: PolicyItem FrequencyInterval should be set when `PolicyItems` is set")
						_, _ = logger.Warnf("Update - error: %+v", err)
						return handler.ProgressEvent{
							OperationStatus:  handler.Failed,
							Message:          err.Error(),
							HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
					}
					if policyItem.FrequencyType == nil {
						err := errors.New("error updating cloud backup schedule: PolicyItem FrequencyType should be set when `PolicyItems` is set")
						_, _ = logger.Warnf("Update - error: %+v", err)
						return handler.ProgressEvent{
							OperationStatus:  handler.Failed,
							Message:          err.Error(),
							HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
					}
					if policyItem.RetentionUnit == nil {
						err := errors.New("error updating cloud backup schedule: PolicyItem RetentionUnit should be set when `PolicyItems` is set")
						_, _ = logger.Warnf("Update - error: %+v", err)
						return handler.ProgressEvent{
							OperationStatus:  handler.Failed,
							Message:          err.Error(),
							HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
					}
					if policyItem.RetentionValue == nil {
						err := errors.New("error updating cloud backup schedule: PolicyItem RetentionValue should be set when `PolicyItems` is set")
						_, _ = logger.Warnf("Update - error: %+v", err)
						return handler.ProgressEvent{
							OperationStatus:  handler.Failed,
							Message:          err.Error(),
							HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
					}
				}
			}
		}
	}
	cloudBackupScheduleRequest := modelToCLoudBackupSchedule(currentModel)
	// API call to Create/Update cloud backup schedule
	clusterBackupScheduled, resp, err := client.CloudProviderSnapshotBackupPolicies.Update(context.Background(), *projectID, *clusterName, cloudBackupScheduleRequest)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error deleting cloud backup schedule : %s", err.Error()),
			resp.Response), nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "PATCH Complete",
		ResourceModel:   backupPolicyToModel(*currentModel, clusterBackupScheduled),
	}, nil
}

func castNO64(i *int64) *int {
	x := cast.ToInt(&i)
	return &x
}
func cast64(i *int) *int64 {
	x := cast.ToInt64(&i)
	return &x
}

func isPolicySchedulePresent(backupPolicy *mongodbatlas.CloudProviderSnapshotBackupPolicy) bool {
	return (backupPolicy.Policies != nil || len(backupPolicy.Policies) > 0) && len(backupPolicy.Policies[0].PolicyItems) > 0
}

// function to check if cloud backup policy already exist.
func isPolicyItemsExists(currentModel *Model, client *mongodbatlas.Client) bool {
	var isExists bool
	backupPolicy, _, err := client.CloudProviderSnapshotBackupPolicies.Get(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName)
	if err != nil {
		return isExists
	}
	if (backupPolicy.Policies != nil || len(backupPolicy.Policies) > 0) && len(backupPolicy.Policies[0].PolicyItems) > 0 {
		isExists = true
	}
	return isExists
}

// function to converts model 'Export' class to mongodb 'Export' class.
func expandExport(export Export) *mongodbatlas.Export {
	var exportArg mongodbatlas.Export

	if export.ExportBucketId != nil {
		exportArg.ExportBucketID = *export.ExportBucketId
	}
	if export.FrequencyType != nil {
		exportArg.FrequencyType = *export.FrequencyType
	}
	return &exportArg
}

// function to converts model 'ApiPolicyView' class to mongodb 'Policy' class.
func expandPolicies(policies []ApiPolicyView) []mongodbatlas.Policy {
	schedulePolicies := make([]mongodbatlas.Policy, 0)
	for _, s := range policies {
		policy := mongodbatlas.Policy{
			ID:          cast.ToString(s.ID),
			PolicyItems: expandPolicyItems(s.PolicyItems),
		}
		schedulePolicies = append(schedulePolicies, policy)
	}
	return schedulePolicies
}

// function to converts model 'ApiPolicyItemView' class to mongodb 'PolicyItem' class.
func expandPolicyItems(cloudPolicyItems []ApiPolicyItemView) []mongodbatlas.PolicyItem {
	policyItems := make([]mongodbatlas.PolicyItem, 0)
	for _, policyItem := range cloudPolicyItems {
		cPolicyItem := mongodbatlas.PolicyItem{
			ID:                cast.ToString(policyItem.ID),
			FrequencyInterval: cast.ToInt(policyItem.FrequencyInterval),
			FrequencyType:     cast.ToString(policyItem.FrequencyType),
			RetentionUnit:     cast.ToString(policyItem.RetentionUnit),
			RetentionValue:    cast.ToInt(policyItem.RetentionValue),
		}
		policyItems = append(policyItems, cPolicyItem)
	}
	return policyItems
}

// function to converts model 'ApiAtlasDiskBackupCopySettingView' class to mongodb 'CopySetting' class.
func expandCopySettings(copySettings []ApiAtlasDiskBackupCopySettingView) []mongodbatlas.CopySetting {
	cloudCopySettings := make([]mongodbatlas.CopySetting, 0)
	for _, copySetting := range copySettings {
		backupSetting := mongodbatlas.CopySetting{
			CloudProvider:     copySetting.CloudProvider,
			RegionName:        copySetting.RegionName,
			ReplicationSpecID: copySetting.ReplicationSpecId,
			ShouldCopyOplogs:  copySetting.ShouldCopyOplogs,
			Frequencies:       copySetting.Frequencies,
		}
		cloudCopySettings = append(cloudCopySettings, backupSetting)
	}
	return cloudCopySettings
}

// function to converts model 'ApiDeleteCopiedBackupsView' class to mongodb 'DeleteCopiedBackup' class.
func expandDeleteCopiedBackups(deleteCopiedBackups []ApiDeleteCopiedBackupsView) []mongodbatlas.DeleteCopiedBackup {
	cloudDeleteCopiedBackups := make([]mongodbatlas.DeleteCopiedBackup, 0)
	for _, deleteCopiedBackup := range deleteCopiedBackups {
		copiedSetting := mongodbatlas.DeleteCopiedBackup{
			CloudProvider:     deleteCopiedBackup.CloudProvider,
			RegionName:        deleteCopiedBackup.RegionName,
			ReplicationSpecID: deleteCopiedBackup.ReplicationSpecId,
		}
		cloudDeleteCopiedBackups = append(cloudDeleteCopiedBackups, copiedSetting)
	}
	return cloudDeleteCopiedBackups
}

// function to converts mongodb 'Policy' class to model 'ApiPolicyView' class.
func flattenPolicies(policies []mongodbatlas.Policy) []ApiPolicyView {
	snapPolicies := make([]ApiPolicyView, 0)
	for _, policy := range policies {
		snapPolicy := ApiPolicyView{
			ID:          &policy.ID,
			PolicyItems: flattenPolicyItems(policy.PolicyItems),
		}
		snapPolicies = append(snapPolicies, snapPolicy)
	}
	return snapPolicies
}

// function to converts mongodb 'PolicyItem' class to model 'ApiPolicyItemView' class.
func flattenPolicyItems(policyItems []mongodbatlas.PolicyItem) []ApiPolicyItemView {
	cloudPolicyItems := make([]ApiPolicyItemView, 0)
	for _, policyItem := range policyItems {
		snapPolicy := ApiPolicyItemView{
			ID:                &policyItem.ID,
			FrequencyInterval: &policyItem.FrequencyInterval,
			FrequencyType:     &policyItem.FrequencyType,
			RetentionUnit:     &policyItem.RetentionUnit,
			RetentionValue:    &policyItem.RetentionValue,
		}
		cloudPolicyItems = append(cloudPolicyItems, snapPolicy)
	}
	return cloudPolicyItems
}

// function to converts mongodb 'Export' class to model 'Export' class.
func flattenExport(export *mongodbatlas.Export) *Export {
	if export == nil {
		return nil
	}
	return &Export{
		ExportBucketId: &export.ExportBucketID,
		FrequencyType:  &export.FrequencyType,
	}
}

// converts mongodb 'link' class to model 'link' class.
func flattenLinks(linksResult []*mongodbatlas.Link) []Link {
	links := make([]Link, 0)
	for _, link := range linksResult {
		var lin Link
		lin.Href = &link.Href
		lin.Rel = &link.Rel
		links = append(links, lin)
	}
	return links
}

// converts mongodb 'CopySetting' class to model 'ApiAtlasDiskBackupCopySettingView' class.
func flattenCopySettings(copySettings []mongodbatlas.CopySetting) []ApiAtlasDiskBackupCopySettingView {
	cloudCopySettings := make([]ApiAtlasDiskBackupCopySettingView, 0)
	for _, copySetting := range copySettings {
		cloudCopySettings = append(cloudCopySettings, ApiAtlasDiskBackupCopySettingView{
			CloudProvider:     copySetting.CloudProvider,
			RegionName:        copySetting.RegionName,
			ReplicationSpecId: copySetting.ReplicationSpecID,
			ShouldCopyOplogs:  copySetting.ShouldCopyOplogs,
			Frequencies:       copySetting.Frequencies,
		})
	}
	return cloudCopySettings
}

// function to converts 'currentModel' model class to mongodb 'CloudProviderSnapshotBackupPolicy' class.
func modelToCLoudBackupSchedule(currentModel *Model) (out *mongodbatlas.CloudProviderSnapshotBackupPolicy) {
	out = &mongodbatlas.CloudProviderSnapshotBackupPolicy{}

	if currentModel.AutoExportEnabled != nil {
		out.AutoExportEnabled = currentModel.AutoExportEnabled
		if *currentModel.AutoExportEnabled && currentModel.Export != nil {
			out.Export = expandExport(*currentModel.Export)
		}
	}
	if currentModel.ReferenceHourOfDay != nil {
		out.ReferenceHourOfDay = cast64(currentModel.ReferenceHourOfDay)
	}
	if currentModel.ReferenceMinuteOfHour != nil {
		out.ReferenceMinuteOfHour = cast64(currentModel.ReferenceMinuteOfHour)
	}
	if currentModel.RestoreWindowDays != nil {
		out.RestoreWindowDays = cast64(currentModel.RestoreWindowDays)
	}
	if currentModel.UseOrgAndGroupNamesInExportPrefix != nil {
		out.UseOrgAndGroupNamesInExportPrefix = currentModel.UseOrgAndGroupNamesInExportPrefix
	}
	if currentModel.Policies != nil {
		out.Policies = expandPolicies(currentModel.Policies)
	}
	if currentModel.UpdateSnapshots != nil {
		out.UpdateSnapshots = currentModel.UpdateSnapshots
	}
	if currentModel.CopySettings != nil || len(currentModel.CopySettings) > 0 {
		out.CopySettings = expandCopySettings(currentModel.CopySettings)
	}
	if currentModel.DeleteCopiedBackups != nil || len(currentModel.DeleteCopiedBackups) > 0 {
		out.DeleteCopiedBackups = expandDeleteCopiedBackups(currentModel.DeleteCopiedBackups)
	}
	return out
}

// function to converts  mongodb 'CloudProviderSnapshotBackupPolicy' class to 'currentModel' model class.
func backupPolicyToModel(currentModel Model, backupPolicy *mongodbatlas.CloudProviderSnapshotBackupPolicy) *Model {
	out := &Model{
		ApiKeys:                           currentModel.ApiKeys,
		ProjectId:                         currentModel.ProjectId,
		ClusterId:                         &backupPolicy.ClusterID,
		ClusterName:                       &backupPolicy.ClusterName,
		AutoExportEnabled:                 backupPolicy.AutoExportEnabled,
		Policies:                          flattenPolicies(backupPolicy.Policies),
		ReferenceHourOfDay:                castNO64(backupPolicy.ReferenceHourOfDay),
		ReferenceMinuteOfHour:             castNO64(backupPolicy.ReferenceMinuteOfHour),
		RestoreWindowDays:                 castNO64(backupPolicy.RestoreWindowDays),
		NextSnapshot:                      &backupPolicy.NextSnapshot,
		UseOrgAndGroupNamesInExportPrefix: backupPolicy.UseOrgAndGroupNamesInExportPrefix,
		Links:                             flattenLinks(backupPolicy.Links),
		CopySettings:                      flattenCopySettings(backupPolicy.CopySettings),
	}
	if backupPolicy.Export != nil && *currentModel.AutoExportEnabled {
		out.Export = flattenExport(backupPolicy.Export)
	}
	return out
}
