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

package resource

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func mapClusterToModel(model *Model, cluster *admin.AdvancedClusterDescription) {
	model.Id = cluster.Id
	model.ProjectId = cluster.GroupId
	model.Name = cluster.Name
	model.BackupEnabled = cluster.BackupEnabled
	model.BiConnector = flattenBiConnectorConfig(cluster.BiConnector)
	model.ConnectionStrings = flattenConnectionStrings(cluster.ConnectionStrings)
	model.ClusterType = cluster.ClusterType
	model.CreatedDate = util.TimePtrToStringPtr(cluster.CreateDate)
	model.DiskSizeGB = cluster.DiskSizeGB
	model.EncryptionAtRestProvider = cluster.EncryptionAtRestProvider
	model.GlobalClusterSelfManagedSharding = cluster.GlobalClusterSelfManagedSharding
	model.Labels = flattenLabels(cluster.GetLabels())
	model.MongoDBMajorVersion = cluster.MongoDBMajorVersion
	model.MongoDBVersion = cluster.MongoDBVersion
	model.Paused = cluster.Paused
	model.PitEnabled = cluster.PitEnabled
	model.RootCertType = cluster.RootCertType
	model.ReplicationSpecs = flattenReplicationSpecs(cluster.GetReplicationSpecs())
	model.StateName = cluster.StateName
	model.VersionReleaseSystem = cluster.VersionReleaseSystem
}

func containsLabelOrKey(list []Labels, item Labels) bool {
	for _, v := range list {
		if reflect.DeepEqual(v, item) || *v.Key == *item.Key {
			return true
		}
	}

	return false
}

func expandBiConnector(biConnector *BiConnector) *admin.BiConnector {
	if biConnector == nil {
		return nil
	}
	return &admin.BiConnector{
		Enabled:        biConnector.Enabled,
		ReadPreference: biConnector.ReadPreference,
	}
}

func expandReplicationSpecs(replicationSpecs []AdvancedReplicationSpec) []admin.ReplicationSpec {
	rSpecs := []admin.ReplicationSpec{}

	for i := range replicationSpecs {
		var numShards int

		rSpec := admin.ReplicationSpec{
			NumShards:     &numShards,
			RegionConfigs: expandRegionsConfig(replicationSpecs[i].AdvancedRegionConfigs),
		}

		if util.IsStringPresent(replicationSpecs[i].ID) {
			rSpec.Id = admin.PtrString(cast.ToString(replicationSpecs[i].ID))
		}

		if replicationSpecs[i].NumShards != nil {
			rSpec.NumShards = replicationSpecs[i].NumShards
		}
		if replicationSpecs[i].ZoneName != nil {
			rSpec.ZoneName = admin.PtrString(cast.ToString(replicationSpecs[i].ZoneName))
		}
		rSpecs = append(rSpecs, rSpec)
	}

	fmt.Printf("specs: len %d %+v", len(replicationSpecs), rSpecs)
	return rSpecs
}

func expandAutoScaling(scaling *AdvancedAutoScaling) *admin.AdvancedAutoScalingSettings {
	advAutoScaling := &admin.AdvancedAutoScalingSettings{}
	if scaling == nil {
		return nil
	}
	if scaling.Compute != nil {
		advAutoScaling.Compute = &admin.AdvancedComputeAutoScaling{
			Enabled:          scaling.Compute.Enabled,
			ScaleDownEnabled: scaling.Compute.ScaleDownEnabled,
		}

		if util.IsStringPresent(scaling.Compute.MinInstanceSize) {
			advAutoScaling.Compute.MinInstanceSize = scaling.Compute.MinInstanceSize
		}

		if util.IsStringPresent(scaling.Compute.MaxInstanceSize) {
			advAutoScaling.Compute.MaxInstanceSize = scaling.Compute.MaxInstanceSize
		}
	}

	if scaling.DiskGB != nil {
		advAutoScaling.DiskGB = &admin.DiskGBAutoScaling{Enabled: scaling.DiskGB.Enabled}
	}

	return advAutoScaling
}

