const { awscdk } = require('projen');
const { ReleaseTrigger } = require('projen/lib/release');
const custom_dns_configuration_cluster_aws = new awscdk.AwsCdkConstructLibrary
({
  author: 'MongoDB',
  authorAddress: 'https://mongodb.com',
  authorName: 'MongoDB',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'INTMDB-548',
  majorVersion: 1,
  releaseToNpm: true,
  releaseTrigger: ReleaseTrigger.manual(),
  docgen: true,
  name: '@mongodb-cdk/atlas-custom-dns-configuration-cluster-aws',
  repositoryUrl: 'https://github.com/mongodb/mongodbatlas-cloudformation-resources.git',
  sampleCode: false,
  publishToPypi: {
    distName: 'mongodb-cdk-mongodb-atlas-custom-dns-configuration-cluster-aws',
    module: 'mongodb_cdk_mongodb_atlas_custom-dns-configuration-cluster-aws',
  },
  dotnet: {
    dotNetNamespace: 'MongoDBCdk.MongoDBAtlasCustomDnsConfigurationClusterAws',
    packageId: 'MongoDBCdk.MongoDBAtlasCustomDnsConfigurationClusterAws',
  },
  publishToMaven: {
    javaPackage: 'com.github.mongodb.cdk.mongodb_atlas_custom-dns-configuration-cluster-aws',
    mavenArtifactId: 'mongodb-atlas-custom-dns-configuration-cluster-aws',
    mavenGroupId: 'com.github.mongodb.cdk',
  },
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

custom_dns_configuration_cluster_aws.synth();
