package resource

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

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

func getClusterRequest(model *Model) *mongodbatlas.Cluster {
	autoScaling := mongodbatlas.AutoScaling{
		DiskGBEnabled: model.AutoScaling.DiskGBEnabled,
	}

	clusterRequest := &mongodbatlas.Cluster{
		Name:                     cast.ToString(model.Name),
		EncryptionAtRestProvider: cast.ToString(model.EncryptionAtRestProvider),
		ClusterType:              cast.ToString(model.ClusterType),
		BackupEnabled:            model.BackupEnabled,
		DiskSizeGB:               model.DiskSizeGB,
		ProviderBackupEnabled:    model.ProviderBackupEnabled,
		AutoScaling:              &autoScaling,
		BiConnector:              expandBiConnector(model.BiConnector),
		ProviderSettings:         expandProviderSettings(model.ProviderSettings),
		ReplicationSpecs:         expandReplicationSpecs(model.ReplicationSpecs),
		ReplicationFactor:        cast64(model.ReplicationFactor),
		NumShards:                cast64(model.NumShards),
	}

	if model.MongoDBMajorVersion != nil {
		clusterRequest.MongoDBMajorVersion = formatMongoDBMajorVersion(*model.MongoDBMajorVersion)
	}
	return clusterRequest
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Printf("cluster Create")
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Printf("Error - %+v", err)
		return handler.ProgressEvent{}, err

	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(client, req, currentModel, "IDLE", "CREATING")
	}

	projectID := *currentModel.ProjectId
	log.Printf("cluster Create projectID=%s", projectID)
	if len(currentModel.ReplicationSpecs) > 0 {
		if currentModel.ClusterType != nil {
			return handler.ProgressEvent{}, errors.New("error creating cluster: ClusterType should be set when `ReplicationSpecs` is set")
		}

		if currentModel.NumShards != nil {
			return handler.ProgressEvent{}, errors.New("error creating cluster: NumShards should be set when `ReplicationSpecs` is set")
		}
	}

	var autoScaling *mongodbatlas.AutoScaling
	if currentModel.AutoScaling != nil {
		autoScaling = &mongodbatlas.AutoScaling{
			DiskGBEnabled: currentModel.AutoScaling.DiskGBEnabled,
		}
	}

	clusterRequest := &mongodbatlas.Cluster{
		Name:                     cast.ToString(currentModel.Name),
		EncryptionAtRestProvider: cast.ToString(currentModel.EncryptionAtRestProvider),
		ClusterType:              cast.ToString(currentModel.ClusterType),
		AutoScaling:              autoScaling,
		ReplicationFactor:        cast64(currentModel.ReplicationFactor),
		NumShards:                cast64(currentModel.NumShards),
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

	if currentModel.ReplicationSpecs != nil {
		clusterRequest.ReplicationSpecs = expandReplicationSpecs(currentModel.ReplicationSpecs)
	}

	cluster, resp, err := client.Clusters.Create(context.Background(), projectID, clusterRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating cluster: %w %v", err, &resp)
	}

	currentModel.StateName = &cluster.StateName

	// This is the intial call to Create, so inject a deployment
	// secret for this resource in order to lookup progress properly
	projectResID := &util.ResourceIdentifier{
		ResourceType: "Project",
		ResourceID:   projectID,
	}
	log.Printf("Created projectResID:%s", projectResID)
	resourceID := util.NewResourceIdentifier("Cluster", cluster.Name, projectResID)
	log.Printf("Created resourceID:%s", resourceID)
	resourceProps := map[string]string{
		"ClusterName": cluster.Name,
	}
	secretName, err := util.CreateDeploymentSecret(&req, resourceID, *currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey, &resourceProps)
	if err != nil {
		log.Printf("Error - %+v", err)
		return handler.ProgressEvent{}, err
	}

	log.Printf("Created new deployment secret for cluster. Secert Name = Cluster Id:%s", secretName)
	currentModel.Id = secretName

	return handler.ProgressEvent{
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
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	log.Printf("Read req:%+v, prevModel:%s, currentModel:%s", req, spew.Sdump(prevModel), spew.Sdump(currentModel))
	callback := map[string]interface{}(req.CallbackContext)
	log.Printf("Read -  callback: %v", callback)

	secretName := *currentModel.Id
	log.Printf("Read for Cluster Id/SecretName:%s", secretName)
	key, err := util.GetApiKeyFromDeploymentSecret(&req, secretName)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error lookupSecret: %w", err)
	}
	log.Printf("key:%+v", key)

	client, err := util.CreateMongoDBClient(key.PublicKey, key.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	// currentModel is NOT populated on the Read after long-running Cluster create
	// need to parse pid and cluster name from Id (deployment secret name).

	//projectID := *currentModel.ProjectId
	//clusterName := *currentModel.Name

	// key.ResourceID should == *currentModel.Id
	id, err := util.ParseResourceIdentifier(*currentModel.Id)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error parsing res if (%s): %s", id, err)
	}
	log.Printf("Parsed resource identifier: id:%+v", id)

	projectID := id.Parent.ResourceID
	clusterName := id.ResourceID

	log.Printf("Got projectID:%s, clusterName:%s, from id:%+v", projectID, clusterName, id)
	cluster, _, err := client.Clusters.Get(context.Background(), projectID, clusterName)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error fetching cluster info (%s): %s", clusterName, err)
	}

	currentModel.AutoScaling = &AutoScaling{
		DiskGBEnabled: cluster.AutoScaling.DiskGBEnabled,
	}

	currentModel.BackupEnabled = cluster.BackupEnabled
	currentModel.ProviderBackupEnabled = cluster.ProviderBackupEnabled
	currentModel.ClusterType = &cluster.ClusterType
	currentModel.DiskSizeGB = cluster.DiskSizeGB
	currentModel.EncryptionAtRestProvider = &cluster.EncryptionAtRestProvider
	currentModel.MongoDBMajorVersion = &cluster.MongoDBVersion

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

	currentModel.BiConnector = &BiConnector{
		ReadPreference: &cluster.BiConnector.ReadPreference,
		Enabled:        cluster.BiConnector.Enabled,
	}

	currentModel.ConnectionStrings = &ConnectionStrings{
		Standard:    &cluster.ConnectionStrings.Standard,
		StandardSrv: &cluster.ConnectionStrings.StandardSrv,
		Private:     &cluster.ConnectionStrings.Private,
		PrivateSrv:  &cluster.ConnectionStrings.PrivateSrv,
		//AwsPrivateLink:         &cluster.ConnectionStrings.AwsPrivateLink,
		//AwsPrivateLinkSrv:      &cluster.ConnectionStrings.AwsPrivateLinkSrv,
	}

	if cluster.ProviderSettings != nil {
		currentModel.ProviderSettings = &ProviderSettings{
			BackingProviderName: &cluster.ProviderSettings.BackingProviderName,
			DiskIOPS:            castNO64(cluster.ProviderSettings.DiskIOPS),
			EncryptEBSVolume:    cluster.ProviderSettings.EncryptEBSVolume,
			InstanceSizeName:    &cluster.ProviderSettings.InstanceSizeName,
			RegionName:          &cluster.ProviderSettings.RegionName,
			VolumeType:          &cluster.ProviderSettings.VolumeType,
		}
	}

	currentModel.ReplicationSpecs = flattenReplicationSpecs(cluster.ReplicationSpecs)
	currentModel.ReplicationFactor = castNO64(cluster.ReplicationFactor)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(client, req, currentModel, "IDLE", "UPDATING")
	}

	projectID := *currentModel.ProjectId
	clusterName := *currentModel.Name

	if len(currentModel.ReplicationSpecs) > 0 {
		if currentModel.ClusterType != nil {
			return handler.ProgressEvent{}, errors.New("error updating cluster: ClusterType should be set when `ReplicationSpecs` is set")
		}

		if currentModel.NumShards != nil {
			return handler.ProgressEvent{}, errors.New("error updating cluster: NumShards should be set when `ReplicationSpecs` is set")
		}
	}

	autoScaling := &mongodbatlas.AutoScaling{
		DiskGBEnabled: currentModel.AutoScaling.DiskGBEnabled,
	}

	clusterRequest := &mongodbatlas.Cluster{
		Name:                     cast.ToString(currentModel.Name),
		EncryptionAtRestProvider: cast.ToString(currentModel.EncryptionAtRestProvider),
		ClusterType:              cast.ToString(currentModel.ClusterType),
		BackupEnabled:            currentModel.BackupEnabled,
		DiskSizeGB:               currentModel.DiskSizeGB,
		ProviderBackupEnabled:    currentModel.ProviderBackupEnabled,
		AutoScaling:              autoScaling,
		BiConnector:              expandBiConnector(currentModel.BiConnector),
		ProviderSettings:         expandProviderSettings(currentModel.ProviderSettings),
		ReplicationSpecs:         expandReplicationSpecs(currentModel.ReplicationSpecs),
		ReplicationFactor:        cast64(currentModel.ReplicationFactor),
		NumShards:                cast64(currentModel.NumShards),
	}

	if currentModel.MongoDBMajorVersion != nil {
		clusterRequest.MongoDBMajorVersion = formatMongoDBMajorVersion(*currentModel.MongoDBMajorVersion)
	}

	cluster, _, err := client.Clusters.Update(context.Background(), projectID, clusterName, clusterRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating cluster: %s", err)
	}

	currentModel.Id = &cluster.ID

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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(client, req, currentModel, "DELETED", "DELETING")
	}

	projectID := *currentModel.ProjectId
	clusterName := *currentModel.Name

	_, err = client.Clusters.Delete(context.Background(), projectID, clusterName)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting cluster info (%s): %s", clusterName, err)
	}

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

