import * as atlas from '@mongodbatlas-awscdk/cluster';
import * as user from '@mongodbatlas-awscdk/database-user';
import * as project from '@mongodbatlas-awscdk/project';
import * as ipAccessList from '@mongodbatlas-awscdk/project-ip-access-list';
import { Construct } from 'constructs';

/** @type {*} */
const projectDefaults = {
  projectName: 'atlas-project-',
};
/** @type {*} */
const ipAccessDefaults = {
  accessList: [
    {
      ipAddress: '0.0.0.0/1',
      comment: 'open all ips',
    },
  ],
};
/** @type {*} */
const dbDefaults = {
  dbName: 'admin',
  username: 'atlas-user',
  password: 'atlas-pwd',
  roles: [{
    roleName: 'atlasAdmin',
    databaseName: 'admin',
  }],
};
/** @type {*} */
const clusterDefaults = {
  clusterName: 'atlas-cluster-',
  clusterType: 'REPLICASET',
};

/**
 * @description
 * @export
 * @interface ProjectProps
 */
export interface ProjectProps {
  /**
     * @description
     * @type {string}
     * @memberof ProjectProps
     */
  readonly name?: string;
  /**
     * @description
     * @type {string}
     * @memberof ProjectProps
     * @default auto-generated
     */
  readonly orgId: string;
  /**
     * @description
     * @type {string}
     * @memberof ProjectProps
     */
  readonly projectOwnerId?: string;
  /**
     * @description
     * @type {boolean}
     * @memberof ProjectProps
     */
  readonly withDefaultAlertsSettings?: boolean;
  /**
     * @description
     * @type {number}
     * @memberof ProjectProps
     */
  readonly clusterCount?: number;
  /**
     * @description
     * @type {project.ProjectSettings}
     * @memberof ProjectProps
     */
  readonly projectSettings?: project.ProjectSettings;
  /**
     * @description
     * @type {project.ProjectTeam[]}
     * @memberof ProjectProps
     */
  readonly projectTeams?: project.ProjectTeam[];
  /**
     * @description
     * @type {project.ProjectApiKey[]}
     * @memberof ProjectProps
     */
  readonly projectApiKeys?: project.ProjectApiKey[];
}
/**
 * @description
 * @export
 * @interface ClusterProps
 */
export interface ClusterProps {
  /**
     * @description
     * @type {atlas.ProcessArgs}
     * @memberof ClusterProps
     */
  readonly advancedSettings?: atlas.ProcessArgs;
  /**
     * @description
     * @type {atlas.ApiKeyDefinition}
     * @memberof ClusterProps
     */
  readonly apiKeys?: atlas.ApiKeyDefinition;
  /**
     * @description
     * @type {boolean}
     * @memberof ClusterProps
     */
  readonly backupEnabled?: boolean;
  /**
     * @description
     * @type {atlas.CfnClusterPropsBiConnector}
     * @memberof ClusterProps
     */
  readonly biConnector?: atlas.CfnClusterPropsBiConnector;
  /**
     * @description
     * @type {string}
     * @memberof ClusterProps
     */
  readonly clusterType?: string;
  /**
     * @description
     * @type {atlas.ConnectionStrings}
     * @memberof ClusterProps
     * @default REPLICASET
     */
  readonly connectionStrings?: atlas.ConnectionStrings;
  /**
     * @description
     * @type {number}
     * @memberof ClusterProps
     */
  readonly diskSizeGb?: number;
  /**
     * @description
     * @type {atlas.CfnClusterPropsEncryptionAtRestProvider}
     * @memberof ClusterProps
     */
  readonly encryptionAtRestProvider?: atlas.CfnClusterPropsEncryptionAtRestProvider;
  /**
     * @description
     * @type {string}
     * @memberof ClusterProps
     */
  readonly projectId?: string;
  /**
     * @description
     * @type {atlas.CfnClusterPropsLabels[]}
     * @memberof ClusterProps
     */
  readonly labels?: atlas.CfnClusterPropsLabels[];
  /**
     * @description
     * @type {string}
     * @memberof ClusterProps
     */
  readonly mongoDbMajorVersion?: string;
  /**
     * @description
     * @type {string}
     * @memberof ClusterProps
     */
  readonly name?: string;
  /**
     * @description
     * @type {boolean}
     * @memberof ClusterProps
     * @default auto-generated
     */
  readonly paused?: boolean;
  /**
     * @description
     * @type {boolean}
     * @memberof ClusterProps
     */
  readonly pitEnabled?: boolean;
  /**
     * @description
     * @type {atlas.AdvancedReplicationSpec[]}
     * @memberof ClusterProps
     */
  readonly replicationSpecs?: atlas.AdvancedReplicationSpec[];
  /**
     * @description
     * @type {string}
     * @memberof ClusterProps
     */
  readonly rootCertType?: string;
  /**
     * @description
     * @type {string}
     * @memberof ClusterProps
     */
  readonly versionReleaseSystem?: string;
  /**
     * @description
     * @type {boolean}
     * @memberof ClusterProps
     */
  readonly terminationProtectionEnabled?: boolean;
}
/**
 * @description
 * @export
 * @interface DatabaseUserProps
 */
