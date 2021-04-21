package resource

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

func setup() {
	util.SetupLogger("mongodb-atlas-cluster")

}

func castNO64(i *int64) *int {
	x := cast.ToInt(&i)
	return &x
}
func cast64(i *int) *int64 {
	x := cast.ToInt64(&i)
	return &x
}
func boolPtr(i bool) *bool {
	return &i
}
func intPtr(i int) *int {
	return &i
}
func stringPtr(i string) *string {
	return &i
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Create() currentModel:%+v", currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Infof("Create - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil

	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(client, req, currentModel, "IDLE", "CREATING")
	}

	projectID := *currentModel.ProjectId
	log.Infof("cluster Create projectID=%s", projectID)
	if len(currentModel.ReplicationSpecs) > 0 {
		if currentModel.ClusterType != nil {
			err := errors.New("error creating cluster: ClusterType should be set when `ReplicationSpecs` is set")
			log.Infof("Create - error: %+v", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}

		if currentModel.NumShards != nil {
			err := errors.New("error creating cluster: NumShards should be set when `ReplicationSpecs` is set")
			log.Infof("Create - error: %+v", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
	}

	// This is the intial call to Create, so inject a deployment
	// secret for this resource in order to lookup progress properly
	projectResID := &util.ResourceIdentifier{
		ResourceType: "Project",
		ResourceID:   projectID,
	}
	log.Debugf("Created projectResID:%s", projectResID)
	resourceID := util.NewResourceIdentifier("Cluster", *currentModel.Name, projectResID)
	log.Debugf("Created resourceID:%s", resourceID)
	resourceProps := map[string]string{
		"Clust:erName": *currentModel.Name,
	}
	secretName, err := util.CreateDeploymentSecret(&req, resourceID, *currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey, &resourceProps)
	if err != nil {
		log.Infof("Create - CreateDeploymentSecret - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}

	log.Infof("Created new deployment secret for cluster. Secert Name = Cluster Id:%s", *secretName)
	currentModel.Id = secretName

	one := int64(1)
	three := int64(3)
	replicationFactor := &three
	if currentModel.ReplicationFactor != nil {
		rf := int64(*currentModel.ReplicationFactor)
		replicationFactor = &rf
	} else {
		log.Debugf("Default setting ReplicationFactor to 3")
	}

	numShards := &one
	if currentModel.NumShards != nil {
		ns := int64(*currentModel.NumShards)
		numShards = &ns
	} else {
		log.Debugf("Default setting NumShards to 1")
	}

	clusterRequest := &mongodbatlas.Cluster{
		Name:                     cast.ToString(currentModel.Name),
		EncryptionAtRestProvider: cast.ToString(currentModel.EncryptionAtRestProvider),
		ClusterType:              cast.ToString(currentModel.ClusterType),
		ReplicationFactor:        replicationFactor,
		NumShards:                numShards,
	}

	if currentModel.BackupEnabled != nil {
		clusterRequest.BackupEnabled = currentModel.BackupEnabled
	}

	if currentModel.ProviderBackupEnabled != nil {
		clusterRequest.ProviderBackupEnabled = currentModel.ProviderBackupEnabled
	}

	if currentModel.DiskSizeGB != nil {
		currentModel.DiskSizeGB = clusterRequest.DiskSizeGB
	}

	if currentModel.MongoDBMajorVersion != nil {
		clusterRequest.MongoDBMajorVersion = formatMongoDBMajorVersion(*currentModel.MongoDBMajorVersion)
	}

	if currentModel.BiConnector != nil {
		clusterRequest.BiConnector = expandBiConnector(currentModel.BiConnector)
	}

	if currentModel.ProviderSettings != nil {
		clusterRequest.ProviderSettings = expandProviderSettings(currentModel.ProviderSettings)
	}
	log.Debugf("DEBUG: clusterRequest.ProviderSettings: %+v", clusterRequest.ProviderSettings)

	if currentModel.ReplicationSpecs != nil {
		clusterRequest.ReplicationSpecs = expandReplicationSpecs(currentModel.ReplicationSpecs)
	}

	if currentModel.AutoScaling != nil {
		clusterRequest.AutoScaling = &mongodbatlas.AutoScaling{
			DiskGBEnabled: currentModel.AutoScaling.DiskGBEnabled,
		}
		if currentModel.AutoScaling.Compute != nil {
			compute := &mongodbatlas.Compute{}
			if currentModel.AutoScaling.Compute.Enabled != nil {
				compute.Enabled = currentModel.AutoScaling.Compute.Enabled
			}
			if currentModel.AutoScaling.Compute.ScaleDownEnabled != nil {
				compute.ScaleDownEnabled = currentModel.AutoScaling.Compute.ScaleDownEnabled
			}
			if currentModel.AutoScaling.Compute.MinInstanceSize != nil {
				compute.MinInstanceSize = *currentModel.AutoScaling.Compute.MinInstanceSize
			}
			if currentModel.AutoScaling.Compute.MaxInstanceSize != nil {
				compute.MaxInstanceSize = *currentModel.AutoScaling.Compute.MaxInstanceSize
			}
			clusterRequest.AutoScaling.Compute = compute
		}
	}

	log.Debugf("DEBUG: clusterRequest: %+v", clusterRequest)
	cluster, _, err := client.Clusters.Create(context.Background(), projectID, clusterRequest)
	if err != nil {
		log.Infof("Create - Cluster.Create() - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	currentModel.StateName = &cluster.StateName
	log.Debugf("Created cluster currentModel: %+v", currentModel)
	event := handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create Cluster `%s`", cluster.StateName),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 65,
		CallbackContext: map[string]interface{}{
			"stateName":        cluster.StateName,
			"projectId":        projectID,
			"clusterName":      *currentModel.Name,
			"deploymentSecret": secretName,
		},
	}
	log.Debugf("Create() return event:%+v", event)
	return event, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Read() currentModel:%+v", currentModel)

	// Callback is not called - Create() and Update() get recalled on
	// long running operations
	callback := map[string]interface{}(req.CallbackContext)
	log.Debugf("Read -  callback: %v", callback)
	if currentModel.Id == nil {
		err := errors.New("No Id found in currentModel")
		log.Infof("Read - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	secretName := *currentModel.Id
	log.Infof("Read for Cluster Id/SecretName:%s", secretName)
	key, err := util.GetApiKeyFromDeploymentSecret(&req, secretName)
	if err != nil {
		log.Infof("Read - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	log.Debugf("key:%+v", key)

	client, err := util.CreateMongoDBClient(key.PublicKey, key.PrivateKey)
	if err != nil {
		log.Infof("Read - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	// currentModel is NOT populated on the Read after long-running Cluster create
	// need to parse pid and cluster name from Id (deployment secret name).

	//projectID := *currentModel.ProjectId
	//clusterName := *currentModel.Name

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
	log.Debugf("parent --->%+v", id.Parent)

	projectID := id.Parent.ResourceID
	clusterName := id.ResourceID

	log.Debugf("Got projectID:%s, clusterName:%s, from id:%+v", projectID, clusterName, id)
	cluster, resp, err := client.Clusters.Get(context.Background(), projectID, clusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Infof("error 404- err:%+v resp:%+v", err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		} else {
			log.Infof("error cluster get- err:%+v resp:%+v", err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
		}
	}
	if currentModel.AutoScaling != nil {
		currentModel.AutoScaling = &AutoScaling{
			DiskGBEnabled: cluster.AutoScaling.DiskGBEnabled,
		}
		if currentModel.AutoScaling.Compute != nil {
			compute := &Compute{
				Enabled:          cluster.AutoScaling.Compute.Enabled,
				ScaleDownEnabled: cluster.AutoScaling.Compute.ScaleDownEnabled,
				MinInstanceSize:  &cluster.AutoScaling.Compute.MinInstanceSize,
				MaxInstanceSize:  &cluster.AutoScaling.Compute.MaxInstanceSize,
			}
			currentModel.AutoScaling.Compute = compute
		}

	}

	if currentModel.BackupEnabled != nil {
		currentModel.BackupEnabled = cluster.BackupEnabled
	}

	if currentModel.ProviderBackupEnabled != nil {
		currentModel.ProviderBackupEnabled = cluster.ProviderBackupEnabled
	}

	if currentModel.ClusterType != nil {
		currentModel.ClusterType = &cluster.ClusterType
	}
	if currentModel.DiskSizeGB != nil {
		currentModel.DiskSizeGB = cluster.DiskSizeGB
	}
	if currentModel.EncryptionAtRestProvider != nil {
		currentModel.EncryptionAtRestProvider = &cluster.EncryptionAtRestProvider
	}
	if currentModel.MongoDBMajorVersion != nil {
		currentModel.MongoDBMajorVersion = &cluster.MongoDBMajorVersion
	}

	if cluster.NumShards != nil {
		currentModel.NumShards = castNO64(cluster.NumShards)
	}

	currentModel.MongoDBVersion = &cluster.MongoDBVersion
	currentModel.MongoURI = &cluster.MongoURI
	currentModel.MongoURIUpdated = &cluster.MongoURIUpdated
	currentModel.MongoURIWithOptions = &cluster.MongoURIWithOptions
	currentModel.Paused = cluster.Paused
	currentModel.SrvAddress = &cluster.SrvAddress
	currentModel.StateName = &cluster.StateName

	if currentModel.BiConnector != nil {
		currentModel.BiConnector = &BiConnector{
			ReadPreference: &cluster.BiConnector.ReadPreference,
			Enabled:        cluster.BiConnector.Enabled,
		}
	}
	currentModel.ConnectionStrings = &ConnectionStrings{
		Standard:    &cluster.ConnectionStrings.Standard,
		StandardSrv: &cluster.ConnectionStrings.StandardSrv,
		Private:     &cluster.ConnectionStrings.Private,
		PrivateSrv:  &cluster.ConnectionStrings.PrivateSrv,
		//AwsPrivateLink:         &cluster.ConnectionStrings.AwsPrivateLink,
		//AwsPrivateLinkSrv:      &cluster.ConnectionStrings.AwsPrivateLinkSrv,
	}
	log.Debugf("step 2 cluster:+%v", cluster)

	/*
			if cluster.ProviderSettings != nil {
		        ps := &ProviderSettings{
					BackingProviderName: &cluster.ProviderSettings.BackingProviderName,
					DiskIOPS:            castNO64(cluster.ProviderSettings.DiskIOPS),
					EncryptEBSVolume:    cluster.ProviderSettings.EncryptEBSVolume,
					InstanceSizeName:    &cluster.ProviderSettings.InstanceSizeName,
					VolumeType:          &cluster.ProviderSettings.VolumeType,
				}
		        rn := util.EnsureAWSRegion(cluster.ProviderSettings.RegionName)
		        ps.RegionName = &rn
		        if currentModel.ProviderSettings.AutoScaling != nil {
		            ps.AutoScaling = &AutoScaling{
		                DiskGBEnabled: cluster.ProviderSettings.AutoScaling.DiskGBEnabled,
		            }
		            if currentModel.ProviderSettings.AutoScaling.Compute != nil {
		                compute := &Compute{}

		                if currentModel.ProviderSettings.AutoScaling.Compute.Enabled != nil {
		                    compute.Enabled = cluster.ProviderSettings.AutoScaling.Compute.Enabled
		                }
		                if currentModel.ProviderSettings.AutoScaling.Compute.ScaleDownEnabled != nil {
		                    compute.ScaleDownEnabled = cluster.ProviderSettings.AutoScaling.Compute.ScaleDownEnabled
		                }
		                if currentModel.ProviderSettings.AutoScaling.Compute.MinInstanceSize != nil {
		                    compute.MinInstanceSize = &cluster.ProviderSettings.AutoScaling.Compute.MinInstanceSize
		                }
		                if currentModel.ProviderSettings.AutoScaling.Compute.MaxInstanceSize != nil {
		                    compute.MaxInstanceSize = &cluster.ProviderSettings.AutoScaling.Compute.MaxInstanceSize
		                }
		                log.Debugf("compute -- what?> +%v",compute)
		                ps.AutoScaling.Compute = compute
		            }
		        }

		        currentModel.ProviderSettings = ps
			}

		    if currentModel.ReplicationSpecs != nil {
			    currentModel.ReplicationSpecs = flattenReplicationSpecs(cluster.ReplicationSpecs)
		    }

			if currentModel.ReplicationFactor != nil {
			    currentModel.ReplicationFactor = castNO64(cluster.ReplicationFactor)
		    }
	*/
	log.Debugf("Read() end currentModel:%+v", currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Update() currentModel:%+v", currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Infof("Update - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(client, req, currentModel, "IDLE", "UPDATING")
	}

	projectID := *currentModel.ProjectId
	clusterName := *currentModel.Name
	log.Debugf("Update - clusterName:%s", clusterName)
	if len(currentModel.ReplicationSpecs) > 0 {
		if currentModel.ClusterType != nil {
			err := errors.New("error creating cluster: ClusterType should be set when `ReplicationSpecs` is set")
			log.Infof("Update - error: %+v", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}

		if currentModel.NumShards != nil {
			err := errors.New("error creating cluster: NumShards should be set when `ReplicationSpecs` is set")
			log.Infof("Update - error: %+v", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
	}

	var autoScaling *mongodbatlas.AutoScaling
	if currentModel.AutoScaling != nil {
		autoScaling = &mongodbatlas.AutoScaling{
			DiskGBEnabled: currentModel.AutoScaling.DiskGBEnabled,
		}
	} else {
		autoScaling = &mongodbatlas.AutoScaling{}
	}

	log.Debugf("Update - autoScaling:%v", autoScaling)
	clusterRequest := &mongodbatlas.Cluster{
		Name:                     cast.ToString(currentModel.Name),
		EncryptionAtRestProvider: cast.ToString(currentModel.EncryptionAtRestProvider),
		ClusterType:              cast.ToString(currentModel.ClusterType),
		BackupEnabled:            currentModel.BackupEnabled,
		DiskSizeGB:               currentModel.DiskSizeGB,
		ProviderBackupEnabled:    currentModel.ProviderBackupEnabled,
		AutoScaling:              autoScaling,
	}
	if currentModel.BiConnector != nil {
		clusterRequest.BiConnector = expandBiConnector(currentModel.BiConnector)
	}
	if currentModel.ProviderSettings != nil {
		clusterRequest.ProviderSettings = expandProviderSettings(currentModel.ProviderSettings)
	}
	if currentModel.ReplicationSpecs != nil {
		clusterRequest.ReplicationSpecs = expandReplicationSpecs(currentModel.ReplicationSpecs)
	}
	if currentModel.ReplicationFactor != nil {
		clusterRequest.ReplicationFactor = cast64(currentModel.ReplicationFactor)
	}

	if currentModel.NumShards != nil {
		clusterRequest.NumShards = cast64(currentModel.NumShards)
	}

	if currentModel.MongoDBMajorVersion != nil {
		clusterRequest.MongoDBMajorVersion = formatMongoDBMajorVersion(*currentModel.MongoDBMajorVersion)
	}

	log.Debugf("Cluster update clusterRequest:%+v", clusterRequest)
	cluster, resp, err := client.Clusters.Update(context.Background(), projectID, clusterName, clusterRequest)
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
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Update Cluster `%s`", cluster.StateName),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 65,
		CallbackContext: map[string]interface{}{
			"stateName": cluster.StateName,
		},
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Delete() currentModel:%+v", currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Infof("Delete err: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(client, req, currentModel, "DELETED", "DELETING")
	}

	projectID := *currentModel.ProjectId
	clusterName := *currentModel.Name

	resp, err := client.Clusters.Delete(context.Background(), projectID, clusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Infof("Delete 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		} else {
			log.Infof("Delete err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
		}
	}
	mm := fmt.Sprintf("%s-Deleting", projectID)
	currentModel.Id = &mm
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Delete Complete",
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 65,
		CallbackContext: map[string]interface{}{
			"stateName": "DELETING",
		},
	}, nil
}

// List NOOP
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   currentModel,
	}, nil
}

func expandBiConnector(biConnector *BiConnector) *mongodbatlas.BiConnector {
	return &mongodbatlas.BiConnector{
		Enabled:        biConnector.Enabled,
		ReadPreference: cast.ToString(biConnector.ReadPreference),
	}
}

const (
	tenant  = "TENANT"
	atlasM2 = "M2"
	atlasM5 = "M5"
)

func (ps *ProviderSettings) providerName() string {
	if *ps.InstanceSizeName == atlasM2 || *ps.InstanceSizeName == atlasM5 {
		return tenant
	}
	return cast.ToString(ps.ProviderName)
}
func expandProviderSettings(providerSettings *ProviderSettings) *mongodbatlas.ProviderSettings {
	// convert AWS- regions to MDB regions
	log.Debugf("DEBUG: clusterRequest.ProviderSettings MODEL --->: %+v", providerSettings)
	regionName := util.EnsureAtlasRegion(*providerSettings.RegionName)
	providerName := providerSettings.providerName()

	var backingProviderName string
	if providerName == tenant {
		backingProviderName = cast.ToString(providerSettings.ProviderName)
	}

	ps := &mongodbatlas.ProviderSettings{
		EncryptEBSVolume:    providerSettings.EncryptEBSVolume,
		RegionName:          regionName,
		BackingProviderName: backingProviderName,
		InstanceSizeName:    cast.ToString(providerSettings.InstanceSizeName),
		ProviderName:        providerName,
		VolumeType:          cast.ToString(providerSettings.VolumeType),
	}
	if providerSettings.DiskIOPS != nil {
		ps.DiskIOPS = cast64(providerSettings.DiskIOPS)
	}
	log.Debugf("DEBUG: clusterRequest.ProviderSettings Atlas Requst struct --->: %+v", ps)
	return ps

}

func expandReplicationSpecs(replicationSpecs []ReplicationSpec) []mongodbatlas.ReplicationSpec {
	rSpecs := make([]mongodbatlas.ReplicationSpec, 0)

	for _, s := range replicationSpecs {
		rSpec := mongodbatlas.ReplicationSpec{
			ID:            cast.ToString(s.ID),
			NumShards:     cast64(s.NumShards),
			ZoneName:      cast.ToString(s.ZoneName),
			RegionsConfig: expandRegionsConfig(s.RegionsConfig),
		}

		rSpecs = append(rSpecs, rSpec)
	}
	return rSpecs
}

func expandRegionsConfig(regions []RegionsConfig) map[string]mongodbatlas.RegionsConfig {
	regionsConfig := make(map[string]mongodbatlas.RegionsConfig)
	for _, region := range regions {
		regionsConfig[*region.RegionName] = mongodbatlas.RegionsConfig{
			AnalyticsNodes: cast64(region.AnalyticsNodes),
			ElectableNodes: cast64(region.ElectableNodes),
			Priority:       cast64(region.Priority),
			ReadOnlyNodes:  cast64(region.ReadOnlyNodes),
		}
	}
	return regionsConfig
}

func formatMongoDBMajorVersion(val interface{}) string {
	if strings.Contains(val.(string), ".") {
		return val.(string)
	}
	return fmt.Sprintf("%.1f", cast.ToFloat32(val))
}

func flattenReplicationSpecs(rSpecs []mongodbatlas.ReplicationSpec) []ReplicationSpec {
	specs := make([]ReplicationSpec, 0)
	for _, rSpec := range rSpecs {
		spec := ReplicationSpec{
			ID:            &rSpec.ID,
			NumShards:     castNO64(rSpec.NumShards),
			ZoneName:      &rSpec.ZoneName,
			RegionsConfig: flattenRegionsConfig(rSpec.RegionsConfig),
		}
		specs = append(specs, spec)
	}
	return specs
}

func flattenRegionsConfig(regionsConfig map[string]mongodbatlas.RegionsConfig) []RegionsConfig {
	regions := make([]RegionsConfig, 0)

	for regionName, regionConfig := range regionsConfig {
		region := RegionsConfig{
			RegionName:     &regionName,
			Priority:       castNO64(regionConfig.Priority),
			AnalyticsNodes: castNO64(regionConfig.AnalyticsNodes),
			ElectableNodes: castNO64(regionConfig.ElectableNodes),
			ReadOnlyNodes:  castNO64(regionConfig.ReadOnlyNodes),
		}
		regions = append(regions, region)
	}
	return regions
}

func validateProgress(client *mongodbatlas.Client, req handler.Request, currentModel *Model, targetState string, pendingState string) (handler.ProgressEvent, error) {
	log.Debugf(" Cluster validateProgress() currentModel:%+v", currentModel)
	isReady, state, cluster, err := isClusterInTargetState(client, *currentModel.ProjectId, *currentModel.Name, targetState)
	log.Debugf("Cluster validateProgress() isReady:%+v, state:+%v, cluster:%+v", isReady, state, cluster)
	if err != nil {
		log.Debugf("ERROR Cluster validateProgress() err:%+v", err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}

	if !isReady {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = 60
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"stateName": state,
		}
		return p, nil
	}

	p := handler.NewProgressEvent()
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	if targetState != "DELETED" {
		p.ResourceModel = currentModel
	}
	return p, nil
}

func isClusterInTargetState(client *mongodbatlas.Client, projectID, clusterName, targetState string) (bool, string, *mongodbatlas.Cluster, error) {
	cluster, resp, err := client.Clusters.Get(context.Background(), projectID, clusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return "DELETED" == targetState, "DELETED", nil, nil
		}
		return false, "ERROR", nil, fmt.Errorf("error fetching cluster info (%s): %s", clusterName, err)
	}
	return cluster.StateName == targetState, cluster.StateName, cluster, nil
}
