import * as cdk from 'aws-cdk-lib';
import {MicrosoftTeamsIntegration, MicrosoftTeamsIntegrationProps} from './microsoftTeamsIntegration';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'third-party-integration', {
    env: { region: process.env.CDK_DEFAULT_REGION, account: process.env.CDK_DEFAULT_ACCOUNT },
});


const projectId = stack.node.tryGetContext('MONGODB_ATLAS_PROJECT_ID') || process.env.MONGODB_ATLAS_PROJECT_ID;
const webhookURL = stack.node.tryGetContext('TEAMS_WEBHOOK_URL') || process.env.TEAMS_WEBHOOK_URL;

const props : MicrosoftTeamsIntegrationProps = {
    microsoftTeamsWebhookUrl: webhookURL,
    projectId: projectId
}

new MicrosoftTeamsIntegration(stack, 'teams-integration',  props);
