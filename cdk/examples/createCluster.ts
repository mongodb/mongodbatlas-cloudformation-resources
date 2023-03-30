import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { CfnProject } from '@mongodbatlas-awscdk/project';
import { CfnCluster } from '@mongodbatlas-awscdk/cluster';

interface AtlasStackProps {
  readonly orgId: string;
  readonly profile: string;
  readonly projName: string;
  readonly clusterName: string;
  readonly clusterType: string;
  readonly instanceSize: string;
  readonly region: string;
}

export class CdkTestingStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const atlasProps = this.getContextProps();

    const projectRes = new CfnProject(this, 'ProjectResource', {
      name: atlasProps.projName,
      orgId: atlasProps.orgId,
      profile: atlasProps.profile
    });

    const clusterRes = new CfnCluster(this, 'ClusterResource', {
      name: atlasProps.clusterName,
      projectId: projectRes.attrId,
      profile: atlasProps.profile,
      clusterType: atlasProps.clusterType,
      backupEnabled: true,
      pitEnabled: false,
      replicationSpecs: [{
        numShards: 1,
        advancedRegionConfigs: [{
          autoScaling: {
            diskGb: {
              enabled: true,
            },
            compute: {
              enabled: false,
              scaleDownEnabled: false,
            },
          },
          analyticsSpecs: {
            ebsVolumeType: "STANDARD",
            instanceSize: atlasProps.instanceSize,
            nodeCount: 3,
          },
          electableSpecs: {
            ebsVolumeType: "STANDARD",
            instanceSize: atlasProps.instanceSize,
            nodeCount: 3,
          },
          readOnlySpecs: {
            ebsVolumeType: "STANDARD",
            instanceSize: atlasProps.instanceSize,
            nodeCount: 3,
          },
          priority: 7,
          regionName: atlasProps.region,
        }]
      }]
    });

  }

  getContextProps(): AtlasStackProps {
    const orgId = this.node.tryGetContext('orgId');
    if (!orgId){
      throw "No context value specified for orgId. Please specify via the cdk context."
    }
    const projName = this.node.tryGetContext('projName') ?? 'test-proj';
    const profile = this.node.tryGetContext('profile') ?? 'default';
    const clusterName = this.node.tryGetContext('clusterName') ?? 'test-cluster';
    const clusterType = this.node.tryGetContext('clusterType') ?? 'REPLICASET';
    const instanceSize = this.node.tryGetContext('instanceSize') ?? "M10";
    const region = this.node.tryGetContext('region') ?? "US_EAST_1";

    return {
      projName,
      orgId,
      profile,
      clusterName,
      clusterType,
      instanceSize,
      region,
    }
  }
}
