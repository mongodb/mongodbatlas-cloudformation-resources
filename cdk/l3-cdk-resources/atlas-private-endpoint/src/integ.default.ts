import * as cdk from 'aws-cdk-lib';
import {ApiKeyDefinition, AtlasBasicProps} from "@mongodbatlas-awscdk/atlas-basic";
import {AtlasBasicPrivateEndpoint, AtlasPrivateEndpointProps} from "./index";
import {CfnPrivateEndpointProps} from "@mongodbatlas-awscdk/private-endpoint";


const app = new cdk.App();

const stack = new cdk.Stack(app, 'atlas-basic-default', {
    env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});

const apiKeys: ApiKeyDefinition = {
    privateKey: stack.node.tryGetContext('MONGODB_ATLAS_PRIVATE_KEY') || process.env.MONGODB_ATLAS_PRIVATE_KEY,
    publicKey: stack.node.tryGetContext('MONGODB_ATLAS_PUBLIC_KEY') || process.env.MONGODB_ATLAS_PUBLIC_KEY,
};

const orgId = stack.node.tryGetContext('MONGODB_ATLAS_ORG_ID') || process.env.MONGODB_ATLAS_ORG_ID;

const replicationSpecs = [
    {
        numShards: 1,
        advancedRegionConfigs: [
            {
                analyticsSpecs: {
                    ebsVolumeType: 'STANDARD',
                    instanceSize: 'M10',
                    nodeCount: 1,
                },
                electableSpecs: {
                    ebsVolumeType: 'STANDARD',
                    instanceSize: 'M10',
                    nodeCount: 3,
                },
                priority: 7,
                regionName: 'US_EAST_1',
            },
        ],
    },
];

const atlasBasicProps : AtlasBasicProps = {
    apiKeys: apiKeys,
    clusterProps: {
        replicationSpecs: replicationSpecs,
    },
    projectProps: {
        orgId: orgId,
    },
}

const privateEndpointProps : CfnPrivateEndpointProps = {
    groupId: '',
    apiKeys: apiKeys,
    privateEndpoints: [{
        vpcId: '',
        subnetIds: ['']
    }],
    region: 'us-east-1'
}

const props   = {
    apiKeys: apiKeys,
    atlasBasicProps: atlasBasicProps,
    privateEndpointProps: privateEndpointProps,
    groupId: '',
    region: 'us-east-1'
}

new AtlasBasicPrivateEndpoint(stack,'private-endpoint', props)


