package resource

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"log"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.ClusterName}
var ReadRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.ClusterName}
var UpdateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.ClusterName}
var DeleteRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.ClusterName}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.ClusterName}

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
		_, _ = logger.Warnf("Create - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
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
		_, _ = logger.Warnf("CreateMongoDBClient error: %s", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName

	// API call to Get the cloud backup schedule
	backupPolicy, resp, errGet := client.CloudProviderSnapshotBackupPolicies.Get(context.Background(), projectID, clusterName)
	if errGet != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Warnf("error 404- err:%+v resp:%+v", errGet, resp)
			return handler.ProgressEvent{
				Message:          errGet.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = logger.Warnf("error cloud backup policy get- err:%+v resp:%+v", errGet, resp)
		return handler.ProgressEvent{
			Message:          errGet.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	_, _ = logger.Debugf("Read() end currentModel:%+v", currentModel)
	log.Printf("Read() end currentModel:%+v", currentModel)

	// check the policy backup schedule is present for the cluster
	if !isPolicySchedulePresent(backupPolicy) {
		_, _ = logger.Warnf("Error - Read policy backup schedule for cluster(%s)", *currentModel.ClusterName)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
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
		_, _ = logger.Warnf("Update - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	// TODO: Need to refactor
	// API call to Get the cloud backup schedule
	backupPolicy, resp, err := client.CloudProviderSnapshotBackupPolicies.Get(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Warnf("error 404- err:%+v resp:%+v", err, resp)
			log.Printf("error 404- err:%+v resp:%+v", err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = logger.Warnf("error cloud backup policy get- err:%+v resp:%+v", err, resp)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	// check the policy backup schedule is present for the cluster
	if &backupPolicy.Policies[0].ID == nil {
		_, _ = logger.Warnf("Error ---- Read policy backup schedule for cluster(%s)", *currentModel.ClusterName)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
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
		_, _ = logger.Warnf("Create - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	projectID := currentModel.ProjectId
	clusterName := currentModel.ClusterName

	// TODO: Need to refactor
	// API call to Get the cloud backup schedule
	backupPolicy, resp, err := client.CloudProviderSnapshotBackupPolicies.Get(context.Background(), *projectID, *clusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Warnf("error 404- err:%+v resp:%+v", err, resp)
			log.Printf("error 404- err:%+v resp:%+v", err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = logger.Warnf("error cloud backup policy get- err:%+v resp:%+v", err, resp)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}

	// check the policy backup schedule is present for the cluster
	if !isPolicySchedulePresent(backupPolicy) {
		_, _ = logger.Warnf("Error - Read policy backup schedule for cluster(%s)", *currentModel.ClusterName)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	_, _ = logger.Debugf("Deleting all snapshot backup schedules for cluster(%s) from project(%s)", *currentModel.ClusterName, *currentModel.ProjectId)

	// API call to delete cloud backup schedule
	_, resp, err = client.CloudProviderSnapshotBackupPolicies.Delete(context.Background(), *projectID, *clusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Warnf("Delete 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = logger.Warnf("Delete err: %+v", err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
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

	//Delete policies items
	_, _ = logger.Debugf("First deleting cloud backup schedule for clusterName:%s", *clusterName)
	_, _, err := client.CloudProviderSnapshotBackupPolicies.Delete(context.Background(), *projectID, *clusterName)
	if err != nil {
		_, _ = logger.Warnf("error deleting MongoDB Cloud Backup Schedule (%s): %v", *clusterName, err)
		err := errors.New("error deleting MongoDB Cloud Backup Schedule")
		_, _ = logger.Warnf("delete - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
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
	clusterBackupScheduled, resp, errUpdate := client.CloudProviderSnapshotBackupPolicies.Update(context.Background(), *projectID, *clusterName, cloudBackupScheduleRequest)
	if errUpdate != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Warnf("update 404 err: %+v", errUpdate)
			return handler.ProgressEvent{
				Message:          errUpdate.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = logger.Warnf("update err: %+v", errUpdate)
		code := cloudformation.HandlerErrorCodeServiceInternalError
		if strings.Contains(errUpdate.Error(), "not exist") { // cfn test needs 404
			code = cloudformation.HandlerErrorCodeNotFound
		}
		if strings.Contains(errUpdate.Error(), "being deleted") {
			code = cloudformation.HandlerErrorCodeNotFound // cfn test needs 404
		}
		return handler.ProgressEvent{
			Message:          errUpdate.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: code}, nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "PATCH Complete",
		ResourceModel:   backupPolicyToModel(*currentModel, clusterBackupScheduled),
	}, nil
}

func isPolicySchedulePresent(backupPolicy *mongodbatlas.CloudProviderSnapshotBackupPolicy) bool {

	return backupPolicy.Policies != nil && len(backupPolicy.Policies[0].PolicyItems) > 0
}

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

func castNO64(i *int64) *int {
	x := cast.ToInt(&i)
	return &x
}
func cast64(i *int) *int64 {
	x := cast.ToInt64(&i)
	return &x
}

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

func flattenPolicyItems(policyItems []mongodbatlas.PolicyItem) []ApiPolicyItemView {
	cloudPolicyItems := make([]ApiPolicyItemView, 0)
	for _, policyItem := range policyItems {
		cloudPolicyItems = append(cloudPolicyItems, ApiPolicyItemView{
			ID:                &policyItem.ID,
			FrequencyInterval: &policyItem.FrequencyInterval,
			FrequencyType:     &policyItem.FrequencyType,
			RetentionUnit:     &policyItem.RetentionUnit,
			RetentionValue:    &policyItem.RetentionValue,
		})
	}
	return cloudPolicyItems
}

func flattenExport(export *mongodbatlas.Export) *Export {
	if export == nil {
		return nil
	}
	return &Export{
		ExportBucketId: &export.ExportBucketID,
		FrequencyType:  &export.FrequencyType,
	}
}

func modelToCLoudBackupSchedule(currentModel *Model) (out *mongodbatlas.CloudProviderSnapshotBackupPolicy) {
	out = &mongodbatlas.CloudProviderSnapshotBackupPolicy{}

	if currentModel.AutoExportEnabled != nil {
		out.AutoExportEnabled = currentModel.AutoExportEnabled
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
	if *currentModel.AutoExportEnabled && currentModel.Export != nil {
		out.Export = expandExport(*currentModel.Export)
	}
	if currentModel.UpdateSnapshots != nil {
		out.UpdateSnapshots = currentModel.UpdateSnapshots
	}
	return out
}

func backupPolicyToModel(currentModel Model, backupPolicy *mongodbatlas.CloudProviderSnapshotBackupPolicy) *Model {

	out := &Model{
		ApiKeys:                           currentModel.ApiKeys,
		ProjectId:                         currentModel.ProjectId,
		ClusterId:                         &backupPolicy.ClusterID,
		ClusterName:                       &backupPolicy.ClusterName,
		AutoExportEnabled:                 backupPolicy.AutoExportEnabled,
		Export:                            flattenExport(backupPolicy.Export),
		Policies:                          flattenPolicies(backupPolicy.Policies),
		ReferenceHourOfDay:                castNO64(backupPolicy.ReferenceHourOfDay),
		ReferenceMinuteOfHour:             castNO64(backupPolicy.ReferenceMinuteOfHour),
		RestoreWindowDays:                 castNO64(backupPolicy.RestoreWindowDays),
		NextSnapshot:                      &backupPolicy.NextSnapshot,
		UseOrgAndGroupNamesInExportPrefix: backupPolicy.UseOrgAndGroupNamesInExportPrefix,
	}
	if !*currentModel.AutoExportEnabled {
		out.Export = nil
	}
	log.Printf("out---")
	spew.Dump(out)
	return out
}
