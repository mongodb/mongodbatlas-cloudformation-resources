import { App } from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import {AtlasBasic, AtlasBasicProps} from '../src/index';

test('Snapshot', () => {
    const app = new App();
    let apiKeys: {};
    const stack = new AtlasBasic(app, 'test', {
        apiKeys: {},
        clusterProps:{},
        dbUserProps: {},
        ipAccessListProps: {},
        projectProps: {orgId:''}
    });

});