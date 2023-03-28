const { awscdk } = require('projen')
const { ReleaseTrigger } = require('projen/lib/release')
const cloud_provider_access = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasCloudProviderAccess',
    packageId: 'MongoDBCdk.MongoDBAtlasCloudProviderAccess'
  },
  publishToMaven: {
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_cloud_provider_access',
    mavenArtifactId: 'mongodb-atlas-cloud-provider-access',
    mavenGroupId: 'com.github.mongodb.cdk'
  },
  publishToGo: {
    moduleName: 'github.com/mongoDB',
    packageName: 'cloud-provider-access'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-cloud-provider-access',
    module: 'mongodb_cdk_mongodb_cloud_provider_access'
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/atlas-cloud-provider-access',
  repositoryUrl:
    'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  keywords: [
    'cdk',
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
    'cloud-provider-access'
  ],
  description:
    'Retrieves or creates cloud-provider-accesss in any given Atlas organization'
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
})

cloud_provider_access.synth()