func expandRegionsConfig(regionConfigs []AdvancedRegionConfig) *[]admin.CloudRegionConfig {
	regionsConfigs := []admin.CloudRegionConfig{}
	for _, regionCfg := range regionConfigs {
		regionsConfigs = append(regionsConfigs, expandRegionConfig(regionCfg))
	}
	return &regionsConfigs
}

func expandRegionConfig(regionCfg AdvancedRegionConfig) admin.CloudRegionConfig {
	providerName := constants.AWS
	if regionCfg.ProviderName != nil {
		providerName = *regionCfg.ProviderName
	}

	advRegionConfig := admin.CloudRegionConfig{
		ProviderName: &providerName,
		RegionName:   regionCfg.RegionName,
		Priority:     regionCfg.Priority,
	}

	if regionCfg.AutoScaling != nil {
		advRegionConfig.AutoScaling = expandAutoScaling(regionCfg.AutoScaling)
	}
	if regionCfg.AnalyticsAutoScaling != nil {
		advRegionConfig.AnalyticsAutoScaling = expandAutoScaling(regionCfg.AnalyticsAutoScaling)
	}
	if regionCfg.AnalyticsSpecs != nil {
		advRegionConfig.AnalyticsSpecs = expandRegionConfigSpec(regionCfg.AnalyticsSpecs)
	}
	if regionCfg.ElectableSpecs != nil {
		advRegionConfig.ElectableSpecs = NewHardwareSpec(regionCfg.ElectableSpecs)
	}
	if regionCfg.ReadOnlySpecs != nil {
		advRegionConfig.ReadOnlySpecs = expandRegionConfigSpec(regionCfg.ReadOnlySpecs)
	}
	if regionCfg.BackingProviderName != nil {
		advRegionConfig.BackingProviderName = regionCfg.BackingProviderName
	}
	return advRegionConfig
}

func NewHardwareSpec(spec *Specs) *admin.HardwareSpec {
	if spec == nil {
		return nil
	}
	return &admin.HardwareSpec{
		DiskIOPS:      util.StrPtrToIntPtr(spec.DiskIOPS),
		EbsVolumeType: spec.EbsVolumeType,
		InstanceSize:  spec.InstanceSize,
		NodeCount:     spec.NodeCount,
	}
}

func expandRegionConfigSpec(spec *Specs) *admin.DedicatedHardwareSpec {
	if spec == nil {
		return nil
	}
	return &admin.DedicatedHardwareSpec{
		DiskIOPS:      util.StrPtrToIntPtr(spec.DiskIOPS),
		EbsVolumeType: spec.EbsVolumeType,
		InstanceSize:  spec.InstanceSize,
		NodeCount:     spec.NodeCount,
	}
}

func expandLabelSlice(labels []Labels) *[]admin.ComponentLabel {
	res := make([]admin.ComponentLabel, len(labels))

	for i := range labels {
		var key string
		if labels[i].Key != nil {
			key = *labels[i].Key
		}
		var value string
		if labels[i].Key != nil {
			value = *labels[i].Value
		}
		res[i] = admin.ComponentLabel{
			Key:   &key,
			Value: &value,
		}
	}
	return &res
}

func flattenAutoScaling(scaling *admin.AdvancedAutoScalingSettings) *AdvancedAutoScaling {
	if scaling == nil {
		return nil
	}
	advAutoScaling := &AdvancedAutoScaling{}

	if scaling.DiskGB != nil {
		advAutoScaling.DiskGB = &DiskGB{Enabled: scaling.DiskGB.Enabled}
	}
	if scaling.Compute != nil {
		compute := &Compute{}
		if scaling.Compute.Enabled != nil {
			compute.Enabled = scaling.Compute.Enabled
		}
		if scaling.Compute.ScaleDownEnabled != nil {
			compute.ScaleDownEnabled = scaling.Compute.ScaleDownEnabled
		}
		if util.IsStringPresent(scaling.Compute.MinInstanceSize) {
			compute.MinInstanceSize = scaling.Compute.MinInstanceSize
		}
		if util.IsStringPresent(scaling.Compute.MaxInstanceSize) {
			compute.MaxInstanceSize = scaling.Compute.MaxInstanceSize
		}

		advAutoScaling.Compute = compute
	}
	return advAutoScaling
}

