# MongoDB::Atlas::PrivateEndPointRegionalMode

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Private endpoint regional mode L1 CDK constructor
 - Quickstart Privateendpoint


## Prerequisites 
### Resources needed to run the manual QA
All resources are created as part of `cfn-testing-helper.sh`:

- Atlas project

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. Regionalized private endpoints setting should be enabled under Project Settings:

   ![image](https://user-images.githubusercontent.com/122359335/227656275-fd32b882-8b7d-4427-af6c-c4dc68fefd76.png)

2. Ensure general [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.


## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-returnregionalizedprivateendpointstatus)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/security-private-endpoint/#enable-regionalized-private-endpoints-1)


## Local Testing
The local tests are integrated with the AWS sam local and cfn invoke tooling features:

```
sam local start-lambda --skip-pull-image
```

then in another shell:
```bash
#https://www.mongodb.com/docs/mongocli/stable/configure/environment-variables/
cd ${repo_root}/cfn-resources/private-endpoint-regional-mode
./test/private-endpoint-regional-mode.create-sample-cfn-request.sh > test.request.json
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json
cfn invoke DELETE test.request.json
```

Both CREATE & DELETE tests must pass.