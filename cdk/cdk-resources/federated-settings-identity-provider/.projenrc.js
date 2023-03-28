const { awscdk } = require('projen')
const { ReleaseTrigger } = require('projen/lib/release')
const federated_settings_identity_provider = new awscdk.AwsCdkConstructLibrary({
  author: 'MongoDBAtlas',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDBAtlas',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'master',
  publishToNuget: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasFederatedSettingsIdentityProvider',
    packageId: 'MongoDBCdk.MongoDBAtlasFederatedSettingsIdentityProvider'
  },
  publishToMaven: {
    javaPackage:
      'com.github.mongodb.cdk.mongodb_atlas_federated_settings_identity_provider',
    mavenArtifactId: 'mongodb-atlas-federated-settings-identity-provider',
    mavenGroupId: 'com.github.mongodb.cdk'
  },
  publishToGo: {
    moduleName: 'github.com/mongoDB',
    packageName: 'federated-settings-identity-provider'
  },
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-federated-settings-identity-provider',
    module: 'mongodb_cdk_mongodb_federated_settings_identity_provider'
  },
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodbatlas-awscdk/federated-settings-identity-provider',
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
    'federated-settings-identity-provider'
  ],
  description:
    'Retrieves or creates federated-settings-identity-providers in any given Atlas organization'
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
})
federated_settings_identity_provider.eslint?.addOverride({
  files: ['src/**'],
  rules: {
    'max-len': [
      'error',
      {
        code: 180
      }
    ]
  }
})

federated_settings_identity_provider.synth()
