package resource

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/spf13/cast"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectID.Value()

	if len(currentModel.ReplicationSpecs) > 0 {
		if currentModel.ClusterType.Value() != nil {
			return handler.ProgressEvent{}, errors.New("ClusterType should be set when `ReplicationSpecs` is set")
		}

		if currentModel.NumShards.Value() != nil {
			return handler.ProgressEvent{}, errors.New("NumShards should be set when `ReplicationSpecs` is set")
		}
	}

	autoScaling := mongodbatlas.AutoScaling{
		DiskGBEnabled: currentModel.AutoScaling.DiskGBEnabled.Value(),
	}

	clusterRequest := &mongodbatlas.Cluster{
		Name:                     *currentModel.Name.Value(),
		EncryptionAtRestProvider: *currentModel.EncryptionAtRestProvider.Value(),
		ClusterType:              *currentModel.ClusterType.Value(),
		BackupEnabled:            currentModel.BackupEnabled.Value(),
		DiskSizeGB:               currentModel.DiskSizeGB.Value(),
		ProviderBackupEnabled:    currentModel.ProviderBackupEnabled.Value(),
		AutoScaling:              autoScaling,
		BiConnector:              expandBiConnector(currentModel.BiConenctor),
		ProviderSettings:         expandProviderSettings(currentModel.ProviderSettings),
		ReplicationSpecs:         expandReplicationSpecs(currentModel.ReplicationSpecs),
		ReplicationFactor:        currentModel.ReplicationFactor.Value(),
		NumShards:                currentModel.NumShards.Value(),
	}

	if currentModel.MongoDBMajorVersion.Value() != nil {
		clusterRequest.MongoDBMajorVersion = formatMongoDBMajorVersion(*currentModel.MongoDBMajorVersion.Value())
	}

	cluster, _, err := client.Clusters.Create(context.Background(), projectID, clusterRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating cluster: %s", err)
	}

	log.Printf("[DEBUG] Cluster %+v", cluster)

	currentModel.ID = encoding.NewString(cluster.ID)

	if cluster.StateName == "IDLE" {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Create Cluster Complete",
			ResourceModel:   currentModel,
		}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.InProgress,
		Message:         fmt.Sprintf("Create Cluster `%s`", cluster.SrvAddress),
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectID.Value()
	clusterName := *currentModel.Name.Value()

	cluster, _, err := client.Clusters.Get(context.Background(), projectID, clusterName)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error fetching cluster info (%s): %s", clusterName, err)
	}

	currentModel.ID = encoding.NewString(cluster.ID)
	currentModel.AutoScaling = AutoScaling{
		DiskGBEnabled: encoding.NewBool(*cluster.AutoScaling.DiskGBEnabled),
	}

	currentModel.BackupEnabled = encoding.NewBool(*cluster.BackupEnabled)
	currentModel.ProviderBackupEnabled = encoding.NewBool(*cluster.ProviderBackupEnabled)
	currentModel.ClusterType = encoding.NewString(cluster.ClusterType)
	currentModel.DiskSizeGB = encoding.NewFloat(*cluster.DiskSizeGB)
	currentModel.EncryptionAtRestProvider = encoding.NewString(cluster.EncryptionAtRestProvider)
	currentModel.MongoDBMajorVersion = encoding.NewString(cluster.MongoDBVersion)

	if cluster.NumShards != nil {
		currentModel.NumShards = encoding.NewInt(*cluster.NumShards)
	}

	currentModel.MongoDBVersion = encoding.NewString(cluster.MongoDBVersion)
	currentModel.MongoURI = encoding.NewString(cluster.MongoURI)
	currentModel.MongoURIUpdated = encoding.NewString(cluster.MongoURIUpdated)
	currentModel.MongoURIWithOptions = encoding.NewString(cluster.MongoURIWithOptions)
	currentModel.Paused = encoding.NewBool(*cluster.Paused)
	currentModel.SrvAddress = encoding.NewString(cluster.SrvAddress)
	currentModel.StateName = encoding.NewString(cluster.StateName)

	currentModel.BiConenctor = BiConenctor{
		ReadPreference: encoding.NewString(cluster.BiConnector.ReadPreference),
		Enabled:        encoding.NewBool(*cluster.BiConnector.Enabled),
	}

	if cluster.ProviderSettings != nil {
		currentModel.ProviderSettings = ProviderSettings{
			BackingProviderName: encoding.NewString(cluster.ProviderSettings.BackingProviderName),
			DiskIOPS:            encoding.NewInt(*cluster.ProviderSettings.DiskIOPS),
			EncryptEBSVolume:    encoding.NewBool(*cluster.ProviderSettings.EncryptEBSVolume),
			InstanceSizeName:    encoding.NewString(cluster.ProviderSettings.InstanceSizeName),
			RegionName:          encoding.NewString(cluster.ProviderSettings.RegionName),
			VolumeType:          encoding.NewString(cluster.ProviderSettings.VolumeType),
		}
	}

	currentModel.ReplicationSpecs = flattenReplicationSpecs(cluster.ReplicationSpecs)
	currentModel.ReplicationFactor = encoding.NewInt(*cluster.ReplicationFactor)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Add your code here:
	// * Make API calls (use req.Session)
	// * Mutate the model
	// * Check/set any callback context (req.CallbackContext / response.CallbackContext)

	/*
	   // Construct a new handler.ProgressEvent and return it
	   response := handler.ProgressEvent{
	       OperationStatus: handler.Success,
	       Message: "Update complete",
	       ResourceModel: currentModel,
	   }

	   return response, nil
	*/

	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectID.Value()
	clusterName := *currentModel.Name.Value()

	_, err = client.Clusters.Delete(context.Background(), projectID, clusterName)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting cluster info (%s): %s", clusterName, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   currentModel,
	}, nil

}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Add your code here:
	// * Make API calls (use req.Session)
	// * Mutate the model
	// * Check/set any callback context (req.CallbackContext / response.CallbackContext)

	/*
	   // Construct a new handler.ProgressEvent and return it
	   response := handler.ProgressEvent{
	       OperationStatus: handler.Success,
	       Message: "List complete",
	       ResourceModel: currentModel,
	   }

	   return response, nil
	*/

	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: List")
}