func flattenReplicationSpecs(replicationSpecs []admin.ReplicationSpec) []AdvancedReplicationSpec {
	var rSpecs []AdvancedReplicationSpec

	for ind := range replicationSpecs {
		rSpec := AdvancedReplicationSpec{
			ID:                    replicationSpecs[ind].Id,
			NumShards:             replicationSpecs[ind].NumShards,
			ZoneName:              replicationSpecs[ind].ZoneName,
			AdvancedRegionConfigs: flattenRegionsConfig(replicationSpecs[ind].RegionConfigs),
		}
		rSpecs = append(rSpecs, rSpec)
	}
	fmt.Printf("specs: len %d %+v", len(replicationSpecs), rSpecs)
	return rSpecs
}

func flattenRegionsConfig(regionConfigs *[]admin.CloudRegionConfig) []AdvancedRegionConfig {
	if regionConfigs == nil {
		return []AdvancedRegionConfig{}
	}
	adminConfigs := *regionConfigs
	regionsConfigs := make([]AdvancedRegionConfig, 0, len(*regionConfigs))
	for i := range adminConfigs {
		adminConfig := adminConfigs[i]
		regionsConfigs = append(regionsConfigs, flattenRegionConfig(&adminConfig))
	}
	return regionsConfigs
}

func flattenRegionConfig(regionCfg *admin.CloudRegionConfig) AdvancedRegionConfig {
	if regionCfg == nil {
		return AdvancedRegionConfig{}
	}
	advRegConfig := AdvancedRegionConfig{
		AutoScaling:          flattenAutoScaling(regionCfg.AutoScaling),
		AnalyticsAutoScaling: flattenAutoScaling(regionCfg.AnalyticsAutoScaling),
		RegionName:           regionCfg.RegionName,
		Priority:             regionCfg.Priority,
	}
	if regionCfg.AnalyticsSpecs != nil {
		advRegConfig.AnalyticsSpecs = flattenRegionConfigSpec(regionCfg.AnalyticsSpecs)
	}
	if regionCfg.ElectableSpecs != nil {
		advRegConfig.ElectableSpecs = flattenElectableSpecs(regionCfg.ElectableSpecs)
	}

	if regionCfg.ReadOnlySpecs != nil {
		advRegConfig.ReadOnlySpecs = flattenRegionConfigSpec(regionCfg.ReadOnlySpecs)
	}

	return advRegConfig
}

func flattenElectableSpecs(spec *admin.HardwareSpec) *Specs {
	if spec == nil {
		return nil
	}
	return &Specs{
		DiskIOPS:      util.IntPtrToStrPtr(spec.DiskIOPS),
		EbsVolumeType: spec.EbsVolumeType,
		InstanceSize:  spec.InstanceSize,
		NodeCount:     spec.NodeCount,
	}
}

func flattenRegionConfigSpec(spec *admin.DedicatedHardwareSpec) *Specs {
	if spec == nil {
		return nil
	}
	return &Specs{
		DiskIOPS:      util.IntPtrToStrPtr(spec.DiskIOPS),
		EbsVolumeType: spec.EbsVolumeType,
		InstanceSize:  spec.InstanceSize,
		NodeCount:     spec.NodeCount,
	}
}

func flattenBiConnectorConfig(biConnector *admin.BiConnector) *BiConnector {
	if biConnector == nil {
		return nil
	}
	return &BiConnector{
		ReadPreference: biConnector.ReadPreference,
		Enabled:        biConnector.Enabled,
	}
}

type privateEndpointConnectionStrings struct {
	PrivateEndpoints                  []string
	PrivateEndpointsSrv               []string
	SRVShardOptimizedConnectionString []string
}

