# quickstart-mongodb-atlas 

MongoDB Atlas in the AWS Cloud

This Quick Start deploys a MongoDB Atlas Resource Provider for CloudFormation.
The resource provider allows you to create and manage complete Enterprise-ready MongoDB Atlas
deployment directly through CloudFormation. 

![Quick Start architecture for MongoDB Atlas on AWS](docs/images/simple-quickstart-arch.png)

Includes support for:
* MongoDB Atlas Projects
* MongoDB Atlas Clusters
* MongoDB Atlas Database Users via AWS IAM Integration
* VPC Peering
* (Private Link)


For architectural details, best practices, step-by-step instructions, and customization options, see the deployment guide.

To post feedback, submit feature ideas, or report bugs, use the [Issues](/issues) section of this GitHub repo. If you'd like to submit code for this Quick Start, please review the AWS Quick Start Contributor's Kit.

# Getting Started

## Setup AWS & API Keys

If needed, install the awscli and mongocli.

```bash
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "/tmp/awscliv2.zip"
unzip /tmp/awscliv2.zip
sudo /tmp/aws/install
MONGOCLI_VERSION="1.7.0"
curl -L "https://github.com/mongodb/mongocli/releases/download/${MONGOCLI_VERSION}/mongocli_${MONGOCLI_VERSION}_linux_x86_64.tar.gz" -o "/tmp/mongocli_${MONGOCLI_VERSION}_linux_x86_64.tar.gz"
tar xzvf "/tmp/mongocli_${MONGOCLI_VERSION}_linux_x86_64.tar.gz" --directory /tmp
cp "/tmp/mongocli_${MONGOCLI_VERSION}_linux_x86_64/mongocli" "~/.local/bin"
~/.local/bin/mongocli --version
```

Make sure to configure each tool properly.

```bash
aws configure
mongocli config
```

+ Run this helper to setup environment variables for your 
MongoDB Atlas API keys (read from mongocli config)

```bash
source <(./scripts/export-mongocli-config.py)
```

## Deploy the MongoDB Atlas Resource Types into your AWS region.

This quickstart is powered by a set of official MongoDB AWS CloudFormation Resource Types
which connect your AWS CloudFormation control plane directly into
the MongoDB Cloud. Right now, these resources need to be registered in each AWS region prior to use. 

Run this command to install the MongoDB Atlas
Resource Types into the `AWS_REGION` of your choice before running the quickstart.

```bash
TODO
```

# Launch the quickstart stack

The `quickstart-mongodb-atlas.template.yaml` template will
provision a complete you MongoDB Atlas Deployment for you.

This includes the follow resources:
* [MongoDB::Atlas::Project](/cfn-resources/project)
* [MongoDB::Atlas::ProjectIpAccessList](/cfn-resources/project-ip-access-list) 
* [MongoDB::Atlas::Cluster](/cfn-resources/cluster)
* [MongoDB::Atlas::DatabaseUser](/cfn-resources/database-user) 
    * Includes AWS IAM Role Integration 

```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/quickstart-mongodb-atlas/templates/quickstart-mongodb-atlas.template.yaml MongoDB-Atlas-Quickstart ParameterKey=OrgId,ParameterValue=${ATLAS_ORG_ID} 
```

The stack will take ~7-10 minutes to provision. When complete you can find the `mongodb+srv` connection information in the stack outputs.

```
aws cloudformation describe-stacks --stack-name ${STACK_NAME} | jq -r '.Stacks[0]|.Outputs'
```

Currently there are 3 outputs:
TODO: *NEEDS UPDATE*
```
[
  {
    "OutputKey": "AtlasDatabaseUser",
    "OutputValue": "org:5ea0477597999053a5f9cbec,project:5f8723ae20f10f128d3d6a07",
    "Description": "AWS IAM ARN for database user"
  },
  {
    "OutputKey": "SrvHost",
    "OutputValue": "mongodb+srv://cookies-99-5x.cqpb3.mongodb.net",
    "Description": "Hostname for mongodb+srv:// connection string",
    "ExportName": "cookies-99-5x-standardSrv"
  },
  {
    "OutputKey": "AtlasDeployment",
    "OutputValue": "org:5ea0477597999053a5f9cbec,project:5f8723ae20f10f128d3d6a07",
    "Description": "Info on your Atlas deployment"
  }
]
```

## Connect to your database

After the cluster provisions, you can connect with the `mongo` shell or MongoDB Compass.

Fetch the new cluster `mongodb+srv://` host info:

```bash
STACK_NAME="mongodb-atlas-quickstart"
MDB=$(aws cloudformation list-exports |\
 jq -r --arg stackname "${STACK_NAME}" \
 '.Exports[] | select(.Name==$stackname+"-standardSrv") | .Value')
echo "New ${STACK_NAME} database url: ${MDB}"
```
Use this url along with your `aws` cli credentials to seamlessly and securly connect to your new MongoDB Atlas database:

```bash
STACK_ROLE=$(aws cloudformation describe-stack-resources --stack-name "${STACK_NAME}" --logical-resource-id AtlasIAMRole)
ROLE=$(aws iam get-role --role-name $( echo "${STACK_ROLE}" | jq -r '.StackResources[] | .PhysicalResourceId'))
ROLE_ARN=$(echo "${ROLE}" | jq -r '.Role.Arn')
ROLE_CREDS=$(aws sts assume-role --role-session-name test --role-arn ${ROLE_ARN})
mongo "${MDB}/${STACK_NAME}?authSource=%24external&authMechanism=MONGODB-AWS" \
    --username $(echo "${ROLE_CREDS}" | jq -r '.Credentials.AccessKeyId') \
    --password $(echo "${ROLE_CREDS}" | jq -r '.Credentials.SecretAccessKey') \
    --awsIamSessionToken $(echo "${ROLE_CREDS}" | jq -r '.Credentials.SessionToken')
```

see [scripts/aws-iam-mongo-shell.sh](scripts/aws-iam-mongo-shell.sh).


# Launch the quickstart stack with Peering

The `quickstart-mongodb-atlas-peering.template.yaml` stack will 
provision a complete you MongoDB Atlas Deployment with VPC Peering for a given AWS VPC enabled. 

This includes the follow resources:
* [MongoDB::Atlas::Project](/cfn-resources/project)
* [MongoDB::Atlas::ProjectIpAccessList](/cfn-resources/project-ip-access-list) 
* [MongoDB::Atlas::Cluster](/cfn-resources/cluster)
* [MongoDB::Atlas::DatabaseUser](/cfn-resources/database-user) 
    * Includes AWS IAM Role Integration 
* [MongoDB::Atlas::NetworkPeering](/cfn-resources/network-peering)
    * With automatic network container management


