## Mongodb::Atlas::OrgInvitation

### Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- Org invitation L1 CDK constructor


### Resources (and parameters for local tests) needed to manually QA:
The Atlas organization must be manually provided.
- Atlas Organization (ATLAS_ORG_ID)
- Team Id (created by cfn-test-create-inputs.sh)
- username (created by cfn-test-create-inputs.sh)



## Manual QA:

### Prerequisite steps:
1. Create an Atlas Organization if you don’t already have one.
2. Export environment variable ATLAS_ORG_ID with the value as your Organization Id. You can get this Id from your Atlas UI 
by clicking on your organization and note the Organization Id from the URL (which should look like this https://cloud.mongodb.com/v2#/org/<ATLAS_ORG_ID>/projects)


### Steps to test:
1. Ensure prerequisites above for this resource and general [prerequisites](../../../TESTING.md#prerequisites) are complete.
2. Follow [general steps](../../../TESTING.md#steps) to test a CFN resource.

### Success criteria when testing the resource
1. Under Access Manager for your Atlas Organization, you should see a new entry under “Users” tab with status “PENDING INVITE”  

![image](https://user-images.githubusercontent.com/122359335/227275914-4af66737-fa72-49f8-8713-9d298606bc4f.png)

2. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api/organization-create-one-invitation/#invite-one-user-to-an-service-organization)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/invitations/#invitations-to-organizations-and-projects)

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

Both CREATE and DELETE tests must pass.
