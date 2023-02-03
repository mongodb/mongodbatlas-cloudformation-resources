import * as atlas from '@mongodbatlas-awscdk/cluster';
import * as user from '@mongodbatlas-awscdk/database-user';
import * as ipAccessList from '@mongodbatlas-awscdk/project-ip-access-list';
import * as cdk from 'aws-cdk-lib';
// import * as cloudformation from 'aws-cdk-lib/aws-cloudformation';
import { Construct } from 'constructs';
import * as defaults from './defaults.json';


export class AtlasBasic extends cdk.Stack {
  constructor(scope: Construct, id: string, props: AtlasBasicProps) {
    super(scope, id, props);
    // Create a new MongoDB Atlas Project
    // const mProject = new project.CfnProject(this, 'project-'.concat(id), {
    // //    name: defaults.projectName,
    //     apiKeys: props.apiKeys,
    //     ...props.projectProps
    // });
    // Create a new MongoDB Atlas Cluster and pass project ID
    new atlas.CfnCluster(this, 'cluster-'.concat(id),
      {
        //    name: defaults.clusterName,
        //    projectId: mProject.ref,
        clusterType: defaults.clusterType,
        replicationSpecs: defaults.replicationSpecs,
        //    apiKeys: props.apiKeys,
        ...props.clusterProps,
      },
    );
    // Create a new MongoDB Atlas Database User
    new user.CfnDatabaseUser(this, 'cfn-db-user-'.concat(id),
      {
        password: defaults.password,
        ...props.dbUserProps,
      });
    // Create a new MongoDB Atlas Project IP Access List
    new ipAccessList.CfnProjectIpAccessList(this, 'ip-access-list-'.concat(id),
      {
        ...props.ipAccessListProps,
      },
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

// export interface ProjectProps extends Omit<project.CfnProjectProps, 'name'> {
//     readonly  name? : string
// }
//
// export interface ClusterProps extends Omit<atlas.CfnClusterProps, 'projectId' | 'name' | 'apiKeys'> {
//     readonly name? : string
//     readonly projectId? : string
//     readonly apiKeys?: atlas.ApiKeyDefinition;
// }
//
// export interface DatabaseUserProps extends Omit<user.CfnDatabaseUserProps, 'databaseName' | 'roles' | 'username' | 'projectId'> {
//     readonly databaseName? : string
//     readonly roles? : user.RoleDefinition[]
//     readonly username? : string
//     readonly projectId? : string
// }
//
// export interface IpAccessListProps extends Omit<ipAccessList.CfnProjectIpAccessListProps, 'accessList' | 'apiKeys' | 'projectId'> {
//     readonly accessList?: ipAccessList.AccessListDefinition[];
//     readonly apiKeys?: ipAccessList.ApiKeyDefinition;
//     readonly projectId?: string;
// }

export interface AtlasBasicProps extends cdk.StackProps{
  readonly projectId : string;
  readonly apiKeys : atlas.ApiKeyDefinition;
  // readonly projectProps: project.CfnProjectProps;
  readonly clusterProps: atlas.CfnClusterProps;
  readonly dbUserProps: user.CfnDatabaseUserProps;
  readonly ipAccessListProps: ipAccessList.CfnProjectIpAccessListProps;
}

