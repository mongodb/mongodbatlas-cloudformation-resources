## MongoDB::Atlas::Trigger

### Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
- Trigger L1 CDK constructor
- Quickstart SageMaker




### Resources (and parameters for local tests) needed to manually QA:
All these resources need to be manually provided. Follow the steps below.
- Atlas project (PROJECT_ID)
- Atlas Cluster loaded with database and a collection (DB_NAME,COLLECTION_NAME)
- Realm App (APP_ID)
- Realm Service (SERVICE_ID)
- Atlas Function (FUNC_NAME,FUNC_ID)


## Manual QA:

For testing the resource Youll require the next properties:

export PROJECT_ID=
export APP_ID=
export DB_NAME=
export COLLECTION_NAME=
export FUNC_NAME=
export FUNC_ID=
export SERVICE_ID=

### Get the properties:
1. Create a new project
   - 1.a ATLAS UI: Go to your project/create a new project in Atlas. Note your PROJECT_ID from URL
   - 1.b CURL:
   ```
   curl --request POST \
   --url https://cloud.mongodb.com/api/atlas/v1.0/groups \
   --header 'Content-Type: application/json' \
   --data '{
   "name": "test-project",
   "orgId": "org_id",
   "regionUsageRestrictions": "NONE",
   "withDefaultAlertsSettings": true
   }'
   ```

2. Create a Cluster.
   - 2.a ATLAS UI: follow [create-new-cluster guide](https://www.mongodb.com/docs/atlas/tutorial/create-new-cluster/)
   - 2.b CURL:
   ```
   curl --request POST \
   --url https://cloud.mongodb.com/api/atlas/v1.0/groups \
   --header 'Content-Type: application/json' \
   --data '{
   "name": "test-project",
   "orgId": "org_id",
   "regionUsageRestrictions": "NONE",
   "withDefaultAlertsSettings": true
   }'
   ```
3. Create a database inside your cluster and a collection (load sample data for testing) and take note of the DB_NAME and COLLECTION_NAME
4. For the next steps we will have to Login to get our access token:
   - 4.a CURL: we execute the next curl and save the access_token result
   ```
   curl --request POST \
   --url https://realm.mongodb.com/api/admin/v3.0/auth/providers/mongodb-cloud/login \
   --header 'Accept: application/json' \
   --header 'Content-Type: application/json' \
   --data '{"username": "PUBLICKEY",
   "apiKey": "PRIVATEKEY"}'
   ```
4. Create an App in the project:
 - Click on ‘App Services’ tab -> click button ‘Create a new App’
 - Follow UI prompts to create app
 - Note your APP_ID from the URL (this is the identifier after ‘/apps/<id>’)
5) If testing with ‘Database Trigger’ type, a serviceId is also required. This can be found in Atlas UI -> App Services -> click on ‘Linked Data Sources’ in the left menu under ‘Manage’. Note SERVICE_ID from URL
6) After creating above resources you can also get above IDs from CLI by following:
```
Get an access token for your account (required by realm client):
curl --request POST \
  --header 'Content-Type: application/json' \
  --header 'Accept: application/json' \
  --data '{"username": "<AtlasPublicKey>", "apiKey": "<AtlasPrivateKey>"}' \
  https://realm.mongodb.com/api/admin/v3.0/auth/providers/mongodb-cloud/login

Retrieve AppId in your project:
curl \                                                                                                                 16.18.0 bazel 5.4.0
  --header 'Content-Type: application/json' \
  --header 'Accept: application/json' \
  --header "Authorization: Bearer ${ACCESS_TOKEN}" \
  https://realm.mongodb.com/api/admin/v3.0/groups/<ProjectId>/apps

Retrieve serviceId:
  curl \
  --header 'Content-Type: application/json' \
  --header 'Accept: application/json' \
  --header "Authorization: Bearer ${ACCESS_TOKEN}" \
https://realm.mongodb.com/api/admin/v3.0/groups/<ProjectId>/apps/<AppId>/services

```
### Steps to test:
1. Follow general [prerequisites](../../../TESTING.md#prerequisites) for testing CFN resources.
2. Update trigger.json under cfn-resources/examples/
 - Replace ProjectId, AppId, ServiceId (for DatabaseTrigger only) with your data.
 - If your EventProcessor is an Atlas Function, you will need to create a function and get its FUNC_ID and FUNC_NAME from Atlas UI under App Services in your Project.
 - Alternatively, your EventProcessor can be AWS Eventbridge. Add your AWS AccountId and Region in the example for this.
3. Set any additional required configuration options as per your needs.
4. To update test cases, update and run cfn-test-create-inputs.sh with required params from above.
5. Follow [general steps](../../../TESTING.md#steps) to test CFN resources.

### Success criteria when testing the resource
1. Trigger should be set up in your Atlas account as per configuration specified in the inputs/example.

   ![image](https://user-images.githubusercontent.com/122359335/227495196-59063691-c475-449c-b6b1-f206f4404715.png) 

2. General [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) should be satisfied.

## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/app-services/admin/api/v3/#tag/triggers)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/triggers/#service-functions-provide-server-side-logic)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
source <(${repo_root}/quickstart-mongodb-atlas/scripts/export-mongocli-config.py)
cd ${repo_root}/cfn-resources/trigger
./test/cluster.create-sample-cfn-request.sh YourProjectID YourClusterName > test.request.json 
echo "Sample request:"
cat test.request.json
cfn invoke resource CREATE test.request.json 
cfn invoke resource DELETE test.request.json 
cd -
```

Both CREATE & DELETE tests must pass.