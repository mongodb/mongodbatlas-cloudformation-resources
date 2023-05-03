# MongoDB::Atlas::LDAPConfiguration

An LDAP configuration defines settings for Atlas to connect to your LDAP server over TLS for user authentication and authorization.
Your LDAP server must be visible to the internet or connected to your Atlas cluster with VPC Peering.
In addition, your LDAP server must use TLS.

## Attributes and Parameters

See the [Resource Docs](./docs/README.md).

## Example
You can use [`LDAPConfiguration.json`](../../examples/ldap-configuration/LDAPConfiguration.json)
as a CloudFormation template to create a new LDAP Configuration resource.

## Installation

```
TAGS=logging make
cfn submit --verbose --set-default
```

## Usage

The [launch-x-quickstart.sh](https://github.com/aws-quickstart/quickstart-mongodb-atlas/blob/master/scripts/launch-x-quickstart.sh)  script
can be used to safely inject your MongoDB Cloud ApiKey environment variables into an example
CloudFormation stack template along with the other neccessary parameters.

You can use the `project.sample-template.yaml` to create a stack using the resource.
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

See the MongoDB Atlas API Endpoint [LDAP-Configuration](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/LDAP-Configuration) documentation.