func flattenConnectionStrings(clusterConnStrings *admin.ClusterConnectionStrings) (connStrings *ConnectionStrings) {
	if clusterConnStrings != nil {
		privateEndpoints := flattenPrivateEndpoint(clusterConnStrings.PrivateEndpoint)
		connStrings = &ConnectionStrings{
			Standard:                          clusterConnStrings.Standard,
			StandardSrv:                       clusterConnStrings.StandardSrv,
			Private:                           clusterConnStrings.Private,
			PrivateSrv:                        clusterConnStrings.PrivateSrv,
			PrivateEndpoints:                  privateEndpoints.PrivateEndpoints,
			PrivateEndpointsSrv:               privateEndpoints.PrivateEndpointsSrv,
			SRVShardOptimizedConnectionString: privateEndpoints.SRVShardOptimizedConnectionString,
		}
	}
	return
}

func flattenPrivateEndpoint(pes *[]admin.ClusterDescriptionConnectionStringsPrivateEndpoint) privateEndpointConnectionStrings {
	privateEndpoints := privateEndpointConnectionStrings{
		PrivateEndpoints:                  make([]string, 0),
		PrivateEndpointsSrv:               make([]string, 0),
		SRVShardOptimizedConnectionString: make([]string, 0),
	}
	if pes == nil {
		return privateEndpoints
	}

	for _, pe := range *pes {
		if util.IsStringPresent(pe.ConnectionString) {
			privateEndpoints.PrivateEndpoints = append(privateEndpoints.PrivateEndpoints, *pe.ConnectionString)
		}

		if util.IsStringPresent(pe.SrvConnectionString) {
			privateEndpoints.PrivateEndpointsSrv = append(privateEndpoints.PrivateEndpointsSrv, *pe.SrvConnectionString)
		}

		if util.IsStringPresent(pe.SrvShardOptimizedConnectionString) {
			privateEndpoints.SRVShardOptimizedConnectionString = append(privateEndpoints.SRVShardOptimizedConnectionString, *pe.SrvShardOptimizedConnectionString)
		}
	}
	return privateEndpoints
}

func flattenProcessArgs(p *admin.ClusterDescriptionProcessArgs) *ProcessArgs {
	return &ProcessArgs{
		DefaultReadConcern:               p.DefaultReadConcern,
		DefaultWriteConcern:              p.DefaultWriteConcern,
		FailIndexKeyTooLong:              p.FailIndexKeyTooLong,
		JavascriptEnabled:                p.JavascriptEnabled,
		MinimumEnabledTLSProtocol:        p.MinimumEnabledTlsProtocol,
		NoTableScan:                      p.NoTableScan,
		OplogSizeMB:                      p.OplogSizeMB,
		SampleSizeBIConnector:            p.SampleSizeBIConnector,
		SampleRefreshIntervalBIConnector: p.SampleRefreshIntervalBIConnector,
		OplogMinRetentionHours:           p.OplogMinRetentionHours,
		TransactionLifetimeLimitSeconds:  util.Int64PtrToIntPtr(p.TransactionLifetimeLimitSeconds),
	}
}

func flattenLabels(clusterLabels []admin.ComponentLabel) []Labels {
	labels := make([]Labels, len(clusterLabels))
	for i := range clusterLabels {
		labels[i] = Labels{
			Key:   clusterLabels[i].Key,
			Value: clusterLabels[i].Value,
		}
	}
	return labels
}

