package resource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressEvents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/openlyinc/pointy"
	"go.mongodb.org/atlas/mongodbatlas"
)

var RequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ClusterName, constants.ProjectID}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Create snapshot for Request() currentModel:%+v", currentModel)

	// validate the request
	event, client, err := ValidateRequest(currentModel)
	if err != nil {
		return event, err
	}

	// Create Atlas API Request Object
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName

	nameSpaces := currentModel.ManagedNamespaces
	if len(nameSpaces) > 0 {
		err = addManagedNamespaces(context.Background(), client, nameSpaces, projectID, clusterName)
		if err != nil {
			_, _ = logger.Debugf("error creating MongoDB Global Cluster Configuration: %s", err)
			var target *mongodbatlas.ErrorResponse
			if errors.As(err, &target) && target.ErrorCode != "DUPLICATE_MANAGED_NAMESPACE" {
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          err.Error(),
					HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError,
				}, nil
			}
		}
	}
	zoneMappings := currentModel.CustomZoneMappings
	if len(zoneMappings) > 0 {
		customZoneMappings := modelToCustomZoneMappings(zoneMappings)
		// API call to create
		clusterDetail, _, err := client.GlobalClusters.AddCustomZoneMappings(context.Background(), projectID, clusterName, &mongodbatlas.CustomZoneMappingsRequest{
			CustomZoneMappings: customZoneMappings,
		})
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError,
			}, nil
		}
		responseData, _ := json.Marshal(clusterDetail)
		_, _ = logger.Debugf("Response Object: %s", responseData)
	}
	_, _ = logger.Debugln("Created Successfully -")

	event = handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel,
	}
	_, _ = logger.Debugf("Create() return event:%+v", event)
	return event, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	fmt.Printf("Read snapshot for Request() currentModel:%+v", currentModel)

	event, client, err := ValidateRequest(currentModel)
	if err != nil {
		return event, err
	}

	// Check if  already exist
	if !isExist(*client, currentModel) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// method call to read configuration
	config, _, err := ReadConfig(*client, currentModel)
	if err != nil {
		_, _ = logger.Debugf("error reading MongoDB Global Cluster Configuration (%s): %v", *currentModel.ClusterName, err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	_, _ = logger.Debugf("Response Value:%+v", config)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   config,
	}, nil
}
func ReadConfig(client mongodbatlas.Client, currentModel *Model) (*Model, handler.ProgressEvent, error) {
	setup() // logger setup
	projectID := *currentModel.ProjectId
	ClusterName := *currentModel.ClusterName

	// API call to read configuration
	globalCluster, resp, err := client.GlobalClusters.Get(context.Background(), projectID, ClusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		return nil, handler.ProgressEvent{}, err
	}
	nameSpaces := globalCluster.ManagedNamespaces
	zoneMappings := globalCluster.CustomZoneMapping

	if len(nameSpaces) == 0 && len(zoneMappings) == 0 {
		return nil, handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	readModel := new(Model)
	readModel.ProjectId = currentModel.ProjectId
	readModel.ClusterName = currentModel.ClusterName
	maps := flattenManagedNamespaces(globalCluster.ManagedNamespaces)
	zones := customZoneToModelMappings(globalCluster.CustomZoneMapping)
	readModel.CustomZoneMappings = zones
	readModel.ManagedNamespaces = maps
	readModel.ApiKeys = currentModel.ApiKeys
	readModel.RemoveAllZoneMapping = currentModel.RemoveAllZoneMapping
	_, _ = logger.Debugf("response Object:%+v", readModel)
	return readModel, handler.ProgressEvent{}, nil
}

func flattenManagedNamespaces(managedNamespaces []mongodbatlas.ManagedNamespace) []ManagedNamespace {
	var results []ManagedNamespace
	if len(managedNamespaces) > 0 {
		results = make([]ManagedNamespace, len(managedNamespaces))
		for k, managedNamespace := range managedNamespaces {
			results[k] = ManagedNamespace{
				Db:                     pointy.String(managedNamespace.Db),
				Collection:             pointy.String(managedNamespace.Collection),
				CustomShardKey:         pointy.String(managedNamespace.CustomShardKey),
				IsCustomShardKeyHashed: managedNamespace.IsCustomShardKeyHashed,
				IsShardKeyUnique:       managedNamespace.IsShardKeyUnique,
			}
		}
	}
	return results
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// No OP
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// No OP
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
	}, nil
}
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	_, _ = logger.Debugf("Delete snapshot for Request() currentModel:%+v", currentModel)

	// ValidateRequest function to validate the request
	event, client, err := ValidateRequest(currentModel)
	if err != nil {
		return event, err
	}

	// Check if  already exist
	if !isExist(*client, currentModel) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName
	remove := currentModel.ManagedNamespaces
	removeFlag := currentModel.RemoveAllZoneMapping
	// check if remove all zone mapping flag is enabled or name space array (to be removed) exist
	if (len(remove) == 0) && (removeFlag != nil && !*currentModel.RemoveAllZoneMapping) {
		// nothing to be removed .so raising invalid request
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "request doest not contain any item to remove",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if len(remove) > 0 {
		if err := removeManagedNamespaces(context.Background(), client, remove, projectID, clusterName); err != nil {
			_, _ = logger.Debugln("Error -Unable to remove Namespace")
		}
	}
	if currentModel.RemoveAllZoneMapping != nil && *currentModel.RemoveAllZoneMapping {
		_, _, err := client.GlobalClusters.DeleteCustomZoneMappings(context.Background(), projectID, clusterName)
		if err != nil {
			return progressEvents.GetFailedEventByCode(fmt.Sprintf("Failed to custom zones : %s", err.Error()),
				cloudformation.HandlerErrorCodeInvalidRequest), nil
		}
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func setup() {
	util.SetupLogger("mongodb-atlas-global-cluster-config")
}

func isExist(client mongodbatlas.Client, currentModel *Model) bool {
	setup()
	config, _, err := ReadConfig(client, currentModel)
	_, _ = logger.Debugf("%v", config)
	if err == nil && config != nil {
		return true
	}
	return false
}

// ValidateRequest function to validate the request
func ValidateRequest(currentModel *Model) (handler.ProgressEvent, *mongodbatlas.Client, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Request - snapshot restore starts currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil, errors.New("required field not found")
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressEvents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil, err
	}
	return handler.ProgressEvent{}, client, nil
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func removeManagedNamespaces(ctx context.Context, conn *mongodbatlas.Client, remove []ManagedNamespace, projectID, clusterName string) error {
	for _, m := range remove {
		addManagedNamespace := &mongodbatlas.ManagedNamespace{
			Collection:     *m.Collection,
			Db:             *m.Db,
			CustomShardKey: *m.CustomShardKey,
		}
		addManagedNamespace.IsCustomShardKeyHashed = pointy.Bool(*m.IsCustomShardKeyHashed)
		addManagedNamespace.IsShardKeyUnique = pointy.Bool(*m.IsShardKeyUnique)
		_, _, err := conn.GlobalClusters.DeleteManagedNamespace(ctx, projectID, clusterName, addManagedNamespace)

		if err != nil {
			_, _ = logger.Warnf("error while removing namespace:%+v", err)
		}
	}

	return nil
}

func addManagedNamespaces(ctx context.Context, conn *mongodbatlas.Client, add []ManagedNamespace, projectID, clusterName string) error {
	for _, m := range add {
		mn := m
		addManagedNamespace := &mongodbatlas.ManagedNamespace{
			Collection:     *mn.Collection,
			Db:             *mn.Db,
			CustomShardKey: *mn.CustomShardKey,
		}
		addManagedNamespace.IsCustomShardKeyHashed = mn.IsCustomShardKeyHashed
		addManagedNamespace.IsShardKeyUnique = mn.IsShardKeyUnique
		_, _, err := conn.GlobalClusters.AddManagedNamespace(ctx, projectID, clusterName, addManagedNamespace)
		if err != nil {
			_, _ = logger.Warnf("error while adding namespace:%+v", err)
		}
	}

	return nil
}

func modelToCustomZoneMapping(tfMap ZoneMapping) *mongodbatlas.CustomZoneMapping {
	apiObject := &mongodbatlas.CustomZoneMapping{
		Location: *tfMap.Location,
		Zone:     *tfMap.Zone,
	}

	return apiObject
}

func modelToCustomZoneMappings(tfList []ZoneMapping) []mongodbatlas.CustomZoneMapping {
	if len(tfList) == 0 {
		return nil
	}
	apiObjects := make([]mongodbatlas.CustomZoneMapping, len(tfList))
	if len(tfList) > 0 {
		for i, tfMapRaw := range tfList {
			if tfMapRaw == (ZoneMapping{}) || tfMapRaw.Location == nil || tfMapRaw.Zone == nil {
				continue
			}
			apiObject := modelToCustomZoneMapping(tfMapRaw)
			apiObjects[i] = *apiObject
		}
	}
	return apiObjects
}
func customZoneToModelMapping(location string, zone string) *ZoneMapping {
	if location == "" {
		return nil
	}
	return &ZoneMapping{
		Location: &location,
		Zone:     &zone,
	}
}

func customZoneToModelMappings(tfList map[string]string) []ZoneMapping {
	if len(tfList) == 0 {
		return nil
	}
	apiObjects := make([]ZoneMapping, len(tfList))
	var i = 0
	if len(tfList) > 0 {
		for k, v := range tfList {
			fmt.Printf("key[%s] value[%s]\n", k, v)
			apiObject := customZoneToModelMapping(k, v)
			if apiObject == nil {
				continue
			}
			apiObjects[i] = *apiObject
			i++
		}
	}
	return apiObjects
}
