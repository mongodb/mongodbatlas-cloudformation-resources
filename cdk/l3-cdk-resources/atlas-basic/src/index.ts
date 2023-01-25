import * as cdk from 'aws-cdk-lib';
import * as atlas from '@mongodbatlas-awscdk/cluster';
import * as project from '@mongodbatlas-awscdk/project';
import * as user from '@mongodbatlas-awscdk/database-user';
import * as ipAccessList from '@mongodbatlas-awscdk/project-ip-access-list';

class ClusterL3 extends cdk.Stack {
    constructor(scope: cdk.App, id: string, props?: cdk.StackProps) {
        super(scope, id, props);

        // Create a new MongoDB Atlas Project
        const myProject = new project.CfnProject(this, 'MyProject', {
            name: 'my-project',
            orgId: ""
        });

        // Create a new MongoDB Atlas Cluster and pass project ID
        const myCluster = new atlas.CfnCluster(this, 'MyCluster', {
            apiKeys: undefined,
            name: "",
            projectId: myProject.ref,
        });

        // Create a new MongoDB Atlas Database User
        new user.CfnDatabaseUser(this, 'MyUser', {
            databaseName: "",
            projectId: myProject.ref,
            roles: [],
            username: ""
        });

        // Create a new MongoDB Atlas Project IP Access List
        new ipAccessList.CfnProjectIpAccessList(this, 'MyIpAccessList', {
            accessList: [],
            apiKeys: undefined,
            projectId: myProject.ref,
        });
    }
}