export interface DatabaseUserProps {
  /**
     * @description
     * @type {string}
     * @memberof DatabaseUserProps
     */
  readonly deleteAfterDate?: string;
  /**
     * @description
     * @type {user.CfnDatabaseUserPropsAwsiamType}
     * @memberof DatabaseUserProps
     */
  readonly awsiamType?:user.CfnDatabaseUserPropsAwsiamType;
  /**
     * @description
     * @type {string}
     * @memberof DatabaseUserProps
     */
  readonly databaseName?: string;
  /**
     * @description
     * @type {user.LabelDefinition[]}
     * @memberof DatabaseUserProps
     * @default admin
     */
  readonly labels?: user.LabelDefinition[];
  /**
     * @description
     * @type {user.CfnDatabaseUserPropsLdapAuthType}
     * @memberof DatabaseUserProps
     */
  readonly ldapAuthType?: user.CfnDatabaseUserPropsLdapAuthType;
  /**
     * @description
     * @type {user.CfnDatabaseUserPropsX509Type}
     * @memberof DatabaseUserProps
     */
  readonly x509Type?: user.CfnDatabaseUserPropsX509Type;
  /**
     * @description
     * @type {string}
     * @memberof DatabaseUserProps
     */
  readonly password?: string;
  /**
     * @description
     * @type {string}
     * @default cdk-pwd
     * @memberof DatabaseUserProps
     */
  readonly projectId?: string;
  /**
     * @description
     * @type {user.RoleDefinition[]}
     * @memberof DatabaseUserProps
     */
  readonly roles?: user.RoleDefinition[];
  /**
     * @description
     * @type {user.ScopeDefinition[]}
     * @memberof DatabaseUserProps
     */
  readonly scopes?: user.ScopeDefinition[];
  /**
     * @description
     * @type {string}
     * @memberof DatabaseUserProps
     * @default cdk-user
     */
  readonly username?: string;
}
/**
 * @description
 * @export
 * @interface IpAccessListProps
 */
export interface IpAccessListProps {
  /**
     * @description
     * @type {ipAccessList.AccessListDefinition[]}
     * @memberof IpAccessListProps
     */
  readonly accessList: ipAccessList.AccessListDefinition[];
  /**
     * @description
     * @type {string}
     * @memberof IpAccessListProps
     * @default allow-all
     */
  readonly projectId?: string;
  /**
     * @description
     * @type {number}
     * @memberof IpAccessListProps
     */
  readonly totalCount?: number;
  /**
     * @description
     * @type {ipAccessList.ListOptions}
     * @memberof IpAccessListProps
     */
  readonly listOptions?: ipAccessList.ListOptions;
}
/**
 * @description
 * @export
 * @interface ApiKeyDefinition
 */
export interface ApiKeyDefinition {
  /**
     * @description
     * @type {string}
     * @memberof ApiKeyDefinition
     */
  readonly privateKey?: string;
  /**
     * @description
     * @type {string}
     * @memberof ApiKeyDefinition
     */
  readonly publicKey?: string;

}
/**
 * @description
 * @export
 * @interface AtlasBasicProps
 */
export interface AtlasBasicProps {
  /**
     * @description
     * @type {ApiKeyDefinition}
     * @memberof AtlasBasicProps
     */
  readonly apiKeys : ApiKeyDefinition;
  /**
     * @description
     * @type {ProjectProps}
     * @memberof AtlasBasicProps
     */
  readonly projectProps: ProjectProps;
  /**
     * @description
     * @type {ClusterProps}
     * @memberof AtlasBasicProps
     */
  readonly clusterProps: ClusterProps;
  /**
     * @description
     * @type {DatabaseUserProps}
     * @memberof AtlasBasicProps
     */
  readonly dbUserProps?: DatabaseUserProps;
  /**
     * @description
     * @type {IpAccessListProps}
     * @memberof AtlasBasicProps
     */
  readonly ipAccessListProps?: IpAccessListProps;
}
/**
 * @description
 * @export
 * @class AtlasBasic
 * @extends {Construct}
 */
export class AtlasBasic extends Construct {
  /**
   * Creates an instance of AtlasBasic.
   * @param {Construct} scope
   * @param {string} id
   * @param {AtlasBasicProps} props
   * @memberof AtlasBasic
   */
  constructor(scope: Construct, id: string, props: AtlasBasicProps) {
    super(scope, id);
    //Create a new MongoDB Atlas Project
    const mProject = new project.CfnProject(this, 'project-'.concat(id), {
      apiKeys: props.apiKeys,
      name: props.projectProps.name || projectDefaults.projectName.concat(String(randomNumber())),
      ...props.projectProps,
    });
    // Create a new MongoDB Atlas Cluster and pass project ID
    new atlas.CfnCluster(this, 'cluster-'.concat(id),
      {
        apiKeys: props.apiKeys,
        name: props.clusterProps.name || clusterDefaults.clusterName.concat(String(randomNumber())),
        projectId: mProject.ref,
        clusterType: clusterDefaults.clusterType,
        ...props.clusterProps,
      },
    );
    // Create a new MongoDB Atlas Database User
    new user.CfnDatabaseUser(this, 'db-user-'.concat(id),
      {
        apiKeys: props.apiKeys,
        databaseName: props.dbUserProps?.databaseName || dbDefaults.dbName,
        projectId: mProject.ref,
        username: props.dbUserProps?.username || dbDefaults.username,
        roles: props.dbUserProps?.roles || dbDefaults.roles,
        password: props.dbUserProps?.password || dbDefaults.password,
        ...props.dbUserProps,
      });
    // Create a new MongoDB Atlas Project IP Access List
    new ipAccessList.CfnProjectIpAccessList(this, 'ip-access-list-'.concat(id),
      {
        apiKeys: props.apiKeys,
        projectId: mProject.ref,
        accessList: props.ipAccessListProps?.accessList || ipAccessDefaults.accessList,
        ...props.ipAccessListProps,
      },
    );

  }
}

function randomNumber() {
  const min = 10;
  const max = 9999999;
  return Math.floor(Math.random() * (max - min + 1) + min);
}
