import * as atlas from '@mongodbatlas-awscdk/cluster';
import * as user from '@mongodbatlas-awscdk/database-user';
import * as project from '@mongodbatlas-awscdk/project';
import * as ipAccessList from '@mongodbatlas-awscdk/project-ip-access-list';
import * as cdk from 'aws-cdk-lib';
// import * as cloudformation from 'aws-cdk-lib/aws-cloudformation';
import { Construct } from 'constructs';



const defaults = {
    projectName: "cdk-project",
    clusterName: "cdk-cluster",
    dbName: "admin",
    username: "cdk-user",
    password: "cdk-pwd",
    roles: [{
        "roleName": "atlasAdmin",
        "databaseName": "admin"
    }],
    accessList: [
        {
            "ipAddress": "0.0.0.0/1",
            "comment": "Testing open all ips"
        }
    ],
    clusterType: "REPLICASET",
};

export interface ProjectProps {
    readonly name?: string;
    readonly orgId: string;
    readonly projectOwnerId?: string;
    readonly withDefaultAlertsSettings?: boolean;
    readonly clusterCount?: number;
    readonly projectSettings?: project.ProjectSettings;
    readonly projectTeams?: project.ProjectTeam[];
    readonly projectApiKeys?: project.ProjectApiKey[];
}
export interface ClusterProps {
    readonly advancedSettings?: atlas.ProcessArgs;
    readonly apiKeys?: atlas.ApiKeyDefinition;
    readonly backupEnabled?: boolean;
    readonly biConnector?: atlas.CfnClusterPropsBiConnector;
    readonly clusterType?: string;
    readonly connectionStrings?: atlas.ConnectionStrings;
    readonly diskSizeGb?: number;
    readonly encryptionAtRestProvider?: atlas.CfnClusterPropsEncryptionAtRestProvider;
    readonly projectId?: string;
    readonly labels?: atlas.CfnClusterPropsLabels[];
    readonly mongoDbMajorVersion?: string;
    readonly name?: string;
    readonly paused?: boolean;
    readonly pitEnabled?: boolean;
    readonly replicationSpecs?: atlas.AdvancedReplicationSpec[];
    readonly rootCertType?: string;
    readonly versionReleaseSystem?: string;
    readonly terminationProtectionEnabled?: boolean;
}
export interface DatabaseUserProps {
    readonly deleteAfterDate?: string;
    readonly awsiamType?:user.CfnDatabaseUserPropsAwsiamType;
    readonly databaseName?: string;
    readonly labels?: user.LabelDefinition[];
    readonly ldapAuthType?: user.CfnDatabaseUserPropsLdapAuthType;
    readonly x509Type?: user.CfnDatabaseUserPropsX509Type;
    readonly password?: string;
    readonly projectId?: string;
    readonly roles?: user.RoleDefinition[];
    readonly scopes?: user.ScopeDefinition[];
    readonly username?: string;
}
export interface IpAccessListProps {
    readonly accessList: ipAccessList.AccessListDefinition[];
    readonly projectId?: string;
    readonly totalCount?: number;
    readonly listOptions?: ipAccessList.ListOptions;
}
export interface ApiKeyDefinition {
    readonly privateKey?: string;
    readonly publicKey?: string;

}

export interface AtlasBasicProps extends cdk.StackProps{
    readonly apiKeys : ApiKeyDefinition;
    readonly projectProps: ProjectProps;
    readonly clusterProps: ClusterProps;
    readonly dbUserProps?: DatabaseUserProps;
    readonly ipAccessListProps?: IpAccessListProps;
}

export class AtlasBasic extends cdk.Stack {
  constructor(scope: Construct, id: string, props: AtlasBasicProps) {
    super(scope, id, props);
    //Create a new MongoDB Atlas Project
  const mProject = new project.CfnProject(this, 'project-'.concat(id), {
      apiKeys: props.apiKeys,
      name: props.projectProps.name || defaults.projectName,
      ...props.projectProps
  });
    // Create a new MongoDB Atlas Cluster and pass project ID
    new atlas.CfnCluster(this, 'cluster-'.concat(id),
      {
          apiKeys: props.apiKeys,
          name: props.clusterProps.name || defaults.clusterName,
          projectId: mProject.ref,
          clusterType: defaults.clusterType,
          ...props.clusterProps,
      }
    );
    // Create a new MongoDB Atlas Database User
    new user.CfnDatabaseUser(this, 'cfn-db-user-'.concat(id),
      {
          apiKeys: props.apiKeys,
          databaseName: props.dbUserProps?.databaseName || defaults.dbName,
          projectId: mProject.ref,
          username: props.dbUserProps?.username || defaults.username,
          roles: props.dbUserProps?.roles || defaults.roles,
          password: props.dbUserProps?.password || defaults.password,
          ...props.dbUserProps,
      });
    // Create a new MongoDB Atlas Project IP Access List
    new ipAccessList.CfnProjectIpAccessList(this, 'ip-access-list-'.concat(id),
        {
            apiKeys: props.apiKeys,
            projectId: mProject.ref,
            accessList: props.ipAccessListProps?.accessList || defaults.accessList,
            ...props.ipAccessListProps,
        }
    );

    //
    // new cloudformation.CfnTypeActivation(this, 'ProjectActivation', {
    //     type: mProject.cfnResourceType,
    //     executionRoleArn: ""
    // });
    //
    // new cloudformation.CfnTypeActivation(this, 'ClusterActivation', {
    //     type: mCluster.cfnResourceType,
    //     executionRoleArn: ""
    // });
    //
    // new cloudformation.CfnTypeActivation(this, 'DatabaseUserActivation', {
    //     type: dbUser.cfnResourceType,
    //     executionRoleArn: "",
    // });
    //
    // new cloudformation.CfnTypeActivation(this, 'ProjectIpAccessListActivation', {
    //     type: ipAccess.cfnResourceType,
    //     executionRoleArn: "",
    // });

  }
}

