import * as atlas from '@mongodbatlas-awscdk/atlas-basic';
import * as cluster from '@mongodbatlas-awscdk/cluster';
import * as databaseUser from '@mongodbatlas-awscdk/database-user';
import * as encryption from '@mongodbatlas-awscdk/encryption-at-rest';
import * as accessList from '@mongodbatlas-awscdk/project-ip-access-list';
import { Construct } from 'constructs';


const NODE_COUNT = 3;
const NODE_COUNT_ANALYTICS = 1;
const BACKUP_ENABLED = true;
const INSTANCE_SIZE = 'M30';
const MONGODB_VERSION = '5.0';
const ENCRYPTIN_AT_REST_PROVIDER = cluster.CfnClusterPropsEncryptionAtRestProvider.AWS;
const REGION = 'US_EAST_1';
const EBS_VOLUME_TYPE = 'STANDARD';
const ENABLE_ENCRYPTION_AT_REST = true;
const CLUSTER_TYPE = 'REPLICASET';
const DB_NAME = 'admin';
const USERNAME = 'cdkUser';
const ROLE = [{
  roleName: 'atlasAdmin',
  databaseName: 'admin',
}];


function randomNumber() {
  const min = 10;
  const max = 9999999;
  return Math.floor(Math.random() * (max - min + 1) + min);
}

function getClusterProps(inputClusterProps: atlas.ClusterProps): cluster.CfnClusterProps {
  const clusterProps: cluster.CfnClusterProps = {
    name: inputClusterProps.name || 'atlas-cluster-'.concat(String(randomNumber())),
    mongoDbMajorVersion: inputClusterProps.mongoDbMajorVersion || MONGODB_VERSION,
    backupEnabled: inputClusterProps.backupEnabled || BACKUP_ENABLED,
    diskSizeGb: inputClusterProps.diskSizeGb,
    clusterType: inputClusterProps.clusterType || CLUSTER_TYPE,
    biConnector: inputClusterProps.biConnector,
    connectionStrings: inputClusterProps.connectionStrings,
    encryptionAtRestProvider: inputClusterProps.encryptionAtRestProvider || ENCRYPTIN_AT_REST_PROVIDER,
    labels: inputClusterProps.labels,
    paused: inputClusterProps.paused,
    pitEnabled: inputClusterProps.pitEnabled,
    rootCertType: inputClusterProps.rootCertType,
    terminationProtectionEnabled: inputClusterProps.terminationProtectionEnabled,
    versionReleaseSystem: inputClusterProps.versionReleaseSystem,
    advancedSettings: inputClusterProps.advancedSettings,
    replicationSpecs: inputClusterProps.replicationSpecs || getDefaultClusterReplicationSpec(),
    projectId: '',
  };

  return clusterProps;
}

function getDefaultClusterReplicationSpec(): cluster.AdvancedReplicationSpec[] {

  const replicationSpecs: cluster.AdvancedReplicationSpec[] = [
    {
      numShards: 1,
      advancedRegionConfigs: [
        {
          analyticsSpecs: {
            ebsVolumeType: EBS_VOLUME_TYPE,
            instanceSize: INSTANCE_SIZE,
            nodeCount: NODE_COUNT_ANALYTICS,
          },
          electableSpecs: {
            ebsVolumeType: EBS_VOLUME_TYPE,
            instanceSize: INSTANCE_SIZE,
            nodeCount: NODE_COUNT,
          },
          priority: 7,
          regionName: REGION,
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

  /**
     * The AWS region in which the AWS customer master key exists.
     *
     * @schema AwsKms#Region
     */
  readonly region?: string;
}

export interface AtlasEncryptionAtRestExpressProps {
  readonly cluster ?: atlas.ClusterProps;
  readonly accessList ?: atlas.IpAccessListProps;
  readonly encryptionAtRest: AtlasEncryptionAtRestProps;
  readonly databaseUser?: atlas.DatabaseUserProps;
  readonly profile?: string;
  readonly projectId : string;
}

export class AtlasEncryptionAtRestExpress extends Construct {
  readonly encryptionAtRest: encryption.CfnEncryptionAtRest;
  readonly cluster?: cluster.CfnCluster;
  readonly accessList?: accessList.CfnProjectIpAccessList;
  readonly databaseUser?: databaseUser.CfnDatabaseUser;

  constructor(scope: Construct, id: string, props: AtlasEncryptionAtRestExpressProps) {
    super(scope, id);

    this.encryptionAtRest = new encryption.CfnEncryptionAtRest(this, 'encryption-at-rest-'.concat(id), {
      awsKms: {
        customerMasterKeyId: props.encryptionAtRest.customerMasterKeyId,
        roleId: props.encryptionAtRest.roleId,
        enabled: props.encryptionAtRest.enabledEncryptionAtRest ? props.encryptionAtRest.enabledEncryptionAtRest : ENABLE_ENCRYPTION_AT_REST,
        region: props.encryptionAtRest.region ? props.encryptionAtRest.region : REGION,
      },
      projectId: props.projectId,
      profile: props.profile,
    });

    if (props.cluster) {
      // Create a new MongoDB Atlas Cluster and pass project ID
      const clusterProps = getClusterProps(props.cluster);
      this.cluster = new cluster.CfnCluster(this, 'cluster-'.concat(id),
        {
          ...clusterProps,
          profile: props.profile,
          projectId: props.projectId,
        });
    }

    if (props.databaseUser) {
      // Create a new MongoDB Atlas Database User
      this.databaseUser = new databaseUser.CfnDatabaseUser(this, 'db-user-'.concat(id),
        {
          ...props.databaseUser,
          profile: props.profile,
          databaseName: props.databaseUser?.databaseName || DB_NAME,
          projectId: props.projectId,
          username: props.databaseUser?.username || USERNAME,
          roles: props.databaseUser?.roles || ROLE,
          password: props.databaseUser.password,
        });
    }


    if (props.accessList) {
      // Create a new MongoDB Atlas Project IP Access List
      this.accessList = new accessList.CfnProjectIpAccessList(this, 'access-list-'.concat(id),
        {
          ...props.accessList,
          profile: props.profile,
          projectId: props.projectId,
        });
    }
  }
}