func expandAdvancedSettings(processArgs ProcessArgs) *admin.ClusterDescriptionProcessArgs {
	var args admin.ClusterDescriptionProcessArgs

	if processArgs.DefaultReadConcern != nil {
		args.DefaultReadConcern = processArgs.DefaultReadConcern
	}
	args.FailIndexKeyTooLong = processArgs.FailIndexKeyTooLong
	if processArgs.DefaultWriteConcern != nil {
		args.DefaultWriteConcern = processArgs.DefaultWriteConcern
	}
	args.JavascriptEnabled = processArgs.JavascriptEnabled
	if processArgs.MinimumEnabledTLSProtocol != nil {
		args.MinimumEnabledTlsProtocol = processArgs.MinimumEnabledTLSProtocol
	}
	args.NoTableScan = processArgs.NoTableScan

	if processArgs.OplogSizeMB != nil {
		args.OplogSizeMB = processArgs.OplogSizeMB
	}
	if processArgs.SampleSizeBIConnector != nil {
		args.SampleSizeBIConnector = processArgs.SampleSizeBIConnector
	}
	if processArgs.SampleRefreshIntervalBIConnector != nil {
		args.SampleRefreshIntervalBIConnector = processArgs.SampleRefreshIntervalBIConnector
	}

	if processArgs.OplogMinRetentionHours != nil {
		args.OplogMinRetentionHours = processArgs.OplogMinRetentionHours
	}

	if processArgs.TransactionLifetimeLimitSeconds != nil {
		args.TransactionLifetimeLimitSeconds = cast64(processArgs.TransactionLifetimeLimitSeconds)
	}

	return &args
}

func flattenTags(clusterTags []admin.ResourceTag) (tags []Tag) {
	for ind := range clusterTags {
		tags = append(tags, Tag{
			Key:   &clusterTags[ind].Key,
			Value: &clusterTags[ind].Value,
		})
	}
	return
}

func expandTags(tags []Tag) (*[]admin.ResourceTag, error) {
	clusterTags := []admin.ResourceTag{}
	for ind := range tags {
		key := tags[ind].Key
		value := tags[ind].Value
		if key == nil {
			return &clusterTags, errors.New("tags Key is undefined")
		}
		if value == nil {
			return &clusterTags, fmt.Errorf("tags Value is undefined for %s", *key)
		}
		clusterTags = append(clusterTags, admin.ResourceTag{
			Key:   *key,
			Value: *value,
		})
	}
	return &clusterTags, nil
}

func setClusterData(currentModel *Model, cluster *admin.AdvancedClusterDescription) {
	if cluster == nil {
		return
	}

	currentModel.ProjectId = cluster.GroupId
	currentModel.Name = cluster.Name
	currentModel.Id = cluster.Id

	if currentModel.BackupEnabled != nil {
		currentModel.BackupEnabled = cluster.BackupEnabled
	}
	if currentModel.BiConnector != nil {
		currentModel.BiConnector = flattenBiConnectorConfig(cluster.BiConnector)
	}
	// Readonly
	currentModel.ConnectionStrings = flattenConnectionStrings(cluster.ConnectionStrings)
	if currentModel.ClusterType != nil {
		currentModel.ClusterType = cluster.ClusterType
	}
	// Readonly
	currentModel.CreatedDate = util.TimePtrToStringPtr(cluster.CreateDate)
	if currentModel.DiskSizeGB != nil {
		currentModel.DiskSizeGB = cluster.DiskSizeGB
	}
	if currentModel.EncryptionAtRestProvider != nil {
		currentModel.EncryptionAtRestProvider = cluster.EncryptionAtRestProvider
	}
	if currentModel.Labels != nil {
		currentModel.Labels = flattenLabels(cluster.GetLabels())
	}
	if currentModel.MongoDBMajorVersion != nil {
		currentModel.MongoDBMajorVersion = cluster.MongoDBMajorVersion
	}
	// Readonly
	currentModel.MongoDBVersion = cluster.MongoDBVersion

	if currentModel.Paused != nil {
		currentModel.Paused = cluster.Paused
	}
	if currentModel.PitEnabled != nil {
		currentModel.PitEnabled = cluster.PitEnabled
	}
	if currentModel.RootCertType != nil {
		currentModel.RootCertType = cluster.RootCertType
	}
	if currentModel.ReplicationSpecs != nil {
		currentModel.ReplicationSpecs = flattenReplicationSpecs(cluster.GetReplicationSpecs())
	}
	// Readonly
	if currentModel.GlobalClusterSelfManagedSharding == nil {
		currentModel.GlobalClusterSelfManagedSharding = cluster.GlobalClusterSelfManagedSharding
	}
	// Readonly
	currentModel.StateName = cluster.StateName
	if currentModel.VersionReleaseSystem != nil {
		currentModel.VersionReleaseSystem = cluster.VersionReleaseSystem
	}

	currentModel.TerminationProtectionEnabled = cluster.TerminationProtectionEnabled
	currentModel.Tags = flattenTags(cluster.GetTags())
}

