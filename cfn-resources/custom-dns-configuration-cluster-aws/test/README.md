# MongoDB::Atlas::CustomDnsConfigurationClusterAws

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Custom DNS configuration L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
All resources are created as part of `cfn-testing-helper.sh`:

- Atlas project

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
1. AWS custom DNS should be enabled for the project. This can be validated via a [GET API](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-getawscustomdns) call as:
```
https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/<ATLAS_PROJECT_ID>/awsCustomDNS
```

![image](https://user-images.githubusercontent.com/122359335/227661815-d48398a9-aaa3-4978-9de4-736acab6ddf8.png)

2. Ensure general [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.


## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-toggleawscustomdns)


## Local Testing

The local tests are integrated with the AWS sam local and cfn invoke tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
cd ${repo_root}/cfn-resources/custom-dns-configuration-cluster-aws
./test/custom-dns-config-cluster-aws.create-sample-cfn-request.sh > test.request.json
echo "Sample request:"
cat test.request.json
cfn invoke CREATE test.request.json
cfn invoke DELETE test.request.json
```

Both CREATE & DELETE tests must pass.
