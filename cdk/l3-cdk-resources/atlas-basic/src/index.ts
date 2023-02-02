import * as cdk from 'aws-cdk-lib';
import * as atlas from '@mongodbatlas-awscdk/cluster';
import * as project from '@mongodbatlas-awscdk/project';
import * as user from '@mongodbatlas-awscdk/database-user';
import * as ipAccessList from '@mongodbatlas-awscdk/project-ip-access-list';
import * as cloudformation from 'aws-cdk-lib/aws-cloudformation';
import * as defaults from "./defaults.json";
import * as constructs from 'constructs';


export class AtlasBasic extends cdk.Stack {
    constructor(scope: constructs.Construct, id: string, props: AtlasBasicProps) {
        super(scope, id, props);
        // Create a new MongoDB Atlas Project
        const mProject = new project.CfnProject(this, 'project-'.concat(id), {
            name: defaults.projectName,
            apiKeys: props.apiKeys,
            ...props.projectProps
        });
        // Create a new MongoDB Atlas Cluster and pass project ID
        const mCluster = new atlas.CfnCluster(this, 'cluster-'.concat(id),
            {
                name: defaults.clusterName,
                projectId: mProject.ref,
                clusterType: defaults.clusterType,
                replicationSpecs: defaults.replicationSpecs,
                apiKeys: props.apiKeys,
                ...props.clusterProps,
            }
        );
        // Create a new MongoDB Atlas Database User
        const dbUser = new user.CfnDatabaseUser(this, 'cfn-db-user-'.concat(id),
            {
                apiKeys: props.apiKeys,
                databaseName: defaults.databaseName,
                projectId: mProject.ref,
                roles: defaults.roles,
                username: defaults.username,
                password: defaults.password,
                ...props.dbUserProps
            });
        // Create a new MongoDB Atlas Project IP Access List
        const ipAccess = new ipAccessList.CfnProjectIpAccessList(this, 'ip-access-list-'.concat(id),
            {
                accessList: defaults.accessList,
                projectId: mProject.ref,
                apiKeys: props.apiKeys,
                ...props.ipAccessListProps
            }
        );
        new cloudformation.CfnTypeActivation(this, 'ProjectActivation', {
            type: mProject.cfnResourceType,
            executionRoleArn: ""
        });

        new cloudformation.CfnTypeActivation(this, 'ClusterActivation', {
            type: mCluster.cfnResourceType,
            executionRoleArn: ""
        });

        new cloudformation.CfnTypeActivation(this, 'DatabaseUserActivation', {
            type: dbUser.cfnResourceType,
            executionRoleArn: "",
        });

        new cloudformation.CfnTypeActivation(this, 'ProjectIpAccessListActivation', {
            type: ipAccess.cfnResourceType,
            executionRoleArn: "",
        });

    }
}

export interface ProjectProps extends Omit<project.CfnProjectProps, 'name'> {
    name? : string
}

export interface ClusterProps extends Omit<atlas.CfnClusterProps, 'projectId' | 'name' | 'apiKeys'> {
    name? : string
    projectId? : string
    apiKeys?: atlas.ApiKeyDefinition;
}

export interface DatabaseUserProps extends Omit<user.CfnDatabaseUserProps, 'databaseName' | 'roles' | 'username' | 'projectId'> {
    databaseName? : string
    roles? : user.RoleDefinition[]
    username? : string
    projectId? : string
}

export interface IpAccessListProps extends Omit<ipAccessList.CfnProjectIpAccessListProps, 'accessList' | 'apiKeys' | 'projectId'> {
    accessList?: ipAccessList.AccessListDefinition[];
    apiKeys?: ipAccessList.ApiKeyDefinition;
    projectId?: string;
}

export interface AtlasBasicProps extends cdk.StackProps{
    apiKeys : atlas.ApiKeyDefinition
    projectProps: ProjectProps;
    clusterProps?: ClusterProps;
    dbUserProps?:  DatabaseUserProps;
    ipAccessListProps?: IpAccessListProps;
}