const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const custom_dns_configuration_cluster_aws = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',  
  publishToGo: {
   moduleName: 'github.com/mongoDB',
   packageName : 'custom-dns-configuration-cluster-aws'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-custom-dns-configuration-cluster-aws',
    module: 'mongodb_cdk_mongodb_custom_dns_configuration_cluster_aws',
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/custom-dns-configuration-cluster-aws',
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  keywords: ['cdk',
    'awscdk',
    'aws-cdk',
    'cloudformation',
    'cfn',
    'extensions',
    'constructs',
    'cfn-resources',
    'cloudformation-registry',
    'l1',
    'mongodb',
    'atlas',
    'custom-dns-configuration-cluster-aws'],
  description: 'Retrieves or creates custom-dns-configuration-cluster-awss in any given Atlas organization',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

custom_dns_configuration_cluster_aws.eslint?.addOverride({
  files: ['src/**'],
  rules: {
    'max-len': [
      'error', {
        code: 180,
      },
    ],
  },
});

custom_dns_configuration_cluster_aws.synth();