func expandProviderSettings(providerSettings *ProviderSettings) *mongodbatlas.ProviderSettings {
	// convert AWS- regions to MDB regions
	regionName := util.EnsureAtlasRegion(*providerSettings.RegionName)
	ps := &mongodbatlas.ProviderSettings{
		EncryptEBSVolume:    providerSettings.EncryptEBSVolume,
		RegionName:          regionName,
		BackingProviderName: cast.ToString(providerSettings.BackingProviderName),
		InstanceSizeName:    cast.ToString(providerSettings.InstanceSizeName),
		ProviderName:        "AWS",
		VolumeType:          cast.ToString(providerSettings.VolumeType),
	}
	if providerSettings.DiskIOPS != nil {
		ps.DiskIOPS = cast64(providerSettings.DiskIOPS)
	}
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
	isReady, state, err := isClusterInTargetState(client, *currentModel.ProjectId, *currentModel.Name, targetState)
	if err != nil {
		return handler.ProgressEvent{}, err
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
	p.ResourceModel = currentModel
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

func isClusterInTargetState(client *mongodbatlas.Client, projectID, clusterName, targetState string) (bool, string, error) {
	cluster, resp, err := client.Clusters.Get(context.Background(), projectID, clusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return "DELETED" == targetState, "DELETED", nil
		}
		return false, "ERROR", fmt.Errorf("error fetching cluster info (%s): %s", clusterName, err)
	}
	return cluster.StateName == targetState, cluster.StateName, nil
}
