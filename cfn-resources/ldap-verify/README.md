# MongoDB::Atlas::LDAPVerify

Requests a verification of an LDAP configuration over TLS for an Atlas project.
Pass the requestId in the response object to the Verify |ldap| Configuration endpoint to get the status of a verification request. Atlas retains only the most recent request for each project.

If the cloud formation stack gets deleted, the current resource, executes a validation again, in other to replace the original validation

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

## Example
This file can be used as a CloudFormation template to create a new Maintenance Window resource

[LDAPVerify.json](../../examples/LDAPVerify/LDAPVerify.json)

## Local Testing

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/project
./test/project.create-sample-cfn-request.sh YourProjectName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
```

Both CREATE & DELETE tests must pass.

## Installation
TAGS=logging make
cfn submit --verbose --set-default

## Usage

The [launch-x-quickstart.sh](../../quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh) script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other neccessary parameters.

You can use the project.sample-template.yaml to create a stack using the resource.
Similar to the local testing described above you can follow the logs for the deployed
lambda function which handles the request for the Resource Type.

In one shell session:
```
aws logs tail mongodb-atlas-project-logs --follow
```

And then you can create the stack with a helper script it insert the apikeys for you:


```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
${repo_root}/quickstart-mongodb-atlas/scripts/launch-x-quickstart.sh ${repo_root}/cfn-resources/project/test/project.sample-template.yaml SampleProject1 ParameterKey=OrgId,ParameterValue=${ATLAS_ORG_ID}
```

## For More Information
See the MongoDB Atlas API [LDAP Verification Docs](https://www.mongodb.com/docs/atlas/reference/api/ldaps-configuration-request-verification/) Documentation.
