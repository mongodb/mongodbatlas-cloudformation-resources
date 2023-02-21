const { awscdk, javascript, JsonFile } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.64.0',
  defaultReleaseBranch: 'master',
  description: 'Creates a project, cluster, dbuser, ipaccess list and set encryption-at-rest in MongoDB Atlas',
  docgen: true,
  keywords: ['cdk', 
  'awscdk', 
  'aws-cdk', 
  'cloudformation', 
  'cfn', 
  'extensions', 
  'constructs', 
  'cfn-resources', 
  'cloudformation-registry', 
  'l3', 
  'mongodb', 
  'atlas', 
  'encryption-at-rest-express',
  'encryption-at-rest'],
  majorVersion: 1,
  name: '@mongodbatlas-awscdk/encryption-at-rest-express',
  npmAccess: javascript.NpmAccess.PUBLIC,
  releaseToNpm: true,
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  peerDeps: [
    '@mongodbatlas-awscdk/atlas-basic',
    '@mongodbatlas-awscdk/atlas-encryption-at-rest',
  ],
  devDeps: ['aws-cdk', 'ts-node'],
  stability: 'experimental',
});

new JsonFile(project, 'cdk.json', {
  obj: {
    app: 'npx ts-node --prefer-ts-exts src/integ.default.ts',
  },
});

const common_exclude = ['cdk.out', 'cdk.context.json', 'yarn-error.log'];
project.npmignore.exclude(...common_exclude);
project.gitignore.exclude(...common_exclude);


project.synth();