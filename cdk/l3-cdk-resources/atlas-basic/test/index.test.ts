import { App } from 'aws-cdk-lib';


test('Synth', () => {
  const app = new App();

  app.synth();
});