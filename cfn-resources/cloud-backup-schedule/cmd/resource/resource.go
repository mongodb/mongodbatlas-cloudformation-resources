package resource

import (
	"context"
	"errors"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
	"strings"
)

var CreateRequiredFields = []string{"ProjectId", "ClusterName", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var ReadRequiredFields = []string{"ProjectId", "ClusterName", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var UpdateRequiredFields = []string{"ProjectId", "ClusterName", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var DeleteRequiredFields = []string{"ProjectId", "ClusterName", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}

// logger setup function
func setup() {
	util.SetupLogger("mongodb-atlas-project")
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	//logger setup
	setup()
	log.Debugf("Get the current snapshot schedule and retention settings for the cluster:%+v", currentModel.ClusterName)
	// Validate required fields in the request
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Infof("Read - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName

	// API call to Get the cloud backup schedule
	backupPolicy, resp, err := client.CloudProviderSnapshotBackupPolicies.Get(context.Background(), projectID, clusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Infof("error 404- err:%+v resp:%+v", err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		} else {
			log.Infof("error cloud backup policy get- err:%+v resp:%+v", err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
		}
	}

	if currentModel.ReferenceHourOfDay != nil {
		currentModel.ReferenceHourOfDay = castNO64(backupPolicy.ReferenceHourOfDay)
	}

	if currentModel.ReferenceMinuteOfHour != nil {
		currentModel.ReferenceMinuteOfHour = castNO64(backupPolicy.ReferenceMinuteOfHour)
	}

	if currentModel.RestoreWindowDays != nil {
		currentModel.RestoreWindowDays = castNO64(backupPolicy.RestoreWindowDays)
	}

	if currentModel.UpdateSnapshots != nil {
		currentModel.UpdateSnapshots = backupPolicy.UpdateSnapshots
	}

	if currentModel.NextSnapshot != nil {
		currentModel.NextSnapshot = &backupPolicy.NextSnapshot
	}

	if currentModel.IdPolicy != nil {
		currentModel.IdPolicy = &backupPolicy.Policies[0].ID
	}

	if currentModel.Export.ExportBucketId != nil {
		currentModel.Export.ExportBucketId = &backupPolicy.Export.ExportBucketID
	}

	if currentModel.Export.FrequencyType != nil {
		currentModel.Export.FrequencyType = &backupPolicy.Export.FrequencyType
	}

	if currentModel.AutoExportEnabled != nil {
		currentModel.AutoExportEnabled = backupPolicy.AutoExportEnabled
	}

	if currentModel.UseOrgAndGroupNamesInExportPrefix != nil {
		currentModel.UseOrgAndGroupNamesInExportPrefix = backupPolicy.UseOrgAndGroupNamesInExportPrefix
	}

	if currentModel.Policies != nil {
		currentModel.Policies = flattenPolicies(backupPolicy.Policies)
	}

	log.Debugf("Read() end currentModel:%+v", currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	//logger setup
	setup()
	// Validate required fields in the request
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	log.Debugf("Create() currentModel:%+v", currentModel)
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Infof("Create - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	// API call to Create cloud backup schedule
	handlerEvent, err := CloudBackupScheduleCreateOrUpdate(req, prevModel, currentModel, client)
	if err != nil {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: handlerEvent.HandlerErrorCode,
		}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	//logger setup
	setup()
	// Validate required fields in the request
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	log.Debugf("Update() currentModel:%+v", currentModel)
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Infof("Update - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	// API call to Update cloud backup schedule
	handlerEvent, err := CloudBackupScheduleCreateOrUpdate(req, prevModel, currentModel, client)
	if err != nil {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: handlerEvent.HandlerErrorCode,
		}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	//logger setup
	setup()
	log.Debugf("Delete all the snapshot schedules for the cluster with currentModel:%+v", currentModel)
	// Validate required fields in the request
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Errorf("Create - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	// key.ResourceID should == *currentModel.Id
	id, err := util.ParseResourceIdentifier(*currentModel.Id)
	if err != nil {
		log.Infof("Read - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	log.Debugf("Parsed resource identifier: id:%+v", id)

	projectID := id.Parent.ResourceID
	clusterName := id.ResourceID

	log.Debugf("Deleting all snapshot schedules for (%s)", *currentModel.Id)

	// API call to delete cloud backup schedule
	_, resp, err := client.CloudProviderSnapshotBackupPolicies.Delete(context.Background(), projectID, clusterName)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Errorf("Delete 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		} else {
			log.Errorf("Delete err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("not implemented")
}

// handles the Create/Update event from the Cloudformation service.
func CloudBackupScheduleCreateOrUpdate(req handler.Request, prevModel *Model, currentModel *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {

	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName
	log.Debugf("Update - clusterName:%s", clusterName)

	// Delete policies items
	_, _, err := client.CloudProviderSnapshotBackupPolicies.Delete(context.Background(), projectID, clusterName)
	if err != nil {
		log.Errorf("error deleting MongoDB Cloud Backup Schedule (%s): %s", clusterName, err)
		err := errors.New("error deleting MongoDB Cloud Backup Schedule")
		log.Infof("delete - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if (currentModel.Export.ExportBucketId) != nil {
		err := errors.New("error updating cloud backup schedule: ExportBucketId should be set when `Export` is set")
		log.Infof("Update - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	if (currentModel.Export.FrequencyType) != nil {
		err := errors.New("error updating cloud backup schedule: FrequencyType should be set when `Export` is set")
		log.Infof("Update - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if len(currentModel.Policies) > 0 {
		if currentModel.Policies[0].ID != nil {
			err := errors.New("error updating cloud backup schedule: Policies ID should be set when `Policies` is set")
			log.Infof("Update - error: %+v", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}

		if len(currentModel.Policies[0].PolicyItems) > 0 {
			if currentModel.Policies[0].PolicyItems[0].Id != nil {
				err := errors.New("error updating cloud backup schedule: PolicyItem ID should be set when `PolicyItems` is set")
				log.Infof("Update - error: %+v", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          err.Error(),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
			if currentModel.Policies[0].PolicyItems[0].FrequencyInterval != nil {
				err := errors.New("error updating cloud backup schedule: PolicyItem FrequencyInterval should be set when `PolicyItems` is set")
				log.Infof("Update - error: %+v", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          err.Error(),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
			if currentModel.Policies[0].PolicyItems[0].FrequencyType != nil {
				err := errors.New("error updating cloud backup schedule: PolicyItem FrequencyType should be set when `PolicyItems` is set")
				log.Infof("Update - error: %+v", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          err.Error(),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
			if currentModel.Policies[0].PolicyItems[0].RetentionUnit != nil {
				err := errors.New("error updating cloud backup schedule: PolicyItem RetentionUnit should be set when `PolicyItems` is set")
				log.Infof("Update - error: %+v", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          err.Error(),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
			if currentModel.Policies[0].PolicyItems[0].RetentionValue != nil {
				err := errors.New("error updating cloud backup schedule: PolicyItem RetentionValue should be set when `PolicyItems` is set")
				log.Infof("Update - error: %+v", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          err.Error(),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}
	}
	export := &mongodbatlas.Export{
		ExportBucketID: *currentModel.Export.ExportBucketId,
		FrequencyType:  *currentModel.Export.FrequencyType,
	}

	cloudBackupScheduleRequest := &mongodbatlas.CloudProviderSnapshotBackupPolicy{
		ReferenceHourOfDay:    cast64(currentModel.ReferenceHourOfDay),
		ReferenceMinuteOfHour: cast64(currentModel.ReferenceMinuteOfHour),
		RestoreWindowDays:     cast64(currentModel.RestoreWindowDays),
		Export:                export,
		UpdateSnapshots:       currentModel.UpdateSnapshots,
		AutoExportEnabled:     currentModel.AutoExportEnabled,
	}

	if currentModel.Policies != nil {
		cloudBackupScheduleRequest.Policies = expandPolicies(currentModel.Policies)
	}
	// API call to Create/Update cloud backup schedule
	_, resp, err := client.CloudProviderSnapshotBackupPolicies.Update(context.Background(), projectID, clusterName, cloudBackupScheduleRequest)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Infof("update 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		} else {
			log.Infof("update err: %+v", err)
			code := cloudformation.HandlerErrorCodeServiceInternalError
			if strings.Contains(err.Error(), "not exist") { // cfn test needs 404
				code = cloudformation.HandlerErrorCodeNotFound
			}
			if strings.Contains(err.Error(), "being deleted") {
				code = cloudformation.HandlerErrorCodeNotFound // cfn test needs 404
			}
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: code}, nil
		}
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

func expandPolicies(policies []Policies) []mongodbatlas.Policy {
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

func expandPolicyItems(cloudPolicyItems []PolicyItems) []mongodbatlas.PolicyItem {
	policyItems := make([]mongodbatlas.PolicyItem, 0)
	for _, policyItem := range cloudPolicyItems {
		cPolicyItem := mongodbatlas.PolicyItem{
			ID:                cast.ToString(policyItem.Id),
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

// validateModel inputs based on the method
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func flattenPolicies(rSpecs []mongodbatlas.Policy) []Policies {
	specs := make([]Policies, 0)
	for _, rSpec := range rSpecs {
		spec := Policies{
			ID:          &rSpec.ID,
			PolicyItems: flattenPolicyItems(rSpec.PolicyItems),
		}
		specs = append(specs, spec)
	}
	return specs
}

func flattenPolicyItems(policyItems []mongodbatlas.PolicyItem) []PolicyItems {
	cloudPolicyItems := make([]PolicyItems, 0)

	for _, policyItem := range policyItems {
		policyItemVal := PolicyItems{
			Id:                &policyItem.ID,
			FrequencyInterval: &policyItem.FrequencyInterval,
			FrequencyType:     &policyItem.FrequencyType,
			RetentionValue:    &policyItem.RetentionValue,
			RetentionUnit:     &policyItem.RetentionUnit,
		}
		cloudPolicyItems = append(cloudPolicyItems, policyItemVal)
	}
	return cloudPolicyItems
}
