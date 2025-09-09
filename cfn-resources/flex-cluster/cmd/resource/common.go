// Copyright 2025 MongoDB Inc
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
	"context"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	cluster "github.com/mongodb/mongodbatlas-cloudformation-resources/cluster/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"go.mongodb.org/atlas-sdk/v20250312006/admin"
)

// ClusterToFlexModel transforms a Cluster Model to a FlexCluster Model representation.
// Returns nil if the cluster is not a flex cluster.
func ClusterToFlexModel(c *cluster.Model) *Model {
	if len(c.ReplicationSpecs) != 1 || len(c.ReplicationSpecs[0].AdvancedRegionConfigs) != 1 {
		return nil
	}
	firstRegion := c.ReplicationSpecs[0].AdvancedRegionConfigs[0]
	if firstRegion.ProviderName != nil &&
		*firstRegion.ProviderName != "FLEX" {
		return nil
	}

	f := &Model{
		Profile:                      c.Profile,
		ProjectId:                    c.ProjectId,
		Name:                         c.Name,
		Id:                           c.Id,
		StateName:                    c.StateName,
		ClusterType:                  c.ClusterType,
		CreateDate:                   c.CreatedDate, // Note: CreatedDate in cluster maps to CreateDate in flex
		MongoDBVersion:               c.MongoDBVersion,
		TerminationProtectionEnabled: c.TerminationProtectionEnabled,
		VersionReleaseSystem:         c.VersionReleaseSystem,
		ProviderSettings: &ProviderSettings{
			BackingProviderName: firstRegion.BackingProviderName,
			RegionName:          firstRegion.RegionName,
			ProviderName:        firstRegion.ProviderName,
			DiskSizeGB:          c.DiskSizeGB,
		},
	}

	if c.BackupEnabled != nil {
		f.BackupSettings = &BackupSettings{
			Enabled: c.BackupEnabled,
		}
	}

	if c.ConnectionStrings != nil {
		f.ConnectionStrings = &ConnectionStrings{
			Standard:    c.ConnectionStrings.Standard,
			StandardSrv: c.ConnectionStrings.StandardSrv,
		}
	}

	if len(c.Tags) > 0 {
		f.Tags = make([]Tag, len(c.Tags))
		for i, tag := range c.Tags {
			f.Tags[i] = Tag{
				Key:   tag.Key,
				Value: tag.Value,
			}
		}
	}

	return f
}

// FlexToClusterModel transforms a FlexCluster Model to a Cluster Model representation.
// This maps the subset of FlexCluster fields to their corresponding Cluster model fields.
// Creates a single replication spec from flex provider settings.
func FlexToClusterModel(f *Model) *cluster.Model {
	if f == nil {
		return nil
	}

	c := &cluster.Model{
		Profile:                      f.Profile,
		ProjectId:                    f.ProjectId,
		Name:                         f.Name,
		Id:                           f.Id,
		StateName:                    f.StateName,
		ClusterType:                  f.ClusterType,
		CreatedDate:                  f.CreateDate,
		MongoDBVersion:               f.MongoDBVersion,
		TerminationProtectionEnabled: f.TerminationProtectionEnabled,
		VersionReleaseSystem:         f.VersionReleaseSystem,
	}

	if f.BackupSettings != nil {
		c.BackupEnabled = f.BackupSettings.Enabled
	}

	if f.ConnectionStrings != nil {
		c.ConnectionStrings = &cluster.ConnectionStrings{
			Standard:    f.ConnectionStrings.Standard,
			StandardSrv: f.ConnectionStrings.StandardSrv,
		}
	}

	if len(f.Tags) > 0 {
		c.Tags = make([]cluster.Tag, len(f.Tags))
		for i, tag := range f.Tags {
			c.Tags[i] = cluster.Tag{
				Key:   tag.Key,
				Value: tag.Value,
			}
		}
	}

	if f.ProviderSettings != nil {
		// Create a single replication spec with one region config
		regionConfig := cluster.AdvancedRegionConfig{
			BackingProviderName: f.ProviderSettings.BackingProviderName,
			RegionName:          f.ProviderSettings.RegionName,
			ProviderName:        f.ProviderSettings.ProviderName,
			Priority:            util.Pointer(7), // Default priority for flex clusters
		}

		// Add basic electable specs for flex cluster
		regionConfig.ElectableSpecs = &cluster.Specs{
			InstanceSize: util.Pointer("FLEX"), // Flex clusters use FLEX instance size
		}

		replicationSpec := cluster.AdvancedReplicationSpec{
			NumShards:             util.Pointer(1), // Flex clusters typically have 1 shard
			AdvancedRegionConfigs: []cluster.AdvancedRegionConfig{regionConfig},
			ZoneName:              util.Pointer("Zone 1"), // Default zone name
		}

		c.ReplicationSpecs = []cluster.AdvancedReplicationSpec{replicationSpec}

		// Set DiskSizeGB at cluster level
		c.DiskSizeGB = f.ProviderSettings.DiskSizeGB
	}

	return c
}

// CreateFlexCluster calls Atlas API to create a flex cluster.
func CreateFlexCluster(client *util.MongoDBClient, model *Model) (*admin.FlexClusterDescription20241113, *handler.ProgressEvent) {
	flexReq := &admin.FlexClusterDescriptionCreate20241113{
		Name: *model.Name,
		ProviderSettings: admin.FlexProviderSettingsCreate20241113{
			BackingProviderName: *model.ProviderSettings.BackingProviderName,
			RegionName:          *model.ProviderSettings.RegionName,
		},
		TerminationProtectionEnabled: model.TerminationProtectionEnabled,
		Tags:                         expandTags(model.Tags),
	}
	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.CreateFlexCluster(context.Background(), *model.ProjectId, flexReq).Execute()
	return flexResp, util.HandleClusterError(err, resp)
}
