# Cloud backup snapshot export bucket


## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Cloud backup snapshot export bucket L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
- Amazon S3 bucket
- IAMRoleId

All resources are created as part of `cfn-testing-helper.sh`


## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
-  The endpoint `api/atlas/v1.0/groups/<ProjectId>/backup/exportBuckets` should return your bucket:

```bash
curl --user "<PublicKey>:<PrivateKey>" -X GET --digest \
     --header "Accept: application/json" \
     --header "Content-Type: application/json" \
     "https://cloud.mongodb.com/api/atlas/v1.0/groups/<ProjectId>/backup/exportBuckets?pretty=true"

{
  "links" : [ {
    "href" : "https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/6414908c207f4d22f4d8f232/backup/exportBuckets?pretty=true&pageNum=1&itemsPerPage=100",
    "rel" : "self"
  } ],
  "results" : [ {
    "_id" : "641dddc051ed5c6792399422",
    "bucketName" : "andrea-angiolillo-mongocli",
    "cloudProvider" : "AWS",
    "iamRoleId" : "641dd86151ed5c67923984f1"
  } ],
  "totalCount" : 1
}

```

## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-listexportbuckets)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/data-federation/config/config-aws-s3/)

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