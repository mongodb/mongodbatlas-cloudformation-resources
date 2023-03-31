# MongoDB::Atlas::ProjectInvitation

## Description
Returns, adds, edits, and removes project-invitations.

## Attributes & Parameters

Please consult the [Resource Docs](docs/README.md)

## Local Testing

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
cd ${repo_root}/cfn-resources/project-invitation
./test/project-invitation.create-sample-cfn-request.sh > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json 
cfn invoke DELETE test.request.json 
(Note: DELETE Invoke needs Id parameter from Create Output)
```

Both CREATE & DELETE tests must pass.

## Installation
```
TAGS=logging make
cfn submit --verbose --set-default
```

## Example and Usage 

Example AWS CloudFormation template is available here [example template.](../../examples/project-invitation/project-invitation.json)


```bash
#Configure you AWS Credentials to create Cloudformation Stack
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export AWS_REGION=""
export AWS_DEFAULT_REGION=""

#Command to deploy the sample Project-Invitation stack (Before this step "cfn submit" should have been executed successfully)
./examples/project-invitation/deploy.sh
```


For more information see: MongoDB Atlas API [Project Invitation Endpoint](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Projects/operation/returnAllProjectInvitations) Documentation.



