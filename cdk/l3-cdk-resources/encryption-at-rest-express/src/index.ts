import * as atlas from '@mongodbatlas-awscdk/atlas-basic';
import * as encryption from '@mongodbatlas-awscdk/atlas-encryption-at-rest';
import * as cluster from '@mongodbatlas-awscdk/cluster';
import { Construct } from 'constructs';


const clusterDefaults = {
   nodeCount: 3,
   nodeCountAnalytics: 1,
   backupEnabled: true,
   instanceSize: 'M30',
   mongoDBVersion: '5.0',
   encryptionAtRestProvider: cluster.CfnClusterPropsEncryptionAtRestProvider.AWS,
   region: 'US_EAST_1',
   ebsVolumeType: 'STANDARD',
 };


function getClusterProps(atlasBasicProps: atlas.AtlasBasicProps): atlas.ClusterProps {
   const inputClusterProps = atlasBasicProps.clusterProps;
   
   const clusterProps: atlas.ClusterProps = {
         name: inputClusterProps.name,
         mongoDbMajorVersion: inputClusterProps.mongoDbMajorVersion ? inputClusterProps.mongoDbMajorVersion : clusterDefaults.mongoDBVersion,
         apiKeys: atlasBasicProps.apiKeys,
         backupEnabled: inputClusterProps.backupEnabled ? inputClusterProps.backupEnabled : clusterDefaults.backupEnabled,
         diskSizeGb: inputClusterProps.diskSizeGb,
         clusterType: inputClusterProps.clusterType,
         biConnector: inputClusterProps.biConnector,
         connectionStrings: inputClusterProps.connectionStrings,
         encryptionAtRestProvider: inputClusterProps.encryptionAtRestProvider ? inputClusterProps.encryptionAtRestProvider : clusterDefaults.encryptionAtRestProvider,
         labels: inputClusterProps.labels,
         paused: inputClusterProps.paused,
         pitEnabled: inputClusterProps.pitEnabled,
         projectId: inputClusterProps.projectId,
         rootCertType: inputClusterProps.rootCertType,
         terminationProtectionEnabled: inputClusterProps.terminationProtectionEnabled,
         versionReleaseSystem: inputClusterProps.versionReleaseSystem,
         advancedSettings: inputClusterProps.advancedSettings,
         replicationSpecs: inputClusterProps.replicationSpecs ? inputClusterProps.replicationSpecs : getDefaultClusterReplicationSpec(),

   }

   return clusterProps;
}

function getDefaultClusterReplicationSpec(): cluster.AdvancedReplicationSpec[]{

   const replicationSpecs: cluster.AdvancedReplicationSpec[] = [
      {
        numShards: 1,
        advancedRegionConfigs: [
          {
            analyticsSpecs: {
              ebsVolumeType: clusterDefaults.ebsVolumeType,
              instanceSize: clusterDefaults.instanceSize,
              nodeCount: clusterDefaults.nodeCountAnalytics,
            },
            electableSpecs: {
              ebsVolumeType: clusterDefaults.ebsVolumeType,
              instanceSize: clusterDefaults.instanceSize,
              nodeCount: clusterDefaults.nodeCount,
            },
            priority: 7,
            regionName: clusterDefaults.region,
          },
        ],
      },
    ];

    return replicationSpecs;

   }



export interface AtlasEncryptionAtRestProps{
   /**
       * ID of an AWS IAM role authorized to manage an AWS customer master key.
       *
       * @schema AwsKms#RoleID
    */
   readonly roleId: string;
   /**
      * The AWS customer master key used to encrypt and decrypt the MongoDB master keys.
      *
      * @schema AwsKms#CustomerMasterKeyID
      */
   readonly customerMasterKeyId: string;
   /**
      * Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
      * Default Value: true
      * @schema AwsKms#Enabled
      */
   readonly enabledEncryptionAtRest?: boolean;
}

export interface AtlasEncryptionAtRestExpressProps extends atlas.AtlasBasicProps {
   readonly encryptionAtRest: AtlasEncryptionAtRestProps;
 
}

export class AtlasEncryptionAtRestExpress  extends Construct {
  readonly atlasBasic: atlas.AtlasBasic;
  readonly encryptionAtRest: encryption.AtlasEncryptionAtRest;

   constructor(scope: Construct, id: string, props: AtlasEncryptionAtRestExpressProps) {
      super(scope, id);
      this.atlasBasic = new atlas.AtlasBasic(this, 'atlas-basic-'.concat(id), {
         apiKeys: props.apiKeys,
         clusterProps: getClusterProps(props),
         projectProps: props.projectProps,
         dbUserProps: props.dbUserProps,
         ipAccessListProps: props.ipAccessListProps
       });

       this.encryptionAtRest = new encryption.AtlasEncryptionAtRest(this, 'atlas-encryption-at-rest-'.concat(id), {
         customerMasterKeyId: props.encryptionAtRest.customerMasterKeyId,
         roleId: props.encryptionAtRest.roleId,
         enabled: props.encryptionAtRest.enabledEncryptionAtRest,
         projectId: this.atlasBasic.mProject.ref,
         privateKey: props.apiKeys.privateKey!,
         publicKey: props.apiKeys.publicKey!,
       })
   }
}


