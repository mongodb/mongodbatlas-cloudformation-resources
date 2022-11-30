# Auto Generate cfn template and generic GO Code

## How to generate Resource setup and Respective Go template for a new Resource

### Step 1:

Update mapping.json in schema folder with new resource details

eg:
    `{
        "resources" : [
            {
                "typeName" :"cloudProviderAccess",
                "openApiPath" : [
                "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes",
                "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName}",
                "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
                ]
            }
        ]
    }`


### step 2:
Run below command to generate schema file that is required for ``cfn generagte`` in the next step.

    make schema

### Step 3:
Run below command to generate Resource template which reads
/configs/<resource>-schema.json and /configs/<resource>-req.json

eg: to create search-index resource run below command.
search-indexes is the directory name for resource to create
indexes is the ResourceName that can be found in the end of the OpenAPI URL.
eg: /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes

    make create dirName=search-indexes resourceName=indexes


