
![atlas-cfn](https://gist.githubusercontent.com/jasonmimick/68f402c378ed364ea1684fda1a7ef5d2/raw/3094b7a2f77d4a8d3e0a8acf4a876844acb39685/atlas-cfn.png "atlas-cfn") MongoDB Atlas Cloud Formation Custom Resources
=====

This note is intended for early adopters and testers interested in trying out the brand new MongoDB Atlas custom resources for Amazon AWS Cloud Formation. 
Note: this software and the Cloud Formation custom resource SDK from AWS are both still Beta products.

# Getting Started

1. Clone our CFN resources from the source repository: 

```bash
git clone https://github.com/mongodb/mongodbatlas-cloudformation-resources
```

2. Setup the `aws` cli as usual, then you can use the helper script to push the resource into the AWS region of your choice.

```bash
cd mongodbatlas-cloudformation-resources
docker build -t mongodb/atlas-cfn-deploy .github/actions/atlas-cfn-deploy/Dockerfile

./util/atlas-cfn-deploy/atlas-cfn-deploy.py --region=us-east-1 all+
```

There is a helper tool to scrub all the "mongodb-atlas-*-role-stack" stacks: [/util/atlas-cfn-stack-cleaner.sh](/util/atlas-cfn-stack-cleaner.sh).

3. Create a secret with a MongoDB Atlas API Key for CFN to use:

```bash
APIKEY=$(cat <<APIKEY
{
  "AtlasMongoDBPublicKey": "<PASTE_ATLAS_PUBLIC_APIKEY_HERE>"
 ,"AtlasMongoDBPrivateKey": "<PASTE_PRIVATE_APIKEY_HERE>" 
 ,"AtlasMongoDBOrgID": "<PASTE_ATLAS_ORGANIZATION_ID_HERE>" 
 }
APIKEY
)
aws secretsmanager create-secret --name "mongodb/atlas/key" \
--secret-string "${APIKEY}" --region us-east-2
```

*Note* the name of the secret we used here is `mongodb/atlas/key` (this is the default in the template) but we recommend a more suitable name for your project. We will refactor the names of the "keys" in this json format too.

4. Select, edit, and deploy the sample CFN Template of your choice. We suggest the [MongoDB Cloud CFN Quickstart](/quickstart-mongodbatlas/templates/mongodbatlas-cfn-quickstart.template.json). This template will provision an Atlas Project, Cluster, Database User, and (optionally) a VPC Peering to the AWS VPC of your choice (BYO-VPC).

Here's an example. Create a file for your parameters, like this:

```json
 [
 {
    "ParameterKey": "MongoDBAtlasClusterName",
    "ParameterValue": "mongodb-atlas-cfn-rocks"
 },
 {
    "ParameterKey": "MongoDBAtlasUsername",
    "ParameterValue": "User123"
 },
 {
    "ParameterKey": "MongoDBAtlasPassword",
    "ParameterValue": ""
 },
 {
    "ParameterKey": "MongoDBAtlasAPIKeySecretName",
    "ParameterValue": "mongodb/atlas/key"
 },
 {
    "ParameterKey": "MongoDBAtlasInstanceSize",
    "ParameterValue": "M10"
 }
]
```

Then refer to that file and the template to lauch the stack:

``` bash
aws cloudformation create-stack --disable-rollback --stack-name=naughty-goldwasser --region us-east-2 --template-body=file://mongodbatlas-existingvpc.template.json --parameters=file:///tmp/tmp.T6a93x6n9u
```

Consult the [templates/create-stack.sh](template/create-stack.sh) helper for an example.

The [.github/workflows/launch-quickstart-mongodbatlas.yml](.github/workflows/launch-quickstart-mongodbatlas.yml) Github workflow is another useful example.

4. Check the Atlas UI, cloud.mongodb.com, for your new CloudFormation build Atlas stack.

5. Connect a simple Flask app from EC2.

:construction:

# Getting Help & Getting Involved

+ If you found an error, very very likely - please file an [Issue](https://github.com/mongodb/mongodbatlas-cloudformation-resources/issues/new)

+ Share your results to @MongoDBAtlasCFN

