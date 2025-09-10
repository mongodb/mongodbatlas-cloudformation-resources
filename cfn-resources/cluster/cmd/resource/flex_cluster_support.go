// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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
	flex "github.com/mongodb/mongodbatlas-cloudformation-resources/flex-cluster/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

const (
	defaultPriority = 7
	defaultZoneName = "Zone 1"
	flexProvider    = "FLEX"
)

// clusterToFlexModelIdentifier transforms a cluster model to a flex cluster model representation.
// It's used for Read and Delete where only the identifier is passed (project id and cluster name).
// As regions are not passed, Atlas calls are made to learn if it's a flex cluster.
// Returns nil if the cluster is not a flex cluster.
func clusterToFlexModelIdentifier(client *util.MongoDBClient, c *Model) *flex.Model {
	f := &flex.Model{
		Profile:   c.Profile,
		ProjectId: c.ProjectId,
		Name:      c.Name,
	}
	_, _, errFlex := client.AtlasSDK.FlexClustersApi.GetFlexCluster(context.Background(), *c.ProjectId, *c.Name).Execute()
	existingFlex := errFlex == nil
	if !existingFlex {
		return nil
	}
	return f
}

// clusterToFlexModelFull transforms a cluster model to a flex cluster model representation.
// It's used for Create and Update where the full model is passed.
// Regions are passed so not Atlas calls are needed to learn if it's a flex cluster.
// Returns nil if the cluster is not a flex cluster.
func clusterToFlexModelFull(c *Model) *flex.Model {
	if len(c.ReplicationSpecs) != 1 || len(c.ReplicationSpecs[0].AdvancedRegionConfigs) != 1 {
		return nil
	}
	firstRegion := c.ReplicationSpecs[0].AdvancedRegionConfigs[0]
	f := &flex.Model{
		Profile:                      c.Profile,
		ProjectId:                    c.ProjectId,
		Name:                         c.Name,
		Id:                           c.Id,
		StateName:                    c.StateName,
		ClusterType:                  c.ClusterType,
		CreateDate:                   c.CreatedDate,
		MongoDBVersion:               c.MongoDBVersion,
		TerminationProtectionEnabled: c.TerminationProtectionEnabled,
		VersionReleaseSystem:         c.VersionReleaseSystem,
		ProviderSettings: &flex.ProviderSettings{
			BackingProviderName: firstRegion.BackingProviderName,
			RegionName:          firstRegion.RegionName,
			ProviderName:        firstRegion.ProviderName,
		},
	}
	if c.BackupEnabled != nil {
		f.BackupSettings = &flex.BackupSettings{
			Enabled: c.BackupEnabled,
		}
	}
	if c.ConnectionStrings != nil {
		f.ConnectionStrings = &flex.ConnectionStrings{
			Standard:    c.ConnectionStrings.Standard,
			StandardSrv: c.ConnectionStrings.StandardSrv,
		}
	}
	f.Tags = make([]flex.Tag, len(c.Tags))
	for i, tag := range c.Tags {
		f.Tags[i] = flex.Tag{
			Key:   tag.Key,
			Value: tag.Value,
		}
	}
	return f
}

// fillModelForFlex updates the flex model into the cluster model.
func fillModelForFlex(pe *handler.ProgressEvent, c *Model) {
	if pe.ResourceModel == nil {
		return
	}
	f := pe.ResourceModel.(*flex.Model) // will panic if not a flex model
	pe.ResourceModel = c

	c.Profile = f.Profile
	c.ProjectId = f.ProjectId
	c.Name = f.Name
	c.Id = f.Id
	c.StateName = f.StateName
	c.ClusterType = f.ClusterType
	c.CreatedDate = f.CreateDate
	c.MongoDBVersion = f.MongoDBVersion
	c.TerminationProtectionEnabled = f.TerminationProtectionEnabled
	c.VersionReleaseSystem = f.VersionReleaseSystem
	if f.BackupSettings != nil {
		c.BackupEnabled = f.BackupSettings.Enabled
	} else {
		c.BackupEnabled = nil
	}
	if f.ConnectionStrings != nil {
		c.ConnectionStrings = &ConnectionStrings{
			Standard:    f.ConnectionStrings.Standard,
			StandardSrv: f.ConnectionStrings.StandardSrv,
		}
	} else {
		c.ConnectionStrings = nil
	}
	c.Tags = make([]Tag, len(f.Tags))
	for i, tag := range f.Tags {
		c.Tags[i] = Tag{
			Key:   tag.Key,
			Value: tag.Value,
		}
	}
	if f.ProviderSettings != nil {
		regionConfig := AdvancedRegionConfig{
			BackingProviderName: f.ProviderSettings.BackingProviderName,
			RegionName:          f.ProviderSettings.RegionName,
			ProviderName:        f.ProviderSettings.ProviderName,
			Priority:            util.Pointer(defaultPriority),
		}
		replicationSpec := AdvancedReplicationSpec{
			AdvancedRegionConfigs: []AdvancedRegionConfig{regionConfig},
			ZoneName:              util.Pointer(defaultZoneName),
		}
		c.ReplicationSpecs = []AdvancedReplicationSpec{replicationSpec}
	} else {
		c.ReplicationSpecs = nil
	}
}

// fillModelForFlexList update the cluster list with flex clusters.
func fillModelForFlexList(pe *handler.ProgressEvent, clusters []*Model) []*Model {
	if pe.ResourceModel == nil {
		return clusters
	}
	list := pe.ResourceModel.([]*flex.Model) // will panic if not a flex model array
	for _, f := range list {
		c := &Model{}
		fillModelForFlex(&handler.ProgressEvent{ResourceModel: f}, c)
		clusters = append(clusters, c)
	}
	return clusters
}