func expandBiConnector(biConnector BiConenctor) mongodbatlas.BiConnector {
	return mongodbatlas.BiConnector{
		Enabled:        biConnector.Enabled.Value(),
		ReadPreference: *biConnector.ReadPreference.Value(),
	}
}

func expandProviderSettings(providerSettings ProviderSettings) *mongodbatlas.ProviderSettings {
	return &mongodbatlas.ProviderSettings{
		DiskIOPS:            providerSettings.DiskIOPS.Value(),
		EncryptEBSVolume:    providerSettings.EncryptEBSVolume.Value(),
		RegionName:          *providerSettings.RegionName.Value(),
		BackingProviderName: *providerSettings.BackingProviderName.Value(),
		InstanceSizeName:    *providerSettings.InstanceSizeName.Value(),
		ProviderName:        "AWS",
		VolumeType:          *providerSettings.VolumeType.Value(),
	}
}

func expandReplicationSpecs(replicationSpecs []ReplicationSpec) []mongodbatlas.ReplicationSpec {
	rSpecs := make([]mongodbatlas.ReplicationSpec, 0)

	for _, s := range replicationSpecs {
		rSpec := mongodbatlas.ReplicationSpec{
			ID:            *s.ID.Value(),
			NumShards:     s.NumShards.Value(),
			ZoneName:      *s.ZoneName.Value(),
			RegionsConfig: expandRegionsConfig(s.RegionsConfig),
		}

		rSpecs = append(rSpecs, rSpec)
	}
	return rSpecs
}

func expandRegionsConfig(regions []RegionsConfig) map[string]mongodbatlas.RegionsConfig {
	regionsConfig := make(map[string]mongodbatlas.RegionsConfig)
	for _, region := range regions {
		regionsConfig[*region.RegionName.Value()] = mongodbatlas.RegionsConfig{
			AnalyticsNodes: region.AnalyticsNodes.Value(),
			ElectableNodes: region.ElectableNodes.Value(),
			Priority:       region.Priority.Value(),
			ReadOnlyNodes:  region.ReadOnlyNodes.Value(),
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
			ID:            encoding.NewString(rSpec.ID),
			NumShards:     encoding.NewInt(*rSpec.NumShards),
			ZoneName:      encoding.NewString(rSpec.ZoneName),
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
			RegionName:     encoding.NewString(regionName),
			Priority:       encoding.NewInt(*regionConfig.Priority),
			AnalyticsNodes: encoding.NewInt(*regionConfig.AnalyticsNodes),
			ElectableNodes: encoding.NewInt(*regionConfig.ElectableNodes),
			ReadOnlyNodes:  encoding.NewInt(*regionConfig.ReadOnlyNodes),
		}
		regions = append(regions, region)
	}
	return regions
}
