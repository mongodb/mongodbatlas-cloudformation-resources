import * as cdk from 'aws-cdk-lib';
import {ApiKeyDefinition, AtlasBasicProps} from "@mongodbatlas-awscdk/atlas-basic";
import {AtlasBasicPrivateEndpoint, PrivateEndpointProps} from "./index";

const app = new cdk.App();

const stack = new cdk.Stack(app, 'atlas-basic-default', {
    env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});

const apiKeys: ApiKeyDefinition = {
    privateKey: stack.node.tryGetContext('MONGODB_ATLAS_PRIVATE_KEY') || process.env.MONGODB_ATLAS_PRIVATE_KEY,
    publicKey: stack.node.tryGetContext('MONGODB_ATLAS_PUBLIC_KEY') || process.env.MONGODB_ATLAS_PUBLIC_KEY,
};

const orgId = stack.node.tryGetContext('MONGODB_ATLAS_ORG_ID') || process.env.MONGODB_ATLAS_ORG_ID;
const vpcId = stack.node.tryGetContext('AWS_VPC_ID') || process.env.AWS_VPC_ID;
const subnetId = stack.node.tryGetContext('AWS_SUBNET_ID') || process.env.AWS_SUBNET_ID;
const awsRegion = stack.node.tryGetContext('AWS_REGION') || process.env.AWS_REGION;

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

const privateEndpointProps : PrivateEndpointProps = {
    privateEndpoints: [{
        vpcId: vpcId,
        subnetIds: [subnetId]
    }],
}

const props   = {
    apiKeys: apiKeys,
    atlasBasicProps: atlasBasicProps,
    privateEndpointProps: privateEndpointProps,
    region: awsRegion
}

new AtlasBasicPrivateEndpoint(stack,'private-endpoint', props)




