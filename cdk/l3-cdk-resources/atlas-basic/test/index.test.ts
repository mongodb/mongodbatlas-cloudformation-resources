import { App } from 'aws-cdk-lib';
import {AtlasBasic} from '../src/index';

test('Synth', () => {
    const app = new App();
    new AtlasBasic(app, 'test-app', {
        apiKeys: {privateKey: 'e6c4bac8-8312-4add-bfca-ee750d4798e4', publicKey: 'hynkfzcw'},
        projectProps: {orgId:'5fe4ea50d1a2b617175ee3d4'}
    });
    app.synth();
});