func setClusterRequest(currentModel *Model) (*admin.AdvancedClusterDescription, *handler.ProgressEvent) {
	clusterRequest := &admin.AdvancedClusterDescription{
		Name: currentModel.Name,
	}
	if currentModel.ReplicationSpecs != nil {
		adminRepSpecs := expandReplicationSpecs(currentModel.ReplicationSpecs)
		clusterRequest.ReplicationSpecs = &adminRepSpecs
	}

	if currentModel.EncryptionAtRestProvider != nil {
		clusterRequest.EncryptionAtRestProvider = currentModel.EncryptionAtRestProvider
	}

	if currentModel.ClusterType != nil {
		clusterRequest.ClusterType = currentModel.ClusterType
	}

	if currentModel.BackupEnabled != nil {
		clusterRequest.BackupEnabled = currentModel.BackupEnabled
	}

	if currentModel.BiConnector != nil {
		clusterRequest.BiConnector = expandBiConnector(currentModel.BiConnector)
	}

	if currentModel.DiskSizeGB != nil {
		clusterRequest.DiskSizeGB = currentModel.DiskSizeGB
	}

	if currentModel.GlobalClusterSelfManagedSharding != nil {
		clusterRequest.GlobalClusterSelfManagedSharding = currentModel.GlobalClusterSelfManagedSharding
	}

	if currentModel.Labels != nil {
		clusterRequest.Labels = expandLabelSlice(currentModel.Labels)
	}

	if currentModel.MongoDBMajorVersion != nil {
		clusterRequest.MongoDBMajorVersion = admin.PtrString(formatMongoDBMajorVersion(*currentModel.MongoDBMajorVersion))
	}

	if currentModel.PitEnabled != nil {
		clusterRequest.PitEnabled = currentModel.PitEnabled
	}

	if currentModel.VersionReleaseSystem != nil {
		clusterRequest.VersionReleaseSystem = currentModel.VersionReleaseSystem
	}

	if currentModel.RootCertType != nil {
		clusterRequest.RootCertType = currentModel.RootCertType
	}
	tags, err := expandTags(currentModel.Tags)
	if err != nil {
		return clusterRequest, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
		}
	}
	clusterRequest.Tags = tags

	clusterRequest.TerminationProtectionEnabled = currentModel.TerminationProtectionEnabled
	return clusterRequest, nil
}

func AddReplicationSpecIDs(src, dest []admin.ReplicationSpec) *[]admin.ReplicationSpec {
	zoneToID := map[string]string{}
	providerRegionToID := map[string]string{}
	usedIDs := map[string]bool{}

	for _, spec := range src {
		specID := spec.GetId()
		if specID == "" {
			continue
		}
		if zoneName := spec.GetZoneName(); zoneName != "" {
			zoneToID[zoneName] = specID
		}
		if providerRegion := asProviderRegion(spec); providerRegion != "" {
			providerRegionToID[providerRegion] = spec.GetId()
		}
	}
	for i, spec := range dest {
		specID := spec.GetId()
		if specID != "" {
			continue
		}
		idZone, foundZone := zoneToID[spec.GetZoneName()]
		zoneUsed := usedIDs[idZone]
		if foundZone && !zoneUsed {
			usedIDs[idZone] = true
			dest[i].SetId(idZone)
			continue
		}
		idProvider, foundProvider := providerRegionToID[asProviderRegion(spec)]
		providerUsed := usedIDs[idProvider]
		if foundProvider && !providerUsed {
			usedIDs[idProvider] = true
			dest[i].SetId(idProvider)
			continue
		}
	}
	return &dest
}

func asProviderRegion(spec admin.ReplicationSpec) string {
	configs := spec.GetRegionConfigs()
	if len(configs) == 0 {
		return ""
	}
	return fmt.Sprintf("%s-%s", configs[0].GetProviderName(), configs[0].GetRegionName())
}
