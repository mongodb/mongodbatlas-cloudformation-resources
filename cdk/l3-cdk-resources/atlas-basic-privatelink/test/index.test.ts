import { App } from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { ClusterL3 } from '../src/index';

test('Snapshot', () => {
  const app = new App();
  const stack = new ClusterL3(app, 'test');

  const template = Template.fromStack(stack);
  expect(template.toJSON()).toMatchSnapshot();
});